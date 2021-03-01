package client

const (
	TLSProfileSpecType       = "tlsProfileSpec"
	TLSProfileSpecFieldHosts = "hosts"
	TLSProfileSpecFieldTLS   = "tls"
)

type TLSProfileSpec struct {
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	TLS   *TLS     `json:"tls,omitempty" yaml:"tls,omitempty"`
}
