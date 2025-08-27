# ContainerImageRegistryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password to authenticate with the container image registry | 
**Username** | **string** | The username to authenticate with the container image registry | 

## Methods

### NewContainerImageRegistryCredentials

`func NewContainerImageRegistryCredentials(password string, username string, ) *ContainerImageRegistryCredentials`

NewContainerImageRegistryCredentials instantiates a new ContainerImageRegistryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageRegistryCredentialsWithDefaults

`func NewContainerImageRegistryCredentialsWithDefaults() *ContainerImageRegistryCredentials`

NewContainerImageRegistryCredentialsWithDefaults instantiates a new ContainerImageRegistryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ContainerImageRegistryCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContainerImageRegistryCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContainerImageRegistryCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ContainerImageRegistryCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ContainerImageRegistryCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ContainerImageRegistryCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


