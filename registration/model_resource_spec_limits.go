/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
)

// checks if the ResourceSpecLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSpecLimits{}

// ResourceSpecLimits The maximum amount of CPU and memory that the container can use
type ResourceSpecLimits struct {
	// The maximum amount of CPU that the container can use
	Cpu *string `json:"cpu,omitempty"`
	// The maximum amount of memory that the container can use
	Memory *string `json:"memory,omitempty"`
}

// NewResourceSpecLimits instantiates a new ResourceSpecLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSpecLimits() *ResourceSpecLimits {
	this := ResourceSpecLimits{}
	return &this
}

// NewResourceSpecLimitsWithDefaults instantiates a new ResourceSpecLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSpecLimitsWithDefaults() *ResourceSpecLimits {
	this := ResourceSpecLimits{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ResourceSpecLimits) GetCpu() string {
	if o == nil || IsNil(o.Cpu) {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSpecLimits) GetCpuOk() (*string, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *ResourceSpecLimits) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ResourceSpecLimits) GetMemory() string {
	if o == nil || IsNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSpecLimits) GetMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *ResourceSpecLimits) SetMemory(v string) {
	o.Memory = &v
}

func (o ResourceSpecLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSpecLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableResourceSpecLimits struct {
	value *ResourceSpecLimits
	isSet bool
}

func (v NullableResourceSpecLimits) Get() *ResourceSpecLimits {
	return v.value
}

func (v *NullableResourceSpecLimits) Set(val *ResourceSpecLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSpecLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSpecLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSpecLimits(val *ResourceSpecLimits) *NullableResourceSpecLimits {
	return &NullableResourceSpecLimits{value: val, isSet: true}
}

func (v NullableResourceSpecLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSpecLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


