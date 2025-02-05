# DebugResourceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugData** | Pointer to **interface{}** | Individual debug data for the resource | [optional] 
**ResourceId** | **string** | ID of a resource | 

## Methods

### NewDebugResourceResult

`func NewDebugResourceResult(resourceId string, ) *DebugResourceResult`

NewDebugResourceResult instantiates a new DebugResourceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugResourceResultWithDefaults

`func NewDebugResourceResultWithDefaults() *DebugResourceResult`

NewDebugResourceResultWithDefaults instantiates a new DebugResourceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugData

`func (o *DebugResourceResult) GetDebugData() interface{}`

GetDebugData returns the DebugData field if non-nil, zero value otherwise.

### GetDebugDataOk

`func (o *DebugResourceResult) GetDebugDataOk() (*interface{}, bool)`

GetDebugDataOk returns a tuple with the DebugData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugData

`func (o *DebugResourceResult) SetDebugData(v interface{})`

SetDebugData sets DebugData field to given value.

### HasDebugData

`func (o *DebugResourceResult) HasDebugData() bool`

HasDebugData returns a boolean if a field has been set.

### SetDebugDataNil

`func (o *DebugResourceResult) SetDebugDataNil(b bool)`

 SetDebugDataNil sets the value for DebugData to be an explicit nil

### UnsetDebugData
`func (o *DebugResourceResult) UnsetDebugData()`

UnsetDebugData ensures that no value is present for DebugData, not even an explicit nil
### GetResourceId

`func (o *DebugResourceResult) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DebugResourceResult) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DebugResourceResult) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


