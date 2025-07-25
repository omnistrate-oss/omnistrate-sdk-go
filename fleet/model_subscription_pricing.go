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

// checks if the SubscriptionPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPricing{}

// SubscriptionPricing struct for SubscriptionPricing
type SubscriptionPricing struct {
	// The time that this pricing was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// ID of a User
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The name of the user that created the price
	CreatedByUserName *string `json:"createdByUserName,omitempty"`
	// Whether this price is a custom price
	CustomPrice *bool `json:"customPrice,omitempty"`
	// The end date of the price
	EndDate *string `json:"endDate,omitempty"`
	// Whether this price should be used for billing
	IsActiveBillingPrice *bool `json:"isActiveBillingPrice,omitempty"`
	// The price per unit for the subscription
	PricePerUnit map[string]interface{} `json:"pricePerUnit,omitempty"`
	// The start date of the price
	StartDate *string `json:"startDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionPricing SubscriptionPricing

// NewSubscriptionPricing instantiates a new SubscriptionPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPricing() *SubscriptionPricing {
	this := SubscriptionPricing{}
	return &this
}

// NewSubscriptionPricingWithDefaults instantiates a new SubscriptionPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPricingWithDefaults() *SubscriptionPricing {
	this := SubscriptionPricing{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SubscriptionPricing) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *SubscriptionPricing) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetCreatedByUserName returns the CreatedByUserName field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetCreatedByUserName() string {
	if o == nil || IsNil(o.CreatedByUserName) {
		var ret string
		return ret
	}
	return *o.CreatedByUserName
}

// GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetCreatedByUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserName) {
		return nil, false
	}
	return o.CreatedByUserName, true
}

// HasCreatedByUserName returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasCreatedByUserName() bool {
	if o != nil && !IsNil(o.CreatedByUserName) {
		return true
	}

	return false
}

// SetCreatedByUserName gets a reference to the given string and assigns it to the CreatedByUserName field.
func (o *SubscriptionPricing) SetCreatedByUserName(v string) {
	o.CreatedByUserName = &v
}

// GetCustomPrice returns the CustomPrice field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetCustomPrice() bool {
	if o == nil || IsNil(o.CustomPrice) {
		var ret bool
		return ret
	}
	return *o.CustomPrice
}

// GetCustomPriceOk returns a tuple with the CustomPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetCustomPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.CustomPrice) {
		return nil, false
	}
	return o.CustomPrice, true
}

// HasCustomPrice returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasCustomPrice() bool {
	if o != nil && !IsNil(o.CustomPrice) {
		return true
	}

	return false
}

// SetCustomPrice gets a reference to the given bool and assigns it to the CustomPrice field.
func (o *SubscriptionPricing) SetCustomPrice(v bool) {
	o.CustomPrice = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SubscriptionPricing) SetEndDate(v string) {
	o.EndDate = &v
}

// GetIsActiveBillingPrice returns the IsActiveBillingPrice field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetIsActiveBillingPrice() bool {
	if o == nil || IsNil(o.IsActiveBillingPrice) {
		var ret bool
		return ret
	}
	return *o.IsActiveBillingPrice
}

// GetIsActiveBillingPriceOk returns a tuple with the IsActiveBillingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetIsActiveBillingPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActiveBillingPrice) {
		return nil, false
	}
	return o.IsActiveBillingPrice, true
}

// HasIsActiveBillingPrice returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasIsActiveBillingPrice() bool {
	if o != nil && !IsNil(o.IsActiveBillingPrice) {
		return true
	}

	return false
}

// SetIsActiveBillingPrice gets a reference to the given bool and assigns it to the IsActiveBillingPrice field.
func (o *SubscriptionPricing) SetIsActiveBillingPrice(v bool) {
	o.IsActiveBillingPrice = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetPricePerUnit() map[string]interface{} {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret map[string]interface{}
		return ret
	}
	return o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetPricePerUnitOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return map[string]interface{}{}, false
	}
	return o.PricePerUnit, true
}

// HasPricePerUnit returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasPricePerUnit() bool {
	if o != nil && !IsNil(o.PricePerUnit) {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given map[string]interface{} and assigns it to the PricePerUnit field.
func (o *SubscriptionPricing) SetPricePerUnit(v map[string]interface{}) {
	o.PricePerUnit = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SubscriptionPricing) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricing) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SubscriptionPricing) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SubscriptionPricing) SetStartDate(v string) {
	o.StartDate = &v
}

func (o SubscriptionPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if !IsNil(o.CreatedByUserName) {
		toSerialize["createdByUserName"] = o.CreatedByUserName
	}
	if !IsNil(o.CustomPrice) {
		toSerialize["customPrice"] = o.CustomPrice
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.IsActiveBillingPrice) {
		toSerialize["isActiveBillingPrice"] = o.IsActiveBillingPrice
	}
	if !IsNil(o.PricePerUnit) {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionPricing) UnmarshalJSON(data []byte) (err error) {
	varSubscriptionPricing := _SubscriptionPricing{}

	err = json.Unmarshal(data, &varSubscriptionPricing)

	if err != nil {
		return err
	}

	*o = SubscriptionPricing(varSubscriptionPricing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdByUserId")
		delete(additionalProperties, "createdByUserName")
		delete(additionalProperties, "customPrice")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "isActiveBillingPrice")
		delete(additionalProperties, "pricePerUnit")
		delete(additionalProperties, "startDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionPricing struct {
	value *SubscriptionPricing
	isSet bool
}

func (v NullableSubscriptionPricing) Get() *SubscriptionPricing {
	return v.value
}

func (v *NullableSubscriptionPricing) Set(val *SubscriptionPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPricing(val *SubscriptionPricing) *NullableSubscriptionPricing {
	return &NullableSubscriptionPricing{value: val, isSet: true}
}

func (v NullableSubscriptionPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


