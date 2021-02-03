package client

const (
	ServerAddresseType             = "serverAddresse"
	ServerAddresseFieldDeviceName  = "deviceName"
	ServerAddresseFieldName        = "name"
	ServerAddresseFieldTranslation = "translation"
)

type ServerAddresse struct {
	DeviceName  string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Translation string `json:"translation,omitempty" yaml:"translation,omitempty"`
}
