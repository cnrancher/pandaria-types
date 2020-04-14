package client

const (
	WebhookConfigType             = "webhookConfig"
	WebhookConfigFieldBearerToken = "bearer_token"
	WebhookConfigFieldProxyURL    = "proxyUrl"
	WebhookConfigFieldURL         = "url"
)

type WebhookConfig struct {
	BearerToken string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
