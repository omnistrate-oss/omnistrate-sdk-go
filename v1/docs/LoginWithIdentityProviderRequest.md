# LoginWithIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** | Additional attributes for the user | [optional] 
**AuthorizationCode** | Pointer to **string** | The authorization code from the Identity Provider | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**DeviceCode** | Pointer to **string** | The device code from the Identity Provider | [optional] 
**IdentityProviderName** | **string** | The name of the identity provider | 
**InvitedEmail** | Pointer to **string** | Email address that the user was invited with | [optional] 
**LegalCompanyName** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI used to get the authorization code | [optional] 

## Methods

### NewLoginWithIdentityProviderRequest

`func NewLoginWithIdentityProviderRequest(identityProviderName string, ) *LoginWithIdentityProviderRequest`

NewLoginWithIdentityProviderRequest instantiates a new LoginWithIdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithIdentityProviderRequestWithDefaults

`func NewLoginWithIdentityProviderRequestWithDefaults() *LoginWithIdentityProviderRequest`

NewLoginWithIdentityProviderRequestWithDefaults instantiates a new LoginWithIdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LoginWithIdentityProviderRequest) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LoginWithIdentityProviderRequest) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LoginWithIdentityProviderRequest) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LoginWithIdentityProviderRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthorizationCode

`func (o *LoginWithIdentityProviderRequest) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *LoginWithIdentityProviderRequest) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *LoginWithIdentityProviderRequest) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *LoginWithIdentityProviderRequest) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *LoginWithIdentityProviderRequest) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *LoginWithIdentityProviderRequest) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *LoginWithIdentityProviderRequest) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *LoginWithIdentityProviderRequest) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *LoginWithIdentityProviderRequest) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *LoginWithIdentityProviderRequest) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *LoginWithIdentityProviderRequest) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *LoginWithIdentityProviderRequest) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetDeviceCode

`func (o *LoginWithIdentityProviderRequest) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *LoginWithIdentityProviderRequest) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *LoginWithIdentityProviderRequest) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *LoginWithIdentityProviderRequest) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *LoginWithIdentityProviderRequest) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *LoginWithIdentityProviderRequest) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *LoginWithIdentityProviderRequest) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetInvitedEmail

`func (o *LoginWithIdentityProviderRequest) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *LoginWithIdentityProviderRequest) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *LoginWithIdentityProviderRequest) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.

### HasInvitedEmail

`func (o *LoginWithIdentityProviderRequest) HasInvitedEmail() bool`

HasInvitedEmail returns a boolean if a field has been set.

### GetLegalCompanyName

`func (o *LoginWithIdentityProviderRequest) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *LoginWithIdentityProviderRequest) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *LoginWithIdentityProviderRequest) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *LoginWithIdentityProviderRequest) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *LoginWithIdentityProviderRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *LoginWithIdentityProviderRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *LoginWithIdentityProviderRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *LoginWithIdentityProviderRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


