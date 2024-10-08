# ListEndCustomerEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EndCustomerEvent**](EndCustomerEvent.md) | List of events | 

## Methods

### NewListEndCustomerEventsResult

`func NewListEndCustomerEventsResult(events []EndCustomerEvent, ) *ListEndCustomerEventsResult`

NewListEndCustomerEventsResult instantiates a new ListEndCustomerEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEndCustomerEventsResultWithDefaults

`func NewListEndCustomerEventsResultWithDefaults() *ListEndCustomerEventsResult`

NewListEndCustomerEventsResultWithDefaults instantiates a new ListEndCustomerEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListEndCustomerEventsResult) GetEvents() []EndCustomerEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEndCustomerEventsResult) GetEventsOk() (*[]EndCustomerEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEndCustomerEventsResult) SetEvents(v []EndCustomerEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


