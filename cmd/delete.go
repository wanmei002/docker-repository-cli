package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wanmei002/docker-repository-cli/requests"
)

var (
	deleteImageRepo, deleteImageTag string
)

var deleteImage = &cobra.Command{
	Use:   "delete-image",
	Short: "Delete image",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteImageRepo == "" || deleteImageTag == "" {
			fmt.Println("repo or tag is required")
			return
		}
		err := requests.DeleteImage(host, basicBase64, deleteImageRepo, deleteImageTag)
		if err != nil {
			fmt.Println(err)
		}
		return
	},
}

func DeleteImageInit() {
	flags := deleteImage.Flags()
	flags.StringVarP(&deleteImageRepo, "repo", "r", "", "docker repository name")
	flags.StringVarP(&deleteImageTag, "tag", "t", "", "docker repository tag name")
}
