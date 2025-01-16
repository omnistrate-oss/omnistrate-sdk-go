# CreateIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Client ID of the Identity Provider | 
**ClientSecret** | **string** | The Client Secret of the Identity Provider | 
**IdentityProviderName** | **string** | The name of the identity provider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateIdentityProviderRequest

`func NewCreateIdentityProviderRequest(clientId string, clientSecret string, identityProviderName string, token string, ) *CreateIdentityProviderRequest`

NewCreateIdentityProviderRequest instantiates a new CreateIdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityProviderRequestWithDefaults

`func NewCreateIdentityProviderRequestWithDefaults() *CreateIdentityProviderRequest`

NewCreateIdentityProviderRequestWithDefaults instantiates a new CreateIdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreateIdentityProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateIdentityProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateIdentityProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateIdentityProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateIdentityProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateIdentityProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetIdentityProviderName

`func (o *CreateIdentityProviderRequest) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *CreateIdentityProviderRequest) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *CreateIdentityProviderRequest) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetToken

`func (o *CreateIdentityProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateIdentityProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateIdentityProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


