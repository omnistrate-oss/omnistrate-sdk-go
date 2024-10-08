/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
)

// checks if the ListVUnitsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVUnitsResult{}

// ListVUnitsResult struct for ListVUnitsResult
type ListVUnitsResult struct {
	// List of VUnit IDs per Cloud provider
	Ids *map[string][]string `json:"ids,omitempty"`
	// Next page token
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewListVUnitsResult instantiates a new ListVUnitsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVUnitsResult() *ListVUnitsResult {
	this := ListVUnitsResult{}
	return &this
}

// NewListVUnitsResultWithDefaults instantiates a new ListVUnitsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVUnitsResultWithDefaults() *ListVUnitsResult {
	this := ListVUnitsResult{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListVUnitsResult) GetIds() map[string][]string {
	if o == nil || IsNil(o.Ids) {
		var ret map[string][]string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVUnitsResult) GetIdsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// SetIds gets a reference to the given map[string][]string and assigns it to the Ids field.
func (o *ListVUnitsResult) SetIds(v map[string][]string) {
	o.Ids = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListVUnitsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVUnitsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListVUnitsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListVUnitsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVUnitsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableListVUnitsResult struct {
	value *ListVUnitsResult
	isSet bool
}

func (v NullableListVUnitsResult) Get() *ListVUnitsResult {
	return v.value
}

func (v *NullableListVUnitsResult) Set(val *ListVUnitsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListVUnitsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListVUnitsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVUnitsResult(val *ListVUnitsResult) *NullableListVUnitsResult {
	return &NullableListVUnitsResult{value: val, isSet: true}
}

func (v NullableListVUnitsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVUnitsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


