# WorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action the task performed, e.g. apply|patch|delete|get (CRD), install|upgrade (helm). | [optional] 
**AttemptCount** | Pointer to **int64** | Consecutive attempts observed for this task, when the event corresponds to a retrying task. | [optional] 
**DisplayMessage** | Pointer to **string** | A concise human-readable summary derived from the error code, when present. | [optional] 
**ErrorCode** | Pointer to **string** | Stable error code from the workflow error taxonomy, present on failure events. | [optional] 
**EventTime** | **string** | Time of the event | 
**EventType** | **string** | The type of the workflow event | 
**FirstSeenAt** | Pointer to **string** | When this error signature was first observed for the current task attempt, RFC3339. | [optional] 
**GateExpression** | Pointer to **string** | The success condition gating task completion, when present. | [optional] 
**GateLastObserved** | Pointer to **string** | The last observed value of the gate expression. | [optional] 
**InfraDetail** | Pointer to [**WorkflowTaskInfraDetail**](WorkflowTaskInfraDetail.md) |  | [optional] 
**Message** | **string** | Details of the event | 
**NextRetryAt** | Pointer to **string** | The time of the next scheduled retry, RFC3339, when the task is awaiting retry. | [optional] 
**ResourceType** | Pointer to **string** | operatorCRD|genericCRD|helm|terraform|workload|cloudInfra|job|infraStack, when the event corresponds to a task. | [optional] 
**State** | Pointer to **string** | Live task lifecycle state for step/task events: Pending|Applying|AwaitingCondition|DriftMismatch|Failed|Succeeded. | [optional] 
**WorkloadDetail** | Pointer to [**WorkflowTaskWorkloadDetail**](WorkflowTaskWorkloadDetail.md) |  | [optional] 

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

### GetAction

`func (o *WorkflowEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAttemptCount

`func (o *WorkflowEvent) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *WorkflowEvent) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *WorkflowEvent) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *WorkflowEvent) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetDisplayMessage

`func (o *WorkflowEvent) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *WorkflowEvent) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *WorkflowEvent) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *WorkflowEvent) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *WorkflowEvent) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *WorkflowEvent) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *WorkflowEvent) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *WorkflowEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

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


### GetFirstSeenAt

`func (o *WorkflowEvent) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *WorkflowEvent) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *WorkflowEvent) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *WorkflowEvent) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetGateExpression

`func (o *WorkflowEvent) GetGateExpression() string`

GetGateExpression returns the GateExpression field if non-nil, zero value otherwise.

### GetGateExpressionOk

`func (o *WorkflowEvent) GetGateExpressionOk() (*string, bool)`

GetGateExpressionOk returns a tuple with the GateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateExpression

`func (o *WorkflowEvent) SetGateExpression(v string)`

SetGateExpression sets GateExpression field to given value.

### HasGateExpression

`func (o *WorkflowEvent) HasGateExpression() bool`

HasGateExpression returns a boolean if a field has been set.

### GetGateLastObserved

`func (o *WorkflowEvent) GetGateLastObserved() string`

GetGateLastObserved returns the GateLastObserved field if non-nil, zero value otherwise.

### GetGateLastObservedOk

`func (o *WorkflowEvent) GetGateLastObservedOk() (*string, bool)`

GetGateLastObservedOk returns a tuple with the GateLastObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateLastObserved

`func (o *WorkflowEvent) SetGateLastObserved(v string)`

SetGateLastObserved sets GateLastObserved field to given value.

### HasGateLastObserved

`func (o *WorkflowEvent) HasGateLastObserved() bool`

HasGateLastObserved returns a boolean if a field has been set.

### GetInfraDetail

`func (o *WorkflowEvent) GetInfraDetail() WorkflowTaskInfraDetail`

GetInfraDetail returns the InfraDetail field if non-nil, zero value otherwise.

### GetInfraDetailOk

`func (o *WorkflowEvent) GetInfraDetailOk() (*WorkflowTaskInfraDetail, bool)`

GetInfraDetailOk returns a tuple with the InfraDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDetail

`func (o *WorkflowEvent) SetInfraDetail(v WorkflowTaskInfraDetail)`

SetInfraDetail sets InfraDetail field to given value.

### HasInfraDetail

`func (o *WorkflowEvent) HasInfraDetail() bool`

HasInfraDetail returns a boolean if a field has been set.

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


### GetNextRetryAt

`func (o *WorkflowEvent) GetNextRetryAt() string`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *WorkflowEvent) GetNextRetryAtOk() (*string, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *WorkflowEvent) SetNextRetryAt(v string)`

SetNextRetryAt sets NextRetryAt field to given value.

### HasNextRetryAt

`func (o *WorkflowEvent) HasNextRetryAt() bool`

HasNextRetryAt returns a boolean if a field has been set.

### GetResourceType

`func (o *WorkflowEvent) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WorkflowEvent) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WorkflowEvent) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WorkflowEvent) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetState

`func (o *WorkflowEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkflowEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkloadDetail

`func (o *WorkflowEvent) GetWorkloadDetail() WorkflowTaskWorkloadDetail`

GetWorkloadDetail returns the WorkloadDetail field if non-nil, zero value otherwise.

### GetWorkloadDetailOk

`func (o *WorkflowEvent) GetWorkloadDetailOk() (*WorkflowTaskWorkloadDetail, bool)`

GetWorkloadDetailOk returns a tuple with the WorkloadDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDetail

`func (o *WorkflowEvent) SetWorkloadDetail(v WorkflowTaskWorkloadDetail)`

SetWorkloadDetail sets WorkloadDetail field to given value.

### HasWorkloadDetail

`func (o *WorkflowEvent) HasWorkloadDetail() bool`

HasWorkloadDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


