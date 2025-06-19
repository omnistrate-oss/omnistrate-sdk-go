# CustomerLoginWithIdentityProviderRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCode** | **string** | The authorization code from the Identity Provider | 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**EnvironmentType** | Pointer to **string** | The environment type of the portal that the customer is signing in to | [optional] 
**IdentityProviderName** | Pointer to **string** | The name or type of the Identity Provider | [optional] 
**InvitedEmail** | Pointer to **string** | Email address that the user was invited with | [optional] 
**LegalCompanyName** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI used to get the authorization code | [optional] 
**State** | Pointer to **string** | The state parameter used to prevent CSRF attacks | [optional] 

## Methods

### NewCustomerLoginWithIdentityProviderRequest2

`func NewCustomerLoginWithIdentityProviderRequest2(authorizationCode string, ) *CustomerLoginWithIdentityProviderRequest2`

NewCustomerLoginWithIdentityProviderRequest2 instantiates a new CustomerLoginWithIdentityProviderRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLoginWithIdentityProviderRequest2WithDefaults

`func NewCustomerLoginWithIdentityProviderRequest2WithDefaults() *CustomerLoginWithIdentityProviderRequest2`

NewCustomerLoginWithIdentityProviderRequest2WithDefaults instantiates a new CustomerLoginWithIdentityProviderRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequest2) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequest2) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.


### GetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest2) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest2) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequest2) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest2) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest2) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequest2) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequest2) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequest2) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.

### HasIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequest2) HasIdentityProviderName() bool`

HasIdentityProviderName returns a boolean if a field has been set.

### GetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest2) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest2) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.

### HasInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequest2) HasInvitedEmail() bool`

HasInvitedEmail returns a boolean if a field has been set.

### GetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest2) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest2) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequest2) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest2) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest2) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequest2) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetState

`func (o *CustomerLoginWithIdentityProviderRequest2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerLoginWithIdentityProviderRequest2) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerLoginWithIdentityProviderRequest2) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerLoginWithIdentityProviderRequest2) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


