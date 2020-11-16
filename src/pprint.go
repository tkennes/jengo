package jengo_src

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

func TableJobs(jobs [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Color", "Buildable", "URL", "Description"})

	for _, v := range jobs {
		table.Append(v)
	}
	table.Render()
}

func TableBuilds(jobs [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Number", "Time", "Result", "URL", "E_Duration", "Duration", "Description"})

	for _, v := range jobs {
		table.Append(v)
	}
	table.Render()
}

func YAMLJob(job JobResponse) {
	d, err := yaml.Marshal(&job)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf(string(d))
}

func YAMLBuild(build BuildResponse) {
	d, err := yaml.Marshal(&build)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf(string(d))
}
