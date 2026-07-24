# ResourceInstanceCustomWorkflowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of an operation | [optional] 
**WorkflowExecutionId** | **string** | ID of a Workflow | 
**WorkflowId** | **string** | ID of a Custom Workflow | 

## Methods

### NewResourceInstanceCustomWorkflowResult

`func NewResourceInstanceCustomWorkflowResult(workflowExecutionId string, workflowId string, ) *ResourceInstanceCustomWorkflowResult`

NewResourceInstanceCustomWorkflowResult instantiates a new ResourceInstanceCustomWorkflowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceCustomWorkflowResultWithDefaults

`func NewResourceInstanceCustomWorkflowResultWithDefaults() *ResourceInstanceCustomWorkflowResult`

NewResourceInstanceCustomWorkflowResultWithDefaults instantiates a new ResourceInstanceCustomWorkflowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResourceInstanceCustomWorkflowResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceInstanceCustomWorkflowResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceInstanceCustomWorkflowResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceInstanceCustomWorkflowResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowExecutionId

`func (o *ResourceInstanceCustomWorkflowResult) GetWorkflowExecutionId() string`

GetWorkflowExecutionId returns the WorkflowExecutionId field if non-nil, zero value otherwise.

### GetWorkflowExecutionIdOk

`func (o *ResourceInstanceCustomWorkflowResult) GetWorkflowExecutionIdOk() (*string, bool)`

GetWorkflowExecutionIdOk returns a tuple with the WorkflowExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecutionId

`func (o *ResourceInstanceCustomWorkflowResult) SetWorkflowExecutionId(v string)`

SetWorkflowExecutionId sets WorkflowExecutionId field to given value.


### GetWorkflowId

`func (o *ResourceInstanceCustomWorkflowResult) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ResourceInstanceCustomWorkflowResult) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ResourceInstanceCustomWorkflowResult) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


