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

// checks if the ListSaaSPortalsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSaaSPortalsResult{}

// ListSaaSPortalsResult struct for ListSaaSPortalsResult
type ListSaaSPortalsResult struct {
	// The list of saas portals
	SaasPortals []SaaSPortal `json:"saasPortals"`
	AdditionalProperties map[string]interface{}
}

type _ListSaaSPortalsResult ListSaaSPortalsResult

// NewListSaaSPortalsResult instantiates a new ListSaaSPortalsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSaaSPortalsResult(saasPortals []SaaSPortal) *ListSaaSPortalsResult {
	this := ListSaaSPortalsResult{}
	this.SaasPortals = saasPortals
	return &this
}

// NewListSaaSPortalsResultWithDefaults instantiates a new ListSaaSPortalsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSaaSPortalsResultWithDefaults() *ListSaaSPortalsResult {
	this := ListSaaSPortalsResult{}
	return &this
}

// GetSaasPortals returns the SaasPortals field value
func (o *ListSaaSPortalsResult) GetSaasPortals() []SaaSPortal {
	if o == nil {
		var ret []SaaSPortal
		return ret
	}

	return o.SaasPortals
}

// GetSaasPortalsOk returns a tuple with the SaasPortals field value
// and a boolean to check if the value has been set.
func (o *ListSaaSPortalsResult) GetSaasPortalsOk() ([]SaaSPortal, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaasPortals, true
}

// SetSaasPortals sets field value
func (o *ListSaaSPortalsResult) SetSaasPortals(v []SaaSPortal) {
	o.SaasPortals = v
}

func (o ListSaaSPortalsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSaaSPortalsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["saasPortals"] = o.SaasPortals

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListSaaSPortalsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"saasPortals",
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

	varListSaaSPortalsResult := _ListSaaSPortalsResult{}

	err = json.Unmarshal(data, &varListSaaSPortalsResult)

	if err != nil {
		return err
	}

	*o = ListSaaSPortalsResult(varListSaaSPortalsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "saasPortals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListSaaSPortalsResult struct {
	value *ListSaaSPortalsResult
	isSet bool
}

func (v NullableListSaaSPortalsResult) Get() *ListSaaSPortalsResult {
	return v.value
}

func (v *NullableListSaaSPortalsResult) Set(val *ListSaaSPortalsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListSaaSPortalsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListSaaSPortalsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSaaSPortalsResult(val *ListSaaSPortalsResult) *NullableListSaaSPortalsResult {
	return &NullableListSaaSPortalsResult{value: val, isSet: true}
}

func (v NullableListSaaSPortalsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSaaSPortalsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


