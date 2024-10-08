# HelmPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package | 
**ChartVersion** | **string** | The chart version of the Helm package | 
**Namespace** | **string** | The namespace where the Helm package is installed | 
**RepoURL** | **string** | The repository URL of the Helm package | 
**Values** | Pointer to **map[string]interface{}** | Custom values for the helm package | [optional] 

## Methods

### NewHelmPackage

`func NewHelmPackage(chartName string, chartVersion string, namespace string, repoURL string, ) *HelmPackage`

NewHelmPackage instantiates a new HelmPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPackageWithDefaults

`func NewHelmPackageWithDefaults() *HelmPackage`

NewHelmPackageWithDefaults instantiates a new HelmPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmPackage) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmPackage) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmPackage) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmPackage) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmPackage) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmPackage) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetNamespace

`func (o *HelmPackage) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmPackage) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmPackage) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetRepoURL

`func (o *HelmPackage) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *HelmPackage) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *HelmPackage) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.


### GetValues

`func (o *HelmPackage) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HelmPackage) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HelmPackage) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HelmPackage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


