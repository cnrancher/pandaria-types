package client

const (
	AlibabaCloudConfigType                 = "alibabaCloudConfig"
	AlibabaCloudConfigFieldAccessKeyID     = "accessKeyID"
	AlibabaCloudConfigFieldAccessKeySecret = "accessKeySecret"
	AlibabaCloudConfigFieldSignName        = "signName"
	AlibabaCloudConfigFieldTemplateCode    = "templateCode"
	AlibabaCloudConfigFieldTo              = "to"
)

type AlibabaCloudConfig struct {
	AccessKeyID     string   `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	AccessKeySecret string   `json:"accessKeySecret,omitempty" yaml:"accessKeySecret,omitempty"`
	SignName        string   `json:"signName,omitempty" yaml:"signName,omitempty"`
	TemplateCode    string   `json:"templateCode,omitempty" yaml:"templateCode,omitempty"`
	To              []string `json:"to,omitempty" yaml:"to,omitempty"`
}
