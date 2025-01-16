# UpdateStorageVolumeConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the context for the storage volume pool | [optional] 
**DisableBackup** | Pointer to **bool** | Disable backup for the storage volume | [optional] 
**Id** | **string** | ID of a Storage Volume Config | 
**Name** | Pointer to **string** | Name of the storage volume pool | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateStorageVolumeConfigRequest

`func NewUpdateStorageVolumeConfigRequest(id string, serviceId string, token string, ) *UpdateStorageVolumeConfigRequest`

NewUpdateStorageVolumeConfigRequest instantiates a new UpdateStorageVolumeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumeConfigRequestWithDefaults

`func NewUpdateStorageVolumeConfigRequestWithDefaults() *UpdateStorageVolumeConfigRequest`

NewUpdateStorageVolumeConfigRequestWithDefaults instantiates a new UpdateStorageVolumeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStorageVolumeConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageVolumeConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageVolumeConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageVolumeConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableBackup

`func (o *UpdateStorageVolumeConfigRequest) GetDisableBackup() bool`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *UpdateStorageVolumeConfigRequest) GetDisableBackupOk() (*bool, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *UpdateStorageVolumeConfigRequest) SetDisableBackup(v bool)`

SetDisableBackup sets DisableBackup field to given value.

### HasDisableBackup

`func (o *UpdateStorageVolumeConfigRequest) HasDisableBackup() bool`

HasDisableBackup returns a boolean if a field has been set.

### GetId

`func (o *UpdateStorageVolumeConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStorageVolumeConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStorageVolumeConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateStorageVolumeConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageVolumeConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageVolumeConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageVolumeConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateStorageVolumeConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateStorageVolumeConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *UpdateStorageVolumeConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateStorageVolumeConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateStorageVolumeConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


