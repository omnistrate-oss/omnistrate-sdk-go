/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OperatorHelmChartDependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorHelmChartDependency{}

// OperatorHelmChartDependency struct for OperatorHelmChartDependency
type OperatorHelmChartDependency struct {
	// The name of the helm chart
	ChartName string `json:"chartName"`
	// The version of the helm chart
	ChartVersion string `json:"chartVersion"`
}

type _OperatorHelmChartDependency OperatorHelmChartDependency

// NewOperatorHelmChartDependency instantiates a new OperatorHelmChartDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorHelmChartDependency(chartName string, chartVersion string) *OperatorHelmChartDependency {
	this := OperatorHelmChartDependency{}
	this.ChartName = chartName
	this.ChartVersion = chartVersion
	return &this
}

// NewOperatorHelmChartDependencyWithDefaults instantiates a new OperatorHelmChartDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorHelmChartDependencyWithDefaults() *OperatorHelmChartDependency {
	this := OperatorHelmChartDependency{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *OperatorHelmChartDependency) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *OperatorHelmChartDependency) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *OperatorHelmChartDependency) SetChartName(v string) {
	o.ChartName = v
}

// GetChartVersion returns the ChartVersion field value
func (o *OperatorHelmChartDependency) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *OperatorHelmChartDependency) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *OperatorHelmChartDependency) SetChartVersion(v string) {
	o.ChartVersion = v
}

func (o OperatorHelmChartDependency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorHelmChartDependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartVersion"] = o.ChartVersion
	return toSerialize, nil
}

func (o *OperatorHelmChartDependency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chartName",
		"chartVersion",
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

	varOperatorHelmChartDependency := _OperatorHelmChartDependency{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperatorHelmChartDependency)

	if err != nil {
		return err
	}

	*o = OperatorHelmChartDependency(varOperatorHelmChartDependency)

	return err
}

type NullableOperatorHelmChartDependency struct {
	value *OperatorHelmChartDependency
	isSet bool
}

func (v NullableOperatorHelmChartDependency) Get() *OperatorHelmChartDependency {
	return v.value
}

func (v *NullableOperatorHelmChartDependency) Set(val *OperatorHelmChartDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorHelmChartDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorHelmChartDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorHelmChartDependency(val *OperatorHelmChartDependency) *NullableOperatorHelmChartDependency {
	return &NullableOperatorHelmChartDependency{value: val, isSet: true}
}

func (v NullableOperatorHelmChartDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorHelmChartDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


