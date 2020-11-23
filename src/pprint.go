package jengo_src

import (
	"os"
	"log"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)


func Table(jobs [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	// ALIGN LEFT
	table.SetAlignment(3)

	for _, v := range jobs {
		table.Append(v)
	}
	table.Render()
}

func YAML(obj interface{}) {
	d, err := yaml.Marshal(&obj)
	if err != nil {
		ErrorLog(err)
	}
	// Log without datetime prefix
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.Printf(string(d))
}
