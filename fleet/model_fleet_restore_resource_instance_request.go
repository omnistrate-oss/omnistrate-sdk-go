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

// checks if the FleetRestoreResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetRestoreResourceInstanceRequest{}

// FleetRestoreResourceInstanceRequest struct for FleetRestoreResourceInstanceRequest
type FleetRestoreResourceInstanceRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Resource Instance
	InstanceId string `json:"instanceId"`
	// The network type
	NetworkType *string `json:"network_type,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The target restore time
	TargetRestoreTime string `json:"targetRestoreTime"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetRestoreResourceInstanceRequest FleetRestoreResourceInstanceRequest

// NewFleetRestoreResourceInstanceRequest instantiates a new FleetRestoreResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetRestoreResourceInstanceRequest(environmentId string, instanceId string, serviceId string, targetRestoreTime string, token string) *FleetRestoreResourceInstanceRequest {
	this := FleetRestoreResourceInstanceRequest{}
	this.EnvironmentId = environmentId
	this.InstanceId = instanceId
	this.ServiceId = serviceId
	this.TargetRestoreTime = targetRestoreTime
	this.Token = token
	return &this
}

// NewFleetRestoreResourceInstanceRequestWithDefaults instantiates a new FleetRestoreResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetRestoreResourceInstanceRequestWithDefaults() *FleetRestoreResourceInstanceRequest {
	this := FleetRestoreResourceInstanceRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetRestoreResourceInstanceRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetRestoreResourceInstanceRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceId returns the InstanceId field value
func (o *FleetRestoreResourceInstanceRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FleetRestoreResourceInstanceRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *FleetRestoreResourceInstanceRequest) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *FleetRestoreResourceInstanceRequest) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *FleetRestoreResourceInstanceRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetServiceId returns the ServiceId field value
func (o *FleetRestoreResourceInstanceRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetRestoreResourceInstanceRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetTargetRestoreTime returns the TargetRestoreTime field value
func (o *FleetRestoreResourceInstanceRequest) GetTargetRestoreTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRestoreTime
}

// GetTargetRestoreTimeOk returns a tuple with the TargetRestoreTime field value
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetTargetRestoreTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetRestoreTime, true
}

// SetTargetRestoreTime sets field value
func (o *FleetRestoreResourceInstanceRequest) SetTargetRestoreTime(v string) {
	o.TargetRestoreTime = v
}

// GetToken returns the Token field value
func (o *FleetRestoreResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetRestoreResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetRestoreResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetRestoreResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetRestoreResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["instanceId"] = o.InstanceId
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["targetRestoreTime"] = o.TargetRestoreTime
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetRestoreResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"instanceId",
		"serviceId",
		"targetRestoreTime",
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

	varFleetRestoreResourceInstanceRequest := _FleetRestoreResourceInstanceRequest{}

	err = json.Unmarshal(data, &varFleetRestoreResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = FleetRestoreResourceInstanceRequest(varFleetRestoreResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "network_type")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "targetRestoreTime")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetRestoreResourceInstanceRequest struct {
	value *FleetRestoreResourceInstanceRequest
	isSet bool
}

func (v NullableFleetRestoreResourceInstanceRequest) Get() *FleetRestoreResourceInstanceRequest {
	return v.value
}

func (v *NullableFleetRestoreResourceInstanceRequest) Set(val *FleetRestoreResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetRestoreResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetRestoreResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetRestoreResourceInstanceRequest(val *FleetRestoreResourceInstanceRequest) *NullableFleetRestoreResourceInstanceRequest {
	return &NullableFleetRestoreResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableFleetRestoreResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetRestoreResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


