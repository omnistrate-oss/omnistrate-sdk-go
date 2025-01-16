/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the VUnitHealthReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VUnitHealthReport{}

// VUnitHealthReport struct for VUnitHealthReport
type VUnitHealthReport struct {
	// The health of each vunit under this service environment
	Vunits map[string]interface{} `json:"vunits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VUnitHealthReport VUnitHealthReport

// NewVUnitHealthReport instantiates a new VUnitHealthReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVUnitHealthReport() *VUnitHealthReport {
	this := VUnitHealthReport{}
	return &this
}

// NewVUnitHealthReportWithDefaults instantiates a new VUnitHealthReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVUnitHealthReportWithDefaults() *VUnitHealthReport {
	this := VUnitHealthReport{}
	return &this
}

// GetVunits returns the Vunits field value if set, zero value otherwise.
func (o *VUnitHealthReport) GetVunits() map[string]interface{} {
	if o == nil || IsNil(o.Vunits) {
		var ret map[string]interface{}
		return ret
	}
	return o.Vunits
}

// GetVunitsOk returns a tuple with the Vunits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VUnitHealthReport) GetVunitsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Vunits) {
		return map[string]interface{}{}, false
	}
	return o.Vunits, true
}

// SetVunits gets a reference to the given map[string]interface{} and assigns it to the Vunits field.
func (o *VUnitHealthReport) SetVunits(v map[string]interface{}) {
	o.Vunits = v
}

func (o VUnitHealthReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VUnitHealthReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vunits) {
		toSerialize["vunits"] = o.Vunits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VUnitHealthReport) UnmarshalJSON(data []byte) (err error) {
	varVUnitHealthReport := _VUnitHealthReport{}

	err = json.Unmarshal(data, &varVUnitHealthReport)

	if err != nil {
		return err
	}

	*o = VUnitHealthReport(varVUnitHealthReport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vunits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVUnitHealthReport struct {
	value *VUnitHealthReport
	isSet bool
}

func (v NullableVUnitHealthReport) Get() *VUnitHealthReport {
	return v.value
}

func (v *NullableVUnitHealthReport) Set(val *VUnitHealthReport) {
	v.value = val
	v.isSet = true
}

func (v NullableVUnitHealthReport) IsSet() bool {
	return v.isSet
}

func (v *NullableVUnitHealthReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVUnitHealthReport(val *VUnitHealthReport) *NullableVUnitHealthReport {
	return &NullableVUnitHealthReport{value: val, isSet: true}
}

func (v NullableVUnitHealthReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVUnitHealthReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


