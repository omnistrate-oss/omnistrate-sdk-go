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

// checks if the OneOffPatchResourceInstanceRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneOffPatchResourceInstanceRequest2{}

// OneOffPatchResourceInstanceRequest2 struct for OneOffPatchResourceInstanceRequest2
type OneOffPatchResourceInstanceRequest2 struct {
	// The resource ID.
	ResourceId string `json:"resourceId"`
	// The resource override configuration for one-off patching.
	ResourceOverrideConfiguration *map[string]ResourceOneOffPatchConfigurationOverride `json:"resourceOverrideConfiguration,omitempty"`
	// The target resource version.
	TargetTierVersion *string `json:"targetTierVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OneOffPatchResourceInstanceRequest2 OneOffPatchResourceInstanceRequest2

// NewOneOffPatchResourceInstanceRequest2 instantiates a new OneOffPatchResourceInstanceRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneOffPatchResourceInstanceRequest2(resourceId string) *OneOffPatchResourceInstanceRequest2 {
	this := OneOffPatchResourceInstanceRequest2{}
	this.ResourceId = resourceId
	return &this
}

// NewOneOffPatchResourceInstanceRequest2WithDefaults instantiates a new OneOffPatchResourceInstanceRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneOffPatchResourceInstanceRequest2WithDefaults() *OneOffPatchResourceInstanceRequest2 {
	this := OneOffPatchResourceInstanceRequest2{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *OneOffPatchResourceInstanceRequest2) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *OneOffPatchResourceInstanceRequest2) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *OneOffPatchResourceInstanceRequest2) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceOverrideConfiguration returns the ResourceOverrideConfiguration field value if set, zero value otherwise.
func (o *OneOffPatchResourceInstanceRequest2) GetResourceOverrideConfiguration() map[string]ResourceOneOffPatchConfigurationOverride {
	if o == nil || IsNil(o.ResourceOverrideConfiguration) {
		var ret map[string]ResourceOneOffPatchConfigurationOverride
		return ret
	}
	return *o.ResourceOverrideConfiguration
}

// GetResourceOverrideConfigurationOk returns a tuple with the ResourceOverrideConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneOffPatchResourceInstanceRequest2) GetResourceOverrideConfigurationOk() (*map[string]ResourceOneOffPatchConfigurationOverride, bool) {
	if o == nil || IsNil(o.ResourceOverrideConfiguration) {
		return nil, false
	}
	return o.ResourceOverrideConfiguration, true
}

// HasResourceOverrideConfiguration returns a boolean if a field has been set.
func (o *OneOffPatchResourceInstanceRequest2) HasResourceOverrideConfiguration() bool {
	if o != nil && !IsNil(o.ResourceOverrideConfiguration) {
		return true
	}

	return false
}

// SetResourceOverrideConfiguration gets a reference to the given map[string]ResourceOneOffPatchConfigurationOverride and assigns it to the ResourceOverrideConfiguration field.
func (o *OneOffPatchResourceInstanceRequest2) SetResourceOverrideConfiguration(v map[string]ResourceOneOffPatchConfigurationOverride) {
	o.ResourceOverrideConfiguration = &v
}

// GetTargetTierVersion returns the TargetTierVersion field value if set, zero value otherwise.
func (o *OneOffPatchResourceInstanceRequest2) GetTargetTierVersion() string {
	if o == nil || IsNil(o.TargetTierVersion) {
		var ret string
		return ret
	}
	return *o.TargetTierVersion
}

// GetTargetTierVersionOk returns a tuple with the TargetTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneOffPatchResourceInstanceRequest2) GetTargetTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetTierVersion) {
		return nil, false
	}
	return o.TargetTierVersion, true
}

// HasTargetTierVersion returns a boolean if a field has been set.
func (o *OneOffPatchResourceInstanceRequest2) HasTargetTierVersion() bool {
	if o != nil && !IsNil(o.TargetTierVersion) {
		return true
	}

	return false
}

// SetTargetTierVersion gets a reference to the given string and assigns it to the TargetTierVersion field.
func (o *OneOffPatchResourceInstanceRequest2) SetTargetTierVersion(v string) {
	o.TargetTierVersion = &v
}

func (o OneOffPatchResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OneOffPatchResourceInstanceRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.ResourceOverrideConfiguration) {
		toSerialize["resourceOverrideConfiguration"] = o.ResourceOverrideConfiguration
	}
	if !IsNil(o.TargetTierVersion) {
		toSerialize["targetTierVersion"] = o.TargetTierVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OneOffPatchResourceInstanceRequest2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
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

	varOneOffPatchResourceInstanceRequest2 := _OneOffPatchResourceInstanceRequest2{}

	err = json.Unmarshal(data, &varOneOffPatchResourceInstanceRequest2)

	if err != nil {
		return err
	}

	*o = OneOffPatchResourceInstanceRequest2(varOneOffPatchResourceInstanceRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceOverrideConfiguration")
		delete(additionalProperties, "targetTierVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOneOffPatchResourceInstanceRequest2 struct {
	value *OneOffPatchResourceInstanceRequest2
	isSet bool
}

func (v NullableOneOffPatchResourceInstanceRequest2) Get() *OneOffPatchResourceInstanceRequest2 {
	return v.value
}

func (v *NullableOneOffPatchResourceInstanceRequest2) Set(val *OneOffPatchResourceInstanceRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableOneOffPatchResourceInstanceRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableOneOffPatchResourceInstanceRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneOffPatchResourceInstanceRequest2(val *OneOffPatchResourceInstanceRequest2) *NullableOneOffPatchResourceInstanceRequest2 {
	return &NullableOneOffPatchResourceInstanceRequest2{value: val, isSet: true}
}

func (v NullableOneOffPatchResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneOffPatchResourceInstanceRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


