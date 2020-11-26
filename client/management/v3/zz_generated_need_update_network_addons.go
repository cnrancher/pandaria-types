package client

const (
	NeedUpdateNetworkAddonsType            = "needUpdateNetworkAddons"
	NeedUpdateNetworkAddonsFieldMessage    = "message"
	NeedUpdateNetworkAddonsFieldNeedUpdate = "needUpdate"
)

type NeedUpdateNetworkAddons struct {
	Message    string `json:"message,omitempty" yaml:"message,omitempty"`
	NeedUpdate bool   `json:"needUpdate,omitempty" yaml:"needUpdate,omitempty"`
}
