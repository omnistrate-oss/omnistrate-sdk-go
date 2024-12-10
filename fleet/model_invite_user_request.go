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

// checks if the InviteUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteUserRequest{}

// InviteUserRequest struct for InviteUserRequest
type InviteUserRequest struct {
	Email string `json:"email"`
	// Type of the role
	RoleType string `json:"roleType"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _InviteUserRequest InviteUserRequest

// NewInviteUserRequest instantiates a new InviteUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteUserRequest(email string, roleType string, token string) *InviteUserRequest {
	this := InviteUserRequest{}
	this.Email = email
	this.RoleType = roleType
	this.Token = token
	return &this
}

// NewInviteUserRequestWithDefaults instantiates a new InviteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteUserRequestWithDefaults() *InviteUserRequest {
	this := InviteUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetRoleType returns the RoleType field value
func (o *InviteUserRequest) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *InviteUserRequest) SetRoleType(v string) {
	o.RoleType = v
}

// GetToken returns the Token field value
func (o *InviteUserRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *InviteUserRequest) SetToken(v string) {
	o.Token = v
}

func (o InviteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["roleType"] = o.RoleType
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InviteUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"roleType",
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

	varInviteUserRequest := _InviteUserRequest{}

	err = json.Unmarshal(data, &varInviteUserRequest)

	if err != nil {
		return err
	}

	*o = InviteUserRequest(varInviteUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "roleType")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInviteUserRequest struct {
	value *InviteUserRequest
	isSet bool
}

func (v NullableInviteUserRequest) Get() *InviteUserRequest {
	return v.value
}

func (v *NullableInviteUserRequest) Set(val *InviteUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteUserRequest(val *InviteUserRequest) *NullableInviteUserRequest {
	return &NullableInviteUserRequest{value: val, isSet: true}
}

func (v NullableInviteUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


