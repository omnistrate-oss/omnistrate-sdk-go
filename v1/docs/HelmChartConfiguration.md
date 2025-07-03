# HelmChartConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package | 
**ChartRepoName** | **string** | The chart repository name of the Helm package | 
**ChartRepoUrl** | **string** | The chart repository URL of the Helm package | 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm package (mutually exclusive with layeredChartValues) | [optional] 
**ChartVersion** | **string** | The chart version of the Helm package | 
**EndpointConfiguration** | Pointer to [**map[string]Endpoint**](Endpoint.md) | The endpoints from the Helm Deployment to expose to the customer | [optional] 
**LayeredChartValues** | Pointer to [**[]ChartValuesRef**](ChartValuesRef.md) | Layered chart values configuration with conditional scoping (mutually exclusive with chartValues). Values are processed in order - later entries override earlier ones for the same keys. | [optional] 
**Password** | Pointer to **string** | The password to authenticate with the registry | [optional] 
**RuntimeConfiguration** | Pointer to [**HelmRuntimeConfiguration**](HelmRuntimeConfiguration.md) |  | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry | [optional] 

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


### GetEndpointConfiguration

`func (o *HelmChartConfiguration) GetEndpointConfiguration() map[string]Endpoint`

GetEndpointConfiguration returns the EndpointConfiguration field if non-nil, zero value otherwise.

### GetEndpointConfigurationOk

`func (o *HelmChartConfiguration) GetEndpointConfigurationOk() (*map[string]Endpoint, bool)`

GetEndpointConfigurationOk returns a tuple with the EndpointConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConfiguration

`func (o *HelmChartConfiguration) SetEndpointConfiguration(v map[string]Endpoint)`

SetEndpointConfiguration sets EndpointConfiguration field to given value.

### HasEndpointConfiguration

`func (o *HelmChartConfiguration) HasEndpointConfiguration() bool`

HasEndpointConfiguration returns a boolean if a field has been set.

### GetLayeredChartValues

`func (o *HelmChartConfiguration) GetLayeredChartValues() []ChartValuesRef`

GetLayeredChartValues returns the LayeredChartValues field if non-nil, zero value otherwise.

### GetLayeredChartValuesOk

`func (o *HelmChartConfiguration) GetLayeredChartValuesOk() (*[]ChartValuesRef, bool)`

GetLayeredChartValuesOk returns a tuple with the LayeredChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayeredChartValues

`func (o *HelmChartConfiguration) SetLayeredChartValues(v []ChartValuesRef)`

SetLayeredChartValues sets LayeredChartValues field to given value.

### HasLayeredChartValues

`func (o *HelmChartConfiguration) HasLayeredChartValues() bool`

HasLayeredChartValues returns a boolean if a field has been set.

### GetPassword

`func (o *HelmChartConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HelmChartConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HelmChartConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HelmChartConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRuntimeConfiguration

`func (o *HelmChartConfiguration) GetRuntimeConfiguration() HelmRuntimeConfiguration`

GetRuntimeConfiguration returns the RuntimeConfiguration field if non-nil, zero value otherwise.

### GetRuntimeConfigurationOk

`func (o *HelmChartConfiguration) GetRuntimeConfigurationOk() (*HelmRuntimeConfiguration, bool)`

GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConfiguration

`func (o *HelmChartConfiguration) SetRuntimeConfiguration(v HelmRuntimeConfiguration)`

SetRuntimeConfiguration sets RuntimeConfiguration field to given value.

### HasRuntimeConfiguration

`func (o *HelmChartConfiguration) HasRuntimeConfiguration() bool`

HasRuntimeConfiguration returns a boolean if a field has been set.

### GetUsername

`func (o *HelmChartConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HelmChartConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HelmChartConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HelmChartConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


