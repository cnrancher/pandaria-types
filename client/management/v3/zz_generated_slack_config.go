package client

const (
	SlackConfigType                  = "slackConfig"
	SlackConfigFieldBearerToken      = "bearerToken"
	SlackConfigFieldDefaultRecipient = "defaultRecipient"
	SlackConfigFieldProxyURL         = "proxyUrl"
	SlackConfigFieldURL              = "url"
)

type SlackConfig struct {
	BearerToken      string `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	DefaultRecipient string `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	ProxyURL         string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL              string `json:"url,omitempty" yaml:"url,omitempty"`
}
