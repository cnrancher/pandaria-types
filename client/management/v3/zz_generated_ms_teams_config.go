package client

const (
	MSTeamsConfigType             = "msTeamsConfig"
	MSTeamsConfigFieldBearerToken = "bearer_token"
	MSTeamsConfigFieldProxyURL    = "proxyUrl"
	MSTeamsConfigFieldURL         = "url"
)

type MSTeamsConfig struct {
	BearerToken string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
