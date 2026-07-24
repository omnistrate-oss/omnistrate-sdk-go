# WorkflowTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action performed by the task, e.g. apply|patch|delete|get (CRD), install|upgrade (helm). | [optional] 
**AttemptCount** | Pointer to **int64** | The number of attempts made for this task. | [optional] 
**EndTime** | Pointer to **string** | The time the task ended, in RFC3339 format. | [optional] 
**GateExpression** | Pointer to **string** | The success condition gating task completion, when present. | [optional] 
**GateLastObserved** | Pointer to **string** | The last observed value of the gate expression. | [optional] 
**HelmDetail** | Pointer to [**WorkflowTaskHelmDetail**](WorkflowTaskHelmDetail.md) |  | [optional] 
**InfraDetail** | Pointer to [**WorkflowTaskInfraDetail**](WorkflowTaskInfraDetail.md) |  | [optional] 
**LastAttemptError** | Pointer to [**TaskAttemptError**](TaskAttemptError.md) |  | [optional] 
**Name** | **string** | The authored DAG task name; the infrastructure stack name for infraStack tasks. | 
**NextRetryAt** | Pointer to **string** | The time of the next scheduled retry, in RFC3339 format. | [optional] 
**ResourceType** | Pointer to **string** | operatorCRD|genericCRD|helm|terraform|workload|cloudInfra|job|infraStack | [optional] 
**StartTime** | Pointer to **string** | The time the task started, in RFC3339 format. | [optional] 
**State** | **string** | Pending|Applying|AwaitingCondition|DriftMismatch|Failed|Succeeded | 
**TerraformDetail** | Pointer to [**WorkflowTaskTerraformDetail**](WorkflowTaskTerraformDetail.md) |  | [optional] 
**WorkloadDetail** | Pointer to [**WorkflowTaskWorkloadDetail**](WorkflowTaskWorkloadDetail.md) |  | [optional] 

## Methods

### NewWorkflowTask

`func NewWorkflowTask(name string, state string, ) *WorkflowTask`

NewWorkflowTask instantiates a new WorkflowTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskWithDefaults

`func NewWorkflowTaskWithDefaults() *WorkflowTask`

NewWorkflowTaskWithDefaults instantiates a new WorkflowTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowTask) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowTask) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowTask) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowTask) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAttemptCount

