/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
)

// checks if the WarmPoolConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarmPoolConfiguration{}

// WarmPoolConfiguration Warm pool configuration for compute nodes
type WarmPoolConfiguration struct {
	// Minimum number of compute nodes in pool
	MinimumNodesInPool *int64 `json:"minimumNodesInPool,omitempty"`
	// Minimum percentage of compute nodes in pool
	MinimumPercentageNodesInPool *int64 `json:"minimumPercentageNodesInPool,omitempty"`
}

// NewWarmPoolConfiguration instantiates a new WarmPoolConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarmPoolConfiguration() *WarmPoolConfiguration {
	this := WarmPoolConfiguration{}
	var minimumNodesInPool int64 = 1
	this.MinimumNodesInPool = &minimumNodesInPool
	var minimumPercentageNodesInPool int64 = 0
	this.MinimumPercentageNodesInPool = &minimumPercentageNodesInPool
	return &this
}

// NewWarmPoolConfigurationWithDefaults instantiates a new WarmPoolConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarmPoolConfigurationWithDefaults() *WarmPoolConfiguration {
	this := WarmPoolConfiguration{}
	var minimumNodesInPool int64 = 1
	this.MinimumNodesInPool = &minimumNodesInPool
	var minimumPercentageNodesInPool int64 = 0
	this.MinimumPercentageNodesInPool = &minimumPercentageNodesInPool
	return &this
}

// GetMinimumNodesInPool returns the MinimumNodesInPool field value if set, zero value otherwise.
func (o *WarmPoolConfiguration) GetMinimumNodesInPool() int64 {
	if o == nil || IsNil(o.MinimumNodesInPool) {
		var ret int64
		return ret
	}
	return *o.MinimumNodesInPool
}

// GetMinimumNodesInPoolOk returns a tuple with the MinimumNodesInPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarmPoolConfiguration) GetMinimumNodesInPoolOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumNodesInPool) {
		return nil, false
	}
	return o.MinimumNodesInPool, true
}

// HasMinimumNodesInPool returns a boolean if a field has been set.
func (o *WarmPoolConfiguration) HasMinimumNodesInPool() bool {
	if o != nil && !IsNil(o.MinimumNodesInPool) {
		return true
	}

	return false
}

// SetMinimumNodesInPool gets a reference to the given int64 and assigns it to the MinimumNodesInPool field.
func (o *WarmPoolConfiguration) SetMinimumNodesInPool(v int64) {
	o.MinimumNodesInPool = &v
}

// GetMinimumPercentageNodesInPool returns the MinimumPercentageNodesInPool field value if set, zero value otherwise.
func (o *WarmPoolConfiguration) GetMinimumPercentageNodesInPool() int64 {
	if o == nil || IsNil(o.MinimumPercentageNodesInPool) {
		var ret int64
		return ret
	}
	return *o.MinimumPercentageNodesInPool
}

// GetMinimumPercentageNodesInPoolOk returns a tuple with the MinimumPercentageNodesInPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarmPoolConfiguration) GetMinimumPercentageNodesInPoolOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumPercentageNodesInPool) {
		return nil, false
	}
	return o.MinimumPercentageNodesInPool, true
}

// HasMinimumPercentageNodesInPool returns a boolean if a field has been set.
func (o *WarmPoolConfiguration) HasMinimumPercentageNodesInPool() bool {
	if o != nil && !IsNil(o.MinimumPercentageNodesInPool) {
		return true
	}

	return false
}

// SetMinimumPercentageNodesInPool gets a reference to the given int64 and assigns it to the MinimumPercentageNodesInPool field.
func (o *WarmPoolConfiguration) SetMinimumPercentageNodesInPool(v int64) {
	o.MinimumPercentageNodesInPool = &v
}

func (o WarmPoolConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarmPoolConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinimumNodesInPool) {
		toSerialize["minimumNodesInPool"] = o.MinimumNodesInPool
	}
	if !IsNil(o.MinimumPercentageNodesInPool) {
		toSerialize["minimumPercentageNodesInPool"] = o.MinimumPercentageNodesInPool
	}
	return toSerialize, nil
}

type NullableWarmPoolConfiguration struct {
	value *WarmPoolConfiguration
	isSet bool
}

func (v NullableWarmPoolConfiguration) Get() *WarmPoolConfiguration {
	return v.value
}

func (v *NullableWarmPoolConfiguration) Set(val *WarmPoolConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableWarmPoolConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableWarmPoolConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarmPoolConfiguration(val *WarmPoolConfiguration) *NullableWarmPoolConfiguration {
	return &NullableWarmPoolConfiguration{value: val, isSet: true}
}

func (v NullableWarmPoolConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarmPoolConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


