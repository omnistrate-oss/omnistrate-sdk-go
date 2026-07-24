# WorkflowTaskHelmDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedHook** | Pointer to **string** | The Helm hook that failed, when applicable. | [optional] 
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


