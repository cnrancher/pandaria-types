package client

const (
	FilterType                 = "filter"
	FilterFieldAPIGroups       = "apiGroups"
	FilterFieldFields          = "fields"
	FilterFieldNonResourceURLs = "nonResourceURLs"
	FilterFieldResourceNames   = "resourceNames"
	FilterFieldResources       = "resources"
	FilterFieldRoleScope       = "roleScope"
	FilterFieldRoles           = "roles"
	FilterFieldVerbs           = "verbs"
)

type Filter struct {
	APIGroups       []string `json:"apiGroups,omitempty" yaml:"apiGroups,omitempty"`
	Fields          []string `json:"fields,omitempty" yaml:"fields,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty" yaml:"nonResourceURLs,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty" yaml:"resourceNames,omitempty"`
	Resources       []string `json:"resources,omitempty" yaml:"resources,omitempty"`
	RoleScope       string   `json:"roleScope,omitempty" yaml:"roleScope,omitempty"`
	Roles           string   `json:"roles,omitempty" yaml:"roles,omitempty"`
	Verbs           []string `json:"verbs,omitempty" yaml:"verbs,omitempty"`
}
