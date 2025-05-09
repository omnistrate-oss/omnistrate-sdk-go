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

// checks if the ListHelmPackagesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHelmPackagesResult{}

// ListHelmPackagesResult struct for ListHelmPackagesResult
type ListHelmPackagesResult struct {
	// List of Helm packages
	HelmPackages []HelmPackage `json:"helmPackages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListHelmPackagesResult ListHelmPackagesResult

// NewListHelmPackagesResult instantiates a new ListHelmPackagesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHelmPackagesResult() *ListHelmPackagesResult {
	this := ListHelmPackagesResult{}
	return &this
}

// NewListHelmPackagesResultWithDefaults instantiates a new ListHelmPackagesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHelmPackagesResultWithDefaults() *ListHelmPackagesResult {
	this := ListHelmPackagesResult{}
	return &this
}

// GetHelmPackages returns the HelmPackages field value if set, zero value otherwise.
func (o *ListHelmPackagesResult) GetHelmPackages() []HelmPackage {
	if o == nil || IsNil(o.HelmPackages) {
		var ret []HelmPackage
		return ret
	}
	return o.HelmPackages
}

// GetHelmPackagesOk returns a tuple with the HelmPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHelmPackagesResult) GetHelmPackagesOk() ([]HelmPackage, bool) {
	if o == nil || IsNil(o.HelmPackages) {
		return nil, false
	}
	return o.HelmPackages, true
}

// SetHelmPackages gets a reference to the given []HelmPackage and assigns it to the HelmPackages field.
func (o *ListHelmPackagesResult) SetHelmPackages(v []HelmPackage) {
	o.HelmPackages = v
}

func (o ListHelmPackagesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHelmPackagesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HelmPackages) {
		toSerialize["helmPackages"] = o.HelmPackages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListHelmPackagesResult) UnmarshalJSON(data []byte) (err error) {
	varListHelmPackagesResult := _ListHelmPackagesResult{}

	err = json.Unmarshal(data, &varListHelmPackagesResult)

	if err != nil {
		return err
	}

	*o = ListHelmPackagesResult(varListHelmPackagesResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "helmPackages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListHelmPackagesResult struct {
	value *ListHelmPackagesResult
	isSet bool
}

func (v NullableListHelmPackagesResult) Get() *ListHelmPackagesResult {
	return v.value
}

func (v *NullableListHelmPackagesResult) Set(val *ListHelmPackagesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListHelmPackagesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListHelmPackagesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHelmPackagesResult(val *ListHelmPackagesResult) *NullableListHelmPackagesResult {
	return &NullableListHelmPackagesResult{value: val, isSet: true}
}

func (v NullableListHelmPackagesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHelmPackagesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


