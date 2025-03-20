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

// checks if the DescribeConsumptionUserBillingDetailsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeConsumptionUserBillingDetailsResult{}

// DescribeConsumptionUserBillingDetailsResult struct for DescribeConsumptionUserBillingDetailsResult
type DescribeConsumptionUserBillingDetailsResult struct {
	// DEPRECATED
	BillingEmbedURL *string `json:"billingEmbedURL,omitempty"`
	// DEPRECATED: The name of the user
	Name *string `json:"name,omitempty"`
	// Whether the customer has configured their payment information.
	PaymentConfigured bool `json:"paymentConfigured"`
	// The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information
	PaymentInfoPortalURL *string `json:"paymentInfoPortalURL,omitempty"`
	// ID of a User
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribeConsumptionUserBillingDetailsResult DescribeConsumptionUserBillingDetailsResult

// NewDescribeConsumptionUserBillingDetailsResult instantiates a new DescribeConsumptionUserBillingDetailsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeConsumptionUserBillingDetailsResult(paymentConfigured bool) *DescribeConsumptionUserBillingDetailsResult {
	this := DescribeConsumptionUserBillingDetailsResult{}
	this.PaymentConfigured = paymentConfigured
	return &this
}

// NewDescribeConsumptionUserBillingDetailsResultWithDefaults instantiates a new DescribeConsumptionUserBillingDetailsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeConsumptionUserBillingDetailsResultWithDefaults() *DescribeConsumptionUserBillingDetailsResult {
	this := DescribeConsumptionUserBillingDetailsResult{}
	return &this
}

// GetBillingEmbedURL returns the BillingEmbedURL field value if set, zero value otherwise.
func (o *DescribeConsumptionUserBillingDetailsResult) GetBillingEmbedURL() string {
	if o == nil || IsNil(o.BillingEmbedURL) {
		var ret string
		return ret
	}
	return *o.BillingEmbedURL
}

// GetBillingEmbedURLOk returns a tuple with the BillingEmbedURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConsumptionUserBillingDetailsResult) GetBillingEmbedURLOk() (*string, bool) {
	if o == nil || IsNil(o.BillingEmbedURL) {
		return nil, false
	}
	return o.BillingEmbedURL, true
}

// SetBillingEmbedURL gets a reference to the given string and assigns it to the BillingEmbedURL field.
func (o *DescribeConsumptionUserBillingDetailsResult) SetBillingEmbedURL(v string) {
	o.BillingEmbedURL = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DescribeConsumptionUserBillingDetailsResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConsumptionUserBillingDetailsResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DescribeConsumptionUserBillingDetailsResult) SetName(v string) {
	o.Name = &v
}

// GetPaymentConfigured returns the PaymentConfigured field value
func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PaymentConfigured
}

// GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field value
// and a boolean to check if the value has been set.
func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentConfigured, true
}

// SetPaymentConfigured sets field value
func (o *DescribeConsumptionUserBillingDetailsResult) SetPaymentConfigured(v bool) {
	o.PaymentConfigured = v
}

// GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field value if set, zero value otherwise.
func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentInfoPortalURL() string {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		var ret string
		return ret
	}
	return *o.PaymentInfoPortalURL
}

// GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentInfoPortalURLOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		return nil, false
	}
	return o.PaymentInfoPortalURL, true
}

// SetPaymentInfoPortalURL gets a reference to the given string and assigns it to the PaymentInfoPortalURL field.
func (o *DescribeConsumptionUserBillingDetailsResult) SetPaymentInfoPortalURL(v string) {
	o.PaymentInfoPortalURL = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DescribeConsumptionUserBillingDetailsResult) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConsumptionUserBillingDetailsResult) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DescribeConsumptionUserBillingDetailsResult) SetUserId(v string) {
	o.UserId = &v
}

func (o DescribeConsumptionUserBillingDetailsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeConsumptionUserBillingDetailsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingEmbedURL) {
		toSerialize["billingEmbedURL"] = o.BillingEmbedURL
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["paymentConfigured"] = o.PaymentConfigured
	if !IsNil(o.PaymentInfoPortalURL) {
		toSerialize["paymentInfoPortalURL"] = o.PaymentInfoPortalURL
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeConsumptionUserBillingDetailsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"paymentConfigured",
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

	varDescribeConsumptionUserBillingDetailsResult := _DescribeConsumptionUserBillingDetailsResult{}

	err = json.Unmarshal(data, &varDescribeConsumptionUserBillingDetailsResult)

	if err != nil {
		return err
	}

	*o = DescribeConsumptionUserBillingDetailsResult(varDescribeConsumptionUserBillingDetailsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billingEmbedURL")
		delete(additionalProperties, "name")
		delete(additionalProperties, "paymentConfigured")
		delete(additionalProperties, "paymentInfoPortalURL")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeConsumptionUserBillingDetailsResult struct {
	value *DescribeConsumptionUserBillingDetailsResult
	isSet bool
}

func (v NullableDescribeConsumptionUserBillingDetailsResult) Get() *DescribeConsumptionUserBillingDetailsResult {
	return v.value
}

func (v *NullableDescribeConsumptionUserBillingDetailsResult) Set(val *DescribeConsumptionUserBillingDetailsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeConsumptionUserBillingDetailsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeConsumptionUserBillingDetailsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeConsumptionUserBillingDetailsResult(val *DescribeConsumptionUserBillingDetailsResult) *NullableDescribeConsumptionUserBillingDetailsResult {
	return &NullableDescribeConsumptionUserBillingDetailsResult{value: val, isSet: true}
}

func (v NullableDescribeConsumptionUserBillingDetailsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeConsumptionUserBillingDetailsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


