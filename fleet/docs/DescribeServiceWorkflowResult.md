# DescribeServiceWorkflowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ResourceDeploymentStatus**](ResourceDeploymentStatus.md) | List of resources with deployment status. | [optional] 
**Workflow** | [**ServiceWorkflow**](ServiceWorkflow.md) |  | 
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**ServiceId** | **string** | The service ID this workflow belongs to. | 

## Methods

### NewDescribeServiceWorkflowResult

`func NewDescribeServiceWorkflowResult(workflow ServiceWorkflow, environmentId string, serviceId string, ) *DescribeServiceWorkflowResult`

NewDescribeServiceWorkflowResult instantiates a new DescribeServiceWorkflowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceWorkflowResultWithDefaults

`func NewDescribeServiceWorkflowResultWithDefaults() *DescribeServiceWorkflowResult`

NewDescribeServiceWorkflowResultWithDefaults instantiates a new DescribeServiceWorkflowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *DescribeServiceWorkflowResult) GetResources() []ResourceDeploymentStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DescribeServiceWorkflowResult) GetResourcesOk() (*[]ResourceDeploymentStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DescribeServiceWorkflowResult) SetResources(v []ResourceDeploymentStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DescribeServiceWorkflowResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetWorkflow

`func (o *DescribeServiceWorkflowResult) GetWorkflow() ServiceWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *DescribeServiceWorkflowResult) GetWorkflowOk() (*ServiceWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *DescribeServiceWorkflowResult) SetWorkflow(v ServiceWorkflow)`

SetWorkflow sets Workflow field to given value.


### GetEnvironmentId

`func (o *DescribeServiceWorkflowResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeServiceWorkflowResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeServiceWorkflowResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServiceId

`func (o *DescribeServiceWorkflowResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeServiceWorkflowResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeServiceWorkflowResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


