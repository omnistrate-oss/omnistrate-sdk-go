/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeUserBillingDetailsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeUserBillingDetailsResult{}

// DescribeUserBillingDetailsResult struct for DescribeUserBillingDetailsResult
type DescribeUserBillingDetailsResult struct {
	// The URL from the billing provider to embed in an iframe to show the user's billing information
	BillingEmbedURL string `json:"billingEmbedURL"`
	// The name of the user
	Name string `json:"name"`
	// Whether the customer has configured their payment information.
	PaymentConfigured bool `json:"paymentConfigured"`
	// The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information
	PaymentInfoPortalURL *string `json:"paymentInfoPortalURL,omitempty"`
	// The User ID
	UserId string `json:"userId"`
}

type _DescribeUserBillingDetailsResult DescribeUserBillingDetailsResult

// NewDescribeUserBillingDetailsResult instantiates a new DescribeUserBillingDetailsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeUserBillingDetailsResult(billingEmbedURL string, name string, paymentConfigured bool, userId string) *DescribeUserBillingDetailsResult {
	this := DescribeUserBillingDetailsResult{}
	this.BillingEmbedURL = billingEmbedURL
	this.Name = name
	this.PaymentConfigured = paymentConfigured
	this.UserId = userId
	return &this
}

// NewDescribeUserBillingDetailsResultWithDefaults instantiates a new DescribeUserBillingDetailsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeUserBillingDetailsResultWithDefaults() *DescribeUserBillingDetailsResult {
	this := DescribeUserBillingDetailsResult{}
	return &this
}

// GetBillingEmbedURL returns the BillingEmbedURL field value
func (o *DescribeUserBillingDetailsResult) GetBillingEmbedURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingEmbedURL
}

// GetBillingEmbedURLOk returns a tuple with the BillingEmbedURL field value
// and a boolean to check if the value has been set.
func (o *DescribeUserBillingDetailsResult) GetBillingEmbedURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingEmbedURL, true
}

// SetBillingEmbedURL sets field value
func (o *DescribeUserBillingDetailsResult) SetBillingEmbedURL(v string) {
	o.BillingEmbedURL = v
}

// GetName returns the Name field value
func (o *DescribeUserBillingDetailsResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeUserBillingDetailsResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeUserBillingDetailsResult) SetName(v string) {
	o.Name = v
}

// GetPaymentConfigured returns the PaymentConfigured field value
func (o *DescribeUserBillingDetailsResult) GetPaymentConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PaymentConfigured
}

// GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field value
// and a boolean to check if the value has been set.
func (o *DescribeUserBillingDetailsResult) GetPaymentConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentConfigured, true
}

// SetPaymentConfigured sets field value
func (o *DescribeUserBillingDetailsResult) SetPaymentConfigured(v bool) {
	o.PaymentConfigured = v
}

// GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field value if set, zero value otherwise.
func (o *DescribeUserBillingDetailsResult) GetPaymentInfoPortalURL() string {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		var ret string
		return ret
	}
	return *o.PaymentInfoPortalURL
}

// GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeUserBillingDetailsResult) GetPaymentInfoPortalURLOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentInfoPortalURL) {
		return nil, false
	}
	return o.PaymentInfoPortalURL, true
}

// HasPaymentInfoPortalURL returns a boolean if a field has been set.
func (o *DescribeUserBillingDetailsResult) HasPaymentInfoPortalURL() bool {
	if o != nil && !IsNil(o.PaymentInfoPortalURL) {
		return true
	}

	return false
}

// SetPaymentInfoPortalURL gets a reference to the given string and assigns it to the PaymentInfoPortalURL field.
func (o *DescribeUserBillingDetailsResult) SetPaymentInfoPortalURL(v string) {
	o.PaymentInfoPortalURL = &v
}

// GetUserId returns the UserId field value
func (o *DescribeUserBillingDetailsResult) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DescribeUserBillingDetailsResult) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DescribeUserBillingDetailsResult) SetUserId(v string) {
	o.UserId = v
}

func (o DescribeUserBillingDetailsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeUserBillingDetailsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billingEmbedURL"] = o.BillingEmbedURL
	toSerialize["name"] = o.Name
	toSerialize["paymentConfigured"] = o.PaymentConfigured
	if !IsNil(o.PaymentInfoPortalURL) {
		toSerialize["paymentInfoPortalURL"] = o.PaymentInfoPortalURL
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *DescribeUserBillingDetailsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billingEmbedURL",
		"name",
		"paymentConfigured",
		"userId",
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

	varDescribeUserBillingDetailsResult := _DescribeUserBillingDetailsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeUserBillingDetailsResult)

	if err != nil {
		return err
	}

	*o = DescribeUserBillingDetailsResult(varDescribeUserBillingDetailsResult)

	return err
}

type NullableDescribeUserBillingDetailsResult struct {
	value *DescribeUserBillingDetailsResult
	isSet bool
}

func (v NullableDescribeUserBillingDetailsResult) Get() *DescribeUserBillingDetailsResult {
	return v.value
}

func (v *NullableDescribeUserBillingDetailsResult) Set(val *DescribeUserBillingDetailsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeUserBillingDetailsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeUserBillingDetailsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeUserBillingDetailsResult(val *DescribeUserBillingDetailsResult) *NullableDescribeUserBillingDetailsResult {
	return &NullableDescribeUserBillingDetailsResult{value: val, isSet: true}
}

func (v NullableDescribeUserBillingDetailsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeUserBillingDetailsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


