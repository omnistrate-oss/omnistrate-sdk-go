# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time the user was created. | 
**Email** | **string** | The user email. | 
**Enabled** | Pointer to **bool** | Is the user enabled. | [optional] 
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**InstanceCount** | **int64** | The number of active instances the user has. | 
**LastModifiedAt** | **string** | The last modified time of the user. | 
**LastModifiedByUserID** | Pointer to **string** | The user ID of the last modifier. | [optional] 
**LastModifiedByUserName** | Pointer to **string** | The user name of the last modifier. | [optional] 
**OrgId** | **string** | The organization ID. | 
**OrgName** | **string** | The organization name. | 
**OrgUrl** | Pointer to **string** | The organization URL. | [optional] 
**ServiceId** | **string** | The service ID this workflow belongs to. | 
**Status** | **string** | The status of the user. | 
**SubscriptionCount** | Pointer to **int64** | The number of subscriptions the user has. | [optional] 
**Token** | Pointer to **string** | Token to validate the user, if the user is not enabled. | [optional] 
**UserId** | **string** | The user ID. | 
**UserName** | **string** | The user name. | 
**UserSubscriptionRole** | Pointer to **string** | The user subscription role. | [optional] 

## Methods

### NewUser

`func NewUser(createdAt string, email string, environmentId string, instanceCount int64, lastModifiedAt string, orgId string, orgName string, serviceId string, status string, userId string, userName string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *User) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *User) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *User) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceCount

`func (o *User) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *User) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *User) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.


### GetLastModifiedAt

`func (o *User) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *User) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *User) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.


### GetLastModifiedByUserID

`func (o *User) GetLastModifiedByUserID() string`

GetLastModifiedByUserID returns the LastModifiedByUserID field if non-nil, zero value otherwise.

### GetLastModifiedByUserIDOk

`func (o *User) GetLastModifiedByUserIDOk() (*string, bool)`

GetLastModifiedByUserIDOk returns a tuple with the LastModifiedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserID

`func (o *User) SetLastModifiedByUserID(v string)`

SetLastModifiedByUserID sets LastModifiedByUserID field to given value.

### HasLastModifiedByUserID

`func (o *User) HasLastModifiedByUserID() bool`

HasLastModifiedByUserID returns a boolean if a field has been set.

### GetLastModifiedByUserName

`func (o *User) GetLastModifiedByUserName() string`

GetLastModifiedByUserName returns the LastModifiedByUserName field if non-nil, zero value otherwise.

### GetLastModifiedByUserNameOk

`func (o *User) GetLastModifiedByUserNameOk() (*string, bool)`

GetLastModifiedByUserNameOk returns a tuple with the LastModifiedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserName

`func (o *User) SetLastModifiedByUserName(v string)`

SetLastModifiedByUserName sets LastModifiedByUserName field to given value.

### HasLastModifiedByUserName

`func (o *User) HasLastModifiedByUserName() bool`

HasLastModifiedByUserName returns a boolean if a field has been set.

### GetOrgId

`func (o *User) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *User) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *User) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *User) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *User) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *User) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgUrl

`func (o *User) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *User) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *User) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.

### HasOrgUrl

`func (o *User) HasOrgUrl() bool`

HasOrgUrl returns a boolean if a field has been set.

### GetServiceId

`func (o *User) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *User) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *User) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionCount

`func (o *User) GetSubscriptionCount() int64`

GetSubscriptionCount returns the SubscriptionCount field if non-nil, zero value otherwise.

### GetSubscriptionCountOk

`func (o *User) GetSubscriptionCountOk() (*int64, bool)`

GetSubscriptionCountOk returns a tuple with the SubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCount

`func (o *User) SetSubscriptionCount(v int64)`

SetSubscriptionCount sets SubscriptionCount field to given value.

### HasSubscriptionCount

`func (o *User) HasSubscriptionCount() bool`

HasSubscriptionCount returns a boolean if a field has been set.

### GetToken

`func (o *User) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *User) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *User) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *User) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserId

`func (o *User) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *User) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *User) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *User) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *User) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *User) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserSubscriptionRole

`func (o *User) GetUserSubscriptionRole() string`

GetUserSubscriptionRole returns the UserSubscriptionRole field if non-nil, zero value otherwise.

### GetUserSubscriptionRoleOk

`func (o *User) GetUserSubscriptionRoleOk() (*string, bool)`

GetUserSubscriptionRoleOk returns a tuple with the UserSubscriptionRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubscriptionRole

`func (o *User) SetUserSubscriptionRole(v string)`

SetUserSubscriptionRole sets UserSubscriptionRole field to given value.

### HasUserSubscriptionRole

`func (o *User) HasUserSubscriptionRole() bool`

HasUserSubscriptionRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


