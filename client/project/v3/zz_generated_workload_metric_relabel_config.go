package client

const (
	WorkloadMetricRelabelConfigType              = "workloadMetricRelabelConfig"
	WorkloadMetricRelabelConfigFieldAction       = "action"
	WorkloadMetricRelabelConfigFieldRegex        = "regex"
	WorkloadMetricRelabelConfigFieldRelabelType  = "relabelType"
	WorkloadMetricRelabelConfigFieldReplacement  = "replacement"
	WorkloadMetricRelabelConfigFieldSourceLabels = "sourceLabels"
	WorkloadMetricRelabelConfigFieldTargetLabel  = "targetLabel"
)

type WorkloadMetricRelabelConfig struct {
	Action       string   `json:"action,omitempty" yaml:"action,omitempty"`
	Regex        string   `json:"regex,omitempty" yaml:"regex,omitempty"`
	RelabelType  string   `json:"relabelType,omitempty" yaml:"relabelType,omitempty"`
	Replacement  string   `json:"replacement,omitempty" yaml:"replacement,omitempty"`
	SourceLabels []string `json:"sourceLabels,omitempty" yaml:"sourceLabels,omitempty"`
	TargetLabel  string   `json:"targetLabel,omitempty" yaml:"targetLabel,omitempty"`
}
