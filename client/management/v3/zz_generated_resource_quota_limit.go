package client

const (
	ResourceQuotaLimitType                             = "resourceQuotaLimit"
	ResourceQuotaLimitFieldConfigMaps                  = "configMaps"
	ResourceQuotaLimitFieldLimitsCPU                   = "limitsCpu"
	ResourceQuotaLimitFieldLimitsMemory                = "limitsMemory"
	ResourceQuotaLimitFieldPersistentVolumeClaims      = "persistentVolumeClaims"
	ResourceQuotaLimitFieldPods                        = "pods"
	ResourceQuotaLimitFieldReplicationControllers      = "replicationControllers"
	ResourceQuotaLimitFieldRequestsCPU                 = "requestsCpu"
	ResourceQuotaLimitFieldRequestsGPUCount            = "requestsGpuCount"
	ResourceQuotaLimitFieldRequestsGPUMemory           = "requestsGpuMemory"
	ResourceQuotaLimitFieldRequestsMemory              = "requestsMemory"
	ResourceQuotaLimitFieldRequestsStorage             = "requestsStorage"
	ResourceQuotaLimitFieldRequestsStorageClassPVC     = "requestsStorageClassPVC"
	ResourceQuotaLimitFieldRequestsStorageClassStorage = "requestsStorageClassStorage"
	ResourceQuotaLimitFieldSecrets                     = "secrets"
	ResourceQuotaLimitFieldServices                    = "services"
	ResourceQuotaLimitFieldServicesLoadBalancers       = "servicesLoadBalancers"
	ResourceQuotaLimitFieldServicesNodePorts           = "servicesNodePorts"
)

type ResourceQuotaLimit struct {
	ConfigMaps                  string            `json:"configMaps,omitempty" yaml:"configMaps,omitempty"`
	LimitsCPU                   string            `json:"limitsCpu,omitempty" yaml:"limitsCpu,omitempty"`
	LimitsMemory                string            `json:"limitsMemory,omitempty" yaml:"limitsMemory,omitempty"`
	PersistentVolumeClaims      string            `json:"persistentVolumeClaims,omitempty" yaml:"persistentVolumeClaims,omitempty"`
	Pods                        string            `json:"pods,omitempty" yaml:"pods,omitempty"`
	ReplicationControllers      string            `json:"replicationControllers,omitempty" yaml:"replicationControllers,omitempty"`
	RequestsCPU                 string            `json:"requestsCpu,omitempty" yaml:"requestsCpu,omitempty"`
	RequestsGPUCount            string            `json:"requestsGpuCount,omitempty" yaml:"requestsGpuCount,omitempty"`
	RequestsGPUMemory           string            `json:"requestsGpuMemory,omitempty" yaml:"requestsGpuMemory,omitempty"`
	RequestsMemory              string            `json:"requestsMemory,omitempty" yaml:"requestsMemory,omitempty"`
	RequestsStorage             string            `json:"requestsStorage,omitempty" yaml:"requestsStorage,omitempty"`
	RequestsStorageClassPVC     map[string]string `json:"requestsStorageClassPVC,omitempty" yaml:"requestsStorageClassPVC,omitempty"`
	RequestsStorageClassStorage map[string]string `json:"requestsStorageClassStorage,omitempty" yaml:"requestsStorageClassStorage,omitempty"`
	Secrets                     string            `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Services                    string            `json:"services,omitempty" yaml:"services,omitempty"`
	ServicesLoadBalancers       string            `json:"servicesLoadBalancers,omitempty" yaml:"servicesLoadBalancers,omitempty"`
	ServicesNodePorts           string            `json:"servicesNodePorts,omitempty" yaml:"servicesNodePorts,omitempty"`
}
