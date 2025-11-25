# CustomerSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** | Additional attributes for the user signup | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**Email** | **string** | Email address of the end-user | 
**LegalCompanyName** | Pointer to **string** |  | [optional] [default to ""]
**Name** | **string** | Name of the end-user | 
**Password** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCustomerSignupRequest

`func NewCustomerSignupRequest(email string, name string, password string, token string, ) *CustomerSignupRequest`

NewCustomerSignupRequest instantiates a new CustomerSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSignupRequestWithDefaults

`func NewCustomerSignupRequestWithDefaults() *CustomerSignupRequest`

NewCustomerSignupRequestWithDefaults instantiates a new CustomerSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CustomerSignupRequest) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSignupRequest) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerSignupRequest) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CustomerSignupRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *CustomerSignupRequest) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CustomerSignupRequest) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CustomerSignupRequest) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CustomerSignupRequest) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CustomerSignupRequest) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CustomerSignupRequest) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CustomerSignupRequest) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CustomerSignupRequest) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerSignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerSignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerSignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLegalCompanyName

`func (o *CustomerSignupRequest) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *CustomerSignupRequest) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *CustomerSignupRequest) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *CustomerSignupRequest) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetName

`func (o *CustomerSignupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerSignupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerSignupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CustomerSignupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerSignupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerSignupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *CustomerSignupRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerSignupRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerSignupRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


