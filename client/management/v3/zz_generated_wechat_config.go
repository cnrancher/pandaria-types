package client

const (
	WechatConfigType                  = "wechatConfig"
	WechatConfigFieldAPIURL           = "apiUrl"
	WechatConfigFieldAgent            = "agent"
	WechatConfigFieldBasicAuth        = "basic_auth"
	WechatConfigFieldBearerToken      = "bearer_token"
	WechatConfigFieldCorp             = "corp"
	WechatConfigFieldDefaultRecipient = "defaultRecipient"
	WechatConfigFieldProxyURL         = "proxyUrl"
	WechatConfigFieldRecipientType    = "recipientType"
	WechatConfigFieldSecret           = "secret"
)

type WechatConfig struct {
	APIURL           string     `json:"apiUrl,omitempty" yaml:"apiUrl,omitempty"`
	Agent            string     `json:"agent,omitempty" yaml:"agent,omitempty"`
	BasicAuth        *BasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`
	BearerToken      string     `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	Corp             string     `json:"corp,omitempty" yaml:"corp,omitempty"`
	DefaultRecipient string     `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	ProxyURL         string     `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	RecipientType    string     `json:"recipientType,omitempty" yaml:"recipientType,omitempty"`
	Secret           string     `json:"secret,omitempty" yaml:"secret,omitempty"`
}
