# ResourceDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterMap** | Pointer to **map[string]string** | A map of the source parameter to the resource dependency parameter | [optional] 
**ResourceId** | **string** | ID of a resource | 

## Methods

### NewResourceDependency

`func NewResourceDependency(resourceId string, ) *ResourceDependency`

NewResourceDependency instantiates a new ResourceDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDependencyWithDefaults

`func NewResourceDependencyWithDefaults() *ResourceDependency`

NewResourceDependencyWithDefaults instantiates a new ResourceDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterMap

`func (o *ResourceDependency) GetParameterMap() map[string]string`

GetParameterMap returns the ParameterMap field if non-nil, zero value otherwise.

### GetParameterMapOk

`func (o *ResourceDependency) GetParameterMapOk() (*map[string]string, bool)`

GetParameterMapOk returns a tuple with the ParameterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterMap

`func (o *ResourceDependency) SetParameterMap(v map[string]string)`

SetParameterMap sets ParameterMap field to given value.

### HasParameterMap

`func (o *ResourceDependency) HasParameterMap() bool`

HasParameterMap returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceDependency) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceDependency) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceDependency) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


