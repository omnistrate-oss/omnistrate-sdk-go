# CustomerInviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCustomerInviteUserRequest

`func NewCustomerInviteUserRequest(email string, token string, ) *CustomerInviteUserRequest`

NewCustomerInviteUserRequest instantiates a new CustomerInviteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInviteUserRequestWithDefaults

`func NewCustomerInviteUserRequestWithDefaults() *CustomerInviteUserRequest`

NewCustomerInviteUserRequestWithDefaults instantiates a new CustomerInviteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomerInviteUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInviteUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInviteUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetToken

`func (o *CustomerInviteUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerInviteUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerInviteUserRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


