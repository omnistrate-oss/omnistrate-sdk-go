# DescribeImageConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomImageCommandsAndArgs** | Pointer to [**CustomImageCommandsAndArgs**](CustomImageCommandsAndArgs.md) |  | [optional] 
**Description** | **string** | A brief description of the image configuration | 
**Id** | **string** | The image configuration ID | 
**ImageName** | **string** | Name of the container image | 
**ImageRegistryId** | **string** | The image registry ID to use for the infra | 
**ImageSignaturePublicKeyPEM** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the container image | [optional] 
**ImageTag** | **string** | Tag representing the software image version that is currently preferred | [default to "latest"]
**ServiceEnvironmentId** | **string** | The service environment ID | 
**ServiceId** | **string** | The service ID | 

## Methods

### NewDescribeImageConfigResult

`func NewDescribeImageConfigResult(description string, id string, imageName string, imageRegistryId string, imageTag string, serviceEnvironmentId string, serviceId string, ) *DescribeImageConfigResult`

NewDescribeImageConfigResult instantiates a new DescribeImageConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeImageConfigResultWithDefaults

`func NewDescribeImageConfigResultWithDefaults() *DescribeImageConfigResult`

NewDescribeImageConfigResultWithDefaults instantiates a new DescribeImageConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageCommandsAndArgs

`func (o *DescribeImageConfigResult) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs`

GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field if non-nil, zero value otherwise.

### GetCustomImageCommandsAndArgsOk

`func (o *DescribeImageConfigResult) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool)`

GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageCommandsAndArgs

`func (o *DescribeImageConfigResult) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs)`

SetCustomImageCommandsAndArgs sets CustomImageCommandsAndArgs field to given value.

### HasCustomImageCommandsAndArgs

`func (o *DescribeImageConfigResult) HasCustomImageCommandsAndArgs() bool`

HasCustomImageCommandsAndArgs returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeImageConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeImageConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeImageConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeImageConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeImageConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeImageConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetImageName

`func (o *DescribeImageConfigResult) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DescribeImageConfigResult) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DescribeImageConfigResult) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageRegistryId

`func (o *DescribeImageConfigResult) GetImageRegistryId() string`

GetImageRegistryId returns the ImageRegistryId field if non-nil, zero value otherwise.

### GetImageRegistryIdOk

`func (o *DescribeImageConfigResult) GetImageRegistryIdOk() (*string, bool)`

GetImageRegistryIdOk returns a tuple with the ImageRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistryId

`func (o *DescribeImageConfigResult) SetImageRegistryId(v string)`

SetImageRegistryId sets ImageRegistryId field to given value.


### GetImageSignaturePublicKeyPEM

`func (o *DescribeImageConfigResult) GetImageSignaturePublicKeyPEM() string`

GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field if non-nil, zero value otherwise.

### GetImageSignaturePublicKeyPEMOk

`func (o *DescribeImageConfigResult) GetImageSignaturePublicKeyPEMOk() (*string, bool)`

GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSignaturePublicKeyPEM

`func (o *DescribeImageConfigResult) SetImageSignaturePublicKeyPEM(v string)`

SetImageSignaturePublicKeyPEM sets ImageSignaturePublicKeyPEM field to given value.

### HasImageSignaturePublicKeyPEM

`func (o *DescribeImageConfigResult) HasImageSignaturePublicKeyPEM() bool`

HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.

### GetImageTag

`func (o *DescribeImageConfigResult) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *DescribeImageConfigResult) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *DescribeImageConfigResult) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetServiceEnvironmentId

`func (o *DescribeImageConfigResult) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *DescribeImageConfigResult) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *DescribeImageConfigResult) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *DescribeImageConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeImageConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeImageConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


