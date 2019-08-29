package client

const (
	SSOConfigType                     = "ssoConfig"
	SSOConfigFieldAccessMode          = "accessMode"
	SSOConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	SSOConfigFieldAnnotations         = "annotations"
	SSOConfigFieldClientID            = "clientId"
	SSOConfigFieldClientSecret        = "clientSecret"
	SSOConfigFieldCreated             = "created"
	SSOConfigFieldCreatorID           = "creatorId"
	SSOConfigFieldEnabled             = "enabled"
	SSOConfigFieldHostname            = "hostname"
	SSOConfigFieldLabels              = "labels"
	SSOConfigFieldName                = "name"
	SSOConfigFieldOwnerReferences     = "ownerReferences"
	SSOConfigFieldRemoved             = "removed"
	SSOConfigFieldTLS                 = "tls"
	SSOConfigFieldType                = "type"
	SSOConfigFieldUUID                = "uuid"
)

type SSOConfig struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID            string            `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret        string            `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname            string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS                 bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
