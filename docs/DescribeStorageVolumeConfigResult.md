# DescribeStorageVolumeConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | Pointer to **string** | The cloud provider name for fixed storage types | [optional] 
**ClusterStorageType** | Pointer to **string** | The type of the fixed storage for the cluster | [optional] 
**Description** | **string** | A brief description of the context for the storage volume pool | 
**DisableBackup** | Pointer to **bool** | Disable backup for the storage volume | [optional] 
**Id** | **string** | The storage volume config ID | 
**InstanceStorageIops** | Pointer to **string** | The IOPS provisioned for the configured instance storage type | [optional] 
**InstanceStorageSizeGi** | Pointer to **string** | The fixed storage size (in Gi) provisioned for the configured instance storage type | [optional] 
**InstanceStorageThroughputMiBps** | Pointer to **string** | The throughput (in MiBps) provisioned for the configured instance storage type | [optional] 
**InstanceStorageType** | Pointer to **string** | The type of the fixed storage for a compute instance | [optional] 
**Name** | **string** | Name of the storage volume pool | 
**ServiceId** | **string** | The service ID | 
**StorageResourceID** | Pointer to **string** | The storage resource ID | [optional] 

## Methods

### NewDescribeStorageVolumeConfigResult

`func NewDescribeStorageVolumeConfigResult(description string, id string, name string, serviceId string, ) *DescribeStorageVolumeConfigResult`

NewDescribeStorageVolumeConfigResult instantiates a new DescribeStorageVolumeConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeStorageVolumeConfigResultWithDefaults

`func NewDescribeStorageVolumeConfigResultWithDefaults() *DescribeStorageVolumeConfigResult`

NewDescribeStorageVolumeConfigResultWithDefaults instantiates a new DescribeStorageVolumeConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *DescribeStorageVolumeConfigResult) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *DescribeStorageVolumeConfigResult) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *DescribeStorageVolumeConfigResult) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.

### HasCloudProviderName

`func (o *DescribeStorageVolumeConfigResult) HasCloudProviderName() bool`

HasCloudProviderName returns a boolean if a field has been set.

### GetClusterStorageType

`func (o *DescribeStorageVolumeConfigResult) GetClusterStorageType() string`

GetClusterStorageType returns the ClusterStorageType field if non-nil, zero value otherwise.

### GetClusterStorageTypeOk

`func (o *DescribeStorageVolumeConfigResult) GetClusterStorageTypeOk() (*string, bool)`

GetClusterStorageTypeOk returns a tuple with the ClusterStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorageType

`func (o *DescribeStorageVolumeConfigResult) SetClusterStorageType(v string)`

SetClusterStorageType sets ClusterStorageType field to given value.

### HasClusterStorageType

`func (o *DescribeStorageVolumeConfigResult) HasClusterStorageType() bool`

HasClusterStorageType returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeStorageVolumeConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeStorageVolumeConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeStorageVolumeConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisableBackup

`func (o *DescribeStorageVolumeConfigResult) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *DescribeStorageVolumeConfigResult) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *DescribeStorageVolumeConfigResult) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *DescribeStorageVolumeConfigResult) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetId

`func (o *DescribeStorageVolumeConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeStorageVolumeConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeStorageVolumeConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceStorageIops

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageIops() string`

GetInstanceStorageIops returns the InstanceStorageIops field if non-nil, zero value otherwise.

### GetInstanceStorageIopsOk

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageIopsOk() (*string, bool)`

GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageIops

`func (o *DescribeStorageVolumeConfigResult) SetInstanceStorageIops(v string)`

SetInstanceStorageIops sets InstanceStorageIops field to given value.

### HasInstanceStorageIops

`func (o *DescribeStorageVolumeConfigResult) HasInstanceStorageIops() bool`

HasInstanceStorageIops returns a boolean if a field has been set.

### GetInstanceStorageSizeGi

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageSizeGi() string`

GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field if non-nil, zero value otherwise.

### GetInstanceStorageSizeGiOk

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageSizeGiOk() (*string, bool)`

GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageSizeGi

`func (o *DescribeStorageVolumeConfigResult) SetInstanceStorageSizeGi(v string)`

SetInstanceStorageSizeGi sets InstanceStorageSizeGi field to given value.

### HasInstanceStorageSizeGi

`func (o *DescribeStorageVolumeConfigResult) HasInstanceStorageSizeGi() bool`

HasInstanceStorageSizeGi returns a boolean if a field has been set.

### GetInstanceStorageThroughputMiBps

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageThroughputMiBps() string`

GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field if non-nil, zero value otherwise.

### GetInstanceStorageThroughputMiBpsOk

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageThroughputMiBpsOk() (*string, bool)`

GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageThroughputMiBps

`func (o *DescribeStorageVolumeConfigResult) SetInstanceStorageThroughputMiBps(v string)`

SetInstanceStorageThroughputMiBps sets InstanceStorageThroughputMiBps field to given value.

### HasInstanceStorageThroughputMiBps

`func (o *DescribeStorageVolumeConfigResult) HasInstanceStorageThroughputMiBps() bool`

HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.

### GetInstanceStorageType

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageType() string`

GetInstanceStorageType returns the InstanceStorageType field if non-nil, zero value otherwise.

### GetInstanceStorageTypeOk

`func (o *DescribeStorageVolumeConfigResult) GetInstanceStorageTypeOk() (*string, bool)`

GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageType

`func (o *DescribeStorageVolumeConfigResult) SetInstanceStorageType(v string)`

SetInstanceStorageType sets InstanceStorageType field to given value.

### HasInstanceStorageType

`func (o *DescribeStorageVolumeConfigResult) HasInstanceStorageType() bool`

HasInstanceStorageType returns a boolean if a field has been set.

### GetName

`func (o *DescribeStorageVolumeConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeStorageVolumeConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeStorageVolumeConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *DescribeStorageVolumeConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeStorageVolumeConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeStorageVolumeConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageResourceID

`func (o *DescribeStorageVolumeConfigResult) GetStorageResourceID() string`

GetStorageResourceID returns the StorageResourceID field if non-nil, zero value otherwise.

### GetStorageResourceIDOk

`func (o *DescribeStorageVolumeConfigResult) GetStorageResourceIDOk() (*string, bool)`

GetStorageResourceIDOk returns a tuple with the StorageResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceID

`func (o *DescribeStorageVolumeConfigResult) SetStorageResourceID(v string)`

SetStorageResourceID sets StorageResourceID field to given value.

### HasStorageResourceID

`func (o *DescribeStorageVolumeConfigResult) HasStorageResourceID() bool`

HasStorageResourceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


