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

// checks if the EnvironmentHealthReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentHealthReport{}

// EnvironmentHealthReport struct for EnvironmentHealthReport
type EnvironmentHealthReport struct {
	// Health report for each model in the environment
	Models *map[string]VUnitHealthReport `json:"models,omitempty"`
}

// NewEnvironmentHealthReport instantiates a new EnvironmentHealthReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentHealthReport() *EnvironmentHealthReport {
	this := EnvironmentHealthReport{}
	return &this
}

// NewEnvironmentHealthReportWithDefaults instantiates a new EnvironmentHealthReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentHealthReportWithDefaults() *EnvironmentHealthReport {
	this := EnvironmentHealthReport{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *EnvironmentHealthReport) GetModels() map[string]VUnitHealthReport {
	if o == nil || IsNil(o.Models) {
		var ret map[string]VUnitHealthReport
		return ret
	}
	return *o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentHealthReport) GetModelsOk() (*map[string]VUnitHealthReport, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// SetModels gets a reference to the given map[string]VUnitHealthReport and assigns it to the Models field.
func (o *EnvironmentHealthReport) SetModels(v map[string]VUnitHealthReport) {
	o.Models = &v
}

func (o EnvironmentHealthReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentHealthReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	return toSerialize, nil
}

type NullableEnvironmentHealthReport struct {
	value *EnvironmentHealthReport
	isSet bool
}

func (v NullableEnvironmentHealthReport) Get() *EnvironmentHealthReport {
	return v.value
}

func (v *NullableEnvironmentHealthReport) Set(val *EnvironmentHealthReport) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentHealthReport) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentHealthReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentHealthReport(val *EnvironmentHealthReport) *NullableEnvironmentHealthReport {
	return &NullableEnvironmentHealthReport{value: val, isSet: true}
}

func (v NullableEnvironmentHealthReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentHealthReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


