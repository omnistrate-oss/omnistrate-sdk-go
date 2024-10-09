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

// checks if the ReleaseTierVersionSetRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseTierVersionSetRequestBody{}

// ReleaseTierVersionSetRequestBody struct for ReleaseTierVersionSetRequestBody
type ReleaseTierVersionSetRequestBody struct {
	// Indicates whether this version set is preferred.
	IsPreferred *bool `json:"isPreferred,omitempty"`
	// The name of the product-tier version set.
	Name *string `json:"name,omitempty"`
}

// NewReleaseTierVersionSetRequestBody instantiates a new ReleaseTierVersionSetRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseTierVersionSetRequestBody() *ReleaseTierVersionSetRequestBody {
	this := ReleaseTierVersionSetRequestBody{}
	return &this
}

// NewReleaseTierVersionSetRequestBodyWithDefaults instantiates a new ReleaseTierVersionSetRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseTierVersionSetRequestBodyWithDefaults() *ReleaseTierVersionSetRequestBody {
	this := ReleaseTierVersionSetRequestBody{}
	return &this
}

// GetIsPreferred returns the IsPreferred field value if set, zero value otherwise.
func (o *ReleaseTierVersionSetRequestBody) GetIsPreferred() bool {
	if o == nil || IsNil(o.IsPreferred) {
		var ret bool
		return ret
	}
	return *o.IsPreferred
}

// GetIsPreferredOk returns a tuple with the IsPreferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseTierVersionSetRequestBody) GetIsPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPreferred) {
		return nil, false
	}
	return o.IsPreferred, true
}

// SetIsPreferred gets a reference to the given bool and assigns it to the IsPreferred field.
func (o *ReleaseTierVersionSetRequestBody) SetIsPreferred(v bool) {
	o.IsPreferred = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReleaseTierVersionSetRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseTierVersionSetRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReleaseTierVersionSetRequestBody) SetName(v string) {
	o.Name = &v
}

func (o ReleaseTierVersionSetRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseTierVersionSetRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPreferred) {
		toSerialize["isPreferred"] = o.IsPreferred
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableReleaseTierVersionSetRequestBody struct {
	value *ReleaseTierVersionSetRequestBody
	isSet bool
}

func (v NullableReleaseTierVersionSetRequestBody) Get() *ReleaseTierVersionSetRequestBody {
	return v.value
}

func (v *NullableReleaseTierVersionSetRequestBody) Set(val *ReleaseTierVersionSetRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseTierVersionSetRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseTierVersionSetRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseTierVersionSetRequestBody(val *ReleaseTierVersionSetRequestBody) *NullableReleaseTierVersionSetRequestBody {
	return &NullableReleaseTierVersionSetRequestBody{value: val, isSet: true}
}

func (v NullableReleaseTierVersionSetRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseTierVersionSetRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

