package jengo_src

import (
	"encoding/json"
	"fmt"
)

type NodesResponse struct {
	BusyExecutors  int        `yaml:"busyExecutors"`
	Computer       []Computer `yaml:"computer"`
	DisplayName    string     `yaml:"displayName"`
	TotalExecutors int        `yaml:"totalExecutors"`
}
type AssignedLabels struct {
	Name string `yaml:"name"`
}
type Executors struct {
}
type LoadStatistics struct {
	Class string `yaml:"_class"`
}
type HudsonNodeMonitorsSwapSpaceMonitor struct {
	Class                   string `yaml:"_class"`
	AvailablePhysicalMemory int    `yaml:"availablePhysicalMemory"`
	AvailableSwapSpace      int    `yaml:"availableSwapSpace"`
	TotalPhysicalMemory     int    `yaml:"totalPhysicalMemory"`
	TotalSwapSpace          int    `yaml:"totalSwapSpace"`
}
type HudsonNodeMonitorsTemporarySpaceMonitor struct {
	Class     string `yaml:"_class"`
	Timestamp int64  `yaml:"timestamp"`
	Path      string `yaml:"path"`
	Size      int64  `yaml:"size"`
}
type HudsonNodeMonitorsDiskSpaceMonitor struct {
	Class     string `yaml:"_class"`
	Timestamp int64  `yaml:"timestamp"`
	Path      string `yaml:"path"`
	Size      int64  `yaml:"size"`
}
type HudsonNodeMonitorsResponseTimeMonitor struct {
	Class     string `yaml:"_class"`
	Timestamp int64  `yaml:"timestamp"`
	Average   int    `yaml:"average"`
}
type HudsonNodeMonitorsClockMonitor struct {
	Class string `yaml:"_class"`
	Diff  int    `yaml:"diff"`
}
type MonitorData struct {
	HudsonNodeMonitorsSwapSpaceMonitor      HudsonNodeMonitorsSwapSpaceMonitor      `yaml:"hudson.node_monitors.SwapSpaceMonitor"`
	HudsonNodeMonitorsTemporarySpaceMonitor HudsonNodeMonitorsTemporarySpaceMonitor `yaml:"hudson.node_monitors.TemporarySpaceMonitor"`
	HudsonNodeMonitorsDiskSpaceMonitor      HudsonNodeMonitorsDiskSpaceMonitor      `yaml:"hudson.node_monitors.DiskSpaceMonitor"`
	HudsonNodeMonitorsArchitectureMonitor   string                                  `yaml:"hudson.node_monitors.ArchitectureMonitor"`
	HudsonNodeMonitorsResponseTimeMonitor   HudsonNodeMonitorsResponseTimeMonitor   `yaml:"hudson.node_monitors.ResponseTimeMonitor"`
	HudsonNodeMonitorsClockMonitor          HudsonNodeMonitorsClockMonitor          `yaml:"hudson.node_monitors.ClockMonitor"`
}
type Computer struct {
	Class               string           `yaml:"_class"`
	Actions             []interface{}    `yaml:"actions"`
	AssignedLabels      []AssignedLabels `yaml:"assignedLabels"`
	Description         string           `yaml:"description"`
	DisplayName         string           `yaml:"displayName"`
	Executors           []Executors      `yaml:"executors"`
	Icon                string           `yaml:"icon"`
	IconClassName       string           `yaml:"iconClassName"`
	Idle                bool             `yaml:"idle"`
	JnlpAgent           bool             `yaml:"jnlpAgent"`
	LaunchSupported     bool             `yaml:"launchSupported"`
	LoadStatistics      LoadStatistics   `yaml:"loadStatistics"`
	ManualLaunchAllowed bool             `yaml:"manualLaunchAllowed"`
	MonitorData         MonitorData      `yaml:"monitorData"`
	NumExecutors        int              `yaml:"numExecutors"`
	Offline             bool             `yaml:"offline"`
	OfflineCause        interface{}      `yaml:"offlineCause"`
	OfflineCauseReason  string           `yaml:"offlineCauseReason"`
	OneOffExecutors     []interface{}    `yaml:"oneOffExecutors"`
	TemporarilyOffline  bool             `yaml:"temporarilyOffline"`
}

func ListNodes() (obj NodesResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "nodes"})
	if err != nil {
		Error.Println(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		Error.Println(err)
	}
	return
}

func GetNode(node_name string) (obj Computer) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "node", "node_name": node_name})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(responseData))
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		Error.Println(err)
	}

	return
}