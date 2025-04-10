# FleetCreateConsumptionUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyUrl** | Pointer to **string** | Company URL of the user. | [optional] 
**Email** | **string** | Email address of the user | 
**EnableAutoVerification** | **bool** | Whether to enable auto verification for the user. | 
**LegalCompanyName** | **string** | Legal company name of the user. | 
**Name** | **string** | Name of the user | 
**Password** | **string** | Password of the user | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateConsumptionUserRequest

`func NewFleetCreateConsumptionUserRequest(email string, enableAutoVerification bool, legalCompanyName string, name string, password string, token string, ) *FleetCreateConsumptionUserRequest`

NewFleetCreateConsumptionUserRequest instantiates a new FleetCreateConsumptionUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateConsumptionUserRequestWithDefaults

`func NewFleetCreateConsumptionUserRequestWithDefaults() *FleetCreateConsumptionUserRequest`

NewFleetCreateConsumptionUserRequestWithDefaults instantiates a new FleetCreateConsumptionUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyUrl

`func (o *FleetCreateConsumptionUserRequest) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *FleetCreateConsumptionUserRequest) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *FleetCreateConsumptionUserRequest) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *FleetCreateConsumptionUserRequest) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetEmail

`func (o *FleetCreateConsumptionUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FleetCreateConsumptionUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FleetCreateConsumptionUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnableAutoVerification

`func (o *FleetCreateConsumptionUserRequest) GetEnableAutoVerification() bool`

GetEnableAutoVerification returns the EnableAutoVerification field if non-nil, zero value otherwise.

### GetEnableAutoVerificationOk

`func (o *FleetCreateConsumptionUserRequest) GetEnableAutoVerificationOk() (*bool, bool)`

GetEnableAutoVerificationOk returns a tuple with the EnableAutoVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoVerification

`func (o *FleetCreateConsumptionUserRequest) SetEnableAutoVerification(v bool)`

SetEnableAutoVerification sets EnableAutoVerification field to given value.


### GetLegalCompanyName

`func (o *FleetCreateConsumptionUserRequest) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *FleetCreateConsumptionUserRequest) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *FleetCreateConsumptionUserRequest) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.


### GetName

`func (o *FleetCreateConsumptionUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetCreateConsumptionUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetCreateConsumptionUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *FleetCreateConsumptionUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FleetCreateConsumptionUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FleetCreateConsumptionUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *FleetCreateConsumptionUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateConsumptionUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateConsumptionUserRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


