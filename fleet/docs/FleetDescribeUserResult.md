# FleetDescribeUserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time the user was created. | 
**Subscriptions** | [**[]UserSubscription**](UserSubscription.md) | List of subscriptions associated with the user. | 
**UserId** | **string** | ID of a User | 

## Methods

### NewFleetDescribeUserResult

`func NewFleetDescribeUserResult(createdAt string, subscriptions []UserSubscription, userId string, ) *FleetDescribeUserResult`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


