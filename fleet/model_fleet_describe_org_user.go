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

// checks if the FleetDescribeOrgUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetDescribeOrgUser{}

// FleetDescribeOrgUser struct for FleetDescribeOrgUser
type FleetDescribeOrgUser struct {
	// JWT token used to perform authorization
	Token string `json:"token"`
	// ID of a User
	UserId string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _FleetDescribeOrgUser FleetDescribeOrgUser

// NewFleetDescribeOrgUser instantiates a new FleetDescribeOrgUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetDescribeOrgUser(token string, userId string) *FleetDescribeOrgUser {
	this := FleetDescribeOrgUser{}
	this.Token = token
	this.UserId = userId
	return &this
}

// NewFleetDescribeOrgUserWithDefaults instantiates a new FleetDescribeOrgUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetDescribeOrgUserWithDefaults() *FleetDescribeOrgUser {
	this := FleetDescribeOrgUser{}
	return &this
}

// GetToken returns the Token field value
func (o *FleetDescribeOrgUser) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeOrgUser) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetDescribeOrgUser) SetToken(v string) {
	o.Token = v
}

// GetUserId returns the UserId field value
func (o *FleetDescribeOrgUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeOrgUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FleetDescribeOrgUser) SetUserId(v string) {
	o.UserId = v
}

func (o FleetDescribeOrgUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetDescribeOrgUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetDescribeOrgUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"userId",
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

	varFleetDescribeOrgUser := _FleetDescribeOrgUser{}

	err = json.Unmarshal(data, &varFleetDescribeOrgUser)

	if err != nil {
		return err
	}

	*o = FleetDescribeOrgUser(varFleetDescribeOrgUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetDescribeOrgUser struct {
	value *FleetDescribeOrgUser
	isSet bool
}

func (v NullableFleetDescribeOrgUser) Get() *FleetDescribeOrgUser {
	return v.value
}

func (v *NullableFleetDescribeOrgUser) Set(val *FleetDescribeOrgUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetDescribeOrgUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetDescribeOrgUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetDescribeOrgUser(val *FleetDescribeOrgUser) *NullableFleetDescribeOrgUser {
	return &NullableFleetDescribeOrgUser{value: val, isSet: true}
}

func (v NullableFleetDescribeOrgUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetDescribeOrgUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


