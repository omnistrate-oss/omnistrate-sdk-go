# UpdateIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | Pointer to **string** | The authorization endpoint of the Identity Provider | [optional] 
**ClientId** | Pointer to **string** | The Client ID of the Identity Provider | [optional] 
**ClientSecret** | Pointer to **string** | The Client Secret of the Identity Provider | [optional] 
**Disabled** | Pointer to **bool** | Whether the Identity Provider is disabled | [optional] 
**EmailIdentifiers** | Pointer to **string** | The email identifiers to use for the Identity Provider | [optional] 
**EnvironmentType** | Pointer to **string** | The type of environment for the Identity Provider | [optional] 
**Id** | **string** | ID of an Identity Provider | 
**LoginButtonIconUrl** | Pointer to **string** | The URL of the icon to use for the login button | [optional] 
**LoginButtonText** | Pointer to **string** | The text to use for the login button | [optional] 
**Name** | Pointer to **string** | The name of the Identity Provider | [optional] 
**Scopes** | Pointer to **string** | The scopes to request from the Identity Provider | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**TokenEndpoint** | Pointer to **string** | The token endpoint of the Identity Provider | [optional] 
**UserInfoEndpoint** | Pointer to **string** | The user info endpoint of the Identity Provider | [optional] 

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

### GetAuthorizationEndpoint

`func (o *UpdateIdentityProviderRequest) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *UpdateIdentityProviderRequest) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *UpdateIdentityProviderRequest) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *UpdateIdentityProviderRequest) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

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

### GetDisabled

`func (o *UpdateIdentityProviderRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateIdentityProviderRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateIdentityProviderRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateIdentityProviderRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmailIdentifiers

`func (o *UpdateIdentityProviderRequest) GetEmailIdentifiers() string`

GetEmailIdentifiers returns the EmailIdentifiers field if non-nil, zero value otherwise.

### GetEmailIdentifiersOk

`func (o *UpdateIdentityProviderRequest) GetEmailIdentifiersOk() (*string, bool)`

GetEmailIdentifiersOk returns a tuple with the EmailIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdentifiers

`func (o *UpdateIdentityProviderRequest) SetEmailIdentifiers(v string)`

SetEmailIdentifiers sets EmailIdentifiers field to given value.

### HasEmailIdentifiers

`func (o *UpdateIdentityProviderRequest) HasEmailIdentifiers() bool`

HasEmailIdentifiers returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *UpdateIdentityProviderRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UpdateIdentityProviderRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UpdateIdentityProviderRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *UpdateIdentityProviderRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

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


### GetLoginButtonIconUrl

`func (o *UpdateIdentityProviderRequest) GetLoginButtonIconUrl() string`

GetLoginButtonIconUrl returns the LoginButtonIconUrl field if non-nil, zero value otherwise.

### GetLoginButtonIconUrlOk

`func (o *UpdateIdentityProviderRequest) GetLoginButtonIconUrlOk() (*string, bool)`

GetLoginButtonIconUrlOk returns a tuple with the LoginButtonIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIconUrl

`func (o *UpdateIdentityProviderRequest) SetLoginButtonIconUrl(v string)`

SetLoginButtonIconUrl sets LoginButtonIconUrl field to given value.

### HasLoginButtonIconUrl

`func (o *UpdateIdentityProviderRequest) HasLoginButtonIconUrl() bool`

HasLoginButtonIconUrl returns a boolean if a field has been set.

### GetLoginButtonText

`func (o *UpdateIdentityProviderRequest) GetLoginButtonText() string`

GetLoginButtonText returns the LoginButtonText field if non-nil, zero value otherwise.

### GetLoginButtonTextOk

`func (o *UpdateIdentityProviderRequest) GetLoginButtonTextOk() (*string, bool)`

GetLoginButtonTextOk returns a tuple with the LoginButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonText

`func (o *UpdateIdentityProviderRequest) SetLoginButtonText(v string)`

SetLoginButtonText sets LoginButtonText field to given value.

### HasLoginButtonText

`func (o *UpdateIdentityProviderRequest) HasLoginButtonText() bool`

HasLoginButtonText returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentityProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentityProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentityProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIdentityProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateIdentityProviderRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateIdentityProviderRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateIdentityProviderRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateIdentityProviderRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

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


### GetTokenEndpoint

`func (o *UpdateIdentityProviderRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *UpdateIdentityProviderRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *UpdateIdentityProviderRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *UpdateIdentityProviderRequest) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *UpdateIdentityProviderRequest) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *UpdateIdentityProviderRequest) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *UpdateIdentityProviderRequest) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *UpdateIdentityProviderRequest) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


