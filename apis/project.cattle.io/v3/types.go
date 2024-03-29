package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceAccountToken struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	AccountName string `json:"accountName"`
	AccountUID  string `json:"accountUid"`
	Description string `json:"description"`
	Token       string `json:"token" norman:"writeOnly"`
	CACRT       string `json:"caCrt"`
}
type NamespacedServiceAccountToken ServiceAccountToken

type DockerCredential struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string                        `json:"description"`
	Registries  map[string]RegistryCredential `json:"registries"`
}
type NamespacedDockerCredential DockerCredential

type RegistryCredential struct {
	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password" norman:"writeOnly"`
	Auth        string `json:"auth" norman:"writeOnly"`
	Email       string `json:"email"`
}

type Certificate struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	Certs       string `json:"certs"`
	Key         string `json:"key" norman:"writeOnly"`

	CertFingerprint         string   `json:"certFingerprint" norman:"nocreate,noupdate"`
	CN                      string   `json:"cn" norman:"nocreate,noupdate"`
	Version                 string   `json:"version" norman:"nocreate,noupdate"`
	ExpiresAt               string   `json:"expiresAt" norman:"nocreate,noupdate"`
	Issuer                  string   `json:"issuer" norman:"nocreate,noupdate"`
	IssuedAt                string   `json:"issuedAt" norman:"nocreate,noupdate"`
	Algorithm               string   `json:"algorithm" norman:"nocreate,noupdate"`
	SerialNumber            string   `json:"serialNumber" norman:"nocreate,noupdate"`
	KeySize                 string   `json:"keySize" norman:"nocreate,noupdate"`
	SubjectAlternativeNames []string `json:"subjectAlternativeNames" norman:"nocreate,noupdate"`
}
type NamespacedCertificate Certificate

type BasicAuth struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password" norman:"writeOnly"`
}
type NamespacedBasicAuth BasicAuth

type SSHAuth struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	PrivateKey  string `json:"privateKey" norman:"writeOnly"`
	Fingerprint string `json:"certFingerprint" norman:"nocreate,noupdate"`
}
type NamespacedSSHAuth SSHAuth

type PublicEndpoint struct {
	NodeName  string   `json:"nodeName,omitempty" norman:"type=reference[/v3/schemas/node],nocreate,noupdate"`
	Addresses []string `json:"addresses,omitempty" norman:"nocreate,noupdate"`
	Port      int32    `json:"port,omitempty" norman:"nocreate,noupdate"`
	Protocol  string   `json:"protocol,omitempty" norman:"nocreate,noupdate"`
	// for node port service endpoint
	ServiceName string `json:"serviceName,omitempty" norman:"type=reference[service],nocreate,noupdate"`
	// for host port endpoint
	PodName string `json:"podName,omitempty" norman:"type=reference[pod],nocreate,noupdate"`
	// for ingress endpoint. ServiceName, podName, ingressName are mutually exclusive
	IngressName string `json:"ingressName,omitempty" norman:"type=reference[ingress],nocreate,noupdate"`
	// Hostname/path are set for Ingress endpoints
	Hostname string `json:"hostname,omitempty" norman:"nocreate,noupdate"`
	Path     string `json:"path,omitempty" norman:"nocreate,noupdate"`
	// True when endpoint is exposed on every node
	AllNodes bool `json:"allNodes" norman:"nocreate,noupdate"`
}

type Workload struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

type DeploymentRollbackInput struct {
	ReplicaSetID string `json:"replicaSetId" norman:"type=reference[replicaSet]"`
}

type WorkloadMetric struct {
	Port                        int32                         `json:"port,omitempty"`
	Path                        string                        `json:"path,omitempty"`
	Schema                      string                        `json:"schema,omitempty" norman:"type=enum,options=HTTP|HTTPS"`
	WorkloadMetricRelabelConfig []WorkloadMetricRelabelConfig `json:"workloadMetricRelabelConfig,omitempty"`
}

type WorkloadMetricRelabelConfig struct {
	RelabelType  string   `json:"relabelType,omitempty" norman:"type=enum,options=metricRelabel|relabel,default=metricRelabel"`
	TargetLabel  string   `json:"targetLabel,omitempty"`
	SourceLabels []string `json:"sourceLabels,omitempty"`
	Replacement  string   `json:"replacement,omitempty"`
	Regex        string   `json:"regex,omitempty"`
	Action       string   `json:"action,omitempty" norman:"type=enum,options=replace|keep|drop|labelmap|labeldrop,default=labeldrop"`
}
