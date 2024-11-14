package cmd

import (
	"github.com/spf13/cobra"
)

var host, basicBase64 string

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func Init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&host, "host", "H", "127.0.0.1:5000", "the host to connect to")
	flags.StringVarP(&basicBase64, "user-password", "u", "", "basic auth user, get value: echo -n username:password | base64")

	rootCmd.AddCommand(catalog)

	TagInit()
	rootCmd.AddCommand(tag)

	ManifestInit()
	rootCmd.AddCommand(manifest)

	DeleteImageInit()
	rootCmd.AddCommand(deleteImage)
}
