package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tomkennes/jengo/src"
)

var BuildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "Get builds",
	Run: func(cmd *cobra.Command, args []string) {
		for _, job := range args {
			src.TableBuilds(src.ListBuilds(job))
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildsCmd)
}
