package client

const (
	VirtualServerSpecType                        = "virtualServerSpec"
	VirtualServerSpecFieldHTTPTraffic            = "httpTraffic"
	VirtualServerSpecFieldHost                   = "host"
	VirtualServerSpecFieldPools                  = "pools"
	VirtualServerSpecFieldRewriteAppRoot         = "rewriteAppRoot"
	VirtualServerSpecFieldSNAT                   = "snat"
	VirtualServerSpecFieldTLSProfileName         = "tlsProfileName"
	VirtualServerSpecFieldVirtualServerAddress   = "virtualServerAddress"
	VirtualServerSpecFieldVirtualServerHTTPPort  = "virtualServerHTTPPort"
	VirtualServerSpecFieldVirtualServerHTTPSPort = "virtualServerHTTPSPort"
	VirtualServerSpecFieldVirtualServerName      = "virtualServerName"
	VirtualServerSpecFieldWAF                    = "waf"
)

type VirtualServerSpec struct {
	HTTPTraffic            string `json:"httpTraffic,omitempty" yaml:"httpTraffic,omitempty"`
	Host                   string `json:"host,omitempty" yaml:"host,omitempty"`
	Pools                  []Pool `json:"pools,omitempty" yaml:"pools,omitempty"`
	RewriteAppRoot         string `json:"rewriteAppRoot,omitempty" yaml:"rewriteAppRoot,omitempty"`
	SNAT                   string `json:"snat,omitempty" yaml:"snat,omitempty"`
	TLSProfileName         string `json:"tlsProfileName,omitempty" yaml:"tlsProfileName,omitempty"`
	VirtualServerAddress   string `json:"virtualServerAddress,omitempty" yaml:"virtualServerAddress,omitempty"`
	VirtualServerHTTPPort  int64  `json:"virtualServerHTTPPort,omitempty" yaml:"virtualServerHTTPPort,omitempty"`
	VirtualServerHTTPSPort int64  `json:"virtualServerHTTPSPort,omitempty" yaml:"virtualServerHTTPSPort,omitempty"`
	VirtualServerName      string `json:"virtualServerName,omitempty" yaml:"virtualServerName,omitempty"`
	WAF                    string `json:"waf,omitempty" yaml:"waf,omitempty"`
}
