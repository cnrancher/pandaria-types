package client

const (
	DingtalkConfigType             = "dingtalkConfig"
	DingtalkConfigFieldBearerToken = "bearer_token"
	DingtalkConfigFieldProxyURL    = "proxyUrl"
	DingtalkConfigFieldSecret      = "secret"
	DingtalkConfigFieldURL         = "url"
)

type DingtalkConfig struct {
	BearerToken string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	Secret      string `json:"secret,omitempty" yaml:"secret,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
