# WorkflowTaskWorkloadDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockingPods** | Pointer to [**[]PodProblem**](PodProblem.md) | The pods currently blocking the workload from becoming ready. | [optional] 
**DesiredReplicas** | Pointer to **int64** | The desired number of replicas for the workload. | [optional] 
**Kind** | Pointer to **string** | The workload kind, e.g. Deployment, StatefulSet, Job. | [optional] 
**ReadyReplicas** | Pointer to **int64** | The number of ready replicas for the workload. | [optional] 
**WarningEvents** | Pointer to [**[]WarningEvent**](WarningEvent.md) | Warning events observed for the workload. | [optional] 

## Methods

### NewWorkflowTaskWorkloadDetail

`func NewWorkflowTaskWorkloadDetail() *WorkflowTaskWorkloadDetail`

NewWorkflowTaskWorkloadDetail instantiates a new WorkflowTaskWorkloadDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskWorkloadDetailWithDefaults

`func NewWorkflowTaskWorkloadDetailWithDefaults() *WorkflowTaskWorkloadDetail`

NewWorkflowTaskWorkloadDetailWithDefaults instantiates a new WorkflowTaskWorkloadDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockingPods

`func (o *WorkflowTaskWorkloadDetail) GetBlockingPods() []PodProblem`

GetBlockingPods returns the BlockingPods field if non-nil, zero value otherwise.

### GetBlockingPodsOk

`func (o *WorkflowTaskWorkloadDetail) GetBlockingPodsOk() (*[]PodProblem, bool)`

GetBlockingPodsOk returns a tuple with the BlockingPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingPods

`func (o *WorkflowTaskWorkloadDetail) SetBlockingPods(v []PodProblem)`

SetBlockingPods sets BlockingPods field to given value.

### HasBlockingPods

`func (o *WorkflowTaskWorkloadDetail) HasBlockingPods() bool`

HasBlockingPods returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *WorkflowTaskWorkloadDetail) GetDesiredReplicas() int64`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *WorkflowTaskWorkloadDetail) GetDesiredReplicasOk() (*int64, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *WorkflowTaskWorkloadDetail) SetDesiredReplicas(v int64)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *WorkflowTaskWorkloadDetail) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetKind

`func (o *WorkflowTaskWorkloadDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkflowTaskWorkloadDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkflowTaskWorkloadDetail) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkflowTaskWorkloadDetail) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *WorkflowTaskWorkloadDetail) GetReadyReplicas() int64`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *WorkflowTaskWorkloadDetail) GetReadyReplicasOk() (*int64, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *WorkflowTaskWorkloadDetail) SetReadyReplicas(v int64)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *WorkflowTaskWorkloadDetail) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetWarningEvents

`func (o *WorkflowTaskWorkloadDetail) GetWarningEvents() []WarningEvent`

GetWarningEvents returns the WarningEvents field if non-nil, zero value otherwise.

### GetWarningEventsOk

`func (o *WorkflowTaskWorkloadDetail) GetWarningEventsOk() (*[]WarningEvent, bool)`

GetWarningEventsOk returns a tuple with the WarningEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningEvents

`func (o *WorkflowTaskWorkloadDetail) SetWarningEvents(v []WarningEvent)`

SetWarningEvents sets WarningEvents field to given value.

### HasWarningEvents

`func (o *WorkflowTaskWorkloadDetail) HasWarningEvents() bool`

HasWarningEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


