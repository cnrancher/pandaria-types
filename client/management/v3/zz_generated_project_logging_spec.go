package client

const (
	ProjectLoggingSpecType                           = "projectLoggingSpec"
	ProjectLoggingSpecFieldCustomBuffer              = "customBuffer"
	ProjectLoggingSpecFieldCustomTargetConfig        = "customTargetConfig"
	ProjectLoggingSpecFieldDisplayName               = "displayName"
	ProjectLoggingSpecFieldElasticsearchConfig       = "elasticsearchConfig"
	ProjectLoggingSpecFieldEnableExceptionStackMatch = "enableExceptionStackMatch"
	ProjectLoggingSpecFieldEnableJSONParsing         = "enableJSONParsing"
	ProjectLoggingSpecFieldEnableMultiLineFilter     = "enableMultiLineFilter"
	ProjectLoggingSpecFieldFluentForwarderConfig     = "fluentForwarderConfig"
	ProjectLoggingSpecFieldKafkaConfig               = "kafkaConfig"
	ProjectLoggingSpecFieldMultiLineContinuousRegexp = "multiLineContinuousRegexp"
	ProjectLoggingSpecFieldMultiLineEndRegexp        = "multiLineEndRegexp"
	ProjectLoggingSpecFieldMultiLineSeparator        = "multiLineSeparator"
	ProjectLoggingSpecFieldMultiLineStartRegexp      = "multiLineStartRegexp"
	ProjectLoggingSpecFieldOutputFlushInterval       = "outputFlushInterval"
	ProjectLoggingSpecFieldOutputTags                = "outputTags"
	ProjectLoggingSpecFieldProjectID                 = "projectId"
	ProjectLoggingSpecFieldSplunkConfig              = "splunkConfig"
	ProjectLoggingSpecFieldSyslogConfig              = "syslogConfig"
)

type ProjectLoggingSpec struct {
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
	ProjectID                 string                 `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	SplunkConfig              *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunkConfig,omitempty"`
	SyslogConfig              *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslogConfig,omitempty"`
}
