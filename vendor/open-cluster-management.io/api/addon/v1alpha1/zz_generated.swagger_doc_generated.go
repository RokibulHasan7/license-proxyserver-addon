package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AddOnDeploymentConfig = map[string]string{
	"":     "AddOnDeploymentConfig represents a deployment configuration for an add-on.",
	"spec": "spec represents a desired configuration for an add-on.",
}

func (AddOnDeploymentConfig) SwaggerDoc() map[string]string {
	return map_AddOnDeploymentConfig
}

var map_AddOnDeploymentConfigList = map[string]string{
	"":         "AddOnDeploymentConfigList is a collection of add-on deployment config.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of add-on deployment config.",
}

func (AddOnDeploymentConfigList) SwaggerDoc() map[string]string {
	return map_AddOnDeploymentConfigList
}

var map_AddOnDeploymentConfigSpec = map[string]string{
	"customizedVariables":   "CustomizedVariables is a list of name-value variables for the current add-on deployment. The add-on implementation can use these variables to render its add-on deployment. The default is an empty list.",
	"nodePlacement":         "NodePlacement enables explicit control over the scheduling of the add-on agents on the managed cluster. All add-on agent pods are expected to comply with this node placement. If the placement is nil, the placement is not specified, it will be omitted. If the placement is an empty object, the placement will match all nodes and tolerate nothing.",
	"registries":            "Registries describes how to override images used by the addon agent on the managed cluster. the following example will override image \"quay.io/open-cluster-management/addon-agent\" to \"quay.io/ocm/addon-agent\" when deploying the addon agent\n\nregistries:\n  - source: quay.io/open-cluster-management/addon-agent\n    mirror: quay.io/ocm/addon-agent",
	"proxyConfig":           "ProxyConfig holds proxy settings for add-on agent on the managed cluster. Empty means no proxy settings is available.",
	"agentInstallNamespace": "AgentInstallNamespace is the namespace where the add-on agent should be installed on the managed cluster.",
}

func (AddOnDeploymentConfigSpec) SwaggerDoc() map[string]string {
	return map_AddOnDeploymentConfigSpec
}

var map_CustomizedVariable = map[string]string{
	"":      "CustomizedVariable represents a customized variable for add-on deployment.",
	"name":  "Name of this variable.",
	"value": "Value of this variable.",
}

func (CustomizedVariable) SwaggerDoc() map[string]string {
	return map_CustomizedVariable
}

var map_ImageMirror = map[string]string{
	"":       "ImageMirror describes how to mirror images from a source",
	"mirror": "Mirror is the mirrored registry of the Source. Will be ignored if Mirror is empty.",
	"source": "Source is the source registry. All image registries will be replaced by Mirror if Source is empty.",
}

func (ImageMirror) SwaggerDoc() map[string]string {
	return map_ImageMirror
}

var map_NodePlacement = map[string]string{
	"":             "NodePlacement describes node scheduling configuration for the pods.",
	"nodeSelector": "NodeSelector defines which Nodes the Pods are scheduled on. If the selector is an empty list, it will match all nodes. The default is an empty list.",
	"tolerations":  "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. If the tolerations is an empty list, it will tolerate nothing. The default is an empty list.",
}

func (NodePlacement) SwaggerDoc() map[string]string {
	return map_NodePlacement
}

var map_ProxyConfig = map[string]string{
	"":           "ProxyConfig describes the proxy settings for the add-on agent",
	"httpProxy":  "HTTPProxy is the URL of the proxy for HTTP requests",
	"httpsProxy": "HTTPSProxy is the URL of the proxy for HTTPS requests",
	"caBundle":   "CABundle is a CA certificate bundle to verify the proxy server. And it's only useful when HTTPSProxy is set and a HTTPS proxy server is specified.",
	"noProxy":    "NoProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used.",
}

func (ProxyConfig) SwaggerDoc() map[string]string {
	return map_ProxyConfig
}

