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

// checks if the BuildServiceFromComposeSpecResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildServiceFromComposeSpecResult{}

// BuildServiceFromComposeSpecResult struct for BuildServiceFromComposeSpecResult
type BuildServiceFromComposeSpecResult struct {
	// ID of a Product Tier
	ProductTierID string `json:"productTierID"`
	// ID of a Service Environment
	ServiceEnvironmentID string `json:"serviceEnvironmentID"`
	// ID of a Service
	ServiceID string `json:"serviceID"`
	// Resources that appear in the service plan but were not defined in the spec. These resources were not processed during the build.
	UndefinedResources *map[string]string `json:"undefinedResources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildServiceFromComposeSpecResult BuildServiceFromComposeSpecResult

// NewBuildServiceFromComposeSpecResult instantiates a new BuildServiceFromComposeSpecResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildServiceFromComposeSpecResult(productTierID string, serviceEnvironmentID string, serviceID string) *BuildServiceFromComposeSpecResult {
	this := BuildServiceFromComposeSpecResult{}
	this.ProductTierID = productTierID
	this.ServiceEnvironmentID = serviceEnvironmentID
	this.ServiceID = serviceID
	return &this
}

// NewBuildServiceFromComposeSpecResultWithDefaults instantiates a new BuildServiceFromComposeSpecResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildServiceFromComposeSpecResultWithDefaults() *BuildServiceFromComposeSpecResult {
	this := BuildServiceFromComposeSpecResult{}
	return &this
}

// GetProductTierID returns the ProductTierID field value
func (o *BuildServiceFromComposeSpecResult) GetProductTierID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierID
}

// GetProductTierIDOk returns a tuple with the ProductTierID field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecResult) GetProductTierIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierID, true
}

// SetProductTierID sets field value
func (o *BuildServiceFromComposeSpecResult) SetProductTierID(v string) {
	o.ProductTierID = v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value
func (o *BuildServiceFromComposeSpecResult) GetServiceEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecResult) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID sets field value
func (o *BuildServiceFromComposeSpecResult) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = v
}

// GetServiceID returns the ServiceID field value
func (o *BuildServiceFromComposeSpecResult) GetServiceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecResult) GetServiceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceID, true
}

// SetServiceID sets field value
func (o *BuildServiceFromComposeSpecResult) SetServiceID(v string) {
	o.ServiceID = v
}

// GetUndefinedResources returns the UndefinedResources field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecResult) GetUndefinedResources() map[string]string {
	if o == nil || IsNil(o.UndefinedResources) {
		var ret map[string]string
		return ret
	}
	return *o.UndefinedResources
}

// GetUndefinedResourcesOk returns a tuple with the UndefinedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecResult) GetUndefinedResourcesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UndefinedResources) {
		return nil, false
	}
	return o.UndefinedResources, true
}

// SetUndefinedResources gets a reference to the given map[string]string and assigns it to the UndefinedResources field.
func (o *BuildServiceFromComposeSpecResult) SetUndefinedResources(v map[string]string) {
	o.UndefinedResources = &v
}

func (o BuildServiceFromComposeSpecResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildServiceFromComposeSpecResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productTierID"] = o.ProductTierID
	toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	toSerialize["serviceID"] = o.ServiceID
	if !IsNil(o.UndefinedResources) {
		toSerialize["undefinedResources"] = o.UndefinedResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BuildServiceFromComposeSpecResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierID",
		"serviceEnvironmentID",
		"serviceID",
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

	varBuildServiceFromComposeSpecResult := _BuildServiceFromComposeSpecResult{}

	err = json.Unmarshal(data, &varBuildServiceFromComposeSpecResult)

	if err != nil {
		return err
	}

	*o = BuildServiceFromComposeSpecResult(varBuildServiceFromComposeSpecResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productTierID")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "undefinedResources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildServiceFromComposeSpecResult struct {
	value *BuildServiceFromComposeSpecResult
	isSet bool
}

func (v NullableBuildServiceFromComposeSpecResult) Get() *BuildServiceFromComposeSpecResult {
	return v.value
}

func (v *NullableBuildServiceFromComposeSpecResult) Set(val *BuildServiceFromComposeSpecResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildServiceFromComposeSpecResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildServiceFromComposeSpecResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildServiceFromComposeSpecResult(val *BuildServiceFromComposeSpecResult) *NullableBuildServiceFromComposeSpecResult {
	return &NullableBuildServiceFromComposeSpecResult{value: val, isSet: true}
}

func (v NullableBuildServiceFromComposeSpecResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildServiceFromComposeSpecResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


