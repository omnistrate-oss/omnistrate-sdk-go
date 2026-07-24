# WorkflowExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | The time the execution ended, in RFC3339 format. | [optional] 
**ExecutionId** | **string** | The unique ID of this workflow execution. | 
**FailureCode** | Pointer to **string** | Stable error code from the workflow error taxonomy, present when the execution failed. | [optional] 
**FailureMessage** | Pointer to **string** | The failure detail, present when the execution failed. | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 
**InvokedBy** | Pointer to **string** | The identity that invoked this workflow execution. | [optional] 
**Outputs** | Pointer to **map[string]interface{}** | The typed output parameters produced by this execution, keyed by parameter name. | [optional] 
**RequestParams** | Pointer to **map[string]interface{}** | The input parameters supplied for this execution. | [optional] 
**StartTime** | Pointer to **string** | The time the execution started, in RFC3339 format. | [optional] 
**Status** | **string** | PENDING|RUNNING|SUCCEEDED|FAILED | 
**Tasks** | Pointer to [**[]WorkflowTask**](WorkflowTask.md) | The live per-task state of the execution&#39;s authored DAG tasks, when available. | [optional] 
**Verb** | **string** | The provider-defined verb that addressed this workflow. | 
**WorkflowId** | **string** | ID of a Custom Workflow | 

## Methods

### NewWorkflowExecution

`func NewWorkflowExecution(executionId string, instanceId string, status string, verb string, workflowId string, ) *WorkflowExecution`

NewWorkflowExecution instantiates a new WorkflowExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionWithDefaults

`func NewWorkflowExecutionWithDefaults() *WorkflowExecution`

NewWorkflowExecutionWithDefaults instantiates a new WorkflowExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *WorkflowExecution) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowExecution) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowExecution) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowExecution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExecutionId

`func (o *WorkflowExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetFailureCode

`func (o *WorkflowExecution) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *WorkflowExecution) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *WorkflowExecution) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *WorkflowExecution) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureMessage

`func (o *WorkflowExecution) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *WorkflowExecution) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *WorkflowExecution) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *WorkflowExecution) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetInstanceId

`func (o *WorkflowExecution) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *WorkflowExecution) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *WorkflowExecution) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInvokedBy

`func (o *WorkflowExecution) GetInvokedBy() string`

GetInvokedBy returns the InvokedBy field if non-nil, zero value otherwise.

### GetInvokedByOk

`func (o *WorkflowExecution) GetInvokedByOk() (*string, bool)`

GetInvokedByOk returns a tuple with the InvokedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokedBy

`func (o *WorkflowExecution) SetInvokedBy(v string)`

SetInvokedBy sets InvokedBy field to given value.

### HasInvokedBy

`func (o *WorkflowExecution) HasInvokedBy() bool`

HasInvokedBy returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowExecution) GetOutputs() map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowExecution) GetOutputsOk() (*map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowExecution) SetOutputs(v map[string]interface{})`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowExecution) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetRequestParams

`func (o *WorkflowExecution) GetRequestParams() map[string]interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *WorkflowExecution) GetRequestParamsOk() (*map[string]interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *WorkflowExecution) SetRequestParams(v map[string]interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *WorkflowExecution) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowExecution) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowExecution) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowExecution) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowExecution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecution) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTasks

`func (o *WorkflowExecution) GetTasks() []WorkflowTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowExecution) GetTasksOk() (*[]WorkflowTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowExecution) SetTasks(v []WorkflowTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowExecution) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetVerb

`func (o *WorkflowExecution) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *WorkflowExecution) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *WorkflowExecution) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetWorkflowId

`func (o *WorkflowExecution) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowExecution) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowExecution) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


