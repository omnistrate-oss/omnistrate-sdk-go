# HelmPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package | 
**ChartRepoName** | **string** | The chart repository name of the Helm package | 
**ChartRepoUrl** | **string** | The chart repository URL of the Helm package | 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm package (mutually exclusive with layeredChartValues) | [optional] 
**ChartVersion** | **string** | The chart version of the Helm package | 
**LayeredChartValues** | Pointer to [**[]ChartValuesRef**](ChartValuesRef.md) | Layered chart values configuration allowing multiple conditional value sets (mutually exclusive with chartValues). | [optional] 
**Namespace** | **string** | The namespace where the Helm package should be installed | 
**Password** | Pointer to **string** | The password to authenticate with the registry | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry | [optional] 

## Methods

### NewHelmPackage

`func NewHelmPackage(chartName string, chartRepoName string, chartRepoUrl string, chartVersion string, namespace string, ) *HelmPackage`

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


### GetChartRepoName

`func (o *HelmPackage) GetChartRepoName() string`

GetChartRepoName returns the ChartRepoName field if non-nil, zero value otherwise.

### GetChartRepoNameOk

`func (o *HelmPackage) GetChartRepoNameOk() (*string, bool)`

GetChartRepoNameOk returns a tuple with the ChartRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoName

`func (o *HelmPackage) SetChartRepoName(v string)`

SetChartRepoName sets ChartRepoName field to given value.


### GetChartRepoUrl

`func (o *HelmPackage) GetChartRepoUrl() string`

GetChartRepoUrl returns the ChartRepoUrl field if non-nil, zero value otherwise.

### GetChartRepoUrlOk

`func (o *HelmPackage) GetChartRepoUrlOk() (*string, bool)`

GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoUrl

`func (o *HelmPackage) SetChartRepoUrl(v string)`

SetChartRepoUrl sets ChartRepoUrl field to given value.


### GetChartValues

`func (o *HelmPackage) GetChartValues() map[string]interface{}`

GetChartValues returns the ChartValues field if non-nil, zero value otherwise.

### GetChartValuesOk

`func (o *HelmPackage) GetChartValuesOk() (*map[string]interface{}, bool)`

GetChartValuesOk returns a tuple with the ChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartValues

`func (o *HelmPackage) SetChartValues(v map[string]interface{})`

SetChartValues sets ChartValues field to given value.

### HasChartValues

`func (o *HelmPackage) HasChartValues() bool`

HasChartValues returns a boolean if a field has been set.

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


### GetLayeredChartValues

`func (o *HelmPackage) GetLayeredChartValues() []ChartValuesRef`

GetLayeredChartValues returns the LayeredChartValues field if non-nil, zero value otherwise.

### GetLayeredChartValuesOk

`func (o *HelmPackage) GetLayeredChartValuesOk() (*[]ChartValuesRef, bool)`

GetLayeredChartValuesOk returns a tuple with the LayeredChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayeredChartValues

`func (o *HelmPackage) SetLayeredChartValues(v []ChartValuesRef)`

SetLayeredChartValues sets LayeredChartValues field to given value.

### HasLayeredChartValues

`func (o *HelmPackage) HasLayeredChartValues() bool`

HasLayeredChartValues returns a boolean if a field has been set.

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


### GetPassword

`func (o *HelmPackage) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HelmPackage) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HelmPackage) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HelmPackage) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *HelmPackage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HelmPackage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HelmPackage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HelmPackage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


