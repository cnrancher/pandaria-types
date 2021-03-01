package client

const (
	TLSType             = "tls"
	TLSFieldClientSSL   = "clientSSL"
	TLSFieldReference   = "reference"
	TLSFieldServerSSL   = "serverSSL"
	TLSFieldTermination = "termination"
)

type TLS struct {
	ClientSSL   string `json:"clientSSL,omitempty" yaml:"clientSSL,omitempty"`
	Reference   string `json:"reference,omitempty" yaml:"reference,omitempty"`
	ServerSSL   string `json:"serverSSL,omitempty" yaml:"serverSSL,omitempty"`
	Termination string `json:"termination,omitempty" yaml:"termination,omitempty"`
}
