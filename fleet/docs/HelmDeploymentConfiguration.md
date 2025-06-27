# HelmDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | Name of the Helm chart to be deployed | 
**ChartVersion** | **string** | Version of the Helm chart to be deployed | 
**DeploymentErrors** | Pointer to **string** | Errors encountered during the Helm deployment | [optional] 
**PodEvents** | Pointer to [**map[string][]PodEvent**](array.md) | Events associated with the pods in the generic resource deployment | [optional] 
**PodStatus** | Pointer to **map[string]string** | Status of the pods in the generic resource deployment | [optional] 
**PodToHostMapping** | Pointer to **map[string]string** | Mapping of pods to k8s node names for the generic resource deployment | [optional] 
**ReleaseName** | **string** | Name of the Helm release | 
**ReleaseNamespace** | **string** | Namespace where the Helm release will be deployed | 
**ReleaseStatus** | **string** | Status of the Helm release | 
**RepositoryURL** | **string** | URL of the Helm chart repository | 
**RuntimeConfiguration** | Pointer to **map[string]interface{}** | Runtime configuration for the Helm deployment | [optional] 
**Values** | Pointer to **map[string]interface{}** | Deployed values for the Helm chart in YAML format | [optional] 

## Methods

### NewHelmDeploymentConfiguration

`func NewHelmDeploymentConfiguration(chartName string, chartVersion string, releaseName string, releaseNamespace string, releaseStatus string, repositoryURL string, ) *HelmDeploymentConfiguration`

NewHelmDeploymentConfiguration instantiates a new HelmDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmDeploymentConfigurationWithDefaults

`func NewHelmDeploymentConfigurationWithDefaults() *HelmDeploymentConfiguration`

NewHelmDeploymentConfigurationWithDefaults instantiates a new HelmDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmDeploymentConfiguration) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmDeploymentConfiguration) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmDeploymentConfiguration) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmDeploymentConfiguration) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmDeploymentConfiguration) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmDeploymentConfiguration) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetDeploymentErrors

`func (o *HelmDeploymentConfiguration) GetDeploymentErrors() string`

GetDeploymentErrors returns the DeploymentErrors field if non-nil, zero value otherwise.

### GetDeploymentErrorsOk

`func (o *HelmDeploymentConfiguration) GetDeploymentErrorsOk() (*string, bool)`

GetDeploymentErrorsOk returns a tuple with the DeploymentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentErrors

`func (o *HelmDeploymentConfiguration) SetDeploymentErrors(v string)`

SetDeploymentErrors sets DeploymentErrors field to given value.

### HasDeploymentErrors

`func (o *HelmDeploymentConfiguration) HasDeploymentErrors() bool`

HasDeploymentErrors returns a boolean if a field has been set.

### GetPodEvents

`func (o *HelmDeploymentConfiguration) GetPodEvents() map[string][]PodEvent`

GetPodEvents returns the PodEvents field if non-nil, zero value otherwise.

### GetPodEventsOk

`func (o *HelmDeploymentConfiguration) GetPodEventsOk() (*map[string][]PodEvent, bool)`

GetPodEventsOk returns a tuple with the PodEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodEvents

`func (o *HelmDeploymentConfiguration) SetPodEvents(v map[string][]PodEvent)`

SetPodEvents sets PodEvents field to given value.

### HasPodEvents

`func (o *HelmDeploymentConfiguration) HasPodEvents() bool`

HasPodEvents returns a boolean if a field has been set.

### GetPodStatus

`func (o *HelmDeploymentConfiguration) GetPodStatus() map[string]string`

GetPodStatus returns the PodStatus field if non-nil, zero value otherwise.

### GetPodStatusOk

`func (o *HelmDeploymentConfiguration) GetPodStatusOk() (*map[string]string, bool)`

GetPodStatusOk returns a tuple with the PodStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodStatus

`func (o *HelmDeploymentConfiguration) SetPodStatus(v map[string]string)`

SetPodStatus sets PodStatus field to given value.

### HasPodStatus

`func (o *HelmDeploymentConfiguration) HasPodStatus() bool`

HasPodStatus returns a boolean if a field has been set.

### GetPodToHostMapping

`func (o *HelmDeploymentConfiguration) GetPodToHostMapping() map[string]string`

GetPodToHostMapping returns the PodToHostMapping field if non-nil, zero value otherwise.

### GetPodToHostMappingOk

`func (o *HelmDeploymentConfiguration) GetPodToHostMappingOk() (*map[string]string, bool)`

GetPodToHostMappingOk returns a tuple with the PodToHostMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodToHostMapping

`func (o *HelmDeploymentConfiguration) SetPodToHostMapping(v map[string]string)`

SetPodToHostMapping sets PodToHostMapping field to given value.

### HasPodToHostMapping

`func (o *HelmDeploymentConfiguration) HasPodToHostMapping() bool`

HasPodToHostMapping returns a boolean if a field has been set.

### GetReleaseName

`func (o *HelmDeploymentConfiguration) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *HelmDeploymentConfiguration) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *HelmDeploymentConfiguration) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.


### GetReleaseNamespace

`func (o *HelmDeploymentConfiguration) GetReleaseNamespace() string`

GetReleaseNamespace returns the ReleaseNamespace field if non-nil, zero value otherwise.

### GetReleaseNamespaceOk

`func (o *HelmDeploymentConfiguration) GetReleaseNamespaceOk() (*string, bool)`

GetReleaseNamespaceOk returns a tuple with the ReleaseNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNamespace

`func (o *HelmDeploymentConfiguration) SetReleaseNamespace(v string)`

SetReleaseNamespace sets ReleaseNamespace field to given value.


### GetReleaseStatus

`func (o *HelmDeploymentConfiguration) GetReleaseStatus() string`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *HelmDeploymentConfiguration) GetReleaseStatusOk() (*string, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *HelmDeploymentConfiguration) SetReleaseStatus(v string)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetRepositoryURL

`func (o *HelmDeploymentConfiguration) GetRepositoryURL() string`

GetRepositoryURL returns the RepositoryURL field if non-nil, zero value otherwise.

### GetRepositoryURLOk

`func (o *HelmDeploymentConfiguration) GetRepositoryURLOk() (*string, bool)`

GetRepositoryURLOk returns a tuple with the RepositoryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryURL

`func (o *HelmDeploymentConfiguration) SetRepositoryURL(v string)`

SetRepositoryURL sets RepositoryURL field to given value.


### GetRuntimeConfiguration

`func (o *HelmDeploymentConfiguration) GetRuntimeConfiguration() map[string]interface{}`

GetRuntimeConfiguration returns the RuntimeConfiguration field if non-nil, zero value otherwise.

### GetRuntimeConfigurationOk

`func (o *HelmDeploymentConfiguration) GetRuntimeConfigurationOk() (*map[string]interface{}, bool)`

GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConfiguration

`func (o *HelmDeploymentConfiguration) SetRuntimeConfiguration(v map[string]interface{})`

SetRuntimeConfiguration sets RuntimeConfiguration field to given value.

### HasRuntimeConfiguration

`func (o *HelmDeploymentConfiguration) HasRuntimeConfiguration() bool`

HasRuntimeConfiguration returns a boolean if a field has been set.

### GetValues

`func (o *HelmDeploymentConfiguration) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HelmDeploymentConfiguration) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HelmDeploymentConfiguration) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HelmDeploymentConfiguration) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


