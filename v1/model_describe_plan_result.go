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

// checks if the DescribePlanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePlanResult{}

// DescribePlanResult struct for DescribePlanResult
type DescribePlanResult struct {
	// The time the plan was last modified
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// Whether the customer has configured their payment information.
	PaymentConfigured *bool `json:"paymentConfigured,omitempty"`
	// The URL from Stripe to configure payment information
	PaymentInfoPortalURL *string `json:"paymentInfoPortalURL,omitempty"`
	// The cost per core hour of this plan
	PlanCoreHourCost *float64 `json:"planCoreHourCost,omitempty"`
	// The description of the plan
	PlanDescription *string `json:"planDescription,omitempty"`
	// This parameter tells you if the plan is charged monthly or yearly
	PlanFrequency *string `json:"planFrequency,omitempty"`
	// The minimum monthly cost of this plan
	PlanMonthlyCost *float64 `json:"planMonthlyCost,omitempty"`
	// This parameter is used to select the appropriate Product Plan
	PlanName string `json:"planName"`
	// The date that the plan starts
	StartDate *string `json:"startDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribePlanResult DescribePlanResult

// NewDescribePlanResult instantiates a new DescribePlanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePlanResult(planName string) *DescribePlanResult {
	this := DescribePlanResult{}
	this.PlanName = planName
	return &this
}

// NewDescribePlanResultWithDefaults instantiates a new DescribePlanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePlanResultWithDefaults() *DescribePlanResult {
	this := DescribePlanResult{}
	return &this
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DescribePlanResult) GetModifiedAt() string {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetModifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *DescribePlanResult) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetPaymentConfigured returns the PaymentConfigured field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPaymentConfigured() bool {
	if o == nil || IsNil(o.PaymentConfigured) {
		var ret bool
		return ret
	}
	return *o.PaymentConfigured
}

// GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPaymentConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.PaymentConfigured) {
		return nil, false
	}
	return o.PaymentConfigured, true
}

// SetPaymentConfigured gets a reference to the given bool and assigns it to the PaymentConfigured field.
func (o *DescribePlanResult) SetPaymentConfigured(v bool) {
	o.PaymentConfigured = &v
}

// GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPaymentInfoPortalURL() string {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		var ret string
		return ret
	}
	return *o.PaymentInfoPortalURL
}

// GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPaymentInfoPortalURLOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		return nil, false
	}
	return o.PaymentInfoPortalURL, true
}

// SetPaymentInfoPortalURL gets a reference to the given string and assigns it to the PaymentInfoPortalURL field.
func (o *DescribePlanResult) SetPaymentInfoPortalURL(v string) {
	o.PaymentInfoPortalURL = &v
}

// GetPlanCoreHourCost returns the PlanCoreHourCost field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPlanCoreHourCost() float64 {
	if o == nil || IsNil(o.PlanCoreHourCost) {
		var ret float64
		return ret
	}
	return *o.PlanCoreHourCost
}

// GetPlanCoreHourCostOk returns a tuple with the PlanCoreHourCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPlanCoreHourCostOk() (*float64, bool) {
	if o == nil || IsNil(o.PlanCoreHourCost) {
		return nil, false
	}
	return o.PlanCoreHourCost, true
}

// SetPlanCoreHourCost gets a reference to the given float64 and assigns it to the PlanCoreHourCost field.
func (o *DescribePlanResult) SetPlanCoreHourCost(v float64) {
	o.PlanCoreHourCost = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPlanDescription() string {
	if o == nil || IsNil(o.PlanDescription) {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDescription) {
		return nil, false
	}
	return o.PlanDescription, true
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *DescribePlanResult) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetPlanFrequency returns the PlanFrequency field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPlanFrequency() string {
	if o == nil || IsNil(o.PlanFrequency) {
		var ret string
		return ret
	}
	return *o.PlanFrequency
}

// GetPlanFrequencyOk returns a tuple with the PlanFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPlanFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.PlanFrequency) {
		return nil, false
	}
	return o.PlanFrequency, true
}

// SetPlanFrequency gets a reference to the given string and assigns it to the PlanFrequency field.
func (o *DescribePlanResult) SetPlanFrequency(v string) {
	o.PlanFrequency = &v
}

// GetPlanMonthlyCost returns the PlanMonthlyCost field value if set, zero value otherwise.
func (o *DescribePlanResult) GetPlanMonthlyCost() float64 {
	if o == nil || IsNil(o.PlanMonthlyCost) {
		var ret float64
		return ret
	}
	return *o.PlanMonthlyCost
}

// GetPlanMonthlyCostOk returns a tuple with the PlanMonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPlanMonthlyCostOk() (*float64, bool) {
	if o == nil || IsNil(o.PlanMonthlyCost) {
		return nil, false
	}
	return o.PlanMonthlyCost, true
}

// SetPlanMonthlyCost gets a reference to the given float64 and assigns it to the PlanMonthlyCost field.
func (o *DescribePlanResult) SetPlanMonthlyCost(v float64) {
	o.PlanMonthlyCost = &v
}

// GetPlanName returns the PlanName field value
func (o *DescribePlanResult) GetPlanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanName, true
}

// SetPlanName sets field value
func (o *DescribePlanResult) SetPlanName(v string) {
	o.PlanName = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DescribePlanResult) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DescribePlanResult) SetStartDate(v string) {
	o.StartDate = &v
}

func (o DescribePlanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePlanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.PaymentConfigured) {
		toSerialize["paymentConfigured"] = o.PaymentConfigured
	}
	if !IsNil(o.PaymentInfoPortalURL) {
		toSerialize["paymentInfoPortalURL"] = o.PaymentInfoPortalURL
	}
	if !IsNil(o.PlanCoreHourCost) {
		toSerialize["planCoreHourCost"] = o.PlanCoreHourCost
	}
	if !IsNil(o.PlanDescription) {
		toSerialize["planDescription"] = o.PlanDescription
	}
	if !IsNil(o.PlanFrequency) {
		toSerialize["planFrequency"] = o.PlanFrequency
	}
	if !IsNil(o.PlanMonthlyCost) {
		toSerialize["planMonthlyCost"] = o.PlanMonthlyCost
	}
	toSerialize["planName"] = o.PlanName
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribePlanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planName",
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

	varDescribePlanResult := _DescribePlanResult{}

	err = json.Unmarshal(data, &varDescribePlanResult)

	if err != nil {
		return err
	}

	*o = DescribePlanResult(varDescribePlanResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "modifiedAt")
		delete(additionalProperties, "paymentConfigured")
		delete(additionalProperties, "paymentInfoPortalURL")
		delete(additionalProperties, "planCoreHourCost")
		delete(additionalProperties, "planDescription")
		delete(additionalProperties, "planFrequency")
		delete(additionalProperties, "planMonthlyCost")
		delete(additionalProperties, "planName")
		delete(additionalProperties, "startDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribePlanResult struct {
	value *DescribePlanResult
	isSet bool
}

func (v NullableDescribePlanResult) Get() *DescribePlanResult {
	return v.value
}

func (v *NullableDescribePlanResult) Set(val *DescribePlanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePlanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePlanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePlanResult(val *DescribePlanResult) *NullableDescribePlanResult {
	return &NullableDescribePlanResult{value: val, isSet: true}
}

func (v NullableDescribePlanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePlanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


