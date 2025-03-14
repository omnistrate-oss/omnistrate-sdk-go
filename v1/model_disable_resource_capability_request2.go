/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the DisableResourceCapabilityRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableResourceCapabilityRequest2{}

// DisableResourceCapabilityRequest2 struct for DisableResourceCapabilityRequest2
type DisableResourceCapabilityRequest2 struct {
	// The capability to disable
	Capability string `json:"capability"`
	AdditionalProperties map[string]interface{}
}

type _DisableResourceCapabilityRequest2 DisableResourceCapabilityRequest2

// NewDisableResourceCapabilityRequest2 instantiates a new DisableResourceCapabilityRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableResourceCapabilityRequest2(capability string) *DisableResourceCapabilityRequest2 {
	this := DisableResourceCapabilityRequest2{}
	this.Capability = capability
	return &this
}

// NewDisableResourceCapabilityRequest2WithDefaults instantiates a new DisableResourceCapabilityRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableResourceCapabilityRequest2WithDefaults() *DisableResourceCapabilityRequest2 {
	this := DisableResourceCapabilityRequest2{}
	return &this
}

// GetCapability returns the Capability field value
func (o *DisableResourceCapabilityRequest2) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *DisableResourceCapabilityRequest2) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *DisableResourceCapabilityRequest2) SetCapability(v string) {
	o.Capability = v
}

func (o DisableResourceCapabilityRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableResourceCapabilityRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capability"] = o.Capability

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DisableResourceCapabilityRequest2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capability",
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

	varDisableResourceCapabilityRequest2 := _DisableResourceCapabilityRequest2{}

	err = json.Unmarshal(data, &varDisableResourceCapabilityRequest2)

	if err != nil {
		return err
	}

	*o = DisableResourceCapabilityRequest2(varDisableResourceCapabilityRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capability")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDisableResourceCapabilityRequest2 struct {
	value *DisableResourceCapabilityRequest2
	isSet bool
}

func (v NullableDisableResourceCapabilityRequest2) Get() *DisableResourceCapabilityRequest2 {
	return v.value
}

func (v *NullableDisableResourceCapabilityRequest2) Set(val *DisableResourceCapabilityRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableResourceCapabilityRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableResourceCapabilityRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableResourceCapabilityRequest2(val *DisableResourceCapabilityRequest2) *NullableDisableResourceCapabilityRequest2 {
	return &NullableDisableResourceCapabilityRequest2{value: val, isSet: true}
}

func (v NullableDisableResourceCapabilityRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableResourceCapabilityRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


