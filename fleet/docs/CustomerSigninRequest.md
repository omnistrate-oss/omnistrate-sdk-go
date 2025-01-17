# CustomerSigninRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**HashedPassword** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCustomerSigninRequest

`func NewCustomerSigninRequest(email string, token string, ) *CustomerSigninRequest`

NewCustomerSigninRequest instantiates a new CustomerSigninRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSigninRequestWithDefaults

`func NewCustomerSigninRequestWithDefaults() *CustomerSigninRequest`

NewCustomerSigninRequestWithDefaults instantiates a new CustomerSigninRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomerSigninRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerSigninRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerSigninRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnvironmentType

`func (o *CustomerSigninRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CustomerSigninRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CustomerSigninRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CustomerSigninRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetHashedPassword

`func (o *CustomerSigninRequest) GetHashedPassword() string`

GetHashedPassword returns the HashedPassword field if non-nil, zero value otherwise.

### GetHashedPasswordOk

`func (o *CustomerSigninRequest) GetHashedPasswordOk() (*string, bool)`

GetHashedPasswordOk returns a tuple with the HashedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedPassword

`func (o *CustomerSigninRequest) SetHashedPassword(v string)`

SetHashedPassword sets HashedPassword field to given value.

### HasHashedPassword

`func (o *CustomerSigninRequest) HasHashedPassword() bool`

HasHashedPassword returns a boolean if a field has been set.

### GetPassword

`func (o *CustomerSigninRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerSigninRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerSigninRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CustomerSigninRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CustomerSigninRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerSigninRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerSigninRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


