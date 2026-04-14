# InstanceUpgradeHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The timestamp when this history record was created. | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ScheduledAt** | Pointer to **string** | The time at which this upgrade was scheduled. | [optional] 
**SourceVersion** | **string** | The source product tier version before the upgrade. | 
**Status** | **string** | The status of the upgrade path. | 
**StatusMessage** | Pointer to **string** | A human-readable message describing the upgrade outcome. | [optional] 
**TargetVersion** | **string** | The target product tier version for the upgrade. | 
**Type** | **string** | The type of the upgrade path. | 
**UpdatedAt** | **string** | The timestamp when this history record was last updated. | 
**UpgradeEndTime** | Pointer to **string** | The timestamp when the upgrade ended. | [optional] 
**UpgradePathId** | **string** | ID of an Upgrade Path | 
**UpgradeStartTime** | Pointer to **string** | The timestamp when the upgrade started. | [optional] 
**WorkflowID** | **string** | The workflow ID associated with this instance upgrade. | 

## Methods

### NewInstanceUpgradeHistory

`func NewInstanceUpgradeHistory(createdAt string, instanceId string, productTierId string, sourceVersion string, status string, targetVersion string, type_ string, updatedAt string, upgradePathId string, workflowID string, ) *InstanceUpgradeHistory`

NewInstanceUpgradeHistory instantiates a new InstanceUpgradeHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpgradeHistoryWithDefaults

`func NewInstanceUpgradeHistoryWithDefaults() *InstanceUpgradeHistory`

NewInstanceUpgradeHistoryWithDefaults instantiates a new InstanceUpgradeHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InstanceUpgradeHistory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceUpgradeHistory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceUpgradeHistory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetInstanceId

`func (o *InstanceUpgradeHistory) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceUpgradeHistory) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceUpgradeHistory) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetProductTierId

`func (o *InstanceUpgradeHistory) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *InstanceUpgradeHistory) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *InstanceUpgradeHistory) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetScheduledAt

`func (o *InstanceUpgradeHistory) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *InstanceUpgradeHistory) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *InstanceUpgradeHistory) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *InstanceUpgradeHistory) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetSourceVersion

`func (o *InstanceUpgradeHistory) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *InstanceUpgradeHistory) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *InstanceUpgradeHistory) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetStatus

`func (o *InstanceUpgradeHistory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceUpgradeHistory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceUpgradeHistory) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *InstanceUpgradeHistory) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InstanceUpgradeHistory) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InstanceUpgradeHistory) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InstanceUpgradeHistory) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTargetVersion

`func (o *InstanceUpgradeHistory) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *InstanceUpgradeHistory) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *InstanceUpgradeHistory) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetType

`func (o *InstanceUpgradeHistory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceUpgradeHistory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceUpgradeHistory) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *InstanceUpgradeHistory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceUpgradeHistory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceUpgradeHistory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpgradeEndTime

`func (o *InstanceUpgradeHistory) GetUpgradeEndTime() string`

GetUpgradeEndTime returns the UpgradeEndTime field if non-nil, zero value otherwise.

### GetUpgradeEndTimeOk

`func (o *InstanceUpgradeHistory) GetUpgradeEndTimeOk() (*string, bool)`

GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeEndTime

`func (o *InstanceUpgradeHistory) SetUpgradeEndTime(v string)`

SetUpgradeEndTime sets UpgradeEndTime field to given value.

### HasUpgradeEndTime

`func (o *InstanceUpgradeHistory) HasUpgradeEndTime() bool`

HasUpgradeEndTime returns a boolean if a field has been set.

### GetUpgradePathId

`func (o *InstanceUpgradeHistory) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *InstanceUpgradeHistory) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *InstanceUpgradeHistory) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.


### GetUpgradeStartTime

`func (o *InstanceUpgradeHistory) GetUpgradeStartTime() string`

GetUpgradeStartTime returns the UpgradeStartTime field if non-nil, zero value otherwise.

### GetUpgradeStartTimeOk

`func (o *InstanceUpgradeHistory) GetUpgradeStartTimeOk() (*string, bool)`

GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStartTime

`func (o *InstanceUpgradeHistory) SetUpgradeStartTime(v string)`

SetUpgradeStartTime sets UpgradeStartTime field to given value.

### HasUpgradeStartTime

`func (o *InstanceUpgradeHistory) HasUpgradeStartTime() bool`

HasUpgradeStartTime returns a boolean if a field has been set.

### GetWorkflowID

`func (o *InstanceUpgradeHistory) GetWorkflowID() string`

GetWorkflowID returns the WorkflowID field if non-nil, zero value otherwise.

### GetWorkflowIDOk

`func (o *InstanceUpgradeHistory) GetWorkflowIDOk() (*string, bool)`

GetWorkflowIDOk returns a tuple with the WorkflowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowID

`func (o *InstanceUpgradeHistory) SetWorkflowID(v string)`

SetWorkflowID sets WorkflowID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


