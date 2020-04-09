package client

const (
	WebhookConfigType             = "webhookConfig"
	WebhookConfigFieldBearerToken = "bearerToken"
	WebhookConfigFieldProxyURL    = "proxyUrl"
	WebhookConfigFieldURL         = "url"
)

type WebhookConfig struct {
	BearerToken string `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
