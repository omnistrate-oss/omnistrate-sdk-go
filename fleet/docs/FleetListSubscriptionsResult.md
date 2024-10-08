# FleetListSubscriptionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of subscription IDs | 
**NextPageToken** | Pointer to **string** | The next token to use when listing subscriptions | [optional] 
**Subscriptions** | Pointer to [**[]FleetDescribeSubscriptionResult**](FleetDescribeSubscriptionResult.md) | List of subscriptions | [optional] 

## Methods

### NewFleetListSubscriptionsResult

`func NewFleetListSubscriptionsResult(ids []string, ) *FleetListSubscriptionsResult`

NewFleetListSubscriptionsResult instantiates a new FleetListSubscriptionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListSubscriptionsResultWithDefaults

`func NewFleetListSubscriptionsResultWithDefaults() *FleetListSubscriptionsResult`

NewFleetListSubscriptionsResultWithDefaults instantiates a new FleetListSubscriptionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *FleetListSubscriptionsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *FleetListSubscriptionsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *FleetListSubscriptionsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *FleetListSubscriptionsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListSubscriptionsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListSubscriptionsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListSubscriptionsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSubscriptions

`func (o *FleetListSubscriptionsResult) GetSubscriptions() []FleetDescribeSubscriptionResult`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FleetListSubscriptionsResult) GetSubscriptionsOk() (*[]FleetDescribeSubscriptionResult, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FleetListSubscriptionsResult) SetSubscriptions(v []FleetDescribeSubscriptionResult)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FleetListSubscriptionsResult) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


