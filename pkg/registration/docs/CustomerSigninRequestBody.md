# CustomerSigninRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**EnvironmentType** | Pointer to **string** | The environment type of the portal that the customer is signing in to | [optional] 
**HashedPassword** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerSigninRequestBody

`func NewCustomerSigninRequestBody(email string, ) *CustomerSigninRequestBody`

NewCustomerSigninRequestBody instantiates a new CustomerSigninRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSigninRequestBodyWithDefaults

`func NewCustomerSigninRequestBodyWithDefaults() *CustomerSigninRequestBody`

NewCustomerSigninRequestBodyWithDefaults instantiates a new CustomerSigninRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomerSigninRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerSigninRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerSigninRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnvironmentType

`func (o *CustomerSigninRequestBody) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CustomerSigninRequestBody) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CustomerSigninRequestBody) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *CustomerSigninRequestBody) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetHashedPassword

`func (o *CustomerSigninRequestBody) GetHashedPassword() string`

GetHashedPassword returns the HashedPassword field if non-nil, zero value otherwise.

### GetHashedPasswordOk

`func (o *CustomerSigninRequestBody) GetHashedPasswordOk() (*string, bool)`

GetHashedPasswordOk returns a tuple with the HashedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedPassword

`func (o *CustomerSigninRequestBody) SetHashedPassword(v string)`

SetHashedPassword sets HashedPassword field to given value.

### HasHashedPassword

`func (o *CustomerSigninRequestBody) HasHashedPassword() bool`

HasHashedPassword returns a boolean if a field has been set.

### GetPassword

`func (o *CustomerSigninRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerSigninRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerSigninRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CustomerSigninRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


