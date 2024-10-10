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

// checks if the SearchServiceInventoryRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchServiceInventoryRequestBody{}

// SearchServiceInventoryRequestBody struct for SearchServiceInventoryRequestBody
type SearchServiceInventoryRequestBody struct {
	// The search query.
	Query string `json:"query"`
	AdditionalProperties map[string]interface{}
}

type _SearchServiceInventoryRequestBody SearchServiceInventoryRequestBody

// NewSearchServiceInventoryRequestBody instantiates a new SearchServiceInventoryRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchServiceInventoryRequestBody(query string) *SearchServiceInventoryRequestBody {
	this := SearchServiceInventoryRequestBody{}
	this.Query = query
	return &this
}

// NewSearchServiceInventoryRequestBodyWithDefaults instantiates a new SearchServiceInventoryRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchServiceInventoryRequestBodyWithDefaults() *SearchServiceInventoryRequestBody {
	this := SearchServiceInventoryRequestBody{}
	return &this
}

// GetQuery returns the Query field value
func (o *SearchServiceInventoryRequestBody) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchServiceInventoryRequestBody) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchServiceInventoryRequestBody) SetQuery(v string) {
	o.Query = v
}

func (o SearchServiceInventoryRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchServiceInventoryRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchServiceInventoryRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varSearchServiceInventoryRequestBody := _SearchServiceInventoryRequestBody{}

	err = json.Unmarshal(data, &varSearchServiceInventoryRequestBody)

	if err != nil {
		return err
	}

	*o = SearchServiceInventoryRequestBody(varSearchServiceInventoryRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchServiceInventoryRequestBody struct {
	value *SearchServiceInventoryRequestBody
	isSet bool
}

func (v NullableSearchServiceInventoryRequestBody) Get() *SearchServiceInventoryRequestBody {
	return v.value
}

func (v *NullableSearchServiceInventoryRequestBody) Set(val *SearchServiceInventoryRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchServiceInventoryRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchServiceInventoryRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchServiceInventoryRequestBody(val *SearchServiceInventoryRequestBody) *NullableSearchServiceInventoryRequestBody {
	return &NullableSearchServiceInventoryRequestBody{value: val, isSet: true}
}

func (v NullableSearchServiceInventoryRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchServiceInventoryRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


