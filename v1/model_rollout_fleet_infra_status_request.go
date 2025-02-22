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

// checks if the RolloutFleetInfraStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolloutFleetInfraStatusRequest{}

// RolloutFleetInfraStatusRequest struct for RolloutFleetInfraStatusRequest
type RolloutFleetInfraStatusRequest struct {
	// ID of an Infra Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RolloutFleetInfraStatusRequest RolloutFleetInfraStatusRequest

// NewRolloutFleetInfraStatusRequest instantiates a new RolloutFleetInfraStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolloutFleetInfraStatusRequest(id string, serviceId string, token string) *RolloutFleetInfraStatusRequest {
	this := RolloutFleetInfraStatusRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewRolloutFleetInfraStatusRequestWithDefaults instantiates a new RolloutFleetInfraStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolloutFleetInfraStatusRequestWithDefaults() *RolloutFleetInfraStatusRequest {
	this := RolloutFleetInfraStatusRequest{}
	return &this
}

// GetId returns the Id field value
func (o *RolloutFleetInfraStatusRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RolloutFleetInfraStatusRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RolloutFleetInfraStatusRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *RolloutFleetInfraStatusRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *RolloutFleetInfraStatusRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *RolloutFleetInfraStatusRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *RolloutFleetInfraStatusRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RolloutFleetInfraStatusRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RolloutFleetInfraStatusRequest) SetToken(v string) {
	o.Token = v
}

func (o RolloutFleetInfraStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolloutFleetInfraStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RolloutFleetInfraStatusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRolloutFleetInfraStatusRequest := _RolloutFleetInfraStatusRequest{}

	err = json.Unmarshal(data, &varRolloutFleetInfraStatusRequest)

	if err != nil {
		return err
	}

	*o = RolloutFleetInfraStatusRequest(varRolloutFleetInfraStatusRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRolloutFleetInfraStatusRequest struct {
	value *RolloutFleetInfraStatusRequest
	isSet bool
}

func (v NullableRolloutFleetInfraStatusRequest) Get() *RolloutFleetInfraStatusRequest {
	return v.value
}

func (v *NullableRolloutFleetInfraStatusRequest) Set(val *RolloutFleetInfraStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutFleetInfraStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutFleetInfraStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutFleetInfraStatusRequest(val *RolloutFleetInfraStatusRequest) *NullableRolloutFleetInfraStatusRequest {
	return &NullableRolloutFleetInfraStatusRequest{value: val, isSet: true}
}

func (v NullableRolloutFleetInfraStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutFleetInfraStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


