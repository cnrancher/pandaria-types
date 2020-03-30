package client

const (
	GlobalAlertRuleSpecType                       = "globalAlertRuleSpec"
	GlobalAlertRuleSpecFieldDisplayName           = "displayName"
	GlobalAlertRuleSpecFieldGroupID               = "groupId"
	GlobalAlertRuleSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	GlobalAlertRuleSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	GlobalAlertRuleSpecFieldInherited             = "inherited"
	GlobalAlertRuleSpecFieldMetricRule            = "metricRule"
	GlobalAlertRuleSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	GlobalAlertRuleSpecFieldSeverity              = "severity"
)

type GlobalAlertRuleSpec struct {
	DisplayName           string      `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupID               string      `json:"groupId,omitempty" yaml:"groupId,omitempty"`
	GroupIntervalSeconds  int64       `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64       `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Inherited             *bool       `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	MetricRule            *MetricRule `json:"metricRule,omitempty" yaml:"metricRule,omitempty"`
	RepeatIntervalSeconds int64       `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	Severity              string      `json:"severity,omitempty" yaml:"severity,omitempty"`
}
