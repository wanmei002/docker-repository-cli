package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanmei002/docker-repository-cli/requests"
	"log"
)

var catalog = &cobra.Command{
	Use:   "catalog",
	Short: "Get all catalog images",
	Run: func(cmd *cobra.Command, args []string) {
		err := requests.Catalog(host, basicBase64)
		if err != nil {
			log.Fatal(err)
		}
	},
}
