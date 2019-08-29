package client

const (
	SSOConfigApplyInputType           = "ssoConfigApplyInput"
	SSOConfigApplyInputFieldCode      = "code"
	SSOConfigApplyInputFieldEnabled   = "enabled"
	SSOConfigApplyInputFieldSSOConfig = "ssoConfig"
)

type SSOConfigApplyInput struct {
	Code      string     `json:"code,omitempty" yaml:"code,omitempty"`
	Enabled   bool       `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	SSOConfig *SSOConfig `json:"ssoConfig,omitempty" yaml:"ssoConfig,omitempty"`
}
