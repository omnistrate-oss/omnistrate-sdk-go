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

// checks if the OmnistrateServiceHealthResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnistrateServiceHealthResult{}

// OmnistrateServiceHealthResult struct for OmnistrateServiceHealthResult
type OmnistrateServiceHealthResult struct {
	// The status of an operation
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _OmnistrateServiceHealthResult OmnistrateServiceHealthResult

// NewOmnistrateServiceHealthResult instantiates a new OmnistrateServiceHealthResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnistrateServiceHealthResult(status string) *OmnistrateServiceHealthResult {
	this := OmnistrateServiceHealthResult{}
	this.Status = status
	return &this
}

// NewOmnistrateServiceHealthResultWithDefaults instantiates a new OmnistrateServiceHealthResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnistrateServiceHealthResultWithDefaults() *OmnistrateServiceHealthResult {
	this := OmnistrateServiceHealthResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *OmnistrateServiceHealthResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OmnistrateServiceHealthResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OmnistrateServiceHealthResult) SetStatus(v string) {
	o.Status = v
}

func (o OmnistrateServiceHealthResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnistrateServiceHealthResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OmnistrateServiceHealthResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varOmnistrateServiceHealthResult := _OmnistrateServiceHealthResult{}

	err = json.Unmarshal(data, &varOmnistrateServiceHealthResult)

	if err != nil {
		return err
	}

	*o = OmnistrateServiceHealthResult(varOmnistrateServiceHealthResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOmnistrateServiceHealthResult struct {
	value *OmnistrateServiceHealthResult
	isSet bool
}

func (v NullableOmnistrateServiceHealthResult) Get() *OmnistrateServiceHealthResult {
	return v.value
}

func (v *NullableOmnistrateServiceHealthResult) Set(val *OmnistrateServiceHealthResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnistrateServiceHealthResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnistrateServiceHealthResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnistrateServiceHealthResult(val *OmnistrateServiceHealthResult) *NullableOmnistrateServiceHealthResult {
	return &NullableOmnistrateServiceHealthResult{value: val, isSet: true}
}

func (v NullableOmnistrateServiceHealthResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnistrateServiceHealthResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


