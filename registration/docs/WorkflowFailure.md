# WorkflowFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTime** | **string** | The time of the event | 
**Message** | **string** | Details of the event | 

## Methods

### NewWorkflowFailure

`func NewWorkflowFailure(eventTime string, message string, ) *WorkflowFailure`

NewWorkflowFailure instantiates a new WorkflowFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFailureWithDefaults

`func NewWorkflowFailureWithDefaults() *WorkflowFailure`

NewWorkflowFailureWithDefaults instantiates a new WorkflowFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTime

`func (o *WorkflowFailure) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *WorkflowFailure) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *WorkflowFailure) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetMessage

`func (o *WorkflowFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowFailure) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


