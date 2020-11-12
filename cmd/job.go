package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomkennes/jengo/common"
)

var JobCmd = &cobra.Command{
	Use:   "job",
	Short: "Get job info",
	Run: func(cmd *cobra.Command, args []string) {
		for i, job := range args {
			if i > 0 {
				fmt.Println("---\n")
			}
			common.YAMLJob(common.GetJob(job))
		}
	},
}

func init() {
	RootCmd.AddCommand(JobCmd)
}
