/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
)

// checks if the DescribeServiceOfferingResourceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceOfferingResourceResult{}

// DescribeServiceOfferingResourceResult struct for DescribeServiceOfferingResourceResult
type DescribeServiceOfferingResourceResult struct {
	// The APIs
	Apis []APIEntity `json:"apis,omitempty"`
}

// NewDescribeServiceOfferingResourceResult instantiates a new DescribeServiceOfferingResourceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceOfferingResourceResult() *DescribeServiceOfferingResourceResult {
	this := DescribeServiceOfferingResourceResult{}
	return &this
}

// NewDescribeServiceOfferingResourceResultWithDefaults instantiates a new DescribeServiceOfferingResourceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceOfferingResourceResultWithDefaults() *DescribeServiceOfferingResourceResult {
	this := DescribeServiceOfferingResourceResult{}
	return &this
}

// GetApis returns the Apis field value if set, zero value otherwise.
func (o *DescribeServiceOfferingResourceResult) GetApis() []APIEntity {
	if o == nil || IsNil(o.Apis) {
		var ret []APIEntity
		return ret
	}
	return o.Apis
}

// GetApisOk returns a tuple with the Apis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResourceResult) GetApisOk() ([]APIEntity, bool) {
	if o == nil || IsNil(o.Apis) {
		return nil, false
	}
	return o.Apis, true
}

// HasApis returns a boolean if a field has been set.
func (o *DescribeServiceOfferingResourceResult) HasApis() bool {
	if o != nil && !IsNil(o.Apis) {
		return true
	}

	return false
}

// SetApis gets a reference to the given []APIEntity and assigns it to the Apis field.
func (o *DescribeServiceOfferingResourceResult) SetApis(v []APIEntity) {
	o.Apis = v
}

func (o DescribeServiceOfferingResourceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceOfferingResourceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apis) {
		toSerialize["apis"] = o.Apis
	}
	return toSerialize, nil
}

type NullableDescribeServiceOfferingResourceResult struct {
	value *DescribeServiceOfferingResourceResult
	isSet bool
}

func (v NullableDescribeServiceOfferingResourceResult) Get() *DescribeServiceOfferingResourceResult {
	return v.value
}

func (v *NullableDescribeServiceOfferingResourceResult) Set(val *DescribeServiceOfferingResourceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceOfferingResourceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceOfferingResourceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceOfferingResourceResult(val *DescribeServiceOfferingResourceResult) *NullableDescribeServiceOfferingResourceResult {
	return &NullableDescribeServiceOfferingResourceResult{value: val, isSet: true}
}

func (v NullableDescribeServiceOfferingResourceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceOfferingResourceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


