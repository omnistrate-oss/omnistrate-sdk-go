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

// checks if the RevokeConsumptionUserRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokeConsumptionUserRoleRequest{}

// RevokeConsumptionUserRoleRequest struct for RevokeConsumptionUserRoleRequest
type RevokeConsumptionUserRoleRequest struct {
	Email string `json:"email"`
	// Type of the role
	RoleType string `json:"roleType"`
	// The subscription ID
	SubscriptionId string `json:"subscriptionId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RevokeConsumptionUserRoleRequest RevokeConsumptionUserRoleRequest

// NewRevokeConsumptionUserRoleRequest instantiates a new RevokeConsumptionUserRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeConsumptionUserRoleRequest(email string, roleType string, subscriptionId string, token string) *RevokeConsumptionUserRoleRequest {
	this := RevokeConsumptionUserRoleRequest{}
	this.Email = email
	this.RoleType = roleType
	this.SubscriptionId = subscriptionId
	this.Token = token
	return &this
}

// NewRevokeConsumptionUserRoleRequestWithDefaults instantiates a new RevokeConsumptionUserRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeConsumptionUserRoleRequestWithDefaults() *RevokeConsumptionUserRoleRequest {
	this := RevokeConsumptionUserRoleRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *RevokeConsumptionUserRoleRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RevokeConsumptionUserRoleRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RevokeConsumptionUserRoleRequest) SetEmail(v string) {
	o.Email = v
}

// GetRoleType returns the RoleType field value
func (o *RevokeConsumptionUserRoleRequest) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *RevokeConsumptionUserRoleRequest) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *RevokeConsumptionUserRoleRequest) SetRoleType(v string) {
	o.RoleType = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *RevokeConsumptionUserRoleRequest) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *RevokeConsumptionUserRoleRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *RevokeConsumptionUserRoleRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetToken returns the Token field value
func (o *RevokeConsumptionUserRoleRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RevokeConsumptionUserRoleRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RevokeConsumptionUserRoleRequest) SetToken(v string) {
	o.Token = v
}

func (o RevokeConsumptionUserRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokeConsumptionUserRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["roleType"] = o.RoleType
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokeConsumptionUserRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"roleType",
		"subscriptionId",
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

	varRevokeConsumptionUserRoleRequest := _RevokeConsumptionUserRoleRequest{}

	err = json.Unmarshal(data, &varRevokeConsumptionUserRoleRequest)

	if err != nil {
		return err
	}

	*o = RevokeConsumptionUserRoleRequest(varRevokeConsumptionUserRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "roleType")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokeConsumptionUserRoleRequest struct {
	value *RevokeConsumptionUserRoleRequest
	isSet bool
}

func (v NullableRevokeConsumptionUserRoleRequest) Get() *RevokeConsumptionUserRoleRequest {
	return v.value
}

func (v *NullableRevokeConsumptionUserRoleRequest) Set(val *RevokeConsumptionUserRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeConsumptionUserRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeConsumptionUserRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeConsumptionUserRoleRequest(val *RevokeConsumptionUserRoleRequest) *NullableRevokeConsumptionUserRoleRequest {
	return &NullableRevokeConsumptionUserRoleRequest{value: val, isSet: true}
}

func (v NullableRevokeConsumptionUserRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeConsumptionUserRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


