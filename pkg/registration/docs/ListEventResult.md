# ListEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]DescribeEventResult**](DescribeEventResult.md) | The list of events | [optional] 
**Ids** | **[]string** | The list of event IDs | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 

## Methods

### NewListEventResult

`func NewListEventResult(ids []string, ) *ListEventResult`

NewListEventResult instantiates a new ListEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventResultWithDefaults

`func NewListEventResultWithDefaults() *ListEventResult`

NewListEventResultWithDefaults instantiates a new ListEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListEventResult) GetEvents() []DescribeEventResult`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEventResult) GetEventsOk() (*[]DescribeEventResult, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEventResult) SetEvents(v []DescribeEventResult)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListEventResult) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIds

`func (o *ListEventResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListEventResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListEventResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListEventResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListEventResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListEventResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListEventResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


