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

// checks if the FleetStopResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetStopResourceInstanceRequest{}

// FleetStopResourceInstanceRequest struct for FleetStopResourceInstanceRequest
type FleetStopResourceInstanceRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Resource Instance
	InstanceId string `json:"instanceId"`
	// ID of a resource
	ResourceId string `json:"resourceId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetStopResourceInstanceRequest FleetStopResourceInstanceRequest

// NewFleetStopResourceInstanceRequest instantiates a new FleetStopResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetStopResourceInstanceRequest(environmentId string, instanceId string, resourceId string, serviceId string, token string) *FleetStopResourceInstanceRequest {
	this := FleetStopResourceInstanceRequest{}
	this.EnvironmentId = environmentId
	this.InstanceId = instanceId
	this.ResourceId = resourceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewFleetStopResourceInstanceRequestWithDefaults instantiates a new FleetStopResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetStopResourceInstanceRequestWithDefaults() *FleetStopResourceInstanceRequest {
	this := FleetStopResourceInstanceRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetStopResourceInstanceRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetStopResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetStopResourceInstanceRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceId returns the InstanceId field value
func (o *FleetStopResourceInstanceRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetStopResourceInstanceRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FleetStopResourceInstanceRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetResourceId returns the ResourceId field value
func (o *FleetStopResourceInstanceRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *FleetStopResourceInstanceRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *FleetStopResourceInstanceRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetStopResourceInstanceRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetStopResourceInstanceRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetStopResourceInstanceRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *FleetStopResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetStopResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetStopResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetStopResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetStopResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetStopResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"instanceId",
		"resourceId",
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

	varFleetStopResourceInstanceRequest := _FleetStopResourceInstanceRequest{}

	err = json.Unmarshal(data, &varFleetStopResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = FleetStopResourceInstanceRequest(varFleetStopResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetStopResourceInstanceRequest struct {
	value *FleetStopResourceInstanceRequest
	isSet bool
}

func (v NullableFleetStopResourceInstanceRequest) Get() *FleetStopResourceInstanceRequest {
	return v.value
}

func (v *NullableFleetStopResourceInstanceRequest) Set(val *FleetStopResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetStopResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetStopResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetStopResourceInstanceRequest(val *FleetStopResourceInstanceRequest) *NullableFleetStopResourceInstanceRequest {
	return &NullableFleetStopResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableFleetStopResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetStopResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


