# AccessSideUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time the user was created. | 
**Email** | **string** | The user email. | 
**Enabled** | **bool** | Is the user enabled. | 
**InstanceCount** | **int64** | The number of active instances the user has. | 
**LastModifiedAt** | **string** | The last modified time of the user. | 
**LastModifiedByUserID** | Pointer to **string** | ID of a User | [optional] 
**LastModifiedByUserName** | Pointer to **string** | The user name of the last modifier. | [optional] 
**OrgId** | **string** | ID of an Org | 
**OrgName** | **string** | The organization name. | 
**OrgUrl** | Pointer to **string** | The organization URL. | [optional] 
**Status** | **string** | The status of the user. | 
**SubscriptionCount** | **int64** | The number of subscriptions the user has. | 
**Token** | Pointer to **string** | Token to validate the user, if the user is not enabled. | [optional] 
**UserId** | **string** | ID of a User | 
**UserName** | **string** | The user name. | 

## Methods

### NewAccessSideUser

`func NewAccessSideUser(createdAt string, email string, enabled bool, instanceCount int64, lastModifiedAt string, orgId string, orgName string, status string, subscriptionCount int64, userId string, userName string, ) *AccessSideUser`

NewAccessSideUser instantiates a new AccessSideUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessSideUserWithDefaults

`func NewAccessSideUserWithDefaults() *AccessSideUser`

NewAccessSideUserWithDefaults instantiates a new AccessSideUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AccessSideUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessSideUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessSideUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *AccessSideUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccessSideUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccessSideUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *AccessSideUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccessSideUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccessSideUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInstanceCount

`func (o *AccessSideUser) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AccessSideUser) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AccessSideUser) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.


### GetLastModifiedAt

`func (o *AccessSideUser) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *AccessSideUser) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *AccessSideUser) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.


### GetLastModifiedByUserID

`func (o *AccessSideUser) GetLastModifiedByUserID() string`

GetLastModifiedByUserID returns the LastModifiedByUserID field if non-nil, zero value otherwise.

### GetLastModifiedByUserIDOk

`func (o *AccessSideUser) GetLastModifiedByUserIDOk() (*string, bool)`

GetLastModifiedByUserIDOk returns a tuple with the LastModifiedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserID

`func (o *AccessSideUser) SetLastModifiedByUserID(v string)`

SetLastModifiedByUserID sets LastModifiedByUserID field to given value.

### HasLastModifiedByUserID

`func (o *AccessSideUser) HasLastModifiedByUserID() bool`

HasLastModifiedByUserID returns a boolean if a field has been set.

### GetLastModifiedByUserName

`func (o *AccessSideUser) GetLastModifiedByUserName() string`

GetLastModifiedByUserName returns the LastModifiedByUserName field if non-nil, zero value otherwise.

### GetLastModifiedByUserNameOk

`func (o *AccessSideUser) GetLastModifiedByUserNameOk() (*string, bool)`

GetLastModifiedByUserNameOk returns a tuple with the LastModifiedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserName

`func (o *AccessSideUser) SetLastModifiedByUserName(v string)`

SetLastModifiedByUserName sets LastModifiedByUserName field to given value.

### HasLastModifiedByUserName

`func (o *AccessSideUser) HasLastModifiedByUserName() bool`

HasLastModifiedByUserName returns a boolean if a field has been set.

### GetOrgId

`func (o *AccessSideUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AccessSideUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AccessSideUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *AccessSideUser) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AccessSideUser) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AccessSideUser) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgUrl

`func (o *AccessSideUser) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *AccessSideUser) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *AccessSideUser) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.

### HasOrgUrl

`func (o *AccessSideUser) HasOrgUrl() bool`

HasOrgUrl returns a boolean if a field has been set.

### GetStatus

`func (o *AccessSideUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessSideUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessSideUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionCount

`func (o *AccessSideUser) GetSubscriptionCount() int64`

GetSubscriptionCount returns the SubscriptionCount field if non-nil, zero value otherwise.

### GetSubscriptionCountOk

`func (o *AccessSideUser) GetSubscriptionCountOk() (*int64, bool)`

GetSubscriptionCountOk returns a tuple with the SubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCount

`func (o *AccessSideUser) SetSubscriptionCount(v int64)`

SetSubscriptionCount sets SubscriptionCount field to given value.


### GetToken

`func (o *AccessSideUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessSideUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessSideUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccessSideUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserId

`func (o *AccessSideUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessSideUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessSideUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *AccessSideUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AccessSideUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AccessSideUser) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


