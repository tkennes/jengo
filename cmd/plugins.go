package cmd

import (
	"github.com/spf13/cobra"

	src "github.com/tkennes/jengo/src"
)

var PluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Get plugins",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAMLPlugins(src.ListPlugins())
	},
}

func init() {
	RootCmd.AddCommand(PluginsCmd)
}
