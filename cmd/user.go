package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Get user",
	Run: func(cmd *cobra.Command, args []string) {
		for _, user := range args {
			src.YAML(src.GetUser(user))
		}
	},
}

func init() {
	RootCmd.AddCommand(UserCmd)
}
