package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthToken struct {
	AccessToken string `json:"access_token"`
	Token       string `json:"token"`
}

func Catalog(host, basicBase64 string) error {
	err := comm(commParam{
		Url:         fmt.Sprintf("http://%s/v2/_catalog?n=1000", host),
		BasicBase64: basicBase64,
		Handle: func(bytes []byte, response *http.Response) {
			type imageList struct {
				Registry []string `json:"repositories"`
			}
			list := &imageList{}
			err := json.Unmarshal(bytes, list)
			if err != nil {
				log.Printf("catalog response body unmarshal failed: %v\n", err)
				return
			}
			printLog(list.Registry)
			printLog("start print registry")
			for _, image := range list.Registry {
				fmt.Println(image)
			}
		},
	})
	return err
}
