# UpdatePasswordRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** |  | [optional] 
**CurrentPasswordHash** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewUpdatePasswordRequest2

`func NewUpdatePasswordRequest2(password string, ) *UpdatePasswordRequest2`

NewUpdatePasswordRequest2 instantiates a new UpdatePasswordRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordRequest2WithDefaults

`func NewUpdatePasswordRequest2WithDefaults() *UpdatePasswordRequest2`

NewUpdatePasswordRequest2WithDefaults instantiates a new UpdatePasswordRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdatePasswordRequest2) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdatePasswordRequest2) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdatePasswordRequest2) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdatePasswordRequest2) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetCurrentPasswordHash

`func (o *UpdatePasswordRequest2) GetCurrentPasswordHash() string`

GetCurrentPasswordHash returns the CurrentPasswordHash field if non-nil, zero value otherwise.

### GetCurrentPasswordHashOk

`func (o *UpdatePasswordRequest2) GetCurrentPasswordHashOk() (*string, bool)`

GetCurrentPasswordHashOk returns a tuple with the CurrentPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPasswordHash

`func (o *UpdatePasswordRequest2) SetCurrentPasswordHash(v string)`

SetCurrentPasswordHash sets CurrentPasswordHash field to given value.

### HasCurrentPasswordHash

`func (o *UpdatePasswordRequest2) HasCurrentPasswordHash() bool`

HasCurrentPasswordHash returns a boolean if a field has been set.

### GetPassword

`func (o *UpdatePasswordRequest2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordRequest2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordRequest2) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


