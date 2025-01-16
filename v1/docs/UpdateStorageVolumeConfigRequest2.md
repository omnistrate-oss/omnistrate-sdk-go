# UpdateStorageVolumeConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the context for the storage volume pool | [optional] 
**DisableBackup** | Pointer to **bool** | Disable backup for the storage volume | [optional] 
**Name** | Pointer to **string** | Name of the storage volume pool | [optional] 

## Methods

### NewUpdateStorageVolumeConfigRequest2

`func NewUpdateStorageVolumeConfigRequest2() *UpdateStorageVolumeConfigRequest2`

NewUpdateStorageVolumeConfigRequest2 instantiates a new UpdateStorageVolumeConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumeConfigRequest2WithDefaults

`func NewUpdateStorageVolumeConfigRequest2WithDefaults() *UpdateStorageVolumeConfigRequest2`

NewUpdateStorageVolumeConfigRequest2WithDefaults instantiates a new UpdateStorageVolumeConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStorageVolumeConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageVolumeConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageVolumeConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageVolumeConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableBackup

`func (o *UpdateStorageVolumeConfigRequest2) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *UpdateStorageVolumeConfigRequest2) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *UpdateStorageVolumeConfigRequest2) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *UpdateStorageVolumeConfigRequest2) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetName

`func (o *UpdateStorageVolumeConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageVolumeConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageVolumeConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageVolumeConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


