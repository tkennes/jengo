package jengo_src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Context struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

type Config struct {
	CurrentContext string `yaml:"current_context"`
	Contexts        []Context `yaml:"contexts"`
}

const (
	ConfigFile = ".jenkins.yaml"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Getters
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetConfig(configfile_location string) Config {
	var configData []byte
	if configfile_location == "" {
		home, _ := homedir.Dir()
		configData = GetFileBytes(filepath.Join(home, ConfigFile))
	} else {
		configData = GetFileBytes(configfile_location)
	}

	var config Config
	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		ErrorLog(err)
	}

	return config
}

func GetCurrentContext(configfile_location string) (context Context) {
	config := GetConfig(configfile_location)
	context = config.GetCurrent()
	return
}

func GetBaseURL(configfile_location string) string {
	jenkins := GetCurrentContext(configfile_location)
	return jenkins.URL
}

func GetName(configfile_location string) string {
	jenkins := GetCurrentContext(configfile_location)
	return jenkins.Name
}

func GetUserName(configfile_location string) string {
	jenkins := GetCurrentContext(configfile_location)
	return jenkins.Username
}

func GetToken(configfile_location string) string {
	jenkins := GetCurrentContext(configfile_location)
	return jenkins.Token
}

func (config Config) GetCurrent() Context {
	for _, c := range config.Contexts {
		if c.Name == config.CurrentContext {
			return c
		}
	}

	// If not return, the current context does not exist
	ErrorLog(errors.New("Current context does not or no longer exist. Create it, or change the context"))
	return Context{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Show-ers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAndShowConfig(configfile_location string) (config Config) {
	config = GetConfig(configfile_location)
	config.hideTokens()
	return
}

func GetAndShowContexts(configfile_location string) (contexts []Context) {
	contexts = GetAndShowConfig(configfile_location).Contexts
	return
}

func GetAndShowContext(configfile_location string, context_name string) Context {
	config := GetConfig(configfile_location)
	config.hideTokens()
	contexts := config.Contexts
	for _, context := range contexts {
		if context.Name == context_name {
			return context
		}
	}
	ErrorLog(errors.New(fmt.Sprintf("Could not find context: %s", context_name)))
	return Context{}
}

func GetAndShowCurrentContext(configfile_location string) (context Context) {
	context = GetCurrentContext(configfile_location)
	context = hideToken(context)
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Updaters
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func DoAddContext(configfile_location string, context_name string, token string, url string, username string) {
	config := GetConfig(configfile_location)
	config = AddContext(config, context_name, token, url, username)
	config.saveConfig(configfile_location)
}

func DoRemoveContext(configfile_location string, context_name string) {
	config := GetConfig(configfile_location)
	config = RemoveContext(config, context_name)
	config.saveConfig(configfile_location)
}

func DoUpdateContext(configfile_location string, context_name string, new_context_name string, token string, url string, username string) {
	config := GetConfig(configfile_location)
	config = UpdateContext(config, context_name, new_context_name, token, url, username)
	config.saveConfig(configfile_location)
}

func DoUpdateManualContext(configfile_location string, context_name string) {
	config := GetConfig(configfile_location)
	config = ManualUpdateContext(config, context_name)
	config.saveConfig(configfile_location)
}

func DoUpdateCurrentContext(configfile_location string, new_context_name string) {
	config := GetConfig(configfile_location)
	config = UpdateCurrentContext(config, new_context_name)
	config.saveConfig(configfile_location)
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Config Handling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func AddContext(config Config, context_name string, token string, url string, username string) Config {
	new_context := PerformCreateContext(context_name, token, url, username)
	config.Contexts = append(config.Contexts, new_context)
	return config
}

func RemoveContext(config Config, context_name string) Config {
	var contexts []Context
	for _, context := range config.Contexts {
		if context.Name != context_name {
			contexts = append(contexts, context)
		}
	}
	config.Contexts = contexts
	return config
}

func UpdateCurrentContext(config Config, new_context string) Config {
	config.CurrentContext = new_context
	return config
}

func UpdateContext(config Config, context_name string, new_context_name string, token string, url string, username string) Config {
	found := false
	for i, _ := range config.Contexts {
		if config.Contexts[i].Name == context_name {
			config.Contexts[i] = PerformUpdateContext(config.Contexts[i], new_context_name, token, url, username)
			found = true
			return config
		}
	}
	if !found {
		ErrorLog(errors.New(fmt.Sprintf("Did not found context: %s", context_name)))
	}
	return config
}

func ManualUpdateContext(config Config, context_name string) Config {
	found := false
	for i, _ := range config.Contexts {
		if config.Contexts[i].Name == context_name {
			config.Contexts[i] = PerformManualUpdateContext(config.Contexts[i])
			found = true
			return config
		}
	}
	if !found {
		ErrorLog(errors.New(fmt.Sprintf("Did not found context: %s", context_name)))
	}
	return config
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Internal helpers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (config Config) hideTokens() Config {
	for i, _ := range config.Contexts {
		config.Contexts[i] = hideToken(config.Contexts[i])
	}
	return config
}

func hideToken(context Context) Context{
	context.Token = "***"
	return context
}

func (config Config) saveConfig(configfile_location string){
	configData, err := yaml.Marshal(&config)
	if err != nil {
		ErrorLog(errors.New("Error in marshalling config"))
	}
	if configfile_location == "" {
		home, _ := homedir.Dir()
		configfile_location = filepath.Join(home, ConfigFile)
	}
	err = ioutil.WriteFile(configfile_location, configData, 0644)
	if err != nil {
		Info.Printf("Could not open file %s", configfile_location)
		os.Exit(1)
	}
}

func PerformUpdateContext(context Context, context_name string, token string, url string, username string) Context {
	if context_name != "" {
		context = updateName(context, context_name)
	}
	if url != "" {
		context = updateUrl(context, url)
	}
	if username != "" {
		context = updateUsername(context, username)
	}
	if token != "" {
		context = updateToken(context, token)
	}
	return context
}

func PerformManualUpdateContext(context Context) Context {
	context = updateName(context, "")
	context = updateUrl(context, "")
	context = updateUsername(context, "")
	context = updateToken(context, "")

	return context
}

func PerformCreateContext(context_name string, token string, url string, username string) Context {
	var context Context

	context = updateName(context, context_name)
	context = updateUrl(context, url)
	context = updateUsername(context, username)
	context = updateToken(context, token)

	return context
}

func updateName(context Context, name string) Context {
	var answer string
	if name == "" {
		fmt.Println("Enter context-name")
		fmt.Scanln(&answer)
		context.Name = answer
	} else {
		context.Name = name
	}
	return context
}

func updateToken(context Context, token string) Context {
	var answer string
	if token == "" {
		fmt.Println("Enter token:")
		fmt.Scanln(&answer)
			context.Token = answer
	} else {
		context.Token = token
	}
	return context
}

func updateUrl(context Context, url string) Context {
	var answer string
	if url == "" {
		fmt.Println("Enter jenkins-server url (default: http://localhost:8000):")
		fmt.Scanln(&answer)
		if answer == "" {
			context.URL = "http://localhost:8000"
		} else {
			context.URL = answer
		}
	} else {
		context.URL = url
	}
	return context
}

func updateUsername(context Context, username string) Context {
	var answer string
	if username == "" {
		fmt.Println("Enter username (default: admin):")
		fmt.Scanln(&answer)
		if answer == "" {
			context.Username = "admin"
		} else {
			context.Username = answer
		}
	} else {
		context.Username = username
	}
	return context
}
