# TerminateDeploymentCellWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | ID of the Host Cluster | 
**Token** | **string** | JWT token used to perform authorization | 
**WorkflowID** | **string** | ID of the Deployment Cell Workflow | 

## Methods

### NewTerminateDeploymentCellWorkflowRequest

`func NewTerminateDeploymentCellWorkflowRequest(hostClusterID string, token string, workflowID string, ) *TerminateDeploymentCellWorkflowRequest`

NewTerminateDeploymentCellWorkflowRequest instantiates a new TerminateDeploymentCellWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateDeploymentCellWorkflowRequestWithDefaults

`func NewTerminateDeploymentCellWorkflowRequestWithDefaults() *TerminateDeploymentCellWorkflowRequest`

NewTerminateDeploymentCellWorkflowRequestWithDefaults instantiates a new TerminateDeploymentCellWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *TerminateDeploymentCellWorkflowRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *TerminateDeploymentCellWorkflowRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *TerminateDeploymentCellWorkflowRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetToken

`func (o *TerminateDeploymentCellWorkflowRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TerminateDeploymentCellWorkflowRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TerminateDeploymentCellWorkflowRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWorkflowID

`func (o *TerminateDeploymentCellWorkflowRequest) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *TerminateDeploymentCellWorkflowRequest) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *TerminateDeploymentCellWorkflowRequest) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


