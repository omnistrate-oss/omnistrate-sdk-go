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

// checks if the SubscriptionUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUsers{}

// SubscriptionUsers struct for SubscriptionUsers
type SubscriptionUsers struct {
	// The email of the user
	Email string `json:"email"`
	// The name of the user
	Name string `json:"name"`
	RoleType string `json:"roleType"`
	// The User ID
	UserId string `json:"userId"`
}

type _SubscriptionUsers SubscriptionUsers

// NewSubscriptionUsers instantiates a new SubscriptionUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUsers(email string, name string, roleType string, userId string) *SubscriptionUsers {
	this := SubscriptionUsers{}
	this.Email = email
	this.Name = name
	this.RoleType = roleType
	this.UserId = userId
	return &this
}

// NewSubscriptionUsersWithDefaults instantiates a new SubscriptionUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUsersWithDefaults() *SubscriptionUsers {
	this := SubscriptionUsers{}
	return &this
}

// GetEmail returns the Email field value
func (o *SubscriptionUsers) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsers) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SubscriptionUsers) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *SubscriptionUsers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionUsers) SetName(v string) {
	o.Name = v
}

// GetRoleType returns the RoleType field value
func (o *SubscriptionUsers) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsers) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *SubscriptionUsers) SetRoleType(v string) {
	o.RoleType = v
}

// GetUserId returns the UserId field value
func (o *SubscriptionUsers) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsers) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SubscriptionUsers) SetUserId(v string) {
	o.UserId = v
}

func (o SubscriptionUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["roleType"] = o.RoleType
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SubscriptionUsers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"name",
		"roleType",
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

	varSubscriptionUsers := _SubscriptionUsers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionUsers)

	if err != nil {
		return err
	}

	*o = SubscriptionUsers(varSubscriptionUsers)

	return err
}

type NullableSubscriptionUsers struct {
	value *SubscriptionUsers
	isSet bool
}

func (v NullableSubscriptionUsers) Get() *SubscriptionUsers {
	return v.value
}

func (v *NullableSubscriptionUsers) Set(val *SubscriptionUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUsers(val *SubscriptionUsers) *NullableSubscriptionUsers {
	return &NullableSubscriptionUsers{value: val, isSet: true}
}

func (v NullableSubscriptionUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


