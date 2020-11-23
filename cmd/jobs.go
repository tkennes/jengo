package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)


var JobRootCmd = &cobra.Command{
	Use:   "job",
	Short: "The entrypoint for the job interface",
}

var GetJobCmd = &cobra.Command{
	Use:   "get",
	Short: "Get job info",
	Run: func(cmd *cobra.Command, args []string) {
		for _, job := range args {
			src.YAML(src.GetJob(job))
		}
	},
}

var ListJobsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the jobs",
	Run: func(cmd *cobra.Command, args []string) {
		header := []string{"Name", "Color", "Buildable", "URL", "Description"}
		src.Table(src.ListJobs(), header)
	},
}

func init() {
	// Toplevel
	RootCmd.AddCommand(JobRootCmd)

	// Sublevel
	JobRootCmd.AddCommand(GetJobCmd)
	JobRootCmd.AddCommand(ListJobsCmd)
}
