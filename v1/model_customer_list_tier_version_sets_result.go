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

// checks if the CustomerListTierVersionSetsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerListTierVersionSetsResult{}

// CustomerListTierVersionSetsResult struct for CustomerListTierVersionSetsResult
type CustomerListTierVersionSetsResult struct {
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// List of product-tier version sets.
	TierVersionSets []TierVersionSet `json:"tierVersionSets"`
	AdditionalProperties map[string]interface{}
}

type _CustomerListTierVersionSetsResult CustomerListTierVersionSetsResult

// NewCustomerListTierVersionSetsResult instantiates a new CustomerListTierVersionSetsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerListTierVersionSetsResult(tierVersionSets []TierVersionSet) *CustomerListTierVersionSetsResult {
	this := CustomerListTierVersionSetsResult{}
	this.TierVersionSets = tierVersionSets
	return &this
}

// NewCustomerListTierVersionSetsResultWithDefaults instantiates a new CustomerListTierVersionSetsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerListTierVersionSetsResultWithDefaults() *CustomerListTierVersionSetsResult {
	this := CustomerListTierVersionSetsResult{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerListTierVersionSetsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListTierVersionSetsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerListTierVersionSetsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTierVersionSets returns the TierVersionSets field value
func (o *CustomerListTierVersionSetsResult) GetTierVersionSets() []TierVersionSet {
	if o == nil {
		var ret []TierVersionSet
		return ret
	}

	return o.TierVersionSets
}

// GetTierVersionSetsOk returns a tuple with the TierVersionSets field value
// and a boolean to check if the value has been set.
func (o *CustomerListTierVersionSetsResult) GetTierVersionSetsOk() ([]TierVersionSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TierVersionSets, true
}

// SetTierVersionSets sets field value
func (o *CustomerListTierVersionSetsResult) SetTierVersionSets(v []TierVersionSet) {
	o.TierVersionSets = v
}

func (o CustomerListTierVersionSetsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerListTierVersionSetsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["tierVersionSets"] = o.TierVersionSets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerListTierVersionSetsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tierVersionSets",
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

	varCustomerListTierVersionSetsResult := _CustomerListTierVersionSetsResult{}

	err = json.Unmarshal(data, &varCustomerListTierVersionSetsResult)

	if err != nil {
		return err
	}

	*o = CustomerListTierVersionSetsResult(varCustomerListTierVersionSetsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "tierVersionSets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerListTierVersionSetsResult struct {
	value *CustomerListTierVersionSetsResult
	isSet bool
}

func (v NullableCustomerListTierVersionSetsResult) Get() *CustomerListTierVersionSetsResult {
	return v.value
}

func (v *NullableCustomerListTierVersionSetsResult) Set(val *CustomerListTierVersionSetsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerListTierVersionSetsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerListTierVersionSetsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerListTierVersionSetsResult(val *CustomerListTierVersionSetsResult) *NullableCustomerListTierVersionSetsResult {
	return &NullableCustomerListTierVersionSetsResult{value: val, isSet: true}
}

func (v NullableCustomerListTierVersionSetsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerListTierVersionSetsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


