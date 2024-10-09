/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the ListInvoicesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInvoicesResult{}

// ListInvoicesResult struct for ListInvoicesResult
type ListInvoicesResult struct {
	// List of Invoices
	Invoices []Invoice `json:"invoices,omitempty"`
}

// NewListInvoicesResult instantiates a new ListInvoicesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvoicesResult() *ListInvoicesResult {
	this := ListInvoicesResult{}
	return &this
}

// NewListInvoicesResultWithDefaults instantiates a new ListInvoicesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvoicesResultWithDefaults() *ListInvoicesResult {
	this := ListInvoicesResult{}
	return &this
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *ListInvoicesResult) GetInvoices() []Invoice {
	if o == nil || IsNil(o.Invoices) {
		var ret []Invoice
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInvoicesResult) GetInvoicesOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// SetInvoices gets a reference to the given []Invoice and assigns it to the Invoices field.
func (o *ListInvoicesResult) SetInvoices(v []Invoice) {
	o.Invoices = v
}

func (o ListInvoicesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInvoicesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	return toSerialize, nil
}

type NullableListInvoicesResult struct {
	value *ListInvoicesResult
	isSet bool
}

func (v NullableListInvoicesResult) Get() *ListInvoicesResult {
	return v.value
}

func (v *NullableListInvoicesResult) Set(val *ListInvoicesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListInvoicesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListInvoicesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInvoicesResult(val *ListInvoicesResult) *NullableListInvoicesResult {
	return &NullableListInvoicesResult{value: val, isSet: true}
}

func (v NullableListInvoicesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInvoicesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

