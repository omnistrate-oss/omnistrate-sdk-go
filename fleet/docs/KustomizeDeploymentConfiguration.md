# KustomizeDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **string** | Base path for Kustomize configuration | 
**DeploymentErrors** | Pointer to **string** | Errors encountered during the Kustomize deployment | [optional] 
**Overlays** | **map[string]string** | Overlays for Kustomize deployment | 

## Methods

### NewKustomizeDeploymentConfiguration

`func NewKustomizeDeploymentConfiguration(basePath string, overlays map[string]string, ) *KustomizeDeploymentConfiguration`

NewKustomizeDeploymentConfiguration instantiates a new KustomizeDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKustomizeDeploymentConfigurationWithDefaults

`func NewKustomizeDeploymentConfigurationWithDefaults() *KustomizeDeploymentConfiguration`

NewKustomizeDeploymentConfigurationWithDefaults instantiates a new KustomizeDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *KustomizeDeploymentConfiguration) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *KustomizeDeploymentConfiguration) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *KustomizeDeploymentConfiguration) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetDeploymentErrors

`func (o *KustomizeDeploymentConfiguration) GetDeploymentErrors() string`

GetDeploymentErrors returns the DeploymentErrors field if non-nil, zero value otherwise.

### GetDeploymentErrorsOk

`func (o *KustomizeDeploymentConfiguration) GetDeploymentErrorsOk() (*string, bool)`

GetDeploymentErrorsOk returns a tuple with the DeploymentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentErrors

`func (o *KustomizeDeploymentConfiguration) SetDeploymentErrors(v string)`

SetDeploymentErrors sets DeploymentErrors field to given value.

### HasDeploymentErrors

`func (o *KustomizeDeploymentConfiguration) HasDeploymentErrors() bool`

HasDeploymentErrors returns a boolean if a field has been set.

### GetOverlays

`func (o *KustomizeDeploymentConfiguration) GetOverlays() map[string]string`

GetOverlays returns the Overlays field if non-nil, zero value otherwise.

### GetOverlaysOk

`func (o *KustomizeDeploymentConfiguration) GetOverlaysOk() (*map[string]string, bool)`

GetOverlaysOk returns a tuple with the Overlays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlays

`func (o *KustomizeDeploymentConfiguration) SetOverlays(v map[string]string)`

SetOverlays sets Overlays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


