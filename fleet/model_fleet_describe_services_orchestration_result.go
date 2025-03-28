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

// checks if the FleetDescribeServicesOrchestrationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetDescribeServicesOrchestrationResult{}

// FleetDescribeServicesOrchestrationResult struct for FleetDescribeServicesOrchestrationResult
type FleetDescribeServicesOrchestrationResult struct {
	AccessServicesOrchestration *DescribeServicesOrchestrationResult `json:"accessServicesOrchestration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FleetDescribeServicesOrchestrationResult FleetDescribeServicesOrchestrationResult

// NewFleetDescribeServicesOrchestrationResult instantiates a new FleetDescribeServicesOrchestrationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetDescribeServicesOrchestrationResult() *FleetDescribeServicesOrchestrationResult {
	this := FleetDescribeServicesOrchestrationResult{}
	return &this
}

// NewFleetDescribeServicesOrchestrationResultWithDefaults instantiates a new FleetDescribeServicesOrchestrationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetDescribeServicesOrchestrationResultWithDefaults() *FleetDescribeServicesOrchestrationResult {
	this := FleetDescribeServicesOrchestrationResult{}
	return &this
}

// GetAccessServicesOrchestration returns the AccessServicesOrchestration field value if set, zero value otherwise.
func (o *FleetDescribeServicesOrchestrationResult) GetAccessServicesOrchestration() DescribeServicesOrchestrationResult {
	if o == nil || IsNil(o.AccessServicesOrchestration) {
		var ret DescribeServicesOrchestrationResult
		return ret
	}
	return *o.AccessServicesOrchestration
}

// GetAccessServicesOrchestrationOk returns a tuple with the AccessServicesOrchestration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeServicesOrchestrationResult) GetAccessServicesOrchestrationOk() (*DescribeServicesOrchestrationResult, bool) {
	if o == nil || IsNil(o.AccessServicesOrchestration) {
		return nil, false
	}
	return o.AccessServicesOrchestration, true
}

// HasAccessServicesOrchestration returns a boolean if a field has been set.
func (o *FleetDescribeServicesOrchestrationResult) HasAccessServicesOrchestration() bool {
	if o != nil && !IsNil(o.AccessServicesOrchestration) {
		return true
	}

	return false
}

// SetAccessServicesOrchestration gets a reference to the given DescribeServicesOrchestrationResult and assigns it to the AccessServicesOrchestration field.
func (o *FleetDescribeServicesOrchestrationResult) SetAccessServicesOrchestration(v DescribeServicesOrchestrationResult) {
	o.AccessServicesOrchestration = &v
}

func (o FleetDescribeServicesOrchestrationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetDescribeServicesOrchestrationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessServicesOrchestration) {
		toSerialize["accessServicesOrchestration"] = o.AccessServicesOrchestration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetDescribeServicesOrchestrationResult) UnmarshalJSON(data []byte) (err error) {
	varFleetDescribeServicesOrchestrationResult := _FleetDescribeServicesOrchestrationResult{}

	err = json.Unmarshal(data, &varFleetDescribeServicesOrchestrationResult)

	if err != nil {
		return err
	}

	*o = FleetDescribeServicesOrchestrationResult(varFleetDescribeServicesOrchestrationResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessServicesOrchestration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetDescribeServicesOrchestrationResult struct {
	value *FleetDescribeServicesOrchestrationResult
	isSet bool
}

func (v NullableFleetDescribeServicesOrchestrationResult) Get() *FleetDescribeServicesOrchestrationResult {
	return v.value
}

func (v *NullableFleetDescribeServicesOrchestrationResult) Set(val *FleetDescribeServicesOrchestrationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetDescribeServicesOrchestrationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetDescribeServicesOrchestrationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetDescribeServicesOrchestrationResult(val *FleetDescribeServicesOrchestrationResult) *NullableFleetDescribeServicesOrchestrationResult {
	return &NullableFleetDescribeServicesOrchestrationResult{value: val, isSet: true}
}

func (v NullableFleetDescribeServicesOrchestrationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetDescribeServicesOrchestrationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


