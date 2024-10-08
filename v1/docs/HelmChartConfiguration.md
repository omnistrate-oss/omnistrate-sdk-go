# HelmChartConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package | 
**ChartRepoName** | **string** | The chart repository name of the Helm package | 
**ChartRepoUrl** | **string** | The chart repository URL of the Helm package | 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm package | [optional] 
**ChartVersion** | **string** | The chart version of the Helm package | 

## Methods

### NewHelmChartConfiguration

`func NewHelmChartConfiguration(chartName string, chartRepoName string, chartRepoUrl string, chartVersion string, ) *HelmChartConfiguration`

NewHelmChartConfiguration instantiates a new HelmChartConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartConfigurationWithDefaults

`func NewHelmChartConfigurationWithDefaults() *HelmChartConfiguration`

NewHelmChartConfigurationWithDefaults instantiates a new HelmChartConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmChartConfiguration) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmChartConfiguration) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmChartConfiguration) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartRepoName

`func (o *HelmChartConfiguration) GetChartRepoName() string`

GetChartRepoName returns the ChartRepoName field if non-nil, zero value otherwise.

### GetChartRepoNameOk

`func (o *HelmChartConfiguration) GetChartRepoNameOk() (*string, bool)`

GetChartRepoNameOk returns a tuple with the ChartRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoName

`func (o *HelmChartConfiguration) SetChartRepoName(v string)`

SetChartRepoName sets ChartRepoName field to given value.


### GetChartRepoUrl

`func (o *HelmChartConfiguration) GetChartRepoUrl() string`

GetChartRepoUrl returns the ChartRepoUrl field if non-nil, zero value otherwise.

### GetChartRepoUrlOk

`func (o *HelmChartConfiguration) GetChartRepoUrlOk() (*string, bool)`

GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoUrl

`func (o *HelmChartConfiguration) SetChartRepoUrl(v string)`

SetChartRepoUrl sets ChartRepoUrl field to given value.


### GetChartValues

`func (o *HelmChartConfiguration) GetChartValues() map[string]interface{}`

GetChartValues returns the ChartValues field if non-nil, zero value otherwise.

### GetChartValuesOk

`func (o *HelmChartConfiguration) GetChartValuesOk() (*map[string]interface{}, bool)`

GetChartValuesOk returns a tuple with the ChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartValues

`func (o *HelmChartConfiguration) SetChartValues(v map[string]interface{})`

SetChartValues sets ChartValues field to given value.

### HasChartValues

`func (o *HelmChartConfiguration) HasChartValues() bool`

HasChartValues returns a boolean if a field has been set.

### GetChartVersion

`func (o *HelmChartConfiguration) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmChartConfiguration) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmChartConfiguration) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


