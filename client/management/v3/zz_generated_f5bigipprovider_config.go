package client

const (
	F5BIGIPProviderConfigType                       = "f5BIGIPProviderConfig"
	F5BIGIPProviderConfigFieldF5BIGIPDatacenterName = "f5BigipDatacenterName"
	F5BIGIPProviderConfigFieldF5BIGIPDeviceIPS      = "f5BigipDeviceIPs"
	F5BIGIPProviderConfigFieldF5BIGIPHost           = "f5BigipHost"
	F5BIGIPProviderConfigFieldF5BIGIPPasswd         = "f5BigipPasswd"
	F5BIGIPProviderConfigFieldF5BIGIPPort           = "f5BigipPort"
	F5BIGIPProviderConfigFieldF5BIGIPServerName     = "f5BigipServerName"
	F5BIGIPProviderConfigFieldF5BIGIPUser           = "f5BigipUser"
)

type F5BIGIPProviderConfig struct {
	F5BIGIPDatacenterName string           `json:"f5BigipDatacenterName,omitempty" yaml:"f5BigipDatacenterName,omitempty"`
	F5BIGIPDeviceIPS      []ServerAddresse `json:"f5BigipDeviceIPs,omitempty" yaml:"f5BigipDeviceIPs,omitempty"`
	F5BIGIPHost           string           `json:"f5BigipHost,omitempty" yaml:"f5BigipHost,omitempty"`
	F5BIGIPPasswd         string           `json:"f5BigipPasswd,omitempty" yaml:"f5BigipPasswd,omitempty"`
	F5BIGIPPort           string           `json:"f5BigipPort,omitempty" yaml:"f5BigipPort,omitempty"`
	F5BIGIPServerName     string           `json:"f5BigipServerName,omitempty" yaml:"f5BigipServerName,omitempty"`
	F5BIGIPUser           string           `json:"f5BigipUser,omitempty" yaml:"f5BigipUser,omitempty"`
}
