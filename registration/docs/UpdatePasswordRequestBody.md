# UpdatePasswordRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** |  | [optional] 
**CurrentPasswordHash** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewUpdatePasswordRequestBody

`func NewUpdatePasswordRequestBody(password string, ) *UpdatePasswordRequestBody`

NewUpdatePasswordRequestBody instantiates a new UpdatePasswordRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordRequestBodyWithDefaults

`func NewUpdatePasswordRequestBodyWithDefaults() *UpdatePasswordRequestBody`

NewUpdatePasswordRequestBodyWithDefaults instantiates a new UpdatePasswordRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdatePasswordRequestBody) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdatePasswordRequestBody) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdatePasswordRequestBody) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdatePasswordRequestBody) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetCurrentPasswordHash

`func (o *UpdatePasswordRequestBody) GetCurrentPasswordHash() string`

GetCurrentPasswordHash returns the CurrentPasswordHash field if non-nil, zero value otherwise.

### GetCurrentPasswordHashOk

`func (o *UpdatePasswordRequestBody) GetCurrentPasswordHashOk() (*string, bool)`

GetCurrentPasswordHashOk returns a tuple with the CurrentPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPasswordHash

`func (o *UpdatePasswordRequestBody) SetCurrentPasswordHash(v string)`

SetCurrentPasswordHash sets CurrentPasswordHash field to given value.

### HasCurrentPasswordHash

`func (o *UpdatePasswordRequestBody) HasCurrentPasswordHash() bool`

HasCurrentPasswordHash returns a boolean if a field has been set.

### GetPassword

`func (o *UpdatePasswordRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


