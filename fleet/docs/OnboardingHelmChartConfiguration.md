# OnboardingHelmChartConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactsLocalPath** | Pointer to **string** | The local path where the helm chart artifacts are stored. | [optional] 
**AutoDiscoverImagesTag** | Pointer to **string** | The tag to auto-discover and update in the Helm chart values. | [optional] 
**ChartName** | **string** | The chart name of the Helm package. | 
**ChartRepoName** | Pointer to **string** | The chart repository name. | [optional] 
**ChartRepoUrl** | Pointer to **string** | The chart repository URL. | [optional] 
**ChartValues** | Pointer to **map[string]interface{}** | The values of the Helm package (mutually exclusive with layeredChartValues). | [optional] 
**ChartVersion** | **string** | The chart version of the Helm package. | 
**DefaultNamespace** | Pointer to **string** | The default namespace to deploy into. | [optional] 
**EndpointConfiguration** | Pointer to [**map[string]OnboardingEndpoint**](OnboardingEndpoint.md) | Endpoints from the Helm deployment to expose. | [optional] 
**LayeredChartValues** | Pointer to [**[]ChartValuesRef**](ChartValuesRef.md) | Layered chart values configuration (mutually exclusive with chartValues). | [optional] 
**Password** | Pointer to **string** | The password to authenticate with the registry. | [optional] 
**ReleaseName** | Pointer to **string** | The release name of the Helm package. | [optional] 
**RuntimeConfiguration** | Pointer to [**OnboardingHelmRuntimeConfiguration**](OnboardingHelmRuntimeConfiguration.md) |  | [optional] 
**Username** | Pointer to **string** | The username to authenticate with the registry. | [optional] 

## Methods

### NewOnboardingHelmChartConfiguration

`func NewOnboardingHelmChartConfiguration(chartName string, chartVersion string, ) *OnboardingHelmChartConfiguration`

NewOnboardingHelmChartConfiguration instantiates a new OnboardingHelmChartConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingHelmChartConfigurationWithDefaults

`func NewOnboardingHelmChartConfigurationWithDefaults() *OnboardingHelmChartConfiguration`

NewOnboardingHelmChartConfigurationWithDefaults instantiates a new OnboardingHelmChartConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactsLocalPath

`func (o *OnboardingHelmChartConfiguration) GetArtifactsLocalPath() string`

GetArtifactsLocalPath returns the ArtifactsLocalPath field if non-nil, zero value otherwise.

### GetArtifactsLocalPathOk

`func (o *OnboardingHelmChartConfiguration) GetArtifactsLocalPathOk() (*string, bool)`

GetArtifactsLocalPathOk returns a tuple with the ArtifactsLocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsLocalPath

`func (o *OnboardingHelmChartConfiguration) SetArtifactsLocalPath(v string)`

SetArtifactsLocalPath sets ArtifactsLocalPath field to given value.

### HasArtifactsLocalPath

`func (o *OnboardingHelmChartConfiguration) HasArtifactsLocalPath() bool`

HasArtifactsLocalPath returns a boolean if a field has been set.

### GetAutoDiscoverImagesTag

`func (o *OnboardingHelmChartConfiguration) GetAutoDiscoverImagesTag() string`

GetAutoDiscoverImagesTag returns the AutoDiscoverImagesTag field if non-nil, zero value otherwise.

### GetAutoDiscoverImagesTagOk

`func (o *OnboardingHelmChartConfiguration) GetAutoDiscoverImagesTagOk() (*string, bool)`

GetAutoDiscoverImagesTagOk returns a tuple with the AutoDiscoverImagesTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoverImagesTag

`func (o *OnboardingHelmChartConfiguration) SetAutoDiscoverImagesTag(v string)`

SetAutoDiscoverImagesTag sets AutoDiscoverImagesTag field to given value.

### HasAutoDiscoverImagesTag

`func (o *OnboardingHelmChartConfiguration) HasAutoDiscoverImagesTag() bool`

HasAutoDiscoverImagesTag returns a boolean if a field has been set.

### GetChartName

