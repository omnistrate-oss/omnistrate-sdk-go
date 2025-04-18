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

// checks if the DescribeUsersBySubscriptionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeUsersBySubscriptionResult{}

// DescribeUsersBySubscriptionResult struct for DescribeUsersBySubscriptionResult
type DescribeUsersBySubscriptionResult struct {
	// The ID of the subscription
	Id string `json:"Id"`
	// The users associated with the corresponding subscription
	SubscriptionUsers []SubscriptionUsers `json:"subscriptionUsers"`
	AdditionalProperties map[string]interface{}
}

type _DescribeUsersBySubscriptionResult DescribeUsersBySubscriptionResult

// NewDescribeUsersBySubscriptionResult instantiates a new DescribeUsersBySubscriptionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeUsersBySubscriptionResult(id string, subscriptionUsers []SubscriptionUsers) *DescribeUsersBySubscriptionResult {
	this := DescribeUsersBySubscriptionResult{}
	this.Id = id
	this.SubscriptionUsers = subscriptionUsers
	return &this
}

// NewDescribeUsersBySubscriptionResultWithDefaults instantiates a new DescribeUsersBySubscriptionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeUsersBySubscriptionResultWithDefaults() *DescribeUsersBySubscriptionResult {
	this := DescribeUsersBySubscriptionResult{}
	return &this
}

// GetId returns the Id field value
func (o *DescribeUsersBySubscriptionResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeUsersBySubscriptionResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeUsersBySubscriptionResult) SetId(v string) {
	o.Id = v
}

// GetSubscriptionUsers returns the SubscriptionUsers field value
func (o *DescribeUsersBySubscriptionResult) GetSubscriptionUsers() []SubscriptionUsers {
	if o == nil {
		var ret []SubscriptionUsers
		return ret
	}

	return o.SubscriptionUsers
}

// GetSubscriptionUsersOk returns a tuple with the SubscriptionUsers field value
// and a boolean to check if the value has been set.
func (o *DescribeUsersBySubscriptionResult) GetSubscriptionUsersOk() ([]SubscriptionUsers, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionUsers, true
}

// SetSubscriptionUsers sets field value
func (o *DescribeUsersBySubscriptionResult) SetSubscriptionUsers(v []SubscriptionUsers) {
	o.SubscriptionUsers = v
}

func (o DescribeUsersBySubscriptionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeUsersBySubscriptionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["subscriptionUsers"] = o.SubscriptionUsers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeUsersBySubscriptionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Id",
		"subscriptionUsers",
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

	varDescribeUsersBySubscriptionResult := _DescribeUsersBySubscriptionResult{}

	err = json.Unmarshal(data, &varDescribeUsersBySubscriptionResult)

	if err != nil {
		return err
	}

	*o = DescribeUsersBySubscriptionResult(varDescribeUsersBySubscriptionResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "subscriptionUsers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeUsersBySubscriptionResult struct {
	value *DescribeUsersBySubscriptionResult
	isSet bool
}

func (v NullableDescribeUsersBySubscriptionResult) Get() *DescribeUsersBySubscriptionResult {
	return v.value
}

func (v *NullableDescribeUsersBySubscriptionResult) Set(val *DescribeUsersBySubscriptionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeUsersBySubscriptionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeUsersBySubscriptionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeUsersBySubscriptionResult(val *DescribeUsersBySubscriptionResult) *NullableDescribeUsersBySubscriptionResult {
	return &NullableDescribeUsersBySubscriptionResult{value: val, isSet: true}
}

func (v NullableDescribeUsersBySubscriptionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeUsersBySubscriptionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


