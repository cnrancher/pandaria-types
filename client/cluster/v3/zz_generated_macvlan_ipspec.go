package client

const (
	MacvlanIPSpecType       = "macvlanIPSpec"
	MacvlanIPSpecFieldCIDR  = "cidr"
	MacvlanIPSpecFieldMAC   = "mac"
	MacvlanIPSpecFieldPodID = "podId"
	MacvlanIPSpecFieldVLAN  = "vlan"
)

type MacvlanIPSpec struct {
	CIDR  string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	MAC   string `json:"mac,omitempty" yaml:"mac,omitempty"`
	PodID string `json:"podId,omitempty" yaml:"podId,omitempty"`
	VLAN  string `json:"vlan,omitempty" yaml:"vlan,omitempty"`
}
