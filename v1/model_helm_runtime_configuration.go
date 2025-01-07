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

// checks if the HelmRuntimeConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRuntimeConfiguration{}

// HelmRuntimeConfiguration struct for HelmRuntimeConfiguration
type HelmRuntimeConfiguration struct {
	// Disable Helm hooks
	DisableHooks bool `json:"disableHooks"`
	// Recreate the Helm package if it already exists
	Recreate bool `json:"recreate"`
	// Reset then reuse values for the Helm package before applying
	ResetThenReuseValues bool `json:"resetThenReuseValues"`
	// Reset values for the Helm package before applying
	ResetValues bool `json:"resetValues"`
	// Reuse values for the Helm package before applying
	ReuseValues bool `json:"reuseValues"`
	// Skip CRDs for the Helm package
	SkipCRDs bool `json:"skipCRDs"`
	// Timeout (nanos) for the Helm package to be deployed
	TimeoutNanos int64 `json:"timeoutNanos"`
	// Upgrade CRDs for the Helm package
	UpgradeCRDs bool `json:"upgradeCRDs"`
	// Wait for the Helm package to be deployed
	Wait bool `json:"wait"`
	// Wait for all jobs to be completed
	WaitForJobs bool `json:"waitForJobs"`
	AdditionalProperties map[string]interface{}
}

type _HelmRuntimeConfiguration HelmRuntimeConfiguration

// NewHelmRuntimeConfiguration instantiates a new HelmRuntimeConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRuntimeConfiguration(disableHooks bool, recreate bool, resetThenReuseValues bool, resetValues bool, reuseValues bool, skipCRDs bool, timeoutNanos int64, upgradeCRDs bool, wait bool, waitForJobs bool) *HelmRuntimeConfiguration {
	this := HelmRuntimeConfiguration{}
	this.DisableHooks = disableHooks
	this.Recreate = recreate
	this.ResetThenReuseValues = resetThenReuseValues
	this.ResetValues = resetValues
	this.ReuseValues = reuseValues
	this.SkipCRDs = skipCRDs
	this.TimeoutNanos = timeoutNanos
	this.UpgradeCRDs = upgradeCRDs
	this.Wait = wait
	this.WaitForJobs = waitForJobs
	return &this
}

// NewHelmRuntimeConfigurationWithDefaults instantiates a new HelmRuntimeConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRuntimeConfigurationWithDefaults() *HelmRuntimeConfiguration {
	this := HelmRuntimeConfiguration{}
	return &this
}

// GetDisableHooks returns the DisableHooks field value
func (o *HelmRuntimeConfiguration) GetDisableHooks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableHooks
}

// GetDisableHooksOk returns a tuple with the DisableHooks field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetDisableHooksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableHooks, true
}

// SetDisableHooks sets field value
func (o *HelmRuntimeConfiguration) SetDisableHooks(v bool) {
	o.DisableHooks = v
}

// GetRecreate returns the Recreate field value
func (o *HelmRuntimeConfiguration) GetRecreate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recreate
}

// GetRecreateOk returns a tuple with the Recreate field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetRecreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recreate, true
}

// SetRecreate sets field value
func (o *HelmRuntimeConfiguration) SetRecreate(v bool) {
	o.Recreate = v
}

// GetResetThenReuseValues returns the ResetThenReuseValues field value
func (o *HelmRuntimeConfiguration) GetResetThenReuseValues() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResetThenReuseValues
}

// GetResetThenReuseValuesOk returns a tuple with the ResetThenReuseValues field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetResetThenReuseValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetThenReuseValues, true
}

// SetResetThenReuseValues sets field value
func (o *HelmRuntimeConfiguration) SetResetThenReuseValues(v bool) {
	o.ResetThenReuseValues = v
}

// GetResetValues returns the ResetValues field value
func (o *HelmRuntimeConfiguration) GetResetValues() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResetValues
}

// GetResetValuesOk returns a tuple with the ResetValues field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetResetValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetValues, true
}

// SetResetValues sets field value
func (o *HelmRuntimeConfiguration) SetResetValues(v bool) {
	o.ResetValues = v
}

// GetReuseValues returns the ReuseValues field value
func (o *HelmRuntimeConfiguration) GetReuseValues() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReuseValues
}

// GetReuseValuesOk returns a tuple with the ReuseValues field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetReuseValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReuseValues, true
}

// SetReuseValues sets field value
func (o *HelmRuntimeConfiguration) SetReuseValues(v bool) {
	o.ReuseValues = v
}

