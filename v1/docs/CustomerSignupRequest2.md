# CustomerSignupRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**Email** | **string** | Email address of the end-user | 
**LegalCompanyName** | Pointer to **string** |  | [optional] [default to ""]
**Name** | **string** | Name of the end-user | 
**Password** | **string** |  | 

## Methods

### NewCustomerSignupRequest2

`func NewCustomerSignupRequest2(email string, name string, password string, ) *CustomerSignupRequest2`

NewCustomerSignupRequest2 instantiates a new CustomerSignupRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSignupRequest2WithDefaults

`func NewCustomerSignupRequest2WithDefaults() *CustomerSignupRequest2`

NewCustomerSignupRequest2WithDefaults instantiates a new CustomerSignupRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyDescription

`func (o *CustomerSignupRequest2) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CustomerSignupRequest2) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CustomerSignupRequest2) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CustomerSignupRequest2) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *CustomerSignupRequest2) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *CustomerSignupRequest2) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *CustomerSignupRequest2) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *CustomerSignupRequest2) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerSignupRequest2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerSignupRequest2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerSignupRequest2) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLegalCompanyName

`func (o *CustomerSignupRequest2) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *CustomerSignupRequest2) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *CustomerSignupRequest2) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.

### HasLegalCompanyName

`func (o *CustomerSignupRequest2) HasLegalCompanyName() bool`

HasLegalCompanyName returns a boolean if a field has been set.

### GetName

`func (o *CustomerSignupRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerSignupRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerSignupRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CustomerSignupRequest2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerSignupRequest2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerSignupRequest2) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


