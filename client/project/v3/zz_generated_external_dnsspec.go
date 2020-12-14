package client

const (
	ExternalDNSSpecType                   = "externalDNSSpec"
	ExternalDNSSpecFieldDNSRecordType     = "dnsRecordType"
	ExternalDNSSpecFieldDomainName        = "domainName"
	ExternalDNSSpecFieldLoadBalanceMethod = "loadBalanceMethod"
	ExternalDNSSpecFieldPools             = "pools"
)

type ExternalDNSSpec struct {
	DNSRecordType     string    `json:"dnsRecordType,omitempty" yaml:"dnsRecordType,omitempty"`
	DomainName        string    `json:"domainName,omitempty" yaml:"domainName,omitempty"`
	LoadBalanceMethod string    `json:"loadBalanceMethod,omitempty" yaml:"loadBalanceMethod,omitempty"`
	Pools             []DNSPool `json:"pools,omitempty" yaml:"pools,omitempty"`
}
