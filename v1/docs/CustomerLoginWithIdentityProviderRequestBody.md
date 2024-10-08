# CustomerLoginWithIdentityProviderRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationCode** | **string** | The authorization code from the Identity Provider | 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**EnvironmentType** | Pointer to **string** | The environment type of the portal that the customer is signing in to | [optional] 
**IdentityProviderName** | **string** | The name of the Identity Provider | 
**InvitedEmail** | Pointer to **string** | Email address that the user was invited with | [optional] 
**LegalCompanyName** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI used to get the authorization code | [optional] 

## Methods

### NewCustomerLoginWithIdentityProviderRequestBody

`func NewCustomerLoginWithIdentityProviderRequestBody(authorizationCode string, identityProviderName string, ) *CustomerLoginWithIdentityProviderRequestBody`

NewCustomerLoginWithIdentityProviderRequestBody instantiates a new CustomerLoginWithIdentityProviderRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLoginWithIdentityProviderRequestBodyWithDefaults

`func NewCustomerLoginWithIdentityProviderRequestBodyWithDefaults() *CustomerLoginWithIdentityProviderRequestBody`

NewCustomerLoginWithIdentityProviderRequestBodyWithDefaults instantiates a new CustomerLoginWithIdentityProviderRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.


### GetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetInvitedEmail() string`

GetInvitedEmail returns the InvitedEmail field if non-nil, zero value otherwise.

### GetInvitedEmailOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetInvitedEmailOk() (*string, bool)`

GetInvitedEmailOk returns a tuple with the InvitedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetInvitedEmail(v string)`

SetInvitedEmail sets InvitedEmail field to given value.

### HasInvitedEmail

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasInvitedEmail() bool`

HasInvitedEmail returns a boolean if a field has been set.

### GetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CustomerLoginWithIdentityProviderRequestBody) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequestBody) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CustomerLoginWithIdentityProviderRequestBody) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


