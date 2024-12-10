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

// checks if the ReleaseServiceModelResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseServiceModelResult{}

// ReleaseServiceModelResult struct for ReleaseServiceModelResult
type ReleaseServiceModelResult struct {
	// The release status of the service model
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _ReleaseServiceModelResult ReleaseServiceModelResult

// NewReleaseServiceModelResult instantiates a new ReleaseServiceModelResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseServiceModelResult(status string) *ReleaseServiceModelResult {
	this := ReleaseServiceModelResult{}
	this.Status = status
	return &this
}

// NewReleaseServiceModelResultWithDefaults instantiates a new ReleaseServiceModelResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseServiceModelResultWithDefaults() *ReleaseServiceModelResult {
	this := ReleaseServiceModelResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *ReleaseServiceModelResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReleaseServiceModelResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReleaseServiceModelResult) SetStatus(v string) {
	o.Status = v
}

func (o ReleaseServiceModelResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseServiceModelResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReleaseServiceModelResult) UnmarshalJSON(data []byte) (err error) {
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

	varReleaseServiceModelResult := _ReleaseServiceModelResult{}

	err = json.Unmarshal(data, &varReleaseServiceModelResult)

	if err != nil {
		return err
	}

	*o = ReleaseServiceModelResult(varReleaseServiceModelResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReleaseServiceModelResult struct {
	value *ReleaseServiceModelResult
	isSet bool
}

func (v NullableReleaseServiceModelResult) Get() *ReleaseServiceModelResult {
	return v.value
}

func (v *NullableReleaseServiceModelResult) Set(val *ReleaseServiceModelResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseServiceModelResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseServiceModelResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseServiceModelResult(val *ReleaseServiceModelResult) *NullableReleaseServiceModelResult {
	return &NullableReleaseServiceModelResult{value: val, isSet: true}
}

func (v NullableReleaseServiceModelResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseServiceModelResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


