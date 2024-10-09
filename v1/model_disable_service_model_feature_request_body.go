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

// checks if the DisableServiceModelFeatureRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableServiceModelFeatureRequestBody{}

// DisableServiceModelFeatureRequestBody struct for DisableServiceModelFeatureRequestBody
type DisableServiceModelFeatureRequestBody struct {
	Feature string `json:"feature"`
}

type _DisableServiceModelFeatureRequestBody DisableServiceModelFeatureRequestBody

// NewDisableServiceModelFeatureRequestBody instantiates a new DisableServiceModelFeatureRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableServiceModelFeatureRequestBody(feature string) *DisableServiceModelFeatureRequestBody {
	this := DisableServiceModelFeatureRequestBody{}
	this.Feature = feature
	return &this
}

// NewDisableServiceModelFeatureRequestBodyWithDefaults instantiates a new DisableServiceModelFeatureRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableServiceModelFeatureRequestBodyWithDefaults() *DisableServiceModelFeatureRequestBody {
	this := DisableServiceModelFeatureRequestBody{}
	return &this
}

// GetFeature returns the Feature field value
func (o *DisableServiceModelFeatureRequestBody) GetFeature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
func (o *DisableServiceModelFeatureRequestBody) GetFeatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feature, true
}

// SetFeature sets field value
func (o *DisableServiceModelFeatureRequestBody) SetFeature(v string) {
	o.Feature = v
}

func (o DisableServiceModelFeatureRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableServiceModelFeatureRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feature"] = o.Feature
	return toSerialize, nil
}

func (o *DisableServiceModelFeatureRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"feature",
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

	varDisableServiceModelFeatureRequestBody := _DisableServiceModelFeatureRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisableServiceModelFeatureRequestBody)

	if err != nil {
		return err
	}

	*o = DisableServiceModelFeatureRequestBody(varDisableServiceModelFeatureRequestBody)

	return err
}

type NullableDisableServiceModelFeatureRequestBody struct {
	value *DisableServiceModelFeatureRequestBody
	isSet bool
}

func (v NullableDisableServiceModelFeatureRequestBody) Get() *DisableServiceModelFeatureRequestBody {
	return v.value
}

func (v *NullableDisableServiceModelFeatureRequestBody) Set(val *DisableServiceModelFeatureRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableServiceModelFeatureRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableServiceModelFeatureRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableServiceModelFeatureRequestBody(val *DisableServiceModelFeatureRequestBody) *NullableDisableServiceModelFeatureRequestBody {
	return &NullableDisableServiceModelFeatureRequestBody{value: val, isSet: true}
}

func (v NullableDisableServiceModelFeatureRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableServiceModelFeatureRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

