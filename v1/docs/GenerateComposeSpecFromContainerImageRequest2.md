# GenerateComposeSpecFromContainerImageRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | Runtime environment variables needed to run the image | [optional] 
**Image** | **string** | Name of the image along with the tag. Include the repository name if the image is not from the official repository | 
**ImageRegistry** | **string** | Registry where the image is stored | 
**Password** | Pointer to **string** | Password to access the image registry | [optional] 
**Username** | Pointer to **string** | Username to access the image registry | [optional] 

## Methods

### NewGenerateComposeSpecFromContainerImageRequest2

`func NewGenerateComposeSpecFromContainerImageRequest2(image string, imageRegistry string, ) *GenerateComposeSpecFromContainerImageRequest2`

NewGenerateComposeSpecFromContainerImageRequest2 instantiates a new GenerateComposeSpecFromContainerImageRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateComposeSpecFromContainerImageRequest2WithDefaults

`func NewGenerateComposeSpecFromContainerImageRequest2WithDefaults() *GenerateComposeSpecFromContainerImageRequest2`

NewGenerateComposeSpecFromContainerImageRequest2WithDefaults instantiates a new GenerateComposeSpecFromContainerImageRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest2) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GenerateComposeSpecFromContainerImageRequest2) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetImage

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GenerateComposeSpecFromContainerImageRequest2) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetImageRegistry() string`

GetImageRegistry returns the ImageRegistry field if non-nil, zero value otherwise.

### GetImageRegistryOk

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetImageRegistryOk() (*string, bool)`

GetImageRegistryOk returns a tuple with the ImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistry

`func (o *GenerateComposeSpecFromContainerImageRequest2) SetImageRegistry(v string)`

SetImageRegistry sets ImageRegistry field to given value.


### GetPassword

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GenerateComposeSpecFromContainerImageRequest2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GenerateComposeSpecFromContainerImageRequest2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GenerateComposeSpecFromContainerImageRequest2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GenerateComposeSpecFromContainerImageRequest2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GenerateComposeSpecFromContainerImageRequest2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


