# DescribeWorkflowTaskObjectsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExecutionId** | **string** |  | 
**LiveObservedAt** | Pointer to **string** | When live state was read, in RFC3339 format | [optional] 
**Objects** | [**[]WorkflowTaskObject**](WorkflowTaskObject.md) | Objects the task touched, in the order it touched them | 
**ServiceId** | **string** | ID of a Service | 
**TaskName** | **string** |  | 

## Methods

### NewDescribeWorkflowTaskObjectsResult

`func NewDescribeWorkflowTaskObjectsResult(environmentId string, executionId string, objects []WorkflowTaskObject, serviceId string, taskName string, ) *DescribeWorkflowTaskObjectsResult`

NewDescribeWorkflowTaskObjectsResult instantiates a new DescribeWorkflowTaskObjectsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkflowTaskObjectsResultWithDefaults

`func NewDescribeWorkflowTaskObjectsResultWithDefaults() *DescribeWorkflowTaskObjectsResult`

NewDescribeWorkflowTaskObjectsResultWithDefaults instantiates a new DescribeWorkflowTaskObjectsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *DescribeWorkflowTaskObjectsResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeWorkflowTaskObjectsResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeWorkflowTaskObjectsResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExecutionId

`func (o *DescribeWorkflowTaskObjectsResult) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DescribeWorkflowTaskObjectsResult) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DescribeWorkflowTaskObjectsResult) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetLiveObservedAt

`func (o *DescribeWorkflowTaskObjectsResult) GetLiveObservedAt() string`

GetLiveObservedAt returns the LiveObservedAt field if non-nil, zero value otherwise.

### GetLiveObservedAtOk

`func (o *DescribeWorkflowTaskObjectsResult) GetLiveObservedAtOk() (*string, bool)`

GetLiveObservedAtOk returns a tuple with the LiveObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveObservedAt

`func (o *DescribeWorkflowTaskObjectsResult) SetLiveObservedAt(v string)`

SetLiveObservedAt sets LiveObservedAt field to given value.

### HasLiveObservedAt

`func (o *DescribeWorkflowTaskObjectsResult) HasLiveObservedAt() bool`

HasLiveObservedAt returns a boolean if a field has been set.

### GetObjects

`func (o *DescribeWorkflowTaskObjectsResult) GetObjects() []WorkflowTaskObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DescribeWorkflowTaskObjectsResult) GetObjectsOk() (*[]WorkflowTaskObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DescribeWorkflowTaskObjectsResult) SetObjects(v []WorkflowTaskObject)`

SetObjects sets Objects field to given value.


### GetServiceId

`func (o *DescribeWorkflowTaskObjectsResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeWorkflowTaskObjectsResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeWorkflowTaskObjectsResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTaskName

`func (o *DescribeWorkflowTaskObjectsResult) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *DescribeWorkflowTaskObjectsResult) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *DescribeWorkflowTaskObjectsResult) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


