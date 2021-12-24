package client

const (
	ContainerResourceLimitType                = "containerResourceLimit"
	ContainerResourceLimitFieldLimitsCPU      = "limitsCpu"
	ContainerResourceLimitFieldLimitsMemory   = "limitsMemory"
	ContainerResourceLimitFieldMaxCPU         = "maxCpu"
	ContainerResourceLimitFieldMaxMemory      = "maxMemory"
	ContainerResourceLimitFieldMinCPU         = "minCpu"
	ContainerResourceLimitFieldMinMemory      = "minMemory"
	ContainerResourceLimitFieldRequestsCPU    = "requestsCpu"
	ContainerResourceLimitFieldRequestsMemory = "requestsMemory"
)

type ContainerResourceLimit struct {
	LimitsCPU      string `json:"limitsCpu,omitempty" yaml:"limitsCpu,omitempty"`
	LimitsMemory   string `json:"limitsMemory,omitempty" yaml:"limitsMemory,omitempty"`
	MaxCPU         string `json:"maxCpu,omitempty" yaml:"maxCpu,omitempty"`
	MaxMemory      string `json:"maxMemory,omitempty" yaml:"maxMemory,omitempty"`
	MinCPU         string `json:"minCpu,omitempty" yaml:"minCpu,omitempty"`
	MinMemory      string `json:"minMemory,omitempty" yaml:"minMemory,omitempty"`
	RequestsCPU    string `json:"requestsCpu,omitempty" yaml:"requestsCpu,omitempty"`
	RequestsMemory string `json:"requestsMemory,omitempty" yaml:"requestsMemory,omitempty"`
}
