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

// checks if the CompleteOAuthConnectionRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteOAuthConnectionRequest2{}

// CompleteOAuthConnectionRequest2 struct for CompleteOAuthConnectionRequest2
type CompleteOAuthConnectionRequest2 struct {
	// Authorization code from Stripe
	Code *string `json:"code,omitempty"`
	// Random string used on the authorize URL
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompleteOAuthConnectionRequest2 CompleteOAuthConnectionRequest2

// NewCompleteOAuthConnectionRequest2 instantiates a new CompleteOAuthConnectionRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteOAuthConnectionRequest2() *CompleteOAuthConnectionRequest2 {
	this := CompleteOAuthConnectionRequest2{}
	return &this
}

// NewCompleteOAuthConnectionRequest2WithDefaults instantiates a new CompleteOAuthConnectionRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteOAuthConnectionRequest2WithDefaults() *CompleteOAuthConnectionRequest2 {
	this := CompleteOAuthConnectionRequest2{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CompleteOAuthConnectionRequest2) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOAuthConnectionRequest2) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CompleteOAuthConnectionRequest2) SetCode(v string) {
	o.Code = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CompleteOAuthConnectionRequest2) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOAuthConnectionRequest2) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CompleteOAuthConnectionRequest2) SetState(v string) {
	o.State = &v
}

func (o CompleteOAuthConnectionRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteOAuthConnectionRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompleteOAuthConnectionRequest2) UnmarshalJSON(data []byte) (err error) {
	varCompleteOAuthConnectionRequest2 := _CompleteOAuthConnectionRequest2{}

	err = json.Unmarshal(data, &varCompleteOAuthConnectionRequest2)

	if err != nil {
		return err
	}

	*o = CompleteOAuthConnectionRequest2(varCompleteOAuthConnectionRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteOAuthConnectionRequest2 struct {
	value *CompleteOAuthConnectionRequest2
	isSet bool
}

func (v NullableCompleteOAuthConnectionRequest2) Get() *CompleteOAuthConnectionRequest2 {
	return v.value
}

func (v *NullableCompleteOAuthConnectionRequest2) Set(val *CompleteOAuthConnectionRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteOAuthConnectionRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteOAuthConnectionRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteOAuthConnectionRequest2(val *CompleteOAuthConnectionRequest2) *NullableCompleteOAuthConnectionRequest2 {
	return &NullableCompleteOAuthConnectionRequest2{value: val, isSet: true}
}

func (v NullableCompleteOAuthConnectionRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteOAuthConnectionRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


