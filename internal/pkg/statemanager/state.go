package statemanager

type ComponentStatus string
type ComponentAction string

const (
	// We should delete the state of the "uninstalled" tool at StatesMap.
	StatusUninstalled ComponentStatus = "uninstalled"
	StatusInstalled   ComponentStatus = "installed"
	StatusRunning     ComponentStatus = "running"
	StatusFailed      ComponentStatus = "Failed"
)

const (
	ActionInstall   ComponentAction = "install"
	ActionReinstall ComponentAction = "reinstall"
	ActionUninstall ComponentAction = "uninstall"
)

// State is the single component's state.
type State struct {
	Name          string          `yaml:"name"`
	Version       string          `yaml:"version"`
	Dependencies  []string        `yaml:"dependencies"`
	Status        ComponentStatus `yaml:"status"`
	LastOperation *Operation      `yaml:"lastOperation"`
}

type Operation struct {
	Action   ComponentAction        `yaml:"action"`
	Time     string                 `yaml:"time"`
	Metadata map[string]interface{} `yaml:"metadata"`
}

func NewState(name, version string, deps []string, status ComponentStatus, lastOpr *Operation) *State {
	return &State{
		Name:          name,
		Version:       version,
		Dependencies:  deps,
		Status:        status,
		LastOperation: lastOpr,
	}
}
