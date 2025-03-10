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

// checks if the CompleteOAuthConnectionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteOAuthConnectionResult{}

// CompleteOAuthConnectionResult struct for CompleteOAuthConnectionResult
type CompleteOAuthConnectionResult struct {
	// Stripe User ID
	StripeUserID         *string `json:"stripeUserID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompleteOAuthConnectionResult CompleteOAuthConnectionResult

// NewCompleteOAuthConnectionResult instantiates a new CompleteOAuthConnectionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteOAuthConnectionResult() *CompleteOAuthConnectionResult {
	this := CompleteOAuthConnectionResult{}
	return &this
}

// NewCompleteOAuthConnectionResultWithDefaults instantiates a new CompleteOAuthConnectionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteOAuthConnectionResultWithDefaults() *CompleteOAuthConnectionResult {
	this := CompleteOAuthConnectionResult{}
	return &this
}

// GetStripeUserID returns the StripeUserID field value if set, zero value otherwise.
func (o *CompleteOAuthConnectionResult) GetStripeUserID() string {
	if o == nil || IsNil(o.StripeUserID) {
		var ret string
		return ret
	}
	return *o.StripeUserID
}

// GetStripeUserIDOk returns a tuple with the StripeUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOAuthConnectionResult) GetStripeUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.StripeUserID) {
		return nil, false
	}
	return o.StripeUserID, true
}

// HasStripeUserID returns a boolean if a field has been set.
func (o *CompleteOAuthConnectionResult) HasStripeUserID() bool {
	if o != nil && !IsNil(o.StripeUserID) {
		return true
	}

	return false
}

// SetStripeUserID gets a reference to the given string and assigns it to the StripeUserID field.
func (o *CompleteOAuthConnectionResult) SetStripeUserID(v string) {
	o.StripeUserID = &v
}

func (o CompleteOAuthConnectionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteOAuthConnectionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StripeUserID) {
		toSerialize["stripeUserID"] = o.StripeUserID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompleteOAuthConnectionResult) UnmarshalJSON(data []byte) (err error) {
	varCompleteOAuthConnectionResult := _CompleteOAuthConnectionResult{}

	err = json.Unmarshal(data, &varCompleteOAuthConnectionResult)

	if err != nil {
		return err
	}

	*o = CompleteOAuthConnectionResult(varCompleteOAuthConnectionResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stripeUserID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteOAuthConnectionResult struct {
	value *CompleteOAuthConnectionResult
	isSet bool
}

func (v NullableCompleteOAuthConnectionResult) Get() *CompleteOAuthConnectionResult {
	return v.value
}

func (v *NullableCompleteOAuthConnectionResult) Set(val *CompleteOAuthConnectionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteOAuthConnectionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteOAuthConnectionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteOAuthConnectionResult(val *CompleteOAuthConnectionResult) *NullableCompleteOAuthConnectionResult {
	return &NullableCompleteOAuthConnectionResult{value: val, isSet: true}
}

func (v NullableCompleteOAuthConnectionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteOAuthConnectionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
