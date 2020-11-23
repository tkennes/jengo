package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var NodeRootCmd = &cobra.Command{
	Use:   "node",
	Short: "The entrypoint for the node interface",
}

var GetNodeCmd = &cobra.Command{
	Use:   "get",
	Short: "Get node info",
	Run: func(cmd *cobra.Command, args []string) {
		for _, node := range args {
			src.YAML(src.GetNode(node))
		}
	},
}

var ListNodesCmd = &cobra.Command{
	Use:   "list",
	Short: "Get Nodes",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAML(src.ListNodes())
	},
}

func init() {
	// Toplevel
	RootCmd.AddCommand(NodeRootCmd)

	// Sublevel
	NodeRootCmd.AddCommand(GetNodeCmd)
	NodeRootCmd.AddCommand(ListNodesCmd)
}
