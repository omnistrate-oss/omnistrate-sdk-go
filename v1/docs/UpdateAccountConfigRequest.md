# UpdateAccountConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The updated description for the account | [optional] 
**Id** | **string** | ID of an Account Config | 
**Name** | Pointer to **string** | The updated name of the account | [optional] 
**NebiusBindings** | Pointer to [**[]NebiusAccountBindingInput**](NebiusAccountBindingInput.md) | Full replacement set of Nebius project/service-account bindings for an existing tenant-scoped Nebius account configuration | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateAccountConfigRequest

`func NewUpdateAccountConfigRequest(id string, token string, ) *UpdateAccountConfigRequest`

NewUpdateAccountConfigRequest instantiates a new UpdateAccountConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigRequestWithDefaults

`func NewUpdateAccountConfigRequestWithDefaults() *UpdateAccountConfigRequest`

NewUpdateAccountConfigRequestWithDefaults instantiates a new UpdateAccountConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateAccountConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAccountConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAccountConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAccountConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateAccountConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAccountConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAccountConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAccountConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccountConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNebiusBindings

`func (o *UpdateAccountConfigRequest) GetNebiusBindings() []NebiusAccountBindingInput`

GetNebiusBindings returns the NebiusBindings field if non-nil, zero value otherwise.

### GetNebiusBindingsOk

`func (o *UpdateAccountConfigRequest) GetNebiusBindingsOk() (*[]NebiusAccountBindingInput, bool)`

GetNebiusBindingsOk returns a tuple with the NebiusBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusBindings

`func (o *UpdateAccountConfigRequest) SetNebiusBindings(v []NebiusAccountBindingInput)`

SetNebiusBindings sets NebiusBindings field to given value.

### HasNebiusBindings

`func (o *UpdateAccountConfigRequest) HasNebiusBindings() bool`

HasNebiusBindings returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAccountConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAccountConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAccountConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


