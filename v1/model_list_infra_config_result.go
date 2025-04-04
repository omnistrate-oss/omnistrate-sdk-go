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

// checks if the ListInfraConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInfraConfigResult{}

// ListInfraConfigResult List of Infra Config IDs
type ListInfraConfigResult struct {
	Ids []string `json:"ids"`
	// Token to use for the next page
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListInfraConfigResult ListInfraConfigResult

// NewListInfraConfigResult instantiates a new ListInfraConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInfraConfigResult(ids []string) *ListInfraConfigResult {
	this := ListInfraConfigResult{}
	this.Ids = ids
	return &this
}

// NewListInfraConfigResultWithDefaults instantiates a new ListInfraConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInfraConfigResultWithDefaults() *ListInfraConfigResult {
	this := ListInfraConfigResult{}
	return &this
}

// GetIds returns the Ids field value
func (o *ListInfraConfigResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListInfraConfigResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListInfraConfigResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListInfraConfigResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInfraConfigResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListInfraConfigResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListInfraConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInfraConfigResult) ToMap() (map[string]interface{}, error) {
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

func (o *ListInfraConfigResult) UnmarshalJSON(data []byte) (err error) {
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

	varListInfraConfigResult := _ListInfraConfigResult{}

	err = json.Unmarshal(data, &varListInfraConfigResult)

	if err != nil {
		return err
	}

	*o = ListInfraConfigResult(varListInfraConfigResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListInfraConfigResult struct {
	value *ListInfraConfigResult
	isSet bool
}

func (v NullableListInfraConfigResult) Get() *ListInfraConfigResult {
	return v.value
}

func (v *NullableListInfraConfigResult) Set(val *ListInfraConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListInfraConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListInfraConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInfraConfigResult(val *ListInfraConfigResult) *NullableListInfraConfigResult {
	return &NullableListInfraConfigResult{value: val, isSet: true}
}

func (v NullableListInfraConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInfraConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


