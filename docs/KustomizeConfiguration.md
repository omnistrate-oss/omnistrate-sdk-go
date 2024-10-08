# KustomizeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitConfiguration** | Pointer to [**GitConfiguration**](GitConfiguration.md) |  | [optional] 
**HelmChartDependencies** | Pointer to [**[]OperatorHelmChartDependency**](OperatorHelmChartDependency.md) | The helm chart dependencies for the CRD - Optional | [optional] 
**KustomizePath** | **string** | The path to the kustomize directory | 

## Methods

### NewKustomizeConfiguration

`func NewKustomizeConfiguration(kustomizePath string, ) *KustomizeConfiguration`

NewKustomizeConfiguration instantiates a new KustomizeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKustomizeConfigurationWithDefaults

`func NewKustomizeConfigurationWithDefaults() *KustomizeConfiguration`

NewKustomizeConfigurationWithDefaults instantiates a new KustomizeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitConfiguration

`func (o *KustomizeConfiguration) GetGitConfiguration() GitConfiguration`

GetGitConfiguration returns the GitConfiguration field if non-nil, zero value otherwise.

### GetGitConfigurationOk

`func (o *KustomizeConfiguration) GetGitConfigurationOk() (*GitConfiguration, bool)`

GetGitConfigurationOk returns a tuple with the GitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitConfiguration

`func (o *KustomizeConfiguration) SetGitConfiguration(v GitConfiguration)`

SetGitConfiguration sets GitConfiguration field to given value.

### HasGitConfiguration

`func (o *KustomizeConfiguration) HasGitConfiguration() bool`

HasGitConfiguration returns a boolean if a field has been set.

### GetHelmChartDependencies

`func (o *KustomizeConfiguration) GetHelmChartDependencies() []OperatorHelmChartDependency`

GetHelmChartDependencies returns the HelmChartDependencies field if non-nil, zero value otherwise.

### GetHelmChartDependenciesOk

`func (o *KustomizeConfiguration) GetHelmChartDependenciesOk() (*[]OperatorHelmChartDependency, bool)`

GetHelmChartDependenciesOk returns a tuple with the HelmChartDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartDependencies

`func (o *KustomizeConfiguration) SetHelmChartDependencies(v []OperatorHelmChartDependency)`

SetHelmChartDependencies sets HelmChartDependencies field to given value.

### HasHelmChartDependencies

`func (o *KustomizeConfiguration) HasHelmChartDependencies() bool`

HasHelmChartDependencies returns a boolean if a field has been set.

### GetKustomizePath

`func (o *KustomizeConfiguration) GetKustomizePath() string`

GetKustomizePath returns the KustomizePath field if non-nil, zero value otherwise.

### GetKustomizePathOk

`func (o *KustomizeConfiguration) GetKustomizePathOk() (*string, bool)`

GetKustomizePathOk returns a tuple with the KustomizePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizePath

`func (o *KustomizeConfiguration) SetKustomizePath(v string)`

SetKustomizePath sets KustomizePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


