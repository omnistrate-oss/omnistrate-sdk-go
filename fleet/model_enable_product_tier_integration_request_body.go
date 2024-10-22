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

// checks if the EnableProductTierIntegrationRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableProductTierIntegrationRequestBody{}

// EnableProductTierIntegrationRequestBody struct for EnableProductTierIntegrationRequestBody
type EnableProductTierIntegrationRequestBody struct {
	// The configuration parameters for the integration provider.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// Name of the product tier integration provider.
	IntegrationProviderName string `json:"integrationProviderName"`
	// Type of the product tier integration.
	IntegrationType string `json:"integrationType"`
	AdditionalProperties map[string]interface{}
}

type _EnableProductTierIntegrationRequestBody EnableProductTierIntegrationRequestBody

// NewEnableProductTierIntegrationRequestBody instantiates a new EnableProductTierIntegrationRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableProductTierIntegrationRequestBody(integrationProviderName string, integrationType string) *EnableProductTierIntegrationRequestBody {
	this := EnableProductTierIntegrationRequestBody{}
	this.IntegrationProviderName = integrationProviderName
	this.IntegrationType = integrationType
	return &this
}

// NewEnableProductTierIntegrationRequestBodyWithDefaults instantiates a new EnableProductTierIntegrationRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableProductTierIntegrationRequestBodyWithDefaults() *EnableProductTierIntegrationRequestBody {
	this := EnableProductTierIntegrationRequestBody{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *EnableProductTierIntegrationRequestBody) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableProductTierIntegrationRequestBody) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *EnableProductTierIntegrationRequestBody) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *EnableProductTierIntegrationRequestBody) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetIntegrationProviderName returns the IntegrationProviderName field value
func (o *EnableProductTierIntegrationRequestBody) GetIntegrationProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationProviderName
}

// GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field value
// and a boolean to check if the value has been set.
func (o *EnableProductTierIntegrationRequestBody) GetIntegrationProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationProviderName, true
}

// SetIntegrationProviderName sets field value
func (o *EnableProductTierIntegrationRequestBody) SetIntegrationProviderName(v string) {
	o.IntegrationProviderName = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *EnableProductTierIntegrationRequestBody) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *EnableProductTierIntegrationRequestBody) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *EnableProductTierIntegrationRequestBody) SetIntegrationType(v string) {
	o.IntegrationType = v
}

func (o EnableProductTierIntegrationRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableProductTierIntegrationRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	toSerialize["integrationProviderName"] = o.IntegrationProviderName
	toSerialize["integrationType"] = o.IntegrationType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableProductTierIntegrationRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"integrationProviderName",
		"integrationType",
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

	varEnableProductTierIntegrationRequestBody := _EnableProductTierIntegrationRequestBody{}

	err = json.Unmarshal(data, &varEnableProductTierIntegrationRequestBody)

	if err != nil {
		return err
	}

	*o = EnableProductTierIntegrationRequestBody(varEnableProductTierIntegrationRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "integrationProviderName")
		delete(additionalProperties, "integrationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableProductTierIntegrationRequestBody struct {
	value *EnableProductTierIntegrationRequestBody
	isSet bool
}

func (v NullableEnableProductTierIntegrationRequestBody) Get() *EnableProductTierIntegrationRequestBody {
	return v.value
}

func (v *NullableEnableProductTierIntegrationRequestBody) Set(val *EnableProductTierIntegrationRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableProductTierIntegrationRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableProductTierIntegrationRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableProductTierIntegrationRequestBody(val *EnableProductTierIntegrationRequestBody) *NullableEnableProductTierIntegrationRequestBody {
	return &NullableEnableProductTierIntegrationRequestBody{value: val, isSet: true}
}

func (v NullableEnableProductTierIntegrationRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableProductTierIntegrationRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


