# VerifyIdentityProviderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Client ID of the Identity Provider | 
**ClientSecret** | Pointer to **string** | The Client Secret of the Identity Provider | [optional] 
**Id** | **string** | ID of an Identity Provider | 
**IdentityProviderName** | **string** | The name of the identity provider | 
**Name** | Pointer to **string** | The name of the Identity Provider | [optional] 
**Status** | **string** | The status of an operation | 

## Methods

### NewVerifyIdentityProviderResult

`func NewVerifyIdentityProviderResult(clientId string, id string, identityProviderName string, status string, ) *VerifyIdentityProviderResult`

NewVerifyIdentityProviderResult instantiates a new VerifyIdentityProviderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyIdentityProviderResultWithDefaults

`func NewVerifyIdentityProviderResultWithDefaults() *VerifyIdentityProviderResult`

NewVerifyIdentityProviderResultWithDefaults instantiates a new VerifyIdentityProviderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *VerifyIdentityProviderResult) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VerifyIdentityProviderResult) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VerifyIdentityProviderResult) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *VerifyIdentityProviderResult) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VerifyIdentityProviderResult) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VerifyIdentityProviderResult) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *VerifyIdentityProviderResult) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetId

`func (o *VerifyIdentityProviderResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyIdentityProviderResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyIdentityProviderResult) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityProviderName

`func (o *VerifyIdentityProviderResult) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *VerifyIdentityProviderResult) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *VerifyIdentityProviderResult) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetName

`func (o *VerifyIdentityProviderResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerifyIdentityProviderResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerifyIdentityProviderResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VerifyIdentityProviderResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VerifyIdentityProviderResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyIdentityProviderResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyIdentityProviderResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


