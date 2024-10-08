# UpdateFileMetadataRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the file | [optional] 
**MountPath** | Pointer to **string** | The mount path of the file | [optional] 
**Name** | Pointer to **string** | The name of the file | [optional] 

## Methods

### NewUpdateFileMetadataRequestBody

`func NewUpdateFileMetadataRequestBody() *UpdateFileMetadataRequestBody`

NewUpdateFileMetadataRequestBody instantiates a new UpdateFileMetadataRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileMetadataRequestBodyWithDefaults

`func NewUpdateFileMetadataRequestBodyWithDefaults() *UpdateFileMetadataRequestBody`

NewUpdateFileMetadataRequestBodyWithDefaults instantiates a new UpdateFileMetadataRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateFileMetadataRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFileMetadataRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFileMetadataRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFileMetadataRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMountPath

`func (o *UpdateFileMetadataRequestBody) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *UpdateFileMetadataRequestBody) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *UpdateFileMetadataRequestBody) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *UpdateFileMetadataRequestBody) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetName

`func (o *UpdateFileMetadataRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFileMetadataRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFileMetadataRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFileMetadataRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


