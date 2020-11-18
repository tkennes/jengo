package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Get node info",
	Run: func(cmd *cobra.Command, args []string) {
		for _, node := range args {
			src.YAMLNode(src.GetNode(node))
		}
	},
}

func init() {
	RootCmd.AddCommand(NodeCmd)
}
