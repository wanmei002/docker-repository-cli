package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wanmei002/docker-repository-cli/requests"
)

var (
	manifestRepo, manifestTag string
)

var manifest = &cobra.Command{
	Use:   "manifest",
	Short: "Get image manifest",
	Run: func(cmd *cobra.Command, args []string) {
		if manifestRepo == "" || manifestTag == "" {
			fmt.Println("manifest repo and manifest tag are required")
			return
		}
		_, err := requests.GetManifest(host, basicBase64, manifestRepo, manifestTag, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func ManifestInit() {
	flags := manifest.Flags()
	flags.StringVarP(&manifestRepo, "repo", "r", "", "repo")
	flags.StringVarP(&manifestTag, "tag", "t", "", "tag")
}
