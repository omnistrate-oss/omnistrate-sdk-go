/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListInputParametersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInputParametersResult{}

// ListInputParametersResult struct for ListInputParametersResult
type ListInputParametersResult struct {
	// List of input parameter IDs
	Ids []string `json:"ids"`
	// The input parameters
	InputParameters []DescribeInputParameterResult `json:"inputParameters,omitempty"`
	// Token to retrieve the next page of results
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

type _ListInputParametersResult ListInputParametersResult

// NewListInputParametersResult instantiates a new ListInputParametersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInputParametersResult(ids []string) *ListInputParametersResult {
	this := ListInputParametersResult{}
	this.Ids = ids
	return &this
}

// NewListInputParametersResultWithDefaults instantiates a new ListInputParametersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInputParametersResultWithDefaults() *ListInputParametersResult {
	this := ListInputParametersResult{}
	return &this
}

// GetIds returns the Ids field value
func (o *ListInputParametersResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListInputParametersResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListInputParametersResult) SetIds(v []string) {
	o.Ids = v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise.
func (o *ListInputParametersResult) GetInputParameters() []DescribeInputParameterResult {
	if o == nil || IsNil(o.InputParameters) {
		var ret []DescribeInputParameterResult
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInputParametersResult) GetInputParametersOk() ([]DescribeInputParameterResult, bool) {
	if o == nil || IsNil(o.InputParameters) {
		return nil, false
	}
	return o.InputParameters, true
}

// SetInputParameters gets a reference to the given []DescribeInputParameterResult and assigns it to the InputParameters field.
func (o *ListInputParametersResult) SetInputParameters(v []DescribeInputParameterResult) {
	o.InputParameters = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListInputParametersResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInputParametersResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListInputParametersResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListInputParametersResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInputParametersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	if !IsNil(o.InputParameters) {
		toSerialize["inputParameters"] = o.InputParameters
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *ListInputParametersResult) UnmarshalJSON(data []byte) (err error) {
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

	varListInputParametersResult := _ListInputParametersResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListInputParametersResult)

	if err != nil {
		return err
	}

	*o = ListInputParametersResult(varListInputParametersResult)

	return err
}

type NullableListInputParametersResult struct {
	value *ListInputParametersResult
	isSet bool
}

func (v NullableListInputParametersResult) Get() *ListInputParametersResult {
	return v.value
}

func (v *NullableListInputParametersResult) Set(val *ListInputParametersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListInputParametersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListInputParametersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInputParametersResult(val *ListInputParametersResult) *NullableListInputParametersResult {
	return &NullableListInputParametersResult{value: val, isSet: true}
}

func (v NullableListInputParametersResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInputParametersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

