package jengo_src

import (
	"encoding/json"
	"errors"
	"regexp"
	"sort"
	"strconv"

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

type ShortNameSorter []Plugin

func (s ShortNameSorter) Len() int           { return len(s) }
func (s ShortNameSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ShortNameSorter) Less(i, j int) bool { return s[i].ShortName < s[j].ShortName }

var (
	HeadersPlugins = []string{"active" , "enabled", "shortName", "version",  "url"}
	HeadersPlugin = []string{"active" , "backupVersion" , "bundled" , "deleted" , "detached" , "downgradable" , "enabled" , "hasUpdate" , "longName" , "minimumJavaVersion" , "pinned" , "requiredCoreVersion" , "shortName" , "supportsDynamicLoad" , "url" , "version"}
)

func GetAllPlugins(regex_filter string) (obj PluginsResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "plugins"})
	if err != nil {
		ErrorLog(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		ErrorLog(err)
	}
	sort.Sort(ShortNameSorter(obj.Plugins))

	// Regex type filter
	if regex_filter != "" {
		re := regexp.MustCompile(regex_filter)
		obj.Plugins = FilterRegexPluginNames(obj.Plugins, re)
	}
	return
}

func ListPlugins(regex_filter string) (res [][]string) {
	plugin_response := GetAllPlugins(regex_filter)

	for _, plugin := range plugin_response.Plugins {
		active := strconv.FormatBool(plugin.Active)
		enabled := strconv.FormatBool(plugin.Enabled)
		res = append(res,
			[]string{active,
				enabled,
				plugin.ShortName,
				plugin.Version,
				plugin.URL})
	}

	return
}

func GetPlugin(plugin_names []string) (plugins []Plugin) {
	plugin_response := GetAllPlugins("")
	for _, plugin_name := range plugin_names {
		p, err := FilterPlugin(plugin_response.Plugins, plugin_name)
		if err != nil {
			ErrorLog(err)
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

func FilterRegexPluginNames(plugins []Plugin, re *regexp.Regexp) (res []Plugin) {
	for _, plugin := range plugins {
		if len(re.FindStringSubmatch(plugin.ShortName)) > 0 || len(re.FindStringSubmatch(plugin.LongName)) > 0 {
			res = append(res, plugin)
		}
	}
	return
}
