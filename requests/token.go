package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getToken(basicBase64, realm, service, scope string) (string, error) {
	authUrl := fmt.Sprintf("%s?service=%s&scope=%s", realm, service, scope)
	printLog(fmt.Sprintf("authUrl: %s", authUrl))
	authRequest, err := http.NewRequest(http.MethodGet, authUrl, nil)
	if err != nil {
		log.Printf("create auth request err: %v\n", err)
		return "", err
	}
	authRequest.Header.Add("Authorization", "Basic "+basicBase64)
	printLog(fmt.Sprintf("auth request header:%v", authRequest.Header))
	authResponse, err := http.DefaultClient.Do(authRequest)
	if err != nil {
		log.Printf("auth request err: %v\n", err)
		return "", err
	}
	defer authResponse.Body.Close()
	body, err := io.ReadAll(authResponse.Body)
	if err != nil {
		log.Printf("read auth response body err: %v\n", err)
		return "", err
	}
	printLog(fmt.Sprintf("auth response body: %s", string(body)))
	if authResponse.StatusCode != http.StatusOK {
		log.Printf("auth request status code: %d\n", authResponse.StatusCode)
		return "", errors.New("auth request status code error")
	}
	token := &AuthToken{}
	err = json.Unmarshal(body, token)
	if err != nil {
		log.Printf("auth json unmarshal err: %v\n", err)
		return "", err
	}
	return token.Token, nil
}
