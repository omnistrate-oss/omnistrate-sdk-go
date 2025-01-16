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

// checks if the DescribeHelmPackageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeHelmPackageRequest{}

// DescribeHelmPackageRequest struct for DescribeHelmPackageRequest
type DescribeHelmPackageRequest struct {
	// The chart name of the Helm package to describe
	ChartName string `json:"chartName"`
	// The chart version of the Helm package to describe
	ChartVersion string `json:"chartVersion"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeHelmPackageRequest DescribeHelmPackageRequest

// NewDescribeHelmPackageRequest instantiates a new DescribeHelmPackageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeHelmPackageRequest(chartName string, chartVersion string, token string) *DescribeHelmPackageRequest {
	this := DescribeHelmPackageRequest{}
	this.ChartName = chartName
	this.ChartVersion = chartVersion
	this.Token = token
	return &this
}

// NewDescribeHelmPackageRequestWithDefaults instantiates a new DescribeHelmPackageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeHelmPackageRequestWithDefaults() *DescribeHelmPackageRequest {
	this := DescribeHelmPackageRequest{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *DescribeHelmPackageRequest) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *DescribeHelmPackageRequest) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *DescribeHelmPackageRequest) SetChartName(v string) {
	o.ChartName = v
}

// GetChartVersion returns the ChartVersion field value
func (o *DescribeHelmPackageRequest) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *DescribeHelmPackageRequest) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *DescribeHelmPackageRequest) SetChartVersion(v string) {
	o.ChartVersion = v
}

// GetToken returns the Token field value
func (o *DescribeHelmPackageRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeHelmPackageRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeHelmPackageRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeHelmPackageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeHelmPackageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartVersion"] = o.ChartVersion
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeHelmPackageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chartName",
		"chartVersion",
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

	varDescribeHelmPackageRequest := _DescribeHelmPackageRequest{}

	err = json.Unmarshal(data, &varDescribeHelmPackageRequest)

	if err != nil {
		return err
	}

	*o = DescribeHelmPackageRequest(varDescribeHelmPackageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chartName")
		delete(additionalProperties, "chartVersion")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeHelmPackageRequest struct {
	value *DescribeHelmPackageRequest
	isSet bool
}

func (v NullableDescribeHelmPackageRequest) Get() *DescribeHelmPackageRequest {
	return v.value
}

func (v *NullableDescribeHelmPackageRequest) Set(val *DescribeHelmPackageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeHelmPackageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeHelmPackageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeHelmPackageRequest(val *DescribeHelmPackageRequest) *NullableDescribeHelmPackageRequest {
	return &NullableDescribeHelmPackageRequest{value: val, isSet: true}
}

func (v NullableDescribeHelmPackageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeHelmPackageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


