package client

const (
	PagerdutyConfigType             = "pagerdutyConfig"
	PagerdutyConfigFieldBearerToken = "bearer_token"
	PagerdutyConfigFieldProxyURL    = "proxyUrl"
	PagerdutyConfigFieldServiceKey  = "serviceKey"
)

type PagerdutyConfig struct {
	BearerToken string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	ServiceKey  string `json:"serviceKey,omitempty" yaml:"serviceKey,omitempty"`
}