var map_AddOnTemplate = map[string]string{
	"":     "AddOnTemplate is the Custom Resource object, it is used to describe how to deploy the addon agent and how to register the addon.\n\nAddOnTemplate is a cluster-scoped resource, and will only be used on the hub cluster.",
	"spec": "spec holds the registration configuration for the addon and the addon agent resources yaml description.",
}

func (AddOnTemplate) SwaggerDoc() map[string]string {
	return map_AddOnTemplate
}

var map_AddOnTemplateList = map[string]string{
	"":         "AddOnTemplateList is a collection of addon templates.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of addon templates.",
}

func (AddOnTemplateList) SwaggerDoc() map[string]string {
	return map_AddOnTemplateList
}

var map_AddOnTemplateSpec = map[string]string{
	"":             "AddOnTemplateSpec defines the template of an addon agent which will be deployed on managed clusters.",
	"addonName":    "AddonName represents the name of the addon which the template belongs to",
	"agentSpec":    "AgentSpec describes what/how the kubernetes resources of the addon agent to be deployed on a managed cluster.",
	"registration": "Registration holds the registration configuration for the addon",
}

func (AddOnTemplateSpec) SwaggerDoc() map[string]string {
	return map_AddOnTemplateSpec
}

var map_CurrentClusterBindingConfig = map[string]string{
	"clusterRoleName": "ClusterRoleName is the name of the clusterrole the addon agent is bound. A rolebinding will be created referring to this cluster role in each cluster namespace. The user must make sure the clusterrole exists on the hub cluster.",
}

func (CurrentClusterBindingConfig) SwaggerDoc() map[string]string {
	return map_CurrentClusterBindingConfig
}

var map_CustomSignerRegistrationConfig = map[string]string{
	"signerName": "signerName is the name of signer that addon agent will use to create csr.",
	"subject":    "Subject is the user subject of the addon agent to be registered to the hub. If it is not set, the addon agent will have the default subject \"subject\": {\n  \"user\": \"system:open-cluster-management:cluster:{clusterName}:addon:{addonName}:agent:{agentName}\",\n  \"groups: [\"system:open-cluster-management:cluster:{clusterName}:addon:{addonName}\",\n            \"system:open-cluster-management:addon:{addonName}\", \"system:authenticated\"]\n}",
	"signingCA":  "SigningCA represents the reference of the secret on the hub cluster to sign the CSR the secret must be in the namespace where the addon-manager is located, and the secret type must be \"kubernetes.io/tls\" Note: The addon manager will not have permission to access the secret by default, so the user must grant the permission to the addon manager(by creating rolebinding for the addon-manager serviceaccount \"addon-manager-controller-sa\").",
}

func (CustomSignerRegistrationConfig) SwaggerDoc() map[string]string {
	return map_CustomSignerRegistrationConfig
}

var map_HubPermissionConfig = map[string]string{
	"":                "HubPermissionConfig configures the permission of the addon agent to access the hub cluster. Will create a RoleBinding in the same namespace as the managedClusterAddon to bind the user provided ClusterRole/Role to the \"system:open-cluster-management:cluster:<cluster-name>:addon:<addon-name>\" Group.",
	"type":            "Type of the permissions setting. It defines how to bind the roleRef on the hub cluster. It can be: - CurrentCluster: Bind the roleRef to the namespace with the same name as the managedCluster. - SingleNamespace: Bind the roleRef to the namespace specified by SingleNamespaceBindingConfig.",
	"currentCluster":  "CurrentCluster contains the configuration of CurrentCluster type binding. It is required when the type is CurrentCluster.",
	"singleNamespace": "SingleNamespace contains the configuration of SingleNamespace type binding. It is required when the type is SingleNamespace",
}

func (HubPermissionConfig) SwaggerDoc() map[string]string {
	return map_HubPermissionConfig
}

var map_KubeClientRegistrationConfig = map[string]string{
	"hubPermissions": "HubPermissions represent the permission configurations of the addon agent to access the hub cluster",
}

func (KubeClientRegistrationConfig) SwaggerDoc() map[string]string {
	return map_KubeClientRegistrationConfig
}

