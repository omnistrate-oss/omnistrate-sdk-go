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

// checks if the FleetListSubscriptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetListSubscriptionsRequest{}

// FleetListSubscriptionsRequest struct for FleetListSubscriptionsRequest
type FleetListSubscriptionsRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetListSubscriptionsRequest FleetListSubscriptionsRequest

// NewFleetListSubscriptionsRequest instantiates a new FleetListSubscriptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetListSubscriptionsRequest(environmentId string, serviceId string, token string) *FleetListSubscriptionsRequest {
	this := FleetListSubscriptionsRequest{}
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewFleetListSubscriptionsRequestWithDefaults instantiates a new FleetListSubscriptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetListSubscriptionsRequestWithDefaults() *FleetListSubscriptionsRequest {
	this := FleetListSubscriptionsRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetListSubscriptionsRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetListSubscriptionsRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetListSubscriptionsRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetListSubscriptionsRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetListSubscriptionsRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetListSubscriptionsRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *FleetListSubscriptionsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetListSubscriptionsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetListSubscriptionsRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetListSubscriptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetListSubscriptionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
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

	varFleetListSubscriptionsRequest := _FleetListSubscriptionsRequest{}

	err = json.Unmarshal(data, &varFleetListSubscriptionsRequest)

	if err != nil {
		return err
	}

	*o = FleetListSubscriptionsRequest(varFleetListSubscriptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetListSubscriptionsRequest struct {
	value *FleetListSubscriptionsRequest
	isSet bool
}

func (v NullableFleetListSubscriptionsRequest) Get() *FleetListSubscriptionsRequest {
	return v.value
}

func (v *NullableFleetListSubscriptionsRequest) Set(val *FleetListSubscriptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetListSubscriptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetListSubscriptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetListSubscriptionsRequest(val *FleetListSubscriptionsRequest) *NullableFleetListSubscriptionsRequest {
	return &NullableFleetListSubscriptionsRequest{value: val, isSet: true}
}

func (v NullableFleetListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetListSubscriptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


