# GenericResourceDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Image for the generic resource deployment | [optional] 
**PodEvents** | Pointer to [**map[string][]PodEvent**](array.md) | Events associated with the pods in the generic resource deployment | [optional] 
**PodStatus** | Pointer to **map[string]string** | Status of the pods in the generic resource deployment | [optional] 
**PodToHostMapping** | Pointer to **map[string]string** | Mapping of pods to k8s node names for the generic resource deployment | [optional] 

## Methods

### NewGenericResourceDeploymentConfiguration

`func NewGenericResourceDeploymentConfiguration() *GenericResourceDeploymentConfiguration`

NewGenericResourceDeploymentConfiguration instantiates a new GenericResourceDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericResourceDeploymentConfigurationWithDefaults

`func NewGenericResourceDeploymentConfigurationWithDefaults() *GenericResourceDeploymentConfiguration`

NewGenericResourceDeploymentConfigurationWithDefaults instantiates a new GenericResourceDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *GenericResourceDeploymentConfiguration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GenericResourceDeploymentConfiguration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GenericResourceDeploymentConfiguration) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GenericResourceDeploymentConfiguration) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPodEvents

`func (o *GenericResourceDeploymentConfiguration) GetPodEvents() map[string][]PodEvent`

GetPodEvents returns the PodEvents field if non-nil, zero value otherwise.

### GetPodEventsOk

`func (o *GenericResourceDeploymentConfiguration) GetPodEventsOk() (*map[string][]PodEvent, bool)`

GetPodEventsOk returns a tuple with the PodEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodEvents

`func (o *GenericResourceDeploymentConfiguration) SetPodEvents(v map[string][]PodEvent)`

SetPodEvents sets PodEvents field to given value.

### HasPodEvents

`func (o *GenericResourceDeploymentConfiguration) HasPodEvents() bool`

HasPodEvents returns a boolean if a field has been set.

### GetPodStatus

`func (o *GenericResourceDeploymentConfiguration) GetPodStatus() map[string]string`

GetPodStatus returns the PodStatus field if non-nil, zero value otherwise.

### GetPodStatusOk

`func (o *GenericResourceDeploymentConfiguration) GetPodStatusOk() (*map[string]string, bool)`

GetPodStatusOk returns a tuple with the PodStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodStatus

`func (o *GenericResourceDeploymentConfiguration) SetPodStatus(v map[string]string)`

SetPodStatus sets PodStatus field to given value.

### HasPodStatus

`func (o *GenericResourceDeploymentConfiguration) HasPodStatus() bool`

HasPodStatus returns a boolean if a field has been set.

### GetPodToHostMapping

`func (o *GenericResourceDeploymentConfiguration) GetPodToHostMapping() map[string]string`

GetPodToHostMapping returns the PodToHostMapping field if non-nil, zero value otherwise.

### GetPodToHostMappingOk

`func (o *GenericResourceDeploymentConfiguration) GetPodToHostMappingOk() (*map[string]string, bool)`

GetPodToHostMappingOk returns a tuple with the PodToHostMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodToHostMapping

`func (o *GenericResourceDeploymentConfiguration) SetPodToHostMapping(v map[string]string)`

SetPodToHostMapping sets PodToHostMapping field to given value.

### HasPodToHostMapping

`func (o *GenericResourceDeploymentConfiguration) HasPodToHostMapping() bool`

HasPodToHostMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


