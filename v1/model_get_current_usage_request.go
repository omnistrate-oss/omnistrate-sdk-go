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

// checks if the GetCurrentUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrentUsageRequest{}

// GetCurrentUsageRequest struct for GetCurrentUsageRequest
type GetCurrentUsageRequest struct {
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _GetCurrentUsageRequest GetCurrentUsageRequest

// NewGetCurrentUsageRequest instantiates a new GetCurrentUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentUsageRequest(token string) *GetCurrentUsageRequest {
	this := GetCurrentUsageRequest{}
	this.Token = token
	return &this
}

// NewGetCurrentUsageRequestWithDefaults instantiates a new GetCurrentUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentUsageRequestWithDefaults() *GetCurrentUsageRequest {
	this := GetCurrentUsageRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *GetCurrentUsageRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetCurrentUsageRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetCurrentUsageRequest) SetToken(v string) {
	o.Token = v
}

func (o GetCurrentUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCurrentUsageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGetCurrentUsageRequest := _GetCurrentUsageRequest{}

	err = json.Unmarshal(data, &varGetCurrentUsageRequest)

	if err != nil {
		return err
	}

	*o = GetCurrentUsageRequest(varGetCurrentUsageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCurrentUsageRequest struct {
	value *GetCurrentUsageRequest
	isSet bool
}

func (v NullableGetCurrentUsageRequest) Get() *GetCurrentUsageRequest {
	return v.value
}

func (v *NullableGetCurrentUsageRequest) Set(val *GetCurrentUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentUsageRequest(val *GetCurrentUsageRequest) *NullableGetCurrentUsageRequest {
	return &NullableGetCurrentUsageRequest{value: val, isSet: true}
}

func (v NullableGetCurrentUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


