package cmd

import (
	"log"

	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Get builds",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString("job")
		if err != nil {
			log.Fatal(err)
		}

		for _, build := range args {
			src.YAMLBuild(src.GetBuild(job, build))
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildCmd)
	BuildCmd.Flags().String("job", "", "Job of interest")
}
