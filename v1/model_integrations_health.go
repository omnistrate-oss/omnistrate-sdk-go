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

// checks if the IntegrationsHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationsHealth{}

// IntegrationsHealth struct for IntegrationsHealth
type IntegrationsHealth struct {
	// The health status of the integration with customer provider observability stack on BYOA
	CustomerObservabilityHealth *string `json:"CustomerObservabilityHealth,omitempty"`
	// The health status of the integration with service provider observability stack
	InternalObservabilityHealth *string `json:"InternalObservabilityHealth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationsHealth IntegrationsHealth

// NewIntegrationsHealth instantiates a new IntegrationsHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationsHealth() *IntegrationsHealth {
	this := IntegrationsHealth{}
	return &this
}

// NewIntegrationsHealthWithDefaults instantiates a new IntegrationsHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationsHealthWithDefaults() *IntegrationsHealth {
	this := IntegrationsHealth{}
	return &this
}

// GetCustomerObservabilityHealth returns the CustomerObservabilityHealth field value if set, zero value otherwise.
func (o *IntegrationsHealth) GetCustomerObservabilityHealth() string {
	if o == nil || IsNil(o.CustomerObservabilityHealth) {
		var ret string
		return ret
	}
	return *o.CustomerObservabilityHealth
}

// GetCustomerObservabilityHealthOk returns a tuple with the CustomerObservabilityHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationsHealth) GetCustomerObservabilityHealthOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerObservabilityHealth) {
		return nil, false
	}
	return o.CustomerObservabilityHealth, true
}

// SetCustomerObservabilityHealth gets a reference to the given string and assigns it to the CustomerObservabilityHealth field.
func (o *IntegrationsHealth) SetCustomerObservabilityHealth(v string) {
	o.CustomerObservabilityHealth = &v
}

// GetInternalObservabilityHealth returns the InternalObservabilityHealth field value if set, zero value otherwise.
func (o *IntegrationsHealth) GetInternalObservabilityHealth() string {
	if o == nil || IsNil(o.InternalObservabilityHealth) {
		var ret string
		return ret
	}
	return *o.InternalObservabilityHealth
}

// GetInternalObservabilityHealthOk returns a tuple with the InternalObservabilityHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationsHealth) GetInternalObservabilityHealthOk() (*string, bool) {
	if o == nil || IsNil(o.InternalObservabilityHealth) {
		return nil, false
	}
	return o.InternalObservabilityHealth, true
}

// SetInternalObservabilityHealth gets a reference to the given string and assigns it to the InternalObservabilityHealth field.
func (o *IntegrationsHealth) SetInternalObservabilityHealth(v string) {
	o.InternalObservabilityHealth = &v
}

func (o IntegrationsHealth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationsHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerObservabilityHealth) {
		toSerialize["CustomerObservabilityHealth"] = o.CustomerObservabilityHealth
	}
	if !IsNil(o.InternalObservabilityHealth) {
		toSerialize["InternalObservabilityHealth"] = o.InternalObservabilityHealth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationsHealth) UnmarshalJSON(data []byte) (err error) {
	varIntegrationsHealth := _IntegrationsHealth{}

	err = json.Unmarshal(data, &varIntegrationsHealth)

	if err != nil {
		return err
	}

	*o = IntegrationsHealth(varIntegrationsHealth)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "CustomerObservabilityHealth")
		delete(additionalProperties, "InternalObservabilityHealth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationsHealth struct {
	value *IntegrationsHealth
	isSet bool
}

func (v NullableIntegrationsHealth) Get() *IntegrationsHealth {
	return v.value
}

func (v *NullableIntegrationsHealth) Set(val *IntegrationsHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationsHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationsHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationsHealth(val *IntegrationsHealth) *NullableIntegrationsHealth {
	return &NullableIntegrationsHealth{value: val, isSet: true}
}

func (v NullableIntegrationsHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationsHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

