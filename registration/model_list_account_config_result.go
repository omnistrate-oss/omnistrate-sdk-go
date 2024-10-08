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

// checks if the ListAccountConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccountConfigResult{}

// ListAccountConfigResult struct for ListAccountConfigResult
type ListAccountConfigResult struct {
	// The list of account configs
	AccountConfigs []DescribeAccountConfigResult `json:"accountConfigs,omitempty"`
	Ids []string `json:"ids,omitempty"`
	// Token to use for the next page
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewListAccountConfigResult instantiates a new ListAccountConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccountConfigResult() *ListAccountConfigResult {
	this := ListAccountConfigResult{}
	return &this
}

// NewListAccountConfigResultWithDefaults instantiates a new ListAccountConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccountConfigResultWithDefaults() *ListAccountConfigResult {
	this := ListAccountConfigResult{}
	return &this
}

// GetAccountConfigs returns the AccountConfigs field value if set, zero value otherwise.
func (o *ListAccountConfigResult) GetAccountConfigs() []DescribeAccountConfigResult {
	if o == nil || IsNil(o.AccountConfigs) {
		var ret []DescribeAccountConfigResult
		return ret
	}
	return o.AccountConfigs
}

// GetAccountConfigsOk returns a tuple with the AccountConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountConfigResult) GetAccountConfigsOk() ([]DescribeAccountConfigResult, bool) {
	if o == nil || IsNil(o.AccountConfigs) {
		return nil, false
	}
	return o.AccountConfigs, true
}

// HasAccountConfigs returns a boolean if a field has been set.
func (o *ListAccountConfigResult) HasAccountConfigs() bool {
	if o != nil && !IsNil(o.AccountConfigs) {
		return true
	}

	return false
}

// SetAccountConfigs gets a reference to the given []DescribeAccountConfigResult and assigns it to the AccountConfigs field.
func (o *ListAccountConfigResult) SetAccountConfigs(v []DescribeAccountConfigResult) {
	o.AccountConfigs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListAccountConfigResult) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountConfigResult) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListAccountConfigResult) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ListAccountConfigResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListAccountConfigResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountConfigResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListAccountConfigResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListAccountConfigResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListAccountConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccountConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountConfigs) {
		toSerialize["accountConfigs"] = o.AccountConfigs
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableListAccountConfigResult struct {
	value *ListAccountConfigResult
	isSet bool
}

func (v NullableListAccountConfigResult) Get() *ListAccountConfigResult {
	return v.value
}

func (v *NullableListAccountConfigResult) Set(val *ListAccountConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccountConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccountConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccountConfigResult(val *ListAccountConfigResult) *NullableListAccountConfigResult {
	return &NullableListAccountConfigResult{value: val, isSet: true}
}

func (v NullableListAccountConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccountConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


