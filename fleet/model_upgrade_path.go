/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the UpgradePath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradePath{}

// UpgradePath struct for UpgradePath
type UpgradePath struct {
	// The timestamp when the upgrade path was completed.
	CompletedAt *string `json:"completedAt,omitempty"`
	// The number of instances that have completed the upgrade.
	CompletedCount int64 `json:"completedCount"`
	// The timestamp when the upgrade path was created.
	CreatedAt string `json:"createdAt"`
	// The name of the user who created the upgrade path.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The number of instances that have failed the upgrade.
	FailedCount int64 `json:"failedCount"`
	// The number of instances that are in progress of the upgrade.
	InProgressCount int64 `json:"inProgressCount"`
	// The name of the user who created the upgrade path.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// The action to perform on an ongoing resource workflow
	LastRequestedAction *string `json:"lastRequestedAction,omitempty"`
	// Whether to notify the end customer about the upgrade progress.
	NotifyCustomer *bool `json:"notifyCustomer,omitempty"`
	// The number of instances that are pending the upgrade.
	PendingCount int64 `json:"pendingCount"`
	// The date when the upgrade was/is planned to be executed.
	PlannedExecutionDate *string `json:"plannedExecutionDate,omitempty"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// The timestamp when the upgrade path was released.
	ReleasedAt string `json:"releasedAt"`
	// The number of instances that are scheduled to upgrade.
	ScheduledCount *int64 `json:"scheduledCount,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The number of instances that are skipped the upgrade.
	SkippedCount int64 `json:"skippedCount"`
	// The source version of the upgrade path.
	SourceVersion string `json:"sourceVersion"`
	// The status of the upgrade path.
	Status string `json:"status"`
	// The target version of the upgrade path.
	TargetVersion string `json:"targetVersion"`
	// The total number of instances that are eligible for the upgrade.
	TotalCount int64 `json:"totalCount"`
	// The type of the upgrade path.
	Type string `json:"type"`
	// The timestamp when the upgrade path was updated.
	UpdatedAt string `json:"updatedAt"`
	// ID of an Upgrade Path
	UpgradePathId string `json:"upgradePathId"`
	AdditionalProperties map[string]interface{}
}

type _UpgradePath UpgradePath

// NewUpgradePath instantiates a new UpgradePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradePath(completedCount int64, createdAt string, failedCount int64, inProgressCount int64, pendingCount int64, productTierId string, releasedAt string, serviceId string, skippedCount int64, sourceVersion string, status string, targetVersion string, totalCount int64, type_ string, updatedAt string, upgradePathId string) *UpgradePath {
	this := UpgradePath{}
	this.CompletedCount = completedCount
	this.CreatedAt = createdAt
	this.FailedCount = failedCount
	this.InProgressCount = inProgressCount
	this.PendingCount = pendingCount
	this.ProductTierId = productTierId
	this.ReleasedAt = releasedAt
	this.ServiceId = serviceId
	this.SkippedCount = skippedCount
	this.SourceVersion = sourceVersion
	this.Status = status
	this.TargetVersion = targetVersion
	this.TotalCount = totalCount
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.UpgradePathId = upgradePathId
	return &this
}

// NewUpgradePathWithDefaults instantiates a new UpgradePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradePathWithDefaults() *UpgradePath {
	this := UpgradePath{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *UpgradePath) GetCompletedAt() string {
	if o == nil || IsNil(o.CompletedAt) {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetCompletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *UpgradePath) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *UpgradePath) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetCompletedCount returns the CompletedCount field value
func (o *UpgradePath) GetCompletedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CompletedCount
}

// GetCompletedCountOk returns a tuple with the CompletedCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetCompletedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedCount, true
}

// SetCompletedCount sets field value
func (o *UpgradePath) SetCompletedCount(v int64) {
	o.CompletedCount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UpgradePath) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UpgradePath) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *UpgradePath) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UpgradePath) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *UpgradePath) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetFailedCount returns the FailedCount field value
func (o *UpgradePath) GetFailedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetFailedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedCount, true
}

// SetFailedCount sets field value
func (o *UpgradePath) SetFailedCount(v int64) {
	o.FailedCount = v
}

// GetInProgressCount returns the InProgressCount field value
func (o *UpgradePath) GetInProgressCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InProgressCount
}

// GetInProgressCountOk returns a tuple with the InProgressCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetInProgressCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProgressCount, true
}

// SetInProgressCount sets field value
func (o *UpgradePath) SetInProgressCount(v int64) {
	o.InProgressCount = v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *UpgradePath) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *UpgradePath) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *UpgradePath) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastRequestedAction returns the LastRequestedAction field value if set, zero value otherwise.