var map_RegistrationSpec = map[string]string{
	"":             "RegistrationSpec describes how to register an addon agent to the hub cluster. With the registration defined, The addon agent can access to kube apiserver with kube style API or other endpoints on hub cluster with client certificate authentication. During the addon registration process, a csr will be created for each Registration on the hub cluster. The CSR will be approved automatically, After the csr is approved on the hub cluster, the klusterlet agent will create a secret in the installNamespace for the addon agent. If the RegistrationType type is KubeClient, the secret name will be \"{addon name}-hub-kubeconfig\" whose content includes key/cert and kubeconfig. Otherwise, If the RegistrationType type is CustomSigner the secret name will be \"{addon name}-{signer name}-client-cert\" whose content includes key/cert.",
	"type":         "Type of the registration configuration, it supports: - KubeClient: the addon agent can access the hub kube apiserver with kube style API.\n  the signer name should be \"kubernetes.io/kube-apiserver-client\". When this type is\n  used, the KubeClientRegistrationConfig can be used to define the permission of the\n  addon agent to access the hub cluster\n- CustomSigner: the addon agent can access the hub cluster through user-defined endpoints.\n  When this type is used, the CustomSignerRegistrationConfig can be used to define how\n  to issue the client certificate for the addon agent.",
	"kubeClient":   "KubeClient holds the configuration of the KubeClient type registration",
	"customSigner": "CustomSigner holds the configuration of the CustomSigner type registration required when the Type is CustomSigner",
}

func (RegistrationSpec) SwaggerDoc() map[string]string {
	return map_RegistrationSpec
}

var map_SigningCARef = map[string]string{
	"":     "SigningCARef is the reference to the signing CA secret which type must be \"kubernetes.io/tls\" and which namespace must be the same as the addon-manager.",
	"name": "Name of the signing CA secret",
}

func (SigningCARef) SwaggerDoc() map[string]string {
	return map_SigningCARef
}

var map_SingleNamespaceBindingConfig = map[string]string{
	"namespace": "Namespace is the namespace the addon agent has permissions to bind to. A rolebinding will be created in this namespace referring to the RoleRef.",
	"roleRef":   "RoleRef is an reference to the permission resource. it could be a role or a cluster role, the user must make sure it exist on the hub cluster.",
}

func (SingleNamespaceBindingConfig) SwaggerDoc() map[string]string {
	return map_SingleNamespaceBindingConfig
}

var map_AddOnMeta = map[string]string{
	"":            "AddOnMeta represents a collection of metadata information for the add-on.",
	"displayName": "displayName represents the name of add-on that will be displayed.",
	"description": "description represents the detailed description of the add-on.",
}

func (AddOnMeta) SwaggerDoc() map[string]string {
	return map_AddOnMeta
}

var map_ClusterManagementAddOn = map[string]string{
	"":       "ClusterManagementAddOn represents the registration of an add-on to the cluster manager. This resource allows the user to discover which add-on is available for the cluster manager and also provides metadata information about the add-on. This resource also provides a linkage to ManagedClusterAddOn, the name of the ClusterManagementAddOn resource will be used for the namespace-scoped ManagedClusterAddOn resource. ClusterManagementAddOn is a cluster-scoped resource.",
	"spec":   "spec represents a desired configuration for the agent on the cluster management add-on.",
	"status": "status represents the current status of cluster management add-on.",
}

func (ClusterManagementAddOn) SwaggerDoc() map[string]string {
	return map_ClusterManagementAddOn
}

var map_ClusterManagementAddOnList = map[string]string{
	"":         "ClusterManagementAddOnList is a collection of cluster management add-ons.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of cluster management add-ons.",
}

func (ClusterManagementAddOnList) SwaggerDoc() map[string]string {
	return map_ClusterManagementAddOnList
}

var map_ClusterManagementAddOnSpec = map[string]string{
	"":                   "ClusterManagementAddOnSpec provides information for the add-on.",
	"addOnMeta":          "addOnMeta is a reference to the metadata information for the add-on.",
	"addOnConfiguration": "Deprecated: Use supportedConfigs filed instead addOnConfiguration is a reference to configuration information for the add-on. In scenario where a multiple add-ons share the same add-on CRD, multiple ClusterManagementAddOn resources need to be created and reference the same AddOnConfiguration.",
	"supportedConfigs":   "supportedConfigs is a list of configuration types supported by add-on. An empty list means the add-on does not require configurations. The default is an empty list",
	"installStrategy":    "InstallStrategy represents that related ManagedClusterAddOns should be installed on certain clusters.",
}

