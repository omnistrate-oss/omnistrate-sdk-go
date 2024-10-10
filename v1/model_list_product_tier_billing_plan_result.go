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

// checks if the ListProductTierBillingPlanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProductTierBillingPlanResult{}

// ListProductTierBillingPlanResult struct for ListProductTierBillingPlanResult
type ListProductTierBillingPlanResult struct {
	// List of product tier billing plans
	BillingPlans []BillingPlan `json:"billingPlans"`
	AdditionalProperties map[string]interface{}
}

type _ListProductTierBillingPlanResult ListProductTierBillingPlanResult

// NewListProductTierBillingPlanResult instantiates a new ListProductTierBillingPlanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductTierBillingPlanResult(billingPlans []BillingPlan) *ListProductTierBillingPlanResult {
	this := ListProductTierBillingPlanResult{}
	this.BillingPlans = billingPlans
	return &this
}

// NewListProductTierBillingPlanResultWithDefaults instantiates a new ListProductTierBillingPlanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductTierBillingPlanResultWithDefaults() *ListProductTierBillingPlanResult {
	this := ListProductTierBillingPlanResult{}
	return &this
}

// GetBillingPlans returns the BillingPlans field value
func (o *ListProductTierBillingPlanResult) GetBillingPlans() []BillingPlan {
	if o == nil {
		var ret []BillingPlan
		return ret
	}

	return o.BillingPlans
}

// GetBillingPlansOk returns a tuple with the BillingPlans field value
// and a boolean to check if the value has been set.
func (o *ListProductTierBillingPlanResult) GetBillingPlansOk() ([]BillingPlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingPlans, true
}

// SetBillingPlans sets field value
func (o *ListProductTierBillingPlanResult) SetBillingPlans(v []BillingPlan) {
	o.BillingPlans = v
}

func (o ListProductTierBillingPlanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProductTierBillingPlanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billingPlans"] = o.BillingPlans

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListProductTierBillingPlanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billingPlans",
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

	varListProductTierBillingPlanResult := _ListProductTierBillingPlanResult{}

	err = json.Unmarshal(data, &varListProductTierBillingPlanResult)

	if err != nil {
		return err
	}

	*o = ListProductTierBillingPlanResult(varListProductTierBillingPlanResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billingPlans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListProductTierBillingPlanResult struct {
	value *ListProductTierBillingPlanResult
	isSet bool
}

func (v NullableListProductTierBillingPlanResult) Get() *ListProductTierBillingPlanResult {
	return v.value
}

func (v *NullableListProductTierBillingPlanResult) Set(val *ListProductTierBillingPlanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductTierBillingPlanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductTierBillingPlanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductTierBillingPlanResult(val *ListProductTierBillingPlanResult) *NullableListProductTierBillingPlanResult {
	return &NullableListProductTierBillingPlanResult{value: val, isSet: true}
}

func (v NullableListProductTierBillingPlanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductTierBillingPlanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


