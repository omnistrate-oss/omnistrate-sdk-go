# UpdateImageConfigRequest2

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

### NewUpdateImageConfigRequest2

`func NewUpdateImageConfigRequest2() *UpdateImageConfigRequest2`

NewUpdateImageConfigRequest2 instantiates a new UpdateImageConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageConfigRequest2WithDefaults

`func NewUpdateImageConfigRequest2WithDefaults() *UpdateImageConfigRequest2`

NewUpdateImageConfigRequest2WithDefaults instantiates a new UpdateImageConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequest2) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *UpdateImageConfigRequest2) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequest2) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *UpdateImageConfigRequest2) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateImageConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImageConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImageConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImageConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageName

`func (o *UpdateImageConfigRequest2) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *UpdateImageConfigRequest2) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *UpdateImageConfigRequest2) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *UpdateImageConfigRequest2) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegistryId

`func (o *UpdateImageConfigRequest2) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *UpdateImageConfigRequest2) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *UpdateImageConfigRequest2) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.

### HasImageRegistryId

`func (o *UpdateImageConfigRequest2) HasImageRegistryId() bool`

HasImageRegistryId returns a boolean if a field has been set.

### GetImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequest2) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *UpdateImageConfigRequest2) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequest2) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *UpdateImageConfigRequest2) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *UpdateImageConfigRequest2) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *UpdateImageConfigRequest2) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *UpdateImageConfigRequest2) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *UpdateImageConfigRequest2) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


