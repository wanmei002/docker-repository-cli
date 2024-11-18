package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetManifest(host, basicBase64, repo, tag string, isPrint bool) (string, error) {
	url := fmt.Sprintf("http://%s/v2/%s/manifests/%s", host, repo, tag)
	sha256 := ""
	err := comm(commParam{
		Url:         url,
		BasicBase64: basicBase64,
		Handle: func(strB []byte, resp *http.Response) {
			if isPrint {
				var str bytes.Buffer
				_ = json.Indent(&str, strB, "", "\t")
				fmt.Println(str.String())
			}
			sha256 = resp.Header.Get("Etag")
			printLog(fmt.Sprintf("manifest request header: %#v", resp.Request.Header))
			printLog(fmt.Sprintf("manifest header: %#v", resp.Header))
		},
		Header: map[string][]string{
			"Accept": {
				"application/vnd.docker.distribution.manifest.v2+json",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	return strings.Trim(sha256, "\""), err
}
