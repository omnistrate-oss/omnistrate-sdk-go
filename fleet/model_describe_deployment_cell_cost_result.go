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

// checks if the DescribeDeploymentCellCostResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeDeploymentCellCostResult{}

// DescribeDeploymentCellCostResult struct for DescribeDeploymentCellCostResult
type DescribeDeploymentCellCostResult struct {
	// The cost data for each deployment cell
	DeploymentCellCosts map[string]interface{} `json:"deploymentCellCosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribeDeploymentCellCostResult DescribeDeploymentCellCostResult

// NewDescribeDeploymentCellCostResult instantiates a new DescribeDeploymentCellCostResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDeploymentCellCostResult() *DescribeDeploymentCellCostResult {
	this := DescribeDeploymentCellCostResult{}
	return &this
}

// NewDescribeDeploymentCellCostResultWithDefaults instantiates a new DescribeDeploymentCellCostResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDeploymentCellCostResultWithDefaults() *DescribeDeploymentCellCostResult {
	this := DescribeDeploymentCellCostResult{}
	return &this
}

// GetDeploymentCellCosts returns the DeploymentCellCosts field value if set, zero value otherwise.
func (o *DescribeDeploymentCellCostResult) GetDeploymentCellCosts() map[string]interface{} {
	if o == nil || IsNil(o.DeploymentCellCosts) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeploymentCellCosts
}

// GetDeploymentCellCostsOk returns a tuple with the DeploymentCellCosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDeploymentCellCostResult) GetDeploymentCellCostsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeploymentCellCosts) {
		return map[string]interface{}{}, false
	}
	return o.DeploymentCellCosts, true
}

// HasDeploymentCellCosts returns a boolean if a field has been set.
func (o *DescribeDeploymentCellCostResult) HasDeploymentCellCosts() bool {
	if o != nil && !IsNil(o.DeploymentCellCosts) {
		return true
	}

	return false
}

// SetDeploymentCellCosts gets a reference to the given map[string]interface{} and assigns it to the DeploymentCellCosts field.
func (o *DescribeDeploymentCellCostResult) SetDeploymentCellCosts(v map[string]interface{}) {
	o.DeploymentCellCosts = v
}

func (o DescribeDeploymentCellCostResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDeploymentCellCostResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentCellCosts) {
		toSerialize["deploymentCellCosts"] = o.DeploymentCellCosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeDeploymentCellCostResult) UnmarshalJSON(data []byte) (err error) {
	varDescribeDeploymentCellCostResult := _DescribeDeploymentCellCostResult{}

	err = json.Unmarshal(data, &varDescribeDeploymentCellCostResult)

	if err != nil {
		return err
	}

	*o = DescribeDeploymentCellCostResult(varDescribeDeploymentCellCostResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deploymentCellCosts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeDeploymentCellCostResult struct {
	value *DescribeDeploymentCellCostResult
	isSet bool
}

func (v NullableDescribeDeploymentCellCostResult) Get() *DescribeDeploymentCellCostResult {
	return v.value
}

func (v *NullableDescribeDeploymentCellCostResult) Set(val *DescribeDeploymentCellCostResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDeploymentCellCostResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDeploymentCellCostResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDeploymentCellCostResult(val *DescribeDeploymentCellCostResult) *NullableDescribeDeploymentCellCostResult {
	return &NullableDescribeDeploymentCellCostResult{value: val, isSet: true}
}

func (v NullableDescribeDeploymentCellCostResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDeploymentCellCostResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


