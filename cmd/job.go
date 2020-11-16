package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tomkennes/jengo/src"
)

var JobCmd = &cobra.Command{
	Use:   "job",
	Short: "Get job info",
	Run: func(cmd *cobra.Command, args []string) {
		for _, job := range args {
			src.YAMLJob(src.GetJob(job))
		}
	},
}

func init() {
	RootCmd.AddCommand(JobCmd)
}
