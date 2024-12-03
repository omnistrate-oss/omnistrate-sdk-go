# OperatorHelmChartDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The name of the helm chart | 
**ChartRepoName** | Pointer to **string** | The repository name of the Helm chart | [optional] 
**ChartRepoUrl** | Pointer to **string** | The repository URL of the Helm chart | [optional] 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm chart | [optional] 
**ChartVersion** | **string** | The version of the helm chart | 
**Password** | Pointer to **string** | The password to authenticate with the registry | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry | [optional] 

## Methods

### NewOperatorHelmChartDependency

`func NewOperatorHelmChartDependency(chartName string, chartVersion string, ) *OperatorHelmChartDependency`

NewOperatorHelmChartDependency instantiates a new OperatorHelmChartDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorHelmChartDependencyWithDefaults

`func NewOperatorHelmChartDependencyWithDefaults() *OperatorHelmChartDependency`

NewOperatorHelmChartDependencyWithDefaults instantiates a new OperatorHelmChartDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *OperatorHelmChartDependency) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *OperatorHelmChartDependency) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *OperatorHelmChartDependency) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartRepoName

`func (o *OperatorHelmChartDependency) GetChartRepoName() string`

GetChartRepoName returns the ChartRepoName field if non-nil, zero value otherwise.

### GetChartRepoNameOk

`func (o *OperatorHelmChartDependency) GetChartRepoNameOk() (*string, bool)`

GetChartRepoNameOk returns a tuple with the ChartRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoName

`func (o *OperatorHelmChartDependency) SetChartRepoName(v string)`

SetChartRepoName sets ChartRepoName field to given value.

### HasChartRepoName

`func (o *OperatorHelmChartDependency) HasChartRepoName() bool`

HasChartRepoName returns a boolean if a field has been set.

### GetChartRepoUrl

`func (o *OperatorHelmChartDependency) GetChartRepoUrl() string`

GetChartRepoUrl returns the ChartRepoUrl field if non-nil, zero value otherwise.

### GetChartRepoUrlOk

`func (o *OperatorHelmChartDependency) GetChartRepoUrlOk() (*string, bool)`

GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoUrl

`func (o *OperatorHelmChartDependency) SetChartRepoUrl(v string)`

SetChartRepoUrl sets ChartRepoUrl field to given value.

### HasChartRepoUrl

`func (o *OperatorHelmChartDependency) HasChartRepoUrl() bool`

HasChartRepoUrl returns a boolean if a field has been set.

### GetChartValues

`func (o *OperatorHelmChartDependency) GetChartValues() map[string]interface{}`

GetChartValues returns the ChartValues field if non-nil, zero value otherwise.

### GetChartValuesOk

`func (o *OperatorHelmChartDependency) GetChartValuesOk() (*map[string]interface{}, bool)`

GetChartValuesOk returns a tuple with the ChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartValues

`func (o *OperatorHelmChartDependency) SetChartValues(v map[string]interface{})`

SetChartValues sets ChartValues field to given value.

### HasChartValues

`func (o *OperatorHelmChartDependency) HasChartValues() bool`

HasChartValues returns a boolean if a field has been set.

### GetChartVersion

`func (o *OperatorHelmChartDependency) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *OperatorHelmChartDependency) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *OperatorHelmChartDependency) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetPassword

`func (o *OperatorHelmChartDependency) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OperatorHelmChartDependency) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OperatorHelmChartDependency) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OperatorHelmChartDependency) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *OperatorHelmChartDependency) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OperatorHelmChartDependency) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OperatorHelmChartDependency) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OperatorHelmChartDependency) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


