# AddStorageVolumeConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Storage Config | 
**MountPath** | Pointer to **string** | The path where the storage volume will be mounted | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StorageVolumeConfigId** | **string** | ID of a Storage Volume Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAddStorageVolumeConfigRequest

`func NewAddStorageVolumeConfigRequest(id string, serviceId string, storageVolumeConfigId string, token string, ) *AddStorageVolumeConfigRequest`

NewAddStorageVolumeConfigRequest instantiates a new AddStorageVolumeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageVolumeConfigRequestWithDefaults

`func NewAddStorageVolumeConfigRequestWithDefaults() *AddStorageVolumeConfigRequest`

NewAddStorageVolumeConfigRequestWithDefaults instantiates a new AddStorageVolumeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddStorageVolumeConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddStorageVolumeConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddStorageVolumeConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMountPath

`func (o *AddStorageVolumeConfigRequest) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *AddStorageVolumeConfigRequest) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *AddStorageVolumeConfigRequest) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *AddStorageVolumeConfigRequest) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetServiceId

`func (o *AddStorageVolumeConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddStorageVolumeConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageVolumeConfigId

`func (o *AddStorageVolumeConfigRequest) GetStorageVolumeConfigId() string`

GetStorageVolumeConfigId returns the StorageVolumeConfigId field if non-nil, zero value otherwise.

### GetStorageVolumeConfigIdOk

`func (o *AddStorageVolumeConfigRequest) GetStorageVolumeConfigIdOk() (*string, bool)`

GetStorageVolumeConfigIdOk returns a tuple with the StorageVolumeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeConfigId

`func (o *AddStorageVolumeConfigRequest) SetStorageVolumeConfigId(v string)`

SetStorageVolumeConfigId sets StorageVolumeConfigId field to given value.


### GetToken

`func (o *AddStorageVolumeConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddStorageVolumeConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddStorageVolumeConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


