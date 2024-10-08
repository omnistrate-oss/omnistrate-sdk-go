# OrgUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | 
**InvitedAt** | **string** | The time when the user was invited | 
**Name** | **string** | The name of the user | 
**RoleType** | **string** |  | 
**UserId** | **string** | The User ID | 

## Methods

### NewOrgUsers

`func NewOrgUsers(email string, invitedAt string, name string, roleType string, userId string, ) *OrgUsers`

NewOrgUsers instantiates a new OrgUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUsersWithDefaults

`func NewOrgUsersWithDefaults() *OrgUsers`

NewOrgUsersWithDefaults instantiates a new OrgUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrgUsers) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgUsers) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgUsers) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInvitedAt

`func (o *OrgUsers) GetInvitedAt() string`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *OrgUsers) GetInvitedAtOk() (*string, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *OrgUsers) SetInvitedAt(v string)`

SetInvitedAt sets InvitedAt field to given value.


### GetName

`func (o *OrgUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgUsers) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *OrgUsers) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *OrgUsers) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *OrgUsers) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetUserId

`func (o *OrgUsers) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgUsers) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgUsers) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


