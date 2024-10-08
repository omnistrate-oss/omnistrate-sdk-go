# ListServiceWorkflowsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**ServiceId** | **string** | The service ID this workflow belongs to. | 
**Workflows** | [**[]ServiceWorkflow**](ServiceWorkflow.md) | List of workflows. | 

## Methods

### NewListServiceWorkflowsResult

`func NewListServiceWorkflowsResult(environmentId string, serviceId string, workflows []ServiceWorkflow, ) *ListServiceWorkflowsResult`

NewListServiceWorkflowsResult instantiates a new ListServiceWorkflowsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceWorkflowsResultWithDefaults

`func NewListServiceWorkflowsResultWithDefaults() *ListServiceWorkflowsResult`

NewListServiceWorkflowsResultWithDefaults instantiates a new ListServiceWorkflowsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListServiceWorkflowsResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListServiceWorkflowsResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListServiceWorkflowsResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetNextPageToken

`func (o *ListServiceWorkflowsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServiceWorkflowsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServiceWorkflowsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServiceWorkflowsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *ListServiceWorkflowsResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListServiceWorkflowsResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListServiceWorkflowsResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetWorkflows

`func (o *ListServiceWorkflowsResult) GetWorkflows() []ServiceWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ListServiceWorkflowsResult) GetWorkflowsOk() (*[]ServiceWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ListServiceWorkflowsResult) SetWorkflows(v []ServiceWorkflow)`

SetWorkflows sets Workflows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


