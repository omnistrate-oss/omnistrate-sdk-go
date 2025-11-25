# DeploymentCellWorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | The error message if the event represents a failure | [optional] 
**EventTime** | **string** | The time the event occurred in RFC3339 format | 
**EventType** | **string** | The type of the workflow event | 
**Message** | **string** | The event message | 

## Methods

### NewDeploymentCellWorkflowEvent

`func NewDeploymentCellWorkflowEvent(eventTime string, eventType string, message string, ) *DeploymentCellWorkflowEvent`

NewDeploymentCellWorkflowEvent instantiates a new DeploymentCellWorkflowEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellWorkflowEventWithDefaults

`func NewDeploymentCellWorkflowEventWithDefaults() *DeploymentCellWorkflowEvent`

NewDeploymentCellWorkflowEventWithDefaults instantiates a new DeploymentCellWorkflowEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeploymentCellWorkflowEvent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeploymentCellWorkflowEvent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeploymentCellWorkflowEvent) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeploymentCellWorkflowEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEventTime

`func (o *DeploymentCellWorkflowEvent) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *DeploymentCellWorkflowEvent) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *DeploymentCellWorkflowEvent) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetEventType

`func (o *DeploymentCellWorkflowEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DeploymentCellWorkflowEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DeploymentCellWorkflowEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMessage

`func (o *DeploymentCellWorkflowEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentCellWorkflowEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentCellWorkflowEvent) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


