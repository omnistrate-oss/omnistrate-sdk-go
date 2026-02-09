# CreateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**AgentConfiguration** | Pointer to [**AgentConfiguration**](AgentConfiguration.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**BlobStorageConfiguration** | Pointer to [**BlobStorageConfiguration**](BlobStorageConfiguration.md) |  | [optional] 
**ContainerImagesRegistryCopyConfiguration** | Pointer to [**ContainerImagesRegistryCopyConfiguration**](ContainerImagesRegistryCopyConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Description** | **string** | A brief description of the resource | 
**Disable** | Pointer to **string** | Allows resource to be disabled. Expression returning true/false can be used as value. Resource is enabled by default, if this property is not set. | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**FileSystemConfiguration** | Pointer to [**FileSystemConfiguration**](FileSystemConfiguration.md) |  | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**ImageConfigId** | Pointer to **string** | ID of an Image Config | [optional] 
**InfraConfigId** | Pointer to **string** | ID of an Infra Config | [optional] 
**Internal** | Pointer to **bool** | Whether this resource is internal or not | [optional] [default to false]
**IsProxy** | Pointer to **bool** | Whether this resource is a proxy or not | [optional] [default to false]
**JobConfig** | Pointer to [**JobConfig**](JobConfig.md) |  | [optional] 
**Key** | Pointer to **string** | The key of the resource | [optional] 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the resource | 
**OnPremTerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for an OnPrem platform | [optional] 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProxyType** | Pointer to **string** | A proxy type of resource | [optional] [default to "PortsBasedProxy"]
**ResourceDependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource | [optional] 
**ServiceId** | **string** | ID of a Service | 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for cloud providers | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateResourceRequest

`func NewCreateResourceRequest(description string, name string, productTierId string, serviceId string, token string, ) *CreateResourceRequest`

NewCreateResourceRequest instantiates a new CreateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceRequestWithDefaults

`func NewCreateResourceRequestWithDefaults() *CreateResourceRequest`

NewCreateResourceRequestWithDefaults instantiates a new CreateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSecurityContext

`func (o *CreateResourceRequest) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *CreateResourceRequest) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *CreateResourceRequest) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *CreateResourceRequest) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetAgentConfiguration

`func (o *CreateResourceRequest) GetAgentConfiguration() AgentConfiguration`

GetAgentConfiguration returns the AgentConfiguration field if non-nil, zero value otherwise.

### GetAgentConfigurationOk

`func (o *CreateResourceRequest) GetAgentConfigurationOk() (*AgentConfiguration, bool)`

GetAgentConfigurationOk returns a tuple with the AgentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfiguration

`func (o *CreateResourceRequest) SetAgentConfiguration(v AgentConfiguration)`

SetAgentConfiguration sets AgentConfiguration field to given value.

### HasAgentConfiguration

`func (o *CreateResourceRequest) HasAgentConfiguration() bool`

HasAgentConfiguration returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *CreateResourceRequest) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *CreateResourceRequest) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *CreateResourceRequest) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *CreateResourceRequest) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetBlobStorageConfiguration

`func (o *CreateResourceRequest) GetBlobStorageConfiguration() BlobStorageConfiguration`

GetBlobStorageConfiguration returns the BlobStorageConfiguration field if non-nil, zero value otherwise.

### GetBlobStorageConfigurationOk

`func (o *CreateResourceRequest) GetBlobStorageConfigurationOk() (*BlobStorageConfiguration, bool)`

GetBlobStorageConfigurationOk returns a tuple with the BlobStorageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStorageConfiguration

`func (o *CreateResourceRequest) SetBlobStorageConfiguration(v BlobStorageConfiguration)`

SetBlobStorageConfiguration sets BlobStorageConfiguration field to given value.

### HasBlobStorageConfiguration

`func (o *CreateResourceRequest) HasBlobStorageConfiguration() bool`

HasBlobStorageConfiguration returns a boolean if a field has been set.

### GetContainerImagesRegistryCopyConfiguration

`func (o *CreateResourceRequest) GetContainerImagesRegistryCopyConfiguration() ContainerImagesRegistryCopyConfiguration`

GetContainerImagesRegistryCopyConfiguration returns the ContainerImagesRegistryCopyConfiguration field if non-nil, zero value otherwise.

### GetContainerImagesRegistryCopyConfigurationOk

`func (o *CreateResourceRequest) GetContainerImagesRegistryCopyConfigurationOk() (*ContainerImagesRegistryCopyConfiguration, bool)`

GetContainerImagesRegistryCopyConfigurationOk returns a tuple with the ContainerImagesRegistryCopyConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImagesRegistryCopyConfiguration

`func (o *CreateResourceRequest) SetContainerImagesRegistryCopyConfiguration(v ContainerImagesRegistryCopyConfiguration)`

SetContainerImagesRegistryCopyConfiguration sets ContainerImagesRegistryCopyConfiguration field to given value.

### HasContainerImagesRegistryCopyConfiguration

`func (o *CreateResourceRequest) HasContainerImagesRegistryCopyConfiguration() bool`

HasContainerImagesRegistryCopyConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *CreateResourceRequest) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *CreateResourceRequest) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *CreateResourceRequest) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *CreateResourceRequest) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *CreateResourceRequest) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *CreateResourceRequest) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *CreateResourceRequest) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *CreateResourceRequest) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *CreateResourceRequest) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *CreateResourceRequest) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *CreateResourceRequest) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *CreateResourceRequest) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDescription

`func (o *CreateResourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisable

`func (o *CreateResourceRequest) GetDisable() string`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *CreateResourceRequest) GetDisableOk() (*string, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *CreateResourceRequest) SetDisable(v string)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *CreateResourceRequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *CreateResourceRequest) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *CreateResourceRequest) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *CreateResourceRequest) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *CreateResourceRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetFileSystemConfiguration

