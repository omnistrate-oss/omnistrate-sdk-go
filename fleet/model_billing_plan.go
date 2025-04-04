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

// checks if the BillingPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPlan{}

// BillingPlan struct for BillingPlan
type BillingPlan struct {
	// Allow creates when payment not configured
	AllowCreatesWhenPaymentNotConfigured bool `json:"allowCreatesWhenPaymentNotConfigured"`
	// ID of a Product Tier Billing Plan
	Id string `json:"id"`
	// Maximum number of instances
	MaxNumberofInstances int64 `json:"maxNumberofInstances"`
	// Plan name
	PlanName string `json:"planName"`
	// Pricing in dollars.
	Pricing interface{} `json:"pricing"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	AdditionalProperties map[string]interface{}
}

type _BillingPlan BillingPlan

// NewBillingPlan instantiates a new BillingPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPlan(allowCreatesWhenPaymentNotConfigured bool, id string, maxNumberofInstances int64, planName string, pricing interface{}, productTierId string, serviceId string) *BillingPlan {
	this := BillingPlan{}
	this.AllowCreatesWhenPaymentNotConfigured = allowCreatesWhenPaymentNotConfigured
	this.Id = id
	this.MaxNumberofInstances = maxNumberofInstances
	this.PlanName = planName
	this.Pricing = pricing
	this.ProductTierId = productTierId
	this.ServiceId = serviceId
	return &this
}

// NewBillingPlanWithDefaults instantiates a new BillingPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPlanWithDefaults() *BillingPlan {
	this := BillingPlan{}
	return &this
}

// GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field value
func (o *BillingPlan) GetAllowCreatesWhenPaymentNotConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowCreatesWhenPaymentNotConfigured
}

// GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowCreatesWhenPaymentNotConfigured, true
}

// SetAllowCreatesWhenPaymentNotConfigured sets field value
func (o *BillingPlan) SetAllowCreatesWhenPaymentNotConfigured(v bool) {
	o.AllowCreatesWhenPaymentNotConfigured = v
}

// GetId returns the Id field value
func (o *BillingPlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillingPlan) SetId(v string) {
	o.Id = v
}

// GetMaxNumberofInstances returns the MaxNumberofInstances field value
func (o *BillingPlan) GetMaxNumberofInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxNumberofInstances
}

// GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetMaxNumberofInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumberofInstances, true
}

// SetMaxNumberofInstances sets field value
func (o *BillingPlan) SetMaxNumberofInstances(v int64) {
	o.MaxNumberofInstances = v
}

// GetPlanName returns the PlanName field value
func (o *BillingPlan) GetPlanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanName, true
}

// SetPlanName sets field value
func (o *BillingPlan) SetPlanName(v string) {
	o.PlanName = v
}

// GetPricing returns the Pricing field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BillingPlan) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingPlan) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing sets field value
func (o *BillingPlan) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetProductTierId returns the ProductTierId field value
func (o *BillingPlan) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *BillingPlan) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetServiceId returns the ServiceId field value
func (o *BillingPlan) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *BillingPlan) SetServiceId(v string) {
	o.ServiceId = v
}

func (o BillingPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowCreatesWhenPaymentNotConfigured"] = o.AllowCreatesWhenPaymentNotConfigured
	toSerialize["id"] = o.Id
	toSerialize["maxNumberofInstances"] = o.MaxNumberofInstances
	toSerialize["planName"] = o.PlanName
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["serviceId"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowCreatesWhenPaymentNotConfigured",
		"id",
		"maxNumberofInstances",
		"planName",
		"pricing",
		"productTierId",
		"serviceId",
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

	varBillingPlan := _BillingPlan{}

	err = json.Unmarshal(data, &varBillingPlan)

	if err != nil {
		return err
	}

	*o = BillingPlan(varBillingPlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowCreatesWhenPaymentNotConfigured")
		delete(additionalProperties, "id")
		delete(additionalProperties, "maxNumberofInstances")
		delete(additionalProperties, "planName")
		delete(additionalProperties, "pricing")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "serviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingPlan struct {
	value *BillingPlan
	isSet bool
}

func (v NullableBillingPlan) Get() *BillingPlan {
	return v.value
}

func (v *NullableBillingPlan) Set(val *BillingPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPlan(val *BillingPlan) *NullableBillingPlan {
	return &NullableBillingPlan{value: val, isSet: true}
}

func (v NullableBillingPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


