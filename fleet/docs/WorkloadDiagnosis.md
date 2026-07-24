# WorkloadDiagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockingPods** | Pointer to [**[]DebugPodProblem**](DebugPodProblem.md) | Pods preventing the workload from reaching its desired state. | [optional] 
**DesiredReplicas** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ReadyReplicas** | Pointer to **int64** |  | [optional] 
**WarningEvents** | Pointer to [**[]DebugWarningEvent**](DebugWarningEvent.md) | Recent Warning events, deduplicated by reason+message. | [optional] 

## Methods

### NewWorkloadDiagnosis

`func NewWorkloadDiagnosis() *WorkloadDiagnosis`

NewWorkloadDiagnosis instantiates a new WorkloadDiagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDiagnosisWithDefaults

`func NewWorkloadDiagnosisWithDefaults() *WorkloadDiagnosis`

NewWorkloadDiagnosisWithDefaults instantiates a new WorkloadDiagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockingPods

`func (o *WorkloadDiagnosis) GetBlockingPods() []DebugPodProblem`

GetBlockingPods returns the BlockingPods field if non-nil, zero value otherwise.

### GetBlockingPodsOk

`func (o *WorkloadDiagnosis) GetBlockingPodsOk() (*[]DebugPodProblem, bool)`

GetBlockingPodsOk returns a tuple with the BlockingPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingPods

`func (o *WorkloadDiagnosis) SetBlockingPods(v []DebugPodProblem)`

SetBlockingPods sets BlockingPods field to given value.

### HasBlockingPods

`func (o *WorkloadDiagnosis) HasBlockingPods() bool`

HasBlockingPods returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *WorkloadDiagnosis) GetDesiredReplicas() int64`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *WorkloadDiagnosis) GetDesiredReplicasOk() (*int64, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *WorkloadDiagnosis) SetDesiredReplicas(v int64)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *WorkloadDiagnosis) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadDiagnosis) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadDiagnosis) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadDiagnosis) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadDiagnosis) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *WorkloadDiagnosis) GetReadyReplicas() int64`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *WorkloadDiagnosis) GetReadyReplicasOk() (*int64, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *WorkloadDiagnosis) SetReadyReplicas(v int64)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *WorkloadDiagnosis) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetWarningEvents

`func (o *WorkloadDiagnosis) GetWarningEvents() []DebugWarningEvent`

GetWarningEvents returns the WarningEvents field if non-nil, zero value otherwise.

### GetWarningEventsOk

`func (o *WorkloadDiagnosis) GetWarningEventsOk() (*[]DebugWarningEvent, bool)`

GetWarningEventsOk returns a tuple with the WarningEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningEvents

`func (o *WorkloadDiagnosis) SetWarningEvents(v []DebugWarningEvent)`

SetWarningEvents sets WarningEvents field to given value.

### HasWarningEvents

`func (o *WorkloadDiagnosis) HasWarningEvents() bool`

HasWarningEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


