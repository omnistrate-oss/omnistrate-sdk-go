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

// checks if the ListSaaSPortalsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSaaSPortalsRequest{}

// ListSaaSPortalsRequest struct for ListSaaSPortalsRequest
type ListSaaSPortalsRequest struct {
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListSaaSPortalsRequest ListSaaSPortalsRequest

// NewListSaaSPortalsRequest instantiates a new ListSaaSPortalsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSaaSPortalsRequest(token string) *ListSaaSPortalsRequest {
	this := ListSaaSPortalsRequest{}
	this.Token = token
	return &this
}

// NewListSaaSPortalsRequestWithDefaults instantiates a new ListSaaSPortalsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSaaSPortalsRequestWithDefaults() *ListSaaSPortalsRequest {
	this := ListSaaSPortalsRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *ListSaaSPortalsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListSaaSPortalsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListSaaSPortalsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListSaaSPortalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSaaSPortalsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListSaaSPortalsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varListSaaSPortalsRequest := _ListSaaSPortalsRequest{}

	err = json.Unmarshal(data, &varListSaaSPortalsRequest)

	if err != nil {
		return err
	}

	*o = ListSaaSPortalsRequest(varListSaaSPortalsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListSaaSPortalsRequest struct {
	value *ListSaaSPortalsRequest
	isSet bool
}

func (v NullableListSaaSPortalsRequest) Get() *ListSaaSPortalsRequest {
	return v.value
}

func (v *NullableListSaaSPortalsRequest) Set(val *ListSaaSPortalsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListSaaSPortalsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListSaaSPortalsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSaaSPortalsRequest(val *ListSaaSPortalsRequest) *NullableListSaaSPortalsRequest {
	return &NullableListSaaSPortalsRequest{value: val, isSet: true}
}

func (v NullableListSaaSPortalsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSaaSPortalsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


