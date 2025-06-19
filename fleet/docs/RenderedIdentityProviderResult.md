# RenderedIdentityProviderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailIdentifiers** | **string** | The email identifiers to use to filter the Identity Provider available for login | 
**Id** | **string** | ID of an Identity Provider | 
**IdentityProviderName** | **string** | The name of the identity provider | 
**LoginButtonIconUrl** | **string** | The URL of the icon to use for the login button | 
**LoginButtonText** | **string** | The text to use for the login button | 
**Name** | **string** | The name of the Identity Provider | 
**RenderedAuthorizationEndpoint** | **string** | The rendered authorization endpoint of the Identity Provider | 
**State** | Pointer to **string** | The state parameter used to prevent CSRF attacks | [optional] 

## Methods

### NewRenderedIdentityProviderResult

`func NewRenderedIdentityProviderResult(emailIdentifiers string, id string, identityProviderName string, loginButtonIconUrl string, loginButtonText string, name string, renderedAuthorizationEndpoint string, ) *RenderedIdentityProviderResult`

NewRenderedIdentityProviderResult instantiates a new RenderedIdentityProviderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedIdentityProviderResultWithDefaults

`func NewRenderedIdentityProviderResultWithDefaults() *RenderedIdentityProviderResult`

NewRenderedIdentityProviderResultWithDefaults instantiates a new RenderedIdentityProviderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailIdentifiers

`func (o *RenderedIdentityProviderResult) GetEmailIdentifiers() string`

GetEmailIdentifiers returns the EmailIdentifiers field if non-nil, zero value otherwise.

### GetEmailIdentifiersOk

`func (o *RenderedIdentityProviderResult) GetEmailIdentifiersOk() (*string, bool)`

GetEmailIdentifiersOk returns a tuple with the EmailIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdentifiers

`func (o *RenderedIdentityProviderResult) SetEmailIdentifiers(v string)`

SetEmailIdentifiers sets EmailIdentifiers field to given value.


### GetId

`func (o *RenderedIdentityProviderResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenderedIdentityProviderResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenderedIdentityProviderResult) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityProviderName

`func (o *RenderedIdentityProviderResult) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *RenderedIdentityProviderResult) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *RenderedIdentityProviderResult) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetLoginButtonIconUrl

`func (o *RenderedIdentityProviderResult) GetLoginButtonIconUrl() string`

GetLoginButtonIconUrl returns the LoginButtonIconUrl field if non-nil, zero value otherwise.

### GetLoginButtonIconUrlOk

`func (o *RenderedIdentityProviderResult) GetLoginButtonIconUrlOk() (*string, bool)`

GetLoginButtonIconUrlOk returns a tuple with the LoginButtonIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIconUrl

`func (o *RenderedIdentityProviderResult) SetLoginButtonIconUrl(v string)`

SetLoginButtonIconUrl sets LoginButtonIconUrl field to given value.


### GetLoginButtonText

`func (o *RenderedIdentityProviderResult) GetLoginButtonText() string`

GetLoginButtonText returns the LoginButtonText field if non-nil, zero value otherwise.

### GetLoginButtonTextOk

`func (o *RenderedIdentityProviderResult) GetLoginButtonTextOk() (*string, bool)`

GetLoginButtonTextOk returns a tuple with the LoginButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonText

`func (o *RenderedIdentityProviderResult) SetLoginButtonText(v string)`

SetLoginButtonText sets LoginButtonText field to given value.


### GetName

`func (o *RenderedIdentityProviderResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RenderedIdentityProviderResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RenderedIdentityProviderResult) SetName(v string)`

SetName sets Name field to given value.


### GetRenderedAuthorizationEndpoint

`func (o *RenderedIdentityProviderResult) GetRenderedAuthorizationEndpoint() string`

GetRenderedAuthorizationEndpoint returns the RenderedAuthorizationEndpoint field if non-nil, zero value otherwise.

### GetRenderedAuthorizationEndpointOk

`func (o *RenderedIdentityProviderResult) GetRenderedAuthorizationEndpointOk() (*string, bool)`

GetRenderedAuthorizationEndpointOk returns a tuple with the RenderedAuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedAuthorizationEndpoint

`func (o *RenderedIdentityProviderResult) SetRenderedAuthorizationEndpoint(v string)`

SetRenderedAuthorizationEndpoint sets RenderedAuthorizationEndpoint field to given value.


### GetState

`func (o *RenderedIdentityProviderResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RenderedIdentityProviderResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RenderedIdentityProviderResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RenderedIdentityProviderResult) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


