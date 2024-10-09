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

// checks if the InventoryListServiceOfferingsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryListServiceOfferingsResult{}

// InventoryListServiceOfferingsResult struct for InventoryListServiceOfferingsResult
type InventoryListServiceOfferingsResult struct {
	ConsumptionListServiceOfferingsResult *ListServiceOfferingsResult `json:"ConsumptionListServiceOfferingsResult,omitempty"`
}

// NewInventoryListServiceOfferingsResult instantiates a new InventoryListServiceOfferingsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListServiceOfferingsResult() *InventoryListServiceOfferingsResult {
	this := InventoryListServiceOfferingsResult{}
	return &this
}

// NewInventoryListServiceOfferingsResultWithDefaults instantiates a new InventoryListServiceOfferingsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListServiceOfferingsResultWithDefaults() *InventoryListServiceOfferingsResult {
	this := InventoryListServiceOfferingsResult{}
	return &this
}

// GetConsumptionListServiceOfferingsResult returns the ConsumptionListServiceOfferingsResult field value if set, zero value otherwise.
func (o *InventoryListServiceOfferingsResult) GetConsumptionListServiceOfferingsResult() ListServiceOfferingsResult {
	if o == nil || IsNil(o.ConsumptionListServiceOfferingsResult) {
		var ret ListServiceOfferingsResult
		return ret
	}
	return *o.ConsumptionListServiceOfferingsResult
}

// GetConsumptionListServiceOfferingsResultOk returns a tuple with the ConsumptionListServiceOfferingsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListServiceOfferingsResult) GetConsumptionListServiceOfferingsResultOk() (*ListServiceOfferingsResult, bool) {
	if o == nil || IsNil(o.ConsumptionListServiceOfferingsResult) {
		return nil, false
	}
	return o.ConsumptionListServiceOfferingsResult, true
}

// HasConsumptionListServiceOfferingsResult returns a boolean if a field has been set.
func (o *InventoryListServiceOfferingsResult) HasConsumptionListServiceOfferingsResult() bool {
	if o != nil && !IsNil(o.ConsumptionListServiceOfferingsResult) {
		return true
	}

	return false
}

// SetConsumptionListServiceOfferingsResult gets a reference to the given ListServiceOfferingsResult and assigns it to the ConsumptionListServiceOfferingsResult field.
func (o *InventoryListServiceOfferingsResult) SetConsumptionListServiceOfferingsResult(v ListServiceOfferingsResult) {
	o.ConsumptionListServiceOfferingsResult = &v
}

func (o InventoryListServiceOfferingsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryListServiceOfferingsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionListServiceOfferingsResult) {
		toSerialize["ConsumptionListServiceOfferingsResult"] = o.ConsumptionListServiceOfferingsResult
	}
	return toSerialize, nil
}

type NullableInventoryListServiceOfferingsResult struct {
	value *InventoryListServiceOfferingsResult
	isSet bool
}

func (v NullableInventoryListServiceOfferingsResult) Get() *InventoryListServiceOfferingsResult {
	return v.value
}

func (v *NullableInventoryListServiceOfferingsResult) Set(val *InventoryListServiceOfferingsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListServiceOfferingsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListServiceOfferingsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListServiceOfferingsResult(val *InventoryListServiceOfferingsResult) *NullableInventoryListServiceOfferingsResult {
	return &NullableInventoryListServiceOfferingsResult{value: val, isSet: true}
}

func (v NullableInventoryListServiceOfferingsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListServiceOfferingsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

