# ListDeploymentCellWorkflowsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | ID of the Host Cluster | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**TotalCount** | **int64** | Total number of Deployment Cell Workflows matching the criteria | 
**Workflows** | [**[]DeploymentCellWorkflow**](DeploymentCellWorkflow.md) | List of Deployment Cell Workflows | 

## Methods

### NewListDeploymentCellWorkflowsResult

`func NewListDeploymentCellWorkflowsResult(hostClusterID string, totalCount int64, workflows []DeploymentCellWorkflow, ) *ListDeploymentCellWorkflowsResult`

NewListDeploymentCellWorkflowsResult instantiates a new ListDeploymentCellWorkflowsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentCellWorkflowsResultWithDefaults

`func NewListDeploymentCellWorkflowsResultWithDefaults() *ListDeploymentCellWorkflowsResult`

NewListDeploymentCellWorkflowsResultWithDefaults instantiates a new ListDeploymentCellWorkflowsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *ListDeploymentCellWorkflowsResult) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *ListDeploymentCellWorkflowsResult) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *ListDeploymentCellWorkflowsResult) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetNextPageToken

`func (o *ListDeploymentCellWorkflowsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListDeploymentCellWorkflowsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListDeploymentCellWorkflowsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListDeploymentCellWorkflowsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListDeploymentCellWorkflowsResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDeploymentCellWorkflowsResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDeploymentCellWorkflowsResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetWorkflows

`func (o *ListDeploymentCellWorkflowsResult) GetWorkflows() []DeploymentCellWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ListDeploymentCellWorkflowsResult) GetWorkflowsOk() (*[]DeploymentCellWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ListDeploymentCellWorkflowsResult) SetWorkflows(v []DeploymentCellWorkflow)`

SetWorkflows sets Workflows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


