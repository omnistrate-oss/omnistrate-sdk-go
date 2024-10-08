# CreateImageConfigRequestBody

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

### NewCreateImageConfigRequestBody

`func NewCreateImageConfigRequestBody(description string, imageName string, serviceEnvironmentId string, ) *CreateImageConfigRequestBody`

NewCreateImageConfigRequestBody instantiates a new CreateImageConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageConfigRequestBodyWithDefaults

`func NewCreateImageConfigRequestBodyWithDefaults() *CreateImageConfigRequestBody`

NewCreateImageConfigRequestBodyWithDefaults instantiates a new CreateImageConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequestBody) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *CreateImageConfigRequestBody) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequestBody) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequestBody) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *CreateImageConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageName

`func (o *CreateImageConfigRequestBody) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateImageConfigRequestBody) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateImageConfigRequestBody) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageRegistryId

`func (o *CreateImageConfigRequestBody) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *CreateImageConfigRequestBody) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *CreateImageConfigRequestBody) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.

### HasImageRegistryId

`func (o *CreateImageConfigRequestBody) HasImageRegistryId() bool`

HasImageRegistryId returns a boolean if a field has been set.

### GetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequestBody) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *CreateImageConfigRequestBody) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequestBody) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequestBody) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *CreateImageConfigRequestBody) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *CreateImageConfigRequestBody) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *CreateImageConfigRequestBody) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *CreateImageConfigRequestBody) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateImageConfigRequestBody) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateImageConfigRequestBody) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateImageConfigRequestBody) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


