# DeploymentCellWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**DeploymentCellId** | Pointer to **string** | The deployment cell&#39;s id for the workflow execution. | [optional] 
**EndTime** | Pointer to **string** | End time of the Deployment Cell Workflow in RFC3339 format | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**HostClusterID** | **string** | ID of the Host Cluster | 
**OrgName** | **string** | The name of the deployment cell owner organization. | 
**ParentId** | Pointer to **string** | The parent workflow&#39;s id for the execution. | [optional] 
**StartTime** | **string** | Start time of the Deployment Cell Workflow in RFC3339 format | 
**Status** | **string** | Status of the Deployment Cell Workflow | 
**WorkflowID** | **string** | ID of the Deployment Cell Workflow | 
**WorkflowType** | **string** | The type of workflow execution. | 

## Methods

### NewDeploymentCellWorkflow

`func NewDeploymentCellWorkflow(cloudProvider string, hostClusterID string, orgName string, startTime string, status string, workflowID string, workflowType string, ) *DeploymentCellWorkflow`

NewDeploymentCellWorkflow instantiates a new DeploymentCellWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellWorkflowWithDefaults

`func NewDeploymentCellWorkflowWithDefaults() *DeploymentCellWorkflow`

NewDeploymentCellWorkflowWithDefaults instantiates a new DeploymentCellWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *DeploymentCellWorkflow) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *DeploymentCellWorkflow) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *DeploymentCellWorkflow) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *DeploymentCellWorkflow) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAzureSubscriptionID

`func (o *DeploymentCellWorkflow) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *DeploymentCellWorkflow) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *DeploymentCellWorkflow) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *DeploymentCellWorkflow) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

### GetCloudProvider

`func (o *DeploymentCellWorkflow) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DeploymentCellWorkflow) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DeploymentCellWorkflow) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDeploymentCellId

`func (o *DeploymentCellWorkflow) GetDeploymentCellId() string`

GetDeploymentCellId returns the DeploymentCellId field if non-nil, zero value otherwise.

### GetDeploymentCellIdOk

`func (o *DeploymentCellWorkflow) GetDeploymentCellIdOk() (*string, bool)`

GetDeploymentCellIdOk returns a tuple with the DeploymentCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellId

`func (o *DeploymentCellWorkflow) SetDeploymentCellId(v string)`

SetDeploymentCellId sets DeploymentCellId field to given value.

### HasDeploymentCellId

`func (o *DeploymentCellWorkflow) HasDeploymentCellId() bool`

HasDeploymentCellId returns a boolean if a field has been set.

### GetEndTime

`func (o *DeploymentCellWorkflow) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DeploymentCellWorkflow) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DeploymentCellWorkflow) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DeploymentCellWorkflow) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *DeploymentCellWorkflow) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *DeploymentCellWorkflow) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *DeploymentCellWorkflow) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *DeploymentCellWorkflow) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetHostClusterID

`func (o *DeploymentCellWorkflow) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DeploymentCellWorkflow) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DeploymentCellWorkflow) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetOrgName

`func (o *DeploymentCellWorkflow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DeploymentCellWorkflow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DeploymentCellWorkflow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetParentId

`func (o *DeploymentCellWorkflow) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DeploymentCellWorkflow) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DeploymentCellWorkflow) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DeploymentCellWorkflow) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetStartTime

`func (o *DeploymentCellWorkflow) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DeploymentCellWorkflow) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DeploymentCellWorkflow) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStatus

`func (o *DeploymentCellWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentCellWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentCellWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWorkflowID

`func (o *DeploymentCellWorkflow) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *DeploymentCellWorkflow) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *DeploymentCellWorkflow) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.


### GetWorkflowType

`func (o *DeploymentCellWorkflow) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *DeploymentCellWorkflow) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *DeploymentCellWorkflow) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


