package client

const (
	F5CISInputType              = "f5CISInput"
	F5CISInputFieldAnswers      = "answers"
	F5CISInputFieldExtraAnswers = "extraAnswers"
	F5CISInputFieldValuesYaml   = "valuesYaml"
	F5CISInputFieldVersion      = "version"
)

type F5CISInput struct {
	Answers      map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ExtraAnswers map[string]string `json:"extraAnswers,omitempty" yaml:"extraAnswers,omitempty"`
	ValuesYaml   string            `json:"valuesYaml,omitempty" yaml:"valuesYaml,omitempty"`
	Version      string            `json:"version,omitempty" yaml:"version,omitempty"`
}
