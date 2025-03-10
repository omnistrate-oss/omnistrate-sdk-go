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

// checks if the FleetUpdateResourceInstanceDebugModeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetUpdateResourceInstanceDebugModeRequest{}

// FleetUpdateResourceInstanceDebugModeRequest struct for FleetUpdateResourceInstanceDebugModeRequest
type FleetUpdateResourceInstanceDebugModeRequest struct {
	// Enable debug mode
	Enable bool `json:"enable"`
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

type _FleetUpdateResourceInstanceDebugModeRequest FleetUpdateResourceInstanceDebugModeRequest

// NewFleetUpdateResourceInstanceDebugModeRequest instantiates a new FleetUpdateResourceInstanceDebugModeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetUpdateResourceInstanceDebugModeRequest(enable bool, environmentId string, instanceId string, serviceId string, token string) *FleetUpdateResourceInstanceDebugModeRequest {
	this := FleetUpdateResourceInstanceDebugModeRequest{}
	this.Enable = enable
	this.EnvironmentId = environmentId
	this.InstanceId = instanceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewFleetUpdateResourceInstanceDebugModeRequestWithDefaults instantiates a new FleetUpdateResourceInstanceDebugModeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetUpdateResourceInstanceDebugModeRequestWithDefaults() *FleetUpdateResourceInstanceDebugModeRequest {
	this := FleetUpdateResourceInstanceDebugModeRequest{}
	return &this
}

// GetEnable returns the Enable field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) SetEnable(v bool) {
	o.Enable = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceId returns the InstanceId field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateResourceInstanceDebugModeRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetUpdateResourceInstanceDebugModeRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetUpdateResourceInstanceDebugModeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetUpdateResourceInstanceDebugModeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetUpdateResourceInstanceDebugModeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable",
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

	varFleetUpdateResourceInstanceDebugModeRequest := _FleetUpdateResourceInstanceDebugModeRequest{}

	err = json.Unmarshal(data, &varFleetUpdateResourceInstanceDebugModeRequest)

	if err != nil {
		return err
	}

	*o = FleetUpdateResourceInstanceDebugModeRequest(varFleetUpdateResourceInstanceDebugModeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable")
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetUpdateResourceInstanceDebugModeRequest struct {
	value *FleetUpdateResourceInstanceDebugModeRequest
	isSet bool
}

func (v NullableFleetUpdateResourceInstanceDebugModeRequest) Get() *FleetUpdateResourceInstanceDebugModeRequest {
	return v.value
}

func (v *NullableFleetUpdateResourceInstanceDebugModeRequest) Set(val *FleetUpdateResourceInstanceDebugModeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetUpdateResourceInstanceDebugModeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetUpdateResourceInstanceDebugModeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetUpdateResourceInstanceDebugModeRequest(val *FleetUpdateResourceInstanceDebugModeRequest) *NullableFleetUpdateResourceInstanceDebugModeRequest {
	return &NullableFleetUpdateResourceInstanceDebugModeRequest{value: val, isSet: true}
}

func (v NullableFleetUpdateResourceInstanceDebugModeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetUpdateResourceInstanceDebugModeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


