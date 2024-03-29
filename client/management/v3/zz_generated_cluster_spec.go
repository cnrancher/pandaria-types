package client

const (
	ClusterSpecType                                     = "clusterSpec"
	ClusterSpecFieldAgentImageOverride                  = "agentImageOverride"
	ClusterSpecFieldAmazonElasticContainerServiceConfig = "amazonElasticContainerServiceConfig"
	ClusterSpecFieldAzureKubernetesServiceConfig        = "azureKubernetesServiceConfig"
	ClusterSpecFieldClusterTemplateAnswers              = "answers"
	ClusterSpecFieldClusterTemplateID                   = "clusterTemplateId"
	ClusterSpecFieldClusterTemplateQuestions            = "questions"
	ClusterSpecFieldClusterTemplateRevisionID           = "clusterTemplateRevisionId"
	ClusterSpecFieldDefaultClusterRoleForProjectMembers = "defaultClusterRoleForProjectMembers"
	ClusterSpecFieldDefaultPodSecurityPolicyTemplateID  = "defaultPodSecurityPolicyTemplateId"
	ClusterSpecFieldDescription                         = "description"
	ClusterSpecFieldDesiredAgentImage                   = "desiredAgentImage"
	ClusterSpecFieldDesiredAuthImage                    = "desiredAuthImage"
	ClusterSpecFieldDisplayName                         = "displayName"
	ClusterSpecFieldDockerRootDir                       = "dockerRootDir"
	ClusterSpecFieldEnableClusterAlerting               = "enableClusterAlerting"
	ClusterSpecFieldEnableClusterMonitoring             = "enableClusterMonitoring"
	ClusterSpecFieldEnableDualStack                     = "enableDualStack"
	ClusterSpecFieldEnableF5CIS                         = "enableF5CIS"
	ClusterSpecFieldEnableGPUManagement                 = "enableGPUManagement"
	ClusterSpecFieldEnableNetworkPolicy                 = "enableNetworkPolicy"
	ClusterSpecFieldFluentdLogDir                       = "fluentdLogDir"
	ClusterSpecFieldGPUSchedulerNodePort                = "gpuSchedulerNodePort"
	ClusterSpecFieldGenericEngineConfig                 = "genericEngineConfig"
	ClusterSpecFieldGoogleKubernetesEngineConfig        = "googleKubernetesEngineConfig"
	ClusterSpecFieldImportedConfig                      = "importedConfig"
	ClusterSpecFieldInternal                            = "internal"
	ClusterSpecFieldK3sConfig                           = "k3sConfig"
	ClusterSpecFieldLocalClusterAuthEndpoint            = "localClusterAuthEndpoint"
	ClusterSpecFieldRancherKubernetesEngineConfig       = "rancherKubernetesEngineConfig"
	ClusterSpecFieldScheduledClusterScan                = "scheduledClusterScan"
	ClusterSpecFieldSystemDefaultRegistry               = "systemDefaultRegistry"
	ClusterSpecFieldWindowsPreferedCluster              = "windowsPreferedCluster"
)

type ClusterSpec struct {
	AgentImageOverride                  string                         `json:"agentImageOverride,omitempty" yaml:"agentImageOverride,omitempty"`
	AmazonElasticContainerServiceConfig map[string]interface{}         `json:"amazonElasticContainerServiceConfig,omitempty" yaml:"amazonElasticContainerServiceConfig,omitempty"`
	AzureKubernetesServiceConfig        map[string]interface{}         `json:"azureKubernetesServiceConfig,omitempty" yaml:"azureKubernetesServiceConfig,omitempty"`
	ClusterTemplateAnswers              *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                   string                         `json:"clusterTemplateId,omitempty" yaml:"clusterTemplateId,omitempty"`
	ClusterTemplateQuestions            []Question                     `json:"questions,omitempty" yaml:"questions,omitempty"`
	ClusterTemplateRevisionID           string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"clusterTemplateRevisionId,omitempty"`
	DefaultClusterRoleForProjectMembers string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityPolicyTemplateID  string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"defaultPodSecurityPolicyTemplateId,omitempty"`
	Description                         string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                   string                         `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DesiredAuthImage                    string                         `json:"desiredAuthImage,omitempty" yaml:"desiredAuthImage,omitempty"`
	DisplayName                         string                         `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	DockerRootDir                       string                         `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	EnableClusterAlerting               bool                           `json:"enableClusterAlerting,omitempty" yaml:"enableClusterAlerting,omitempty"`
	EnableClusterMonitoring             bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enableClusterMonitoring,omitempty"`
	EnableDualStack                     bool                           `json:"enableDualStack,omitempty" yaml:"enableDualStack,omitempty"`
	EnableF5CIS                         bool                           `json:"enableF5CIS,omitempty" yaml:"enableF5CIS,omitempty"`
	EnableGPUManagement                 bool                           `json:"enableGPUManagement,omitempty" yaml:"enableGPUManagement,omitempty"`
	EnableNetworkPolicy                 *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FluentdLogDir                       string                         `json:"fluentdLogDir,omitempty" yaml:"fluentdLogDir,omitempty"`
	GPUSchedulerNodePort                string                         `json:"gpuSchedulerNodePort,omitempty" yaml:"gpuSchedulerNodePort,omitempty"`
	GenericEngineConfig                 map[string]interface{}         `json:"genericEngineConfig,omitempty" yaml:"genericEngineConfig,omitempty"`
	GoogleKubernetesEngineConfig        map[string]interface{}         `json:"googleKubernetesEngineConfig,omitempty" yaml:"googleKubernetesEngineConfig,omitempty"`
	ImportedConfig                      *ImportedConfig                `json:"importedConfig,omitempty" yaml:"importedConfig,omitempty"`
	Internal                            bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	K3sConfig                           *K3sConfig                     `json:"k3sConfig,omitempty" yaml:"k3sConfig,omitempty"`
	LocalClusterAuthEndpoint            *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"localClusterAuthEndpoint,omitempty"`
	RancherKubernetesEngineConfig       *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	ScheduledClusterScan                *ScheduledClusterScan          `json:"scheduledClusterScan,omitempty" yaml:"scheduledClusterScan,omitempty"`
	SystemDefaultRegistry               string                         `json:"systemDefaultRegistry,omitempty" yaml:"systemDefaultRegistry,omitempty"`
	WindowsPreferedCluster              bool                           `json:"windowsPreferedCluster,omitempty" yaml:"windowsPreferedCluster,omitempty"`
}
