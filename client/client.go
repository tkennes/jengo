package client

import (
	"bytes"
	"flag"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type JenkinsClient struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

type YAMLConfig struct {
	current_context string `yaml:"current"`
	Contexts        []JenkinsClient
}

//  Neat way to group input variables
var (
	configFile      = flag.String("conf", ".jenkins.yaml", "The file containing shortened paths to URL's")
	current_context = flag.String("c", "localhost", "The context currently in use")
)

func init() {
	flag.Parse()
}

func GetClient() JenkinsClient {
	home, _ := homedir.Dir()
	jenkins, err := parseYAML(getFileBytes(filepath.Join(home, *configFile)))

	if err != nil {
		panic(err)
	}

	return jenkins
}

func GetBaseURL() string {
	jenkins := GetClient()
	return jenkins.URL
}

func CreateRequest(METHOD string, URL string) *http.Request {
	jenkins := GetClient()
	req, err := http.NewRequest(METHOD, jenkins.URL+URL, nil)
	req.SetBasicAuth(jenkins.Username, jenkins.Token)
	if err != nil {
		panic(err)
	}
	return req
}

func buildMap(client JenkinsClient) (builtMap map[string]string) {
	builtMap = make(map[string]string)
	builtMap["url"] = client.URL
	builtMap["username"] = client.Username
	builtMap["token"] = client.Token
	return
}

func parseYAML(yamlData []byte) (client JenkinsClient, err error) {
	// marshalling or marshaling is the process of transforming the memory representation of
	// an object to a data format suitable for storage or transmission
	var y YAMLConfig
	err = yaml.Unmarshal(yamlData, &y)
	client = y.GetCurrent()
	return
}

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file %s", fileName)
	}

	// Package bytes implements functions for the manipulation of byte slices. It is analogous to the facilities of the strings package.
	// Read in bytes in a single buffer, and convert them into a string when we are done.
	//
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", fileName)
	}

	return buf.Bytes()
}

func (config YAMLConfig) GetCurrent() JenkinsClient {
	for _, c := range config.Contexts {
		if c.Name == *current_context {
			return c
		}
	}
	panic("Context not found in YAML")
}
