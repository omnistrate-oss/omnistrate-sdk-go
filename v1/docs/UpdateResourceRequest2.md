# UpdateResourceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalSecurityContext** | Pointer to [**AdditionalSecurityContext**](AdditionalSecurityContext.md) |  | [optional] 
**AgentConfiguration** | Pointer to [**AgentConfiguration**](AgentConfiguration.md) |  | [optional] 
**BackupConfiguration** | Pointer to [**BackupConfiguration**](BackupConfiguration.md) |  | [optional] 
**ContainerImageConfiguration** | Pointer to [**ContainerImageConfiguration**](ContainerImageConfiguration.md) |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** | Custom labels for the resource | [optional] 
**CustomSysCTLs** | Pointer to **map[string]string** | Custom sysctl settings for the resource | [optional] 
**CustomULimits** | Pointer to [**[]CustomULimits**](CustomULimits.md) | Custom ulimits for the resource | [optional] 
**Dependencies** | Pointer to [**[]ResourceDependency**](ResourceDependency.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the resource | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 
**HelmChartConfiguration** | Pointer to [**HelmChartConfiguration**](HelmChartConfiguration.md) |  | [optional] 
**ImageConfigId** | Pointer to **string** | The ID of the image configuration that this resource refers to | [optional] 
**InfraConfigId** | Pointer to **string** | The ID of the infrastructure configuration that this resource refers to | [optional] 
**JobConfig** | Pointer to [**JobConfig**](JobConfig.md) |  | [optional] 
**KustomizeConfiguration** | Pointer to [**KustomizeConfiguration**](KustomizeConfiguration.md) |  | [optional] 
**L4LoadBalancerConfiguration** | Pointer to [**L4LoadBalancerConfiguration**](L4LoadBalancerConfiguration.md) |  | [optional] 
**L7LoadBalancerConfiguration** | Pointer to [**L7LoadBalancerConfiguration**](L7LoadBalancerConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource | [optional] 
**OnPremTerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for various OnPrem platforms | [optional] 
**OperatorCRDConfiguration** | Pointer to [**OperatorCRDConfiguration**](OperatorCRDConfiguration.md) |  | [optional] 
**TerraformConfigurations** | Pointer to [**map[string]TerraformConfiguration**](TerraformConfiguration.md) | The Terraform configurations for various cloud providers | [optional] 

## Methods

### NewUpdateResourceRequest2

`func NewUpdateResourceRequest2() *UpdateResourceRequest2`

NewUpdateResourceRequest2 instantiates a new UpdateResourceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRequest2WithDefaults

`func NewUpdateResourceRequest2WithDefaults() *UpdateResourceRequest2`

NewUpdateResourceRequest2WithDefaults instantiates a new UpdateResourceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalSecurityContext

`func (o *UpdateResourceRequest2) GetAdditionalSecurityContext() AdditionalSecurityContext`

GetAdditionalSecurityContext returns the AdditionalSecurityContext field if non-nil, zero value otherwise.

### GetAdditionalSecurityContextOk

`func (o *UpdateResourceRequest2) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool)`

GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSecurityContext

`func (o *UpdateResourceRequest2) SetAdditionalSecurityContext(v AdditionalSecurityContext)`

SetAdditionalSecurityContext sets AdditionalSecurityContext field to given value.

### HasAdditionalSecurityContext

`func (o *UpdateResourceRequest2) HasAdditionalSecurityContext() bool`

HasAdditionalSecurityContext returns a boolean if a field has been set.

### GetAgentConfiguration

`func (o *UpdateResourceRequest2) GetAgentConfiguration() AgentConfiguration`

GetAgentConfiguration returns the AgentConfiguration field if non-nil, zero value otherwise.

### GetAgentConfigurationOk

`func (o *UpdateResourceRequest2) GetAgentConfigurationOk() (*AgentConfiguration, bool)`

GetAgentConfigurationOk returns a tuple with the AgentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfiguration

`func (o *UpdateResourceRequest2) SetAgentConfiguration(v AgentConfiguration)`

SetAgentConfiguration sets AgentConfiguration field to given value.

### HasAgentConfiguration

`func (o *UpdateResourceRequest2) HasAgentConfiguration() bool`

HasAgentConfiguration returns a boolean if a field has been set.

### GetBackupConfiguration

`func (o *UpdateResourceRequest2) GetBackupConfiguration() BackupConfiguration`

GetBackupConfiguration returns the BackupConfiguration field if non-nil, zero value otherwise.

### GetBackupConfigurationOk

`func (o *UpdateResourceRequest2) GetBackupConfigurationOk() (*BackupConfiguration, bool)`

GetBackupConfigurationOk returns a tuple with the BackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfiguration

`func (o *UpdateResourceRequest2) SetBackupConfiguration(v BackupConfiguration)`

SetBackupConfiguration sets BackupConfiguration field to given value.

### HasBackupConfiguration

`func (o *UpdateResourceRequest2) HasBackupConfiguration() bool`

HasBackupConfiguration returns a boolean if a field has been set.

### GetContainerImageConfiguration

`func (o *UpdateResourceRequest2) GetContainerImageConfiguration() ContainerImageConfiguration`

GetContainerImageConfiguration returns the ContainerImageConfiguration field if non-nil, zero value otherwise.

### GetContainerImageConfigurationOk

`func (o *UpdateResourceRequest2) GetContainerImageConfigurationOk() (*ContainerImageConfiguration, bool)`

GetContainerImageConfigurationOk returns a tuple with the ContainerImageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImageConfiguration

`func (o *UpdateResourceRequest2) SetContainerImageConfiguration(v ContainerImageConfiguration)`

SetContainerImageConfiguration sets ContainerImageConfiguration field to given value.

### HasContainerImageConfiguration

`func (o *UpdateResourceRequest2) HasContainerImageConfiguration() bool`

HasContainerImageConfiguration returns a boolean if a field has been set.

### GetCustomLabels

`func (o *UpdateResourceRequest2) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *UpdateResourceRequest2) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *UpdateResourceRequest2) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *UpdateResourceRequest2) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetCustomSysCTLs

`func (o *UpdateResourceRequest2) GetCustomSysCTLs() map[string]string`

GetCustomSysCTLs returns the CustomSysCTLs field if non-nil, zero value otherwise.

### GetCustomSysCTLsOk

`func (o *UpdateResourceRequest2) GetCustomSysCTLsOk() (*map[string]string, bool)`

GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSysCTLs

`func (o *UpdateResourceRequest2) SetCustomSysCTLs(v map[string]string)`

SetCustomSysCTLs sets CustomSysCTLs field to given value.

### HasCustomSysCTLs

`func (o *UpdateResourceRequest2) HasCustomSysCTLs() bool`

HasCustomSysCTLs returns a boolean if a field has been set.

### GetCustomULimits

`func (o *UpdateResourceRequest2) GetCustomULimits() []CustomULimits`

GetCustomULimits returns the CustomULimits field if non-nil, zero value otherwise.

### GetCustomULimitsOk

`func (o *UpdateResourceRequest2) GetCustomULimitsOk() (*[]CustomULimits, bool)`

GetCustomULimitsOk returns a tuple with the CustomULimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomULimits

`func (o *UpdateResourceRequest2) SetCustomULimits(v []CustomULimits)`

SetCustomULimits sets CustomULimits field to given value.

### HasCustomULimits

`func (o *UpdateResourceRequest2) HasCustomULimits() bool`

HasCustomULimits returns a boolean if a field has been set.

### GetDependencies

`func (o *UpdateResourceRequest2) GetDependencies() []ResourceDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *UpdateResourceRequest2) GetDependenciesOk() (*[]ResourceDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *UpdateResourceRequest2) SetDependencies(v []ResourceDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *UpdateResourceRequest2) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateResourceRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResourceRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResourceRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResourceRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *UpdateResourceRequest2) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateResourceRequest2) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateResourceRequest2) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateResourceRequest2) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *UpdateResourceRequest2) GetHelmChartConfiguration() HelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *UpdateResourceRequest2) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *UpdateResourceRequest2) SetHelmChartConfiguration(v HelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *UpdateResourceRequest2) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetImageConfigId

