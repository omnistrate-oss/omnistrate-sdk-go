/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InventoryDescribeServiceOfferingResourceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDescribeServiceOfferingResourceResult{}

// InventoryDescribeServiceOfferingResourceResult struct for InventoryDescribeServiceOfferingResourceResult
type InventoryDescribeServiceOfferingResourceResult struct {
	ConsumptionDescribeServiceOfferingResourceResult *DescribeServiceOfferingResourceResult `json:"ConsumptionDescribeServiceOfferingResourceResult,omitempty"`
}

// NewInventoryDescribeServiceOfferingResourceResult instantiates a new InventoryDescribeServiceOfferingResourceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDescribeServiceOfferingResourceResult() *InventoryDescribeServiceOfferingResourceResult {
	this := InventoryDescribeServiceOfferingResourceResult{}
	return &this
}

// NewInventoryDescribeServiceOfferingResourceResultWithDefaults instantiates a new InventoryDescribeServiceOfferingResourceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDescribeServiceOfferingResourceResultWithDefaults() *InventoryDescribeServiceOfferingResourceResult {
	this := InventoryDescribeServiceOfferingResourceResult{}
	return &this
}

// GetConsumptionDescribeServiceOfferingResourceResult returns the ConsumptionDescribeServiceOfferingResourceResult field value if set, zero value otherwise.
func (o *InventoryDescribeServiceOfferingResourceResult) GetConsumptionDescribeServiceOfferingResourceResult() DescribeServiceOfferingResourceResult {
	if o == nil || IsNil(o.ConsumptionDescribeServiceOfferingResourceResult) {
		var ret DescribeServiceOfferingResourceResult
		return ret
	}
	return *o.ConsumptionDescribeServiceOfferingResourceResult
}

// GetConsumptionDescribeServiceOfferingResourceResultOk returns a tuple with the ConsumptionDescribeServiceOfferingResourceResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDescribeServiceOfferingResourceResult) GetConsumptionDescribeServiceOfferingResourceResultOk() (*DescribeServiceOfferingResourceResult, bool) {
	if o == nil || IsNil(o.ConsumptionDescribeServiceOfferingResourceResult) {
		return nil, false
	}
	return o.ConsumptionDescribeServiceOfferingResourceResult, true
}

// HasConsumptionDescribeServiceOfferingResourceResult returns a boolean if a field has been set.
func (o *InventoryDescribeServiceOfferingResourceResult) HasConsumptionDescribeServiceOfferingResourceResult() bool {
	if o != nil && !IsNil(o.ConsumptionDescribeServiceOfferingResourceResult) {
		return true
	}

	return false
}

// SetConsumptionDescribeServiceOfferingResourceResult gets a reference to the given DescribeServiceOfferingResourceResult and assigns it to the ConsumptionDescribeServiceOfferingResourceResult field.
func (o *InventoryDescribeServiceOfferingResourceResult) SetConsumptionDescribeServiceOfferingResourceResult(v DescribeServiceOfferingResourceResult) {
	o.ConsumptionDescribeServiceOfferingResourceResult = &v
}

func (o InventoryDescribeServiceOfferingResourceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryDescribeServiceOfferingResourceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionDescribeServiceOfferingResourceResult) {
		toSerialize["ConsumptionDescribeServiceOfferingResourceResult"] = o.ConsumptionDescribeServiceOfferingResourceResult
	}
	return toSerialize, nil
}

type NullableInventoryDescribeServiceOfferingResourceResult struct {
	value *InventoryDescribeServiceOfferingResourceResult
	isSet bool
}

func (v NullableInventoryDescribeServiceOfferingResourceResult) Get() *InventoryDescribeServiceOfferingResourceResult {
	return v.value
}

func (v *NullableInventoryDescribeServiceOfferingResourceResult) Set(val *InventoryDescribeServiceOfferingResourceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDescribeServiceOfferingResourceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDescribeServiceOfferingResourceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDescribeServiceOfferingResourceResult(val *InventoryDescribeServiceOfferingResourceResult) *NullableInventoryDescribeServiceOfferingResourceResult {
	return &NullableInventoryDescribeServiceOfferingResourceResult{value: val, isSet: true}
}

func (v NullableInventoryDescribeServiceOfferingResourceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDescribeServiceOfferingResourceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


