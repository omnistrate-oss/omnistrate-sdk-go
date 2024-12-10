# ListSubscriptionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of subscription IDs | 
**NextPageToken** | Pointer to **string** | The next token to use when listing subscriptions | [optional] 
**Subscriptions** | Pointer to [**[]DescribeSubscriptionResult**](DescribeSubscriptionResult.md) | List of subscriptions | [optional] 

## Methods

### NewListSubscriptionsResult

`func NewListSubscriptionsResult(ids []string, ) *ListSubscriptionsResult`

NewListSubscriptionsResult instantiates a new ListSubscriptionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionsResultWithDefaults

`func NewListSubscriptionsResultWithDefaults() *ListSubscriptionsResult`

NewListSubscriptionsResultWithDefaults instantiates a new ListSubscriptionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListSubscriptionsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListSubscriptionsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListSubscriptionsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListSubscriptionsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListSubscriptionsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListSubscriptionsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListSubscriptionsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ListSubscriptionsResult) GetSubscriptions() []DescribeSubscriptionResult`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ListSubscriptionsResult) GetSubscriptionsOk() (*[]DescribeSubscriptionResult, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ListSubscriptionsResult) SetSubscriptions(v []DescribeSubscriptionResult)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ListSubscriptionsResult) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