`func (o *CreateResourceRequest) GetFileSystemConfiguration() FileSystemConfiguration`

GetFileSystemConfiguration returns the FileSystemConfiguration field if non-nil, zero value otherwise.

### GetFileSystemConfigurationOk

`func (o *CreateResourceRequest) GetFileSystemConfigurationOk() (*FileSystemConfiguration, bool)`

GetFileSystemConfigurationOk returns a tuple with the FileSystemConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemConfiguration

`func (o *CreateResourceRequest) SetFileSystemConfiguration(v FileSystemConfiguration)`

SetFileSystemConfiguration sets FileSystemConfiguration field to given value.

### HasFileSystemConfiguration

`func (o *CreateResourceRequest) HasFileSystemConfiguration() bool`

HasFileSystemConfiguration returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *CreateResourceRequest) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *CreateResourceRequest) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *CreateResourceRequest) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *CreateResourceRequest) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetImageConfigId

`func (o *CreateResourceRequest) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *CreateResourceRequest) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *CreateResourceRequest) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *CreateResourceRequest) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *CreateResourceRequest) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *CreateResourceRequest) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *CreateResourceRequest) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *CreateResourceRequest) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetInternal

`func (o *CreateResourceRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateResourceRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateResourceRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateResourceRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetIsProxy

`func (o *CreateResourceRequest) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *CreateResourceRequest) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *CreateResourceRequest) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *CreateResourceRequest) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetJobConfig

`func (o *CreateResourceRequest) GetJobConfig() JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *CreateResourceRequest) GetJobConfigOk() (*JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *CreateResourceRequest) SetJobConfig(v JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *CreateResourceRequest) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetKey

`func (o *CreateResourceRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateResourceRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateResourceRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateResourceRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *CreateResourceRequest) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *CreateResourceRequest) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *CreateResourceRequest) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *CreateResourceRequest) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *CreateResourceRequest) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *CreateResourceRequest) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *CreateResourceRequest) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *CreateResourceRequest) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *CreateResourceRequest) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *CreateResourceRequest) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *CreateResourceRequest) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *CreateResourceRequest) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *CreateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnPremTerraformConfigurations

`func (o *CreateResourceRequest) GetOnPremTerraformConfigurations() map[string]TerraformConfiguration`

GetOnPremTerraformConfigurations returns the OnPremTerraformConfigurations field if non-nil, zero value otherwise.

### GetOnPremTerraformConfigurationsOk

`func (o *CreateResourceRequest) GetOnPremTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetOnPremTerraformConfigurationsOk returns a tuple with the OnPremTerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremTerraformConfigurations

`func (o *CreateResourceRequest) SetOnPremTerraformConfigurations(v map[string]TerraformConfiguration)`

SetOnPremTerraformConfigurations sets OnPremTerraformConfigurations field to given value.

### HasOnPremTerraformConfigurations

`func (o *CreateResourceRequest) HasOnPremTerraformConfigurations() bool`

HasOnPremTerraformConfigurations returns a boolean if a field has been set.

### GetOperatorCRDConfiguration

`func (o *CreateResourceRequest) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *CreateResourceRequest) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *CreateResourceRequest) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *CreateResourceRequest) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetProductTierId

`func (o *CreateResourceRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *CreateResourceRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *CreateResourceRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProxyType

`func (o *CreateResourceRequest) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *CreateResourceRequest) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *CreateResourceRequest) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *CreateResourceRequest) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetResourceDependencies

`func (o *CreateResourceRequest) GetResourceDependencies() []ResourceDependency`

GetResourceDependencies returns the ResourceDependencies field if non-nil, zero value otherwise.

### GetResourceDependenciesOk

`func (o *CreateResourceRequest) GetResourceDependenciesOk() (*[]ResourceDependency, bool)`

GetResourceDependenciesOk returns a tuple with the ResourceDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDependencies

`func (o *CreateResourceRequest) SetResourceDependencies(v []ResourceDependency)`

SetResourceDependencies sets ResourceDependencies field to given value.

### HasResourceDependencies

`func (o *CreateResourceRequest) HasResourceDependencies() bool`

HasResourceDependencies returns a boolean if a field has been set.

### GetResourceType

`func (o *CreateResourceRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateResourceRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateResourceRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CreateResourceRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateResourceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateResourceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateResourceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTerraformConfigurations

`func (o *CreateResourceRequest) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *CreateResourceRequest) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *CreateResourceRequest) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *CreateResourceRequest) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.

### GetToken

`func (o *CreateResourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateResourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateResourceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


