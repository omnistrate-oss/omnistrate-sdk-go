# EventsPerWorkflowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]WorkflowEvent**](WorkflowEvent.md) | List of events for the step | 
**StepName** | **string** | Name of the workflow step | 

## Methods

### NewEventsPerWorkflowStep

`func NewEventsPerWorkflowStep(events []WorkflowEvent, stepName string, ) *EventsPerWorkflowStep`

NewEventsPerWorkflowStep instantiates a new EventsPerWorkflowStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPerWorkflowStepWithDefaults

`func NewEventsPerWorkflowStepWithDefaults() *EventsPerWorkflowStep`

NewEventsPerWorkflowStepWithDefaults instantiates a new EventsPerWorkflowStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsPerWorkflowStep) GetEvents() []WorkflowEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsPerWorkflowStep) GetEventsOk() (*[]WorkflowEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsPerWorkflowStep) SetEvents(v []WorkflowEvent)`

SetEvents sets Events field to given value.


### GetStepName

`func (o *EventsPerWorkflowStep) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *EventsPerWorkflowStep) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *EventsPerWorkflowStep) SetStepName(v string)`

SetStepName sets StepName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


