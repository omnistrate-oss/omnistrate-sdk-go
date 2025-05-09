/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the SigninRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigninRequest{}

// SigninRequest struct for SigninRequest
type SigninRequest struct {
	Email string `json:"email"`
	HashedPassword *string `json:"hashedPassword,omitempty"`
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SigninRequest SigninRequest

// NewSigninRequest instantiates a new SigninRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigninRequest(email string) *SigninRequest {
	this := SigninRequest{}
	this.Email = email
	return &this
}

// NewSigninRequestWithDefaults instantiates a new SigninRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigninRequestWithDefaults() *SigninRequest {
	this := SigninRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *SigninRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SigninRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SigninRequest) SetEmail(v string) {
	o.Email = v
}

// GetHashedPassword returns the HashedPassword field value if set, zero value otherwise.
func (o *SigninRequest) GetHashedPassword() string {
	if o == nil || IsNil(o.HashedPassword) {
		var ret string
		return ret
	}
	return *o.HashedPassword
}

// GetHashedPasswordOk returns a tuple with the HashedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigninRequest) GetHashedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.HashedPassword) {
		return nil, false
	}
	return o.HashedPassword, true
}

// SetHashedPassword gets a reference to the given string and assigns it to the HashedPassword field.
func (o *SigninRequest) SetHashedPassword(v string) {
	o.HashedPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SigninRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigninRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SigninRequest) SetPassword(v string) {
	o.Password = &v
}

func (o SigninRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigninRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.HashedPassword) {
		toSerialize["hashedPassword"] = o.HashedPassword
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SigninRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSigninRequest := _SigninRequest{}

	err = json.Unmarshal(data, &varSigninRequest)

	if err != nil {
		return err
	}

	*o = SigninRequest(varSigninRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "hashedPassword")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSigninRequest struct {
	value *SigninRequest
	isSet bool
}

func (v NullableSigninRequest) Get() *SigninRequest {
	return v.value
}

func (v *NullableSigninRequest) Set(val *SigninRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSigninRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSigninRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigninRequest(val *SigninRequest) *NullableSigninRequest {
	return &NullableSigninRequest{value: val, isSet: true}
}

func (v NullableSigninRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigninRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


