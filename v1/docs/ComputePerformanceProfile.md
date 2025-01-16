# ComputePerformanceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | Processor architecture | 
**MaxInstanceAggregateThroughputMBps** | **int64** | Max throughput (in MBps) aggregated across all network interfaces per compute node required | 
**MaxSanAggregateThroughputMBps** | **int64** | Max dedicated network throughput (in MBps) aggregated across all storage volumes per compute node required | 
**MemoryMiB** | **int64** | Amount of memory (in Mi) per compute node required | 
**Tuning** | Pointer to **string** | Tune compute in a specific dimension | [optional] [default to "BALANCED"]
**VirtualCores** | **int64** | Number of virtual CPU cores required | 

## Methods

### NewComputePerformanceProfile

`func NewComputePerformanceProfile(architecture string, maxInstanceAggregateThroughputMBps int64, maxSanAggregateThroughputMBps int64, memoryMiB int64, virtualCores int64, ) *ComputePerformanceProfile`

NewComputePerformanceProfile instantiates a new ComputePerformanceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePerformanceProfileWithDefaults

`func NewComputePerformanceProfileWithDefaults() *ComputePerformanceProfile`

NewComputePerformanceProfileWithDefaults instantiates a new ComputePerformanceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ComputePerformanceProfile) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ComputePerformanceProfile) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ComputePerformanceProfile) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetMaxInstanceAggregateThroughputMBps

`func (o *ComputePerformanceProfile) GetMaxInstanceAggregateThroughputMBps() int64`

GetMaxInstanceAggregateThroughputMBps returns the MaxInstanceAggregateThroughputMBps field if non-nil, zero value otherwise.

### GetMaxInstanceAggregateThroughputMBpsOk

`func (o *ComputePerformanceProfile) GetMaxInstanceAggregateThroughputMBpsOk() (*int64, bool)`

GetMaxInstanceAggregateThroughputMBpsOk returns a tuple with the MaxInstanceAggregateThroughputMBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceAggregateThroughputMBps

`func (o *ComputePerformanceProfile) SetMaxInstanceAggregateThroughputMBps(v int64)`

SetMaxInstanceAggregateThroughputMBps sets MaxInstanceAggregateThroughputMBps field to given value.


### GetMaxSanAggregateThroughputMBps

`func (o *ComputePerformanceProfile) GetMaxSanAggregateThroughputMBps() int64`

GetMaxSanAggregateThroughputMBps returns the MaxSanAggregateThroughputMBps field if non-nil, zero value otherwise.

### GetMaxSanAggregateThroughputMBpsOk

`func (o *ComputePerformanceProfile) GetMaxSanAggregateThroughputMBpsOk() (*int64, bool)`

GetMaxSanAggregateThroughputMBpsOk returns a tuple with the MaxSanAggregateThroughputMBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSanAggregateThroughputMBps

`func (o *ComputePerformanceProfile) SetMaxSanAggregateThroughputMBps(v int64)`

SetMaxSanAggregateThroughputMBps sets MaxSanAggregateThroughputMBps field to given value.


### GetMemoryMiB

`func (o *ComputePerformanceProfile) GetMemoryMiB() int64`

GetMemoryMiB returns the MemoryMiB field if non-nil, zero value otherwise.

### GetMemoryMiBOk

`func (o *ComputePerformanceProfile) GetMemoryMiBOk() (*int64, bool)`

GetMemoryMiBOk returns a tuple with the MemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiB

`func (o *ComputePerformanceProfile) SetMemoryMiB(v int64)`

SetMemoryMiB sets MemoryMiB field to given value.


### GetTuning

`func (o *ComputePerformanceProfile) GetTuning() string`

GetTuning returns the Tuning field if non-nil, zero value otherwise.

### GetTuningOk

`func (o *ComputePerformanceProfile) GetTuningOk() (*string, bool)`

GetTuningOk returns a tuple with the Tuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuning

`func (o *ComputePerformanceProfile) SetTuning(v string)`

SetTuning sets Tuning field to given value.

### HasTuning

`func (o *ComputePerformanceProfile) HasTuning() bool`

HasTuning returns a boolean if a field has been set.

### GetVirtualCores

`func (o *ComputePerformanceProfile) GetVirtualCores() int64`

GetVirtualCores returns the VirtualCores field if non-nil, zero value otherwise.

### GetVirtualCoresOk

`func (o *ComputePerformanceProfile) GetVirtualCoresOk() (*int64, bool)`

GetVirtualCoresOk returns a tuple with the VirtualCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCores

`func (o *ComputePerformanceProfile) SetVirtualCores(v int64)`

SetVirtualCores sets VirtualCores field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


