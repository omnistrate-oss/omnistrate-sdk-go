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

// checks if the DeploymentCellConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentCellConfiguration{}

// DeploymentCellConfiguration struct for DeploymentCellConfiguration
type DeploymentCellConfiguration struct {
	// The amenities available in the deployment cell.
	Amenities []Amenity `json:"Amenities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentCellConfiguration DeploymentCellConfiguration

// NewDeploymentCellConfiguration instantiates a new DeploymentCellConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentCellConfiguration() *DeploymentCellConfiguration {
	this := DeploymentCellConfiguration{}
	return &this
}

// NewDeploymentCellConfigurationWithDefaults instantiates a new DeploymentCellConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentCellConfigurationWithDefaults() *DeploymentCellConfiguration {
	this := DeploymentCellConfiguration{}
	return &this
}

// GetAmenities returns the Amenities field value if set, zero value otherwise.
func (o *DeploymentCellConfiguration) GetAmenities() []Amenity {
	if o == nil || IsNil(o.Amenities) {
		var ret []Amenity
		return ret
	}
	return o.Amenities
}

// GetAmenitiesOk returns a tuple with the Amenities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCellConfiguration) GetAmenitiesOk() ([]Amenity, bool) {
	if o == nil || IsNil(o.Amenities) {
		return nil, false
	}
	return o.Amenities, true
}

// SetAmenities gets a reference to the given []Amenity and assigns it to the Amenities field.
func (o *DeploymentCellConfiguration) SetAmenities(v []Amenity) {
	o.Amenities = v
}

func (o DeploymentCellConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentCellConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amenities) {
		toSerialize["Amenities"] = o.Amenities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentCellConfiguration) UnmarshalJSON(data []byte) (err error) {
	varDeploymentCellConfiguration := _DeploymentCellConfiguration{}

	err = json.Unmarshal(data, &varDeploymentCellConfiguration)

	if err != nil {
		return err
	}

	*o = DeploymentCellConfiguration(varDeploymentCellConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Amenities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentCellConfiguration struct {
	value *DeploymentCellConfiguration
	isSet bool
}

func (v NullableDeploymentCellConfiguration) Get() *DeploymentCellConfiguration {
	return v.value
}

func (v *NullableDeploymentCellConfiguration) Set(val *DeploymentCellConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentCellConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentCellConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentCellConfiguration(val *DeploymentCellConfiguration) *NullableDeploymentCellConfiguration {
	return &NullableDeploymentCellConfiguration{value: val, isSet: true}
}

func (v NullableDeploymentCellConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentCellConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