func (o *UpgradePath) GetLastRequestedAction() string {
	if o == nil || IsNil(o.LastRequestedAction) {
		var ret string
		return ret
	}
	return *o.LastRequestedAction
}

// GetLastRequestedActionOk returns a tuple with the LastRequestedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetLastRequestedActionOk() (*string, bool) {
	if o == nil || IsNil(o.LastRequestedAction) {
		return nil, false
	}
	return o.LastRequestedAction, true
}

// HasLastRequestedAction returns a boolean if a field has been set.
func (o *UpgradePath) HasLastRequestedAction() bool {
	if o != nil && !IsNil(o.LastRequestedAction) {
		return true
	}

	return false
}

// SetLastRequestedAction gets a reference to the given string and assigns it to the LastRequestedAction field.
func (o *UpgradePath) SetLastRequestedAction(v string) {
	o.LastRequestedAction = &v
}

// GetNotifyCustomer returns the NotifyCustomer field value if set, zero value otherwise.
func (o *UpgradePath) GetNotifyCustomer() bool {
	if o == nil || IsNil(o.NotifyCustomer) {
		var ret bool
		return ret
	}
	return *o.NotifyCustomer
}

// GetNotifyCustomerOk returns a tuple with the NotifyCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetNotifyCustomerOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyCustomer) {
		return nil, false
	}
	return o.NotifyCustomer, true
}

// HasNotifyCustomer returns a boolean if a field has been set.
func (o *UpgradePath) HasNotifyCustomer() bool {
	if o != nil && !IsNil(o.NotifyCustomer) {
		return true
	}

	return false
}

// SetNotifyCustomer gets a reference to the given bool and assigns it to the NotifyCustomer field.
func (o *UpgradePath) SetNotifyCustomer(v bool) {
	o.NotifyCustomer = &v
}

// GetPendingCount returns the PendingCount field value
func (o *UpgradePath) GetPendingCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetPendingCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingCount, true
}

// SetPendingCount sets field value
func (o *UpgradePath) SetPendingCount(v int64) {
	o.PendingCount = v
}

// GetPlannedExecutionDate returns the PlannedExecutionDate field value if set, zero value otherwise.
func (o *UpgradePath) GetPlannedExecutionDate() string {
	if o == nil || IsNil(o.PlannedExecutionDate) {
		var ret string
		return ret
	}
	return *o.PlannedExecutionDate
}

// GetPlannedExecutionDateOk returns a tuple with the PlannedExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetPlannedExecutionDateOk() (*string, bool) {
	if o == nil || IsNil(o.PlannedExecutionDate) {
		return nil, false
	}
	return o.PlannedExecutionDate, true
}

// HasPlannedExecutionDate returns a boolean if a field has been set.
func (o *UpgradePath) HasPlannedExecutionDate() bool {
	if o != nil && !IsNil(o.PlannedExecutionDate) {
		return true
	}

	return false
}

// SetPlannedExecutionDate gets a reference to the given string and assigns it to the PlannedExecutionDate field.
func (o *UpgradePath) SetPlannedExecutionDate(v string) {
	o.PlannedExecutionDate = &v
}

// GetProductTierId returns the ProductTierId field value
func (o *UpgradePath) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *UpgradePath) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetReleasedAt returns the ReleasedAt field value
func (o *UpgradePath) GetReleasedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleasedAt
}

// GetReleasedAtOk returns a tuple with the ReleasedAt field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetReleasedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleasedAt, true
}

// SetReleasedAt sets field value
func (o *UpgradePath) SetReleasedAt(v string) {
	o.ReleasedAt = v
}

// GetScheduledCount returns the ScheduledCount field value if set, zero value otherwise.
func (o *UpgradePath) GetScheduledCount() int64 {
	if o == nil || IsNil(o.ScheduledCount) {
		var ret int64
		return ret
	}
	return *o.ScheduledCount
}

// GetScheduledCountOk returns a tuple with the ScheduledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetScheduledCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ScheduledCount) {
		return nil, false
	}
	return o.ScheduledCount, true
}

// HasScheduledCount returns a boolean if a field has been set.
func (o *UpgradePath) HasScheduledCount() bool {
	if o != nil && !IsNil(o.ScheduledCount) {
		return true
	}

	return false
}

// SetScheduledCount gets a reference to the given int64 and assigns it to the ScheduledCount field.
func (o *UpgradePath) SetScheduledCount(v int64) {
	o.ScheduledCount = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpgradePath) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpgradePath) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSkippedCount returns the SkippedCount field value
