/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DisableProductTierIntegrationRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableProductTierIntegrationRequestBody{}

// DisableProductTierIntegrationRequestBody struct for DisableProductTierIntegrationRequestBody
type DisableProductTierIntegrationRequestBody struct {
	// Name of the product tier integration provider.
	IntegrationProviderName string `json:"integrationProviderName"`
	// Type of the product tier integration.
	IntegrationType string `json:"integrationType"`
}

type _DisableProductTierIntegrationRequestBody DisableProductTierIntegrationRequestBody

// NewDisableProductTierIntegrationRequestBody instantiates a new DisableProductTierIntegrationRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableProductTierIntegrationRequestBody(integrationProviderName string, integrationType string) *DisableProductTierIntegrationRequestBody {
	this := DisableProductTierIntegrationRequestBody{}
	this.IntegrationProviderName = integrationProviderName
	this.IntegrationType = integrationType
	return &this
}

// NewDisableProductTierIntegrationRequestBodyWithDefaults instantiates a new DisableProductTierIntegrationRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableProductTierIntegrationRequestBodyWithDefaults() *DisableProductTierIntegrationRequestBody {
	this := DisableProductTierIntegrationRequestBody{}
	return &this
}

// GetIntegrationProviderName returns the IntegrationProviderName field value
func (o *DisableProductTierIntegrationRequestBody) GetIntegrationProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationProviderName
}

// GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field value
// and a boolean to check if the value has been set.
func (o *DisableProductTierIntegrationRequestBody) GetIntegrationProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationProviderName, true
}

// SetIntegrationProviderName sets field value
func (o *DisableProductTierIntegrationRequestBody) SetIntegrationProviderName(v string) {
	o.IntegrationProviderName = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *DisableProductTierIntegrationRequestBody) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *DisableProductTierIntegrationRequestBody) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *DisableProductTierIntegrationRequestBody) SetIntegrationType(v string) {
	o.IntegrationType = v
}

func (o DisableProductTierIntegrationRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableProductTierIntegrationRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integrationProviderName"] = o.IntegrationProviderName
	toSerialize["integrationType"] = o.IntegrationType
	return toSerialize, nil
}

func (o *DisableProductTierIntegrationRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varDisableProductTierIntegrationRequestBody := _DisableProductTierIntegrationRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisableProductTierIntegrationRequestBody)

	if err != nil {
		return err
	}

	*o = DisableProductTierIntegrationRequestBody(varDisableProductTierIntegrationRequestBody)

	return err
}

type NullableDisableProductTierIntegrationRequestBody struct {
	value *DisableProductTierIntegrationRequestBody
	isSet bool
}

func (v NullableDisableProductTierIntegrationRequestBody) Get() *DisableProductTierIntegrationRequestBody {
	return v.value
}

func (v *NullableDisableProductTierIntegrationRequestBody) Set(val *DisableProductTierIntegrationRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableProductTierIntegrationRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableProductTierIntegrationRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableProductTierIntegrationRequestBody(val *DisableProductTierIntegrationRequestBody) *NullableDisableProductTierIntegrationRequestBody {
	return &NullableDisableProductTierIntegrationRequestBody{value: val, isSet: true}
}

func (v NullableDisableProductTierIntegrationRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableProductTierIntegrationRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


