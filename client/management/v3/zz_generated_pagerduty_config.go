package client

const (
	PagerdutyConfigType             = "pagerdutyConfig"
	PagerdutyConfigFieldBearerToken = "bearerToken"
	PagerdutyConfigFieldProxyURL    = "proxyUrl"
	PagerdutyConfigFieldServiceKey  = "serviceKey"
)

type PagerdutyConfig struct {
	BearerToken string `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	ProxyURL    string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	ServiceKey  string `json:"serviceKey,omitempty" yaml:"serviceKey,omitempty"`
}
