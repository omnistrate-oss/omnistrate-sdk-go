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

// checks if the FleetListLinkedInstancesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetListLinkedInstancesRequest{}

// FleetListLinkedInstancesRequest struct for FleetListLinkedInstancesRequest
type FleetListLinkedInstancesRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Resource Instance
	InstanceId string `json:"instanceId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetListLinkedInstancesRequest FleetListLinkedInstancesRequest

// NewFleetListLinkedInstancesRequest instantiates a new FleetListLinkedInstancesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetListLinkedInstancesRequest(environmentId string, instanceId string, serviceId string, token string) *FleetListLinkedInstancesRequest {
	this := FleetListLinkedInstancesRequest{}
	this.EnvironmentId = environmentId
	this.InstanceId = instanceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewFleetListLinkedInstancesRequestWithDefaults instantiates a new FleetListLinkedInstancesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetListLinkedInstancesRequestWithDefaults() *FleetListLinkedInstancesRequest {
	this := FleetListLinkedInstancesRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetListLinkedInstancesRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetListLinkedInstancesRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceId returns the InstanceId field value
func (o *FleetListLinkedInstancesRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FleetListLinkedInstancesRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetListLinkedInstancesRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetListLinkedInstancesRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *FleetListLinkedInstancesRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetListLinkedInstancesRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetListLinkedInstancesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetListLinkedInstancesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetListLinkedInstancesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"instanceId",
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

	varFleetListLinkedInstancesRequest := _FleetListLinkedInstancesRequest{}

	err = json.Unmarshal(data, &varFleetListLinkedInstancesRequest)

	if err != nil {
		return err
	}

	*o = FleetListLinkedInstancesRequest(varFleetListLinkedInstancesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetListLinkedInstancesRequest struct {
	value *FleetListLinkedInstancesRequest
	isSet bool
}

func (v NullableFleetListLinkedInstancesRequest) Get() *FleetListLinkedInstancesRequest {
	return v.value
}

func (v *NullableFleetListLinkedInstancesRequest) Set(val *FleetListLinkedInstancesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetListLinkedInstancesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetListLinkedInstancesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetListLinkedInstancesRequest(val *FleetListLinkedInstancesRequest) *NullableFleetListLinkedInstancesRequest {
	return &NullableFleetListLinkedInstancesRequest{value: val, isSet: true}
}

func (v NullableFleetListLinkedInstancesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetListLinkedInstancesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


