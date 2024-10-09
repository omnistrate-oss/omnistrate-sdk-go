# UpdateResourceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Dependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the resource | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**ImageConfigId** | Pointer to **string** | The ID of the image configuration that this resource refers to | [optional] 
**InfraConfigId** | Pointer to **string** | The ID of the infrastructure configuration that this resource refers to | [optional] 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource | [optional] 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for various cloud providers | [optional] 

## Methods

### NewUpdateResourceRequestBody

`func NewUpdateResourceRequestBody() *UpdateResourceRequestBody`

NewUpdateResourceRequestBody instantiates a new UpdateResourceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRequestBodyWithDefaults

`func NewUpdateResourceRequestBodyWithDefaults() *UpdateResourceRequestBody`

NewUpdateResourceRequestBodyWithDefaults instantiates a new UpdateResourceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSecurityContext

`func (o *UpdateResourceRequestBody) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *UpdateResourceRequestBody) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *UpdateResourceRequestBody) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *UpdateResourceRequestBody) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *UpdateResourceRequestBody) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *UpdateResourceRequestBody) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *UpdateResourceRequestBody) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *UpdateResourceRequestBody) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *UpdateResourceRequestBody) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *UpdateResourceRequestBody) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *UpdateResourceRequestBody) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *UpdateResourceRequestBody) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *UpdateResourceRequestBody) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *UpdateResourceRequestBody) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *UpdateResourceRequestBody) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *UpdateResourceRequestBody) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *UpdateResourceRequestBody) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *UpdateResourceRequestBody) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *UpdateResourceRequestBody) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *UpdateResourceRequestBody) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDependencies

`func (o *UpdateResourceRequestBody) GetDependencies() []ResourceDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *UpdateResourceRequestBody) GetDependenciesOk() (*[]ResourceDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *UpdateResourceRequestBody) SetDependencies(v []ResourceDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *UpdateResourceRequestBody) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateResourceRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResourceRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResourceRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResourceRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *UpdateResourceRequestBody) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateResourceRequestBody) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateResourceRequestBody) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateResourceRequestBody) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *UpdateResourceRequestBody) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *UpdateResourceRequestBody) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *UpdateResourceRequestBody) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *UpdateResourceRequestBody) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetImageConfigId

`func (o *UpdateResourceRequestBody) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *UpdateResourceRequestBody) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *UpdateResourceRequestBody) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *UpdateResourceRequestBody) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *UpdateResourceRequestBody) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *UpdateResourceRequestBody) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *UpdateResourceRequestBody) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *UpdateResourceRequestBody) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *UpdateResourceRequestBody) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *UpdateResourceRequestBody) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *UpdateResourceRequestBody) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *UpdateResourceRequestBody) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *UpdateResourceRequestBody) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *UpdateResourceRequestBody) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *UpdateResourceRequestBody) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateResourceRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourceRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourceRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourceRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatorCRDConfiguration

`func (o *UpdateResourceRequestBody) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *UpdateResourceRequestBody) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *UpdateResourceRequestBody) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *UpdateResourceRequestBody) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetTerraformConfigurations

`func (o *UpdateResourceRequestBody) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *UpdateResourceRequestBody) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *UpdateResourceRequestBody) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *UpdateResourceRequestBody) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

