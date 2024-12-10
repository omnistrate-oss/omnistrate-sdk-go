# CheckIfContainerImageAccessibleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Name of the image along with the tag. Include the repository name if the image is not from the official repository | 
**ImageRegistry** | **string** | Registry where the image is stored | 
**Password** | Pointer to **string** | Password to access the image registry | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Username** | Pointer to **string** | Username to access the image registry | [optional] 

## Methods

### NewCheckIfContainerImageAccessibleRequest

`func NewCheckIfContainerImageAccessibleRequest(image string, imageRegistry string, token string, ) *CheckIfContainerImageAccessibleRequest`

NewCheckIfContainerImageAccessibleRequest instantiates a new CheckIfContainerImageAccessibleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckIfContainerImageAccessibleRequestWithDefaults

`func NewCheckIfContainerImageAccessibleRequestWithDefaults() *CheckIfContainerImageAccessibleRequest`

NewCheckIfContainerImageAccessibleRequestWithDefaults instantiates a new CheckIfContainerImageAccessibleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *CheckIfContainerImageAccessibleRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CheckIfContainerImageAccessibleRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CheckIfContainerImageAccessibleRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageRegistry

`func (o *CheckIfContainerImageAccessibleRequest) GetImageRegistry() string`

GetImageRegistry returns the ImageRegistry field if non-nil, zero value otherwise.

### GetImageRegistryOk

`func (o *CheckIfContainerImageAccessibleRequest) GetImageRegistryOk() (*string, bool)`

GetImageRegistryOk returns a tuple with the ImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistry

`func (o *CheckIfContainerImageAccessibleRequest) SetImageRegistry(v string)`

SetImageRegistry sets ImageRegistry field to given value.


### GetPassword

`func (o *CheckIfContainerImageAccessibleRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckIfContainerImageAccessibleRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CheckIfContainerImageAccessibleRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CheckIfContainerImageAccessibleRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CheckIfContainerImageAccessibleRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckIfContainerImageAccessibleRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckIfContainerImageAccessibleRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsername

`func (o *CheckIfContainerImageAccessibleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CheckIfContainerImageAccessibleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CheckIfContainerImageAccessibleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CheckIfContainerImageAccessibleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


