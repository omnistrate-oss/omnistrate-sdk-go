# InstanceUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | The name of the cloud provider that the instance is running on. | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**CreatedAt** | **string** | The timestamp when the instance was created. | 
**HealthStatus** | Pointer to **string** | The heath status of a resource | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 
**LifecycleStatus** | **string** | The status of an operation | 
**ManagedResourceType** | Pointer to **string** | The managed resource type of the top-level resource of the instance. | [optional] 
**OrgName** | **string** | The name of the organization that owns the instance. | 
**ResourceName** | **string** | The name of the top-level resource of the instance. | 
**Status** | **string** | The status of the upgrade path. | 
**UpdatedAt** | **string** | The timestamp when the instance was updated. | 
**UpgradeEndTime** | Pointer to **string** | The timestamp when the upgrade ended. | [optional] 
**UpgradeStartTime** | Pointer to **string** | The timestamp when the upgrade started. | [optional] 
**WorkflowID** | **string** | The workflow ID of the instance upgrade. | 

## Methods

### NewInstanceUpgrade

`func NewInstanceUpgrade(cloudProviderName string, cloudProviderRegion string, createdAt string, instanceId string, lifecycleStatus string, orgName string, resourceName string, status string, updatedAt string, workflowID string, ) *InstanceUpgrade`

NewInstanceUpgrade instantiates a new InstanceUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpgradeWithDefaults

`func NewInstanceUpgradeWithDefaults() *InstanceUpgrade`

NewInstanceUpgradeWithDefaults instantiates a new InstanceUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *InstanceUpgrade) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *InstanceUpgrade) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *InstanceUpgrade) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *InstanceUpgrade) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *InstanceUpgrade) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *InstanceUpgrade) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetCreatedAt

`func (o *InstanceUpgrade) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceUpgrade) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceUpgrade) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHealthStatus

`func (o *InstanceUpgrade) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *InstanceUpgrade) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *InstanceUpgrade) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *InstanceUpgrade) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetInstanceId

`func (o *InstanceUpgrade) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceUpgrade) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceUpgrade) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetLifecycleStatus

`func (o *InstanceUpgrade) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *InstanceUpgrade) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *InstanceUpgrade) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.


### GetManagedResourceType

`func (o *InstanceUpgrade) GetManagedResourceType() string`

GetManagedResourceType returns the ManagedResourceType field if non-nil, zero value otherwise.

### GetManagedResourceTypeOk

`func (o *InstanceUpgrade) GetManagedResourceTypeOk() (*string, bool)`

GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceType

`func (o *InstanceUpgrade) SetManagedResourceType(v string)`

SetManagedResourceType sets ManagedResourceType field to given value.

### HasManagedResourceType

`func (o *InstanceUpgrade) HasManagedResourceType() bool`

HasManagedResourceType returns a boolean if a field has been set.

### GetOrgName

`func (o *InstanceUpgrade) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *InstanceUpgrade) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *InstanceUpgrade) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetResourceName

`func (o *InstanceUpgrade) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InstanceUpgrade) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InstanceUpgrade) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetStatus

`func (o *InstanceUpgrade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceUpgrade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceUpgrade) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *InstanceUpgrade) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceUpgrade) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceUpgrade) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpgradeEndTime

`func (o *InstanceUpgrade) GetUpgradeEndTime() string`

GetUpgradeEndTime returns the UpgradeEndTime field if non-nil, zero value otherwise.

### GetUpgradeEndTimeOk

`func (o *InstanceUpgrade) GetUpgradeEndTimeOk() (*string, bool)`

GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeEndTime

`func (o *InstanceUpgrade) SetUpgradeEndTime(v string)`

SetUpgradeEndTime sets UpgradeEndTime field to given value.

### HasUpgradeEndTime

`func (o *InstanceUpgrade) HasUpgradeEndTime() bool`

HasUpgradeEndTime returns a boolean if a field has been set.

### GetUpgradeStartTime

`func (o *InstanceUpgrade) GetUpgradeStartTime() string`

GetUpgradeStartTime returns the UpgradeStartTime field if non-nil, zero value otherwise.

### GetUpgradeStartTimeOk

`func (o *InstanceUpgrade) GetUpgradeStartTimeOk() (*string, bool)`

GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStartTime

`func (o *InstanceUpgrade) SetUpgradeStartTime(v string)`

SetUpgradeStartTime sets UpgradeStartTime field to given value.

### HasUpgradeStartTime

`func (o *InstanceUpgrade) HasUpgradeStartTime() bool`

HasUpgradeStartTime returns a boolean if a field has been set.

### GetWorkflowID

`func (o *InstanceUpgrade) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *InstanceUpgrade) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *InstanceUpgrade) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


