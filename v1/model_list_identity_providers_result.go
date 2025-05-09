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

// checks if the ListIdentityProvidersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIdentityProvidersResult{}

// ListIdentityProvidersResult struct for ListIdentityProvidersResult
type ListIdentityProvidersResult struct {
	// The list of Identity Providers
	IdentityProviders []DescribeIdentityProviderResult `json:"identityProviders"`
	AdditionalProperties map[string]interface{}
}

type _ListIdentityProvidersResult ListIdentityProvidersResult

// NewListIdentityProvidersResult instantiates a new ListIdentityProvidersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIdentityProvidersResult(identityProviders []DescribeIdentityProviderResult) *ListIdentityProvidersResult {
	this := ListIdentityProvidersResult{}
	this.IdentityProviders = identityProviders
	return &this
}

// NewListIdentityProvidersResultWithDefaults instantiates a new ListIdentityProvidersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIdentityProvidersResultWithDefaults() *ListIdentityProvidersResult {
	this := ListIdentityProvidersResult{}
	return &this
}

// GetIdentityProviders returns the IdentityProviders field value
func (o *ListIdentityProvidersResult) GetIdentityProviders() []DescribeIdentityProviderResult {
	if o == nil {
		var ret []DescribeIdentityProviderResult
		return ret
	}

	return o.IdentityProviders
}

// GetIdentityProvidersOk returns a tuple with the IdentityProviders field value
// and a boolean to check if the value has been set.
func (o *ListIdentityProvidersResult) GetIdentityProvidersOk() ([]DescribeIdentityProviderResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProviders, true
}

// SetIdentityProviders sets field value
func (o *ListIdentityProvidersResult) SetIdentityProviders(v []DescribeIdentityProviderResult) {
	o.IdentityProviders = v
}

func (o ListIdentityProvidersResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIdentityProvidersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identityProviders"] = o.IdentityProviders

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListIdentityProvidersResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identityProviders",
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

	varListIdentityProvidersResult := _ListIdentityProvidersResult{}

	err = json.Unmarshal(data, &varListIdentityProvidersResult)

	if err != nil {
		return err
	}

	*o = ListIdentityProvidersResult(varListIdentityProvidersResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identityProviders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListIdentityProvidersResult struct {
	value *ListIdentityProvidersResult
	isSet bool
}

func (v NullableListIdentityProvidersResult) Get() *ListIdentityProvidersResult {
	return v.value
}

func (v *NullableListIdentityProvidersResult) Set(val *ListIdentityProvidersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListIdentityProvidersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListIdentityProvidersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIdentityProvidersResult(val *ListIdentityProvidersResult) *NullableListIdentityProvidersResult {
	return &NullableListIdentityProvidersResult{value: val, isSet: true}
}

func (v NullableListIdentityProvidersResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIdentityProvidersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


