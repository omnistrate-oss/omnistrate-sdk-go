/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SigninRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigninRequestBody{}

// SigninRequestBody struct for SigninRequestBody
type SigninRequestBody struct {
	Email string `json:"email"`
	HashedPassword *string `json:"hashedPassword,omitempty"`
	Password *string `json:"password,omitempty"`
}

type _SigninRequestBody SigninRequestBody

// NewSigninRequestBody instantiates a new SigninRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigninRequestBody(email string) *SigninRequestBody {
	this := SigninRequestBody{}
	this.Email = email
	return &this
}

// NewSigninRequestBodyWithDefaults instantiates a new SigninRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigninRequestBodyWithDefaults() *SigninRequestBody {
	this := SigninRequestBody{}
	return &this
}

// GetEmail returns the Email field value
func (o *SigninRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SigninRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SigninRequestBody) SetEmail(v string) {
	o.Email = v
}

// GetHashedPassword returns the HashedPassword field value if set, zero value otherwise.
func (o *SigninRequestBody) GetHashedPassword() string {
	if o == nil || IsNil(o.HashedPassword) {
		var ret string
		return ret
	}
	return *o.HashedPassword
}

// GetHashedPasswordOk returns a tuple with the HashedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigninRequestBody) GetHashedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.HashedPassword) {
		return nil, false
	}
	return o.HashedPassword, true
}

// HasHashedPassword returns a boolean if a field has been set.
func (o *SigninRequestBody) HasHashedPassword() bool {
	if o != nil && !IsNil(o.HashedPassword) {
		return true
	}

	return false
}

// SetHashedPassword gets a reference to the given string and assigns it to the HashedPassword field.
func (o *SigninRequestBody) SetHashedPassword(v string) {
	o.HashedPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SigninRequestBody) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigninRequestBody) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SigninRequestBody) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SigninRequestBody) SetPassword(v string) {
	o.Password = &v
}

func (o SigninRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigninRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.HashedPassword) {
		toSerialize["hashedPassword"] = o.HashedPassword
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

func (o *SigninRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varSigninRequestBody := _SigninRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSigninRequestBody)

	if err != nil {
		return err
	}

	*o = SigninRequestBody(varSigninRequestBody)

	return err
}

type NullableSigninRequestBody struct {
	value *SigninRequestBody
	isSet bool
}

func (v NullableSigninRequestBody) Get() *SigninRequestBody {
	return v.value
}

func (v *NullableSigninRequestBody) Set(val *SigninRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSigninRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSigninRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigninRequestBody(val *SigninRequestBody) *NullableSigninRequestBody {
	return &NullableSigninRequestBody{value: val, isSet: true}
}

func (v NullableSigninRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigninRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


