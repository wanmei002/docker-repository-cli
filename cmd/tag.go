package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wanmei002/docker-repository-cli/requests"
)

var repo string

var tag = &cobra.Command{
	Use:   "tag",
	Short: "image tags",
	Run: func(cmd *cobra.Command, args []string) {
		if repo == "" {
			fmt.Println("repo is required")
			return
		}
		err := requests.TagList(host, basicBase64, repo)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func TagInit() {
	flags := tag.Flags()
	flags.StringVarP(&repo, "repo", "r", "", "image repo")
}
