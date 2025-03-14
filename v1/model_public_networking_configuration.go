/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the PublicNetworkingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetworkingConfiguration{}

// PublicNetworkingConfiguration The public networking configuration for the network config
type PublicNetworkingConfiguration struct {
	// Enable a single load balancer for all replicas
	EnableClusterLoadBalancer *bool `json:"enableClusterLoadBalancer,omitempty"`
	// Create an external node load balancer per node rather than exposing the node ip directly
	EnableNodeLoadBalancer *bool `json:"enableNodeLoadBalancer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicNetworkingConfiguration PublicNetworkingConfiguration

// NewPublicNetworkingConfiguration instantiates a new PublicNetworkingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetworkingConfiguration() *PublicNetworkingConfiguration {
	this := PublicNetworkingConfiguration{}
	return &this
}

// NewPublicNetworkingConfigurationWithDefaults instantiates a new PublicNetworkingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkingConfigurationWithDefaults() *PublicNetworkingConfiguration {
	this := PublicNetworkingConfiguration{}
	return &this
}

// GetEnableClusterLoadBalancer returns the EnableClusterLoadBalancer field value if set, zero value otherwise.
func (o *PublicNetworkingConfiguration) GetEnableClusterLoadBalancer() bool {
	if o == nil || IsNil(o.EnableClusterLoadBalancer) {
		var ret bool
		return ret
	}
	return *o.EnableClusterLoadBalancer
}

// GetEnableClusterLoadBalancerOk returns a tuple with the EnableClusterLoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkingConfiguration) GetEnableClusterLoadBalancerOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClusterLoadBalancer) {
		return nil, false
	}
	return o.EnableClusterLoadBalancer, true
}

// SetEnableClusterLoadBalancer gets a reference to the given bool and assigns it to the EnableClusterLoadBalancer field.
func (o *PublicNetworkingConfiguration) SetEnableClusterLoadBalancer(v bool) {
	o.EnableClusterLoadBalancer = &v
}

// GetEnableNodeLoadBalancer returns the EnableNodeLoadBalancer field value if set, zero value otherwise.
func (o *PublicNetworkingConfiguration) GetEnableNodeLoadBalancer() bool {
	if o == nil || IsNil(o.EnableNodeLoadBalancer) {
		var ret bool
		return ret
	}
	return *o.EnableNodeLoadBalancer
}

// GetEnableNodeLoadBalancerOk returns a tuple with the EnableNodeLoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkingConfiguration) GetEnableNodeLoadBalancerOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableNodeLoadBalancer) {
		return nil, false
	}
	return o.EnableNodeLoadBalancer, true
}

// SetEnableNodeLoadBalancer gets a reference to the given bool and assigns it to the EnableNodeLoadBalancer field.
func (o *PublicNetworkingConfiguration) SetEnableNodeLoadBalancer(v bool) {
	o.EnableNodeLoadBalancer = &v
}

func (o PublicNetworkingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNetworkingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableClusterLoadBalancer) {
		toSerialize["enableClusterLoadBalancer"] = o.EnableClusterLoadBalancer
	}
	if !IsNil(o.EnableNodeLoadBalancer) {
		toSerialize["enableNodeLoadBalancer"] = o.EnableNodeLoadBalancer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNetworkingConfiguration) UnmarshalJSON(data []byte) (err error) {
	varPublicNetworkingConfiguration := _PublicNetworkingConfiguration{}

	err = json.Unmarshal(data, &varPublicNetworkingConfiguration)

	if err != nil {
		return err
	}

	*o = PublicNetworkingConfiguration(varPublicNetworkingConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enableClusterLoadBalancer")
		delete(additionalProperties, "enableNodeLoadBalancer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNetworkingConfiguration struct {
	value *PublicNetworkingConfiguration
	isSet bool
}

func (v NullablePublicNetworkingConfiguration) Get() *PublicNetworkingConfiguration {
	return v.value
}

func (v *NullablePublicNetworkingConfiguration) Set(val *PublicNetworkingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetworkingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetworkingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetworkingConfiguration(val *PublicNetworkingConfiguration) *NullablePublicNetworkingConfiguration {
	return &NullablePublicNetworkingConfiguration{value: val, isSet: true}
}

func (v NullablePublicNetworkingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetworkingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


