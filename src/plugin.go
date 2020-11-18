package jengo_src

import (
	"encoding/json"
	"errors"

)

type PluginsResponse struct {
	Plugins []Plugin `yaml:"plugins"`
}

type Plugin struct {
	Active              bool           `yaml:"active"`
	BackupVersion       interface{}    `yaml:"backupVersion"`
	Bundled             bool           `yaml:"bundled"`
	Deleted             bool           `yaml:"deleted"`
	Detached            bool           `yaml:"detached"`
	Downgradable        bool           `yaml:"downgradable"`
	Enabled             bool           `yaml:"enabled"`
	HasUpdate           bool           `yaml:"hasUpdate"`
	LongName            string         `yaml:"longName"`
	MinimumJavaVersion  string         `yaml:"minimumJavaVersion"`
	Pinned              bool           `yaml:"pinned"`
	RequiredCoreVersion string         `yaml:"requiredCoreVersion"`
	ShortName           string         `yaml:"shortName"`
	SupportsDynamicLoad string         `yaml:"supportsDynamicLoad"`
	URL                 string         `yaml:"url"`
	Version             string         `yaml:"version"`
}

func ListPlugins() (obj PluginsResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "plugins"})
	if err != nil {
		Error.Println(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		Error.Println(err)
	}
	return
}

func GetPlugin(plugin_names []string) (plugins []Plugin) {
	plugin_response := ListPlugins()
	for _, plugin_name := range plugin_names {
		p, err := FilterPlugin(plugin_response.Plugins, plugin_name)
		if err != nil {
			Error.Println(err)
		}
		plugins = append(plugins, p)
	}
	return
}

func FilterPlugin(plugins []Plugin, plugin_name string) (Plugin, error) {
	for _, plugin := range plugins {
		if plugin.ShortName == plugin_name || plugin.LongName == plugin_name {
			return plugin, nil
		}
	}
	p := Plugin{}
	return p, errors.New("Plugin does not exist")
}
