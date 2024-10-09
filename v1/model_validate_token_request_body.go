/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ValidateTokenRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateTokenRequestBody{}

// ValidateTokenRequestBody struct for ValidateTokenRequestBody
type ValidateTokenRequestBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type _ValidateTokenRequestBody ValidateTokenRequestBody

// NewValidateTokenRequestBody instantiates a new ValidateTokenRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateTokenRequestBody(email string, token string) *ValidateTokenRequestBody {
	this := ValidateTokenRequestBody{}
	this.Email = email
	this.Token = token
	return &this
}

// NewValidateTokenRequestBodyWithDefaults instantiates a new ValidateTokenRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateTokenRequestBodyWithDefaults() *ValidateTokenRequestBody {
	this := ValidateTokenRequestBody{}
	return &this
}

// GetEmail returns the Email field value
func (o *ValidateTokenRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ValidateTokenRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ValidateTokenRequestBody) SetEmail(v string) {
	o.Email = v
}

// GetToken returns the Token field value
func (o *ValidateTokenRequestBody) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ValidateTokenRequestBody) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ValidateTokenRequestBody) SetToken(v string) {
	o.Token = v
}

func (o ValidateTokenRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateTokenRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *ValidateTokenRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"token",
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

	varValidateTokenRequestBody := _ValidateTokenRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidateTokenRequestBody)

	if err != nil {
		return err
	}

	*o = ValidateTokenRequestBody(varValidateTokenRequestBody)

	return err
}

type NullableValidateTokenRequestBody struct {
	value *ValidateTokenRequestBody
	isSet bool
}

func (v NullableValidateTokenRequestBody) Get() *ValidateTokenRequestBody {
	return v.value
}

func (v *NullableValidateTokenRequestBody) Set(val *ValidateTokenRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateTokenRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateTokenRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateTokenRequestBody(val *ValidateTokenRequestBody) *NullableValidateTokenRequestBody {
	return &NullableValidateTokenRequestBody{value: val, isSet: true}
}

func (v NullableValidateTokenRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateTokenRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

