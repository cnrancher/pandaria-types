package client

const (
	PoolType                 = "pool"
	PoolFieldMonitor         = "monitor"
	PoolFieldNodeMemberLabel = "nodeMemberLabel"
	PoolFieldPath            = "path"
	PoolFieldRewrite         = "rewrite"
	PoolFieldService         = "service"
	PoolFieldServicePort     = "servicePort"
)

type Pool struct {
	Monitor         *Monitor `json:"monitor,omitempty" yaml:"monitor,omitempty"`
	NodeMemberLabel string   `json:"nodeMemberLabel,omitempty" yaml:"nodeMemberLabel,omitempty"`
	Path            string   `json:"path,omitempty" yaml:"path,omitempty"`
	Rewrite         string   `json:"rewrite,omitempty" yaml:"rewrite,omitempty"`
	Service         string   `json:"service,omitempty" yaml:"service,omitempty"`
	ServicePort     int64    `json:"servicePort,omitempty" yaml:"servicePort,omitempty"`
}
