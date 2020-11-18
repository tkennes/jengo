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
		Error.Println(err)
	}
	Info.Println(string(d))
}

func YAMLBuild(build BuildResponse) {
	d, err := yaml.Marshal(&build)
	if err != nil {
		Error.Println(err)
	}
	log.Printf(string(d))
}

func YAMLNodes(nodes NodesResponse) {
	d, err := yaml.Marshal(&nodes)
	if err != nil {
		Error.Println(err)
	}
	log.Printf(string(d))
}

func YAMLNode(node Computer) {
	d, err := yaml.Marshal(&node)
	if err != nil {
		Error.Println(err)
	}
	log.Printf(string(d))
}

func YAMLPlugins(plugins PluginsResponse) {
	d, err := yaml.Marshal(&plugins)
	if err != nil {
		Error.Println(err)
	}
	log.Printf(string(d))
}

func YAMLPlugin(plugin []Plugin) {
	d, err := yaml.Marshal(&plugin)
	if err != nil {
		Error.Println(err)
	}
	log.Printf(string(d))
}