// GetSkipCRDs returns the SkipCRDs field value
func (o *HelmRuntimeConfiguration) GetSkipCRDs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SkipCRDs
}

// GetSkipCRDsOk returns a tuple with the SkipCRDs field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetSkipCRDsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipCRDs, true
}

// SetSkipCRDs sets field value
func (o *HelmRuntimeConfiguration) SetSkipCRDs(v bool) {
	o.SkipCRDs = v
}

// GetTimeoutNanos returns the TimeoutNanos field value
func (o *HelmRuntimeConfiguration) GetTimeoutNanos() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimeoutNanos
}

// GetTimeoutNanosOk returns a tuple with the TimeoutNanos field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetTimeoutNanosOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutNanos, true
}

// SetTimeoutNanos sets field value
func (o *HelmRuntimeConfiguration) SetTimeoutNanos(v int64) {
	o.TimeoutNanos = v
}

// GetUpgradeCRDs returns the UpgradeCRDs field value
func (o *HelmRuntimeConfiguration) GetUpgradeCRDs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpgradeCRDs
}

// GetUpgradeCRDsOk returns a tuple with the UpgradeCRDs field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetUpgradeCRDsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradeCRDs, true
}

// SetUpgradeCRDs sets field value
func (o *HelmRuntimeConfiguration) SetUpgradeCRDs(v bool) {
	o.UpgradeCRDs = v
}

// GetWait returns the Wait field value
func (o *HelmRuntimeConfiguration) GetWait() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Wait
}

// GetWaitOk returns a tuple with the Wait field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetWaitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wait, true
}

// SetWait sets field value
func (o *HelmRuntimeConfiguration) SetWait(v bool) {
	o.Wait = v
}

// GetWaitForJobs returns the WaitForJobs field value
func (o *HelmRuntimeConfiguration) GetWaitForJobs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WaitForJobs
}

// GetWaitForJobsOk returns a tuple with the WaitForJobs field value
// and a boolean to check if the value has been set.
func (o *HelmRuntimeConfiguration) GetWaitForJobsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitForJobs, true
}

// SetWaitForJobs sets field value
func (o *HelmRuntimeConfiguration) SetWaitForJobs(v bool) {
	o.WaitForJobs = v
}

func (o HelmRuntimeConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRuntimeConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disableHooks"] = o.DisableHooks
	toSerialize["recreate"] = o.Recreate
	toSerialize["resetThenReuseValues"] = o.ResetThenReuseValues
	toSerialize["resetValues"] = o.ResetValues
	toSerialize["reuseValues"] = o.ReuseValues
	toSerialize["skipCRDs"] = o.SkipCRDs
	toSerialize["timeoutNanos"] = o.TimeoutNanos
	toSerialize["upgradeCRDs"] = o.UpgradeCRDs
	toSerialize["wait"] = o.Wait
	toSerialize["waitForJobs"] = o.WaitForJobs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRuntimeConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disableHooks",
		"recreate",
		"resetThenReuseValues",
		"resetValues",
		"reuseValues",
		"skipCRDs",
		"timeoutNanos",
		"upgradeCRDs",
		"wait",
		"waitForJobs",
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

	varHelmRuntimeConfiguration := _HelmRuntimeConfiguration{}

	err = json.Unmarshal(data, &varHelmRuntimeConfiguration)

	if err != nil {
		return err
	}

	*o = HelmRuntimeConfiguration(varHelmRuntimeConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disableHooks")
		delete(additionalProperties, "recreate")
		delete(additionalProperties, "resetThenReuseValues")
		delete(additionalProperties, "resetValues")
		delete(additionalProperties, "reuseValues")
		delete(additionalProperties, "skipCRDs")
		delete(additionalProperties, "timeoutNanos")
		delete(additionalProperties, "upgradeCRDs")
		delete(additionalProperties, "wait")
		delete(additionalProperties, "waitForJobs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRuntimeConfiguration struct {
	value *HelmRuntimeConfiguration
	isSet bool
}

func (v NullableHelmRuntimeConfiguration) Get() *HelmRuntimeConfiguration {
	return v.value
}

func (v *NullableHelmRuntimeConfiguration) Set(val *HelmRuntimeConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRuntimeConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRuntimeConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRuntimeConfiguration(val *HelmRuntimeConfiguration) *NullableHelmRuntimeConfiguration {
	return &NullableHelmRuntimeConfiguration{value: val, isSet: true}
}

func (v NullableHelmRuntimeConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRuntimeConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

