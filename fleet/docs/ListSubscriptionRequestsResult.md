# ListSubscriptionRequestsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of subscription request IDs | 
**NextPageToken** | Pointer to **string** | The next token to use when listing subscription requests | [optional] 
**SubscriptionRequests** | Pointer to [**[]DescribeSubscriptionRequestResult**](DescribeSubscriptionRequestResult.md) | List of subscription requests | [optional] 

## Methods

### NewListSubscriptionRequestsResult

`func NewListSubscriptionRequestsResult(ids []string, ) *ListSubscriptionRequestsResult`

NewListSubscriptionRequestsResult instantiates a new ListSubscriptionRequestsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionRequestsResultWithDefaults

`func NewListSubscriptionRequestsResultWithDefaults() *ListSubscriptionRequestsResult`

NewListSubscriptionRequestsResultWithDefaults instantiates a new ListSubscriptionRequestsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListSubscriptionRequestsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListSubscriptionRequestsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListSubscriptionRequestsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListSubscriptionRequestsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListSubscriptionRequestsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListSubscriptionRequestsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListSubscriptionRequestsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSubscriptionRequests

`func (o *ListSubscriptionRequestsResult) GetSubscriptionRequests() []DescribeSubscriptionRequestResult`

GetSubscriptionRequests returns the SubscriptionRequests field if non-nil, zero value otherwise.

### GetSubscriptionRequestsOk

`func (o *ListSubscriptionRequestsResult) GetSubscriptionRequestsOk() (*[]DescribeSubscriptionRequestResult, bool)`

GetSubscriptionRequestsOk returns a tuple with the SubscriptionRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRequests

`func (o *ListSubscriptionRequestsResult) SetSubscriptionRequests(v []DescribeSubscriptionRequestResult)`

SetSubscriptionRequests sets SubscriptionRequests field to given value.

### HasSubscriptionRequests

`func (o *ListSubscriptionRequestsResult) HasSubscriptionRequests() bool`

HasSubscriptionRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


