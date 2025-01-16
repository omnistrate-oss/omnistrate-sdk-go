# FleetAuditEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]FleetAuditEvent**](FleetAuditEvent.md) | The list of events | 
**Ids** | **[]string** | The list of event IDs | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 

## Methods

### NewFleetAuditEventsResult

`func NewFleetAuditEventsResult(events []FleetAuditEvent, ids []string, ) *FleetAuditEventsResult`

NewFleetAuditEventsResult instantiates a new FleetAuditEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAuditEventsResultWithDefaults

`func NewFleetAuditEventsResultWithDefaults() *FleetAuditEventsResult`

NewFleetAuditEventsResultWithDefaults instantiates a new FleetAuditEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FleetAuditEventsResult) GetEvents() []FleetAuditEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FleetAuditEventsResult) GetEventsOk() (*[]FleetAuditEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FleetAuditEventsResult) SetEvents(v []FleetAuditEvent)`

SetEvents sets Events field to given value.


### GetIds

`func (o *FleetAuditEventsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *FleetAuditEventsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *FleetAuditEventsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *FleetAuditEventsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetAuditEventsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetAuditEventsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetAuditEventsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


