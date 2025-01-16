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

// checks if the DescribeSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeSubscriptionRequest{}

// DescribeSubscriptionRequest struct for DescribeSubscriptionRequest
type DescribeSubscriptionRequest struct {
	// ID of a Subscription
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeSubscriptionRequest DescribeSubscriptionRequest

// NewDescribeSubscriptionRequest instantiates a new DescribeSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSubscriptionRequest(id string, token string) *DescribeSubscriptionRequest {
	this := DescribeSubscriptionRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewDescribeSubscriptionRequestWithDefaults instantiates a new DescribeSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSubscriptionRequestWithDefaults() *DescribeSubscriptionRequest {
	this := DescribeSubscriptionRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DescribeSubscriptionRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeSubscriptionRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *DescribeSubscriptionRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeSubscriptionRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeSubscriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varDescribeSubscriptionRequest := _DescribeSubscriptionRequest{}

	err = json.Unmarshal(data, &varDescribeSubscriptionRequest)

	if err != nil {
		return err
	}

	*o = DescribeSubscriptionRequest(varDescribeSubscriptionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeSubscriptionRequest struct {
	value *DescribeSubscriptionRequest
	isSet bool
}

func (v NullableDescribeSubscriptionRequest) Get() *DescribeSubscriptionRequest {
	return v.value
}

func (v *NullableDescribeSubscriptionRequest) Set(val *DescribeSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSubscriptionRequest(val *DescribeSubscriptionRequest) *NullableDescribeSubscriptionRequest {
	return &NullableDescribeSubscriptionRequest{value: val, isSet: true}
}

func (v NullableDescribeSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


