# DescribeWorkflowExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Execution** | [**WorkflowExecution**](WorkflowExecution.md) |  | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewDescribeWorkflowExecutionResult

`func NewDescribeWorkflowExecutionResult(environmentId string, execution WorkflowExecution, serviceId string, ) *DescribeWorkflowExecutionResult`

NewDescribeWorkflowExecutionResult instantiates a new DescribeWorkflowExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkflowExecutionResultWithDefaults

`func NewDescribeWorkflowExecutionResultWithDefaults() *DescribeWorkflowExecutionResult`

NewDescribeWorkflowExecutionResultWithDefaults instantiates a new DescribeWorkflowExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *DescribeWorkflowExecutionResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeWorkflowExecutionResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeWorkflowExecutionResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExecution

`func (o *DescribeWorkflowExecutionResult) GetExecution() WorkflowExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *DescribeWorkflowExecutionResult) GetExecutionOk() (*WorkflowExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *DescribeWorkflowExecutionResult) SetExecution(v WorkflowExecution)`

SetExecution sets Execution field to given value.


### GetServiceId

`func (o *DescribeWorkflowExecutionResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeWorkflowExecutionResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeWorkflowExecutionResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


