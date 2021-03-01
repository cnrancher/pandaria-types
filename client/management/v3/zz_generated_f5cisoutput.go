package client

const (
	F5CISOutputType              = "f5CISOutput"
	F5CISOutputFieldAnswers      = "answers"
	F5CISOutputFieldExtraAnswers = "extraAnswers"
	F5CISOutputFieldValuesYaml   = "valuesYaml"
	F5CISOutputFieldVersion      = "version"
)

type F5CISOutput struct {
	Answers      map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ExtraAnswers map[string]string `json:"extraAnswers,omitempty" yaml:"extraAnswers,omitempty"`
	ValuesYaml   string            `json:"valuesYaml,omitempty" yaml:"valuesYaml,omitempty"`
	Version      string            `json:"version,omitempty" yaml:"version,omitempty"`
}
