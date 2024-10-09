# InviteUserRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleType** | **string** |  | 

## Methods

### NewInviteUserRequestBody

`func NewInviteUserRequestBody(email string, roleType string, ) *InviteUserRequestBody`

NewInviteUserRequestBody instantiates a new InviteUserRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserRequestBodyWithDefaults

`func NewInviteUserRequestBodyWithDefaults() *InviteUserRequestBody`

NewInviteUserRequestBodyWithDefaults instantiates a new InviteUserRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUserRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUserRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUserRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleType

`func (o *InviteUserRequestBody) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InviteUserRequestBody) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InviteUserRequestBody) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

