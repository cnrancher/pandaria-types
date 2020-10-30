package v3

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SensitiveFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Filters           []Filter `json:"filters"`
}

type Filter struct {
	rbacv1.PolicyRule `json:",inline"`
	Roles             []string `json:"roles,omitempty" norman:"type=reference[roleTemplate]"`
	RoleScope         string   `json:"roleScope,omitempty"`
	Fields            []string `json:"fields,omitempty"`
}
