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

// checks if the TerraformConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformConfiguration{}

// TerraformConfiguration struct for TerraformConfiguration
type TerraformConfiguration struct {
	GitConfiguration *GitConfiguration `json:"gitConfiguration,omitempty"`
	// The git access tokens for private modules
	PrivateModuleGitAccessTokens *map[string]string `json:"privateModuleGitAccessTokens,omitempty"`
	// The path to the terraform files directory
	TerraformPath string `json:"terraformPath"`
	AdditionalProperties map[string]interface{}
}

type _TerraformConfiguration TerraformConfiguration

// NewTerraformConfiguration instantiates a new TerraformConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformConfiguration(terraformPath string) *TerraformConfiguration {
	this := TerraformConfiguration{}
	this.TerraformPath = terraformPath
	return &this
}

// NewTerraformConfigurationWithDefaults instantiates a new TerraformConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformConfigurationWithDefaults() *TerraformConfiguration {
	this := TerraformConfiguration{}
	return &this
}

// GetGitConfiguration returns the GitConfiguration field value if set, zero value otherwise.
func (o *TerraformConfiguration) GetGitConfiguration() GitConfiguration {
	if o == nil || IsNil(o.GitConfiguration) {
		var ret GitConfiguration
		return ret
	}
	return *o.GitConfiguration
}

// GetGitConfigurationOk returns a tuple with the GitConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformConfiguration) GetGitConfigurationOk() (*GitConfiguration, bool) {
	if o == nil || IsNil(o.GitConfiguration) {
		return nil, false
	}
	return o.GitConfiguration, true
}

// SetGitConfiguration gets a reference to the given GitConfiguration and assigns it to the GitConfiguration field.
func (o *TerraformConfiguration) SetGitConfiguration(v GitConfiguration) {
	o.GitConfiguration = &v
}

// GetPrivateModuleGitAccessTokens returns the PrivateModuleGitAccessTokens field value if set, zero value otherwise.
func (o *TerraformConfiguration) GetPrivateModuleGitAccessTokens() map[string]string {
	if o == nil || IsNil(o.PrivateModuleGitAccessTokens) {
		var ret map[string]string
		return ret
	}
	return *o.PrivateModuleGitAccessTokens
}

// GetPrivateModuleGitAccessTokensOk returns a tuple with the PrivateModuleGitAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformConfiguration) GetPrivateModuleGitAccessTokensOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PrivateModuleGitAccessTokens) {
		return nil, false
	}
	return o.PrivateModuleGitAccessTokens, true
}

// SetPrivateModuleGitAccessTokens gets a reference to the given map[string]string and assigns it to the PrivateModuleGitAccessTokens field.
func (o *TerraformConfiguration) SetPrivateModuleGitAccessTokens(v map[string]string) {
	o.PrivateModuleGitAccessTokens = &v
}

// GetTerraformPath returns the TerraformPath field value
func (o *TerraformConfiguration) GetTerraformPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerraformPath
}

// GetTerraformPathOk returns a tuple with the TerraformPath field value
// and a boolean to check if the value has been set.
func (o *TerraformConfiguration) GetTerraformPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerraformPath, true
}

// SetTerraformPath sets field value
func (o *TerraformConfiguration) SetTerraformPath(v string) {
	o.TerraformPath = v
}

func (o TerraformConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GitConfiguration) {
		toSerialize["gitConfiguration"] = o.GitConfiguration
	}
	if !IsNil(o.PrivateModuleGitAccessTokens) {
		toSerialize["privateModuleGitAccessTokens"] = o.PrivateModuleGitAccessTokens
	}
	toSerialize["terraformPath"] = o.TerraformPath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TerraformConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"terraformPath",
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

	varTerraformConfiguration := _TerraformConfiguration{}

	err = json.Unmarshal(data, &varTerraformConfiguration)

	if err != nil {
		return err
	}

	*o = TerraformConfiguration(varTerraformConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gitConfiguration")
		delete(additionalProperties, "privateModuleGitAccessTokens")
		delete(additionalProperties, "terraformPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerraformConfiguration struct {
	value *TerraformConfiguration
	isSet bool
}

func (v NullableTerraformConfiguration) Get() *TerraformConfiguration {
	return v.value
}

func (v *NullableTerraformConfiguration) Set(val *TerraformConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformConfiguration(val *TerraformConfiguration) *NullableTerraformConfiguration {
	return &NullableTerraformConfiguration{value: val, isSet: true}
}

func (v NullableTerraformConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


