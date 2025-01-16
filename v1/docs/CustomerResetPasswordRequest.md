# CustomerResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCustomerResetPasswordRequest

`func NewCustomerResetPasswordRequest(email string, token string, ) *CustomerResetPasswordRequest`

NewCustomerResetPasswordRequest instantiates a new CustomerResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResetPasswordRequestWithDefaults

`func NewCustomerResetPasswordRequestWithDefaults() *CustomerResetPasswordRequest`

NewCustomerResetPasswordRequestWithDefaults instantiates a new CustomerResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomerResetPasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerResetPasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerResetPasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetToken

`func (o *CustomerResetPasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerResetPasswordRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerResetPasswordRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


