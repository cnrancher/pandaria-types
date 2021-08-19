package client

const (
	AffinityTermType             = "affinityTerm"
	AffinityTermFieldExpressions = "expressions"
	AffinityTermFieldNamespaces  = "namespaces"
	AffinityTermFieldTopology    = "topology"
)

type AffinityTerm struct {
	Expressions []string `json:"expressions,omitempty" yaml:"expressions,omitempty"`
	Namespaces  []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Topology    string   `json:"topology,omitempty" yaml:"topology,omitempty"`
}
