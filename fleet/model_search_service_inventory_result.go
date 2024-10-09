/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SearchServiceInventoryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchServiceInventoryResult{}

// SearchServiceInventoryResult struct for SearchServiceInventoryResult
type SearchServiceInventoryResult struct {
	// The search results
	Results []SearchRecord `json:"results"`
	// The service ID this workflow belongs to.
	ServiceId string `json:"serviceId"`
}

type _SearchServiceInventoryResult SearchServiceInventoryResult

// NewSearchServiceInventoryResult instantiates a new SearchServiceInventoryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchServiceInventoryResult(results []SearchRecord, serviceId string) *SearchServiceInventoryResult {
	this := SearchServiceInventoryResult{}
	this.Results = results
	this.ServiceId = serviceId
	return &this
}

// NewSearchServiceInventoryResultWithDefaults instantiates a new SearchServiceInventoryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchServiceInventoryResultWithDefaults() *SearchServiceInventoryResult {
	this := SearchServiceInventoryResult{}
	return &this
}

// GetResults returns the Results field value
func (o *SearchServiceInventoryResult) GetResults() []SearchRecord {
	if o == nil {
		var ret []SearchRecord
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SearchServiceInventoryResult) GetResultsOk() ([]SearchRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *SearchServiceInventoryResult) SetResults(v []SearchRecord) {
	o.Results = v
}

// GetServiceId returns the ServiceId field value
func (o *SearchServiceInventoryResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *SearchServiceInventoryResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *SearchServiceInventoryResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o SearchServiceInventoryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchServiceInventoryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *SearchServiceInventoryResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
		"serviceId",
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

	varSearchServiceInventoryResult := _SearchServiceInventoryResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchServiceInventoryResult)

	if err != nil {
		return err
	}

	*o = SearchServiceInventoryResult(varSearchServiceInventoryResult)

	return err
}

type NullableSearchServiceInventoryResult struct {
	value *SearchServiceInventoryResult
	isSet bool
}

func (v NullableSearchServiceInventoryResult) Get() *SearchServiceInventoryResult {
	return v.value
}

func (v *NullableSearchServiceInventoryResult) Set(val *SearchServiceInventoryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchServiceInventoryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchServiceInventoryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchServiceInventoryResult(val *SearchServiceInventoryResult) *NullableSearchServiceInventoryResult {
	return &NullableSearchServiceInventoryResult{value: val, isSet: true}
}

func (v NullableSearchServiceInventoryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchServiceInventoryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


