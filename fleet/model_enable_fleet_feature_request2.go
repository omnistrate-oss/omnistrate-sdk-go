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

// checks if the EnableFleetFeatureRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableFleetFeatureRequest2{}

// EnableFleetFeatureRequest2 struct for EnableFleetFeatureRequest2
type EnableFleetFeatureRequest2 struct {
	// The feature to enable.
	Feature *string `json:"feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnableFleetFeatureRequest2 EnableFleetFeatureRequest2

// NewEnableFleetFeatureRequest2 instantiates a new EnableFleetFeatureRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableFleetFeatureRequest2() *EnableFleetFeatureRequest2 {
	this := EnableFleetFeatureRequest2{}
	return &this
}

// NewEnableFleetFeatureRequest2WithDefaults instantiates a new EnableFleetFeatureRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableFleetFeatureRequest2WithDefaults() *EnableFleetFeatureRequest2 {
	this := EnableFleetFeatureRequest2{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *EnableFleetFeatureRequest2) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableFleetFeatureRequest2) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *EnableFleetFeatureRequest2) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *EnableFleetFeatureRequest2) SetFeature(v string) {
	o.Feature = &v
}

func (o EnableFleetFeatureRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableFleetFeatureRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableFleetFeatureRequest2) UnmarshalJSON(data []byte) (err error) {
	varEnableFleetFeatureRequest2 := _EnableFleetFeatureRequest2{}

	err = json.Unmarshal(data, &varEnableFleetFeatureRequest2)

	if err != nil {
		return err
	}

	*o = EnableFleetFeatureRequest2(varEnableFleetFeatureRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableFleetFeatureRequest2 struct {
	value *EnableFleetFeatureRequest2
	isSet bool
}

func (v NullableEnableFleetFeatureRequest2) Get() *EnableFleetFeatureRequest2 {
	return v.value
}

func (v *NullableEnableFleetFeatureRequest2) Set(val *EnableFleetFeatureRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableFleetFeatureRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableFleetFeatureRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableFleetFeatureRequest2(val *EnableFleetFeatureRequest2) *NullableEnableFleetFeatureRequest2 {
	return &NullableEnableFleetFeatureRequest2{value: val, isSet: true}
}

func (v NullableEnableFleetFeatureRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableFleetFeatureRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


