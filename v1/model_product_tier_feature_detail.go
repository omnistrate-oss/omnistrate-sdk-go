/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the ProductTierFeatureDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTierFeatureDetail{}

// ProductTierFeatureDetail struct for ProductTierFeatureDetail
type ProductTierFeatureDetail struct {
	// The configuration parameters of the product tier feature
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// Feature name
	Feature *string `json:"feature,omitempty"`
	// Feature scope
	Scope *string `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductTierFeatureDetail ProductTierFeatureDetail

// NewProductTierFeatureDetail instantiates a new ProductTierFeatureDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTierFeatureDetail() *ProductTierFeatureDetail {
	this := ProductTierFeatureDetail{}
	return &this
}

// NewProductTierFeatureDetailWithDefaults instantiates a new ProductTierFeatureDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTierFeatureDetailWithDefaults() *ProductTierFeatureDetail {
	this := ProductTierFeatureDetail{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ProductTierFeatureDetail) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTierFeatureDetail) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *ProductTierFeatureDetail) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ProductTierFeatureDetail) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTierFeatureDetail) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ProductTierFeatureDetail) SetFeature(v string) {
	o.Feature = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ProductTierFeatureDetail) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTierFeatureDetail) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ProductTierFeatureDetail) SetScope(v string) {
	o.Scope = &v
}

func (o ProductTierFeatureDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductTierFeatureDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductTierFeatureDetail) UnmarshalJSON(data []byte) (err error) {
	varProductTierFeatureDetail := _ProductTierFeatureDetail{}

	err = json.Unmarshal(data, &varProductTierFeatureDetail)

	if err != nil {
		return err
	}

	*o = ProductTierFeatureDetail(varProductTierFeatureDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "feature")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductTierFeatureDetail struct {
	value *ProductTierFeatureDetail
	isSet bool
}

func (v NullableProductTierFeatureDetail) Get() *ProductTierFeatureDetail {
	return v.value
}

func (v *NullableProductTierFeatureDetail) Set(val *ProductTierFeatureDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTierFeatureDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTierFeatureDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTierFeatureDetail(val *ProductTierFeatureDetail) *NullableProductTierFeatureDetail {
	return &NullableProductTierFeatureDetail{value: val, isSet: true}
}

func (v NullableProductTierFeatureDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTierFeatureDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


