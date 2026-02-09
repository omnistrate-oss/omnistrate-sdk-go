# DescribeDeploymentCellWorkflowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | ID of the Host Cluster | 
**Status** | [**DeploymentCellStatus**](DeploymentCellStatus.md) |  | 
**Workflow** | [**DeploymentCellWorkflow**](DeploymentCellWorkflow.md) |  | 
**WorkflowID** | **string** | ID of the Deployment Cell Workflow | 

## Methods

### NewDescribeDeploymentCellWorkflowResult

`func NewDescribeDeploymentCellWorkflowResult(hostClusterID string, status DeploymentCellStatus, workflow DeploymentCellWorkflow, workflowID string, ) *DescribeDeploymentCellWorkflowResult`

NewDescribeDeploymentCellWorkflowResult instantiates a new DescribeDeploymentCellWorkflowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentCellWorkflowResultWithDefaults

`func NewDescribeDeploymentCellWorkflowResultWithDefaults() *DescribeDeploymentCellWorkflowResult`

NewDescribeDeploymentCellWorkflowResultWithDefaults instantiates a new DescribeDeploymentCellWorkflowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *DescribeDeploymentCellWorkflowResult) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DescribeDeploymentCellWorkflowResult) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DescribeDeploymentCellWorkflowResult) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetStatus

`func (o *DescribeDeploymentCellWorkflowResult) GetStatus() DeploymentCellStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeDeploymentCellWorkflowResult) GetStatusOk() (*DeploymentCellStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeDeploymentCellWorkflowResult) SetStatus(v DeploymentCellStatus)`

SetStatus sets Status field to given value.


### GetWorkflow

`func (o *DescribeDeploymentCellWorkflowResult) GetWorkflow() DeploymentCellWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *DescribeDeploymentCellWorkflowResult) GetWorkflowOk() (*DeploymentCellWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *DescribeDeploymentCellWorkflowResult) SetWorkflow(v DeploymentCellWorkflow)`

SetWorkflow sets Workflow field to given value.


### GetWorkflowID

`func (o *DescribeDeploymentCellWorkflowResult) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *DescribeDeploymentCellWorkflowResult) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *DescribeDeploymentCellWorkflowResult) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