func (ClusterManagementAddOnSpec) SwaggerDoc() map[string]string {
	return map_ClusterManagementAddOnSpec
}

var map_ClusterManagementAddOnStatus = map[string]string{
	"":                        "ClusterManagementAddOnStatus represents the current status of cluster management add-on.",
	"defaultconfigReferences": "defaultconfigReferences is a list of current add-on default configuration references.",
	"installProgressions":     "installProgression is a list of current add-on configuration references per placement.",
}

func (ClusterManagementAddOnStatus) SwaggerDoc() map[string]string {
	return map_ClusterManagementAddOnStatus
}

var map_ConfigCoordinates = map[string]string{
	"":                       "ConfigCoordinates represents the information for locating the CRD and CR that configures the add-on.",
	"crdName":                "crdName is the name of the CRD used to configure instances of the managed add-on. This field should be configured if the add-on have a CRD that controls the configuration of the add-on.",
	"crName":                 "crName is the name of the CR used to configure instances of the managed add-on. This field should be configured if add-on CR have a consistent name across the all of the ManagedCluster instaces.",
	"lastObservedGeneration": "lastObservedGeneration is the observed generation of the custom resource for the configuration of the addon.",
}

func (ConfigCoordinates) SwaggerDoc() map[string]string {
	return map_ConfigCoordinates
}

var map_ConfigGroupResource = map[string]string{
	"":         "ConfigGroupResource represents the GroupResource of the add-on configuration",
	"group":    "group of the add-on configuration.",
	"resource": "resource of the add-on configuration.",
}

func (ConfigGroupResource) SwaggerDoc() map[string]string {
	return map_ConfigGroupResource
}

var map_ConfigMeta = map[string]string{
	"":              "ConfigMeta represents a collection of metadata information for add-on configuration.",
	"defaultConfig": "defaultConfig represents the namespace and name of the default add-on configuration. In scenario where all add-ons have a same configuration.",
}

func (ConfigMeta) SwaggerDoc() map[string]string {
	return map_ConfigMeta
}

var map_ConfigReferent = map[string]string{
	"":          "ConfigReferent represents the namespace and name for an add-on configuration.",
	"namespace": "namespace of the add-on configuration. If this field is not set, the configuration is in the cluster scope.",
	"name":      "name of the add-on configuration.",
}

func (ConfigReferent) SwaggerDoc() map[string]string {
	return map_ConfigReferent
}

var map_ConfigSpecHash = map[string]string{
	"":         "ConfigSpecHash represents the namespace,name and spec hash for an add-on configuration.",
	"specHash": "spec hash for an add-on configuration.",
}

func (ConfigSpecHash) SwaggerDoc() map[string]string {
	return map_ConfigSpecHash
}

var map_DefaultConfigReference = map[string]string{
	"":              "DefaultConfigReference is a reference to the current add-on configuration. This resource is used to record the configuration resource for the current add-on.",
	"desiredConfig": "desiredConfig record the desired config spec hash.",
}

func (DefaultConfigReference) SwaggerDoc() map[string]string {
	return map_DefaultConfigReference
}

var map_InstallConfigReference = map[string]string{
	"":                    "InstallConfigReference is a reference to the current add-on configuration. This resource is used to record the configuration resource for the current add-on.",
	"desiredConfig":       "desiredConfig record the desired config name and spec hash.",
	"lastKnownGoodConfig": "lastKnownGoodConfig records the last known good config spec hash. For fresh install or rollout with type UpdateAll or RollingUpdate, the lastKnownGoodConfig is the same as lastAppliedConfig. For rollout with type RollingUpdateWithCanary, the lastKnownGoodConfig is the last successfully applied config spec hash of the canary placement.",
	"lastAppliedConfig":   "lastAppliedConfig records the config spec hash when the all the corresponding ManagedClusterAddOn are applied successfully.",
}

