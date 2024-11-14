package requests

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type commParam struct {
	Url         string
	BasicBase64 string
	Handle      func([]byte, *http.Response)
	Header      http.Header
}

func comm(param commParam) error {
	urlI, err := url.Parse(param.Url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if urlI.Host == "" {
		fmt.Println(urlI)
		return errors.New("arg host is empty")
	}
	if param.BasicBase64 == "" {
		return errors.New("arg basic base64 is empty")
	}
	realm, service, scope, err := getAuthInfo(param.Url, http.MethodGet, nil, param.Header)
	if err != nil {
		return err
	}

	token, err := getToken(param.BasicBase64, realm, service, scope)
	if err != nil {
		log.Printf("getToken err: %v\n", err)
		return err
	}
	printLog("url: ", param.Url)

	request, err := http.NewRequest(http.MethodGet, param.Url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed; url:%s, err: %v\n", param.Url, err)
		return err
	}

	request.Header.Add("Authorization", "Bearer "+token)

	for k, v := range param.Header {
		for _, vv := range v {
			request.Header.Add(k, vv)
		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("http.Do failed; url:%s, err: %v\n", param.Url, err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("resp.StatusCode != http.StatusOK; resp:%+v \n", resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("io.ReadAll failed; url:%s, err: %v\n", param.Url, err)
		return err
	}

	param.Handle(body, resp)
	return nil
}

func getAuthInfo(url, httpMethod string, body io.Reader, header http.Header) (string, string, string, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		log.Printf("http.NewRequest url:%s err: %v\n", url, err)
		return "", "", "", err
	}
	for k, v := range header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http.Do url:%s err: %v\n", url, err)
		return "", "", "", err
	}
	defer resp.Body.Close()
	authReqInfo := resp.Header.Get("Www-Authenticate")
	printLog(fmt.Sprintf("auth request info: %s", authReqInfo))

	authReqInfoList := strings.Split(strings.TrimSpace(authReqInfo), " ")
	if len(authReqInfoList) != 2 {
		log.Printf("authReq format error\n")
		return "", "", "", errors.New("authReq format error")
	}
	reqParam := strings.Split(authReqInfoList[1], ",")
	var realm, service, scope string
	for _, param := range reqParam {
		tmp := strings.Split(param, "=")
		if len(tmp) != 2 {
			log.Printf("param format error %v\n", param)
			continue
		}
		tmpVal := strings.Trim(tmp[1], "\"")
		switch tmp[0] {
		case "realm":
			realm = tmpVal
		case "service":
			service = tmpVal
		case "scope":
			scope = tmpVal
		}
	}

	return realm, service, scope, nil
}
