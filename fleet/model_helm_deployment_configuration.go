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

// checks if the HelmDeploymentConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmDeploymentConfiguration{}

// HelmDeploymentConfiguration struct for HelmDeploymentConfiguration
type HelmDeploymentConfiguration struct {
	// Name of the Helm chart to be deployed
	ChartName string `json:"ChartName"`
	// Version of the Helm chart to be deployed
	ChartVersion string `json:"ChartVersion"`
	// Errors encountered during the Helm deployment
	DeploymentErrors *string `json:"DeploymentErrors,omitempty"`
	// Events associated with the pods in the generic resource deployment
	PodEvents *map[string][]PodEvent `json:"PodEvents,omitempty"`
	// Status of the pods in the generic resource deployment
	PodStatus *map[string]string `json:"PodStatus,omitempty"`
	// Mapping of pods to k8s node names for the generic resource deployment
	PodToHostMapping *map[string]string `json:"PodToHostMapping,omitempty"`
	// Name of the Helm release
	ReleaseName string `json:"ReleaseName"`
	// Namespace where the Helm release will be deployed
	ReleaseNamespace string `json:"ReleaseNamespace"`
	// Status of the Helm release
	ReleaseStatus string `json:"ReleaseStatus"`
	// URL of the Helm chart repository
	RepositoryURL string `json:"RepositoryURL"`
	// Runtime configuration for the Helm deployment
	RuntimeConfiguration map[string]interface{} `json:"RuntimeConfiguration,omitempty"`
	// Deployed values for the Helm chart in YAML format
	Values map[string]interface{} `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmDeploymentConfiguration HelmDeploymentConfiguration

// NewHelmDeploymentConfiguration instantiates a new HelmDeploymentConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmDeploymentConfiguration(chartName string, chartVersion string, releaseName string, releaseNamespace string, releaseStatus string, repositoryURL string) *HelmDeploymentConfiguration {
	this := HelmDeploymentConfiguration{}
	this.ChartName = chartName
	this.ChartVersion = chartVersion
	this.ReleaseName = releaseName
	this.ReleaseNamespace = releaseNamespace
	this.ReleaseStatus = releaseStatus
	this.RepositoryURL = repositoryURL
	return &this
}

// NewHelmDeploymentConfigurationWithDefaults instantiates a new HelmDeploymentConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmDeploymentConfigurationWithDefaults() *HelmDeploymentConfiguration {
	this := HelmDeploymentConfiguration{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *HelmDeploymentConfiguration) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *HelmDeploymentConfiguration) SetChartName(v string) {
	o.ChartName = v
}

// GetChartVersion returns the ChartVersion field value
func (o *HelmDeploymentConfiguration) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *HelmDeploymentConfiguration) SetChartVersion(v string) {
	o.ChartVersion = v
}

// GetDeploymentErrors returns the DeploymentErrors field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetDeploymentErrors() string {
	if o == nil || IsNil(o.DeploymentErrors) {
		var ret string
		return ret
	}
	return *o.DeploymentErrors
}

// GetDeploymentErrorsOk returns a tuple with the DeploymentErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetDeploymentErrorsOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentErrors) {
		return nil, false
	}
	return o.DeploymentErrors, true
}

// HasDeploymentErrors returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasDeploymentErrors() bool {
	if o != nil && !IsNil(o.DeploymentErrors) {
		return true
	}

	return false
}

// SetDeploymentErrors gets a reference to the given string and assigns it to the DeploymentErrors field.
func (o *HelmDeploymentConfiguration) SetDeploymentErrors(v string) {
	o.DeploymentErrors = &v
}

// GetPodEvents returns the PodEvents field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetPodEvents() map[string][]PodEvent {
	if o == nil || IsNil(o.PodEvents) {
		var ret map[string][]PodEvent
		return ret
	}
	return *o.PodEvents
}

// GetPodEventsOk returns a tuple with the PodEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetPodEventsOk() (*map[string][]PodEvent, bool) {
	if o == nil || IsNil(o.PodEvents) {
		return nil, false
	}
	return o.PodEvents, true
}

// HasPodEvents returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasPodEvents() bool {
	if o != nil && !IsNil(o.PodEvents) {
		return true
	}

	return false
}

// SetPodEvents gets a reference to the given map[string][]PodEvent and assigns it to the PodEvents field.
func (o *HelmDeploymentConfiguration) SetPodEvents(v map[string][]PodEvent) {
	o.PodEvents = &v
}

// GetPodStatus returns the PodStatus field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetPodStatus() map[string]string {
	if o == nil || IsNil(o.PodStatus) {
		var ret map[string]string
		return ret
	}
	return *o.PodStatus
}

// GetPodStatusOk returns a tuple with the PodStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetPodStatusOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PodStatus) {
		return nil, false
	}
	return o.PodStatus, true
}

// HasPodStatus returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasPodStatus() bool {
	if o != nil && !IsNil(o.PodStatus) {
		return true
	}

	return false
}

// SetPodStatus gets a reference to the given map[string]string and assigns it to the PodStatus field.
func (o *HelmDeploymentConfiguration) SetPodStatus(v map[string]string) {
	o.PodStatus = &v
}

// GetPodToHostMapping returns the PodToHostMapping field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetPodToHostMapping() map[string]string {
	if o == nil || IsNil(o.PodToHostMapping) {
		var ret map[string]string
		return ret
	}
	return *o.PodToHostMapping
}

// GetPodToHostMappingOk returns a tuple with the PodToHostMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetPodToHostMappingOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PodToHostMapping) {
		return nil, false
	}
	return o.PodToHostMapping, true
}

// HasPodToHostMapping returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasPodToHostMapping() bool {
	if o != nil && !IsNil(o.PodToHostMapping) {
		return true
	}

	return false
}

// SetPodToHostMapping gets a reference to the given map[string]string and assigns it to the PodToHostMapping field.
func (o *HelmDeploymentConfiguration) SetPodToHostMapping(v map[string]string) {
	o.PodToHostMapping = &v
}

// GetReleaseName returns the ReleaseName field value
func (o *HelmDeploymentConfiguration) GetReleaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetReleaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseName, true
}

// SetReleaseName sets field value
func (o *HelmDeploymentConfiguration) SetReleaseName(v string) {
	o.ReleaseName = v
}

// GetReleaseNamespace returns the ReleaseNamespace field value
func (o *HelmDeploymentConfiguration) GetReleaseNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseNamespace
}

// GetReleaseNamespaceOk returns a tuple with the ReleaseNamespace field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetReleaseNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseNamespace, true
}

// SetReleaseNamespace sets field value
func (o *HelmDeploymentConfiguration) SetReleaseNamespace(v string) {
	o.ReleaseNamespace = v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *HelmDeploymentConfiguration) GetReleaseStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetReleaseStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *HelmDeploymentConfiguration) SetReleaseStatus(v string) {
	o.ReleaseStatus = v
}

// GetRepositoryURL returns the RepositoryURL field value
func (o *HelmDeploymentConfiguration) GetRepositoryURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryURL
}

// GetRepositoryURLOk returns a tuple with the RepositoryURL field value
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetRepositoryURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryURL, true
}

// SetRepositoryURL sets field value
func (o *HelmDeploymentConfiguration) SetRepositoryURL(v string) {
	o.RepositoryURL = v
}

// GetRuntimeConfiguration returns the RuntimeConfiguration field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetRuntimeConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.RuntimeConfiguration) {
		var ret map[string]interface{}
		return ret
	}
	return o.RuntimeConfiguration
}

// GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetRuntimeConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RuntimeConfiguration) {
		return map[string]interface{}{}, false
	}
	return o.RuntimeConfiguration, true
}

// HasRuntimeConfiguration returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasRuntimeConfiguration() bool {
	if o != nil && !IsNil(o.RuntimeConfiguration) {
		return true
	}

	return false
}

// SetRuntimeConfiguration gets a reference to the given map[string]interface{} and assigns it to the RuntimeConfiguration field.
func (o *HelmDeploymentConfiguration) SetRuntimeConfiguration(v map[string]interface{}) {
	o.RuntimeConfiguration = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *HelmDeploymentConfiguration) GetValues() map[string]interface{} {
	if o == nil || IsNil(o.Values) {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentConfiguration) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]interface{}{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HelmDeploymentConfiguration) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *HelmDeploymentConfiguration) SetValues(v map[string]interface{}) {
	o.Values = v
}

func (o HelmDeploymentConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmDeploymentConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ChartName"] = o.ChartName
	toSerialize["ChartVersion"] = o.ChartVersion
	if !IsNil(o.DeploymentErrors) {
		toSerialize["DeploymentErrors"] = o.DeploymentErrors
	}
	if !IsNil(o.PodEvents) {
		toSerialize["PodEvents"] = o.PodEvents
	}
	if !IsNil(o.PodStatus) {
		toSerialize["PodStatus"] = o.PodStatus
	}
	if !IsNil(o.PodToHostMapping) {
		toSerialize["PodToHostMapping"] = o.PodToHostMapping
	}
	toSerialize["ReleaseName"] = o.ReleaseName
	toSerialize["ReleaseNamespace"] = o.ReleaseNamespace
	toSerialize["ReleaseStatus"] = o.ReleaseStatus
	toSerialize["RepositoryURL"] = o.RepositoryURL
	if !IsNil(o.RuntimeConfiguration) {
		toSerialize["RuntimeConfiguration"] = o.RuntimeConfiguration
	}
	if !IsNil(o.Values) {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmDeploymentConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ChartName",
		"ChartVersion",
		"ReleaseName",
		"ReleaseNamespace",
		"ReleaseStatus",
		"RepositoryURL",
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

	varHelmDeploymentConfiguration := _HelmDeploymentConfiguration{}

	err = json.Unmarshal(data, &varHelmDeploymentConfiguration)

	if err != nil {
		return err
	}

	*o = HelmDeploymentConfiguration(varHelmDeploymentConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ChartName")
		delete(additionalProperties, "ChartVersion")
		delete(additionalProperties, "DeploymentErrors")
		delete(additionalProperties, "PodEvents")
		delete(additionalProperties, "PodStatus")
		delete(additionalProperties, "PodToHostMapping")
		delete(additionalProperties, "ReleaseName")
		delete(additionalProperties, "ReleaseNamespace")
		delete(additionalProperties, "ReleaseStatus")
		delete(additionalProperties, "RepositoryURL")
		delete(additionalProperties, "RuntimeConfiguration")
		delete(additionalProperties, "Values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmDeploymentConfiguration struct {
	value *HelmDeploymentConfiguration
	isSet bool
}

func (v NullableHelmDeploymentConfiguration) Get() *HelmDeploymentConfiguration {
	return v.value
}

func (v *NullableHelmDeploymentConfiguration) Set(val *HelmDeploymentConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmDeploymentConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmDeploymentConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmDeploymentConfiguration(val *HelmDeploymentConfiguration) *NullableHelmDeploymentConfiguration {
	return &NullableHelmDeploymentConfiguration{value: val, isSet: true}
}

func (v NullableHelmDeploymentConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmDeploymentConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


