# WorkflowTaskHelmDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveHook** | Pointer to **string** | The Helm hook currently executing, when one is. Complements failedHook, which describes a hook that already failed. | [optional] 
**FailedHook** | Pointer to **string** | The Helm hook that failed, when applicable. | [optional] 
**FirstDeployedAt** | Pointer to **string** | The time the release was first deployed, in RFC3339 format, as recorded on the release. Observed evidence for the preparation checkpoint; absent when not recorded. | [optional] 
**HookPhase** | Pointer to **string** | The execution phase of activeHook as recorded on the release: Unknown|Running|Succeeded|Failed. A hook that reads Running with a hookStartedAt well in the past and a terminal release status indicates the process died before recording the hook&#39;s terminal phase; it is not evidence that the hook is still executing. | [optional] 
**HookStartedAt** | Pointer to **string** | The recorded start time of activeHook, in RFC3339 format. Absent when the release carries no start evidence for the hook; an absent time is never replaced by the current poll time or by an epoch value. | [optional] 
**LastDeployedAt** | Pointer to **string** | The time the release was last deployed, in RFC3339 format, as recorded on the release. Absent when not recorded. | [optional] 
**PendingResources** | Pointer to **[]string** | The resources from the release that are still pending. | [optional] 
**Phase** | Pointer to **string** | The Helm release phase. | [optional] 
**Release** | Pointer to **string** | The Helm release name. | [optional] 
**Revision** | Pointer to **int64** | The Helm release revision. | [optional] 

## Methods

### NewWorkflowTaskHelmDetail

`func NewWorkflowTaskHelmDetail() *WorkflowTaskHelmDetail`

NewWorkflowTaskHelmDetail instantiates a new WorkflowTaskHelmDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskHelmDetailWithDefaults

`func NewWorkflowTaskHelmDetailWithDefaults() *WorkflowTaskHelmDetail`

NewWorkflowTaskHelmDetailWithDefaults instantiates a new WorkflowTaskHelmDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveHook

`func (o *WorkflowTaskHelmDetail) GetActiveHook() string`

GetActiveHook returns the ActiveHook field if non-nil, zero value otherwise.

### GetActiveHookOk

`func (o *WorkflowTaskHelmDetail) GetActiveHookOk() (*string, bool)`

GetActiveHookOk returns a tuple with the ActiveHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveHook

`func (o *WorkflowTaskHelmDetail) SetActiveHook(v string)`

SetActiveHook sets ActiveHook field to given value.

### HasActiveHook

`func (o *WorkflowTaskHelmDetail) HasActiveHook() bool`

HasActiveHook returns a boolean if a field has been set.

### GetFailedHook

`func (o *WorkflowTaskHelmDetail) GetFailedHook() string`

GetFailedHook returns the FailedHook field if non-nil, zero value otherwise.

### GetFailedHookOk

`func (o *WorkflowTaskHelmDetail) GetFailedHookOk() (*string, bool)`

GetFailedHookOk returns a tuple with the FailedHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedHook

`func (o *WorkflowTaskHelmDetail) SetFailedHook(v string)`

SetFailedHook sets FailedHook field to given value.

### HasFailedHook

`func (o *WorkflowTaskHelmDetail) HasFailedHook() bool`

HasFailedHook returns a boolean if a field has been set.

### GetFirstDeployedAt

`func (o *WorkflowTaskHelmDetail) GetFirstDeployedAt() string`

GetFirstDeployedAt returns the FirstDeployedAt field if non-nil, zero value otherwise.

### GetFirstDeployedAtOk

`func (o *WorkflowTaskHelmDetail) GetFirstDeployedAtOk() (*string, bool)`

GetFirstDeployedAtOk returns a tuple with the FirstDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDeployedAt

`func (o *WorkflowTaskHelmDetail) SetFirstDeployedAt(v string)`

SetFirstDeployedAt sets FirstDeployedAt field to given value.

### HasFirstDeployedAt

`func (o *WorkflowTaskHelmDetail) HasFirstDeployedAt() bool`

HasFirstDeployedAt returns a boolean if a field has been set.

### GetHookPhase

`func (o *WorkflowTaskHelmDetail) GetHookPhase() string`

GetHookPhase returns the HookPhase field if non-nil, zero value otherwise.

### GetHookPhaseOk

`func (o *WorkflowTaskHelmDetail) GetHookPhaseOk() (*string, bool)`

GetHookPhaseOk returns a tuple with the HookPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookPhase

`func (o *WorkflowTaskHelmDetail) SetHookPhase(v string)`

SetHookPhase sets HookPhase field to given value.

### HasHookPhase

`func (o *WorkflowTaskHelmDetail) HasHookPhase() bool`

HasHookPhase returns a boolean if a field has been set.

### GetHookStartedAt

`func (o *WorkflowTaskHelmDetail) GetHookStartedAt() string`

GetHookStartedAt returns the HookStartedAt field if non-nil, zero value otherwise.

### GetHookStartedAtOk

`func (o *WorkflowTaskHelmDetail) GetHookStartedAtOk() (*string, bool)`

GetHookStartedAtOk returns a tuple with the HookStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookStartedAt

`func (o *WorkflowTaskHelmDetail) SetHookStartedAt(v string)`

SetHookStartedAt sets HookStartedAt field to given value.

### HasHookStartedAt

`func (o *WorkflowTaskHelmDetail) HasHookStartedAt() bool`

HasHookStartedAt returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *WorkflowTaskHelmDetail) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *WorkflowTaskHelmDetail) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *WorkflowTaskHelmDetail) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *WorkflowTaskHelmDetail) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.

### GetPendingResources

`func (o *WorkflowTaskHelmDetail) GetPendingResources() []string`

GetPendingResources returns the PendingResources field if non-nil, zero value otherwise.

### GetPendingResourcesOk

`func (o *WorkflowTaskHelmDetail) GetPendingResourcesOk() (*[]string, bool)`

GetPendingResourcesOk returns a tuple with the PendingResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingResources

`func (o *WorkflowTaskHelmDetail) SetPendingResources(v []string)`

SetPendingResources sets PendingResources field to given value.

### HasPendingResources

`func (o *WorkflowTaskHelmDetail) HasPendingResources() bool`

HasPendingResources returns a boolean if a field has been set.

### GetPhase

`func (o *WorkflowTaskHelmDetail) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *WorkflowTaskHelmDetail) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *WorkflowTaskHelmDetail) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *WorkflowTaskHelmDetail) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRelease

`func (o *WorkflowTaskHelmDetail) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *WorkflowTaskHelmDetail) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *WorkflowTaskHelmDetail) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *WorkflowTaskHelmDetail) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRevision

`func (o *WorkflowTaskHelmDetail) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *WorkflowTaskHelmDetail) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *WorkflowTaskHelmDetail) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *WorkflowTaskHelmDetail) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


