/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the FleetNetworkFeaturesConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetNetworkFeaturesConfiguration{}

// FleetNetworkFeaturesConfiguration struct for FleetNetworkFeaturesConfiguration
type FleetNetworkFeaturesConfiguration struct {
	// Indicates if PrivateLink is enabled for the network
	IsPrivateLinkEnabled *bool `json:"isPrivateLinkEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FleetNetworkFeaturesConfiguration FleetNetworkFeaturesConfiguration

// NewFleetNetworkFeaturesConfiguration instantiates a new FleetNetworkFeaturesConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetNetworkFeaturesConfiguration() *FleetNetworkFeaturesConfiguration {
	this := FleetNetworkFeaturesConfiguration{}
	return &this
}

// NewFleetNetworkFeaturesConfigurationWithDefaults instantiates a new FleetNetworkFeaturesConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetNetworkFeaturesConfigurationWithDefaults() *FleetNetworkFeaturesConfiguration {
	this := FleetNetworkFeaturesConfiguration{}
	return &this
}

// GetIsPrivateLinkEnabled returns the IsPrivateLinkEnabled field value if set, zero value otherwise.
func (o *FleetNetworkFeaturesConfiguration) GetIsPrivateLinkEnabled() bool {
	if o == nil || IsNil(o.IsPrivateLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPrivateLinkEnabled
}

// GetIsPrivateLinkEnabledOk returns a tuple with the IsPrivateLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetNetworkFeaturesConfiguration) GetIsPrivateLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivateLinkEnabled) {
		return nil, false
	}
	return o.IsPrivateLinkEnabled, true
}

// HasIsPrivateLinkEnabled returns a boolean if a field has been set.
func (o *FleetNetworkFeaturesConfiguration) HasIsPrivateLinkEnabled() bool {
	if o != nil && !IsNil(o.IsPrivateLinkEnabled) {
		return true
	}

	return false
}

// SetIsPrivateLinkEnabled gets a reference to the given bool and assigns it to the IsPrivateLinkEnabled field.
func (o *FleetNetworkFeaturesConfiguration) SetIsPrivateLinkEnabled(v bool) {
	o.IsPrivateLinkEnabled = &v
}

func (o FleetNetworkFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetNetworkFeaturesConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPrivateLinkEnabled) {
		toSerialize["isPrivateLinkEnabled"] = o.IsPrivateLinkEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetNetworkFeaturesConfiguration) UnmarshalJSON(data []byte) (err error) {
	varFleetNetworkFeaturesConfiguration := _FleetNetworkFeaturesConfiguration{}

	err = json.Unmarshal(data, &varFleetNetworkFeaturesConfiguration)

	if err != nil {
		return err
	}

	*o = FleetNetworkFeaturesConfiguration(varFleetNetworkFeaturesConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isPrivateLinkEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetNetworkFeaturesConfiguration struct {
	value *FleetNetworkFeaturesConfiguration
	isSet bool
}

func (v NullableFleetNetworkFeaturesConfiguration) Get() *FleetNetworkFeaturesConfiguration {
	return v.value
}

func (v *NullableFleetNetworkFeaturesConfiguration) Set(val *FleetNetworkFeaturesConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetNetworkFeaturesConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetNetworkFeaturesConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetNetworkFeaturesConfiguration(val *FleetNetworkFeaturesConfiguration) *NullableFleetNetworkFeaturesConfiguration {
	return &NullableFleetNetworkFeaturesConfiguration{value: val, isSet: true}
}

func (v NullableFleetNetworkFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetNetworkFeaturesConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


