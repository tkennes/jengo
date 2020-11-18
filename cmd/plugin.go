package cmd

import (
	"github.com/spf13/cobra"

	src "github.com/tkennes/jengo/src"
)

var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Get plugin",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAMLPlugin(src.GetPlugin(args))
	},
}

func init() {
	RootCmd.AddCommand(PluginCmd)
}
