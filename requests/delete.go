package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DeleteImage(host, basicBase64, repo, tag string) error {
	// get sha256
	sha256, err := GetManifest(host, basicBase64, repo, tag, false)
	if err != nil {
		return err
	}
	printLog(fmt.Sprintf("%s:%s SHA256: %s", repo, tag, sha256))
	url := fmt.Sprintf("http://%s/v2/%s/manifests/%s", host, repo, sha256)
	printLog("delete image url: ", url)
	realm, service, scope, err := getAuthInfo(url, http.MethodDelete, nil, map[string][]string{
		"Accept": {"application/vnd.docker.distribution.manifest.v2+json"},
	})
	if err != nil {
		printLog("get auth info failed: ", err)
		return err
	}
	token, err := getToken(basicBase64, realm, service, scope)
	if err != nil {
		printLog(err)
		return err
	}
	printLog("delete image token: ", token)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("%s new request err: %s", url, err)
		return err
	}

	request.Header.Add("Authorization", "Bearer "+token)
	// request.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("%s do request err: %s", url, err)
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		b, _ := io.ReadAll(resp.Body)
		if b != nil {
			log.Println("delete image body: ", string(b))
		}
		log.Printf("%s status code %d", url, resp.StatusCode)
		return err
	}
	fmt.Printf("%s:%s delete success!\n", repo, tag)
	return nil
}
