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

// checks if the InstanceHealthSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceHealthSummary{}

// InstanceHealthSummary struct for InstanceHealthSummary
type InstanceHealthSummary struct {
	// The ID of the instance
	InstanceID string `json:"instanceID"`
	// The lifecycle status of the instance
	LifeCycleStatus string `json:"lifeCycleStatus"`
	// The health summary of the resources by resource ID
	ResourcesHealth map[string]ResourceHealthSummary `json:"resourcesHealth"`
	// The status of the instance
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _InstanceHealthSummary InstanceHealthSummary

// NewInstanceHealthSummary instantiates a new InstanceHealthSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceHealthSummary(instanceID string, lifeCycleStatus string, resourcesHealth map[string]ResourceHealthSummary, status string) *InstanceHealthSummary {
	this := InstanceHealthSummary{}
	this.InstanceID = instanceID
	this.LifeCycleStatus = lifeCycleStatus
	this.ResourcesHealth = resourcesHealth
	this.Status = status
	return &this
}

// NewInstanceHealthSummaryWithDefaults instantiates a new InstanceHealthSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceHealthSummaryWithDefaults() *InstanceHealthSummary {
	this := InstanceHealthSummary{}
	return &this
}

// GetInstanceID returns the InstanceID field value
func (o *InstanceHealthSummary) GetInstanceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value
// and a boolean to check if the value has been set.
func (o *InstanceHealthSummary) GetInstanceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceID, true
}

// SetInstanceID sets field value
func (o *InstanceHealthSummary) SetInstanceID(v string) {
	o.InstanceID = v
}

// GetLifeCycleStatus returns the LifeCycleStatus field value
func (o *InstanceHealthSummary) GetLifeCycleStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LifeCycleStatus
}

// GetLifeCycleStatusOk returns a tuple with the LifeCycleStatus field value
// and a boolean to check if the value has been set.
func (o *InstanceHealthSummary) GetLifeCycleStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifeCycleStatus, true
}

// SetLifeCycleStatus sets field value
func (o *InstanceHealthSummary) SetLifeCycleStatus(v string) {
	o.LifeCycleStatus = v
}

// GetResourcesHealth returns the ResourcesHealth field value
func (o *InstanceHealthSummary) GetResourcesHealth() map[string]ResourceHealthSummary {
	if o == nil {
		var ret map[string]ResourceHealthSummary
		return ret
	}

	return o.ResourcesHealth
}

// GetResourcesHealthOk returns a tuple with the ResourcesHealth field value
// and a boolean to check if the value has been set.
func (o *InstanceHealthSummary) GetResourcesHealthOk() (*map[string]ResourceHealthSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcesHealth, true
}

// SetResourcesHealth sets field value
func (o *InstanceHealthSummary) SetResourcesHealth(v map[string]ResourceHealthSummary) {
	o.ResourcesHealth = v
}

// GetStatus returns the Status field value
func (o *InstanceHealthSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InstanceHealthSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InstanceHealthSummary) SetStatus(v string) {
	o.Status = v
}

func (o InstanceHealthSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceHealthSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceID"] = o.InstanceID
	toSerialize["lifeCycleStatus"] = o.LifeCycleStatus
	toSerialize["resourcesHealth"] = o.ResourcesHealth
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstanceHealthSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceID",
		"lifeCycleStatus",
		"resourcesHealth",
		"status",
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

	varInstanceHealthSummary := _InstanceHealthSummary{}

	err = json.Unmarshal(data, &varInstanceHealthSummary)

	if err != nil {
		return err
	}

	*o = InstanceHealthSummary(varInstanceHealthSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instanceID")
		delete(additionalProperties, "lifeCycleStatus")
		delete(additionalProperties, "resourcesHealth")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstanceHealthSummary struct {
	value *InstanceHealthSummary
	isSet bool
}

func (v NullableInstanceHealthSummary) Get() *InstanceHealthSummary {
	return v.value
}

func (v *NullableInstanceHealthSummary) Set(val *InstanceHealthSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceHealthSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceHealthSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceHealthSummary(val *InstanceHealthSummary) *NullableInstanceHealthSummary {
	return &NullableInstanceHealthSummary{value: val, isSet: true}
}

func (v NullableInstanceHealthSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceHealthSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


