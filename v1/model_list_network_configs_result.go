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

// checks if the ListNetworkConfigsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworkConfigsResult{}

// ListNetworkConfigsResult struct for ListNetworkConfigsResult
type ListNetworkConfigsResult struct {
	// The list of network config IDs
	Ids []string `json:"ids"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListNetworkConfigsResult ListNetworkConfigsResult

// NewListNetworkConfigsResult instantiates a new ListNetworkConfigsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkConfigsResult(ids []string) *ListNetworkConfigsResult {
	this := ListNetworkConfigsResult{}
	this.Ids = ids
	return &this
}

// NewListNetworkConfigsResultWithDefaults instantiates a new ListNetworkConfigsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkConfigsResultWithDefaults() *ListNetworkConfigsResult {
	this := ListNetworkConfigsResult{}
	return &this
}

// GetIds returns the Ids field value
func (o *ListNetworkConfigsResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListNetworkConfigsResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListNetworkConfigsResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListNetworkConfigsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkConfigsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListNetworkConfigsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListNetworkConfigsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkConfigsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListNetworkConfigsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varListNetworkConfigsResult := _ListNetworkConfigsResult{}

	err = json.Unmarshal(data, &varListNetworkConfigsResult)

	if err != nil {
		return err
	}

	*o = ListNetworkConfigsResult(varListNetworkConfigsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListNetworkConfigsResult struct {
	value *ListNetworkConfigsResult
	isSet bool
}

func (v NullableListNetworkConfigsResult) Get() *ListNetworkConfigsResult {
	return v.value
}

func (v *NullableListNetworkConfigsResult) Set(val *ListNetworkConfigsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListNetworkConfigsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListNetworkConfigsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNetworkConfigsResult(val *ListNetworkConfigsResult) *NullableListNetworkConfigsResult {
	return &NullableListNetworkConfigsResult{value: val, isSet: true}
}

func (v NullableListNetworkConfigsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNetworkConfigsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


