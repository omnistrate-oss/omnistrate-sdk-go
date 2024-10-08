/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OperatorCRDConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorCRDConfiguration{}

// OperatorCRDConfiguration struct for OperatorCRDConfiguration
type OperatorCRDConfiguration struct {
	// The helm chart dependencies for the CRD (including charts necessary to manage the operator) - Optional
	HelmChartDependencies []OperatorHelmChartDependency `json:"helmChartDependencies,omitempty"`
	// The output parameters to export to the user from the CRD
	OutputParameters *map[string]string `json:"outputParameters,omitempty"`
	// The readiness conditions to check for the CRD
	ReadinessConditions map[string]interface{} `json:"readinessConditions,omitempty"`
	// The supplemental files to apply with the CRD
	SupplementalFiles []string `json:"supplementalFiles,omitempty"`
	// The template of the CRD to apply on every deployment
	Template string `json:"template"`
}

type _OperatorCRDConfiguration OperatorCRDConfiguration

// NewOperatorCRDConfiguration instantiates a new OperatorCRDConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorCRDConfiguration(template string) *OperatorCRDConfiguration {
	this := OperatorCRDConfiguration{}
	this.Template = template
	return &this
}

// NewOperatorCRDConfigurationWithDefaults instantiates a new OperatorCRDConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorCRDConfigurationWithDefaults() *OperatorCRDConfiguration {
	this := OperatorCRDConfiguration{}
	return &this
}

// GetHelmChartDependencies returns the HelmChartDependencies field value if set, zero value otherwise.
func (o *OperatorCRDConfiguration) GetHelmChartDependencies() []OperatorHelmChartDependency {
	if o == nil || IsNil(o.HelmChartDependencies) {
		var ret []OperatorHelmChartDependency
		return ret
	}
	return o.HelmChartDependencies
}

// GetHelmChartDependenciesOk returns a tuple with the HelmChartDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorCRDConfiguration) GetHelmChartDependenciesOk() ([]OperatorHelmChartDependency, bool) {
	if o == nil || IsNil(o.HelmChartDependencies) {
		return nil, false
	}
	return o.HelmChartDependencies, true
}

// HasHelmChartDependencies returns a boolean if a field has been set.
func (o *OperatorCRDConfiguration) HasHelmChartDependencies() bool {
	if o != nil && !IsNil(o.HelmChartDependencies) {
		return true
	}

	return false
}

// SetHelmChartDependencies gets a reference to the given []OperatorHelmChartDependency and assigns it to the HelmChartDependencies field.
func (o *OperatorCRDConfiguration) SetHelmChartDependencies(v []OperatorHelmChartDependency) {
	o.HelmChartDependencies = v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise.
func (o *OperatorCRDConfiguration) GetOutputParameters() map[string]string {
	if o == nil || IsNil(o.OutputParameters) {
		var ret map[string]string
		return ret
	}
	return *o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorCRDConfiguration) GetOutputParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.OutputParameters) {
		return nil, false
	}
	return o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *OperatorCRDConfiguration) HasOutputParameters() bool {
	if o != nil && !IsNil(o.OutputParameters) {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given map[string]string and assigns it to the OutputParameters field.
func (o *OperatorCRDConfiguration) SetOutputParameters(v map[string]string) {
	o.OutputParameters = &v
}

// GetReadinessConditions returns the ReadinessConditions field value if set, zero value otherwise.
func (o *OperatorCRDConfiguration) GetReadinessConditions() map[string]interface{} {
	if o == nil || IsNil(o.ReadinessConditions) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReadinessConditions
}

// GetReadinessConditionsOk returns a tuple with the ReadinessConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorCRDConfiguration) GetReadinessConditionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReadinessConditions) {
		return map[string]interface{}{}, false
	}
	return o.ReadinessConditions, true
}

// HasReadinessConditions returns a boolean if a field has been set.
func (o *OperatorCRDConfiguration) HasReadinessConditions() bool {
	if o != nil && !IsNil(o.ReadinessConditions) {
		return true
	}

	return false
}

// SetReadinessConditions gets a reference to the given map[string]interface{} and assigns it to the ReadinessConditions field.
func (o *OperatorCRDConfiguration) SetReadinessConditions(v map[string]interface{}) {
	o.ReadinessConditions = v
}

// GetSupplementalFiles returns the SupplementalFiles field value if set, zero value otherwise.
func (o *OperatorCRDConfiguration) GetSupplementalFiles() []string {
	if o == nil || IsNil(o.SupplementalFiles) {
		var ret []string
		return ret
	}
	return o.SupplementalFiles
}

// GetSupplementalFilesOk returns a tuple with the SupplementalFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorCRDConfiguration) GetSupplementalFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupplementalFiles) {
		return nil, false
	}
	return o.SupplementalFiles, true
}

// HasSupplementalFiles returns a boolean if a field has been set.
func (o *OperatorCRDConfiguration) HasSupplementalFiles() bool {
	if o != nil && !IsNil(o.SupplementalFiles) {
		return true
	}

	return false
}

// SetSupplementalFiles gets a reference to the given []string and assigns it to the SupplementalFiles field.
func (o *OperatorCRDConfiguration) SetSupplementalFiles(v []string) {
	o.SupplementalFiles = v
}

// GetTemplate returns the Template field value
func (o *OperatorCRDConfiguration) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *OperatorCRDConfiguration) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *OperatorCRDConfiguration) SetTemplate(v string) {
	o.Template = v
}

func (o OperatorCRDConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorCRDConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HelmChartDependencies) {
		toSerialize["helmChartDependencies"] = o.HelmChartDependencies
	}
	if !IsNil(o.OutputParameters) {
		toSerialize["outputParameters"] = o.OutputParameters
	}
	if !IsNil(o.ReadinessConditions) {
		toSerialize["readinessConditions"] = o.ReadinessConditions
	}
	if !IsNil(o.SupplementalFiles) {
		toSerialize["supplementalFiles"] = o.SupplementalFiles
	}
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *OperatorCRDConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template",
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

	varOperatorCRDConfiguration := _OperatorCRDConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperatorCRDConfiguration)

	if err != nil {
		return err
	}

	*o = OperatorCRDConfiguration(varOperatorCRDConfiguration)

	return err
}

type NullableOperatorCRDConfiguration struct {
	value *OperatorCRDConfiguration
	isSet bool
}

func (v NullableOperatorCRDConfiguration) Get() *OperatorCRDConfiguration {
	return v.value
}

func (v *NullableOperatorCRDConfiguration) Set(val *OperatorCRDConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorCRDConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorCRDConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorCRDConfiguration(val *OperatorCRDConfiguration) *NullableOperatorCRDConfiguration {
	return &NullableOperatorCRDConfiguration{value: val, isSet: true}
}

func (v NullableOperatorCRDConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorCRDConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