func (InstallConfigReference) SwaggerDoc() map[string]string {
	return map_InstallConfigReference
}

var map_InstallProgression = map[string]string{
	"configReferences": "configReferences is a list of current add-on configuration references.",
	"conditions":       "conditions describe the state of the managed and monitored components for the operator.",
}

func (InstallProgression) SwaggerDoc() map[string]string {
	return map_InstallProgression
}

var map_InstallStrategy = map[string]string{
	"":           "InstallStrategy represents that related ManagedClusterAddOns should be installed on certain clusters.",
	"type":       "Type is the type of the install strategy, it can be: - Manual: no automatic install - Placements: install to clusters selected by placements.",
	"placements": "Placements is a list of placement references honored when install strategy type is Placements. All clusters selected by these placements will install the addon If one cluster belongs to multiple placements, it will only apply the strategy defined later in the order. That is to say, The latter strategy overrides the previous one.",
}

func (InstallStrategy) SwaggerDoc() map[string]string {
	return map_InstallStrategy
}

var map_PlacementRef = map[string]string{
	"namespace": "Namespace is the namespace of the placement",
	"name":      "Name is the name of the placement",
}

func (PlacementRef) SwaggerDoc() map[string]string {
	return map_PlacementRef
}

var map_PlacementStrategy = map[string]string{
	"configs":         "Configs is the configuration of managedClusterAddon during installation. User can override the configuration by updating the managedClusterAddon directly.",
	"rolloutStrategy": "The rollout strategy to apply addon configurations change. The rollout strategy only watches the addon configurations defined in ClusterManagementAddOn.",
}

func (PlacementStrategy) SwaggerDoc() map[string]string {
	return map_PlacementStrategy
}

var map_ConfigReference = map[string]string{
	"":                       "ConfigReference is a reference to the current add-on configuration. This resource is used to locate the configuration resource for the current add-on.",
	"lastObservedGeneration": "Deprecated: Use LastAppliedConfig instead lastObservedGeneration is the observed generation of the add-on configuration.",
	"desiredConfig":          "desiredConfig record the desired config spec hash.",
	"lastAppliedConfig":      "lastAppliedConfig record the config spec hash when the corresponding ManifestWork is applied successfully.",
}

func (ConfigReference) SwaggerDoc() map[string]string {
	return map_ConfigReference
}

var map_HealthCheck = map[string]string{
	"mode": "mode indicates which mode will be used to check the healthiness status of the addon.",
}

func (HealthCheck) SwaggerDoc() map[string]string {
	return map_HealthCheck
}

var map_ManagedClusterAddOn = map[string]string{
	"":       "ManagedClusterAddOn is the Custom Resource object which holds the current state of an add-on. This object is used by add-on operators to convey their state. This resource should be created in the ManagedCluster namespace.",
	"spec":   "spec holds configuration that could apply to any operator.",
	"status": "status holds the information about the state of an operator.  It is consistent with status information across the Kubernetes ecosystem.",
}

func (ManagedClusterAddOn) SwaggerDoc() map[string]string {
	return map_ManagedClusterAddOn
}

var map_ManagedClusterAddOnList = map[string]string{
	"": "ManagedClusterAddOnList is a list of ManagedClusterAddOn resources.",
}

func (ManagedClusterAddOnList) SwaggerDoc() map[string]string {
	return map_ManagedClusterAddOnList
}

var map_ManagedClusterAddOnSpec = map[string]string{
	"":                 "ManagedClusterAddOnSpec defines the install configuration of an addon agent on managed cluster.",
	"installNamespace": "installNamespace is the namespace on the managed cluster to install the addon agent. If it is not set, open-cluster-management-agent-addon namespace is used to install the addon agent.",
	"configs":          "configs is a list of add-on configurations. In scenario where the current add-on has its own configurations. An empty list means there are no default configurations for add-on. The default is an empty list",
}

func (ManagedClusterAddOnSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterAddOnSpec
}

