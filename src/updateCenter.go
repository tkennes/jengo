package jengo_src

import (
	"encoding/json"
	"time"
	"regexp"
	"sort"
	"strconv"
)

type UpdateCenterResponse struct {
	Class              string       `yaml:"_class"`
	Availables         []Availables `yaml:"availables"`
	ConnectionCheckURL string       `yaml:"connectionCheckUrl"`
	DataTimestamp      int64        `yaml:"dataTimestamp"`
	HasUpdates         bool         `yaml:"hasUpdates"`
	ID                 string       `yaml:"id"`
	Updates            []Updates    `yaml:"updates"`
	URL                string       `yaml:"url"`
}

type Availables struct {
	Name                           string               `yaml:"name"`
	SourceID                       string               `yaml:"sourceId"`
	URL                            string               `yaml:"url"`
	Version                        string               `yaml:"version"`
	Categories                     []interface{}        `yaml:"categories"`
	Compatible                     bool                 `yaml:"compatible"`
	CompatibleSinceVersion         interface{}          `yaml:"compatibleSinceVersion"`
	CompatibleWithInstalledVersion bool                 `yaml:"compatibleWithInstalledVersion"`
	Excerpt                        string               `yaml:"excerpt"`
	Installed                      interface{}          `yaml:"installed"`
	MinimumJavaVersion             string               `yaml:"minimumJavaVersion"`
	NeededDependencies             []interface{}        `yaml:"neededDependencies"`
	ReleaseTimestamp               int64                `yaml:"releaseTimestamp"`
	RequiredCore                   string               `yaml:"requiredCore"`
	Title                          string               `yaml:"title"`
	Wiki                           string               `yaml:"wiki"`
}

type Updates struct {
	Name                           string               `yaml:"name"`
	SourceID                       string               `yaml:"sourceId"`
	URL                            string               `yaml:"url"`
	Version                        string               `yaml:"version"`
	Categories                     []string             `yaml:"categories"`
	Compatible                     bool                 `yaml:"compatible"`
	CompatibleSinceVersion         interface{}          `yaml:"compatibleSinceVersion"`
	CompatibleWithInstalledVersion bool                 `yaml:"compatibleWithInstalledVersion"`
	Excerpt                        string               `yaml:"excerpt"`
	Installed                      Installed            `yaml:"installed"`
	MinimumJavaVersion             string               `yaml:"minimumJavaVersion"`
	NeededDependencies             []interface{}        `yaml:"neededDependencies"`
	ReleaseTimestamp               int64                `yaml:"releaseTimestamp"`
	RequiredCore                   string               `yaml:"requiredCore"`
	Title                          string               `yaml:"title"`
	Wiki                           string               `yaml:"wiki"`
}

type Installed struct {
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

var (
	HeadersAvailablePlugins = []string{"name", "current-version", "new-version", "compatible", "release-timestamp"}
)

type AvailableNameSorter []Availables
func (a AvailableNameSorter) Len() int           { return len(a) }
func (a AvailableNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AvailableNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

type UpdateNameSorter []Updates
func (u UpdateNameSorter) Len() int           { return len(u) }
func (u UpdateNameSorter) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u UpdateNameSorter) Less(i, j int) bool { return u[i].Name < u[j].Name }


func GetAllUpdateCenter() (obj UpdateCenterResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "update-center"})
	if err != nil {
		ErrorLog(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		ErrorLog(err)
	}

	return
}

func GetAllUpdateablePlugins(regex_filter string) (res [][]string) {
	update_center := GetAllUpdateCenter()

	// Sort availables alphabetically
	sort.Sort(UpdateNameSorter(update_center.Updates))

	current_plugins := getUpdatePlugins(update_center.Updates)

	for i, updates := range update_center.Updates {
		releaseTimestamp := time.Unix(updates.ReleaseTimestamp/1000, 0).Format("2006-02-01 15:04:05")
		compatible := strconv.FormatBool(updates.Compatible)
		res = append(res,
			[]string{updates.Name,
				current_plugins[i].Version,
				updates.Version,
				compatible,
				releaseTimestamp})
	}

	return

}

func GetAllAvailablePlugins(regex_filter string) (res [][]string) {
	update_center := GetAllUpdateCenter()

	// Sort availables alphabetically
	sort.Sort(AvailableNameSorter(update_center.Availables))

	filtered_availables := update_center.Availables
	// Regex type filter
	if regex_filter != "" {
		re := regexp.MustCompile(regex_filter)
		filtered_availables = FilterRegexNames(filtered_availables, re)
	}

	for _, available := range filtered_availables {
		releaseTimestamp := time.Unix(available.ReleaseTimestamp/1000, 0).Format("2006-02-01 15:04:05")
		compatible := strconv.FormatBool(available.Compatible)
		res = append(res,
			[]string{available.Name,
				available.Version,
				compatible,
				releaseTimestamp})
	}

	return
}


func FilterRegexNames(availables []Availables, re *regexp.Regexp) (res []Availables) {
	for _, av := range availables {
		if len(re.FindStringSubmatch(av.Name)) > 0 {
			res = append(res, av)
		}
	}
	return
}

func getUpdatePlugins(updates []Updates) []Plugin {
	names := []string{}
	for _, update := range updates {
		names = append(names, update.Name)
	}

	return GetPlugin(names)
}
