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

// checks if the ListHelmPackagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHelmPackagesRequest{}

// ListHelmPackagesRequest struct for ListHelmPackagesRequest
type ListHelmPackagesRequest struct {
	// The chart name of the Helm package to filter by
	ChartName *string `json:"chartName,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListHelmPackagesRequest ListHelmPackagesRequest

// NewListHelmPackagesRequest instantiates a new ListHelmPackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHelmPackagesRequest(token string) *ListHelmPackagesRequest {
	this := ListHelmPackagesRequest{}
	this.Token = token
	return &this
}

// NewListHelmPackagesRequestWithDefaults instantiates a new ListHelmPackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHelmPackagesRequestWithDefaults() *ListHelmPackagesRequest {
	this := ListHelmPackagesRequest{}
	return &this
}

// GetChartName returns the ChartName field value if set, zero value otherwise.
func (o *ListHelmPackagesRequest) GetChartName() string {
	if o == nil || IsNil(o.ChartName) {
		var ret string
		return ret
	}
	return *o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHelmPackagesRequest) GetChartNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChartName) {
		return nil, false
	}
	return o.ChartName, true
}

// SetChartName gets a reference to the given string and assigns it to the ChartName field.
func (o *ListHelmPackagesRequest) SetChartName(v string) {
	o.ChartName = &v
}

// GetToken returns the Token field value
func (o *ListHelmPackagesRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListHelmPackagesRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListHelmPackagesRequest) SetToken(v string) {
	o.Token = v
}

func (o ListHelmPackagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHelmPackagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartName) {
		toSerialize["chartName"] = o.ChartName
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListHelmPackagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varListHelmPackagesRequest := _ListHelmPackagesRequest{}

	err = json.Unmarshal(data, &varListHelmPackagesRequest)

	if err != nil {
		return err
	}

	*o = ListHelmPackagesRequest(varListHelmPackagesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chartName")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListHelmPackagesRequest struct {
	value *ListHelmPackagesRequest
	isSet bool
}

func (v NullableListHelmPackagesRequest) Get() *ListHelmPackagesRequest {
	return v.value
}

func (v *NullableListHelmPackagesRequest) Set(val *ListHelmPackagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListHelmPackagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListHelmPackagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHelmPackagesRequest(val *ListHelmPackagesRequest) *NullableListHelmPackagesRequest {
	return &NullableListHelmPackagesRequest{value: val, isSet: true}
}

func (v NullableListHelmPackagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHelmPackagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


