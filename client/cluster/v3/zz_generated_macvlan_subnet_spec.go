package client

const (
	MacvlanSubnetSpecType        = "macvlanSubnetSpec"
	MacvlanSubnetSpecFieldCIDR   = "cidr"
	MacvlanSubnetSpecFieldMaster = "master"
	MacvlanSubnetSpecFieldMode   = "mode"
)

type MacvlanSubnetSpec struct {
	CIDR   string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Master string `json:"master,omitempty" yaml:"master,omitempty"`
	Mode   string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
