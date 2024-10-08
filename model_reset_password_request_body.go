/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResetPasswordRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordRequestBody{}

// ResetPasswordRequestBody struct for ResetPasswordRequestBody
type ResetPasswordRequestBody struct {
	Email string `json:"email"`
}

type _ResetPasswordRequestBody ResetPasswordRequestBody

// NewResetPasswordRequestBody instantiates a new ResetPasswordRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordRequestBody(email string) *ResetPasswordRequestBody {
	this := ResetPasswordRequestBody{}
	this.Email = email
	return &this
}

// NewResetPasswordRequestBodyWithDefaults instantiates a new ResetPasswordRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordRequestBodyWithDefaults() *ResetPasswordRequestBody {
	this := ResetPasswordRequestBody{}
	return &this
}

// GetEmail returns the Email field value
func (o *ResetPasswordRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ResetPasswordRequestBody) SetEmail(v string) {
	o.Email = v
}

func (o ResetPasswordRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *ResetPasswordRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResetPasswordRequestBody := _ResetPasswordRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResetPasswordRequestBody)

	if err != nil {
		return err
	}

	*o = ResetPasswordRequestBody(varResetPasswordRequestBody)

	return err
}

type NullableResetPasswordRequestBody struct {
	value *ResetPasswordRequestBody
	isSet bool
}

func (v NullableResetPasswordRequestBody) Get() *ResetPasswordRequestBody {
	return v.value
}

func (v *NullableResetPasswordRequestBody) Set(val *ResetPasswordRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordRequestBody(val *ResetPasswordRequestBody) *NullableResetPasswordRequestBody {
	return &NullableResetPasswordRequestBody{value: val, isSet: true}
}

func (v NullableResetPasswordRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


