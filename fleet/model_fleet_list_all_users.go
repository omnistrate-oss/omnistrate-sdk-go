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

// checks if the FleetListAllUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetListAllUsers{}

// FleetListAllUsers struct for FleetListAllUsers
type FleetListAllUsers struct {
	// ID of a Service
	ServiceId *string `json:"serviceId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetListAllUsers FleetListAllUsers

// NewFleetListAllUsers instantiates a new FleetListAllUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetListAllUsers(token string) *FleetListAllUsers {
	this := FleetListAllUsers{}
	this.Token = token
	return &this
}

// NewFleetListAllUsersWithDefaults instantiates a new FleetListAllUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetListAllUsersWithDefaults() *FleetListAllUsers {
	this := FleetListAllUsers{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *FleetListAllUsers) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetListAllUsers) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *FleetListAllUsers) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *FleetListAllUsers) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetToken returns the Token field value
func (o *FleetListAllUsers) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetListAllUsers) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetListAllUsers) SetToken(v string) {
	o.Token = v
}

func (o FleetListAllUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetListAllUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetListAllUsers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varFleetListAllUsers := _FleetListAllUsers{}

	err = json.Unmarshal(data, &varFleetListAllUsers)

	if err != nil {
		return err
	}

	*o = FleetListAllUsers(varFleetListAllUsers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetListAllUsers struct {
	value *FleetListAllUsers
	isSet bool
}

func (v NullableFleetListAllUsers) Get() *FleetListAllUsers {
	return v.value
}

func (v *NullableFleetListAllUsers) Set(val *FleetListAllUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetListAllUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetListAllUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetListAllUsers(val *FleetListAllUsers) *NullableFleetListAllUsers {
	return &NullableFleetListAllUsers{value: val, isSet: true}
}

func (v NullableFleetListAllUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetListAllUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


