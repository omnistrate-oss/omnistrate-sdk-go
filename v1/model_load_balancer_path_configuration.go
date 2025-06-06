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

// checks if the LoadBalancerPathConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPathConfiguration{}

// LoadBalancerPathConfiguration struct for LoadBalancerPathConfiguration
type LoadBalancerPathConfiguration struct {
	// Override the default target Kubernetes service name with this value
	AssociatedKubernetesServiceName *string `json:"associatedKubernetesServiceName,omitempty"`
	// ID of a resource
	AssociatedResourceID string `json:"associatedResourceID"`
	// The REST API path backed by the load balancer
	Path string `json:"path"`
	// The port to forward traffic to
	Port int64 `json:"port"`
	AdditionalProperties map[string]interface{}
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

// GetAssociatedKubernetesServiceName returns the AssociatedKubernetesServiceName field value if set, zero value otherwise.
func (o *LoadBalancerPathConfiguration) GetAssociatedKubernetesServiceName() string {
	if o == nil || IsNil(o.AssociatedKubernetesServiceName) {
		var ret string
		return ret
	}
	return *o.AssociatedKubernetesServiceName
}

// GetAssociatedKubernetesServiceNameOk returns a tuple with the AssociatedKubernetesServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPathConfiguration) GetAssociatedKubernetesServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedKubernetesServiceName) {
		return nil, false
	}
	return o.AssociatedKubernetesServiceName, true
}

// SetAssociatedKubernetesServiceName gets a reference to the given string and assigns it to the AssociatedKubernetesServiceName field.
func (o *LoadBalancerPathConfiguration) SetAssociatedKubernetesServiceName(v string) {
	o.AssociatedKubernetesServiceName = &v
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
	if !IsNil(o.AssociatedKubernetesServiceName) {
		toSerialize["associatedKubernetesServiceName"] = o.AssociatedKubernetesServiceName
	}
	toSerialize["associatedResourceID"] = o.AssociatedResourceID
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

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

	err = json.Unmarshal(data, &varLoadBalancerPathConfiguration)

	if err != nil {
		return err
	}

	*o = LoadBalancerPathConfiguration(varLoadBalancerPathConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associatedKubernetesServiceName")
		delete(additionalProperties, "associatedResourceID")
		delete(additionalProperties, "path")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

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


