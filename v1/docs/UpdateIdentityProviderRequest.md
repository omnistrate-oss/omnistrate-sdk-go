# UpdateIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The Client ID of the Identity Provider | [optional] 
**ClientSecret** | Pointer to **string** | The Client Secret of the Identity Provider | [optional] 
**Id** | **string** | ID of an Identity Provider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateIdentityProviderRequest

`func NewUpdateIdentityProviderRequest(id string, token string, ) *UpdateIdentityProviderRequest`

NewUpdateIdentityProviderRequest instantiates a new UpdateIdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentityProviderRequestWithDefaults

`func NewUpdateIdentityProviderRequestWithDefaults() *UpdateIdentityProviderRequest`

NewUpdateIdentityProviderRequestWithDefaults instantiates a new UpdateIdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *UpdateIdentityProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateIdentityProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateIdentityProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateIdentityProviderRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *UpdateIdentityProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateIdentityProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateIdentityProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateIdentityProviderRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetId

`func (o *UpdateIdentityProviderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIdentityProviderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIdentityProviderRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *UpdateIdentityProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateIdentityProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateIdentityProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


