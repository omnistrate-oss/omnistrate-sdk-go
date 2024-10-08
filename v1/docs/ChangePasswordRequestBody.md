# ChangePasswordRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Token** | **string** |  | 

## Methods

### NewChangePasswordRequestBody

`func NewChangePasswordRequestBody(email string, password string, token string, ) *ChangePasswordRequestBody`

NewChangePasswordRequestBody instantiates a new ChangePasswordRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordRequestBodyWithDefaults

`func NewChangePasswordRequestBodyWithDefaults() *ChangePasswordRequestBody`

NewChangePasswordRequestBodyWithDefaults instantiates a new ChangePasswordRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ChangePasswordRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChangePasswordRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChangePasswordRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *ChangePasswordRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *ChangePasswordRequestBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChangePasswordRequestBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChangePasswordRequestBody) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


