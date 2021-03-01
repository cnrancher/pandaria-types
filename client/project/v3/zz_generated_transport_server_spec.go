package client

const (
	TransportServerSpecType                      = "transportServerSpec"
	TransportServerSpecFieldMode                 = "mode"
	TransportServerSpecFieldPool                 = "pool"
	TransportServerSpecFieldSNAT                 = "snat"
	TransportServerSpecFieldVirtualServerAddress = "virtualServerAddress"
	TransportServerSpecFieldVirtualServerName    = "virtualServerName"
	TransportServerSpecFieldVirtualServerPort    = "virtualServerPort"
)

type TransportServerSpec struct {
	Mode                 string `json:"mode,omitempty" yaml:"mode,omitempty"`
	Pool                 *Pool  `json:"pool,omitempty" yaml:"pool,omitempty"`
	SNAT                 string `json:"snat,omitempty" yaml:"snat,omitempty"`
	VirtualServerAddress string `json:"virtualServerAddress,omitempty" yaml:"virtualServerAddress,omitempty"`
	VirtualServerName    string `json:"virtualServerName,omitempty" yaml:"virtualServerName,omitempty"`
	VirtualServerPort    int64  `json:"virtualServerPort,omitempty" yaml:"virtualServerPort,omitempty"`
}
