# GetDeploymentCellWorkflowEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsPerWorkflowStep** | [**[]DeploymentCellEventsPerWorkflowStep**](DeploymentCellEventsPerWorkflowStep.md) | List of events per workflow step | 
**HostClusterID** | **string** | ID of the Host Cluster | 
**WorkflowID** | **string** | ID of the Deployment Cell Workflow | 

## Methods

### NewGetDeploymentCellWorkflowEventsResult

`func NewGetDeploymentCellWorkflowEventsResult(eventsPerWorkflowStep []DeploymentCellEventsPerWorkflowStep, hostClusterID string, workflowID string, ) *GetDeploymentCellWorkflowEventsResult`

NewGetDeploymentCellWorkflowEventsResult instantiates a new GetDeploymentCellWorkflowEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentCellWorkflowEventsResultWithDefaults

`func NewGetDeploymentCellWorkflowEventsResultWithDefaults() *GetDeploymentCellWorkflowEventsResult`

NewGetDeploymentCellWorkflowEventsResultWithDefaults instantiates a new GetDeploymentCellWorkflowEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsPerWorkflowStep

`func (o *GetDeploymentCellWorkflowEventsResult) GetEventsPerWorkflowStep() []DeploymentCellEventsPerWorkflowStep`

GetEventsPerWorkflowStep returns the EventsPerWorkflowStep field if non-nil, zero value otherwise.

### GetEventsPerWorkflowStepOk

`func (o *GetDeploymentCellWorkflowEventsResult) GetEventsPerWorkflowStepOk() (*[]DeploymentCellEventsPerWorkflowStep, bool)`

GetEventsPerWorkflowStepOk returns a tuple with the EventsPerWorkflowStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsPerWorkflowStep

`func (o *GetDeploymentCellWorkflowEventsResult) SetEventsPerWorkflowStep(v []DeploymentCellEventsPerWorkflowStep)`

SetEventsPerWorkflowStep sets EventsPerWorkflowStep field to given value.


### GetHostClusterID

`func (o *GetDeploymentCellWorkflowEventsResult) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *GetDeploymentCellWorkflowEventsResult) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *GetDeploymentCellWorkflowEventsResult) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetWorkflowID

`func (o *GetDeploymentCellWorkflowEventsResult) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *GetDeploymentCellWorkflowEventsResult) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *GetDeploymentCellWorkflowEventsResult) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


