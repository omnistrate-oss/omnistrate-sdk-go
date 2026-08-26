# DeploymentCellWorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action the task performed, e.g. apply|patch|delete|get (CRD), install|upgrade (helm). | [optional] 
**AttemptCount** | Pointer to **int64** | Consecutive attempts observed for this task, when the event corresponds to a retrying task. | [optional] 
**DisplayMessage** | Pointer to **string** | A concise human-readable summary derived from the error code, when present. | [optional] 
**Error** | Pointer to **string** | The error message if the event represents a failure | [optional] 
**ErrorCode** | Pointer to **string** | Stable error code from the workflow error taxonomy, present on failure events | [optional] 
**EventTime** | **string** | The time the event occurred in RFC3339 format | 
**EventType** | **string** | The type of the workflow event | 
**FirstSeenAt** | Pointer to **string** | When this error signature was first observed for the current task attempt, RFC3339. | [optional] 
**GateExpression** | Pointer to **string** | The success condition gating task completion, when present. | [optional] 
**GateLastObserved** | Pointer to **string** | The last observed value of the gate expression. | [optional] 
**InfraDetail** | Pointer to [**WorkflowTaskInfraDetail**](WorkflowTaskInfraDetail.md) |  | [optional] 
**Message** | **string** | The event message | 
**NextRetryAt** | Pointer to **string** | The time of the next scheduled retry, RFC3339, when the task is awaiting retry. | [optional] 
**ResourceType** | Pointer to **string** | operatorCRD|genericCRD|helm|terraform|workload|cloudInfra|job|infraStack, when the event corresponds to a task. | [optional] 
**State** | Pointer to **string** | Live task lifecycle state for step/task events: Pending|Applying|AwaitingCondition|DriftMismatch|Failed|Succeeded. | [optional] 
**WorkloadDetail** | Pointer to [**WorkflowTaskWorkloadDetail**](WorkflowTaskWorkloadDetail.md) |  | [optional] 

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

### GetAction

`func (o *DeploymentCellWorkflowEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentCellWorkflowEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentCellWorkflowEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentCellWorkflowEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAttemptCount

`func (o *DeploymentCellWorkflowEvent) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *DeploymentCellWorkflowEvent) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *DeploymentCellWorkflowEvent) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *DeploymentCellWorkflowEvent) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetDisplayMessage

`func (o *DeploymentCellWorkflowEvent) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *DeploymentCellWorkflowEvent) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *DeploymentCellWorkflowEvent) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *DeploymentCellWorkflowEvent) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

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

### GetErrorCode

`func (o *DeploymentCellWorkflowEvent) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DeploymentCellWorkflowEvent) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DeploymentCellWorkflowEvent) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DeploymentCellWorkflowEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

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


### GetFirstSeenAt

`func (o *DeploymentCellWorkflowEvent) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *DeploymentCellWorkflowEvent) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *DeploymentCellWorkflowEvent) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *DeploymentCellWorkflowEvent) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetGateExpression

`func (o *DeploymentCellWorkflowEvent) GetGateExpression() string`

GetGateExpression returns the GateExpression field if non-nil, zero value otherwise.

### GetGateExpressionOk

`func (o *DeploymentCellWorkflowEvent) GetGateExpressionOk() (*string, bool)`

GetGateExpressionOk returns a tuple with the GateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateExpression

`func (o *DeploymentCellWorkflowEvent) SetGateExpression(v string)`

SetGateExpression sets GateExpression field to given value.

### HasGateExpression

`func (o *DeploymentCellWorkflowEvent) HasGateExpression() bool`

HasGateExpression returns a boolean if a field has been set.

### GetGateLastObserved

`func (o *DeploymentCellWorkflowEvent) GetGateLastObserved() string`

GetGateLastObserved returns the GateLastObserved field if non-nil, zero value otherwise.

### GetGateLastObservedOk

`func (o *DeploymentCellWorkflowEvent) GetGateLastObservedOk() (*string, bool)`

GetGateLastObservedOk returns a tuple with the GateLastObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateLastObserved

`func (o *DeploymentCellWorkflowEvent) SetGateLastObserved(v string)`

SetGateLastObserved sets GateLastObserved field to given value.

### HasGateLastObserved

`func (o *DeploymentCellWorkflowEvent) HasGateLastObserved() bool`

HasGateLastObserved returns a boolean if a field has been set.

### GetInfraDetail

`func (o *DeploymentCellWorkflowEvent) GetInfraDetail() WorkflowTaskInfraDetail`

GetInfraDetail returns the InfraDetail field if non-nil, zero value otherwise.

### GetInfraDetailOk

`func (o *DeploymentCellWorkflowEvent) GetInfraDetailOk() (*WorkflowTaskInfraDetail, bool)`

GetInfraDetailOk returns a tuple with the InfraDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDetail

`func (o *DeploymentCellWorkflowEvent) SetInfraDetail(v WorkflowTaskInfraDetail)`

SetInfraDetail sets InfraDetail field to given value.

### HasInfraDetail

`func (o *DeploymentCellWorkflowEvent) HasInfraDetail() bool`

HasInfraDetail returns a boolean if a field has been set.

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


### GetNextRetryAt

`func (o *DeploymentCellWorkflowEvent) GetNextRetryAt() string`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *DeploymentCellWorkflowEvent) GetNextRetryAtOk() (*string, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *DeploymentCellWorkflowEvent) SetNextRetryAt(v string)`

SetNextRetryAt sets NextRetryAt field to given value.

### HasNextRetryAt

`func (o *DeploymentCellWorkflowEvent) HasNextRetryAt() bool`

HasNextRetryAt returns a boolean if a field has been set.

### GetResourceType

`func (o *DeploymentCellWorkflowEvent) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DeploymentCellWorkflowEvent) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DeploymentCellWorkflowEvent) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DeploymentCellWorkflowEvent) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetState

`func (o *DeploymentCellWorkflowEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentCellWorkflowEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentCellWorkflowEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentCellWorkflowEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkloadDetail

`func (o *DeploymentCellWorkflowEvent) GetWorkloadDetail() WorkflowTaskWorkloadDetail`

GetWorkloadDetail returns the WorkloadDetail field if non-nil, zero value otherwise.

### GetWorkloadDetailOk

`func (o *DeploymentCellWorkflowEvent) GetWorkloadDetailOk() (*WorkflowTaskWorkloadDetail, bool)`

GetWorkloadDetailOk returns a tuple with the WorkloadDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDetail

`func (o *DeploymentCellWorkflowEvent) SetWorkloadDetail(v WorkflowTaskWorkloadDetail)`

SetWorkloadDetail sets WorkloadDetail field to given value.

### HasWorkloadDetail

`func (o *DeploymentCellWorkflowEvent) HasWorkloadDetail() bool`

HasWorkloadDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


