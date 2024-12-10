# UpdatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** |  | [optional] 
**CurrentPasswordHash** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdatePasswordRequest

`func NewUpdatePasswordRequest(password string, token string, ) *UpdatePasswordRequest`

NewUpdatePasswordRequest instantiates a new UpdatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordRequestWithDefaults

`func NewUpdatePasswordRequestWithDefaults() *UpdatePasswordRequest`

NewUpdatePasswordRequestWithDefaults instantiates a new UpdatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdatePasswordRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdatePasswordRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdatePasswordRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdatePasswordRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetCurrentPasswordHash

`func (o *UpdatePasswordRequest) GetCurrentPasswordHash() string`

GetCurrentPasswordHash returns the CurrentPasswordHash field if non-nil, zero value otherwise.

### GetCurrentPasswordHashOk

`func (o *UpdatePasswordRequest) GetCurrentPasswordHashOk() (*string, bool)`

GetCurrentPasswordHashOk returns a tuple with the CurrentPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPasswordHash

`func (o *UpdatePasswordRequest) SetCurrentPasswordHash(v string)`

SetCurrentPasswordHash sets CurrentPasswordHash field to given value.

### HasCurrentPasswordHash

`func (o *UpdatePasswordRequest) HasCurrentPasswordHash() bool`

HasCurrentPasswordHash returns a boolean if a field has been set.

### GetPassword

`func (o *UpdatePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *UpdatePasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdatePasswordRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdatePasswordRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


