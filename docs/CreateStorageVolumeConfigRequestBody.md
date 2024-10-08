# CreateStorageVolumeConfigRequestBody

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

### NewCreateStorageVolumeConfigRequestBody

`func NewCreateStorageVolumeConfigRequestBody(description string, name string, ) *CreateStorageVolumeConfigRequestBody`

NewCreateStorageVolumeConfigRequestBody instantiates a new CreateStorageVolumeConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageVolumeConfigRequestBodyWithDefaults

`func NewCreateStorageVolumeConfigRequestBodyWithDefaults() *CreateStorageVolumeConfigRequestBody`

NewCreateStorageVolumeConfigRequestBodyWithDefaults instantiates a new CreateStorageVolumeConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterStorageType

`func (o *CreateStorageVolumeConfigRequestBody) GetClusterStorageType() string`

GetClusterStorageType returns the ClusterStorageType field if non-nil, zero value otherwise.

### GetClusterStorageTypeOk

`func (o *CreateStorageVolumeConfigRequestBody) GetClusterStorageTypeOk() (*string, bool)`

GetClusterStorageTypeOk returns a tuple with the ClusterStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorageType

`func (o *CreateStorageVolumeConfigRequestBody) SetClusterStorageType(v string)`

SetClusterStorageType sets ClusterStorageType field to given value.

### HasClusterStorageType

`func (o *CreateStorageVolumeConfigRequestBody) HasClusterStorageType() bool`

HasClusterStorageType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStorageVolumeConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageVolumeConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageVolumeConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisableBackup

`func (o *CreateStorageVolumeConfigRequestBody) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *CreateStorageVolumeConfigRequestBody) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *CreateStorageVolumeConfigRequestBody) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *CreateStorageVolumeConfigRequestBody) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageSizeGi() string`

GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field if non-nil, zero value otherwise.

### GetInstanceStorageSizeGiOk

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageSizeGiOk() (*string, bool)`

GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageSizeGi(v string)`

SetInstanceStorageSizeGi sets InstanceStorageSizeGi field to given value.

### HasInstanceStorageSizeGi

`func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageSizeGi() bool`

HasInstanceStorageSizeGi returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.

### GetName

`func (o *CreateStorageVolumeConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorageVolumeConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorageVolumeConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetStorageResourceID

`func (o *CreateStorageVolumeConfigRequestBody) GetStorageResourceID() string`

GetStorageResourceID returns the StorageResourceID field if non-nil, zero value otherwise.

### GetStorageResourceIDOk

`func (o *CreateStorageVolumeConfigRequestBody) GetStorageResourceIDOk() (*string, bool)`

GetStorageResourceIDOk returns a tuple with the StorageResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceID

`func (o *CreateStorageVolumeConfigRequestBody) SetStorageResourceID(v string)`

SetStorageResourceID sets StorageResourceID field to given value.

### HasStorageResourceID

`func (o *CreateStorageVolumeConfigRequestBody) HasStorageResourceID() bool`

HasStorageResourceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


