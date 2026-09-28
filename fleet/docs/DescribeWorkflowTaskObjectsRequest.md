# DescribeWorkflowTaskObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExecutionId** | **string** | The workflow execution that ran the task. | 
**Live** | Pointer to **bool** | Read each object&#39;s live state from the cluster. When false, only what was recorded when the task ran is returned. | [optional] [default to true]
**ServiceId** | **string** | ID of a Service | 
**TaskName** | **string** | The workflow task whose objects to describe. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeWorkflowTaskObjectsRequest

`func NewDescribeWorkflowTaskObjectsRequest(environmentId string, executionId string, serviceId string, taskName string, token string, ) *DescribeWorkflowTaskObjectsRequest`

NewDescribeWorkflowTaskObjectsRequest instantiates a new DescribeWorkflowTaskObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkflowTaskObjectsRequestWithDefaults

`func NewDescribeWorkflowTaskObjectsRequestWithDefaults() *DescribeWorkflowTaskObjectsRequest`

NewDescribeWorkflowTaskObjectsRequestWithDefaults instantiates a new DescribeWorkflowTaskObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *DescribeWorkflowTaskObjectsRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeWorkflowTaskObjectsRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExecutionId

`func (o *DescribeWorkflowTaskObjectsRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DescribeWorkflowTaskObjectsRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetLive

`func (o *DescribeWorkflowTaskObjectsRequest) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *DescribeWorkflowTaskObjectsRequest) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *DescribeWorkflowTaskObjectsRequest) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeWorkflowTaskObjectsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeWorkflowTaskObjectsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTaskName

`func (o *DescribeWorkflowTaskObjectsRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *DescribeWorkflowTaskObjectsRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetToken

`func (o *DescribeWorkflowTaskObjectsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeWorkflowTaskObjectsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeWorkflowTaskObjectsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


