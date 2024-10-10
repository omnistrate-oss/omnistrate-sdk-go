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

// checks if the BackendPortConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackendPortConfiguration{}

// BackendPortConfiguration struct for BackendPortConfiguration
type BackendPortConfiguration struct {
	// The IDs of the resource associated with the backend port
	AssociatedResourceIDs []string `json:"associatedResourceIDs"`
	// The port to forward traffic to
	BackendPort int64 `json:"backendPort"`
	// The ingress port to configure on the load balancer
	IngressPort int64 `json:"ingressPort"`
	AdditionalProperties map[string]interface{}
}

type _BackendPortConfiguration BackendPortConfiguration

// NewBackendPortConfiguration instantiates a new BackendPortConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackendPortConfiguration(associatedResourceIDs []string, backendPort int64, ingressPort int64) *BackendPortConfiguration {
	this := BackendPortConfiguration{}
	this.AssociatedResourceIDs = associatedResourceIDs
	this.BackendPort = backendPort
	this.IngressPort = ingressPort
	return &this
}

// NewBackendPortConfigurationWithDefaults instantiates a new BackendPortConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackendPortConfigurationWithDefaults() *BackendPortConfiguration {
	this := BackendPortConfiguration{}
	return &this
}

// GetAssociatedResourceIDs returns the AssociatedResourceIDs field value
func (o *BackendPortConfiguration) GetAssociatedResourceIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AssociatedResourceIDs
}

// GetAssociatedResourceIDsOk returns a tuple with the AssociatedResourceIDs field value
// and a boolean to check if the value has been set.
func (o *BackendPortConfiguration) GetAssociatedResourceIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedResourceIDs, true
}

// SetAssociatedResourceIDs sets field value
func (o *BackendPortConfiguration) SetAssociatedResourceIDs(v []string) {
	o.AssociatedResourceIDs = v
}

// GetBackendPort returns the BackendPort field value
func (o *BackendPortConfiguration) GetBackendPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BackendPort
}

// GetBackendPortOk returns a tuple with the BackendPort field value
// and a boolean to check if the value has been set.
func (o *BackendPortConfiguration) GetBackendPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendPort, true
}

// SetBackendPort sets field value
func (o *BackendPortConfiguration) SetBackendPort(v int64) {
	o.BackendPort = v
}

// GetIngressPort returns the IngressPort field value
func (o *BackendPortConfiguration) GetIngressPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IngressPort
}

// GetIngressPortOk returns a tuple with the IngressPort field value
// and a boolean to check if the value has been set.
func (o *BackendPortConfiguration) GetIngressPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IngressPort, true
}

// SetIngressPort sets field value
func (o *BackendPortConfiguration) SetIngressPort(v int64) {
	o.IngressPort = v
}

func (o BackendPortConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackendPortConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associatedResourceIDs"] = o.AssociatedResourceIDs
	toSerialize["backendPort"] = o.BackendPort
	toSerialize["ingressPort"] = o.IngressPort

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BackendPortConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associatedResourceIDs",
		"backendPort",
		"ingressPort",
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

	varBackendPortConfiguration := _BackendPortConfiguration{}

	err = json.Unmarshal(data, &varBackendPortConfiguration)

	if err != nil {
		return err
	}

	*o = BackendPortConfiguration(varBackendPortConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associatedResourceIDs")
		delete(additionalProperties, "backendPort")
		delete(additionalProperties, "ingressPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBackendPortConfiguration struct {
	value *BackendPortConfiguration
	isSet bool
}

func (v NullableBackendPortConfiguration) Get() *BackendPortConfiguration {
	return v.value
}

func (v *NullableBackendPortConfiguration) Set(val *BackendPortConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableBackendPortConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableBackendPortConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackendPortConfiguration(val *BackendPortConfiguration) *NullableBackendPortConfiguration {
	return &NullableBackendPortConfiguration{value: val, isSet: true}
}

func (v NullableBackendPortConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackendPortConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


