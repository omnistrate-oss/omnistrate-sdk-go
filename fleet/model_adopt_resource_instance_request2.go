/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the AdoptResourceInstanceRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdoptResourceInstanceRequest2{}

// AdoptResourceInstanceRequest2 struct for AdoptResourceInstanceRequest2
type AdoptResourceInstanceRequest2 struct {
	// The resource adoption configuration
	ResourceAdoptionConfiguration *map[string]ResourceAdoptionConfiguration `json:"resourceAdoptionConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdoptResourceInstanceRequest2 AdoptResourceInstanceRequest2

// NewAdoptResourceInstanceRequest2 instantiates a new AdoptResourceInstanceRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdoptResourceInstanceRequest2() *AdoptResourceInstanceRequest2 {
	this := AdoptResourceInstanceRequest2{}
	return &this
}

// NewAdoptResourceInstanceRequest2WithDefaults instantiates a new AdoptResourceInstanceRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdoptResourceInstanceRequest2WithDefaults() *AdoptResourceInstanceRequest2 {
	this := AdoptResourceInstanceRequest2{}
	return &this
}

// GetResourceAdoptionConfiguration returns the ResourceAdoptionConfiguration field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest2) GetResourceAdoptionConfiguration() map[string]ResourceAdoptionConfiguration {
	if o == nil || IsNil(o.ResourceAdoptionConfiguration) {
		var ret map[string]ResourceAdoptionConfiguration
		return ret
	}
	return *o.ResourceAdoptionConfiguration
}

// GetResourceAdoptionConfigurationOk returns a tuple with the ResourceAdoptionConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest2) GetResourceAdoptionConfigurationOk() (*map[string]ResourceAdoptionConfiguration, bool) {
	if o == nil || IsNil(o.ResourceAdoptionConfiguration) {
		return nil, false
	}
	return o.ResourceAdoptionConfiguration, true
}

// HasResourceAdoptionConfiguration returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest2) HasResourceAdoptionConfiguration() bool {
	if o != nil && !IsNil(o.ResourceAdoptionConfiguration) {
		return true
	}

	return false
}

// SetResourceAdoptionConfiguration gets a reference to the given map[string]ResourceAdoptionConfiguration and assigns it to the ResourceAdoptionConfiguration field.
func (o *AdoptResourceInstanceRequest2) SetResourceAdoptionConfiguration(v map[string]ResourceAdoptionConfiguration) {
	o.ResourceAdoptionConfiguration = &v
}

func (o AdoptResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdoptResourceInstanceRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceAdoptionConfiguration) {
		toSerialize["resourceAdoptionConfiguration"] = o.ResourceAdoptionConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdoptResourceInstanceRequest2) UnmarshalJSON(data []byte) (err error) {
	varAdoptResourceInstanceRequest2 := _AdoptResourceInstanceRequest2{}

	err = json.Unmarshal(data, &varAdoptResourceInstanceRequest2)

	if err != nil {
		return err
	}

	*o = AdoptResourceInstanceRequest2(varAdoptResourceInstanceRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceAdoptionConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdoptResourceInstanceRequest2 struct {
	value *AdoptResourceInstanceRequest2
	isSet bool
}

func (v NullableAdoptResourceInstanceRequest2) Get() *AdoptResourceInstanceRequest2 {
	return v.value
}

func (v *NullableAdoptResourceInstanceRequest2) Set(val *AdoptResourceInstanceRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableAdoptResourceInstanceRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableAdoptResourceInstanceRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdoptResourceInstanceRequest2(val *AdoptResourceInstanceRequest2) *NullableAdoptResourceInstanceRequest2 {
	return &NullableAdoptResourceInstanceRequest2{value: val, isSet: true}
}

func (v NullableAdoptResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdoptResourceInstanceRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


