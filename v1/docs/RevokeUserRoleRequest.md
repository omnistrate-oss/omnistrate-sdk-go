# RevokeUserRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleType** | **string** | Type of the role | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRevokeUserRoleRequest

`func NewRevokeUserRoleRequest(email string, roleType string, token string, ) *RevokeUserRoleRequest`

NewRevokeUserRoleRequest instantiates a new RevokeUserRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeUserRoleRequestWithDefaults

`func NewRevokeUserRoleRequestWithDefaults() *RevokeUserRoleRequest`

NewRevokeUserRoleRequestWithDefaults instantiates a new RevokeUserRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RevokeUserRoleRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RevokeUserRoleRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RevokeUserRoleRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleType

`func (o *RevokeUserRoleRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RevokeUserRoleRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RevokeUserRoleRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetToken

`func (o *RevokeUserRoleRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RevokeUserRoleRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RevokeUserRoleRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


