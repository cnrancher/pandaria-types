package client

const (
	CASLoginType              = "casLogin"
	CASLoginFieldDescription  = "description"
	CASLoginFieldResponseType = "responseType"
	CASLoginFieldTTLMillis    = "ttl"
	CASLoginFieldTicket       = "ticket"
)

type CASLogin struct {
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`
	TTLMillis    int64  `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	Ticket       string `json:"ticket,omitempty" yaml:"ticket,omitempty"`
}
