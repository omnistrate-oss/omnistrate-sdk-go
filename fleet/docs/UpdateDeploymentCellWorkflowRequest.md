# UpdateDeploymentCellWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | ID of the Host Cluster | 
**Status** | **string** | Update the status of the deployment cell workflow execution. You can pause, resume or retry a workflow. | 
**Token** | **string** | JWT token used to perform authorization | 
**WorkflowID** | **string** | ID of the Deployment Cell Workflow | 

## Methods

### NewUpdateDeploymentCellWorkflowRequest

`func NewUpdateDeploymentCellWorkflowRequest(hostClusterID string, status string, token string, workflowID string, ) *UpdateDeploymentCellWorkflowRequest`

NewUpdateDeploymentCellWorkflowRequest instantiates a new UpdateDeploymentCellWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentCellWorkflowRequestWithDefaults

`func NewUpdateDeploymentCellWorkflowRequestWithDefaults() *UpdateDeploymentCellWorkflowRequest`

NewUpdateDeploymentCellWorkflowRequestWithDefaults instantiates a new UpdateDeploymentCellWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *UpdateDeploymentCellWorkflowRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *UpdateDeploymentCellWorkflowRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *UpdateDeploymentCellWorkflowRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetStatus

`func (o *UpdateDeploymentCellWorkflowRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDeploymentCellWorkflowRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDeploymentCellWorkflowRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *UpdateDeploymentCellWorkflowRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateDeploymentCellWorkflowRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateDeploymentCellWorkflowRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWorkflowID

`func (o *UpdateDeploymentCellWorkflowRequest) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *UpdateDeploymentCellWorkflowRequest) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *UpdateDeploymentCellWorkflowRequest) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


