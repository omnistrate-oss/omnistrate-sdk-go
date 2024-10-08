# LoginWithIdentityProviderRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCode** | Pointer to **string** | The authorization code from the Identity Provider | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**DeviceCode** | Pointer to **string** | The device code from the Identity Provider | [optional] 
**IdentityProviderName** | **string** | The name of the Identity Provider | 
**InvitedEmail** | Pointer to **string** | Email address that the user was invited with | [optional] 
**LegalCompanyName** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI used to get the authorization code | [optional] 

## Methods

### NewLoginWithIdentityProviderRequestBody

`func NewLoginWithIdentityProviderRequestBody(identityProviderName string, ) *LoginWithIdentityProviderRequestBody`

NewLoginWithIdentityProviderRequestBody instantiates a new LoginWithIdentityProviderRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithIdentityProviderRequestBodyWithDefaults

`func NewLoginWithIdentityProviderRequestBodyWithDefaults() *LoginWithIdentityProviderRequestBody`

NewLoginWithIdentityProviderRequestBodyWithDefaults instantiates a new LoginWithIdentityProviderRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCode

`func (o *LoginWithIdentityProviderRequestBody) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *LoginWithIdentityProviderRequestBody) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *LoginWithIdentityProviderRequestBody) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *LoginWithIdentityProviderRequestBody) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *LoginWithIdentityProviderRequestBody) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *LoginWithIdentityProviderRequestBody) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *LoginWithIdentityProviderRequestBody) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *LoginWithIdentityProviderRequestBody) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *LoginWithIdentityProviderRequestBody) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *LoginWithIdentityProviderRequestBody) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *LoginWithIdentityProviderRequestBody) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *LoginWithIdentityProviderRequestBody) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetDeviceCode

`func (o *LoginWithIdentityProviderRequestBody) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *LoginWithIdentityProviderRequestBody) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *LoginWithIdentityProviderRequestBody) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *LoginWithIdentityProviderRequestBody) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *LoginWithIdentityProviderRequestBody) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *LoginWithIdentityProviderRequestBody) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *LoginWithIdentityProviderRequestBody) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetInvitedEmail

`func (o *LoginWithIdentityProviderRequestBody) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *LoginWithIdentityProviderRequestBody) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *LoginWithIdentityProviderRequestBody) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.

### HasInvitedEmail

`func (o *LoginWithIdentityProviderRequestBody) HasInvitedEmail() bool`

HasInvitedEmail returns a boolean if a field has been set.

### GetLegalCompanyName

`func (o *LoginWithIdentityProviderRequestBody) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *LoginWithIdentityProviderRequestBody) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *LoginWithIdentityProviderRequestBody) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *LoginWithIdentityProviderRequestBody) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *LoginWithIdentityProviderRequestBody) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *LoginWithIdentityProviderRequestBody) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *LoginWithIdentityProviderRequestBody) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *LoginWithIdentityProviderRequestBody) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


