# InfrastructureGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocatable** | Pointer to **int64** |  | [optional] 
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**DriverVersion** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**InstanceAllocated** | Pointer to **int64** |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**MemoryMiB** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**RuntimeVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureGPU

`func NewInfrastructureGPU() *InfrastructureGPU`

NewInfrastructureGPU instantiates a new InfrastructureGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureGPUWithDefaults

`func NewInfrastructureGPUWithDefaults() *InfrastructureGPU`

NewInfrastructureGPUWithDefaults instantiates a new InfrastructureGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatable

`func (o *InfrastructureGPU) GetAllocatable() int64`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *InfrastructureGPU) GetAllocatableOk() (*int64, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *InfrastructureGPU) SetAllocatable(v int64)`

SetAllocatable sets Allocatable field to given value.

### HasAllocatable

`func (o *InfrastructureGPU) HasAllocatable() bool`

HasAllocatable returns a boolean if a field has been set.

### GetAvailability

`func (o *InfrastructureGPU) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureGPU) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureGPU) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureGPU) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCount

`func (o *InfrastructureGPU) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InfrastructureGPU) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InfrastructureGPU) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *InfrastructureGPU) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDriverVersion

`func (o *InfrastructureGPU) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *InfrastructureGPU) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *InfrastructureGPU) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *InfrastructureGPU) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetFamily

`func (o *InfrastructureGPU) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *InfrastructureGPU) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *InfrastructureGPU) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *InfrastructureGPU) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetInstanceAllocated

`func (o *InfrastructureGPU) GetInstanceAllocated() int64`

GetInstanceAllocated returns the InstanceAllocated field if non-nil, zero value otherwise.

### GetInstanceAllocatedOk

`func (o *InfrastructureGPU) GetInstanceAllocatedOk() (*int64, bool)`

GetInstanceAllocatedOk returns a tuple with the InstanceAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAllocated

`func (o *InfrastructureGPU) SetInstanceAllocated(v int64)`

SetInstanceAllocated sets InstanceAllocated field to given value.

### HasInstanceAllocated

`func (o *InfrastructureGPU) HasInstanceAllocated() bool`

HasInstanceAllocated returns a boolean if a field has been set.

### GetMachine

`func (o *InfrastructureGPU) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *InfrastructureGPU) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *InfrastructureGPU) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *InfrastructureGPU) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemoryMiB

`func (o *InfrastructureGPU) GetMemoryMiB() string`

GetMemoryMiB returns the MemoryMiB field if non-nil, zero value otherwise.

### GetMemoryMiBOk

`func (o *InfrastructureGPU) GetMemoryMiBOk() (*string, bool)`

GetMemoryMiBOk returns a tuple with the MemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiB

`func (o *InfrastructureGPU) SetMemoryMiB(v string)`

SetMemoryMiB sets MemoryMiB field to given value.

### HasMemoryMiB

`func (o *InfrastructureGPU) HasMemoryMiB() bool`

HasMemoryMiB returns a boolean if a field has been set.

### GetProduct

`func (o *InfrastructureGPU) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *InfrastructureGPU) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *InfrastructureGPU) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *InfrastructureGPU) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRuntimeVersion

`func (o *InfrastructureGPU) GetRuntimeVersion() string`

GetRuntimeVersion returns the RuntimeVersion field if non-nil, zero value otherwise.

### GetRuntimeVersionOk

`func (o *InfrastructureGPU) GetRuntimeVersionOk() (*string, bool)`

GetRuntimeVersionOk returns a tuple with the RuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeVersion

`func (o *InfrastructureGPU) SetRuntimeVersion(v string)`

SetRuntimeVersion sets RuntimeVersion field to given value.

### HasRuntimeVersion

`func (o *InfrastructureGPU) HasRuntimeVersion() bool`

HasRuntimeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


