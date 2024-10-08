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

// checks if the InventoryDescribeServiceOfferingResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDescribeServiceOfferingResult{}

// InventoryDescribeServiceOfferingResult struct for InventoryDescribeServiceOfferingResult
type InventoryDescribeServiceOfferingResult struct {
	ConsumptionDescribeServiceOfferingResult *DescribeServiceOfferingResult `json:"ConsumptionDescribeServiceOfferingResult,omitempty"`
}

// NewInventoryDescribeServiceOfferingResult instantiates a new InventoryDescribeServiceOfferingResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDescribeServiceOfferingResult() *InventoryDescribeServiceOfferingResult {
	this := InventoryDescribeServiceOfferingResult{}
	return &this
}

// NewInventoryDescribeServiceOfferingResultWithDefaults instantiates a new InventoryDescribeServiceOfferingResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDescribeServiceOfferingResultWithDefaults() *InventoryDescribeServiceOfferingResult {
	this := InventoryDescribeServiceOfferingResult{}
	return &this
}

// GetConsumptionDescribeServiceOfferingResult returns the ConsumptionDescribeServiceOfferingResult field value if set, zero value otherwise.
func (o *InventoryDescribeServiceOfferingResult) GetConsumptionDescribeServiceOfferingResult() DescribeServiceOfferingResult {
	if o == nil || IsNil(o.ConsumptionDescribeServiceOfferingResult) {
		var ret DescribeServiceOfferingResult
		return ret
	}
	return *o.ConsumptionDescribeServiceOfferingResult
}

// GetConsumptionDescribeServiceOfferingResultOk returns a tuple with the ConsumptionDescribeServiceOfferingResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDescribeServiceOfferingResult) GetConsumptionDescribeServiceOfferingResultOk() (*DescribeServiceOfferingResult, bool) {
	if o == nil || IsNil(o.ConsumptionDescribeServiceOfferingResult) {
		return nil, false
	}
	return o.ConsumptionDescribeServiceOfferingResult, true
}

// HasConsumptionDescribeServiceOfferingResult returns a boolean if a field has been set.
func (o *InventoryDescribeServiceOfferingResult) HasConsumptionDescribeServiceOfferingResult() bool {
	if o != nil && !IsNil(o.ConsumptionDescribeServiceOfferingResult) {
		return true
	}

	return false
}

// SetConsumptionDescribeServiceOfferingResult gets a reference to the given DescribeServiceOfferingResult and assigns it to the ConsumptionDescribeServiceOfferingResult field.
func (o *InventoryDescribeServiceOfferingResult) SetConsumptionDescribeServiceOfferingResult(v DescribeServiceOfferingResult) {
	o.ConsumptionDescribeServiceOfferingResult = &v
}

func (o InventoryDescribeServiceOfferingResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryDescribeServiceOfferingResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionDescribeServiceOfferingResult) {
		toSerialize["ConsumptionDescribeServiceOfferingResult"] = o.ConsumptionDescribeServiceOfferingResult
	}
	return toSerialize, nil
}

type NullableInventoryDescribeServiceOfferingResult struct {
	value *InventoryDescribeServiceOfferingResult
	isSet bool
}

func (v NullableInventoryDescribeServiceOfferingResult) Get() *InventoryDescribeServiceOfferingResult {
	return v.value
}

func (v *NullableInventoryDescribeServiceOfferingResult) Set(val *InventoryDescribeServiceOfferingResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDescribeServiceOfferingResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDescribeServiceOfferingResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDescribeServiceOfferingResult(val *InventoryDescribeServiceOfferingResult) *NullableInventoryDescribeServiceOfferingResult {
	return &NullableInventoryDescribeServiceOfferingResult{value: val, isSet: true}
}

func (v NullableInventoryDescribeServiceOfferingResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDescribeServiceOfferingResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


