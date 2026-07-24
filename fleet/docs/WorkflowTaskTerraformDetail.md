# WorkflowTaskTerraformDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAddress** | Pointer to **string** | The Terraform resource address that failed, when applicable. | [optional] 
**Phase** | Pointer to **string** | The Terraform phase, e.g. plan, apply. | [optional] 
**PlanAdd** | Pointer to **int64** | The number of resources to add in the current plan. | [optional] 
**PlanChange** | Pointer to **int64** | The number of resources to change in the current plan. | [optional] 
**PlanDestroy** | Pointer to **int64** | The number of resources to destroy in the current plan. | [optional] 
**StateLocked** | Pointer to **bool** | Whether the Terraform state is currently locked. | [optional] 

## Methods

### NewWorkflowTaskTerraformDetail

`func NewWorkflowTaskTerraformDetail() *WorkflowTaskTerraformDetail`

NewWorkflowTaskTerraformDetail instantiates a new WorkflowTaskTerraformDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskTerraformDetailWithDefaults

`func NewWorkflowTaskTerraformDetailWithDefaults() *WorkflowTaskTerraformDetail`

NewWorkflowTaskTerraformDetailWithDefaults instantiates a new WorkflowTaskTerraformDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAddress

`func (o *WorkflowTaskTerraformDetail) GetFailedAddress() string`

GetFailedAddress returns the FailedAddress field if non-nil, zero value otherwise.

### GetFailedAddressOk

`func (o *WorkflowTaskTerraformDetail) GetFailedAddressOk() (*string, bool)`

GetFailedAddressOk returns a tuple with the FailedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAddress

`func (o *WorkflowTaskTerraformDetail) SetFailedAddress(v string)`

SetFailedAddress sets FailedAddress field to given value.

### HasFailedAddress

`func (o *WorkflowTaskTerraformDetail) HasFailedAddress() bool`

HasFailedAddress returns a boolean if a field has been set.

### GetPhase

`func (o *WorkflowTaskTerraformDetail) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *WorkflowTaskTerraformDetail) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *WorkflowTaskTerraformDetail) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *WorkflowTaskTerraformDetail) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlanAdd

`func (o *WorkflowTaskTerraformDetail) GetPlanAdd() int64`

GetPlanAdd returns the PlanAdd field if non-nil, zero value otherwise.

### GetPlanAddOk

`func (o *WorkflowTaskTerraformDetail) GetPlanAddOk() (*int64, bool)`

GetPlanAddOk returns a tuple with the PlanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAdd

`func (o *WorkflowTaskTerraformDetail) SetPlanAdd(v int64)`

SetPlanAdd sets PlanAdd field to given value.

### HasPlanAdd

`func (o *WorkflowTaskTerraformDetail) HasPlanAdd() bool`

HasPlanAdd returns a boolean if a field has been set.

### GetPlanChange

`func (o *WorkflowTaskTerraformDetail) GetPlanChange() int64`

GetPlanChange returns the PlanChange field if non-nil, zero value otherwise.

### GetPlanChangeOk

`func (o *WorkflowTaskTerraformDetail) GetPlanChangeOk() (*int64, bool)`

GetPlanChangeOk returns a tuple with the PlanChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanChange

`func (o *WorkflowTaskTerraformDetail) SetPlanChange(v int64)`

SetPlanChange sets PlanChange field to given value.

### HasPlanChange

`func (o *WorkflowTaskTerraformDetail) HasPlanChange() bool`

HasPlanChange returns a boolean if a field has been set.

### GetPlanDestroy

`func (o *WorkflowTaskTerraformDetail) GetPlanDestroy() int64`

GetPlanDestroy returns the PlanDestroy field if non-nil, zero value otherwise.

### GetPlanDestroyOk

`func (o *WorkflowTaskTerraformDetail) GetPlanDestroyOk() (*int64, bool)`

GetPlanDestroyOk returns a tuple with the PlanDestroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDestroy

`func (o *WorkflowTaskTerraformDetail) SetPlanDestroy(v int64)`

SetPlanDestroy sets PlanDestroy field to given value.

### HasPlanDestroy

`func (o *WorkflowTaskTerraformDetail) HasPlanDestroy() bool`

HasPlanDestroy returns a boolean if a field has been set.

### GetStateLocked

`func (o *WorkflowTaskTerraformDetail) GetStateLocked() bool`

GetStateLocked returns the StateLocked field if non-nil, zero value otherwise.

### GetStateLockedOk

`func (o *WorkflowTaskTerraformDetail) GetStateLockedOk() (*bool, bool)`

GetStateLockedOk returns a tuple with the StateLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateLocked

`func (o *WorkflowTaskTerraformDetail) SetStateLocked(v bool)`

SetStateLocked sets StateLocked field to given value.

### HasStateLocked

`func (o *WorkflowTaskTerraformDetail) HasStateLocked() bool`

HasStateLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


