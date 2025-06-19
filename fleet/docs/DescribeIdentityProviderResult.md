# DescribeIdentityProviderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | Pointer to **string** | The authorization endpoint of the Identity Provider | [optional] 
**ClientId** | **string** | The Client ID of the Identity Provider | 
**ClientSecret** | Pointer to **string** | The Client Secret of the Identity Provider | [optional] 
**Disabled** | Pointer to **bool** | Whether the Identity Provider is disabled | [optional] 
**EmailIdentifiers** | Pointer to **string** | The email identifiers to use for the Identity Provider | [optional] 
**EnvironmentType** | Pointer to **string** | The type of environment for the Identity Provider | [optional] 
**Id** | **string** | ID of an Identity Provider | 
**IdentityProviderName** | **string** | The name of the identity provider | 
**LoginButtonIconUrl** | Pointer to **string** | The URL of the icon to use for the login button | [optional] 
**LoginButtonText** | Pointer to **string** | The text to use for the login button | [optional] 
**Name** | Pointer to **string** | The name of the Identity Provider | [optional] 
**Scopes** | Pointer to **string** | The scopes to request from the Identity Provider | [optional] 
**Status** | **string** | The status of an operation | 
**TokenEndpoint** | Pointer to **string** | The token endpoint of the Identity Provider | [optional] 
**UserInfoEndpoint** | Pointer to **string** | The user info endpoint of the Identity Provider | [optional] 

## Methods

### NewDescribeIdentityProviderResult

`func NewDescribeIdentityProviderResult(clientId string, id string, identityProviderName string, status string, ) *DescribeIdentityProviderResult`

NewDescribeIdentityProviderResult instantiates a new DescribeIdentityProviderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIdentityProviderResultWithDefaults

`func NewDescribeIdentityProviderResultWithDefaults() *DescribeIdentityProviderResult`

NewDescribeIdentityProviderResultWithDefaults instantiates a new DescribeIdentityProviderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationEndpoint

`func (o *DescribeIdentityProviderResult) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *DescribeIdentityProviderResult) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *DescribeIdentityProviderResult) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *DescribeIdentityProviderResult) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetClientId

`func (o *DescribeIdentityProviderResult) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DescribeIdentityProviderResult) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DescribeIdentityProviderResult) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *DescribeIdentityProviderResult) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DescribeIdentityProviderResult) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DescribeIdentityProviderResult) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *DescribeIdentityProviderResult) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDisabled

`func (o *DescribeIdentityProviderResult) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *DescribeIdentityProviderResult) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *DescribeIdentityProviderResult) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *DescribeIdentityProviderResult) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmailIdentifiers

`func (o *DescribeIdentityProviderResult) GetEmailIdentifiers() string`

GetEmailIdentifiers returns the EmailIdentifiers field if non-nil, zero value otherwise.

### GetEmailIdentifiersOk

`func (o *DescribeIdentityProviderResult) GetEmailIdentifiersOk() (*string, bool)`

GetEmailIdentifiersOk returns a tuple with the EmailIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdentifiers

`func (o *DescribeIdentityProviderResult) SetEmailIdentifiers(v string)`

SetEmailIdentifiers sets EmailIdentifiers field to given value.

### HasEmailIdentifiers

`func (o *DescribeIdentityProviderResult) HasEmailIdentifiers() bool`

HasEmailIdentifiers returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *DescribeIdentityProviderResult) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeIdentityProviderResult) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeIdentityProviderResult) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *DescribeIdentityProviderResult) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetId

`func (o *DescribeIdentityProviderResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeIdentityProviderResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeIdentityProviderResult) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityProviderName

`func (o *DescribeIdentityProviderResult) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *DescribeIdentityProviderResult) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *DescribeIdentityProviderResult) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetLoginButtonIconUrl

`func (o *DescribeIdentityProviderResult) GetLoginButtonIconUrl() string`

GetLoginButtonIconUrl returns the LoginButtonIconUrl field if non-nil, zero value otherwise.

### GetLoginButtonIconUrlOk

`func (o *DescribeIdentityProviderResult) GetLoginButtonIconUrlOk() (*string, bool)`

GetLoginButtonIconUrlOk returns a tuple with the LoginButtonIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIconUrl

`func (o *DescribeIdentityProviderResult) SetLoginButtonIconUrl(v string)`

SetLoginButtonIconUrl sets LoginButtonIconUrl field to given value.

### HasLoginButtonIconUrl

`func (o *DescribeIdentityProviderResult) HasLoginButtonIconUrl() bool`

HasLoginButtonIconUrl returns a boolean if a field has been set.

### GetLoginButtonText

`func (o *DescribeIdentityProviderResult) GetLoginButtonText() string`

GetLoginButtonText returns the LoginButtonText field if non-nil, zero value otherwise.

### GetLoginButtonTextOk

`func (o *DescribeIdentityProviderResult) GetLoginButtonTextOk() (*string, bool)`

GetLoginButtonTextOk returns a tuple with the LoginButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonText

`func (o *DescribeIdentityProviderResult) SetLoginButtonText(v string)`

SetLoginButtonText sets LoginButtonText field to given value.

### HasLoginButtonText

`func (o *DescribeIdentityProviderResult) HasLoginButtonText() bool`

HasLoginButtonText returns a boolean if a field has been set.

### GetName

`func (o *DescribeIdentityProviderResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeIdentityProviderResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeIdentityProviderResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeIdentityProviderResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *DescribeIdentityProviderResult) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DescribeIdentityProviderResult) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DescribeIdentityProviderResult) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DescribeIdentityProviderResult) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeIdentityProviderResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeIdentityProviderResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeIdentityProviderResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenEndpoint

`func (o *DescribeIdentityProviderResult) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *DescribeIdentityProviderResult) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *DescribeIdentityProviderResult) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *DescribeIdentityProviderResult) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *DescribeIdentityProviderResult) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *DescribeIdentityProviderResult) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *DescribeIdentityProviderResult) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *DescribeIdentityProviderResult) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


