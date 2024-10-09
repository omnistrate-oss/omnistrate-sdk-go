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

// checks if the EnableFleetFeatureRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableFleetFeatureRequestBody{}

// EnableFleetFeatureRequestBody struct for EnableFleetFeatureRequestBody
type EnableFleetFeatureRequestBody struct {
	// The feature to enable.
	Feature *string `json:"feature,omitempty"`
}

// NewEnableFleetFeatureRequestBody instantiates a new EnableFleetFeatureRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableFleetFeatureRequestBody() *EnableFleetFeatureRequestBody {
	this := EnableFleetFeatureRequestBody{}
	return &this
}

// NewEnableFleetFeatureRequestBodyWithDefaults instantiates a new EnableFleetFeatureRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableFleetFeatureRequestBodyWithDefaults() *EnableFleetFeatureRequestBody {
	this := EnableFleetFeatureRequestBody{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *EnableFleetFeatureRequestBody) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableFleetFeatureRequestBody) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *EnableFleetFeatureRequestBody) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *EnableFleetFeatureRequestBody) SetFeature(v string) {
	o.Feature = &v
}

func (o EnableFleetFeatureRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableFleetFeatureRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	return toSerialize, nil
}

type NullableEnableFleetFeatureRequestBody struct {
	value *EnableFleetFeatureRequestBody
	isSet bool
}

func (v NullableEnableFleetFeatureRequestBody) Get() *EnableFleetFeatureRequestBody {
	return v.value
}

func (v *NullableEnableFleetFeatureRequestBody) Set(val *EnableFleetFeatureRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableFleetFeatureRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableFleetFeatureRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableFleetFeatureRequestBody(val *EnableFleetFeatureRequestBody) *NullableEnableFleetFeatureRequestBody {
	return &NullableEnableFleetFeatureRequestBody{value: val, isSet: true}
}

func (v NullableEnableFleetFeatureRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableFleetFeatureRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