func (o *UpgradePath) GetSkippedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SkippedCount
}

// GetSkippedCountOk returns a tuple with the SkippedCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetSkippedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkippedCount, true
}

// SetSkippedCount sets field value
func (o *UpgradePath) SetSkippedCount(v int64) {
	o.SkippedCount = v
}

// GetSourceVersion returns the SourceVersion field value
func (o *UpgradePath) GetSourceVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceVersion
}

// GetSourceVersionOk returns a tuple with the SourceVersion field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetSourceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceVersion, true
}

// SetSourceVersion sets field value
func (o *UpgradePath) SetSourceVersion(v string) {
	o.SourceVersion = v
}

// GetStatus returns the Status field value
func (o *UpgradePath) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpgradePath) SetStatus(v string) {
	o.Status = v
}

// GetTargetVersion returns the TargetVersion field value
func (o *UpgradePath) GetTargetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetTargetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetVersion, true
}

// SetTargetVersion sets field value
func (o *UpgradePath) SetTargetVersion(v string) {
	o.TargetVersion = v
}

// GetTotalCount returns the TotalCount field value
func (o *UpgradePath) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *UpgradePath) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetType returns the Type field value
func (o *UpgradePath) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpgradePath) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UpgradePath) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UpgradePath) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUpgradePathId returns the UpgradePathId field value
func (o *UpgradePath) GetUpgradePathId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpgradePathId
}

// GetUpgradePathIdOk returns a tuple with the UpgradePathId field value
// and a boolean to check if the value has been set.
func (o *UpgradePath) GetUpgradePathIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradePathId, true
}

// SetUpgradePathId sets field value
func (o *UpgradePath) SetUpgradePathId(v string) {
	o.UpgradePathId = v
}

func (o UpgradePath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradePath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	toSerialize["completedCount"] = o.CompletedCount
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	toSerialize["failedCount"] = o.FailedCount
	toSerialize["inProgressCount"] = o.InProgressCount
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastRequestedAction) {
		toSerialize["lastRequestedAction"] = o.LastRequestedAction
	}
	if !IsNil(o.NotifyCustomer) {
		toSerialize["notifyCustomer"] = o.NotifyCustomer
	}
	toSerialize["pendingCount"] = o.PendingCount
	if !IsNil(o.PlannedExecutionDate) {
		toSerialize["plannedExecutionDate"] = o.PlannedExecutionDate
	}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["releasedAt"] = o.ReleasedAt
	if !IsNil(o.ScheduledCount) {
		toSerialize["scheduledCount"] = o.ScheduledCount
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["skippedCount"] = o.SkippedCount
	toSerialize["sourceVersion"] = o.SourceVersion
	toSerialize["status"] = o.Status
	toSerialize["targetVersion"] = o.TargetVersion
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["type"] = o.Type
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["upgradePathId"] = o.UpgradePathId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpgradePath) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completedCount",
		"createdAt",
		"failedCount",
		"inProgressCount",
		"pendingCount",
		"productTierId",
		"releasedAt",
		"serviceId",
		"skippedCount",
		"sourceVersion",
		"status",
		"targetVersion",
		"totalCount",
		"type",
		"updatedAt",
		"upgradePathId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpgradePath := _UpgradePath{}

	err = json.Unmarshal(data, &varUpgradePath)

	if err != nil {
		return err
	}

	*o = UpgradePath(varUpgradePath)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completedAt")
		delete(additionalProperties, "completedCount")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "failedCount")
		delete(additionalProperties, "inProgressCount")
		delete(additionalProperties, "lastModifiedBy")
		delete(additionalProperties, "lastRequestedAction")
		delete(additionalProperties, "notifyCustomer")
		delete(additionalProperties, "pendingCount")
		delete(additionalProperties, "plannedExecutionDate")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "releasedAt")
		delete(additionalProperties, "scheduledCount")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "skippedCount")
		delete(additionalProperties, "sourceVersion")
		delete(additionalProperties, "status")
		delete(additionalProperties, "targetVersion")
		delete(additionalProperties, "totalCount")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "upgradePathId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpgradePath struct {
	value *UpgradePath
	isSet bool
}

func (v NullableUpgradePath) Get() *UpgradePath {
	return v.value
}

func (v *NullableUpgradePath) Set(val *UpgradePath) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradePath) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradePath(val *UpgradePath) *NullableUpgradePath {
	return &NullableUpgradePath{value: val, isSet: true}
}

func (v NullableUpgradePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


