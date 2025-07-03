# HelmAdoptionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartRepoURL** | **string** | The Helm chart repository URL | 
**Password** | Pointer to **string** | The password to authenticate with the registry | [optional] 
**ReleaseName** | **string** | The Helm release name | 
**ReleaseNamespace** | **string** | The Helm release namespace | 
**RuntimeConfiguration** | Pointer to [**HelmRuntimeConfiguration**](HelmRuntimeConfiguration.md) |  | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry | [optional] 

## Methods

### NewHelmAdoptionConfiguration

`func NewHelmAdoptionConfiguration(chartRepoURL string, releaseName string, releaseNamespace string, ) *HelmAdoptionConfiguration`

NewHelmAdoptionConfiguration instantiates a new HelmAdoptionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmAdoptionConfigurationWithDefaults

`func NewHelmAdoptionConfigurationWithDefaults() *HelmAdoptionConfiguration`

NewHelmAdoptionConfigurationWithDefaults instantiates a new HelmAdoptionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartRepoURL

`func (o *HelmAdoptionConfiguration) GetChartRepoURL() string`

GetChartRepoURL returns the ChartRepoURL field if non-nil, zero value otherwise.

### GetChartRepoURLOk

`func (o *HelmAdoptionConfiguration) GetChartRepoURLOk() (*string, bool)`

GetChartRepoURLOk returns a tuple with the ChartRepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoURL

`func (o *HelmAdoptionConfiguration) SetChartRepoURL(v string)`

SetChartRepoURL sets ChartRepoURL field to given value.


### GetPassword

`func (o *HelmAdoptionConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HelmAdoptionConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HelmAdoptionConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HelmAdoptionConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetReleaseName

`func (o *HelmAdoptionConfiguration) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *HelmAdoptionConfiguration) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *HelmAdoptionConfiguration) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.


### GetReleaseNamespace

`func (o *HelmAdoptionConfiguration) GetReleaseNamespace() string`

GetReleaseNamespace returns the ReleaseNamespace field if non-nil, zero value otherwise.

### GetReleaseNamespaceOk

`func (o *HelmAdoptionConfiguration) GetReleaseNamespaceOk() (*string, bool)`

GetReleaseNamespaceOk returns a tuple with the ReleaseNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNamespace

`func (o *HelmAdoptionConfiguration) SetReleaseNamespace(v string)`

SetReleaseNamespace sets ReleaseNamespace field to given value.


### GetRuntimeConfiguration

`func (o *HelmAdoptionConfiguration) GetRuntimeConfiguration() HelmRuntimeConfiguration`

GetRuntimeConfiguration returns the RuntimeConfiguration field if non-nil, zero value otherwise.

### GetRuntimeConfigurationOk

`func (o *HelmAdoptionConfiguration) GetRuntimeConfigurationOk() (*HelmRuntimeConfiguration, bool)`

GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConfiguration

`func (o *HelmAdoptionConfiguration) SetRuntimeConfiguration(v HelmRuntimeConfiguration)`

SetRuntimeConfiguration sets RuntimeConfiguration field to given value.

### HasRuntimeConfiguration

`func (o *HelmAdoptionConfiguration) HasRuntimeConfiguration() bool`

HasRuntimeConfiguration returns a boolean if a field has been set.

### GetUsername

`func (o *HelmAdoptionConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HelmAdoptionConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HelmAdoptionConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HelmAdoptionConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


