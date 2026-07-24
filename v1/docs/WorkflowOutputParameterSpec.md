# WorkflowOutputParameterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description of the output parameter. | [optional] 
**Name** | **string** | The output parameter name. | 
**Type** | **string** | The value type used to render and validate the output. One of: string, int, float, bool, json. | 
**ValueRef** | **string** | The reference expression resolving the value from a task resource, e.g. $tasks.&lt;taskName&gt;.resource.&lt;jsonPath&gt;. | 

## Methods

### NewWorkflowOutputParameterSpec

`func NewWorkflowOutputParameterSpec(name string, type_ string, valueRef string, ) *WorkflowOutputParameterSpec`

NewWorkflowOutputParameterSpec instantiates a new WorkflowOutputParameterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputParameterSpecWithDefaults

`func NewWorkflowOutputParameterSpecWithDefaults() *WorkflowOutputParameterSpec`

NewWorkflowOutputParameterSpecWithDefaults instantiates a new WorkflowOutputParameterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowOutputParameterSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowOutputParameterSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowOutputParameterSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowOutputParameterSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowOutputParameterSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowOutputParameterSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowOutputParameterSpec) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WorkflowOutputParameterSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowOutputParameterSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowOutputParameterSpec) SetType(v string)`

SetType sets Type field to given value.


### GetValueRef

`func (o *WorkflowOutputParameterSpec) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *WorkflowOutputParameterSpec) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *WorkflowOutputParameterSpec) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


