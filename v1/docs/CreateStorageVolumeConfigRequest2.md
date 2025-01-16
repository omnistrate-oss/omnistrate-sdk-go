# CreateStorageVolumeConfigRequest2

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
**StorageResourceID** | Pointer to **string** | The storage resource ID | [optional] 

## Methods

### NewCreateStorageVolumeConfigRequest2

`func NewCreateStorageVolumeConfigRequest2(description string, name string, ) *CreateStorageVolumeConfigRequest2`

NewCreateStorageVolumeConfigRequest2 instantiates a new CreateStorageVolumeConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageVolumeConfigRequest2WithDefaults

`func NewCreateStorageVolumeConfigRequest2WithDefaults() *CreateStorageVolumeConfigRequest2`

NewCreateStorageVolumeConfigRequest2WithDefaults instantiates a new CreateStorageVolumeConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterStorageType

`func (o *CreateStorageVolumeConfigRequest2) GetClusterStorageType() string`

GetClusterStorageType returns the ClusterStorageType field if non-nil, zero value otherwise.

### GetClusterStorageTypeOk

`func (o *CreateStorageVolumeConfigRequest2) GetClusterStorageTypeOk() (*string, bool)`

GetClusterStorageTypeOk returns a tuple with the ClusterStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorageType

`func (o *CreateStorageVolumeConfigRequest2) SetClusterStorageType(v string)`

SetClusterStorageType sets ClusterStorageType field to given value.

### HasClusterStorageType

`func (o *CreateStorageVolumeConfigRequest2) HasClusterStorageType() bool`

HasClusterStorageType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStorageVolumeConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageVolumeConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageVolumeConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisableBackup

`func (o *CreateStorageVolumeConfigRequest2) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *CreateStorageVolumeConfigRequest2) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *CreateStorageVolumeConfigRequest2) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *CreateStorageVolumeConfigRequest2) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest2) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequest2) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageSizeGi() string`

GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field if non-nil, zero value otherwise.

### GetInstanceStorageSizeGiOk

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageSizeGiOk() (*string, bool)`

GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest2) SetInstanceStorageSizeGi(v string)`

SetInstanceStorageSizeGi sets InstanceStorageSizeGi field to given value.

### HasInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequest2) HasInstanceStorageSizeGi() bool`

HasInstanceStorageSizeGi returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest2) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequest2) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *CreateStorageVolumeConfigRequest2) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest2) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *CreateStorageVolumeConfigRequest2) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.

### GetName

`func (o *CreateStorageVolumeConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorageVolumeConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorageVolumeConfigRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetStorageResourceID

`func (o *CreateStorageVolumeConfigRequest2) GetStorageResourceID() string`

GetStorageResourceID returns the StorageResourceID field if non-nil, zero value otherwise.

### GetStorageResourceIDOk

`func (o *CreateStorageVolumeConfigRequest2) GetStorageResourceIDOk() (*string, bool)`

GetStorageResourceIDOk returns a tuple with the StorageResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceID

`func (o *CreateStorageVolumeConfigRequest2) SetStorageResourceID(v string)`

SetStorageResourceID sets StorageResourceID field to given value.

### HasStorageResourceID

`func (o *CreateStorageVolumeConfigRequest2) HasStorageResourceID() bool`

HasStorageResourceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


