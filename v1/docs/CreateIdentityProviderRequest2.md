# CreateIdentityProviderRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | Pointer to **string** | The authorization endpoint of the Identity Provider | [optional] 
**ClientId** | **string** | The Client ID of the Identity Provider | 
**ClientSecret** | **string** | The Client Secret of the Identity Provider | 
**Disabled** | Pointer to **bool** | Whether the Identity Provider is disabled | [optional] 
**EmailIdentifiers** | Pointer to **string** | The email identifiers to use for the Identity Provider | [optional] 
**EnvironmentType** | Pointer to **string** | The type of environment to filter costs by | [optional] 
**IdentityProviderName** | **string** | The name of the Identity Provider | 
**LoginButtonIconUrl** | Pointer to **string** | The URL of the icon to use for the login button | [optional] 
**LoginButtonText** | Pointer to **string** | The text to use for the login button | [optional] 
**Name** | Pointer to **string** | The name of the Identity Provider | [optional] 
**Scopes** | Pointer to **[]string** | The scopes to request from the Identity Provider | [optional] 
**TokenEndpoint** | Pointer to **string** | The token endpoint of the Identity Provider | [optional] 
**UserInfoEndpoint** | Pointer to **string** | The user info endpoint of the Identity Provider | [optional] 

## Methods

### NewCreateIdentityProviderRequest2

`func NewCreateIdentityProviderRequest2(clientId string, clientSecret string, identityProviderName string, ) *CreateIdentityProviderRequest2`

NewCreateIdentityProviderRequest2 instantiates a new CreateIdentityProviderRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityProviderRequest2WithDefaults

`func NewCreateIdentityProviderRequest2WithDefaults() *CreateIdentityProviderRequest2`

NewCreateIdentityProviderRequest2WithDefaults instantiates a new CreateIdentityProviderRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationEndpoint

`func (o *CreateIdentityProviderRequest2) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *CreateIdentityProviderRequest2) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *CreateIdentityProviderRequest2) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *CreateIdentityProviderRequest2) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetClientId

`func (o *CreateIdentityProviderRequest2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateIdentityProviderRequest2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateIdentityProviderRequest2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateIdentityProviderRequest2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateIdentityProviderRequest2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateIdentityProviderRequest2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDisabled

`func (o *CreateIdentityProviderRequest2) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateIdentityProviderRequest2) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateIdentityProviderRequest2) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateIdentityProviderRequest2) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmailIdentifiers

`func (o *CreateIdentityProviderRequest2) GetEmailIdentifiers() string`

GetEmailIdentifiers returns the EmailIdentifiers field if non-nil, zero value otherwise.

### GetEmailIdentifiersOk

`func (o *CreateIdentityProviderRequest2) GetEmailIdentifiersOk() (*string, bool)`

GetEmailIdentifiersOk returns a tuple with the EmailIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdentifiers

`func (o *CreateIdentityProviderRequest2) SetEmailIdentifiers(v string)`

SetEmailIdentifiers sets EmailIdentifiers field to given value.

### HasEmailIdentifiers

`func (o *CreateIdentityProviderRequest2) HasEmailIdentifiers() bool`

HasEmailIdentifiers returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *CreateIdentityProviderRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CreateIdentityProviderRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CreateIdentityProviderRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CreateIdentityProviderRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *CreateIdentityProviderRequest2) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *CreateIdentityProviderRequest2) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *CreateIdentityProviderRequest2) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetLoginButtonIconUrl

`func (o *CreateIdentityProviderRequest2) GetLoginButtonIconUrl() string`

GetLoginButtonIconUrl returns the LoginButtonIconUrl field if non-nil, zero value otherwise.

### GetLoginButtonIconUrlOk

`func (o *CreateIdentityProviderRequest2) GetLoginButtonIconUrlOk() (*string, bool)`

GetLoginButtonIconUrlOk returns a tuple with the LoginButtonIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIconUrl

`func (o *CreateIdentityProviderRequest2) SetLoginButtonIconUrl(v string)`

SetLoginButtonIconUrl sets LoginButtonIconUrl field to given value.

### HasLoginButtonIconUrl

`func (o *CreateIdentityProviderRequest2) HasLoginButtonIconUrl() bool`

HasLoginButtonIconUrl returns a boolean if a field has been set.

### GetLoginButtonText

`func (o *CreateIdentityProviderRequest2) GetLoginButtonText() string`

GetLoginButtonText returns the LoginButtonText field if non-nil, zero value otherwise.

### GetLoginButtonTextOk

`func (o *CreateIdentityProviderRequest2) GetLoginButtonTextOk() (*string, bool)`

GetLoginButtonTextOk returns a tuple with the LoginButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonText

`func (o *CreateIdentityProviderRequest2) SetLoginButtonText(v string)`

SetLoginButtonText sets LoginButtonText field to given value.

### HasLoginButtonText

`func (o *CreateIdentityProviderRequest2) HasLoginButtonText() bool`

HasLoginButtonText returns a boolean if a field has been set.

### GetName

`func (o *CreateIdentityProviderRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIdentityProviderRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIdentityProviderRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIdentityProviderRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *CreateIdentityProviderRequest2) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateIdentityProviderRequest2) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateIdentityProviderRequest2) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateIdentityProviderRequest2) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *CreateIdentityProviderRequest2) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *CreateIdentityProviderRequest2) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *CreateIdentityProviderRequest2) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *CreateIdentityProviderRequest2) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *CreateIdentityProviderRequest2) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *CreateIdentityProviderRequest2) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *CreateIdentityProviderRequest2) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *CreateIdentityProviderRequest2) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


