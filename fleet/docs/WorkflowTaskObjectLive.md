# WorkflowTaskObjectLive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**ConditionDetail** | Pointer to **string** | What the evaluation observed, at most 512 bytes; never returned for a Secret | [optional] 
**ConditionResult** | Pointer to **string** | The task&#39;s condition evaluated against the live object now: success, failure, not-yet-met or error | [optional] 
**Conditions** | Pointer to [**[]InfrastructureCondition**](InfrastructureCondition.md) | At most three, messages capped | [optional] 
**ConditionsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Generation** | Pointer to **int64** |  | [optional] 
**LiveUid** | Pointer to **string** |  | [optional] 
**ObservedGeneration** | Pointer to **int64** | The generation the operator has observed | [optional] 
**Replaced** | Pointer to **bool** | The live UID differs from the one the task applied | [optional] 

## Methods

### NewWorkflowTaskObjectLive

`func NewWorkflowTaskObjectLive() *WorkflowTaskObjectLive`

NewWorkflowTaskObjectLive instantiates a new WorkflowTaskObjectLive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskObjectLiveWithDefaults

`func NewWorkflowTaskObjectLiveWithDefaults() *WorkflowTaskObjectLive`

NewWorkflowTaskObjectLiveWithDefaults instantiates a new WorkflowTaskObjectLive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *WorkflowTaskObjectLive) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *WorkflowTaskObjectLive) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *WorkflowTaskObjectLive) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *WorkflowTaskObjectLive) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetConditionDetail

`func (o *WorkflowTaskObjectLive) GetConditionDetail() string`

GetConditionDetail returns the ConditionDetail field if non-nil, zero value otherwise.

### GetConditionDetailOk

`func (o *WorkflowTaskObjectLive) GetConditionDetailOk() (*string, bool)`

GetConditionDetailOk returns a tuple with the ConditionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionDetail

`func (o *WorkflowTaskObjectLive) SetConditionDetail(v string)`

SetConditionDetail sets ConditionDetail field to given value.

### HasConditionDetail

`func (o *WorkflowTaskObjectLive) HasConditionDetail() bool`

HasConditionDetail returns a boolean if a field has been set.

### GetConditionResult

`func (o *WorkflowTaskObjectLive) GetConditionResult() string`

GetConditionResult returns the ConditionResult field if non-nil, zero value otherwise.

### GetConditionResultOk

`func (o *WorkflowTaskObjectLive) GetConditionResultOk() (*string, bool)`

GetConditionResultOk returns a tuple with the ConditionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionResult

`func (o *WorkflowTaskObjectLive) SetConditionResult(v string)`

SetConditionResult sets ConditionResult field to given value.

### HasConditionResult

`func (o *WorkflowTaskObjectLive) HasConditionResult() bool`

HasConditionResult returns a boolean if a field has been set.

### GetConditions

`func (o *WorkflowTaskObjectLive) GetConditions() []InfrastructureCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowTaskObjectLive) GetConditionsOk() (*[]InfrastructureCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowTaskObjectLive) SetConditions(v []InfrastructureCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowTaskObjectLive) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConditionsTruncation

`func (o *WorkflowTaskObjectLive) GetConditionsTruncation() CheckpointTruncation`

GetConditionsTruncation returns the ConditionsTruncation field if non-nil, zero value otherwise.

### GetConditionsTruncationOk

`func (o *WorkflowTaskObjectLive) GetConditionsTruncationOk() (*CheckpointTruncation, bool)`

GetConditionsTruncationOk returns a tuple with the ConditionsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsTruncation

`func (o *WorkflowTaskObjectLive) SetConditionsTruncation(v CheckpointTruncation)`

SetConditionsTruncation sets ConditionsTruncation field to given value.

### HasConditionsTruncation

`func (o *WorkflowTaskObjectLive) HasConditionsTruncation() bool`

HasConditionsTruncation returns a boolean if a field has been set.

### GetGeneration

`func (o *WorkflowTaskObjectLive) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *WorkflowTaskObjectLive) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *WorkflowTaskObjectLive) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *WorkflowTaskObjectLive) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetLiveUid

`func (o *WorkflowTaskObjectLive) GetLiveUid() string`

GetLiveUid returns the LiveUid field if non-nil, zero value otherwise.

### GetLiveUidOk

`func (o *WorkflowTaskObjectLive) GetLiveUidOk() (*string, bool)`

GetLiveUidOk returns a tuple with the LiveUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUid

`func (o *WorkflowTaskObjectLive) SetLiveUid(v string)`

SetLiveUid sets LiveUid field to given value.

### HasLiveUid

`func (o *WorkflowTaskObjectLive) HasLiveUid() bool`

HasLiveUid returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *WorkflowTaskObjectLive) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *WorkflowTaskObjectLive) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *WorkflowTaskObjectLive) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *WorkflowTaskObjectLive) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReplaced

`func (o *WorkflowTaskObjectLive) GetReplaced() bool`

GetReplaced returns the Replaced field if non-nil, zero value otherwise.

### GetReplacedOk

`func (o *WorkflowTaskObjectLive) GetReplacedOk() (*bool, bool)`

GetReplacedOk returns a tuple with the Replaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaced

`func (o *WorkflowTaskObjectLive) SetReplaced(v bool)`

SetReplaced sets Replaced field to given value.

### HasReplaced

`func (o *WorkflowTaskObjectLive) HasReplaced() bool`

HasReplaced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


