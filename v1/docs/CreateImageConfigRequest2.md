# CreateImageConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomImageCommandsAndArgs** | Pointer to [**CustomImageCommandsAndArgs**](CustomImageCommandsAndArgs.md) |  | [optional] 
**Description** | **string** | A brief description of the image configuration | 
**ImageName** | **string** | Name of the container image | 
**ImageRegistryId** | Pointer to **string** | The image registry ID to use for the infra | [optional] 
**ImageSignaturePublicKeyPEM** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the container image | [optional] 
**ImageTag** | Pointer to **string** | Tag representing the specific software image version | [optional] [default to "latest"]
**ServiceEnvironmentId** | **string** | The service environment ID | 

## Methods

### NewCreateImageConfigRequest2

`func NewCreateImageConfigRequest2(description string, imageName string, serviceEnvironmentId string, ) *CreateImageConfigRequest2`

NewCreateImageConfigRequest2 instantiates a new CreateImageConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageConfigRequest2WithDefaults

`func NewCreateImageConfigRequest2WithDefaults() *CreateImageConfigRequest2`

NewCreateImageConfigRequest2WithDefaults instantiates a new CreateImageConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest2) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *CreateImageConfigRequest2) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest2) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest2) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *CreateImageConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageName

`func (o *CreateImageConfigRequest2) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateImageConfigRequest2) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateImageConfigRequest2) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageRegistryId

`func (o *CreateImageConfigRequest2) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *CreateImageConfigRequest2) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *CreateImageConfigRequest2) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.

### HasImageRegistryId

`func (o *CreateImageConfigRequest2) HasImageRegistryId() bool`

HasImageRegistryId returns a boolean if a field has been set.

### GetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest2) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *CreateImageConfigRequest2) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest2) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest2) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *CreateImageConfigRequest2) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *CreateImageConfigRequest2) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *CreateImageConfigRequest2) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *CreateImageConfigRequest2) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateImageConfigRequest2) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateImageConfigRequest2) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateImageConfigRequest2) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


