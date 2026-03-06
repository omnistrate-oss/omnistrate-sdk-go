# WorkflowBreakpointWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **map[string]interface{}** | The conditions to break on | [optional] 
**Id** | **string** | The ID or key of the resource to break on | 
**Status** | **string** | The status of the breakpoint | 

## Methods

### NewWorkflowBreakpointWithStatus

`func NewWorkflowBreakpointWithStatus(id string, status string, ) *WorkflowBreakpointWithStatus`

NewWorkflowBreakpointWithStatus instantiates a new WorkflowBreakpointWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBreakpointWithStatusWithDefaults

`func NewWorkflowBreakpointWithStatusWithDefaults() *WorkflowBreakpointWithStatus`

NewWorkflowBreakpointWithStatusWithDefaults instantiates a new WorkflowBreakpointWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *WorkflowBreakpointWithStatus) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowBreakpointWithStatus) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowBreakpointWithStatus) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowBreakpointWithStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *WorkflowBreakpointWithStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowBreakpointWithStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowBreakpointWithStatus) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *WorkflowBreakpointWithStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowBreakpointWithStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowBreakpointWithStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


