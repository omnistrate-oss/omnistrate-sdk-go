# UpdateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**AgentConfiguration** | Pointer to [**AgentConfiguration**](AgentConfiguration.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**ContainerImagesRegistryCopyConfiguration** | Pointer to [**ContainerImagesRegistryCopyConfiguration**](ContainerImagesRegistryCopyConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Dependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the resource | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**Id** | **string** | ID of a resource | 
**ImageConfigId** | Pointer to **string** | ID of an Image Config | [optional] 
**InfraConfigId** | Pointer to **string** | ID of an Infra Config | [optional] 
**JobConfig** | Pointer to [**JobConfig**](JobConfig.md) |  | [optional] 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource | [optional] 
**OnPremTerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for an OnPrem platform | [optional] 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**ServiceId** | **string** | ID of a Service | 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for cloud providers | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateResourceRequest

`func NewUpdateResourceRequest(id string, serviceId string, token string, ) *UpdateResourceRequest`

NewUpdateResourceRequest instantiates a new UpdateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRequestWithDefaults

`func NewUpdateResourceRequestWithDefaults() *UpdateResourceRequest`

NewUpdateResourceRequestWithDefaults instantiates a new UpdateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSecurityContext

`func (o *UpdateResourceRequest) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *UpdateResourceRequest) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *UpdateResourceRequest) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *UpdateResourceRequest) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetAgentConfiguration

`func (o *UpdateResourceRequest) GetAgentConfiguration() AgentConfiguration`

GetAgentConfiguration returns the AgentConfiguration field if non-nil, zero value otherwise.

### GetAgentConfigurationOk

`func (o *UpdateResourceRequest) GetAgentConfigurationOk() (*AgentConfiguration, bool)`

GetAgentConfigurationOk returns a tuple with the AgentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfiguration

`func (o *UpdateResourceRequest) SetAgentConfiguration(v AgentConfiguration)`

SetAgentConfiguration sets AgentConfiguration field to given value.

### HasAgentConfiguration

`func (o *UpdateResourceRequest) HasAgentConfiguration() bool`

HasAgentConfiguration returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *UpdateResourceRequest) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *UpdateResourceRequest) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *UpdateResourceRequest) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *UpdateResourceRequest) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetContainerImagesRegistryCopyConfiguration

`func (o *UpdateResourceRequest) GetContainerImagesRegistryCopyConfiguration() ContainerImagesRegistryCopyConfiguration`

GetContainerImagesRegistryCopyConfiguration returns the ContainerImagesRegistryCopyConfiguration field if non-nil, zero value otherwise.

### GetContainerImagesRegistryCopyConfigurationOk

`func (o *UpdateResourceRequest) GetContainerImagesRegistryCopyConfigurationOk() (*ContainerImagesRegistryCopyConfiguration, bool)`

GetContainerImagesRegistryCopyConfigurationOk returns a tuple with the ContainerImagesRegistryCopyConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImagesRegistryCopyConfiguration

`func (o *UpdateResourceRequest) SetContainerImagesRegistryCopyConfiguration(v ContainerImagesRegistryCopyConfiguration)`

SetContainerImagesRegistryCopyConfiguration sets ContainerImagesRegistryCopyConfiguration field to given value.

### HasContainerImagesRegistryCopyConfiguration

`func (o *UpdateResourceRequest) HasContainerImagesRegistryCopyConfiguration() bool`

HasContainerImagesRegistryCopyConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *UpdateResourceRequest) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *UpdateResourceRequest) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *UpdateResourceRequest) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *UpdateResourceRequest) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *UpdateResourceRequest) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *UpdateResourceRequest) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *UpdateResourceRequest) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *UpdateResourceRequest) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *UpdateResourceRequest) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *UpdateResourceRequest) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *UpdateResourceRequest) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *UpdateResourceRequest) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDependencies

`func (o *UpdateResourceRequest) GetDependencies() []ResourceDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *UpdateResourceRequest) GetDependenciesOk() (*[]ResourceDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *UpdateResourceRequest) SetDependencies(v []ResourceDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *UpdateResourceRequest) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateResourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *UpdateResourceRequest) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateResourceRequest) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateResourceRequest) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateResourceRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *UpdateResourceRequest) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *UpdateResourceRequest) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *UpdateResourceRequest) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *UpdateResourceRequest) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetId

`func (o *UpdateResourceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateResourceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateResourceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetImageConfigId

`func (o *UpdateResourceRequest) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *UpdateResourceRequest) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *UpdateResourceRequest) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *UpdateResourceRequest) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *UpdateResourceRequest) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *UpdateResourceRequest) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *UpdateResourceRequest) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *UpdateResourceRequest) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetJobConfig

`func (o *UpdateResourceRequest) GetJobConfig() JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *UpdateResourceRequest) GetJobConfigOk() (*JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *UpdateResourceRequest) SetJobConfig(v JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *UpdateResourceRequest) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *UpdateResourceRequest) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *UpdateResourceRequest) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *UpdateResourceRequest) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *UpdateResourceRequest) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *UpdateResourceRequest) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *UpdateResourceRequest) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnPremTerraformConfigurations

`func (o *UpdateResourceRequest) GetOnPremTerraformConfigurations() map[string]TerraformConfiguration`

GetOnPremTerraformConfigurations returns the OnPremTerraformConfigurations field if non-nil, zero value otherwise.

### GetOnPremTerraformConfigurationsOk

`func (o *UpdateResourceRequest) GetOnPremTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetOnPremTerraformConfigurationsOk returns a tuple with the OnPremTerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremTerraformConfigurations

`func (o *UpdateResourceRequest) SetOnPremTerraformConfigurations(v map[string]TerraformConfiguration)`

SetOnPremTerraformConfigurations sets OnPremTerraformConfigurations field to given value.

### HasOnPremTerraformConfigurations

`func (o *UpdateResourceRequest) HasOnPremTerraformConfigurations() bool`

HasOnPremTerraformConfigurations returns a boolean if a field has been set.

### GetOperatorCRDConfiguration

`func (o *UpdateResourceRequest) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *UpdateResourceRequest) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *UpdateResourceRequest) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *UpdateResourceRequest) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateResourceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateResourceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateResourceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTerraformConfigurations

`func (o *UpdateResourceRequest) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *UpdateResourceRequest) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *UpdateResourceRequest) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *UpdateResourceRequest) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.

### GetToken

`func (o *UpdateResourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateResourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateResourceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


