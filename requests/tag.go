package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TagList(host, basicBase64, repo string) error {
	url := fmt.Sprintf("http://%s/v2/%s/tags/list?n=10000", host, repo)
	err := comm(commParam{
		Url:         url,
		BasicBase64: basicBase64,
		Handle: func(bytes []byte, response *http.Response) {
			type TagList struct {
				Name string   `json:"name"`
				Tags []string `json:"tags"`
			}
			printLog(string(bytes))
			list := &TagList{}
			err := json.Unmarshal(bytes, list)
			if err != nil {
				log.Printf("request failed; url:%s; err:%v", url, err)
				fmt.Println(err)
				return
			}
			for i, tag := range list.Tags {
				fmt.Print(tag, ", ")
				if i != 0 && i%6 == 0 {
					fmt.Print("\n")
				}
			}
			fmt.Print("\n")
		},
	})

	return err
}
