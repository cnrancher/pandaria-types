package client

const (
	CloneTargetType           = "cloneTarget"
	CloneTargetFieldNamespace = "namespace"
	CloneTargetFieldProject   = "project"
)

type CloneTarget struct {
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Project   string `json:"project,omitempty" yaml:"project,omitempty"`
}
