# DescribeResourceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionHooks** | Pointer to [**[]ActionHook**](ActionHook.md) | The action hooks that this resource supports | [optional] 
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**AgentConfiguration** | Pointer to [**AgentConfiguration**](AgentConfiguration.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**BlobStorageConfiguration** | Pointer to [**BlobStorageConfiguration**](BlobStorageConfiguration.md) |  | [optional] 
**Capabilities** | Pointer to [**[]ResourceCapability**](ResourceCapability.md) | The capabilities enabled for the resource | [optional] 
**ContainerImagesRegistryCopyConfiguration** | Pointer to [**ContainerImagesRegistryCopyConfiguration**](ContainerImagesRegistryCopyConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Dependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**Description** | **string** | A brief description of the resource | 
**Disable** | Pointer to **string** | Allows resource to be disabled. Expression returning true/false can be used as value. Resource is enabled by default, if this property is not set. | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**FileSystemConfiguration** | Pointer to [**FileSystemConfiguration**](FileSystemConfiguration.md) |  | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**Id** | **string** | ID of a resource | 
**ImageConfigId** | Pointer to **string** | ID of an Image Config | [optional] 
**InfraConfigId** | Pointer to **string** | ID of an Infra Config | [optional] 
**Internal** | **bool** | Whether this resource is internal or not | [default to false]
**IsDeprecated** | **bool** | Whether this resource is deprecated or not | [default to false]
**JobConfig** | Pointer to [**JobConfig**](JobConfig.md) |  | [optional] 
**Key** | **string** | The key of the resource | 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the resource | 
**OnPremTerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for an OnPrem platform | [optional] 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProxyType** | Pointer to **string** | The proxy type of instance | [optional] 
**ResourceType** | **string** | The type of the resource | 
**ServiceId** | **string** | ID of a Service | 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for cloud providers | [optional] 

## Methods

### NewDescribeResourceResult

`func NewDescribeResourceResult(description string, id string, internal bool, isDeprecated bool, key string, name string, productTierId string, resourceType string, serviceId string, ) *DescribeResourceResult`

NewDescribeResourceResult instantiates a new DescribeResourceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceResultWithDefaults

`func NewDescribeResourceResultWithDefaults() *DescribeResourceResult`

NewDescribeResourceResultWithDefaults instantiates a new DescribeResourceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionHooks

`func (o *DescribeResourceResult) GetActionHooks() []ActionHook`

GetActionHooks returns the ActionHooks field if non-nil, zero value otherwise.

### GetActionHooksOk

`func (o *DescribeResourceResult) GetActionHooksOk() (*[]ActionHook, bool)`

GetActionHooksOk returns a tuple with the ActionHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHooks

`func (o *DescribeResourceResult) SetActionHooks(v []ActionHook)`

SetActionHooks sets ActionHooks field to given value.

### HasActionHooks

`func (o *DescribeResourceResult) HasActionHooks() bool`

HasActionHooks returns a boolean if a field has been set.

### GetAdditionalSecurityContext

`func (o *DescribeResourceResult) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *DescribeResourceResult) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *DescribeResourceResult) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *DescribeResourceResult) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetAgentConfiguration

`func (o *DescribeResourceResult) GetAgentConfiguration() AgentConfiguration`

GetAgentConfiguration returns the AgentConfiguration field if non-nil, zero value otherwise.

### GetAgentConfigurationOk

`func (o *DescribeResourceResult) GetAgentConfigurationOk() (*AgentConfiguration, bool)`

GetAgentConfigurationOk returns a tuple with the AgentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfiguration

`func (o *DescribeResourceResult) SetAgentConfiguration(v AgentConfiguration)`

SetAgentConfiguration sets AgentConfiguration field to given value.

### HasAgentConfiguration

`func (o *DescribeResourceResult) HasAgentConfiguration() bool`

HasAgentConfiguration returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *DescribeResourceResult) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *DescribeResourceResult) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *DescribeResourceResult) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *DescribeResourceResult) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetBlobStorageConfiguration

`func (o *DescribeResourceResult) GetBlobStorageConfiguration() BlobStorageConfiguration`

GetBlobStorageConfiguration returns the BlobStorageConfiguration field if non-nil, zero value otherwise.

### GetBlobStorageConfigurationOk

`func (o *DescribeResourceResult) GetBlobStorageConfigurationOk() (*BlobStorageConfiguration, bool)`

GetBlobStorageConfigurationOk returns a tuple with the BlobStorageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStorageConfiguration

`func (o *DescribeResourceResult) SetBlobStorageConfiguration(v BlobStorageConfiguration)`

SetBlobStorageConfiguration sets BlobStorageConfiguration field to given value.

### HasBlobStorageConfiguration

`func (o *DescribeResourceResult) HasBlobStorageConfiguration() bool`

HasBlobStorageConfiguration returns a boolean if a field has been set.

### GetCapabilities

`func (o *DescribeResourceResult) GetCapabilities() []ResourceCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DescribeResourceResult) GetCapabilitiesOk() (*[]ResourceCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DescribeResourceResult) SetCapabilities(v []ResourceCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DescribeResourceResult) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetContainerImagesRegistryCopyConfiguration

`func (o *DescribeResourceResult) GetContainerImagesRegistryCopyConfiguration() ContainerImagesRegistryCopyConfiguration`

GetContainerImagesRegistryCopyConfiguration returns the ContainerImagesRegistryCopyConfiguration field if non-nil, zero value otherwise.

### GetContainerImagesRegistryCopyConfigurationOk

`func (o *DescribeResourceResult) GetContainerImagesRegistryCopyConfigurationOk() (*ContainerImagesRegistryCopyConfiguration, bool)`

GetContainerImagesRegistryCopyConfigurationOk returns a tuple with the ContainerImagesRegistryCopyConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImagesRegistryCopyConfiguration

`func (o *DescribeResourceResult) SetContainerImagesRegistryCopyConfiguration(v ContainerImagesRegistryCopyConfiguration)`

SetContainerImagesRegistryCopyConfiguration sets ContainerImagesRegistryCopyConfiguration field to given value.

### HasContainerImagesRegistryCopyConfiguration

`func (o *DescribeResourceResult) HasContainerImagesRegistryCopyConfiguration() bool`

HasContainerImagesRegistryCopyConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *DescribeResourceResult) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *DescribeResourceResult) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *DescribeResourceResult) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *DescribeResourceResult) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *DescribeResourceResult) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *DescribeResourceResult) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *DescribeResourceResult) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *DescribeResourceResult) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *DescribeResourceResult) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *DescribeResourceResult) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *DescribeResourceResult) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *DescribeResourceResult) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDependencies

