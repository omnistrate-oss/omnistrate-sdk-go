/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InstanceUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceUpgrade{}

// InstanceUpgrade The instance upgrade status in the context of an upgrade path
type InstanceUpgrade struct {
	// The name of the cloud provider that the instance is running on.
	CloudProviderName string `json:"cloudProviderName"`
	// The region of the cloud provider that the instance is running in.
	CloudProviderRegion string `json:"cloudProviderRegion"`
	// The timestamp when the instance was created.
	CreatedAt string `json:"createdAt"`
	// The heath status of the instance.
	HealthStatus *string `json:"healthStatus,omitempty"`
	// The instance ID.
	InstanceId string `json:"instanceId"`
	// Instance lifecycle status.
	LifecycleStatus string `json:"lifecycleStatus"`
	// The managed resource type of the top-level resource of the instance.
	ManagedResourceType *string `json:"managedResourceType,omitempty"`
	// The name of the organization that owns the instance.
	OrgName string `json:"orgName"`
	// The name of the top-level resource of the instance.
	ResourceName string `json:"resourceName"`
	// The status of the instance upgrade.
	Status string `json:"status"`
	// The timestamp when the instance was updated.
	UpdatedAt string `json:"updatedAt"`
	// The timestamp when the upgrade ended.
	UpgradeEndTime *string `json:"upgradeEndTime,omitempty"`
	// The timestamp when the upgrade started.
	UpgradeStartTime *string `json:"upgradeStartTime,omitempty"`
	// The workflow ID of the instance upgrade.
	WorkflowID string `json:"workflowID"`
}

type _InstanceUpgrade InstanceUpgrade

// NewInstanceUpgrade instantiates a new InstanceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceUpgrade(cloudProviderName string, cloudProviderRegion string, createdAt string, instanceId string, lifecycleStatus string, orgName string, resourceName string, status string, updatedAt string, workflowID string) *InstanceUpgrade {
	this := InstanceUpgrade{}
	this.CloudProviderName = cloudProviderName
	this.CloudProviderRegion = cloudProviderRegion
	this.CreatedAt = createdAt
	this.InstanceId = instanceId
	this.LifecycleStatus = lifecycleStatus
	this.OrgName = orgName
	this.ResourceName = resourceName
	this.Status = status
	this.UpdatedAt = updatedAt
	this.WorkflowID = workflowID
	return &this
}

// NewInstanceUpgradeWithDefaults instantiates a new InstanceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceUpgradeWithDefaults() *InstanceUpgrade {
	this := InstanceUpgrade{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *InstanceUpgrade) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *InstanceUpgrade) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCloudProviderRegion returns the CloudProviderRegion field value
func (o *InstanceUpgrade) GetCloudProviderRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderRegion
}

// GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetCloudProviderRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderRegion, true
}

// SetCloudProviderRegion sets field value
func (o *InstanceUpgrade) SetCloudProviderRegion(v string) {
	o.CloudProviderRegion = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InstanceUpgrade) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InstanceUpgrade) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetHealthStatus returns the HealthStatus field value if set, zero value otherwise.
func (o *InstanceUpgrade) GetHealthStatus() string {
	if o == nil || IsNil(o.HealthStatus) {
		var ret string
		return ret
	}
	return *o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetHealthStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HealthStatus) {
		return nil, false
	}
	return o.HealthStatus, true
}

// HasHealthStatus returns a boolean if a field has been set.
func (o *InstanceUpgrade) HasHealthStatus() bool {
	if o != nil && !IsNil(o.HealthStatus) {
		return true
	}

	return false
}

// SetHealthStatus gets a reference to the given string and assigns it to the HealthStatus field.
func (o *InstanceUpgrade) SetHealthStatus(v string) {
	o.HealthStatus = &v
}

// GetInstanceId returns the InstanceId field value
func (o *InstanceUpgrade) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *InstanceUpgrade) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetLifecycleStatus returns the LifecycleStatus field value
func (o *InstanceUpgrade) GetLifecycleStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LifecycleStatus
}

// GetLifecycleStatusOk returns a tuple with the LifecycleStatus field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetLifecycleStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifecycleStatus, true
}

// SetLifecycleStatus sets field value
func (o *InstanceUpgrade) SetLifecycleStatus(v string) {
	o.LifecycleStatus = v
}

// GetManagedResourceType returns the ManagedResourceType field value if set, zero value otherwise.
func (o *InstanceUpgrade) GetManagedResourceType() string {
	if o == nil || IsNil(o.ManagedResourceType) {
		var ret string
		return ret
	}
	return *o.ManagedResourceType
}

// GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetManagedResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedResourceType) {
		return nil, false
	}
	return o.ManagedResourceType, true
}

