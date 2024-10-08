# ResourceSpecLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **string** | The maximum amount of CPU that the container can use | [optional] 
**Memory** | Pointer to **string** | The maximum amount of memory that the container can use | [optional] 

## Methods

### NewResourceSpecLimits

`func NewResourceSpecLimits() *ResourceSpecLimits`

NewResourceSpecLimits instantiates a new ResourceSpecLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSpecLimitsWithDefaults

`func NewResourceSpecLimitsWithDefaults() *ResourceSpecLimits`

NewResourceSpecLimitsWithDefaults instantiates a new ResourceSpecLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceSpecLimits) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceSpecLimits) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceSpecLimits) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceSpecLimits) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ResourceSpecLimits) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceSpecLimits) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceSpecLimits) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceSpecLimits) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


