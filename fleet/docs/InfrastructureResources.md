# InfrastructureResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to **int64** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Pods** | Pointer to **int64** |  | [optional] 

## Methods

### NewInfrastructureResources

`func NewInfrastructureResources() *InfrastructureResources`

NewInfrastructureResources instantiates a new InfrastructureResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureResourcesWithDefaults

`func NewInfrastructureResourcesWithDefaults() *InfrastructureResources`

NewInfrastructureResourcesWithDefaults instantiates a new InfrastructureResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *InfrastructureResources) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InfrastructureResources) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InfrastructureResources) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InfrastructureResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetGpu

`func (o *InfrastructureResources) GetGpu() int64`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *InfrastructureResources) GetGpuOk() (*int64, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *InfrastructureResources) SetGpu(v int64)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *InfrastructureResources) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetMemory

`func (o *InfrastructureResources) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InfrastructureResources) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InfrastructureResources) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InfrastructureResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetPods

`func (o *InfrastructureResources) GetPods() int64`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *InfrastructureResources) GetPodsOk() (*int64, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *InfrastructureResources) SetPods(v int64)`

SetPods sets Pods field to given value.

### HasPods

`func (o *InfrastructureResources) HasPods() bool`

HasPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


