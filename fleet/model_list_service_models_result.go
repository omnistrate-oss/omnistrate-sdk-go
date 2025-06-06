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

// checks if the ListServiceModelsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceModelsResult{}

// ListServiceModelsResult struct for ListServiceModelsResult
type ListServiceModelsResult struct {
	// The service model IDs
	Ids []string `json:"ids"`
	// The next token to use when listing service models
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceModelsResult ListServiceModelsResult

// NewListServiceModelsResult instantiates a new ListServiceModelsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceModelsResult(ids []string) *ListServiceModelsResult {
	this := ListServiceModelsResult{}
	this.Ids = ids
	return &this
}

// NewListServiceModelsResultWithDefaults instantiates a new ListServiceModelsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceModelsResultWithDefaults() *ListServiceModelsResult {
	this := ListServiceModelsResult{}
	return &this
}

// GetIds returns the Ids field value
func (o *ListServiceModelsResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListServiceModelsResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListServiceModelsResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceModelsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceModelsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListServiceModelsResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceModelsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListServiceModelsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceModelsResult) ToMap() (map[string]interface{}, error) {
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

func (o *ListServiceModelsResult) UnmarshalJSON(data []byte) (err error) {
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

	varListServiceModelsResult := _ListServiceModelsResult{}

	err = json.Unmarshal(data, &varListServiceModelsResult)

	if err != nil {
		return err
	}

	*o = ListServiceModelsResult(varListServiceModelsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceModelsResult struct {
	value *ListServiceModelsResult
	isSet bool
}

func (v NullableListServiceModelsResult) Get() *ListServiceModelsResult {
	return v.value
}

func (v *NullableListServiceModelsResult) Set(val *ListServiceModelsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceModelsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceModelsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceModelsResult(val *ListServiceModelsResult) *NullableListServiceModelsResult {
	return &NullableListServiceModelsResult{value: val, isSet: true}
}

func (v NullableListServiceModelsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceModelsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


