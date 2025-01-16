# UpdateImageRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the Image Registry | [optional] 
**Host** | Pointer to **string** | The Image Registry host | [optional] 
**Id** | **string** | ID of an Image Registry | 
**Name** | Pointer to **string** | Name of the Image Registry | [optional] 
**Password** | Pointer to **string** | The password to use when authenticating to the Image Registry | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Username** | Pointer to **string** | The username to use when authenticating to the Image Registry | [optional] 

## Methods

### NewUpdateImageRegistryRequest

`func NewUpdateImageRegistryRequest(id string, token string, ) *UpdateImageRegistryRequest`

NewUpdateImageRegistryRequest instantiates a new UpdateImageRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageRegistryRequestWithDefaults

`func NewUpdateImageRegistryRequestWithDefaults() *UpdateImageRegistryRequest`

NewUpdateImageRegistryRequestWithDefaults instantiates a new UpdateImageRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateImageRegistryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImageRegistryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImageRegistryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImageRegistryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *UpdateImageRegistryRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateImageRegistryRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateImageRegistryRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateImageRegistryRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *UpdateImageRegistryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateImageRegistryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateImageRegistryRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateImageRegistryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateImageRegistryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateImageRegistryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateImageRegistryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateImageRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateImageRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateImageRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateImageRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *UpdateImageRegistryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateImageRegistryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateImageRegistryRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsername

`func (o *UpdateImageRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateImageRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateImageRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateImageRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


