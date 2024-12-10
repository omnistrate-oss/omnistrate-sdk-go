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

// checks if the ValidateTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateTokenRequest{}

// ValidateTokenRequest struct for ValidateTokenRequest
type ValidateTokenRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ValidateTokenRequest ValidateTokenRequest

// NewValidateTokenRequest instantiates a new ValidateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateTokenRequest(email string, token string) *ValidateTokenRequest {
	this := ValidateTokenRequest{}
	this.Email = email
	this.Token = token
	return &this
}

// NewValidateTokenRequestWithDefaults instantiates a new ValidateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateTokenRequestWithDefaults() *ValidateTokenRequest {
	this := ValidateTokenRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ValidateTokenRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ValidateTokenRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ValidateTokenRequest) SetEmail(v string) {
	o.Email = v
}

// GetToken returns the Token field value
func (o *ValidateTokenRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ValidateTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ValidateTokenRequest) SetToken(v string) {
	o.Token = v
}

func (o ValidateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateTokenRequest) UnmarshalJSON(data []byte) (err error) {
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

	varValidateTokenRequest := _ValidateTokenRequest{}

	err = json.Unmarshal(data, &varValidateTokenRequest)

	if err != nil {
		return err
	}

	*o = ValidateTokenRequest(varValidateTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateTokenRequest struct {
	value *ValidateTokenRequest
	isSet bool
}

func (v NullableValidateTokenRequest) Get() *ValidateTokenRequest {
	return v.value
}

func (v *NullableValidateTokenRequest) Set(val *ValidateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateTokenRequest(val *ValidateTokenRequest) *NullableValidateTokenRequest {
	return &NullableValidateTokenRequest{value: val, isSet: true}
}

func (v NullableValidateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


