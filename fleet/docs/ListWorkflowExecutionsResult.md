# ListWorkflowExecutionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Executions** | [**[]WorkflowExecution**](WorkflowExecution.md) | The list of workflow executions. | 
**InstanceId** | **string** | ID of a Resource Instance | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewListWorkflowExecutionsResult

`func NewListWorkflowExecutionsResult(environmentId string, executions []WorkflowExecution, instanceId string, serviceId string, ) *ListWorkflowExecutionsResult`

NewListWorkflowExecutionsResult instantiates a new ListWorkflowExecutionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowExecutionsResultWithDefaults

`func NewListWorkflowExecutionsResultWithDefaults() *ListWorkflowExecutionsResult`

NewListWorkflowExecutionsResultWithDefaults instantiates a new ListWorkflowExecutionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListWorkflowExecutionsResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListWorkflowExecutionsResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListWorkflowExecutionsResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExecutions

`func (o *ListWorkflowExecutionsResult) GetExecutions() []WorkflowExecution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ListWorkflowExecutionsResult) GetExecutionsOk() (*[]WorkflowExecution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ListWorkflowExecutionsResult) SetExecutions(v []WorkflowExecution)`

SetExecutions sets Executions field to given value.


### GetInstanceId

`func (o *ListWorkflowExecutionsResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListWorkflowExecutionsResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListWorkflowExecutionsResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetNextPageToken

`func (o *ListWorkflowExecutionsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListWorkflowExecutionsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListWorkflowExecutionsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListWorkflowExecutionsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *ListWorkflowExecutionsResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListWorkflowExecutionsResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListWorkflowExecutionsResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


