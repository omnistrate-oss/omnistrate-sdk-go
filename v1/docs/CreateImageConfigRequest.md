# CreateImageConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomImageCommandsAndArgs** | Pointer to [**CustomImageCommandsAndArgs**](CustomImageCommandsAndArgs.md) |  | [optional] 
**Description** | **string** | A brief description of the image configuration | 
**ImageName** | **string** | Name of the container image | 
**ImageRegistryId** | Pointer to **string** | ID of an Image Registry | [optional] 
**ImageSignaturePublicKeyPEM** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the container image | [optional] 
**ImageTag** | Pointer to **string** | Tag representing the specific software image version | [optional] [default to "latest"]
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateImageConfigRequest

`func NewCreateImageConfigRequest(description string, imageName string, serviceEnvironmentId string, serviceId string, token string, ) *CreateImageConfigRequest`

NewCreateImageConfigRequest instantiates a new CreateImageConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageConfigRequestWithDefaults

`func NewCreateImageConfigRequestWithDefaults() *CreateImageConfigRequest`

NewCreateImageConfigRequestWithDefaults instantiates a new CreateImageConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *CreateImageConfigRequest) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *CreateImageConfigRequest) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *CreateImageConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageName

`func (o *CreateImageConfigRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateImageConfigRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateImageConfigRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageRegistryId

`func (o *CreateImageConfigRequest) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *CreateImageConfigRequest) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *CreateImageConfigRequest) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.

### HasImageRegistryId

`func (o *CreateImageConfigRequest) HasImageRegistryId() bool`

HasImageRegistryId returns a boolean if a field has been set.

### GetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *CreateImageConfigRequest) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *CreateImageConfigRequest) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *CreateImageConfigRequest) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *CreateImageConfigRequest) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *CreateImageConfigRequest) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *CreateImageConfigRequest) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateImageConfigRequest) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateImageConfigRequest) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateImageConfigRequest) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *CreateImageConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateImageConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateImageConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreateImageConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateImageConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateImageConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


