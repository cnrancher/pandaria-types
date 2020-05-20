package client

const (
	CASProviderType                 = "casProvider"
	CASProviderFieldAnnotations     = "annotations"
	CASProviderFieldCreated         = "created"
	CASProviderFieldCreatorID       = "creatorId"
	CASProviderFieldLabels          = "labels"
	CASProviderFieldLogoutURL       = "logoutUrl"
	CASProviderFieldName            = "name"
	CASProviderFieldOwnerReferences = "ownerReferences"
	CASProviderFieldRedirectURL     = "redirectUrl"
	CASProviderFieldRemoved         = "removed"
	CASProviderFieldType            = "type"
	CASProviderFieldUUID            = "uuid"
)

type CASProvider struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LogoutURL       string            `json:"logoutUrl,omitempty" yaml:"logoutUrl,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty" yaml:"redirectUrl,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type            string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
