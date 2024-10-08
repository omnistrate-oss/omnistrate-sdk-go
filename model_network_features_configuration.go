/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
)

// checks if the NetworkFeaturesConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFeaturesConfiguration{}

// NetworkFeaturesConfiguration struct for NetworkFeaturesConfiguration
type NetworkFeaturesConfiguration struct {
	// Indicates if PrivateLink is enabled for the network
	IsPrivateLinkEnabled *bool `json:"isPrivateLinkEnabled,omitempty"`
}

// NewNetworkFeaturesConfiguration instantiates a new NetworkFeaturesConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeaturesConfiguration() *NetworkFeaturesConfiguration {
	this := NetworkFeaturesConfiguration{}
	return &this
}

// NewNetworkFeaturesConfigurationWithDefaults instantiates a new NetworkFeaturesConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeaturesConfigurationWithDefaults() *NetworkFeaturesConfiguration {
	this := NetworkFeaturesConfiguration{}
	return &this
}

// GetIsPrivateLinkEnabled returns the IsPrivateLinkEnabled field value if set, zero value otherwise.
func (o *NetworkFeaturesConfiguration) GetIsPrivateLinkEnabled() bool {
	if o == nil || IsNil(o.IsPrivateLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPrivateLinkEnabled
}

// GetIsPrivateLinkEnabledOk returns a tuple with the IsPrivateLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeaturesConfiguration) GetIsPrivateLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivateLinkEnabled) {
		return nil, false
	}
	return o.IsPrivateLinkEnabled, true
}

// HasIsPrivateLinkEnabled returns a boolean if a field has been set.
func (o *NetworkFeaturesConfiguration) HasIsPrivateLinkEnabled() bool {
	if o != nil && !IsNil(o.IsPrivateLinkEnabled) {
		return true
	}

	return false
}

// SetIsPrivateLinkEnabled gets a reference to the given bool and assigns it to the IsPrivateLinkEnabled field.
func (o *NetworkFeaturesConfiguration) SetIsPrivateLinkEnabled(v bool) {
	o.IsPrivateLinkEnabled = &v
}

func (o NetworkFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFeaturesConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPrivateLinkEnabled) {
		toSerialize["isPrivateLinkEnabled"] = o.IsPrivateLinkEnabled
	}
	return toSerialize, nil
}

type NullableNetworkFeaturesConfiguration struct {
	value *NetworkFeaturesConfiguration
	isSet bool
}

func (v NullableNetworkFeaturesConfiguration) Get() *NetworkFeaturesConfiguration {
	return v.value
}

func (v *NullableNetworkFeaturesConfiguration) Set(val *NetworkFeaturesConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeaturesConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeaturesConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeaturesConfiguration(val *NetworkFeaturesConfiguration) *NullableNetworkFeaturesConfiguration {
	return &NullableNetworkFeaturesConfiguration{value: val, isSet: true}
}

func (v NullableNetworkFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeaturesConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


