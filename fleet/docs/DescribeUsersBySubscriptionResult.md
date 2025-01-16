# DescribeUsersBySubscriptionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the subscription | 
**SubscriptionUsers** | [**[]SubscriptionUsers**](SubscriptionUsers.md) | The users associated with the corresponding subscription | 

## Methods

### NewDescribeUsersBySubscriptionResult

`func NewDescribeUsersBySubscriptionResult(id string, subscriptionUsers []SubscriptionUsers, ) *DescribeUsersBySubscriptionResult`

NewDescribeUsersBySubscriptionResult instantiates a new DescribeUsersBySubscriptionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersBySubscriptionResultWithDefaults

`func NewDescribeUsersBySubscriptionResultWithDefaults() *DescribeUsersBySubscriptionResult`

NewDescribeUsersBySubscriptionResultWithDefaults instantiates a new DescribeUsersBySubscriptionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeUsersBySubscriptionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeUsersBySubscriptionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeUsersBySubscriptionResult) SetId(v string)`

SetId sets Id field to given value.


### GetSubscriptionUsers

`func (o *DescribeUsersBySubscriptionResult) GetSubscriptionUsers() []SubscriptionUsers`

GetSubscriptionUsers returns the SubscriptionUsers field if non-nil, zero value otherwise.

### GetSubscriptionUsersOk

`func (o *DescribeUsersBySubscriptionResult) GetSubscriptionUsersOk() (*[]SubscriptionUsers, bool)`

GetSubscriptionUsersOk returns a tuple with the SubscriptionUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUsers

`func (o *DescribeUsersBySubscriptionResult) SetSubscriptionUsers(v []SubscriptionUsers)`

SetSubscriptionUsers sets SubscriptionUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


