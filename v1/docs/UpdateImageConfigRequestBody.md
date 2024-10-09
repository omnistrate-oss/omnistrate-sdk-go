# UpdateImageConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomImageCommandsAndArgs** | Pointer to [**CustomImageCommandsAndArgs**](CustomImageCommandsAndArgs.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the image configuration | [optional] 
**ImageName** | Pointer to **string** | Name of the container image | [optional] 
**ImageRegistryId** | Pointer to **string** | The image registry ID to use for the infra | [optional] 
**ImageSignaturePublicKeyPEM** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the container image | [optional] 
**ImageTag** | Pointer to **string** | Tag representing the software image version that is currently preferred | [optional] 

## Methods

### NewUpdateImageConfigRequestBody

`func NewUpdateImageConfigRequestBody() *UpdateImageConfigRequestBody`

NewUpdateImageConfigRequestBody instantiates a new UpdateImageConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageConfigRequestBodyWithDefaults

`func NewUpdateImageConfigRequestBodyWithDefaults() *UpdateImageConfigRequestBody`

NewUpdateImageConfigRequestBodyWithDefaults instantiates a new UpdateImageConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequestBody) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *UpdateImageConfigRequestBody) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequestBody) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequestBody) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateImageConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImageConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImageConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImageConfigRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageName

`func (o *UpdateImageConfigRequestBody) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *UpdateImageConfigRequestBody) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *UpdateImageConfigRequestBody) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *UpdateImageConfigRequestBody) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegistryId

`func (o *UpdateImageConfigRequestBody) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *UpdateImageConfigRequestBody) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *UpdateImageConfigRequestBody) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.

### HasImageRegistryId

`func (o *UpdateImageConfigRequestBody) HasImageRegistryId() bool`

HasImageRegistryId returns a boolean if a field has been set.

### GetImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequestBody) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *UpdateImageConfigRequestBody) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequestBody) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequestBody) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *UpdateImageConfigRequestBody) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *UpdateImageConfigRequestBody) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *UpdateImageConfigRequestBody) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *UpdateImageConfigRequestBody) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

