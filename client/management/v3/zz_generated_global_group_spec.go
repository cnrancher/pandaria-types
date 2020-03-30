package client

const (
	GlobalGroupSpecType                       = "globalGroupSpec"
	GlobalGroupSpecFieldDescription           = "description"
	GlobalGroupSpecFieldDisplayName           = "displayName"
	GlobalGroupSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	GlobalGroupSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	GlobalGroupSpecFieldRecipients            = "recipients"
	GlobalGroupSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
)

type GlobalGroupSpec struct {
	Description           string      `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string      `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupIntervalSeconds  int64       `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64       `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Recipients            []Recipient `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64       `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
}
