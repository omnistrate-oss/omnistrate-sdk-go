/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the ListCustomDomainResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomDomainResult{}

// ListCustomDomainResult List of custom domain IDs
type ListCustomDomainResult struct {
	// The list of custom domains
	CustomDomains []DescribeCustomDomainResult `json:"CustomDomains,omitempty"`
	Ids []string `json:"ids,omitempty"`
	// Token to use for the next page
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCustomDomainResult ListCustomDomainResult

// NewListCustomDomainResult instantiates a new ListCustomDomainResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomDomainResult() *ListCustomDomainResult {
	this := ListCustomDomainResult{}
	return &this
}

// NewListCustomDomainResultWithDefaults instantiates a new ListCustomDomainResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomDomainResultWithDefaults() *ListCustomDomainResult {
	this := ListCustomDomainResult{}
	return &this
}

// GetCustomDomains returns the CustomDomains field value if set, zero value otherwise.
func (o *ListCustomDomainResult) GetCustomDomains() []DescribeCustomDomainResult {
	if o == nil || IsNil(o.CustomDomains) {
		var ret []DescribeCustomDomainResult
		return ret
	}
	return o.CustomDomains
}

// GetCustomDomainsOk returns a tuple with the CustomDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomDomainResult) GetCustomDomainsOk() ([]DescribeCustomDomainResult, bool) {
	if o == nil || IsNil(o.CustomDomains) {
		return nil, false
	}
	return o.CustomDomains, true
}

// SetCustomDomains gets a reference to the given []DescribeCustomDomainResult and assigns it to the CustomDomains field.
func (o *ListCustomDomainResult) SetCustomDomains(v []DescribeCustomDomainResult) {
	o.CustomDomains = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListCustomDomainResult) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomDomainResult) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ListCustomDomainResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListCustomDomainResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomDomainResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListCustomDomainResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListCustomDomainResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomDomainResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomDomains) {
		toSerialize["CustomDomains"] = o.CustomDomains
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCustomDomainResult) UnmarshalJSON(data []byte) (err error) {
	varListCustomDomainResult := _ListCustomDomainResult{}

	err = json.Unmarshal(data, &varListCustomDomainResult)

	if err != nil {
		return err
	}

	*o = ListCustomDomainResult(varListCustomDomainResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "CustomDomains")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCustomDomainResult struct {
	value *ListCustomDomainResult
	isSet bool
}

func (v NullableListCustomDomainResult) Get() *ListCustomDomainResult {
	return v.value
}

func (v *NullableListCustomDomainResult) Set(val *ListCustomDomainResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomDomainResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomDomainResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomDomainResult(val *ListCustomDomainResult) *NullableListCustomDomainResult {
	return &NullableListCustomDomainResult{value: val, isSet: true}
}

func (v NullableListCustomDomainResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomDomainResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


