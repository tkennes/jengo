package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomkennes/jengo/common"
)

var BuildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "Get builds",
	Run: func(cmd *cobra.Command, args []string) {
		for _, job := range args {
			common.TableBuilds(common.ListBuilds(job))
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildsCmd)
}
