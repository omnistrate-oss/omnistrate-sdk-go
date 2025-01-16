# CreateImageRegistryRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the Image Registry | 
**Host** | **string** | The Image Registry host | 
**Name** | **string** | Name of the Image Registry | 
**Password** | Pointer to **string** | The password to use when authenticating to the Image Registry | [optional] 
**Username** | Pointer to **string** | The username to use when authenticating to the Image Registry | [optional] 

## Methods

### NewCreateImageRegistryRequest2

`func NewCreateImageRegistryRequest2(description string, host string, name string, ) *CreateImageRegistryRequest2`

NewCreateImageRegistryRequest2 instantiates a new CreateImageRegistryRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageRegistryRequest2WithDefaults

`func NewCreateImageRegistryRequest2WithDefaults() *CreateImageRegistryRequest2`

NewCreateImageRegistryRequest2WithDefaults instantiates a new CreateImageRegistryRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateImageRegistryRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageRegistryRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageRegistryRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHost

`func (o *CreateImageRegistryRequest2) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateImageRegistryRequest2) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateImageRegistryRequest2) SetHost(v string)`

SetHost sets Host field to given value.


### GetName

`func (o *CreateImageRegistryRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateImageRegistryRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateImageRegistryRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateImageRegistryRequest2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateImageRegistryRequest2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateImageRegistryRequest2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateImageRegistryRequest2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *CreateImageRegistryRequest2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateImageRegistryRequest2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateImageRegistryRequest2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateImageRegistryRequest2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


