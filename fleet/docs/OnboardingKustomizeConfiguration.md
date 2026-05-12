# OnboardingKustomizeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEndpoints** | Pointer to [**map[string]OnboardingEndpoint**](OnboardingEndpoint.md) | Endpoints from the kustomize deployment to expose. | [optional] 
**DisableReconciliation** | Pointer to **bool** | Flag to disable drift reconciliation. | [optional] 
**GitConfiguration** | Pointer to [**GitConfiguration**](GitConfiguration.md) |  | [optional] 
**HelmChartDependencies** | Pointer to [**[]OnboardingChartDependency**](OnboardingChartDependency.md) | Helm chart dependencies for the kustomize deployment. | [optional] 
**KustomizePath** | **string** | The path to the kustomize directory. | 

## Methods

### NewOnboardingKustomizeConfiguration

`func NewOnboardingKustomizeConfiguration(kustomizePath string, ) *OnboardingKustomizeConfiguration`

NewOnboardingKustomizeConfiguration instantiates a new OnboardingKustomizeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingKustomizeConfigurationWithDefaults

`func NewOnboardingKustomizeConfigurationWithDefaults() *OnboardingKustomizeConfiguration`

NewOnboardingKustomizeConfigurationWithDefaults instantiates a new OnboardingKustomizeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEndpoints

`func (o *OnboardingKustomizeConfiguration) GetAdditionalEndpoints() map[string]OnboardingEndpoint`

GetAdditionalEndpoints returns the AdditionalEndpoints field if non-nil, zero value otherwise.

### GetAdditionalEndpointsOk

`func (o *OnboardingKustomizeConfiguration) GetAdditionalEndpointsOk() (*map[string]OnboardingEndpoint, bool)`

GetAdditionalEndpointsOk returns a tuple with the AdditionalEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEndpoints

`func (o *OnboardingKustomizeConfiguration) SetAdditionalEndpoints(v map[string]OnboardingEndpoint)`

SetAdditionalEndpoints sets AdditionalEndpoints field to given value.

### HasAdditionalEndpoints

`func (o *OnboardingKustomizeConfiguration) HasAdditionalEndpoints() bool`

HasAdditionalEndpoints returns a boolean if a field has been set.

### GetDisableReconciliation

`func (o *OnboardingKustomizeConfiguration) GetDisableReconciliation() bool`

GetDisableReconciliation returns the DisableReconciliation field if non-nil, zero value otherwise.

### GetDisableReconciliationOk

`func (o *OnboardingKustomizeConfiguration) GetDisableReconciliationOk() (*bool, bool)`

GetDisableReconciliationOk returns a tuple with the DisableReconciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReconciliation

`func (o *OnboardingKustomizeConfiguration) SetDisableReconciliation(v bool)`

SetDisableReconciliation sets DisableReconciliation field to given value.

### HasDisableReconciliation

`func (o *OnboardingKustomizeConfiguration) HasDisableReconciliation() bool`

HasDisableReconciliation returns a boolean if a field has been set.

### GetGitConfiguration

`func (o *OnboardingKustomizeConfiguration) GetGitConfiguration() GitConfiguration`

GetGitConfiguration returns the GitConfiguration field if non-nil, zero value otherwise.

### GetGitConfigurationOk

`func (o *OnboardingKustomizeConfiguration) GetGitConfigurationOk() (*GitConfiguration, bool)`

GetGitConfigurationOk returns a tuple with the GitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitConfiguration

`func (o *OnboardingKustomizeConfiguration) SetGitConfiguration(v GitConfiguration)`

SetGitConfiguration sets GitConfiguration field to given value.

### HasGitConfiguration

`func (o *OnboardingKustomizeConfiguration) HasGitConfiguration() bool`

HasGitConfiguration returns a boolean if a field has been set.

### GetHelmChartDependencies

`func (o *OnboardingKustomizeConfiguration) GetHelmChartDependencies() []OnboardingChartDependency`

GetHelmChartDependencies returns the HelmChartDependencies field if non-nil, zero value otherwise.

### GetHelmChartDependenciesOk

`func (o *OnboardingKustomizeConfiguration) GetHelmChartDependenciesOk() (*[]OnboardingChartDependency, bool)`

GetHelmChartDependenciesOk returns a tuple with the HelmChartDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartDependencies

`func (o *OnboardingKustomizeConfiguration) SetHelmChartDependencies(v []OnboardingChartDependency)`

SetHelmChartDependencies sets HelmChartDependencies field to given value.

### HasHelmChartDependencies

`func (o *OnboardingKustomizeConfiguration) HasHelmChartDependencies() bool`

HasHelmChartDependencies returns a boolean if a field has been set.

### GetKustomizePath

`func (o *OnboardingKustomizeConfiguration) GetKustomizePath() string`

GetKustomizePath returns the KustomizePath field if non-nil, zero value otherwise.

### GetKustomizePathOk

`func (o *OnboardingKustomizeConfiguration) GetKustomizePathOk() (*string, bool)`

GetKustomizePathOk returns a tuple with the KustomizePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizePath

`func (o *OnboardingKustomizeConfiguration) SetKustomizePath(v string)`

SetKustomizePath sets KustomizePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


