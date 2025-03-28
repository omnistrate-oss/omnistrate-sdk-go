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

// checks if the ReleaseImageConfigRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseImageConfigRequest2{}

// ReleaseImageConfigRequest2 struct for ReleaseImageConfigRequest2
type ReleaseImageConfigRequest2 struct {
	// The product tier ID
	ProductTierId *string `json:"productTierId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReleaseImageConfigRequest2 ReleaseImageConfigRequest2

// NewReleaseImageConfigRequest2 instantiates a new ReleaseImageConfigRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseImageConfigRequest2() *ReleaseImageConfigRequest2 {
	this := ReleaseImageConfigRequest2{}
	return &this
}

// NewReleaseImageConfigRequest2WithDefaults instantiates a new ReleaseImageConfigRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseImageConfigRequest2WithDefaults() *ReleaseImageConfigRequest2 {
	this := ReleaseImageConfigRequest2{}
	return &this
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *ReleaseImageConfigRequest2) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseImageConfigRequest2) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *ReleaseImageConfigRequest2) SetProductTierId(v string) {
	o.ProductTierId = &v
}

func (o ReleaseImageConfigRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseImageConfigRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductTierId) {
		toSerialize["productTierId"] = o.ProductTierId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReleaseImageConfigRequest2) UnmarshalJSON(data []byte) (err error) {
	varReleaseImageConfigRequest2 := _ReleaseImageConfigRequest2{}

	err = json.Unmarshal(data, &varReleaseImageConfigRequest2)

	if err != nil {
		return err
	}

	*o = ReleaseImageConfigRequest2(varReleaseImageConfigRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productTierId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReleaseImageConfigRequest2 struct {
	value *ReleaseImageConfigRequest2
	isSet bool
}

func (v NullableReleaseImageConfigRequest2) Get() *ReleaseImageConfigRequest2 {
	return v.value
}

func (v *NullableReleaseImageConfigRequest2) Set(val *ReleaseImageConfigRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseImageConfigRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseImageConfigRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseImageConfigRequest2(val *ReleaseImageConfigRequest2) *NullableReleaseImageConfigRequest2 {
	return &NullableReleaseImageConfigRequest2{value: val, isSet: true}
}

func (v NullableReleaseImageConfigRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseImageConfigRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


