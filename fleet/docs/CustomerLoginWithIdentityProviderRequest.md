# CustomerLoginWithIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCode** | **string** | The authorization code from the Identity Provider | 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**IdentityProviderName** | **string** | The name of the identity provider | 
**InvitedEmail** | Pointer to **string** | Email address that the user was invited with | [optional] 
**LegalCompanyName** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI used to get the authorization code | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCustomerLoginWithIdentityProviderRequest

`func NewCustomerLoginWithIdentityProviderRequest(authorizationCode string, identityProviderName string, token string, ) *CustomerLoginWithIdentityProviderRequest`

NewCustomerLoginWithIdentityProviderRequest instantiates a new CustomerLoginWithIdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLoginWithIdentityProviderRequestWithDefaults

`func NewCustomerLoginWithIdentityProviderRequestWithDefaults() *CustomerLoginWithIdentityProviderRequest`

NewCustomerLoginWithIdentityProviderRequestWithDefaults instantiates a new CustomerLoginWithIdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequest) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequest) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.


### GetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequest) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequest) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.

### HasInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest) HasInvitedEmail() bool`

HasInvitedEmail returns a boolean if a field has been set.

### GetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetToken

`func (o *CustomerLoginWithIdentityProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerLoginWithIdentityProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerLoginWithIdentityProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


