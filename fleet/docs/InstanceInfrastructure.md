# InstanceInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**InfrastructureCompute**](InfrastructureCompute.md) |  | [optional] 
**Networking** | Pointer to [**InfrastructureNetworking**](InfrastructureNetworking.md) |  | [optional] 
**Release** | Pointer to [**InfrastructureRelease**](InfrastructureRelease.md) |  | [optional] 
**Source** | Pointer to [**InfrastructureSourceInfo**](InfrastructureSourceInfo.md) |  | [optional] 
**Storage** | Pointer to [**InfrastructureStorage**](InfrastructureStorage.md) |  | [optional] 

## Methods

### NewInstanceInfrastructure

`func NewInstanceInfrastructure() *InstanceInfrastructure`

NewInstanceInfrastructure instantiates a new InstanceInfrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInfrastructureWithDefaults

`func NewInstanceInfrastructureWithDefaults() *InstanceInfrastructure`

NewInstanceInfrastructureWithDefaults instantiates a new InstanceInfrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *InstanceInfrastructure) GetCompute() InfrastructureCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *InstanceInfrastructure) GetComputeOk() (*InfrastructureCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *InstanceInfrastructure) SetCompute(v InfrastructureCompute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *InstanceInfrastructure) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetNetworking

`func (o *InstanceInfrastructure) GetNetworking() InfrastructureNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *InstanceInfrastructure) GetNetworkingOk() (*InfrastructureNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *InstanceInfrastructure) SetNetworking(v InfrastructureNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *InstanceInfrastructure) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetRelease

`func (o *InstanceInfrastructure) GetRelease() InfrastructureRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *InstanceInfrastructure) GetReleaseOk() (*InfrastructureRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *InstanceInfrastructure) SetRelease(v InfrastructureRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *InstanceInfrastructure) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSource

`func (o *InstanceInfrastructure) GetSource() InfrastructureSourceInfo`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InstanceInfrastructure) GetSourceOk() (*InfrastructureSourceInfo, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InstanceInfrastructure) SetSource(v InfrastructureSourceInfo)`

SetSource sets Source field to given value.

### HasSource

`func (o *InstanceInfrastructure) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStorage

`func (o *InstanceInfrastructure) GetStorage() InfrastructureStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceInfrastructure) GetStorageOk() (*InfrastructureStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceInfrastructure) SetStorage(v InfrastructureStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstanceInfrastructure) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