`func (o *DescribeResourceResult) GetDependencies() []ResourceDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *DescribeResourceResult) GetDependenciesOk() (*[]ResourceDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *DescribeResourceResult) SetDependencies(v []ResourceDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *DescribeResourceResult) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeResourceResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeResourceResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeResourceResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisable

`func (o *DescribeResourceResult) GetDisable() string`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DescribeResourceResult) GetDisableOk() (*string, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DescribeResourceResult) SetDisable(v string)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DescribeResourceResult) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *DescribeResourceResult) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *DescribeResourceResult) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *DescribeResourceResult) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *DescribeResourceResult) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetFileSystemConfiguration

`func (o *DescribeResourceResult) GetFileSystemConfiguration() FileSystemConfiguration`

GetFileSystemConfiguration returns the FileSystemConfiguration field if non-nil, zero value otherwise.

### GetFileSystemConfigurationOk

`func (o *DescribeResourceResult) GetFileSystemConfigurationOk() (*FileSystemConfiguration, bool)`

GetFileSystemConfigurationOk returns a tuple with the FileSystemConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemConfiguration

`func (o *DescribeResourceResult) SetFileSystemConfiguration(v FileSystemConfiguration)`

SetFileSystemConfiguration sets FileSystemConfiguration field to given value.

### HasFileSystemConfiguration

`func (o *DescribeResourceResult) HasFileSystemConfiguration() bool`

HasFileSystemConfiguration returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *DescribeResourceResult) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *DescribeResourceResult) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *DescribeResourceResult) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *DescribeResourceResult) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetId

`func (o *DescribeResourceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeResourceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeResourceResult) SetId(v string)`

SetId sets Id field to given value.


### GetImageConfigId

`func (o *DescribeResourceResult) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *DescribeResourceResult) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *DescribeResourceResult) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *DescribeResourceResult) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *DescribeResourceResult) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *DescribeResourceResult) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *DescribeResourceResult) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *DescribeResourceResult) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetInternal

`func (o *DescribeResourceResult) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *DescribeResourceResult) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *DescribeResourceResult) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetIsDeprecated

`func (o *DescribeResourceResult) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *DescribeResourceResult) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *DescribeResourceResult) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetJobConfig

`func (o *DescribeResourceResult) GetJobConfig() JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *DescribeResourceResult) GetJobConfigOk() (*JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *DescribeResourceResult) SetJobConfig(v JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *DescribeResourceResult) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetKey

`func (o *DescribeResourceResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeResourceResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeResourceResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetKustomizeConfiguration

`func (o *DescribeResourceResult) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *DescribeResourceResult) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *DescribeResourceResult) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *DescribeResourceResult) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *DescribeResourceResult) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *DescribeResourceResult) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *DescribeResourceResult) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *DescribeResourceResult) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *DescribeResourceResult) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *DescribeResourceResult) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *DescribeResourceResult) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *DescribeResourceResult) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *DescribeResourceResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeResourceResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeResourceResult) SetName(v string)`

SetName sets Name field to given value.


### GetOnPremTerraformConfigurations

`func (o *DescribeResourceResult) GetOnPremTerraformConfigurations() map[string]TerraformConfiguration`

GetOnPremTerraformConfigurations returns the OnPremTerraformConfigurations field if non-nil, zero value otherwise.

### GetOnPremTerraformConfigurationsOk

`func (o *DescribeResourceResult) GetOnPremTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetOnPremTerraformConfigurationsOk returns a tuple with the OnPremTerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremTerraformConfigurations

`func (o *DescribeResourceResult) SetOnPremTerraformConfigurations(v map[string]TerraformConfiguration)`

SetOnPremTerraformConfigurations sets OnPremTerraformConfigurations field to given value.

### HasOnPremTerraformConfigurations

`func (o *DescribeResourceResult) HasOnPremTerraformConfigurations() bool`

HasOnPremTerraformConfigurations returns a boolean if a field has been set.

### GetOperatorCRDConfiguration

`func (o *DescribeResourceResult) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *DescribeResourceResult) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *DescribeResourceResult) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *DescribeResourceResult) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetProductTierId

`func (o *DescribeResourceResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeResourceResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeResourceResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProxyType

`func (o *DescribeResourceResult) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *DescribeResourceResult) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *DescribeResourceResult) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *DescribeResourceResult) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetResourceType

`func (o *DescribeResourceResult) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DescribeResourceResult) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DescribeResourceResult) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetServiceId

`func (o *DescribeResourceResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeResourceResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeResourceResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTerraformConfigurations

`func (o *DescribeResourceResult) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *DescribeResourceResult) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *DescribeResourceResult) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *DescribeResourceResult) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


