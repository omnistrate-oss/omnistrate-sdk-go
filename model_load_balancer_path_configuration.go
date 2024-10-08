/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LoadBalancerPathConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPathConfiguration{}

// LoadBalancerPathConfiguration struct for LoadBalancerPathConfiguration
type LoadBalancerPathConfiguration struct {
	// The ID of the resource associated with the path
	AssociatedResourceID string `json:"associatedResourceID"`
	// The REST API path backed by the load balancer
	Path string `json:"path"`
	// The port to forward traffic to
	Port int64 `json:"port"`
}

type _LoadBalancerPathConfiguration LoadBalancerPathConfiguration

// NewLoadBalancerPathConfiguration instantiates a new LoadBalancerPathConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPathConfiguration(associatedResourceID string, path string, port int64) *LoadBalancerPathConfiguration {
	this := LoadBalancerPathConfiguration{}
	this.AssociatedResourceID = associatedResourceID
	this.Path = path
	this.Port = port
	return &this
}

// NewLoadBalancerPathConfigurationWithDefaults instantiates a new LoadBalancerPathConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPathConfigurationWithDefaults() *LoadBalancerPathConfiguration {
	this := LoadBalancerPathConfiguration{}
	return &this
}

// GetAssociatedResourceID returns the AssociatedResourceID field value
func (o *LoadBalancerPathConfiguration) GetAssociatedResourceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociatedResourceID
}

// GetAssociatedResourceIDOk returns a tuple with the AssociatedResourceID field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPathConfiguration) GetAssociatedResourceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedResourceID, true
}

// SetAssociatedResourceID sets field value
func (o *LoadBalancerPathConfiguration) SetAssociatedResourceID(v string) {
	o.AssociatedResourceID = v
}

// GetPath returns the Path field value
func (o *LoadBalancerPathConfiguration) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPathConfiguration) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *LoadBalancerPathConfiguration) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *LoadBalancerPathConfiguration) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPathConfiguration) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LoadBalancerPathConfiguration) SetPort(v int64) {
	o.Port = v
}

func (o LoadBalancerPathConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPathConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associatedResourceID"] = o.AssociatedResourceID
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

func (o *LoadBalancerPathConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associatedResourceID",
		"path",
		"port",
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

	varLoadBalancerPathConfiguration := _LoadBalancerPathConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoadBalancerPathConfiguration)

	if err != nil {
		return err
	}

	*o = LoadBalancerPathConfiguration(varLoadBalancerPathConfiguration)

	return err
}

type NullableLoadBalancerPathConfiguration struct {
	value *LoadBalancerPathConfiguration
	isSet bool
}

func (v NullableLoadBalancerPathConfiguration) Get() *LoadBalancerPathConfiguration {
	return v.value
}

func (v *NullableLoadBalancerPathConfiguration) Set(val *LoadBalancerPathConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPathConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPathConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPathConfiguration(val *LoadBalancerPathConfiguration) *NullableLoadBalancerPathConfiguration {
	return &NullableLoadBalancerPathConfiguration{value: val, isSet: true}
}

func (v NullableLoadBalancerPathConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPathConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


