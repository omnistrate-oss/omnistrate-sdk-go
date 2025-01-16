# WorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTime** | **string** | Time of the event | 
**EventType** | **string** | The type of the workflow event | 
**Message** | **string** | Details of the event | 

## Methods

### NewWorkflowEvent

`func NewWorkflowEvent(eventTime string, eventType string, message string, ) *WorkflowEvent`

NewWorkflowEvent instantiates a new WorkflowEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEventWithDefaults

`func NewWorkflowEventWithDefaults() *WorkflowEvent`

NewWorkflowEventWithDefaults instantiates a new WorkflowEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTime

`func (o *WorkflowEvent) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *WorkflowEvent) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *WorkflowEvent) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetEventType

`func (o *WorkflowEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMessage

`func (o *WorkflowEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowEvent) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


