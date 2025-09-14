# DebugResourceInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | ID of a Resource Instance | 
**ResourcesDebug** | Pointer to [**map[string]DebugResourceResult**](DebugResourceResult.md) | The debug information for individual instance resources | [optional] 

## Methods

### NewDebugResourceInstanceResult

`func NewDebugResourceInstanceResult(instanceId string, ) *DebugResourceInstanceResult`

NewDebugResourceInstanceResult instantiates a new DebugResourceInstanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugResourceInstanceResultWithDefaults

`func NewDebugResourceInstanceResultWithDefaults() *DebugResourceInstanceResult`

NewDebugResourceInstanceResultWithDefaults instantiates a new DebugResourceInstanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *DebugResourceInstanceResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DebugResourceInstanceResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DebugResourceInstanceResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetResourcesDebug

`func (o *DebugResourceInstanceResult) GetResourcesDebug() map[string]DebugResourceResult`

GetResourcesDebug returns the ResourcesDebug field if non-nil, zero value otherwise.

### GetResourcesDebugOk

`func (o *DebugResourceInstanceResult) GetResourcesDebugOk() (*map[string]DebugResourceResult, bool)`

GetResourcesDebugOk returns a tuple with the ResourcesDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesDebug

`func (o *DebugResourceInstanceResult) SetResourcesDebug(v map[string]DebugResourceResult)`

SetResourcesDebug sets ResourcesDebug field to given value.

### HasResourcesDebug

`func (o *DebugResourceInstanceResult) HasResourcesDebug() bool`

HasResourcesDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


