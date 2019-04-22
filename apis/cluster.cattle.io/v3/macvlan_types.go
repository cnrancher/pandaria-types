package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MacvlanSubnet struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MacvlanSubnetSpec `json:"spec"`
}

type MacvlanSubnetSpec struct {
	Master string `json:"master"`
	CIDR   string `json:"cidr"`
	Mode   string `json:"mode"`
}

type MacvlanIP struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MacvlanIPSpec `json:"spec"`
}

type MacvlanIPSpec struct {
	VLAN  string `json:"vlan"`
	PodID string `json:"podId"`
	CIDR  string `json:"cidr"`
	MAC   string `json:"mac"`
}
