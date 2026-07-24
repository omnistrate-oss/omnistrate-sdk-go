# WorkflowTaskInfraDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentResource** | Pointer to **string** | The infrastructure resource currently being provisioned. | [optional] 
**CurrentSince** | Pointer to **string** | The time the current infrastructure operation began, in RFC3339 format. | [optional] 
**LastDiagnostic** | Pointer to **string** | The most recent provisioning diagnostic message. | [optional] 
**OpsCompleted** | Pointer to **int64** | The number of infrastructure operations completed so far. | [optional] 
**OpsInFlight** | Pointer to **int64** | The number of infrastructure operations currently in flight. | [optional] 
**Stack** | Pointer to **string** | The infrastructure stack name. | [optional] 

## Methods

### NewWorkflowTaskInfraDetail

`func NewWorkflowTaskInfraDetail() *WorkflowTaskInfraDetail`

NewWorkflowTaskInfraDetail instantiates a new WorkflowTaskInfraDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskInfraDetailWithDefaults

`func NewWorkflowTaskInfraDetailWithDefaults() *WorkflowTaskInfraDetail`

NewWorkflowTaskInfraDetailWithDefaults instantiates a new WorkflowTaskInfraDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentResource

`func (o *WorkflowTaskInfraDetail) GetCurrentResource() string`

GetCurrentResource returns the CurrentResource field if non-nil, zero value otherwise.

### GetCurrentResourceOk

`func (o *WorkflowTaskInfraDetail) GetCurrentResourceOk() (*string, bool)`

GetCurrentResourceOk returns a tuple with the CurrentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentResource

`func (o *WorkflowTaskInfraDetail) SetCurrentResource(v string)`

SetCurrentResource sets CurrentResource field to given value.

### HasCurrentResource

`func (o *WorkflowTaskInfraDetail) HasCurrentResource() bool`

HasCurrentResource returns a boolean if a field has been set.

### GetCurrentSince

`func (o *WorkflowTaskInfraDetail) GetCurrentSince() string`

GetCurrentSince returns the CurrentSince field if non-nil, zero value otherwise.

### GetCurrentSinceOk

`func (o *WorkflowTaskInfraDetail) GetCurrentSinceOk() (*string, bool)`

GetCurrentSinceOk returns a tuple with the CurrentSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSince

`func (o *WorkflowTaskInfraDetail) SetCurrentSince(v string)`

SetCurrentSince sets CurrentSince field to given value.

### HasCurrentSince

`func (o *WorkflowTaskInfraDetail) HasCurrentSince() bool`

HasCurrentSince returns a boolean if a field has been set.

### GetLastDiagnostic

`func (o *WorkflowTaskInfraDetail) GetLastDiagnostic() string`

GetLastDiagnostic returns the LastDiagnostic field if non-nil, zero value otherwise.

### GetLastDiagnosticOk

`func (o *WorkflowTaskInfraDetail) GetLastDiagnosticOk() (*string, bool)`

GetLastDiagnosticOk returns a tuple with the LastDiagnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiagnostic

`func (o *WorkflowTaskInfraDetail) SetLastDiagnostic(v string)`

SetLastDiagnostic sets LastDiagnostic field to given value.

### HasLastDiagnostic

`func (o *WorkflowTaskInfraDetail) HasLastDiagnostic() bool`

HasLastDiagnostic returns a boolean if a field has been set.

### GetOpsCompleted

`func (o *WorkflowTaskInfraDetail) GetOpsCompleted() int64`

GetOpsCompleted returns the OpsCompleted field if non-nil, zero value otherwise.

### GetOpsCompletedOk

`func (o *WorkflowTaskInfraDetail) GetOpsCompletedOk() (*int64, bool)`

GetOpsCompletedOk returns a tuple with the OpsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCompleted

`func (o *WorkflowTaskInfraDetail) SetOpsCompleted(v int64)`

SetOpsCompleted sets OpsCompleted field to given value.

### HasOpsCompleted

`func (o *WorkflowTaskInfraDetail) HasOpsCompleted() bool`

HasOpsCompleted returns a boolean if a field has been set.

### GetOpsInFlight

`func (o *WorkflowTaskInfraDetail) GetOpsInFlight() int64`

GetOpsInFlight returns the OpsInFlight field if non-nil, zero value otherwise.

### GetOpsInFlightOk

`func (o *WorkflowTaskInfraDetail) GetOpsInFlightOk() (*int64, bool)`

GetOpsInFlightOk returns a tuple with the OpsInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsInFlight

`func (o *WorkflowTaskInfraDetail) SetOpsInFlight(v int64)`

SetOpsInFlight sets OpsInFlight field to given value.

### HasOpsInFlight

`func (o *WorkflowTaskInfraDetail) HasOpsInFlight() bool`

HasOpsInFlight returns a boolean if a field has been set.

### GetStack

`func (o *WorkflowTaskInfraDetail) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *WorkflowTaskInfraDetail) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *WorkflowTaskInfraDetail) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *WorkflowTaskInfraDetail) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


