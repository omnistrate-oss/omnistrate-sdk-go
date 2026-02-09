# ServiceWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceCount** | Pointer to **int64** | The number of resources updated by the workflow. | [optional] 
**UpdatedBy** | Pointer to **string** | The user who updated the workflow execution. - Pause, Resume, Terminate | [optional] 
**UpdatedReason** | Pointer to **string** | The reason the workflow execution was terminated. | [optional] 
**WorkflowType** | **string** | The type of workflow execution. | 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**EndTime** | Pointer to **string** | The time the workflow execution ended. | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**Id** | **string** | ID of the ServiceWorkflow | 
**ManualOverride** | Pointer to [**ManualOverride**](ManualOverride.md) |  | [optional] 
**OciTenancyID** | Pointer to **string** | The Tenancy OCID for Oracle Cloud Infrastructure | [optional] 
**OrgName** | **string** | The name of the instance owner organization. | 
**ParentId** | Pointer to **string** | The parent workflow&#39;s id for the execution. | [optional] 
**PlanType** | Pointer to **string** | The plan type of the instance owner organization. | [optional] 
**ServicePlanName** | Pointer to **string** | The service plan name for the workflow. | [optional] 
**StartTime** | **string** | The time the workflow execution started. | 
**Status** | **string** | The status of the workflow execution. | 

## Methods

### NewServiceWorkflow

`func NewServiceWorkflow(workflowType string, cloudProvider string, id string, orgName string, startTime string, status string, ) *ServiceWorkflow`

NewServiceWorkflow instantiates a new ServiceWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWorkflowWithDefaults

`func NewServiceWorkflowWithDefaults() *ServiceWorkflow`

NewServiceWorkflowWithDefaults instantiates a new ServiceWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceCount

`func (o *ServiceWorkflow) GetResourceCount() int64`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *ServiceWorkflow) GetResourceCountOk() (*int64, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *ServiceWorkflow) SetResourceCount(v int64)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *ServiceWorkflow) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ServiceWorkflow) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceWorkflow) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceWorkflow) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServiceWorkflow) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedReason

`func (o *ServiceWorkflow) GetUpdatedReason() string`

GetUpdatedReason returns the UpdatedReason field if non-nil, zero value otherwise.

### GetUpdatedReasonOk

`func (o *ServiceWorkflow) GetUpdatedReasonOk() (*string, bool)`

GetUpdatedReasonOk returns a tuple with the UpdatedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReason

`func (o *ServiceWorkflow) SetUpdatedReason(v string)`

SetUpdatedReason sets UpdatedReason field to given value.

### HasUpdatedReason

`func (o *ServiceWorkflow) HasUpdatedReason() bool`

HasUpdatedReason returns a boolean if a field has been set.

### GetWorkflowType

`func (o *ServiceWorkflow) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *ServiceWorkflow) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *ServiceWorkflow) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### GetAwsAccountID

`func (o *ServiceWorkflow) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *ServiceWorkflow) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *ServiceWorkflow) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *ServiceWorkflow) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAzureSubscriptionID

`func (o *ServiceWorkflow) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *ServiceWorkflow) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *ServiceWorkflow) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *ServiceWorkflow) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ServiceWorkflow) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ServiceWorkflow) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ServiceWorkflow) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetEndTime

`func (o *ServiceWorkflow) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ServiceWorkflow) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ServiceWorkflow) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ServiceWorkflow) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *ServiceWorkflow) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *ServiceWorkflow) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *ServiceWorkflow) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *ServiceWorkflow) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetId

`func (o *ServiceWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceWorkflow) SetId(v string)`

SetId sets Id field to given value.


### GetManualOverride

`func (o *ServiceWorkflow) GetManualOverride() ManualOverride`

GetManualOverride returns the ManualOverride field if non-nil, zero value otherwise.

### GetManualOverrideOk

`func (o *ServiceWorkflow) GetManualOverrideOk() (*ManualOverride, bool)`

GetManualOverrideOk returns a tuple with the ManualOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverride

`func (o *ServiceWorkflow) SetManualOverride(v ManualOverride)`

SetManualOverride sets ManualOverride field to given value.

### HasManualOverride

`func (o *ServiceWorkflow) HasManualOverride() bool`

HasManualOverride returns a boolean if a field has been set.

### GetOciTenancyID

`func (o *ServiceWorkflow) GetOciTenancyID() string`

GetOciTenancyID returns the OciTenancyID field if non-nil, zero value otherwise.

### GetOciTenancyIDOk

`func (o *ServiceWorkflow) GetOciTenancyIDOk() (*string, bool)`

GetOciTenancyIDOk returns a tuple with the OciTenancyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciTenancyID

`func (o *ServiceWorkflow) SetOciTenancyID(v string)`

SetOciTenancyID sets OciTenancyID field to given value.

### HasOciTenancyID

`func (o *ServiceWorkflow) HasOciTenancyID() bool`

HasOciTenancyID returns a boolean if a field has been set.

### GetOrgName

`func (o *ServiceWorkflow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ServiceWorkflow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ServiceWorkflow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetParentId

`func (o *ServiceWorkflow) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ServiceWorkflow) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ServiceWorkflow) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ServiceWorkflow) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPlanType

`func (o *ServiceWorkflow) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *ServiceWorkflow) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *ServiceWorkflow) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *ServiceWorkflow) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetServicePlanName

`func (o *ServiceWorkflow) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *ServiceWorkflow) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *ServiceWorkflow) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *ServiceWorkflow) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetStartTime

`func (o *ServiceWorkflow) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ServiceWorkflow) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ServiceWorkflow) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStatus

`func (o *ServiceWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


