# UpdateStorageVolumeConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the context for the storage volume pool | [optional] 
**DisableBackup** | Pointer to **bool** | Disable backup for the storage volume | [optional] 
**Name** | Pointer to **string** | Name of the storage volume pool | [optional] 

## Methods

### NewUpdateStorageVolumeConfigRequestBody

`func NewUpdateStorageVolumeConfigRequestBody() *UpdateStorageVolumeConfigRequestBody`

NewUpdateStorageVolumeConfigRequestBody instantiates a new UpdateStorageVolumeConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumeConfigRequestBodyWithDefaults

`func NewUpdateStorageVolumeConfigRequestBodyWithDefaults() *UpdateStorageVolumeConfigRequestBody`

NewUpdateStorageVolumeConfigRequestBodyWithDefaults instantiates a new UpdateStorageVolumeConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStorageVolumeConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageVolumeConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageVolumeConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageVolumeConfigRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableBackup

`func (o *UpdateStorageVolumeConfigRequestBody) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *UpdateStorageVolumeConfigRequestBody) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *UpdateStorageVolumeConfigRequestBody) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *UpdateStorageVolumeConfigRequestBody) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetName

`func (o *UpdateStorageVolumeConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageVolumeConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageVolumeConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageVolumeConfigRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


