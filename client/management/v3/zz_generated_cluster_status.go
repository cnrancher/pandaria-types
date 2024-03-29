package client

const (
	ClusterStatusType                                      = "clusterStatus"
	ClusterStatusFieldAADClientCertSecret                  = "aadClientCertSecret"
	ClusterStatusFieldAADClientSecret                      = "aadClientSecret"
	ClusterStatusFieldAPIEndpoint                          = "apiEndpoint"
	ClusterStatusFieldAgentFeatures                        = "agentFeatures"
	ClusterStatusFieldAgentImage                           = "agentImage"
	ClusterStatusFieldAllocatable                          = "allocatable"
	ClusterStatusFieldAppliedEnableNetworkPolicy           = "appliedEnableNetworkPolicy"
	ClusterStatusFieldAppliedPodSecurityPolicyTemplateName = "appliedPodSecurityPolicyTemplateId"
	ClusterStatusFieldAppliedSpec                          = "appliedSpec"
	ClusterStatusFieldAuthImage                            = "authImage"
	ClusterStatusFieldCACert                               = "caCert"
	ClusterStatusFieldCapabilities                         = "capabilities"
	ClusterStatusFieldCapacity                             = "capacity"
	ClusterStatusFieldCertificatesExpiration               = "certificatesExpiration"
	ClusterStatusFieldComponentStatuses                    = "componentStatuses"
	ClusterStatusFieldConditions                           = "conditions"
	ClusterStatusFieldCurrentCisRunName                    = "currentCisRunName"
	ClusterStatusFieldDriver                               = "driver"
	ClusterStatusFieldFailedSpec                           = "failedSpec"
	ClusterStatusFieldIstioEnabled                         = "istioEnabled"
	ClusterStatusFieldLimits                               = "limits"
	ClusterStatusFieldMonitoringStatus                     = "monitoringStatus"
	ClusterStatusFieldNodeCount                            = "nodeCount"
	ClusterStatusFieldNodeVersion                          = "nodeVersion"
	ClusterStatusFieldOpenStackSecret                      = "openStackSecret"
	ClusterStatusFieldPrivateRegistrySecret                = "privateRegistrySecret"
	ClusterStatusFieldRequested                            = "requested"
	ClusterStatusFieldS3CredentialSecret                   = "s3CredentialSecret"
	ClusterStatusFieldScheduledClusterScanStatus           = "scheduledClusterScanStatus"
	ClusterStatusFieldVersion                              = "version"
	ClusterStatusFieldVirtualCenterSecret                  = "virtualCenterSecret"
	ClusterStatusFieldVsphereSecret                        = "vsphereSecret"
	ClusterStatusFieldWeavePasswordSecret                  = "weavePasswordSecret"
)

type ClusterStatus struct {
	AADClientCertSecret                  string                      `json:"aadClientCertSecret,omitempty" yaml:"aadClientCertSecret,omitempty"`
	AADClientSecret                      string                      `json:"aadClientSecret,omitempty" yaml:"aadClientSecret,omitempty"`
	APIEndpoint                          string                      `json:"apiEndpoint,omitempty" yaml:"apiEndpoint,omitempty"`
	AgentFeatures                        map[string]bool             `json:"agentFeatures,omitempty" yaml:"agentFeatures,omitempty"`
	AgentImage                           string                      `json:"agentImage,omitempty" yaml:"agentImage,omitempty"`
	Allocatable                          map[string]string           `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	AppliedEnableNetworkPolicy           bool                        `json:"appliedEnableNetworkPolicy,omitempty" yaml:"appliedEnableNetworkPolicy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                      `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"appliedPodSecurityPolicyTemplateId,omitempty"`
	AppliedSpec                          *ClusterSpec                `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	AuthImage                            string                      `json:"authImage,omitempty" yaml:"authImage,omitempty"`
	CACert                               string                      `json:"caCert,omitempty" yaml:"caCert,omitempty"`
	Capabilities                         *Capabilities               `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Capacity                             map[string]string           `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CertificatesExpiration               map[string]CertExpiration   `json:"certificatesExpiration,omitempty" yaml:"certificatesExpiration,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus    `json:"componentStatuses,omitempty" yaml:"componentStatuses,omitempty"`
	Conditions                           []ClusterCondition          `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	CurrentCisRunName                    string                      `json:"currentCisRunName,omitempty" yaml:"currentCisRunName,omitempty"`
	Driver                               string                      `json:"driver,omitempty" yaml:"driver,omitempty"`
	FailedSpec                           *ClusterSpec                `json:"failedSpec,omitempty" yaml:"failedSpec,omitempty"`
	IstioEnabled                         bool                        `json:"istioEnabled,omitempty" yaml:"istioEnabled,omitempty"`
	Limits                               map[string]string           `json:"limits,omitempty" yaml:"limits,omitempty"`
	MonitoringStatus                     *MonitoringStatus           `json:"monitoringStatus,omitempty" yaml:"monitoringStatus,omitempty"`
	NodeCount                            int64                       `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	NodeVersion                          int64                       `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`
	OpenStackSecret                      string                      `json:"openStackSecret,omitempty" yaml:"openStackSecret,omitempty"`
	PrivateRegistrySecret                string                      `json:"privateRegistrySecret,omitempty" yaml:"privateRegistrySecret,omitempty"`
	Requested                            map[string]string           `json:"requested,omitempty" yaml:"requested,omitempty"`
	S3CredentialSecret                   string                      `json:"s3CredentialSecret,omitempty" yaml:"s3CredentialSecret,omitempty"`
	ScheduledClusterScanStatus           *ScheduledClusterScanStatus `json:"scheduledClusterScanStatus,omitempty" yaml:"scheduledClusterScanStatus,omitempty"`
	Version                              *Info                       `json:"version,omitempty" yaml:"version,omitempty"`
	VirtualCenterSecret                  string                      `json:"virtualCenterSecret,omitempty" yaml:"virtualCenterSecret,omitempty"`
	VsphereSecret                        string                      `json:"vsphereSecret,omitempty" yaml:"vsphereSecret,omitempty"`
	WeavePasswordSecret                  string                      `json:"weavePasswordSecret,omitempty" yaml:"weavePasswordSecret,omitempty"`
}
