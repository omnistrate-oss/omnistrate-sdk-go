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

// checks if the ListImageRegistryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageRegistryResult{}

// ListImageRegistryResult List of HTTP API v2 Docker Image Registries
type ListImageRegistryResult struct {
	// List of Image Registry IDs
	Ids []string `json:"ids"`
	// Token to use for the next page
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListImageRegistryResult ListImageRegistryResult

// NewListImageRegistryResult instantiates a new ListImageRegistryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageRegistryResult(ids []string) *ListImageRegistryResult {
	this := ListImageRegistryResult{}
	this.Ids = ids
	return &this
}

// NewListImageRegistryResultWithDefaults instantiates a new ListImageRegistryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageRegistryResultWithDefaults() *ListImageRegistryResult {
	this := ListImageRegistryResult{}
	return &this
}

// GetIds returns the Ids field value
func (o *ListImageRegistryResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListImageRegistryResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListImageRegistryResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListImageRegistryResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageRegistryResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListImageRegistryResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListImageRegistryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListImageRegistryResult) ToMap() (map[string]interface{}, error) {
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

func (o *ListImageRegistryResult) UnmarshalJSON(data []byte) (err error) {
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

	varListImageRegistryResult := _ListImageRegistryResult{}

	err = json.Unmarshal(data, &varListImageRegistryResult)

	if err != nil {
		return err
	}

	*o = ListImageRegistryResult(varListImageRegistryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListImageRegistryResult struct {
	value *ListImageRegistryResult
	isSet bool
}

func (v NullableListImageRegistryResult) Get() *ListImageRegistryResult {
	return v.value
}

func (v *NullableListImageRegistryResult) Set(val *ListImageRegistryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListImageRegistryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListImageRegistryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImageRegistryResult(val *ListImageRegistryResult) *NullableListImageRegistryResult {
	return &NullableListImageRegistryResult{value: val, isSet: true}
}

func (v NullableListImageRegistryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListImageRegistryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


