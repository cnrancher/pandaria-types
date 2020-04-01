package client

const (
	HarborAdminAuthInputType           = "harborAdminAuthInput"
	HarborAdminAuthInputFieldPassword  = "password"
	HarborAdminAuthInputFieldServerURL = "serverURL"
	HarborAdminAuthInputFieldUsername  = "username"
)

type HarborAdminAuthInput struct {
	Password  string `json:"password,omitempty" yaml:"password,omitempty"`
	ServerURL string `json:"serverURL,omitempty" yaml:"serverURL,omitempty"`
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
}
