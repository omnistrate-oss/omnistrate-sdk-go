# UpdateFileMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the file | [optional] 
**FileId** | **string** | ID of a File | 
**Id** | **string** | ID of a resource | 
**MountPath** | Pointer to **string** | The mount path of the file | [optional] 
**Name** | Pointer to **string** | The name of the file | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateFileMetadataRequest

`func NewUpdateFileMetadataRequest(fileId string, id string, serviceId string, token string, ) *UpdateFileMetadataRequest`

NewUpdateFileMetadataRequest instantiates a new UpdateFileMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileMetadataRequestWithDefaults

`func NewUpdateFileMetadataRequestWithDefaults() *UpdateFileMetadataRequest`

NewUpdateFileMetadataRequestWithDefaults instantiates a new UpdateFileMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateFileMetadataRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFileMetadataRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFileMetadataRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFileMetadataRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileId

`func (o *UpdateFileMetadataRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *UpdateFileMetadataRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *UpdateFileMetadataRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetId

`func (o *UpdateFileMetadataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateFileMetadataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateFileMetadataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMountPath

`func (o *UpdateFileMetadataRequest) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *UpdateFileMetadataRequest) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *UpdateFileMetadataRequest) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *UpdateFileMetadataRequest) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetName

`func (o *UpdateFileMetadataRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFileMetadataRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFileMetadataRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFileMetadataRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateFileMetadataRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateFileMetadataRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateFileMetadataRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *UpdateFileMetadataRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateFileMetadataRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateFileMetadataRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


