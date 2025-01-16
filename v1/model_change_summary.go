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

// checks if the ChangeSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeSummary{}

// ChangeSummary struct for ChangeSummary
type ChangeSummary struct {
	// List of individual changes
	Changes []Change `json:"changes"`
	// Summary status of the changes
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _ChangeSummary ChangeSummary

// NewChangeSummary instantiates a new ChangeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeSummary(changes []Change, status string) *ChangeSummary {
	this := ChangeSummary{}
	this.Changes = changes
	this.Status = status
	return &this
}

// NewChangeSummaryWithDefaults instantiates a new ChangeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeSummaryWithDefaults() *ChangeSummary {
	this := ChangeSummary{}
	return &this
}

// GetChanges returns the Changes field value
func (o *ChangeSummary) GetChanges() []Change {
	if o == nil {
		var ret []Change
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *ChangeSummary) GetChangesOk() ([]Change, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *ChangeSummary) SetChanges(v []Change) {
	o.Changes = v
}

// GetStatus returns the Status field value
func (o *ChangeSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChangeSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChangeSummary) SetStatus(v string) {
	o.Status = v
}

func (o ChangeSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["changes"] = o.Changes
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"changes",
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

	varChangeSummary := _ChangeSummary{}

	err = json.Unmarshal(data, &varChangeSummary)

	if err != nil {
		return err
	}

	*o = ChangeSummary(varChangeSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "changes")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeSummary struct {
	value *ChangeSummary
	isSet bool
}

func (v NullableChangeSummary) Get() *ChangeSummary {
	return v.value
}

func (v *NullableChangeSummary) Set(val *ChangeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeSummary(val *ChangeSummary) *NullableChangeSummary {
	return &NullableChangeSummary{value: val, isSet: true}
}

func (v NullableChangeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


