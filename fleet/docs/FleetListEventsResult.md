# FleetListEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]FleetDescribeEventResult**](FleetDescribeEventResult.md) | The list of events | 
**Ids** | **[]string** | The list of event IDs | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 

## Methods

### NewFleetListEventsResult

`func NewFleetListEventsResult(events []FleetDescribeEventResult, ids []string, ) *FleetListEventsResult`

NewFleetListEventsResult instantiates a new FleetListEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListEventsResultWithDefaults

`func NewFleetListEventsResultWithDefaults() *FleetListEventsResult`

NewFleetListEventsResultWithDefaults instantiates a new FleetListEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FleetListEventsResult) GetEvents() []FleetDescribeEventResult`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FleetListEventsResult) GetEventsOk() (*[]FleetDescribeEventResult, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FleetListEventsResult) SetEvents(v []FleetDescribeEventResult)`

SetEvents sets Events field to given value.


### GetIds

`func (o *FleetListEventsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *FleetListEventsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *FleetListEventsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *FleetListEventsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListEventsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListEventsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListEventsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