`func (o *UpdateResourceRequest2) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *UpdateResourceRequest2) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *UpdateResourceRequest2) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *UpdateResourceRequest2) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *UpdateResourceRequest2) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *UpdateResourceRequest2) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *UpdateResourceRequest2) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *UpdateResourceRequest2) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetJobConfig

`func (o *UpdateResourceRequest2) GetJobConfig() JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *UpdateResourceRequest2) GetJobConfigOk() (*JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *UpdateResourceRequest2) SetJobConfig(v JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *UpdateResourceRequest2) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *UpdateResourceRequest2) GetKustomizeConfiguration() KustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *UpdateResourceRequest2) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *UpdateResourceRequest2) SetKustomizeConfiguration(v KustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *UpdateResourceRequest2) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration`

GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL4LoadBalancerConfigurationOk

`func (o *UpdateResourceRequest2) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool)`

GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration)`

SetL4LoadBalancerConfiguration sets L4LoadBalancerConfiguration field to given value.

### HasL4LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) HasL4LoadBalancerConfiguration() bool`

HasL4LoadBalancerConfiguration returns a boolean if a field has been set.

### GetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration`

GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field if non-nil, zero value otherwise.

### GetL7LoadBalancerConfigurationOk

`func (o *UpdateResourceRequest2) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool)`

GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration)`

SetL7LoadBalancerConfiguration sets L7LoadBalancerConfiguration field to given value.

### HasL7LoadBalancerConfiguration

`func (o *UpdateResourceRequest2) HasL7LoadBalancerConfiguration() bool`

HasL7LoadBalancerConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateResourceRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourceRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourceRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourceRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnPremTerraformConfigurations

`func (o *UpdateResourceRequest2) GetOnPremTerraformConfigurations() map[string]TerraformConfiguration`

GetOnPremTerraformConfigurations returns the OnPremTerraformConfigurations field if non-nil, zero value otherwise.

### GetOnPremTerraformConfigurationsOk

`func (o *UpdateResourceRequest2) GetOnPremTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetOnPremTerraformConfigurationsOk returns a tuple with the OnPremTerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremTerraformConfigurations

`func (o *UpdateResourceRequest2) SetOnPremTerraformConfigurations(v map[string]TerraformConfiguration)`

SetOnPremTerraformConfigurations sets OnPremTerraformConfigurations field to given value.

### HasOnPremTerraformConfigurations

`func (o *UpdateResourceRequest2) HasOnPremTerraformConfigurations() bool`

HasOnPremTerraformConfigurations returns a boolean if a field has been set.

### GetOperatorCRDConfiguration

`func (o *UpdateResourceRequest2) GetOperatorCRDConfiguration() OperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *UpdateResourceRequest2) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *UpdateResourceRequest2) SetOperatorCRDConfiguration(v OperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *UpdateResourceRequest2) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetTerraformConfigurations

`func (o *UpdateResourceRequest2) GetTerraformConfigurations() map[string]TerraformConfiguration`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *UpdateResourceRequest2) GetTerraformConfigurationsOk() (*map[string]TerraformConfiguration, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *UpdateResourceRequest2) SetTerraformConfigurations(v map[string]TerraformConfiguration)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *UpdateResourceRequest2) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