`func (o *WorkflowTask) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *WorkflowTask) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *WorkflowTask) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *WorkflowTask) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowTask) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowTask) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowTask) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowTask) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGateExpression

`func (o *WorkflowTask) GetGateExpression() string`

GetGateExpression returns the GateExpression field if non-nil, zero value otherwise.

### GetGateExpressionOk

`func (o *WorkflowTask) GetGateExpressionOk() (*string, bool)`

GetGateExpressionOk returns a tuple with the GateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateExpression

`func (o *WorkflowTask) SetGateExpression(v string)`

SetGateExpression sets GateExpression field to given value.

### HasGateExpression

`func (o *WorkflowTask) HasGateExpression() bool`

HasGateExpression returns a boolean if a field has been set.

### GetGateLastObserved

`func (o *WorkflowTask) GetGateLastObserved() string`

GetGateLastObserved returns the GateLastObserved field if non-nil, zero value otherwise.

### GetGateLastObservedOk

`func (o *WorkflowTask) GetGateLastObservedOk() (*string, bool)`

GetGateLastObservedOk returns a tuple with the GateLastObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateLastObserved

`func (o *WorkflowTask) SetGateLastObserved(v string)`

SetGateLastObserved sets GateLastObserved field to given value.

### HasGateLastObserved

`func (o *WorkflowTask) HasGateLastObserved() bool`

HasGateLastObserved returns a boolean if a field has been set.

### GetHelmDetail

`func (o *WorkflowTask) GetHelmDetail() WorkflowTaskHelmDetail`

GetHelmDetail returns the HelmDetail field if non-nil, zero value otherwise.

### GetHelmDetailOk

`func (o *WorkflowTask) GetHelmDetailOk() (*WorkflowTaskHelmDetail, bool)`

GetHelmDetailOk returns a tuple with the HelmDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmDetail

`func (o *WorkflowTask) SetHelmDetail(v WorkflowTaskHelmDetail)`

SetHelmDetail sets HelmDetail field to given value.

### HasHelmDetail

`func (o *WorkflowTask) HasHelmDetail() bool`

HasHelmDetail returns a boolean if a field has been set.

### GetInfraDetail

`func (o *WorkflowTask) GetInfraDetail() WorkflowTaskInfraDetail`

GetInfraDetail returns the InfraDetail field if non-nil, zero value otherwise.

### GetInfraDetailOk

`func (o *WorkflowTask) GetInfraDetailOk() (*WorkflowTaskInfraDetail, bool)`

GetInfraDetailOk returns a tuple with the InfraDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDetail

`func (o *WorkflowTask) SetInfraDetail(v WorkflowTaskInfraDetail)`

SetInfraDetail sets InfraDetail field to given value.

### HasInfraDetail

`func (o *WorkflowTask) HasInfraDetail() bool`

HasInfraDetail returns a boolean if a field has been set.

### GetLastAttemptError

`func (o *WorkflowTask) GetLastAttemptError() TaskAttemptError`

GetLastAttemptError returns the LastAttemptError field if non-nil, zero value otherwise.

### GetLastAttemptErrorOk

`func (o *WorkflowTask) GetLastAttemptErrorOk() (*TaskAttemptError, bool)`

GetLastAttemptErrorOk returns a tuple with the LastAttemptError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptError

`func (o *WorkflowTask) SetLastAttemptError(v TaskAttemptError)`

SetLastAttemptError sets LastAttemptError field to given value.

### HasLastAttemptError

`func (o *WorkflowTask) HasLastAttemptError() bool`

HasLastAttemptError returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTask) SetName(v string)`

SetName sets Name field to given value.


### GetNextRetryAt

`func (o *WorkflowTask) GetNextRetryAt() string`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *WorkflowTask) GetNextRetryAtOk() (*string, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *WorkflowTask) SetNextRetryAt(v string)`

SetNextRetryAt sets NextRetryAt field to given value.

### HasNextRetryAt

`func (o *WorkflowTask) HasNextRetryAt() bool`

HasNextRetryAt returns a boolean if a field has been set.

### GetResourceType

`func (o *WorkflowTask) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WorkflowTask) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WorkflowTask) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WorkflowTask) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowTask) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowTask) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowTask) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowTask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *WorkflowTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowTask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowTask) SetState(v string)`

SetState sets State field to given value.


### GetTerraformDetail

`func (o *WorkflowTask) GetTerraformDetail() WorkflowTaskTerraformDetail`

GetTerraformDetail returns the TerraformDetail field if non-nil, zero value otherwise.

### GetTerraformDetailOk

`func (o *WorkflowTask) GetTerraformDetailOk() (*WorkflowTaskTerraformDetail, bool)`

GetTerraformDetailOk returns a tuple with the TerraformDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformDetail

`func (o *WorkflowTask) SetTerraformDetail(v WorkflowTaskTerraformDetail)`

SetTerraformDetail sets TerraformDetail field to given value.

### HasTerraformDetail

`func (o *WorkflowTask) HasTerraformDetail() bool`

HasTerraformDetail returns a boolean if a field has been set.

### GetWorkloadDetail

`func (o *WorkflowTask) GetWorkloadDetail() WorkflowTaskWorkloadDetail`

GetWorkloadDetail returns the WorkloadDetail field if non-nil, zero value otherwise.

### GetWorkloadDetailOk

`func (o *WorkflowTask) GetWorkloadDetailOk() (*WorkflowTaskWorkloadDetail, bool)`

GetWorkloadDetailOk returns a tuple with the WorkloadDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadDetail

`func (o *WorkflowTask) SetWorkloadDetail(v WorkflowTaskWorkloadDetail)`

SetWorkloadDetail sets WorkloadDetail field to given value.

### HasWorkloadDetail

`func (o *WorkflowTask) HasWorkloadDetail() bool`

HasWorkloadDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


