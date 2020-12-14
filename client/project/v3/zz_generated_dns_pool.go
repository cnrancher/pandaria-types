package client

const (
	DNSPoolType                   = "dnsPool"
	DNSPoolFieldDNSRecordType     = "dnsRecordType"
	DNSPoolFieldDataServerName    = "dataServerName"
	DNSPoolFieldLoadBalanceMethod = "loadBalanceMethod"
	DNSPoolFieldMonitor           = "monitor"
	DNSPoolFieldName              = "name"
)

type DNSPool struct {
	DNSRecordType     string   `json:"dnsRecordType,omitempty" yaml:"dnsRecordType,omitempty"`
	DataServerName    string   `json:"dataServerName,omitempty" yaml:"dataServerName,omitempty"`
	LoadBalanceMethod string   `json:"loadBalanceMethod,omitempty" yaml:"loadBalanceMethod,omitempty"`
	Monitor           *Monitor `json:"monitor,omitempty" yaml:"monitor,omitempty"`
	Name              string   `json:"name,omitempty" yaml:"name,omitempty"`
}
