package client

const (
	MonitoringOutputType            = "monitoringOutput"
	MonitoringOutputFieldAnswers    = "answers"
	MonitoringOutputFieldValuesYaml = "valuesYaml"
	MonitoringOutputFieldVersion    = "version"
)

type MonitoringOutput struct {
	Answers    map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ValuesYaml string            `json:"valuesYaml,omitempty" yaml:"valuesYaml,omitempty"`
	Version    string            `json:"version,omitempty" yaml:"version,omitempty"`
}
