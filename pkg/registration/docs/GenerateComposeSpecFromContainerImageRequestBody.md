# GenerateComposeSpecFromContainerImageRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | Runtime environment variables needed to run the image | [optional] 
**Image** | **string** | Name of the image along with the tag. Include the repository name if the image is not from the official repository | 
**ImageRegistry** | **string** | Registry where the image is stored | 
**Password** | Pointer to **string** | Password to access the image registry | [optional] 
**Username** | Pointer to **string** | Username to access the image registry | [optional] 

## Methods

### NewGenerateComposeSpecFromContainerImageRequestBody

`func NewGenerateComposeSpecFromContainerImageRequestBody(image string, imageRegistry string, ) *GenerateComposeSpecFromContainerImageRequestBody`

NewGenerateComposeSpecFromContainerImageRequestBody instantiates a new GenerateComposeSpecFromContainerImageRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateComposeSpecFromContainerImageRequestBodyWithDefaults

`func NewGenerateComposeSpecFromContainerImageRequestBodyWithDefaults() *GenerateComposeSpecFromContainerImageRequestBody`

NewGenerateComposeSpecFromContainerImageRequestBodyWithDefaults instantiates a new GenerateComposeSpecFromContainerImageRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequestBody) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequestBody) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetImage

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GenerateComposeSpecFromContainerImageRequestBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageRegistry() string`

GetImageRegistry returns the ImageRegistry field if non-nil, zero value otherwise.

### GetImageRegistryOk

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageRegistryOk() (*string, bool)`

GetImageRegistryOk returns a tuple with the ImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequestBody) SetImageRegistry(v string)`

SetImageRegistry sets ImageRegistry field to given value.


### GetPassword

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GenerateComposeSpecFromContainerImageRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GenerateComposeSpecFromContainerImageRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GenerateComposeSpecFromContainerImageRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GenerateComposeSpecFromContainerImageRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GenerateComposeSpecFromContainerImageRequestBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


