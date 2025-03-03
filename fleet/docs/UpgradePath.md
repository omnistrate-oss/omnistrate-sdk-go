# UpgradePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **string** | The timestamp when the upgrade path was completed. | 
**CompletedCount** | **int64** | The number of instances that have completed the upgrade. | 
**CreatedAt** | **string** | The timestamp when the upgrade path was created. | 
**CreatedBy** | Pointer to **string** | The name of the user who created the upgrade path. | [optional] 
**FailedCount** | **int64** | The number of instances that have failed the upgrade. | 
**InProgressCount** | **int64** | The number of instances that are in progress of the upgrade. | 
**LastModifiedBy** | Pointer to **string** | The name of the user who created the upgrade path. | [optional] 
**LastRequestedAction** | Pointer to **string** | The action to perform on an ongoing resource workflow | [optional] 
**PendingCount** | **int64** | The number of instances that are pending the upgrade. | 
**PlannedExecutionDate** | Pointer to **string** | The date when the upgrade was/is planned to be executed. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ReleasedAt** | **string** | The timestamp when the upgrade path was released. | 
**ScheduledCount** | Pointer to **int64** | The number of instances that are scheduled to upgrade. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SkippedCount** | **int64** | The number of instances that are skipped the upgrade. | 
**SourceVersion** | **string** | The source version of the upgrade path. | 
**Status** | **string** | The status of the upgrade path. | 
**TargetVersion** | **string** | The target version of the upgrade path. | 
**TotalCount** | **int64** | The total number of instances that are eligible for the upgrade. | 
**Type** | **string** | The type of the upgrade path. | 
**UpdatedAt** | **string** | The timestamp when the upgrade path was updated. | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewUpgradePath

`func NewUpgradePath(completedAt string, completedCount int64, createdAt string, failedCount int64, inProgressCount int64, pendingCount int64, productTierId string, releasedAt string, serviceId string, skippedCount int64, sourceVersion string, status string, targetVersion string, totalCount int64, type_ string, updatedAt string, upgradePathId string, ) *UpgradePath`

NewUpgradePath instantiates a new UpgradePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradePathWithDefaults

`func NewUpgradePathWithDefaults() *UpgradePath`

NewUpgradePathWithDefaults instantiates a new UpgradePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *UpgradePath) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *UpgradePath) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *UpgradePath) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### GetCompletedCount

`func (o *UpgradePath) GetCompletedCount() int64`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *UpgradePath) GetCompletedCountOk() (*int64, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *UpgradePath) SetCompletedCount(v int64)`

SetCompletedCount sets CompletedCount field to given value.


### GetCreatedAt

`func (o *UpgradePath) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpgradePath) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpgradePath) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *UpgradePath) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpgradePath) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpgradePath) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpgradePath) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetFailedCount

`func (o *UpgradePath) GetFailedCount() int64`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UpgradePath) GetFailedCountOk() (*int64, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UpgradePath) SetFailedCount(v int64)`

SetFailedCount sets FailedCount field to given value.


### GetInProgressCount

`func (o *UpgradePath) GetInProgressCount() int64`

GetInProgressCount returns the InProgressCount field if non-nil, zero value otherwise.

### GetInProgressCountOk

`func (o *UpgradePath) GetInProgressCountOk() (*int64, bool)`

GetInProgressCountOk returns a tuple with the InProgressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressCount

`func (o *UpgradePath) SetInProgressCount(v int64)`

SetInProgressCount sets InProgressCount field to given value.


### GetLastModifiedBy

`func (o *UpgradePath) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *UpgradePath) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *UpgradePath) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *UpgradePath) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastRequestedAction

`func (o *UpgradePath) GetLastRequestedAction() string`

GetLastRequestedAction returns the LastRequestedAction field if non-nil, zero value otherwise.

### GetLastRequestedActionOk

`func (o *UpgradePath) GetLastRequestedActionOk() (*string, bool)`

GetLastRequestedActionOk returns a tuple with the LastRequestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequestedAction

`func (o *UpgradePath) SetLastRequestedAction(v string)`

SetLastRequestedAction sets LastRequestedAction field to given value.

### HasLastRequestedAction

`func (o *UpgradePath) HasLastRequestedAction() bool`

HasLastRequestedAction returns a boolean if a field has been set.

### GetPendingCount

`func (o *UpgradePath) GetPendingCount() int64`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *UpgradePath) GetPendingCountOk() (*int64, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *UpgradePath) SetPendingCount(v int64)`

SetPendingCount sets PendingCount field to given value.


### GetPlannedExecutionDate

`func (o *UpgradePath) GetPlannedExecutionDate() string`

GetPlannedExecutionDate returns the PlannedExecutionDate field if non-nil, zero value otherwise.

### GetPlannedExecutionDateOk

`func (o *UpgradePath) GetPlannedExecutionDateOk() (*string, bool)`

GetPlannedExecutionDateOk returns a tuple with the PlannedExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedExecutionDate

`func (o *UpgradePath) SetPlannedExecutionDate(v string)`

SetPlannedExecutionDate sets PlannedExecutionDate field to given value.

### HasPlannedExecutionDate

`func (o *UpgradePath) HasPlannedExecutionDate() bool`

HasPlannedExecutionDate returns a boolean if a field has been set.

### GetProductTierId

`func (o *UpgradePath) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *UpgradePath) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *UpgradePath) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetReleasedAt

`func (o *UpgradePath) GetReleasedAt() string`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *UpgradePath) GetReleasedAtOk() (*string, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *UpgradePath) SetReleasedAt(v string)`

SetReleasedAt sets ReleasedAt field to given value.


### GetScheduledCount

`func (o *UpgradePath) GetScheduledCount() int64`

GetScheduledCount returns the ScheduledCount field if non-nil, zero value otherwise.

### GetScheduledCountOk

`func (o *UpgradePath) GetScheduledCountOk() (*int64, bool)`

GetScheduledCountOk returns a tuple with the ScheduledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCount

`func (o *UpgradePath) SetScheduledCount(v int64)`

SetScheduledCount sets ScheduledCount field to given value.

### HasScheduledCount

`func (o *UpgradePath) HasScheduledCount() bool`

HasScheduledCount returns a boolean if a field has been set.

### GetServiceId

`func (o *UpgradePath) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpgradePath) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpgradePath) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSkippedCount

`func (o *UpgradePath) GetSkippedCount() int64`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *UpgradePath) GetSkippedCountOk() (*int64, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *UpgradePath) SetSkippedCount(v int64)`

SetSkippedCount sets SkippedCount field to given value.


### GetSourceVersion

`func (o *UpgradePath) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *UpgradePath) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *UpgradePath) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetStatus

`func (o *UpgradePath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradePath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradePath) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetVersion

`func (o *UpgradePath) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *UpgradePath) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *UpgradePath) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetTotalCount

`func (o *UpgradePath) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UpgradePath) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UpgradePath) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetType

`func (o *UpgradePath) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradePath) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradePath) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *UpgradePath) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpgradePath) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpgradePath) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpgradePathId

`func (o *UpgradePath) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *UpgradePath) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *UpgradePath) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


