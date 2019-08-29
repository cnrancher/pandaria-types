package client

const (
	SSOLoginType                      = "ssoLogin"
	SSOLoginFieldCode                 = "code"
	SSOLoginFieldDescription          = "description"
	SSOLoginFieldDigest               = "digest"
	SSOLoginFieldJwt                  = "jwt"
	SSOLoginFieldRegion               = "region"
	SSOLoginFieldRegionClusterKeyName = "regionClusterKeyName"
	SSOLoginFieldResponseType         = "responseType"
	SSOLoginFieldTTLMillis            = "ttl"
)

type SSOLogin struct {
	Code                 string `json:"code,omitempty" yaml:"code,omitempty"`
	Description          string `json:"description,omitempty" yaml:"description,omitempty"`
	Digest               string `json:"digest,omitempty" yaml:"digest,omitempty"`
	Jwt                  string `json:"jwt,omitempty" yaml:"jwt,omitempty"`
	Region               string `json:"region,omitempty" yaml:"region,omitempty"`
	RegionClusterKeyName string `json:"regionClusterKeyName,omitempty" yaml:"regionClusterKeyName,omitempty"`
	ResponseType         string `json:"responseType,omitempty" yaml:"responseType,omitempty"`
	TTLMillis            int64  `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
