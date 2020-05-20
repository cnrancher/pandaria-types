package client

const (
	CASLogoutOutputType           = "casLogoutOutput"
	CASLogoutOutputFieldLogoutURL = "logoutUrl"
)

type CASLogoutOutput struct {
	LogoutURL string `json:"logoutUrl,omitempty" yaml:"logoutUrl,omitempty"`
}
