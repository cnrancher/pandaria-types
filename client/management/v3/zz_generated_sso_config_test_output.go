package client

const (
	SSOConfigTestOutputType             = "ssoConfigTestOutput"
	SSOConfigTestOutputFieldRedirectURL = "redirectUrl"
)

type SSOConfigTestOutput struct {
	RedirectURL string `json:"redirectUrl,omitempty" yaml:"redirectUrl,omitempty"`
}
