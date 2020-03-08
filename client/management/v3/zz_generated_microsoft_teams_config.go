package client

const (
	MicrosoftTeamsConfigType          = "microsoftTeamsConfig"
	MicrosoftTeamsConfigFieldProxyURL = "proxyUrl"
	MicrosoftTeamsConfigFieldURL      = "url"
)

type MicrosoftTeamsConfig struct {
	ProxyURL string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL      string `json:"url,omitempty" yaml:"url,omitempty"`
}
