# UpdateInstanceStorageVolumeConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Storage Volume Config | 
**InstanceStorageIops** | Pointer to **string** | The IOPS provisioned for the configured instance storage type | [optional] 
**InstanceStorageThroughputMiBps** | Pointer to **string** | The throughput (in MiBps) provisioned for the configured instance storage type | [optional] 
**InstanceStorageType** | Pointer to **string** | The type of the storage for a compute instance | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateInstanceStorageVolumeConfigRequest

`func NewUpdateInstanceStorageVolumeConfigRequest(id string, serviceId string, token string, ) *UpdateInstanceStorageVolumeConfigRequest`

NewUpdateInstanceStorageVolumeConfigRequest instantiates a new UpdateInstanceStorageVolumeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceStorageVolumeConfigRequestWithDefaults

`func NewUpdateInstanceStorageVolumeConfigRequestWithDefaults() *UpdateInstanceStorageVolumeConfigRequest`

NewUpdateInstanceStorageVolumeConfigRequestWithDefaults instantiates a new UpdateInstanceStorageVolumeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *UpdateInstanceStorageVolumeConfigRequest) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *UpdateInstanceStorageVolumeConfigRequest) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *UpdateInstanceStorageVolumeConfigRequest) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateInstanceStorageVolumeConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateInstanceStorageVolumeConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


