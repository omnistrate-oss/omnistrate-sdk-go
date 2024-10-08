/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListSaaSPortalCustomDomainsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSaaSPortalCustomDomainsResult{}

// ListSaaSPortalCustomDomainsResult struct for ListSaaSPortalCustomDomainsResult
type ListSaaSPortalCustomDomainsResult struct {
	// The list of custom domains
	CustomDomains []CustomDomain `json:"customDomains"`
}

type _ListSaaSPortalCustomDomainsResult ListSaaSPortalCustomDomainsResult

// NewListSaaSPortalCustomDomainsResult instantiates a new ListSaaSPortalCustomDomainsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSaaSPortalCustomDomainsResult(customDomains []CustomDomain) *ListSaaSPortalCustomDomainsResult {
	this := ListSaaSPortalCustomDomainsResult{}
	this.CustomDomains = customDomains
	return &this
}

// NewListSaaSPortalCustomDomainsResultWithDefaults instantiates a new ListSaaSPortalCustomDomainsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSaaSPortalCustomDomainsResultWithDefaults() *ListSaaSPortalCustomDomainsResult {
	this := ListSaaSPortalCustomDomainsResult{}
	return &this
}

// GetCustomDomains returns the CustomDomains field value
func (o *ListSaaSPortalCustomDomainsResult) GetCustomDomains() []CustomDomain {
	if o == nil {
		var ret []CustomDomain
		return ret
	}

	return o.CustomDomains
}

// GetCustomDomainsOk returns a tuple with the CustomDomains field value
// and a boolean to check if the value has been set.
func (o *ListSaaSPortalCustomDomainsResult) GetCustomDomainsOk() ([]CustomDomain, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomDomains, true
}

// SetCustomDomains sets field value
func (o *ListSaaSPortalCustomDomainsResult) SetCustomDomains(v []CustomDomain) {
	o.CustomDomains = v
}

func (o ListSaaSPortalCustomDomainsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSaaSPortalCustomDomainsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customDomains"] = o.CustomDomains
	return toSerialize, nil
}

func (o *ListSaaSPortalCustomDomainsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customDomains",
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

	varListSaaSPortalCustomDomainsResult := _ListSaaSPortalCustomDomainsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListSaaSPortalCustomDomainsResult)

	if err != nil {
		return err
	}

	*o = ListSaaSPortalCustomDomainsResult(varListSaaSPortalCustomDomainsResult)

	return err
}

type NullableListSaaSPortalCustomDomainsResult struct {
	value *ListSaaSPortalCustomDomainsResult
	isSet bool
}

func (v NullableListSaaSPortalCustomDomainsResult) Get() *ListSaaSPortalCustomDomainsResult {
	return v.value
}

func (v *NullableListSaaSPortalCustomDomainsResult) Set(val *ListSaaSPortalCustomDomainsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListSaaSPortalCustomDomainsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListSaaSPortalCustomDomainsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSaaSPortalCustomDomainsResult(val *ListSaaSPortalCustomDomainsResult) *NullableListSaaSPortalCustomDomainsResult {
	return &NullableListSaaSPortalCustomDomainsResult{value: val, isSet: true}
}

func (v NullableListSaaSPortalCustomDomainsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSaaSPortalCustomDomainsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


