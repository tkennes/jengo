package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var BuildRootCmd = &cobra.Command{
	Use:   "build",
	Short: "The entrypoint for the build interface",
}

var GetBuild = &cobra.Command{
	Use:   "get",
	Short: "Get build",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString("job")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if job == "" {
			fmt.Println("No job was submitted.\n\nUse --job or -j")
		} else 	if len(args) == 0 {
			fmt.Println("No Build was submitted.\n\nEnter the corresponding builds you want to fetch")
		} else {
			for _, build := range args {
				src.YAML(src.GetBuild(job, build))
			}
		}
	},
}

var ListBuilds = &cobra.Command{
	Use:   "list",
	Short: "List all builds of jobs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No job was submitted.\n\nEnter the corresponding names for the jobs whose builds you want to fetch")
		}
		for _, job := range args {
			src.Table(src.ListBuilds(job), src.HeadersBuilds)
		}
	},
}

func init() {
	// Toplevel
	RootCmd.AddCommand(BuildRootCmd)

	// Sublevel
	var job string
	BuildRootCmd.AddCommand(GetBuild)
	GetBuild.PersistentFlags().StringVarP(&job, "job", "j", "", "Job of interest")

	BuildRootCmd.AddCommand(ListBuilds)
}
