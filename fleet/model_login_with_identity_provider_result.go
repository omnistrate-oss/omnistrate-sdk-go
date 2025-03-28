/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the LoginWithIdentityProviderResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginWithIdentityProviderResult{}

// LoginWithIdentityProviderResult struct for LoginWithIdentityProviderResult
type LoginWithIdentityProviderResult struct {
	// The jwt token
	JwtToken string `json:"jwtToken"`
	AdditionalProperties map[string]interface{}
}

type _LoginWithIdentityProviderResult LoginWithIdentityProviderResult

// NewLoginWithIdentityProviderResult instantiates a new LoginWithIdentityProviderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginWithIdentityProviderResult(jwtToken string) *LoginWithIdentityProviderResult {
	this := LoginWithIdentityProviderResult{}
	this.JwtToken = jwtToken
	return &this
}

// NewLoginWithIdentityProviderResultWithDefaults instantiates a new LoginWithIdentityProviderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginWithIdentityProviderResultWithDefaults() *LoginWithIdentityProviderResult {
	this := LoginWithIdentityProviderResult{}
	return &this
}

// GetJwtToken returns the JwtToken field value
func (o *LoginWithIdentityProviderResult) GetJwtToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwtToken
}

// GetJwtTokenOk returns a tuple with the JwtToken field value
// and a boolean to check if the value has been set.
func (o *LoginWithIdentityProviderResult) GetJwtTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtToken, true
}

// SetJwtToken sets field value
func (o *LoginWithIdentityProviderResult) SetJwtToken(v string) {
	o.JwtToken = v
}

func (o LoginWithIdentityProviderResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginWithIdentityProviderResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jwtToken"] = o.JwtToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginWithIdentityProviderResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jwtToken",
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

	varLoginWithIdentityProviderResult := _LoginWithIdentityProviderResult{}

	err = json.Unmarshal(data, &varLoginWithIdentityProviderResult)

	if err != nil {
		return err
	}

	*o = LoginWithIdentityProviderResult(varLoginWithIdentityProviderResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jwtToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginWithIdentityProviderResult struct {
	value *LoginWithIdentityProviderResult
	isSet bool
}

func (v NullableLoginWithIdentityProviderResult) Get() *LoginWithIdentityProviderResult {
	return v.value
}

func (v *NullableLoginWithIdentityProviderResult) Set(val *LoginWithIdentityProviderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginWithIdentityProviderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginWithIdentityProviderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginWithIdentityProviderResult(val *LoginWithIdentityProviderResult) *NullableLoginWithIdentityProviderResult {
	return &NullableLoginWithIdentityProviderResult{value: val, isSet: true}
}

func (v NullableLoginWithIdentityProviderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginWithIdentityProviderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


