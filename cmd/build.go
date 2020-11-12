package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tomkennes/jengo/common"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Get builds",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString("job")
		if err != nil {
			log.Fatal(err)
		}

		for i, build := range args {
			if i > 0 {
				fmt.Println("---\n")
			}
			common.YAMLBuild(common.GetBuild(job, build))
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildCmd)
	BuildCmd.Flags().String("job", "", "Job of interest")
}
