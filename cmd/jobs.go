package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tomkennes/jengo/src"
)

var JobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Lists all the jobs",
	Run: func(cmd *cobra.Command, args []string) {
		src.TableJobs(src.ListJobs())
	},
}

func init() {
	RootCmd.AddCommand(JobsCmd)
}
