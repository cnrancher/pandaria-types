package client

const (
	WorkloadMetricRelabelConfigType        = "workloadMetricRelabelConfig"
	WorkloadMetricRelabelConfigFieldAction = "action"
	WorkloadMetricRelabelConfigFieldRegex  = "regex"
)

type WorkloadMetricRelabelConfig struct {
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	Regex  string `json:"regex,omitempty" yaml:"regex,omitempty"`
}
