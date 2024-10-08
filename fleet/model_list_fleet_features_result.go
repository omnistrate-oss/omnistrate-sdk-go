/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListFleetFeaturesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFleetFeaturesResult{}

// ListFleetFeaturesResult struct for ListFleetFeaturesResult
type ListFleetFeaturesResult struct {
	// List of features enabled for the service provider.
	Features []FleetFeature `json:"features"`
}

type _ListFleetFeaturesResult ListFleetFeaturesResult

// NewListFleetFeaturesResult instantiates a new ListFleetFeaturesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFleetFeaturesResult(features []FleetFeature) *ListFleetFeaturesResult {
	this := ListFleetFeaturesResult{}
	this.Features = features
	return &this
}

// NewListFleetFeaturesResultWithDefaults instantiates a new ListFleetFeaturesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFleetFeaturesResultWithDefaults() *ListFleetFeaturesResult {
	this := ListFleetFeaturesResult{}
	return &this
}

// GetFeatures returns the Features field value
func (o *ListFleetFeaturesResult) GetFeatures() []FleetFeature {
	if o == nil {
		var ret []FleetFeature
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ListFleetFeaturesResult) GetFeaturesOk() ([]FleetFeature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ListFleetFeaturesResult) SetFeatures(v []FleetFeature) {
	o.Features = v
}

func (o ListFleetFeaturesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFleetFeaturesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *ListFleetFeaturesResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"features",
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

	varListFleetFeaturesResult := _ListFleetFeaturesResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListFleetFeaturesResult)

	if err != nil {
		return err
	}

	*o = ListFleetFeaturesResult(varListFleetFeaturesResult)

	return err
}

type NullableListFleetFeaturesResult struct {
	value *ListFleetFeaturesResult
	isSet bool
}

func (v NullableListFleetFeaturesResult) Get() *ListFleetFeaturesResult {
	return v.value
}

func (v *NullableListFleetFeaturesResult) Set(val *ListFleetFeaturesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListFleetFeaturesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListFleetFeaturesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFleetFeaturesResult(val *ListFleetFeaturesResult) *NullableListFleetFeaturesResult {
	return &NullableListFleetFeaturesResult{value: val, isSet: true}
}

func (v NullableListFleetFeaturesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFleetFeaturesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


