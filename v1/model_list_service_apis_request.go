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

// checks if the ListServiceAPIsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceAPIsRequest{}

// ListServiceAPIsRequest struct for ListServiceAPIsRequest
type ListServiceAPIsRequest struct {
	// ID of a Service Environment
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceAPIsRequest ListServiceAPIsRequest

// NewListServiceAPIsRequest instantiates a new ListServiceAPIsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceAPIsRequest(serviceEnvironmentId string, serviceId string, token string) *ListServiceAPIsRequest {
	this := ListServiceAPIsRequest{}
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewListServiceAPIsRequestWithDefaults instantiates a new ListServiceAPIsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceAPIsRequestWithDefaults() *ListServiceAPIsRequest {
	this := ListServiceAPIsRequest{}
	return &this
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *ListServiceAPIsRequest) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ListServiceAPIsRequest) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *ListServiceAPIsRequest) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *ListServiceAPIsRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListServiceAPIsRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListServiceAPIsRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *ListServiceAPIsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListServiceAPIsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListServiceAPIsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListServiceAPIsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceAPIsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceAPIsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serviceEnvironmentId",
		"serviceId",
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

	varListServiceAPIsRequest := _ListServiceAPIsRequest{}

	err = json.Unmarshal(data, &varListServiceAPIsRequest)

	if err != nil {
		return err
	}

	*o = ListServiceAPIsRequest(varListServiceAPIsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceEnvironmentId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceAPIsRequest struct {
	value *ListServiceAPIsRequest
	isSet bool
}

func (v NullableListServiceAPIsRequest) Get() *ListServiceAPIsRequest {
	return v.value
}

func (v *NullableListServiceAPIsRequest) Set(val *ListServiceAPIsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceAPIsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceAPIsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceAPIsRequest(val *ListServiceAPIsRequest) *NullableListServiceAPIsRequest {
	return &NullableListServiceAPIsRequest{value: val, isSet: true}
}

func (v NullableListServiceAPIsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceAPIsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


