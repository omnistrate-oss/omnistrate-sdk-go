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

// checks if the HelmPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPackage{}

// HelmPackage struct for HelmPackage
type HelmPackage struct {
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
	// The namespace where the Helm package should be installed
	Namespace string `json:"namespace"`
	AdditionalProperties map[string]interface{}
}

type _HelmPackage HelmPackage

// NewHelmPackage instantiates a new HelmPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPackage(chartName string, chartRepoName string, chartRepoUrl string, chartVersion string, namespace string) *HelmPackage {
	this := HelmPackage{}
	this.ChartName = chartName
	this.ChartRepoName = chartRepoName
	this.ChartRepoUrl = chartRepoUrl
	this.ChartVersion = chartVersion
	this.Namespace = namespace
	return &this
}

// NewHelmPackageWithDefaults instantiates a new HelmPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPackageWithDefaults() *HelmPackage {
	this := HelmPackage{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *HelmPackage) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *HelmPackage) SetChartName(v string) {
	o.ChartName = v
}

// GetChartRepoName returns the ChartRepoName field value
func (o *HelmPackage) GetChartRepoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartRepoName
}

// GetChartRepoNameOk returns a tuple with the ChartRepoName field value
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetChartRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartRepoName, true
}

// SetChartRepoName sets field value
func (o *HelmPackage) SetChartRepoName(v string) {
	o.ChartRepoName = v
}

// GetChartRepoUrl returns the ChartRepoUrl field value
func (o *HelmPackage) GetChartRepoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartRepoUrl
}

// GetChartRepoUrlOk returns a tuple with the ChartRepoUrl field value
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetChartRepoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartRepoUrl, true
}

// SetChartRepoUrl sets field value
func (o *HelmPackage) SetChartRepoUrl(v string) {
	o.ChartRepoUrl = v
}

// GetChartValues returns the ChartValues field value if set, zero value otherwise.
func (o *HelmPackage) GetChartValues() map[string]interface{} {
	if o == nil || IsNil(o.ChartValues) {
		var ret map[string]interface{}
		return ret
	}
	return o.ChartValues
}

// GetChartValuesOk returns a tuple with the ChartValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetChartValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChartValues) {
		return map[string]interface{}{}, false
	}
	return o.ChartValues, true
}

// HasChartValues returns a boolean if a field has been set.
func (o *HelmPackage) HasChartValues() bool {
	if o != nil && !IsNil(o.ChartValues) {
		return true
	}

	return false
}

// SetChartValues gets a reference to the given map[string]interface{} and assigns it to the ChartValues field.
func (o *HelmPackage) SetChartValues(v map[string]interface{}) {
	o.ChartValues = v
}

// GetChartVersion returns the ChartVersion field value
func (o *HelmPackage) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *HelmPackage) SetChartVersion(v string) {
	o.ChartVersion = v
}

// GetNamespace returns the Namespace field value
func (o *HelmPackage) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *HelmPackage) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *HelmPackage) SetNamespace(v string) {
	o.Namespace = v
}

func (o HelmPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartRepoName"] = o.ChartRepoName
	toSerialize["chartRepoUrl"] = o.ChartRepoUrl
	if !IsNil(o.ChartValues) {
		toSerialize["chartValues"] = o.ChartValues
	}
	toSerialize["chartVersion"] = o.ChartVersion
	toSerialize["namespace"] = o.Namespace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmPackage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chartName",
		"chartRepoName",
		"chartRepoUrl",
		"chartVersion",
		"namespace",
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

	varHelmPackage := _HelmPackage{}

	err = json.Unmarshal(data, &varHelmPackage)

	if err != nil {
		return err
	}

	*o = HelmPackage(varHelmPackage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chartName")
		delete(additionalProperties, "chartRepoName")
		delete(additionalProperties, "chartRepoUrl")
		delete(additionalProperties, "chartValues")
		delete(additionalProperties, "chartVersion")
		delete(additionalProperties, "namespace")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPackage struct {
	value *HelmPackage
	isSet bool
}

func (v NullableHelmPackage) Get() *HelmPackage {
	return v.value
}

func (v *NullableHelmPackage) Set(val *HelmPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPackage(val *HelmPackage) *NullableHelmPackage {
	return &NullableHelmPackage{value: val, isSet: true}
}

func (v NullableHelmPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


