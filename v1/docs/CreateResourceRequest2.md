# CreateResourceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**BlobStorageConfiguration** | Pointer to [**BlobStorageConfiguration**](BlobStorageConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Description** | **string** | A brief description of the resource | 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**FileSystemConfiguration** | Pointer to [**FileSystemConfiguration**](FileSystemConfiguration.md) |  | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**ImageConfigId** | Pointer to **string** | The ID of the image configuration that this resource refers to | [optional] 
**InfraConfigId** | Pointer to **string** | The ID of the infrastructure configuration that this resource refers to | [optional] 
**Internal** | Pointer to **bool** | Whether this resource is internal or not | [optional] [default to false]
**IsProxy** | Pointer to **bool** | Whether this resource is a proxy or not | [optional] [default to false]
**JobConfig** | Pointer to [**JobConfig**](JobConfig.md) |  | [optional] 
**Key** | Pointer to **string** | The key of the resource | [optional] 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the resource | 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**ProductTierId** | **string** | The product tier ID | 
**ProxyType** | Pointer to **string** | A proxy type of resource | [optional] [default to "PortsBasedProxy"]
**ResourceDependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource | [optional] 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for various cloud providers | [optional] 

## Methods

### NewCreateResourceRequest2

`func NewCreateResourceRequest2(description string, name string, productTierId string, ) *CreateResourceRequest2`

NewCreateResourceRequest2 instantiates a new CreateResourceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceRequest2WithDefaults

`func NewCreateResourceRequest2WithDefaults() *CreateResourceRequest2`

NewCreateResourceRequest2WithDefaults instantiates a new CreateResourceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSecurityContext

`func (o *CreateResourceRequest2) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *CreateResourceRequest2) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *CreateResourceRequest2) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *CreateResourceRequest2) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *CreateResourceRequest2) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *CreateResourceRequest2) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *CreateResourceRequest2) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *CreateResourceRequest2) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetBlobStorageConfiguration

`func (o *CreateResourceRequest2) GetBlobStorageConfiguration() BlobStorageConfiguration`

GetBlobStorageConfiguration returns the BlobStorageConfiguration field if non-nil, zero value otherwise.

### GetBlobStorageConfigurationOk

`func (o *CreateResourceRequest2) GetBlobStorageConfigurationOk() (*BlobStorageConfiguration, bool)`

GetBlobStorageConfigurationOk returns a tuple with the BlobStorageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStorageConfiguration

`func (o *CreateResourceRequest2) SetBlobStorageConfiguration(v BlobStorageConfiguration)`

SetBlobStorageConfiguration sets BlobStorageConfiguration field to given value.

### HasBlobStorageConfiguration

`func (o *CreateResourceRequest2) HasBlobStorageConfiguration() bool`

HasBlobStorageConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *CreateResourceRequest2) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *CreateResourceRequest2) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *CreateResourceRequest2) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *CreateResourceRequest2) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *CreateResourceRequest2) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *CreateResourceRequest2) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *CreateResourceRequest2) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *CreateResourceRequest2) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *CreateResourceRequest2) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *CreateResourceRequest2) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *CreateResourceRequest2) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *CreateResourceRequest2) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDescription

`func (o *CreateResourceRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourceRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourceRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvironmentVariables

`func (o *CreateResourceRequest2) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *CreateResourceRequest2) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *CreateResourceRequest2) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *CreateResourceRequest2) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetFileSystemConfiguration

`func (o *CreateResourceRequest2) GetFileSystemConfiguration() FileSystemConfiguration`

GetFileSystemConfiguration returns the FileSystemConfiguration field if non-nil, zero value otherwise.

### GetFileSystemConfigurationOk

`func (o *CreateResourceRequest2) GetFileSystemConfigurationOk() (*FileSystemConfiguration, bool)`

GetFileSystemConfigurationOk returns a tuple with the FileSystemConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemConfiguration

`func (o *CreateResourceRequest2) SetFileSystemConfiguration(v FileSystemConfiguration)`

SetFileSystemConfiguration sets FileSystemConfiguration field to given value.

### HasFileSystemConfiguration

`func (o *CreateResourceRequest2) HasFileSystemConfiguration() bool`

HasFileSystemConfiguration returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *CreateResourceRequest2) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *CreateResourceRequest2) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *CreateResourceRequest2) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *CreateResourceRequest2) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetImageConfigId

`func (o *CreateResourceRequest2) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *CreateResourceRequest2) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *CreateResourceRequest2) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *CreateResourceRequest2) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *CreateResourceRequest2) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *CreateResourceRequest2) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *CreateResourceRequest2) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *CreateResourceRequest2) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetInternal

`func (o *CreateResourceRequest2) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateResourceRequest2) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateResourceRequest2) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateResourceRequest2) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetIsProxy

`func (o *CreateResourceRequest2) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *CreateResourceRequest2) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *CreateResourceRequest2) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *CreateResourceRequest2) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetJobConfig

`func (o *CreateResourceRequest2) GetJobConfig() JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *CreateResourceRequest2) GetJobConfigOk() (*JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *CreateResourceRequest2) SetJobConfig(v JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *CreateResourceRequest2) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetKey

`func (o *CreateResourceRequest2) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateResourceRequest2) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateResourceRequest2) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateResourceRequest2) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *CreateResourceRequest2) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *CreateResourceRequest2) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *CreateResourceRequest2) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *CreateResourceRequest2) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *CreateResourceRequest2) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *CreateResourceRequest2) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *CreateResourceRequest2) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *CreateResourceRequest2) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *CreateResourceRequest2) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *CreateResourceRequest2) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *CreateResourceRequest2) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *CreateResourceRequest2) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *CreateResourceRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorCRDConfiguration

`func (o *CreateResourceRequest2) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *CreateResourceRequest2) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *CreateResourceRequest2) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *CreateResourceRequest2) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetProductTierId

`func (o *CreateResourceRequest2) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *CreateResourceRequest2) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *CreateResourceRequest2) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProxyType

`func (o *CreateResourceRequest2) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *CreateResourceRequest2) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *CreateResourceRequest2) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *CreateResourceRequest2) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetResourceDependencies

`func (o *CreateResourceRequest2) GetResourceDependencies() []ResourceDependency`

GetResourceDependencies returns the ResourceDependencies field if non-nil, zero value otherwise.

### GetResourceDependenciesOk

`func (o *CreateResourceRequest2) GetResourceDependenciesOk() (*[]ResourceDependency, bool)`

GetResourceDependenciesOk returns a tuple with the ResourceDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDependencies

`func (o *CreateResourceRequest2) SetResourceDependencies(v []ResourceDependency)`

SetResourceDependencies sets ResourceDependencies field to given value.

### HasResourceDependencies

`func (o *CreateResourceRequest2) HasResourceDependencies() bool`

HasResourceDependencies returns a boolean if a field has been set.

### GetResourceType

`func (o *CreateResourceRequest2) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateResourceRequest2) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateResourceRequest2) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CreateResourceRequest2) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTerraformConfigurations

`func (o *CreateResourceRequest2) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *CreateResourceRequest2) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *CreateResourceRequest2) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *CreateResourceRequest2) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