`func (o *OnboardingHelmChartConfiguration) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *OnboardingHelmChartConfiguration) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *OnboardingHelmChartConfiguration) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartRepoName

`func (o *OnboardingHelmChartConfiguration) GetChartRepoName() string`

GetChartRepoName returns the ChartRepoName field if non-nil, zero value otherwise.

### GetChartRepoNameOk

`func (o *OnboardingHelmChartConfiguration) GetChartRepoNameOk() (*string, bool)`

GetChartRepoNameOk returns a tuple with the ChartRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoName

`func (o *OnboardingHelmChartConfiguration) SetChartRepoName(v string)`

SetChartRepoName sets ChartRepoName field to given value.

### HasChartRepoName

`func (o *OnboardingHelmChartConfiguration) HasChartRepoName() bool`

HasChartRepoName returns a boolean if a field has been set.

### GetChartRepoUrl

`func (o *OnboardingHelmChartConfiguration) GetChartRepoUrl() string`

GetChartRepoUrl returns the ChartRepoUrl field if non-nil, zero value otherwise.

### GetChartRepoUrlOk

`func (o *OnboardingHelmChartConfiguration) GetChartRepoUrlOk() (*string, bool)`

GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepoUrl

`func (o *OnboardingHelmChartConfiguration) SetChartRepoUrl(v string)`

SetChartRepoUrl sets ChartRepoUrl field to given value.

### HasChartRepoUrl

`func (o *OnboardingHelmChartConfiguration) HasChartRepoUrl() bool`

HasChartRepoUrl returns a boolean if a field has been set.

### GetChartValues

`func (o *OnboardingHelmChartConfiguration) GetChartValues() map[string]interface{}`

GetChartValues returns the ChartValues field if non-nil, zero value otherwise.

### GetChartValuesOk

`func (o *OnboardingHelmChartConfiguration) GetChartValuesOk() (*map[string]interface{}, bool)`

GetChartValuesOk returns a tuple with the ChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartValues

`func (o *OnboardingHelmChartConfiguration) SetChartValues(v map[string]interface{})`

SetChartValues sets ChartValues field to given value.

### HasChartValues

`func (o *OnboardingHelmChartConfiguration) HasChartValues() bool`

HasChartValues returns a boolean if a field has been set.

### GetChartVersion

`func (o *OnboardingHelmChartConfiguration) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *OnboardingHelmChartConfiguration) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *OnboardingHelmChartConfiguration) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetDefaultNamespace

`func (o *OnboardingHelmChartConfiguration) GetDefaultNamespace() string`

GetDefaultNamespace returns the DefaultNamespace field if non-nil, zero value otherwise.

### GetDefaultNamespaceOk

`func (o *OnboardingHelmChartConfiguration) GetDefaultNamespaceOk() (*string, bool)`

GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNamespace

`func (o *OnboardingHelmChartConfiguration) SetDefaultNamespace(v string)`

SetDefaultNamespace sets DefaultNamespace field to given value.

### HasDefaultNamespace

`func (o *OnboardingHelmChartConfiguration) HasDefaultNamespace() bool`

HasDefaultNamespace returns a boolean if a field has been set.

### GetEndpointConfiguration

`func (o *OnboardingHelmChartConfiguration) GetEndpointConfiguration() map[string]OnboardingEndpoint`

GetEndpointConfiguration returns the EndpointConfiguration field if non-nil, zero value otherwise.

### GetEndpointConfigurationOk

`func (o *OnboardingHelmChartConfiguration) GetEndpointConfigurationOk() (*map[string]OnboardingEndpoint, bool)`

GetEndpointConfigurationOk returns a tuple with the EndpointConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConfiguration

`func (o *OnboardingHelmChartConfiguration) SetEndpointConfiguration(v map[string]OnboardingEndpoint)`

SetEndpointConfiguration sets EndpointConfiguration field to given value.

### HasEndpointConfiguration

`func (o *OnboardingHelmChartConfiguration) HasEndpointConfiguration() bool`

HasEndpointConfiguration returns a boolean if a field has been set.

### GetLayeredChartValues

`func (o *OnboardingHelmChartConfiguration) GetLayeredChartValues() []ChartValuesRef`

GetLayeredChartValues returns the LayeredChartValues field if non-nil, zero value otherwise.

### GetLayeredChartValuesOk

`func (o *OnboardingHelmChartConfiguration) GetLayeredChartValuesOk() (*[]ChartValuesRef, bool)`

GetLayeredChartValuesOk returns a tuple with the LayeredChartValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayeredChartValues

`func (o *OnboardingHelmChartConfiguration) SetLayeredChartValues(v []ChartValuesRef)`

SetLayeredChartValues sets LayeredChartValues field to given value.

### HasLayeredChartValues

`func (o *OnboardingHelmChartConfiguration) HasLayeredChartValues() bool`

HasLayeredChartValues returns a boolean if a field has been set.

### GetPassword

`func (o *OnboardingHelmChartConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OnboardingHelmChartConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OnboardingHelmChartConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OnboardingHelmChartConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetReleaseName

`func (o *OnboardingHelmChartConfiguration) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *OnboardingHelmChartConfiguration) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *OnboardingHelmChartConfiguration) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *OnboardingHelmChartConfiguration) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetRuntimeConfiguration

`func (o *OnboardingHelmChartConfiguration) GetRuntimeConfiguration() OnboardingHelmRuntimeConfiguration`

GetRuntimeConfiguration returns the RuntimeConfiguration field if non-nil, zero value otherwise.

### GetRuntimeConfigurationOk

`func (o *OnboardingHelmChartConfiguration) GetRuntimeConfigurationOk() (*OnboardingHelmRuntimeConfiguration, bool)`

GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConfiguration

`func (o *OnboardingHelmChartConfiguration) SetRuntimeConfiguration(v OnboardingHelmRuntimeConfiguration)`

SetRuntimeConfiguration sets RuntimeConfiguration field to given value.

### HasRuntimeConfiguration

`func (o *OnboardingHelmChartConfiguration) HasRuntimeConfiguration() bool`

HasRuntimeConfiguration returns a boolean if a field has been set.

### GetUsername

`func (o *OnboardingHelmChartConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OnboardingHelmChartConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OnboardingHelmChartConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OnboardingHelmChartConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


