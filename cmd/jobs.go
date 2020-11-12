package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomkennes/jengo/common"
)

var JobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Lists all the jobs",
	Run: func(cmd *cobra.Command, args []string) {
		common.TableJobs(common.ListJobs())
	},
}

func init() {
	RootCmd.AddCommand(JobsCmd)
}
