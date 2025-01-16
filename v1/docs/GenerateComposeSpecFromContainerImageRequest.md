# GenerateComposeSpecFromContainerImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | Runtime environment variables needed to run the image | [optional] 
**Image** | **string** | Name of the image along with the tag. Include the repository name if the image is not from the official repository | 
**ImageRegistry** | **string** | Registry where the image is stored | 
**Password** | Pointer to **string** | Password to access the image registry | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Username** | Pointer to **string** | Username to access the image registry | [optional] 

## Methods

### NewGenerateComposeSpecFromContainerImageRequest

`func NewGenerateComposeSpecFromContainerImageRequest(image string, imageRegistry string, token string, ) *GenerateComposeSpecFromContainerImageRequest`

NewGenerateComposeSpecFromContainerImageRequest instantiates a new GenerateComposeSpecFromContainerImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateComposeSpecFromContainerImageRequestWithDefaults

`func NewGenerateComposeSpecFromContainerImageRequestWithDefaults() *GenerateComposeSpecFromContainerImageRequest`

NewGenerateComposeSpecFromContainerImageRequestWithDefaults instantiates a new GenerateComposeSpecFromContainerImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetImage

`func (o *GenerateComposeSpecFromContainerImageRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GenerateComposeSpecFromContainerImageRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequest) GetImageRegistry() string`

GetImageRegistry returns the ImageRegistry field if non-nil, zero value otherwise.

### GetImageRegistryOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetImageRegistryOk() (*string, bool)`

GetImageRegistryOk returns a tuple with the ImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequest) SetImageRegistry(v string)`

SetImageRegistry sets ImageRegistry field to given value.


### GetPassword

`func (o *GenerateComposeSpecFromContainerImageRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GenerateComposeSpecFromContainerImageRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GenerateComposeSpecFromContainerImageRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *GenerateComposeSpecFromContainerImageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GenerateComposeSpecFromContainerImageRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsername

`func (o *GenerateComposeSpecFromContainerImageRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GenerateComposeSpecFromContainerImageRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GenerateComposeSpecFromContainerImageRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GenerateComposeSpecFromContainerImageRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


