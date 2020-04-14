package client

const (
	SlackConfigType                  = "slackConfig"
	SlackConfigFieldBearerToken      = "bearer_token"
	SlackConfigFieldDefaultRecipient = "defaultRecipient"
	SlackConfigFieldProxyURL         = "proxyUrl"
	SlackConfigFieldURL              = "url"
)

type SlackConfig struct {
	BearerToken      string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	DefaultRecipient string `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	ProxyURL         string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL              string `json:"url,omitempty" yaml:"url,omitempty"`
}