var map_ManagedClusterAddOnStatus = map[string]string{
	"":                   "ManagedClusterAddOnStatus provides information about the status of the operator.",
	"conditions":         "conditions describe the state of the managed and monitored components for the operator.",
	"relatedObjects":     "relatedObjects is a list of objects that are \"interesting\" or related to this operator. Common uses are: 1. the detailed resource driving the operator 2. operator namespaces 3. operand namespaces 4. related ClusterManagementAddon resource",
	"addOnMeta":          "addOnMeta is a reference to the metadata information for the add-on. This should be same as the addOnMeta for the corresponding ClusterManagementAddOn resource.",
	"addOnConfiguration": "Deprecated: Use configReferences instead. addOnConfiguration is a reference to configuration information for the add-on. This resource is used to locate the configuration resource for the add-on.",
	"supportedConfigs":   "SupportedConfigs is a list of configuration types that are allowed to override the add-on configurations defined in ClusterManagementAddOn spec. The default is an empty list, which means the add-on configurations can not be overridden.",
	"configReferences":   "configReferences is a list of current add-on configuration references. This will be overridden by the clustermanagementaddon configuration references.",
	"namespace":          "namespace is the namespace on the managedcluster to put registration secret or lease for the addon. It is required when registration is set or healthcheck mode is Lease.",
	"registrations":      "registrations is the configurations for the addon agent to register to hub. It should be set by each addon controller on hub to define how the addon agent on managedcluster is registered. With the registration defined, The addon agent can access to kube apiserver with kube style API or other endpoints on hub cluster with client certificate authentication. A csr will be created per registration configuration. If more than one registrationConfig is defined, a csr will be created for each registration configuration. It is not allowed that multiple registrationConfigs have the same signer name. After the csr is approved on the hub cluster, the klusterlet agent will create a secret in the installNamespace for the registrationConfig. If the signerName is \"kubernetes.io/kube-apiserver-client\", the secret name will be \"{addon name}-hub-kubeconfig\" whose contents includes key/cert and kubeconfig. Otherwise, the secret name will be \"{addon name}-{signer name}-client-cert\" whose contents includes key/cert.",
	"healthCheck":        "healthCheck indicates how to check the healthiness status of the current addon. It should be set by each addon implementation, by default, the lease mode will be used.",
}

func (ManagedClusterAddOnStatus) SwaggerDoc() map[string]string {
	return map_ManagedClusterAddOnStatus
}

var map_ObjectReference = map[string]string{
	"":          "ObjectReference contains enough information to let you inspect or modify the referred object.",
	"group":     "group of the referent.",
	"resource":  "resource of the referent.",
	"namespace": "namespace of the referent.",
	"name":      "name of the referent.",
}

func (ObjectReference) SwaggerDoc() map[string]string {
	return map_ObjectReference
}

var map_RegistrationConfig = map[string]string{
	"":           "RegistrationConfig defines the configuration of the addon agent to register to hub. The Klusterlet agent will create a csr for the addon agent with the registrationConfig.",
	"signerName": "signerName is the name of signer that addon agent will use to create csr.",
	"subject":    "subject is the user subject of the addon agent to be registered to the hub. If it is not set, the addon agent will have the default subject \"subject\": {\n  \"user\": \"system:open-cluster-management:cluster:{clusterName}:addon:{addonName}:agent:{agentName}\",\n  \"groups: [\"system:open-cluster-management:cluster:{clusterName}:addon:{addonName}\",\n            \"system:open-cluster-management:addon:{addonName}\", \"system:authenticated\"]\n}",
}

func (RegistrationConfig) SwaggerDoc() map[string]string {
	return map_RegistrationConfig
}

var map_Subject = map[string]string{
	"":                 "Subject is the user subject of the addon agent to be registered to the hub.",
	"user":             "user is the user name of the addon agent.",
	"groups":           "groups is the user group of the addon agent.",
	"organizationUnit": "organizationUnit is the ou of the addon agent",
}

func (Subject) SwaggerDoc() map[string]string {
	return map_Subject
}

// AUTO-GENERATED FUNCTIONS END HERE
