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

// checks if the ChangePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangePasswordRequest{}

// ChangePasswordRequest struct for ChangePasswordRequest
type ChangePasswordRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ChangePasswordRequest ChangePasswordRequest

// NewChangePasswordRequest instantiates a new ChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangePasswordRequest(email string, password string, token string) *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	this.Email = email
	this.Password = password
	this.Token = token
	return &this
}

// NewChangePasswordRequestWithDefaults instantiates a new ChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangePasswordRequestWithDefaults() *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ChangePasswordRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ChangePasswordRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *ChangePasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ChangePasswordRequest) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *ChangePasswordRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ChangePasswordRequest) SetToken(v string) {
	o.Token = v
}

func (o ChangePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangePasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
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

	varChangePasswordRequest := _ChangePasswordRequest{}

	err = json.Unmarshal(data, &varChangePasswordRequest)

	if err != nil {
		return err
	}

	*o = ChangePasswordRequest(varChangePasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "password")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangePasswordRequest struct {
	value *ChangePasswordRequest
	isSet bool
}

func (v NullableChangePasswordRequest) Get() *ChangePasswordRequest {
	return v.value
}

func (v *NullableChangePasswordRequest) Set(val *ChangePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangePasswordRequest(val *ChangePasswordRequest) *NullableChangePasswordRequest {
	return &NullableChangePasswordRequest{value: val, isSet: true}
}

func (v NullableChangePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


