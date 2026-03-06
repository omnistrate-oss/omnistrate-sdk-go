# WorkflowBreakpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **map[string]interface{}** | The conditions to break on | [optional] 
**Id** | **string** | The ID or key of the resource to break on | 

## Methods

### NewWorkflowBreakpoint

`func NewWorkflowBreakpoint(id string, ) *WorkflowBreakpoint`

NewWorkflowBreakpoint instantiates a new WorkflowBreakpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBreakpointWithDefaults

`func NewWorkflowBreakpointWithDefaults() *WorkflowBreakpoint`

NewWorkflowBreakpointWithDefaults instantiates a new WorkflowBreakpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *WorkflowBreakpoint) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowBreakpoint) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowBreakpoint) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowBreakpoint) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *WorkflowBreakpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowBreakpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowBreakpoint) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


