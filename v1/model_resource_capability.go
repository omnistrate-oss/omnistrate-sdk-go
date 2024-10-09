/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResourceCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceCapability{}

// ResourceCapability The capability of a resource
type ResourceCapability struct {
	// The type of capability of a resource
	Capability string `json:"capability"`
	// The configuration parameters of a capability of a resource
	Configuration map[string]interface{} `json:"configuration"`
}

type _ResourceCapability ResourceCapability

// NewResourceCapability instantiates a new ResourceCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCapability(capability string, configuration map[string]interface{}) *ResourceCapability {
	this := ResourceCapability{}
	this.Capability = capability
	this.Configuration = configuration
	return &this
}

// NewResourceCapabilityWithDefaults instantiates a new ResourceCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCapabilityWithDefaults() *ResourceCapability {
	this := ResourceCapability{}
	return &this
}

// GetCapability returns the Capability field value
func (o *ResourceCapability) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *ResourceCapability) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *ResourceCapability) SetCapability(v string) {
	o.Capability = v
}

// GetConfiguration returns the Configuration field value
func (o *ResourceCapability) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *ResourceCapability) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value
func (o *ResourceCapability) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o ResourceCapability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capability"] = o.Capability
	toSerialize["configuration"] = o.Configuration
	return toSerialize, nil
}

func (o *ResourceCapability) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capability",
		"configuration",
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

	varResourceCapability := _ResourceCapability{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceCapability)

	if err != nil {
		return err
	}

	*o = ResourceCapability(varResourceCapability)

	return err
}

type NullableResourceCapability struct {
	value *ResourceCapability
	isSet bool
}

func (v NullableResourceCapability) Get() *ResourceCapability {
	return v.value
}

func (v *NullableResourceCapability) Set(val *ResourceCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCapability(val *ResourceCapability) *NullableResourceCapability {
	return &NullableResourceCapability{value: val, isSet: true}
}

func (v NullableResourceCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

