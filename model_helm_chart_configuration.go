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

// checks if the HelmChartConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmChartConfiguration{}

// HelmChartConfiguration struct for HelmChartConfiguration
type HelmChartConfiguration struct {
	// The chart name of the Helm package
	ChartName string `json:"chartName"`
	// The chart repository name of the Helm package
	ChartRepoName string `json:"chartRepoName"`
	// The chart repository URL of the Helm package
	ChartRepoUrl string `json:"chartRepoUrl"`
	// The values of the Helm package
	ChartValues map[string]interface{} `json:"chartValues,omitempty"`
	// The chart version of the Helm package
	ChartVersion string `json:"chartVersion"`
}

type _HelmChartConfiguration HelmChartConfiguration

// NewHelmChartConfiguration instantiates a new HelmChartConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmChartConfiguration(chartName string, chartRepoName string, chartRepoUrl string, chartVersion string) *HelmChartConfiguration {
	this := HelmChartConfiguration{}
	this.ChartName = chartName
	this.ChartRepoName = chartRepoName
	this.ChartRepoUrl = chartRepoUrl
	this.ChartVersion = chartVersion
	return &this
}

// NewHelmChartConfigurationWithDefaults instantiates a new HelmChartConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmChartConfigurationWithDefaults() *HelmChartConfiguration {
	this := HelmChartConfiguration{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *HelmChartConfiguration) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *HelmChartConfiguration) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *HelmChartConfiguration) SetChartName(v string) {
	o.ChartName = v
}

// GetChartRepoName returns the ChartRepoName field value
func (o *HelmChartConfiguration) GetChartRepoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartRepoName
}

// GetChartRepoNameOk returns a tuple with the ChartRepoName field value
// and a boolean to check if the value has been set.
func (o *HelmChartConfiguration) GetChartRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartRepoName, true
}

// SetChartRepoName sets field value
func (o *HelmChartConfiguration) SetChartRepoName(v string) {
	o.ChartRepoName = v
}

// GetChartRepoUrl returns the ChartRepoUrl field value
func (o *HelmChartConfiguration) GetChartRepoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartRepoUrl
}

// GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field value
// and a boolean to check if the value has been set.
func (o *HelmChartConfiguration) GetChartRepoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartRepoUrl, true
}

// SetChartRepoUrl sets field value
func (o *HelmChartConfiguration) SetChartRepoUrl(v string) {
	o.ChartRepoUrl = v
}

// GetChartValues returns the ChartValues field value if set, zero value otherwise.
func (o *HelmChartConfiguration) GetChartValues() map[string]interface{} {
	if o == nil || IsNil(o.ChartValues) {
		var ret map[string]interface{}
		return ret
	}
	return o.ChartValues
}

// GetChartValuesOk returns a tuple with the ChartValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmChartConfiguration) GetChartValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChartValues) {
		return map[string]interface{}{}, false
	}
	return o.ChartValues, true
}

// HasChartValues returns a boolean if a field has been set.
func (o *HelmChartConfiguration) HasChartValues() bool {
	if o != nil && !IsNil(o.ChartValues) {
		return true
	}

	return false
}

// SetChartValues gets a reference to the given map[string]interface{} and assigns it to the ChartValues field.
func (o *HelmChartConfiguration) SetChartValues(v map[string]interface{}) {
	o.ChartValues = v
}

// GetChartVersion returns the ChartVersion field value
func (o *HelmChartConfiguration) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *HelmChartConfiguration) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *HelmChartConfiguration) SetChartVersion(v string) {
	o.ChartVersion = v
}

func (o HelmChartConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmChartConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartRepoName"] = o.ChartRepoName
	toSerialize["chartRepoUrl"] = o.ChartRepoUrl
	if !IsNil(o.ChartValues) {
		toSerialize["chartValues"] = o.ChartValues
	}
	toSerialize["chartVersion"] = o.ChartVersion
	return toSerialize, nil
}

func (o *HelmChartConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chartName",
		"chartRepoName",
		"chartRepoUrl",
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

	varHelmChartConfiguration := _HelmChartConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHelmChartConfiguration)

	if err != nil {
		return err
	}

	*o = HelmChartConfiguration(varHelmChartConfiguration)

	return err
}

type NullableHelmChartConfiguration struct {
	value *HelmChartConfiguration
	isSet bool
}

func (v NullableHelmChartConfiguration) Get() *HelmChartConfiguration {
	return v.value
}

func (v *NullableHelmChartConfiguration) Set(val *HelmChartConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmChartConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmChartConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmChartConfiguration(val *HelmChartConfiguration) *NullableHelmChartConfiguration {
	return &NullableHelmChartConfiguration{value: val, isSet: true}
}

func (v NullableHelmChartConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmChartConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


