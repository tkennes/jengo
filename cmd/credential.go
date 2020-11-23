package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var CredentialCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Get credentials",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAML(src.ListCredentials())
	},
}

func init() {
	RootCmd.AddCommand(CredentialCmd)
}
