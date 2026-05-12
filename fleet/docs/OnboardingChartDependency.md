# OnboardingChartDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The name of the Helm chart. | 
**ChartRepoName** | Pointer to **string** | The repository name of the Helm chart. | [optional] 
**ChartRepoUrl** | Pointer to **string** | The repository URL of the Helm chart. | [optional] 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm chart. | [optional] 
**ChartVersion** | **string** | The version of the Helm chart. | 
**Namespace** | Pointer to **string** | The namespace to deploy the chart into. | [optional] 
**Password** | Pointer to **string** | The password to authenticate with the registry. | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry. | [optional] 

## Methods

### NewOnboardingChartDependency

`func NewOnboardingChartDependency(chartName string, chartVersion string, ) *OnboardingChartDependency`

NewOnboardingChartDependency instantiates a new OnboardingChartDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingChartDependencyWithDefaults

`func NewOnboardingChartDependencyWithDefaults() *OnboardingChartDependency`

NewOnboardingChartDependencyWithDefaults instantiates a new OnboardingChartDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *OnboardingChartDependency) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *OnboardingChartDependency) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *OnboardingChartDependency) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartRepoName

`func (o *OnboardingChartDependency) GetChartRepoName() string`

GetChartRepoName returns the ChartRepoName field if non-nil, zero value otherwise.

### GetChartRepoNameOk

`func (o *OnboardingChartDependency) GetChartRepoNameOk() (*string, bool)`

GetChartRepoNameOk returns a tuple with the ChartRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoName

`func (o *OnboardingChartDependency) SetChartRepoName(v string)`

SetChartRepoName sets ChartRepoName field to given value.

### HasChartRepoName

`func (o *OnboardingChartDependency) HasChartRepoName() bool`

HasChartRepoName returns a boolean if a field has been set.

### GetChartRepoUrl

`func (o *OnboardingChartDependency) GetChartRepoUrl() string`

GetChartRepoUrl returns the ChartRepoUrl field if non-nil, zero value otherwise.

### GetChartRepoUrlOk

`func (o *OnboardingChartDependency) GetChartRepoUrlOk() (*string, bool)`

GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoUrl

`func (o *OnboardingChartDependency) SetChartRepoUrl(v string)`

SetChartRepoUrl sets ChartRepoUrl field to given value.

### HasChartRepoUrl

`func (o *OnboardingChartDependency) HasChartRepoUrl() bool`

HasChartRepoUrl returns a boolean if a field has been set.

### GetChartValues

`func (o *OnboardingChartDependency) GetChartValues() map[string]interface{}`

GetChartValues returns the ChartValues field if non-nil, zero value otherwise.

### GetChartValuesOk

`func (o *OnboardingChartDependency) GetChartValuesOk() (*map[string]interface{}, bool)`

GetChartValuesOk returns a tuple with the ChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartValues

`func (o *OnboardingChartDependency) SetChartValues(v map[string]interface{})`

SetChartValues sets ChartValues field to given value.

### HasChartValues

`func (o *OnboardingChartDependency) HasChartValues() bool`

HasChartValues returns a boolean if a field has been set.

### GetChartVersion

`func (o *OnboardingChartDependency) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *OnboardingChartDependency) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *OnboardingChartDependency) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetNamespace

`func (o *OnboardingChartDependency) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OnboardingChartDependency) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OnboardingChartDependency) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OnboardingChartDependency) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPassword

`func (o *OnboardingChartDependency) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OnboardingChartDependency) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OnboardingChartDependency) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OnboardingChartDependency) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *OnboardingChartDependency) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OnboardingChartDependency) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OnboardingChartDependency) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OnboardingChartDependency) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


