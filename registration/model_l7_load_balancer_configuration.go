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

// checks if the L7LoadBalancerConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L7LoadBalancerConfiguration{}

// L7LoadBalancerConfiguration struct for L7LoadBalancerConfiguration
type L7LoadBalancerConfiguration struct {
	// The paths to configure on the load balancer
	Paths []LoadBalancerPathConfiguration `json:"paths,omitempty"`
}

// NewL7LoadBalancerConfiguration instantiates a new L7LoadBalancerConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL7LoadBalancerConfiguration() *L7LoadBalancerConfiguration {
	this := L7LoadBalancerConfiguration{}
	return &this
}

// NewL7LoadBalancerConfigurationWithDefaults instantiates a new L7LoadBalancerConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL7LoadBalancerConfigurationWithDefaults() *L7LoadBalancerConfiguration {
	this := L7LoadBalancerConfiguration{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *L7LoadBalancerConfiguration) GetPaths() []LoadBalancerPathConfiguration {
	if o == nil || IsNil(o.Paths) {
		var ret []LoadBalancerPathConfiguration
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L7LoadBalancerConfiguration) GetPathsOk() ([]LoadBalancerPathConfiguration, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths gets a reference to the given []LoadBalancerPathConfiguration and assigns it to the Paths field.
func (o *L7LoadBalancerConfiguration) SetPaths(v []LoadBalancerPathConfiguration) {
	o.Paths = v
}

func (o L7LoadBalancerConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L7LoadBalancerConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	return toSerialize, nil
}

type NullableL7LoadBalancerConfiguration struct {
	value *L7LoadBalancerConfiguration
	isSet bool
}

func (v NullableL7LoadBalancerConfiguration) Get() *L7LoadBalancerConfiguration {
	return v.value
}

func (v *NullableL7LoadBalancerConfiguration) Set(val *L7LoadBalancerConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableL7LoadBalancerConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableL7LoadBalancerConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL7LoadBalancerConfiguration(val *L7LoadBalancerConfiguration) *NullableL7LoadBalancerConfiguration {
	return &NullableL7LoadBalancerConfiguration{value: val, isSet: true}
}

func (v NullableL7LoadBalancerConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL7LoadBalancerConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


