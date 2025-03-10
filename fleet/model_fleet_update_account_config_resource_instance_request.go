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

// checks if the FleetUpdateAccountConfigResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetUpdateAccountConfigResourceInstanceRequest{}

// FleetUpdateAccountConfigResourceInstanceRequest struct for FleetUpdateAccountConfigResourceInstanceRequest
type FleetUpdateAccountConfigResourceInstanceRequest struct {
	// Disconnect account config instance or not
	Disconnect *bool `json:"disconnect,omitempty"`
	// ID of a Resource Instance
	InstanceId string `json:"instanceId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _FleetUpdateAccountConfigResourceInstanceRequest FleetUpdateAccountConfigResourceInstanceRequest

// NewFleetUpdateAccountConfigResourceInstanceRequest instantiates a new FleetUpdateAccountConfigResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetUpdateAccountConfigResourceInstanceRequest(instanceId string, serviceId string, token string) *FleetUpdateAccountConfigResourceInstanceRequest {
	this := FleetUpdateAccountConfigResourceInstanceRequest{}
	this.InstanceId = instanceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewFleetUpdateAccountConfigResourceInstanceRequestWithDefaults instantiates a new FleetUpdateAccountConfigResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetUpdateAccountConfigResourceInstanceRequestWithDefaults() *FleetUpdateAccountConfigResourceInstanceRequest {
	this := FleetUpdateAccountConfigResourceInstanceRequest{}
	return &this
}

// GetDisconnect returns the Disconnect field value if set, zero value otherwise.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetDisconnect() bool {
	if o == nil || IsNil(o.Disconnect) {
		var ret bool
		return ret
	}
	return *o.Disconnect
}

// GetDisconnectOk returns a tuple with the Disconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetDisconnectOk() (*bool, bool) {
	if o == nil || IsNil(o.Disconnect) {
		return nil, false
	}
	return o.Disconnect, true
}

// HasDisconnect returns a boolean if a field has been set.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) HasDisconnect() bool {
	if o != nil && !IsNil(o.Disconnect) {
		return true
	}

	return false
}

// SetDisconnect gets a reference to the given bool and assigns it to the Disconnect field.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetDisconnect(v bool) {
	o.Disconnect = &v
}

// GetInstanceId returns the InstanceId field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o FleetUpdateAccountConfigResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetUpdateAccountConfigResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disconnect) {
		toSerialize["disconnect"] = o.Disconnect
	}
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetUpdateAccountConfigResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varFleetUpdateAccountConfigResourceInstanceRequest := _FleetUpdateAccountConfigResourceInstanceRequest{}

	err = json.Unmarshal(data, &varFleetUpdateAccountConfigResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = FleetUpdateAccountConfigResourceInstanceRequest(varFleetUpdateAccountConfigResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disconnect")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetUpdateAccountConfigResourceInstanceRequest struct {
	value *FleetUpdateAccountConfigResourceInstanceRequest
	isSet bool
}

func (v NullableFleetUpdateAccountConfigResourceInstanceRequest) Get() *FleetUpdateAccountConfigResourceInstanceRequest {
	return v.value
}

func (v *NullableFleetUpdateAccountConfigResourceInstanceRequest) Set(val *FleetUpdateAccountConfigResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetUpdateAccountConfigResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetUpdateAccountConfigResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetUpdateAccountConfigResourceInstanceRequest(val *FleetUpdateAccountConfigResourceInstanceRequest) *NullableFleetUpdateAccountConfigResourceInstanceRequest {
	return &NullableFleetUpdateAccountConfigResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableFleetUpdateAccountConfigResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetUpdateAccountConfigResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


