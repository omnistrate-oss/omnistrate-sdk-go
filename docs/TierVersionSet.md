# TierVersionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**BaseVersion** | **string** | The base version of the version set. | 
**CreatedAt** | **string** | The timestamp when the version set was created. | 
**CreatedBy** | Pointer to **string** | The name of the user who created the version set. | [optional] 
**Description** | Pointer to **string** | A brief description of the product-tier version set. | [optional] 
**EnabledFeatures** | [**[]ProductTierFeatureDetail**](ProductTierFeatureDetail.md) | The features that are enabled for this product tier, including scope details and configuration | 
**Features** | **map[string]bool** | The features that are enabled / disabled for this product tier | 
**InstanceCount** | Pointer to **int64** | The number of instances that are currently running this version set. | [optional] 
**LatestUpgradePathId** | Pointer to **string** | The ID of the latest upgrade path away from the tier-version set. | [optional] 
**Name** | Pointer to **string** | The name of the product-tier version set. | [optional] 
**ParentVersion** | Pointer to **string** | The parent version of this release. | [optional] 
**ProductTierId** | **string** | The product tier ID that this version set belongs to. | 
**ReleasedAt** | **string** | The timestamp when the version set was released. | 
**ReleasedBy** | Pointer to **string** | The name of the user who released the version set. | [optional] 
**Resources** | Pointer to [**[]ResourceSummary**](ResourceSummary.md) | List of resources that are part of this version set. | [optional] 
**ServiceId** | **string** | ID of the Service | 
**ServiceModelId** | **string** | The service model ID that this version set belongs to. | 
**Status** | **string** | The tier version set status. | 
**Type** | **string** | The version-set type of the product-tier. | 
**UpdatedAt** | **string** | The timestamp when the version set was updated. | 
**Version** | **string** | The version number for the specific version set. | 

## Methods

### NewTierVersionSet

`func NewTierVersionSet(baseVersion string, createdAt string, enabledFeatures []ProductTierFeatureDetail, features map[string]bool, productTierId string, releasedAt string, serviceId string, serviceModelId string, status string, type_ string, updatedAt string, version string, ) *TierVersionSet`

NewTierVersionSet instantiates a new TierVersionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierVersionSetWithDefaults

`func NewTierVersionSetWithDefaults() *TierVersionSet`

NewTierVersionSetWithDefaults instantiates a new TierVersionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *TierVersionSet) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *TierVersionSet) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *TierVersionSet) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *TierVersionSet) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetBaseVersion

`func (o *TierVersionSet) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *TierVersionSet) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *TierVersionSet) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.


### GetCreatedAt

`func (o *TierVersionSet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TierVersionSet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TierVersionSet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TierVersionSet) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TierVersionSet) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TierVersionSet) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TierVersionSet) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *TierVersionSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TierVersionSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TierVersionSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TierVersionSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabledFeatures

`func (o *TierVersionSet) GetEnabledFeatures() []ProductTierFeatureDetail`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *TierVersionSet) GetEnabledFeaturesOk() (*[]ProductTierFeatureDetail, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *TierVersionSet) SetEnabledFeatures(v []ProductTierFeatureDetail)`

SetEnabledFeatures sets EnabledFeatures field to given value.


### GetFeatures

`func (o *TierVersionSet) GetFeatures() map[string]bool`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TierVersionSet) GetFeaturesOk() (*map[string]bool, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TierVersionSet) SetFeatures(v map[string]bool)`

SetFeatures sets Features field to given value.


### GetInstanceCount

`func (o *TierVersionSet) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *TierVersionSet) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *TierVersionSet) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *TierVersionSet) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetLatestUpgradePathId

`func (o *TierVersionSet) GetLatestUpgradePathId() string`

GetLatestUpgradePathId returns the LatestUpgradePathId field if non-nil, zero value otherwise.

### GetLatestUpgradePathIdOk

`func (o *TierVersionSet) GetLatestUpgradePathIdOk() (*string, bool)`

GetLatestUpgradePathIdOk returns a tuple with the LatestUpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpgradePathId

`func (o *TierVersionSet) SetLatestUpgradePathId(v string)`

SetLatestUpgradePathId sets LatestUpgradePathId field to given value.

### HasLatestUpgradePathId

`func (o *TierVersionSet) HasLatestUpgradePathId() bool`

HasLatestUpgradePathId returns a boolean if a field has been set.

### GetName

`func (o *TierVersionSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierVersionSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierVersionSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TierVersionSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentVersion

`func (o *TierVersionSet) GetParentVersion() string`

GetParentVersion returns the ParentVersion field if non-nil, zero value otherwise.

### GetParentVersionOk

`func (o *TierVersionSet) GetParentVersionOk() (*string, bool)`

GetParentVersionOk returns a tuple with the ParentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersion

`func (o *TierVersionSet) SetParentVersion(v string)`

SetParentVersion sets ParentVersion field to given value.

### HasParentVersion

`func (o *TierVersionSet) HasParentVersion() bool`

HasParentVersion returns a boolean if a field has been set.

### GetProductTierId

`func (o *TierVersionSet) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *TierVersionSet) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *TierVersionSet) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetReleasedAt

`func (o *TierVersionSet) GetReleasedAt() string`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *TierVersionSet) GetReleasedAtOk() (*string, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *TierVersionSet) SetReleasedAt(v string)`

SetReleasedAt sets ReleasedAt field to given value.


### GetReleasedBy

`func (o *TierVersionSet) GetReleasedBy() string`

GetReleasedBy returns the ReleasedBy field if non-nil, zero value otherwise.

### GetReleasedByOk

`func (o *TierVersionSet) GetReleasedByOk() (*string, bool)`

GetReleasedByOk returns a tuple with the ReleasedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedBy

`func (o *TierVersionSet) SetReleasedBy(v string)`

SetReleasedBy sets ReleasedBy field to given value.

### HasReleasedBy

`func (o *TierVersionSet) HasReleasedBy() bool`

HasReleasedBy returns a boolean if a field has been set.

### GetResources

`func (o *TierVersionSet) GetResources() []ResourceSummary`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TierVersionSet) GetResourcesOk() (*[]ResourceSummary, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TierVersionSet) SetResources(v []ResourceSummary)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TierVersionSet) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetServiceId

`func (o *TierVersionSet) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TierVersionSet) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TierVersionSet) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *TierVersionSet) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *TierVersionSet) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *TierVersionSet) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetStatus

`func (o *TierVersionSet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TierVersionSet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TierVersionSet) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *TierVersionSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TierVersionSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TierVersionSet) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *TierVersionSet) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TierVersionSet) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TierVersionSet) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *TierVersionSet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TierVersionSet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TierVersionSet) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


