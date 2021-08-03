package client

const (
	ClusterTestInputType                           = "clusterTestInput"
	ClusterTestInputFieldClusterName               = "clusterId"
	ClusterTestInputFieldCustomBuffer              = "customBuffer"
	ClusterTestInputFieldCustomTargetConfig        = "customTargetConfig"
	ClusterTestInputFieldDisplayName               = "displayName"
	ClusterTestInputFieldElasticsearchConfig       = "elasticsearchConfig"
	ClusterTestInputFieldEnableExceptionStackMatch = "enableExceptionStackMatch"
	ClusterTestInputFieldEnableJSONParsing         = "enableJSONParsing"
	ClusterTestInputFieldEnableMultiLineFilter     = "enableMultiLineFilter"
	ClusterTestInputFieldFluentForwarderConfig     = "fluentForwarderConfig"
	ClusterTestInputFieldKafkaConfig               = "kafkaConfig"
	ClusterTestInputFieldMultiLineContinuousRegexp = "multiLineContinuousRegexp"
	ClusterTestInputFieldMultiLineEndRegexp        = "multiLineEndRegexp"
	ClusterTestInputFieldMultiLineSeparator        = "multiLineSeparator"
	ClusterTestInputFieldMultiLineStartRegexp      = "multiLineStartRegexp"
	ClusterTestInputFieldOutputFlushInterval       = "outputFlushInterval"
	ClusterTestInputFieldOutputTags                = "outputTags"
	ClusterTestInputFieldSplunkConfig              = "splunkConfig"
	ClusterTestInputFieldSyslogConfig              = "syslogConfig"
)

type ClusterTestInput struct {
	ClusterName               string                 `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	CustomBuffer              string                 `json:"customBuffer,omitempty" yaml:"customBuffer,omitempty"`
	CustomTargetConfig        *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"customTargetConfig,omitempty"`
	DisplayName               string                 `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ElasticsearchConfig       *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`
	EnableExceptionStackMatch bool                   `json:"enableExceptionStackMatch,omitempty" yaml:"enableExceptionStackMatch,omitempty"`
	EnableJSONParsing         bool                   `json:"enableJSONParsing,omitempty" yaml:"enableJSONParsing,omitempty"`
	EnableMultiLineFilter     bool                   `json:"enableMultiLineFilter,omitempty" yaml:"enableMultiLineFilter,omitempty"`
	FluentForwarderConfig     *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluentForwarderConfig,omitempty"`
	KafkaConfig               *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafkaConfig,omitempty"`
	MultiLineContinuousRegexp string                 `json:"multiLineContinuousRegexp,omitempty" yaml:"multiLineContinuousRegexp,omitempty"`
	MultiLineEndRegexp        string                 `json:"multiLineEndRegexp,omitempty" yaml:"multiLineEndRegexp,omitempty"`
	MultiLineSeparator        string                 `json:"multiLineSeparator,omitempty" yaml:"multiLineSeparator,omitempty"`
	MultiLineStartRegexp      string                 `json:"multiLineStartRegexp,omitempty" yaml:"multiLineStartRegexp,omitempty"`
	OutputFlushInterval       int64                  `json:"outputFlushInterval,omitempty" yaml:"outputFlushInterval,omitempty"`
	OutputTags                map[string]string      `json:"outputTags,omitempty" yaml:"outputTags,omitempty"`
	SplunkConfig              *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunkConfig,omitempty"`
	SyslogConfig              *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslogConfig,omitempty"`
}
