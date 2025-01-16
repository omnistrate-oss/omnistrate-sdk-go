# RemoveStorageVolumeConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | Pointer to **string** | The specific mount path to remove. If not specified, all mount paths for the storage volume config will be removed | [optional] 

## Methods

### NewRemoveStorageVolumeConfigRequest2

`func NewRemoveStorageVolumeConfigRequest2() *RemoveStorageVolumeConfigRequest2`

NewRemoveStorageVolumeConfigRequest2 instantiates a new RemoveStorageVolumeConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveStorageVolumeConfigRequest2WithDefaults

`func NewRemoveStorageVolumeConfigRequest2WithDefaults() *RemoveStorageVolumeConfigRequest2`

NewRemoveStorageVolumeConfigRequest2WithDefaults instantiates a new RemoveStorageVolumeConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPath

`func (o *RemoveStorageVolumeConfigRequest2) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *RemoveStorageVolumeConfigRequest2) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *RemoveStorageVolumeConfigRequest2) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *RemoveStorageVolumeConfigRequest2) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


