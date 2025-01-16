# CreateStorageVolumeConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterStorageType** | Pointer to **string** | The type of the fixed storage for the cluster | [optional] 
**Description** | **string** | A brief description of the context for the storage volume pool | 
**DisableBackup** | Pointer to **bool** | Disable backup for the storage volume | [optional] 
**InstanceStorageIops** | Pointer to **string** | The IOPS provisioned for the configured instance storage type | [optional] 
**InstanceStorageSizeGi** | Pointer to **string** | The storage size (in Gi) provisioned for the configured instance storage type | [optional] 
**InstanceStorageThroughputMiBps** | Pointer to **string** | The throughput (in MiBps) provisioned for the configured instance storage type | [optional] 
**InstanceStorageType** | Pointer to **string** | The type of the storage for a compute instance | [optional] 
**Name** | **string** | Name of the storage volume pool | 
**ServiceId** | **string** | ID of a Service | 
**StorageResourceID** | Pointer to **string** | The storage resource ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateStorageVolumeConfigRequest

`func NewCreateStorageVolumeConfigRequest(description string, name string, serviceId string, token string, ) *CreateStorageVolumeConfigRequest`

NewCreateStorageVolumeConfigRequest instantiates a new CreateStorageVolumeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageVolumeConfigRequestWithDefaults

`func NewCreateStorageVolumeConfigRequestWithDefaults() *CreateStorageVolumeConfigRequest`

NewCreateStorageVolumeConfigRequestWithDefaults instantiates a new CreateStorageVolumeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterStorageType

`func (o *CreateStorageVolumeConfigRequest) GetClusterStorageType() string`

GetClusterStorageType returns the ClusterStorageType field if non-nil, zero value otherwise.

### GetClusterStorageTypeOk

`func (o *CreateStorageVolumeConfigRequest) GetClusterStorageTypeOk() (*string, bool)`

GetClusterStorageTypeOk returns a tuple with the ClusterStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorageType

`func (o *CreateStorageVolumeConfigRequest) SetClusterStorageType(v string)`

SetClusterStorageType sets ClusterStorageType field to given value.

### HasClusterStorageType

`func (o *CreateStorageVolumeConfigRequest) HasClusterStorageType() bool`

HasClusterStorageType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStorageVolumeConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageVolumeConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageVolumeConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisableBackup

`func (o *CreateStorageVolumeConfigRequest) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *CreateStorageVolumeConfigRequest) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *CreateStorageVolumeConfigRequest) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *CreateStorageVolumeConfigRequest) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageSizeGi() string`

GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field if non-nil, zero value otherwise.

### GetInstanceStorageSizeGiOk

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageSizeGiOk() (*string, bool)`

GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest) SetInstanceStorageSizeGi(v string)`

SetInstanceStorageSizeGi sets InstanceStorageSizeGi field to given value.

### HasInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest) HasInstanceStorageSizeGi() bool`

HasInstanceStorageSizeGi returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *CreateStorageVolumeConfigRequest) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.

### GetName

`func (o *CreateStorageVolumeConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorageVolumeConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorageVolumeConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *CreateStorageVolumeConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateStorageVolumeConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageResourceID

`func (o *CreateStorageVolumeConfigRequest) GetStorageResourceID() string`

GetStorageResourceID returns the StorageResourceID field if non-nil, zero value otherwise.

### GetStorageResourceIDOk

`func (o *CreateStorageVolumeConfigRequest) GetStorageResourceIDOk() (*string, bool)`

GetStorageResourceIDOk returns a tuple with the StorageResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceID

`func (o *CreateStorageVolumeConfigRequest) SetStorageResourceID(v string)`

SetStorageResourceID sets StorageResourceID field to given value.

### HasStorageResourceID

`func (o *CreateStorageVolumeConfigRequest) HasStorageResourceID() bool`

HasStorageResourceID returns a boolean if a field has been set.

### GetToken

`func (o *CreateStorageVolumeConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateStorageVolumeConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateStorageVolumeConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


