/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribePlanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePlanResult{}

// DescribePlanResult struct for DescribePlanResult
type DescribePlanResult struct {
	// The time the plan was last modified
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// The next day stripe will process a charge for this plan
	NextChargeDate *string `json:"nextChargeDate,omitempty"`
	// Whether the customer has configured their payment information.
	PaymentConfigured *bool `json:"paymentConfigured,omitempty"`
	// The URL from Paigo to redirect users to so they can enter their payment information.  Only present when first adding payment information
	PaymentInfoPortalURL *string `json:"paymentInfoPortalURL,omitempty"`
	// The cost per core hour of this plan
	PlanCoreHourCost *float64 `json:"planCoreHourCost,omitempty"`
	// The description of the plan
	PlanDescription *string `json:"planDescription,omitempty"`
	// Whether the plan is charged monthly or yearly
	PlanFrequency *string `json:"planFrequency,omitempty"`
	// The minimum monthly cost of this plan
	PlanMonthlyCost *float64 `json:"planMonthlyCost,omitempty"`
	// The name of the plan this user is changing to
	PlanName string `json:"planName"`
	// The credits remaining for the customer for the customer in Paigo
	RemainingCredits *string `json:"remainingCredits,omitempty"`
	// The date that the plan starts
	StartDate *string `json:"startDate,omitempty"`
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

// HasModifiedAt returns a boolean if a field has been set.
func (o *DescribePlanResult) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *DescribePlanResult) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetNextChargeDate returns the NextChargeDate field value if set, zero value otherwise.
func (o *DescribePlanResult) GetNextChargeDate() string {
	if o == nil || IsNil(o.NextChargeDate) {
		var ret string
		return ret
	}
	return *o.NextChargeDate
}

// GetNextChargeDateOk returns a tuple with the NextChargeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetNextChargeDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextChargeDate) {
		return nil, false
	}
	return o.NextChargeDate, true
}

// HasNextChargeDate returns a boolean if a field has been set.
func (o *DescribePlanResult) HasNextChargeDate() bool {
	if o != nil && !IsNil(o.NextChargeDate) {
		return true
	}

	return false
}

// SetNextChargeDate gets a reference to the given string and assigns it to the NextChargeDate field.
func (o *DescribePlanResult) SetNextChargeDate(v string) {
	o.NextChargeDate = &v
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

// HasPaymentConfigured returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPaymentConfigured() bool {
	if o != nil && !IsNil(o.PaymentConfigured) {
		return true
	}

	return false
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

// HasPaymentInfoPortalURL returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPaymentInfoPortalURL() bool {
	if o != nil && !IsNil(o.PaymentInfoPortalURL) {
		return true
	}

	return false
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

// HasPlanCoreHourCost returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPlanCoreHourCost() bool {
	if o != nil && !IsNil(o.PlanCoreHourCost) {
		return true
	}

	return false
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

// HasPlanDescription returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPlanDescription() bool {
	if o != nil && !IsNil(o.PlanDescription) {
		return true
	}

	return false
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

// HasPlanFrequency returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPlanFrequency() bool {
	if o != nil && !IsNil(o.PlanFrequency) {
		return true
	}

	return false
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

// HasPlanMonthlyCost returns a boolean if a field has been set.
func (o *DescribePlanResult) HasPlanMonthlyCost() bool {
	if o != nil && !IsNil(o.PlanMonthlyCost) {
		return true
	}

	return false
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

// GetRemainingCredits returns the RemainingCredits field value if set, zero value otherwise.
func (o *DescribePlanResult) GetRemainingCredits() string {
	if o == nil || IsNil(o.RemainingCredits) {
		var ret string
		return ret
	}
	return *o.RemainingCredits
}

// GetRemainingCreditsOk returns a tuple with the RemainingCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePlanResult) GetRemainingCreditsOk() (*string, bool) {
	if o == nil || IsNil(o.RemainingCredits) {
		return nil, false
	}
	return o.RemainingCredits, true
}

// HasRemainingCredits returns a boolean if a field has been set.
func (o *DescribePlanResult) HasRemainingCredits() bool {
	if o != nil && !IsNil(o.RemainingCredits) {
		return true
	}

	return false
}

// SetRemainingCredits gets a reference to the given string and assigns it to the RemainingCredits field.
func (o *DescribePlanResult) SetRemainingCredits(v string) {
	o.RemainingCredits = &v
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

// HasStartDate returns a boolean if a field has been set.
func (o *DescribePlanResult) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
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
	if !IsNil(o.NextChargeDate) {
		toSerialize["nextChargeDate"] = o.NextChargeDate
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
	if !IsNil(o.RemainingCredits) {
		toSerialize["remainingCredits"] = o.RemainingCredits
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
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

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribePlanResult)

	if err != nil {
		return err
	}

	*o = DescribePlanResult(varDescribePlanResult)

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


