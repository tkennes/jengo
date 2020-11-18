package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var NodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Get Nodes",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAMLNodes(src.ListNodes())
	},
}

func init() {
	RootCmd.AddCommand(NodesCmd)
}
