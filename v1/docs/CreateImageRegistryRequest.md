# CreateImageRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the Image Registry | 
**Host** | **string** | The Image Registry host | 
**Name** | **string** | Name of the Image Registry | 
**Password** | Pointer to **string** | The password to use when authenticating to the Image Registry | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Username** | Pointer to **string** | The username to use when authenticating to the Image Registry | [optional] 

## Methods

### NewCreateImageRegistryRequest

`func NewCreateImageRegistryRequest(description string, host string, name string, token string, ) *CreateImageRegistryRequest`

NewCreateImageRegistryRequest instantiates a new CreateImageRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageRegistryRequestWithDefaults

`func NewCreateImageRegistryRequestWithDefaults() *CreateImageRegistryRequest`

NewCreateImageRegistryRequestWithDefaults instantiates a new CreateImageRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateImageRegistryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageRegistryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageRegistryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHost

`func (o *CreateImageRegistryRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateImageRegistryRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateImageRegistryRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetName

`func (o *CreateImageRegistryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateImageRegistryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateImageRegistryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateImageRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateImageRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateImageRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateImageRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateImageRegistryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateImageRegistryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateImageRegistryRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsername

`func (o *CreateImageRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateImageRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateImageRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateImageRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


