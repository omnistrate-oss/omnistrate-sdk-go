# AccessSideUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the user was created. | [optional] 
**Email** | Pointer to **string** | The user email. | [optional] 
**Enabled** | Pointer to **bool** | Is the user enabled. | [optional] 
**InstanceCount** | Pointer to **int64** | The number of active instances the user has. | [optional] 
**LastModifiedAt** | Pointer to **string** | The last modified time of the user. | [optional] 
**LastModifiedByUserID** | Pointer to **string** | ID of a User | [optional] 
**LastModifiedByUserName** | Pointer to **string** | The user name of the last modifier. | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**OrgUrl** | Pointer to **string** | The organization URL. | [optional] 
**Status** | Pointer to **string** | The status of the user. | [optional] 
**SubscriptionCount** | Pointer to **int64** | The number of subscriptions the user has. | [optional] 
**Token** | Pointer to **string** | Token to validate the user, if the user is not enabled. | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The user name. | [optional] 

## Methods

### NewAccessSideUser

`func NewAccessSideUser() *AccessSideUser`

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

### HasCreatedAt

`func (o *AccessSideUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasEmail

`func (o *AccessSideUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

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

### HasEnabled

`func (o *AccessSideUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

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

### HasInstanceCount

`func (o *AccessSideUser) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

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

### HasLastModifiedAt

`func (o *AccessSideUser) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

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

### HasOrgId

`func (o *AccessSideUser) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

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

### HasOrgName

`func (o *AccessSideUser) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

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

### HasStatus

`func (o *AccessSideUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### HasSubscriptionCount

`func (o *AccessSideUser) HasSubscriptionCount() bool`

HasSubscriptionCount returns a boolean if a field has been set.

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

### HasUserId

`func (o *AccessSideUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

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

### HasUserName

`func (o *AccessSideUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


