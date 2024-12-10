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

// checks if the ProductTierIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTierIntegration{}

// ProductTierIntegration struct for ProductTierIntegration
type ProductTierIntegration struct {
	// The configuration parameters for the integration provider.
	Configuration map[string]interface{} `json:"configuration"`
	// The details of the integration failure
	IntegrationFailureDetails *string `json:"integrationFailureDetails,omitempty"`
	// The provider offering the integration for the product tier feature.
	IntegrationProviderName string `json:"integrationProviderName"`
	// The status of the integration
	IntegrationStatus string `json:"integrationStatus"`
	// ProductTierFeatureType is to enable / disable features per product tier
	IntegrationType string `json:"integrationType"`
	AdditionalProperties map[string]interface{}
}

type _ProductTierIntegration ProductTierIntegration

// NewProductTierIntegration instantiates a new ProductTierIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTierIntegration(configuration map[string]interface{}, integrationProviderName string, integrationStatus string, integrationType string) *ProductTierIntegration {
	this := ProductTierIntegration{}
	this.Configuration = configuration
	this.IntegrationProviderName = integrationProviderName
	this.IntegrationStatus = integrationStatus
	this.IntegrationType = integrationType
	return &this
}

// NewProductTierIntegrationWithDefaults instantiates a new ProductTierIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTierIntegrationWithDefaults() *ProductTierIntegration {
	this := ProductTierIntegration{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *ProductTierIntegration) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *ProductTierIntegration) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value
func (o *ProductTierIntegration) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetIntegrationFailureDetails returns the IntegrationFailureDetails field value if set, zero value otherwise.
func (o *ProductTierIntegration) GetIntegrationFailureDetails() string {
	if o == nil || IsNil(o.IntegrationFailureDetails) {
		var ret string
		return ret
	}
	return *o.IntegrationFailureDetails
}

// GetIntegrationFailureDetailsOk returns a tuple with the IntegrationFailureDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTierIntegration) GetIntegrationFailureDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationFailureDetails) {
		return nil, false
	}
	return o.IntegrationFailureDetails, true
}

// HasIntegrationFailureDetails returns a boolean if a field has been set.
func (o *ProductTierIntegration) HasIntegrationFailureDetails() bool {
	if o != nil && !IsNil(o.IntegrationFailureDetails) {
		return true
	}

	return false
}

// SetIntegrationFailureDetails gets a reference to the given string and assigns it to the IntegrationFailureDetails field.
func (o *ProductTierIntegration) SetIntegrationFailureDetails(v string) {
	o.IntegrationFailureDetails = &v
}

// GetIntegrationProviderName returns the IntegrationProviderName field value
func (o *ProductTierIntegration) GetIntegrationProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationProviderName
}

// GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field value
// and a boolean to check if the value has been set.
func (o *ProductTierIntegration) GetIntegrationProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationProviderName, true
}

// SetIntegrationProviderName sets field value
func (o *ProductTierIntegration) SetIntegrationProviderName(v string) {
	o.IntegrationProviderName = v
}

// GetIntegrationStatus returns the IntegrationStatus field value
func (o *ProductTierIntegration) GetIntegrationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationStatus
}

// GetIntegrationStatusOk returns a tuple with the IntegrationStatus field value
// and a boolean to check if the value has been set.
func (o *ProductTierIntegration) GetIntegrationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationStatus, true
}

// SetIntegrationStatus sets field value
func (o *ProductTierIntegration) SetIntegrationStatus(v string) {
	o.IntegrationStatus = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *ProductTierIntegration) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *ProductTierIntegration) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *ProductTierIntegration) SetIntegrationType(v string) {
	o.IntegrationType = v
}

func (o ProductTierIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductTierIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	if !IsNil(o.IntegrationFailureDetails) {
		toSerialize["integrationFailureDetails"] = o.IntegrationFailureDetails
	}
	toSerialize["integrationProviderName"] = o.IntegrationProviderName
	toSerialize["integrationStatus"] = o.IntegrationStatus
	toSerialize["integrationType"] = o.IntegrationType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductTierIntegration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configuration",
		"integrationProviderName",
		"integrationStatus",
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

	varProductTierIntegration := _ProductTierIntegration{}

	err = json.Unmarshal(data, &varProductTierIntegration)

	if err != nil {
		return err
	}

	*o = ProductTierIntegration(varProductTierIntegration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "integrationFailureDetails")
		delete(additionalProperties, "integrationProviderName")
		delete(additionalProperties, "integrationStatus")
		delete(additionalProperties, "integrationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductTierIntegration struct {
	value *ProductTierIntegration
	isSet bool
}

func (v NullableProductTierIntegration) Get() *ProductTierIntegration {
	return v.value
}

func (v *NullableProductTierIntegration) Set(val *ProductTierIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTierIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTierIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTierIntegration(val *ProductTierIntegration) *NullableProductTierIntegration {
	return &NullableProductTierIntegration{value: val, isSet: true}
}

func (v NullableProductTierIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTierIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


