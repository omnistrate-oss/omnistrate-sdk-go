# ListServiceProviderEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]ServiceProviderEvent**](ServiceProviderEvent.md) | List of events | 
**EventsSummary** | [**ServiceProviderEventSummary**](ServiceProviderEventSummary.md) |  | 

## Methods

### NewListServiceProviderEventsResult

`func NewListServiceProviderEventsResult(events []ServiceProviderEvent, eventsSummary ServiceProviderEventSummary, ) *ListServiceProviderEventsResult`

NewListServiceProviderEventsResult instantiates a new ListServiceProviderEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceProviderEventsResultWithDefaults

`func NewListServiceProviderEventsResultWithDefaults() *ListServiceProviderEventsResult`

NewListServiceProviderEventsResultWithDefaults instantiates a new ListServiceProviderEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListServiceProviderEventsResult) GetEvents() []ServiceProviderEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListServiceProviderEventsResult) GetEventsOk() (*[]ServiceProviderEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListServiceProviderEventsResult) SetEvents(v []ServiceProviderEvent)`

SetEvents sets Events field to given value.


### GetEventsSummary

`func (o *ListServiceProviderEventsResult) GetEventsSummary() ServiceProviderEventSummary`

GetEventsSummary returns the EventsSummary field if non-nil, zero value otherwise.

### GetEventsSummaryOk

`func (o *ListServiceProviderEventsResult) GetEventsSummaryOk() (*ServiceProviderEventSummary, bool)`

GetEventsSummaryOk returns a tuple with the EventsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSummary

`func (o *ListServiceProviderEventsResult) SetEventsSummary(v ServiceProviderEventSummary)`

SetEventsSummary sets EventsSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


