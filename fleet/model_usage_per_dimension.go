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

// checks if the UsagePerDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsagePerDimension{}

// UsagePerDimension struct for UsagePerDimension
type UsagePerDimension struct {
	// Dimension of usage
	Dimension *string `json:"dimension,omitempty"`
	// Total amount of usage during the period
	Total *float64 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsagePerDimension UsagePerDimension

// NewUsagePerDimension instantiates a new UsagePerDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsagePerDimension() *UsagePerDimension {
	this := UsagePerDimension{}
	return &this
}

// NewUsagePerDimensionWithDefaults instantiates a new UsagePerDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsagePerDimensionWithDefaults() *UsagePerDimension {
	this := UsagePerDimension{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *UsagePerDimension) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsagePerDimension) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *UsagePerDimension) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *UsagePerDimension) SetDimension(v string) {
	o.Dimension = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UsagePerDimension) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsagePerDimension) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UsagePerDimension) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *UsagePerDimension) SetTotal(v float64) {
	o.Total = &v
}

func (o UsagePerDimension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsagePerDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsagePerDimension) UnmarshalJSON(data []byte) (err error) {
	varUsagePerDimension := _UsagePerDimension{}

	err = json.Unmarshal(data, &varUsagePerDimension)

	if err != nil {
		return err
	}

	*o = UsagePerDimension(varUsagePerDimension)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsagePerDimension struct {
	value *UsagePerDimension
	isSet bool
}

func (v NullableUsagePerDimension) Get() *UsagePerDimension {
	return v.value
}

func (v *NullableUsagePerDimension) Set(val *UsagePerDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableUsagePerDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableUsagePerDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsagePerDimension(val *UsagePerDimension) *NullableUsagePerDimension {
	return &NullableUsagePerDimension{value: val, isSet: true}
}

func (v NullableUsagePerDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsagePerDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


