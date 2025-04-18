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

// checks if the CreateTierVersionSetResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTierVersionSetResult{}

// CreateTierVersionSetResult struct for CreateTierVersionSetResult
type CreateTierVersionSetResult struct {
	// The version number for the specific version set.
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _CreateTierVersionSetResult CreateTierVersionSetResult

// NewCreateTierVersionSetResult instantiates a new CreateTierVersionSetResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTierVersionSetResult(version string) *CreateTierVersionSetResult {
	this := CreateTierVersionSetResult{}
	this.Version = version
	return &this
}

// NewCreateTierVersionSetResultWithDefaults instantiates a new CreateTierVersionSetResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTierVersionSetResultWithDefaults() *CreateTierVersionSetResult {
	this := CreateTierVersionSetResult{}
	return &this
}

// GetVersion returns the Version field value
func (o *CreateTierVersionSetResult) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateTierVersionSetResult) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CreateTierVersionSetResult) SetVersion(v string) {
	o.Version = v
}

func (o CreateTierVersionSetResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTierVersionSetResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTierVersionSetResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
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

	varCreateTierVersionSetResult := _CreateTierVersionSetResult{}

	err = json.Unmarshal(data, &varCreateTierVersionSetResult)

	if err != nil {
		return err
	}

	*o = CreateTierVersionSetResult(varCreateTierVersionSetResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTierVersionSetResult struct {
	value *CreateTierVersionSetResult
	isSet bool
}

func (v NullableCreateTierVersionSetResult) Get() *CreateTierVersionSetResult {
	return v.value
}

func (v *NullableCreateTierVersionSetResult) Set(val *CreateTierVersionSetResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTierVersionSetResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTierVersionSetResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTierVersionSetResult(val *CreateTierVersionSetResult) *NullableCreateTierVersionSetResult {
	return &NullableCreateTierVersionSetResult{value: val, isSet: true}
}

func (v NullableCreateTierVersionSetResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTierVersionSetResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


