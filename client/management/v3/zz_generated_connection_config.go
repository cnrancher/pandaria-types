package client

const (
	ConnectionConfigType              = "connectionConfig"
	ConnectionConfigFieldAPIEndpoint  = "apiEndpoint"
	ConnectionConfigFieldClusterID    = "clusterID"
	ConnectionConfigFieldDirectAccess = "directAccess"
)

type ConnectionConfig struct {
	APIEndpoint  string `json:"apiEndpoint,omitempty" yaml:"apiEndpoint,omitempty"`
	ClusterID    string `json:"clusterID,omitempty" yaml:"clusterID,omitempty"`
	DirectAccess string `json:"directAccess,omitempty" yaml:"directAccess,omitempty"`
}
