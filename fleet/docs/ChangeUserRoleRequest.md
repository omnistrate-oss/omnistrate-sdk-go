# ChangeUserRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**NewRoleType** | **string** | Type of the role | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewChangeUserRoleRequest

`func NewChangeUserRoleRequest(email string, newRoleType string, token string, ) *ChangeUserRoleRequest`

NewChangeUserRoleRequest instantiates a new ChangeUserRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeUserRoleRequestWithDefaults

`func NewChangeUserRoleRequestWithDefaults() *ChangeUserRoleRequest`

NewChangeUserRoleRequestWithDefaults instantiates a new ChangeUserRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ChangeUserRoleRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChangeUserRoleRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChangeUserRoleRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewRoleType

`func (o *ChangeUserRoleRequest) GetNewRoleType() string`

GetNewRoleType returns the NewRoleType field if non-nil, zero value otherwise.

### GetNewRoleTypeOk

`func (o *ChangeUserRoleRequest) GetNewRoleTypeOk() (*string, bool)`

GetNewRoleTypeOk returns a tuple with the NewRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRoleType

`func (o *ChangeUserRoleRequest) SetNewRoleType(v string)`

SetNewRoleType sets NewRoleType field to given value.


### GetToken

`func (o *ChangeUserRoleRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChangeUserRoleRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChangeUserRoleRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


