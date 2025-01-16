# UpdateInstanceStorageVolumeConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceStorageIops** | Pointer to **string** | The IOPS provisioned for the configured instance storage type | [optional] 
**InstanceStorageThroughputMiBps** | Pointer to **string** | The throughput (in MiBps) provisioned for the configured instance storage type | [optional] 
**InstanceStorageType** | Pointer to **string** | The type of the storage for a compute instance | [optional] 

## Methods

### NewUpdateInstanceStorageVolumeConfigRequest2

`func NewUpdateInstanceStorageVolumeConfigRequest2() *UpdateInstanceStorageVolumeConfigRequest2`

NewUpdateInstanceStorageVolumeConfigRequest2 instantiates a new UpdateInstanceStorageVolumeConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceStorageVolumeConfigRequest2WithDefaults

`func NewUpdateInstanceStorageVolumeConfigRequest2WithDefaults() *UpdateInstanceStorageVolumeConfigRequest2`

NewUpdateInstanceStorageVolumeConfigRequest2WithDefaults instantiates a new UpdateInstanceStorageVolumeConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest2) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest2) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest2) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest2) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *UpdateInstanceStorageVolumeConfigRequest2) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest2) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest2) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


