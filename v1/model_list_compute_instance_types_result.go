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

// checks if the ListComputeInstanceTypesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListComputeInstanceTypesResult{}

// ListComputeInstanceTypesResult struct for ListComputeInstanceTypesResult
type ListComputeInstanceTypesResult struct {
	// The next page token
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The list of compute instance type IDs
	Types []string `json:"types"`
	AdditionalProperties map[string]interface{}
}

type _ListComputeInstanceTypesResult ListComputeInstanceTypesResult

// NewListComputeInstanceTypesResult instantiates a new ListComputeInstanceTypesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListComputeInstanceTypesResult(types []string) *ListComputeInstanceTypesResult {
	this := ListComputeInstanceTypesResult{}
	this.Types = types
	return &this
}

// NewListComputeInstanceTypesResultWithDefaults instantiates a new ListComputeInstanceTypesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListComputeInstanceTypesResultWithDefaults() *ListComputeInstanceTypesResult {
	this := ListComputeInstanceTypesResult{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListComputeInstanceTypesResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListComputeInstanceTypesResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListComputeInstanceTypesResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTypes returns the Types field value
func (o *ListComputeInstanceTypesResult) GetTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *ListComputeInstanceTypesResult) GetTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *ListComputeInstanceTypesResult) SetTypes(v []string) {
	o.Types = v
}

func (o ListComputeInstanceTypesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListComputeInstanceTypesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["types"] = o.Types

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListComputeInstanceTypesResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"types",
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

	varListComputeInstanceTypesResult := _ListComputeInstanceTypesResult{}

	err = json.Unmarshal(data, &varListComputeInstanceTypesResult)

	if err != nil {
		return err
	}

	*o = ListComputeInstanceTypesResult(varListComputeInstanceTypesResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "types")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListComputeInstanceTypesResult struct {
	value *ListComputeInstanceTypesResult
	isSet bool
}

func (v NullableListComputeInstanceTypesResult) Get() *ListComputeInstanceTypesResult {
	return v.value
}

func (v *NullableListComputeInstanceTypesResult) Set(val *ListComputeInstanceTypesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListComputeInstanceTypesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListComputeInstanceTypesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListComputeInstanceTypesResult(val *ListComputeInstanceTypesResult) *NullableListComputeInstanceTypesResult {
	return &NullableListComputeInstanceTypesResult{value: val, isSet: true}
}

func (v NullableListComputeInstanceTypesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListComputeInstanceTypesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


