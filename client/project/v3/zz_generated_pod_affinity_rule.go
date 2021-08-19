package client

const (
	PodAffinityRuleType           = "podAffinityRule"
	PodAffinityRuleFieldPreferred = "preferred"
	PodAffinityRuleFieldRequired  = "required"
)

type PodAffinityRule struct {
	Preferred []AffinityTerm `json:"preferred,omitempty" yaml:"preferred,omitempty"`
	Required  []AffinityTerm `json:"required,omitempty" yaml:"required,omitempty"`
}
