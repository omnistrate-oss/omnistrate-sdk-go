/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the DescribeCloudProviderCostResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeCloudProviderCostResult{}

// DescribeCloudProviderCostResult struct for DescribeCloudProviderCostResult
type DescribeCloudProviderCostResult struct {
	// The cost data for each cloud provider
	CloudProviderCosts map[string]interface{} `json:"cloudProviderCosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribeCloudProviderCostResult DescribeCloudProviderCostResult

// NewDescribeCloudProviderCostResult instantiates a new DescribeCloudProviderCostResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCloudProviderCostResult() *DescribeCloudProviderCostResult {
	this := DescribeCloudProviderCostResult{}
	return &this
}

// NewDescribeCloudProviderCostResultWithDefaults instantiates a new DescribeCloudProviderCostResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCloudProviderCostResultWithDefaults() *DescribeCloudProviderCostResult {
	this := DescribeCloudProviderCostResult{}
	return &this
}

// GetCloudProviderCosts returns the CloudProviderCosts field value if set, zero value otherwise.
func (o *DescribeCloudProviderCostResult) GetCloudProviderCosts() map[string]interface{} {
	if o == nil || IsNil(o.CloudProviderCosts) {
		var ret map[string]interface{}
		return ret
	}
	return o.CloudProviderCosts
}

// GetCloudProviderCostsOk returns a tuple with the CloudProviderCosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCloudProviderCostResult) GetCloudProviderCostsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CloudProviderCosts) {
		return map[string]interface{}{}, false
	}
	return o.CloudProviderCosts, true
}

// HasCloudProviderCosts returns a boolean if a field has been set.
func (o *DescribeCloudProviderCostResult) HasCloudProviderCosts() bool {
	if o != nil && !IsNil(o.CloudProviderCosts) {
		return true
	}

	return false
}

// SetCloudProviderCosts gets a reference to the given map[string]interface{} and assigns it to the CloudProviderCosts field.
func (o *DescribeCloudProviderCostResult) SetCloudProviderCosts(v map[string]interface{}) {
	o.CloudProviderCosts = v
}

func (o DescribeCloudProviderCostResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCloudProviderCostResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProviderCosts) {
		toSerialize["cloudProviderCosts"] = o.CloudProviderCosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeCloudProviderCostResult) UnmarshalJSON(data []byte) (err error) {
	varDescribeCloudProviderCostResult := _DescribeCloudProviderCostResult{}

	err = json.Unmarshal(data, &varDescribeCloudProviderCostResult)

	if err != nil {
		return err
	}

	*o = DescribeCloudProviderCostResult(varDescribeCloudProviderCostResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderCosts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeCloudProviderCostResult struct {
	value *DescribeCloudProviderCostResult
	isSet bool
}

func (v NullableDescribeCloudProviderCostResult) Get() *DescribeCloudProviderCostResult {
	return v.value
}

func (v *NullableDescribeCloudProviderCostResult) Set(val *DescribeCloudProviderCostResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCloudProviderCostResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCloudProviderCostResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCloudProviderCostResult(val *DescribeCloudProviderCostResult) *NullableDescribeCloudProviderCostResult {
	return &NullableDescribeCloudProviderCostResult{value: val, isSet: true}
}

func (v NullableDescribeCloudProviderCostResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCloudProviderCostResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


