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

// checks if the FleetDescribeUserResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetDescribeUserResult{}

// FleetDescribeUserResult struct for FleetDescribeUserResult
type FleetDescribeUserResult struct {
	// The time the user was created.
	CreatedAt string `json:"createdAt"`
	// List of subscriptions associated with the user.
	Subscriptions []UserSubscription `json:"subscriptions"`
	// The user ID
	UserId string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _FleetDescribeUserResult FleetDescribeUserResult

// NewFleetDescribeUserResult instantiates a new FleetDescribeUserResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetDescribeUserResult(createdAt string, subscriptions []UserSubscription, userId string) *FleetDescribeUserResult {
	this := FleetDescribeUserResult{}
	this.CreatedAt = createdAt
	this.Subscriptions = subscriptions
	this.UserId = userId
	return &this
}

// NewFleetDescribeUserResultWithDefaults instantiates a new FleetDescribeUserResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetDescribeUserResultWithDefaults() *FleetDescribeUserResult {
	this := FleetDescribeUserResult{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FleetDescribeUserResult) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeUserResult) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FleetDescribeUserResult) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *FleetDescribeUserResult) GetSubscriptions() []UserSubscription {
	if o == nil {
		var ret []UserSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeUserResult) GetSubscriptionsOk() ([]UserSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *FleetDescribeUserResult) SetSubscriptions(v []UserSubscription) {
	o.Subscriptions = v
}

// GetUserId returns the UserId field value
func (o *FleetDescribeUserResult) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeUserResult) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FleetDescribeUserResult) SetUserId(v string) {
	o.UserId = v
}

func (o FleetDescribeUserResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetDescribeUserResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetDescribeUserResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"subscriptions",
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

	varFleetDescribeUserResult := _FleetDescribeUserResult{}

	err = json.Unmarshal(data, &varFleetDescribeUserResult)

	if err != nil {
		return err
	}

	*o = FleetDescribeUserResult(varFleetDescribeUserResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "subscriptions")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetDescribeUserResult struct {
	value *FleetDescribeUserResult
	isSet bool
}

func (v NullableFleetDescribeUserResult) Get() *FleetDescribeUserResult {
	return v.value
}

func (v *NullableFleetDescribeUserResult) Set(val *FleetDescribeUserResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetDescribeUserResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetDescribeUserResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetDescribeUserResult(val *FleetDescribeUserResult) *NullableFleetDescribeUserResult {
	return &NullableFleetDescribeUserResult{value: val, isSet: true}
}

func (v NullableFleetDescribeUserResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetDescribeUserResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


