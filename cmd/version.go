package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// StoryArc is a path in a story
type Config struct {
	Version string `json:"version"`
}


var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get Nodes",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := parse_json(getFileBytes("./config/jengo.json"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(config.Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file %s", fileName)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", fileName)
	}

	return buf.Bytes()
}

func parse_json(jsondata []byte) (config Config, err error) {
	err = json.Unmarshal(jsondata, &config)

	return
}
