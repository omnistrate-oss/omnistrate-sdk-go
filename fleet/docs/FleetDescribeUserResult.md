# FleetDescribeUserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the user was created. | [optional] 
**Email** | Pointer to **string** | The user email. | [optional] 
**Enabled** | Pointer to **bool** | Is the user enabled. | [optional] 
**LastModifiedAt** | Pointer to **string** | The last modified time of the user. | [optional] 
**LastModifiedByUserID** | Pointer to **string** | ID of a User | [optional] 
**LastModifiedByUserName** | Pointer to **string** | The user name of the last modifier. | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**OrgUrl** | Pointer to **string** | The organization URL. | [optional] 
**Status** | Pointer to **string** | The status of the user. | [optional] 
**Subscriptions** | Pointer to [**[]UserSubscription**](UserSubscription.md) | List of subscriptions associated with the user. | [optional] 
**Token** | Pointer to **string** | Token to validate the user, if the user is not enabled. | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The user name. | [optional] 

## Methods

### NewFleetDescribeUserResult

`func NewFleetDescribeUserResult() *FleetDescribeUserResult`

NewFleetDescribeUserResult instantiates a new FleetDescribeUserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeUserResultWithDefaults

`func NewFleetDescribeUserResultWithDefaults() *FleetDescribeUserResult`

NewFleetDescribeUserResultWithDefaults instantiates a new FleetDescribeUserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FleetDescribeUserResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetDescribeUserResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetDescribeUserResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FleetDescribeUserResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *FleetDescribeUserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FleetDescribeUserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FleetDescribeUserResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FleetDescribeUserResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *FleetDescribeUserResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FleetDescribeUserResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FleetDescribeUserResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FleetDescribeUserResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *FleetDescribeUserResult) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *FleetDescribeUserResult) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *FleetDescribeUserResult) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *FleetDescribeUserResult) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLastModifiedByUserID

`func (o *FleetDescribeUserResult) GetLastModifiedByUserID() string`

GetLastModifiedByUserID returns the LastModifiedByUserID field if non-nil, zero value otherwise.

### GetLastModifiedByUserIDOk

`func (o *FleetDescribeUserResult) GetLastModifiedByUserIDOk() (*string, bool)`

GetLastModifiedByUserIDOk returns a tuple with the LastModifiedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserID

`func (o *FleetDescribeUserResult) SetLastModifiedByUserID(v string)`

SetLastModifiedByUserID sets LastModifiedByUserID field to given value.

### HasLastModifiedByUserID

`func (o *FleetDescribeUserResult) HasLastModifiedByUserID() bool`

HasLastModifiedByUserID returns a boolean if a field has been set.

### GetLastModifiedByUserName

`func (o *FleetDescribeUserResult) GetLastModifiedByUserName() string`

GetLastModifiedByUserName returns the LastModifiedByUserName field if non-nil, zero value otherwise.

### GetLastModifiedByUserNameOk

`func (o *FleetDescribeUserResult) GetLastModifiedByUserNameOk() (*string, bool)`

GetLastModifiedByUserNameOk returns a tuple with the LastModifiedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserName

`func (o *FleetDescribeUserResult) SetLastModifiedByUserName(v string)`

SetLastModifiedByUserName sets LastModifiedByUserName field to given value.

### HasLastModifiedByUserName

`func (o *FleetDescribeUserResult) HasLastModifiedByUserName() bool`

HasLastModifiedByUserName returns a boolean if a field has been set.

### GetOrgId

`func (o *FleetDescribeUserResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FleetDescribeUserResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FleetDescribeUserResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FleetDescribeUserResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *FleetDescribeUserResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FleetDescribeUserResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FleetDescribeUserResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *FleetDescribeUserResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgUrl

`func (o *FleetDescribeUserResult) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *FleetDescribeUserResult) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *FleetDescribeUserResult) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.

### HasOrgUrl

`func (o *FleetDescribeUserResult) HasOrgUrl() bool`

HasOrgUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FleetDescribeUserResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeUserResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeUserResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetDescribeUserResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptions

`func (o *FleetDescribeUserResult) GetSubscriptions() []UserSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FleetDescribeUserResult) GetSubscriptionsOk() (*[]UserSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FleetDescribeUserResult) SetSubscriptions(v []UserSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FleetDescribeUserResult) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetToken

`func (o *FleetDescribeUserResult) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetDescribeUserResult) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetDescribeUserResult) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FleetDescribeUserResult) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserId

`func (o *FleetDescribeUserResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetDescribeUserResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetDescribeUserResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FleetDescribeUserResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *FleetDescribeUserResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FleetDescribeUserResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FleetDescribeUserResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FleetDescribeUserResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