// HasManagedResourceType returns a boolean if a field has been set.
func (o *InstanceUpgrade) HasManagedResourceType() bool {
	if o != nil && !IsNil(o.ManagedResourceType) {
		return true
	}

	return false
}

// SetManagedResourceType gets a reference to the given string and assigns it to the ManagedResourceType field.
func (o *InstanceUpgrade) SetManagedResourceType(v string) {
	o.ManagedResourceType = &v
}

// GetOrgName returns the OrgName field value
func (o *InstanceUpgrade) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *InstanceUpgrade) SetOrgName(v string) {
	o.OrgName = v
}

// GetResourceName returns the ResourceName field value
func (o *InstanceUpgrade) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *InstanceUpgrade) SetResourceName(v string) {
	o.ResourceName = v
}

// GetStatus returns the Status field value
func (o *InstanceUpgrade) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InstanceUpgrade) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InstanceUpgrade) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InstanceUpgrade) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUpgradeEndTime returns the UpgradeEndTime field value if set, zero value otherwise.
func (o *InstanceUpgrade) GetUpgradeEndTime() string {
	if o == nil || IsNil(o.UpgradeEndTime) {
		var ret string
		return ret
	}
	return *o.UpgradeEndTime
}

// GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetUpgradeEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeEndTime) {
		return nil, false
	}
	return o.UpgradeEndTime, true
}

// HasUpgradeEndTime returns a boolean if a field has been set.
func (o *InstanceUpgrade) HasUpgradeEndTime() bool {
	if o != nil && !IsNil(o.UpgradeEndTime) {
		return true
	}

	return false
}

// SetUpgradeEndTime gets a reference to the given string and assigns it to the UpgradeEndTime field.
func (o *InstanceUpgrade) SetUpgradeEndTime(v string) {
	o.UpgradeEndTime = &v
}

// GetUpgradeStartTime returns the UpgradeStartTime field value if set, zero value otherwise.
func (o *InstanceUpgrade) GetUpgradeStartTime() string {
	if o == nil || IsNil(o.UpgradeStartTime) {
		var ret string
		return ret
	}
	return *o.UpgradeStartTime
}

// GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetUpgradeStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeStartTime) {
		return nil, false
	}
	return o.UpgradeStartTime, true
}

// HasUpgradeStartTime returns a boolean if a field has been set.
func (o *InstanceUpgrade) HasUpgradeStartTime() bool {
	if o != nil && !IsNil(o.UpgradeStartTime) {
		return true
	}

	return false
}

// SetUpgradeStartTime gets a reference to the given string and assigns it to the UpgradeStartTime field.
func (o *InstanceUpgrade) SetUpgradeStartTime(v string) {
	o.UpgradeStartTime = &v
}

// GetWorkflowID returns the WorkflowID field value
func (o *InstanceUpgrade) GetWorkflowID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowID
}

// GetWorkflowIDOk returns a tuple with the WorkflowID field value
// and a boolean to check if the value has been set.
func (o *InstanceUpgrade) GetWorkflowIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowID, true
}

// SetWorkflowID sets field value
func (o *InstanceUpgrade) SetWorkflowID(v string) {
	o.WorkflowID = v
}

func (o InstanceUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["cloudProviderRegion"] = o.CloudProviderRegion
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.HealthStatus) {
		toSerialize["healthStatus"] = o.HealthStatus
	}
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["lifecycleStatus"] = o.LifecycleStatus
	if !IsNil(o.ManagedResourceType) {
		toSerialize["managedResourceType"] = o.ManagedResourceType
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["status"] = o.Status
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.UpgradeEndTime) {
		toSerialize["upgradeEndTime"] = o.UpgradeEndTime
	}
	if !IsNil(o.UpgradeStartTime) {
		toSerialize["upgradeStartTime"] = o.UpgradeStartTime
	}
	toSerialize["workflowID"] = o.WorkflowID
	return toSerialize, nil
}

func (o *InstanceUpgrade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"cloudProviderRegion",
		"createdAt",
		"instanceId",
		"lifecycleStatus",
		"orgName",
		"resourceName",
		"status",
		"updatedAt",
		"workflowID",
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

	varInstanceUpgrade := _InstanceUpgrade{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstanceUpgrade)

	if err != nil {
		return err
	}

	*o = InstanceUpgrade(varInstanceUpgrade)

	return err
}

type NullableInstanceUpgrade struct {
	value *InstanceUpgrade
	isSet bool
}

func (v NullableInstanceUpgrade) Get() *InstanceUpgrade {
	return v.value
}

func (v *NullableInstanceUpgrade) Set(val *InstanceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceUpgrade(val *InstanceUpgrade) *NullableInstanceUpgrade {
	return &NullableInstanceUpgrade{value: val, isSet: true}
}

func (v NullableInstanceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


