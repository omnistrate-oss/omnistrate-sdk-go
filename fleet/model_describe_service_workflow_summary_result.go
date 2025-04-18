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

// checks if the DescribeServiceWorkflowSummaryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceWorkflowSummaryResult{}

// DescribeServiceWorkflowSummaryResult struct for DescribeServiceWorkflowSummaryResult
type DescribeServiceWorkflowSummaryResult struct {
	// Number of active workflows for the given service in the past 24 hours.
	ActiveWorkflowCount *int64 `json:"ActiveWorkflowCount,omitempty"`
	// Number of completed workflows for the given service in the past 24 hours.
	CompletedWorkflowCount *int64 `json:"CompletedWorkflowCount,omitempty"`
	// Number of failed workflows for the given service in the past 24 hours.
	FailedWorkflowCount *int64 `json:"FailedWorkflowCount,omitempty"`
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceWorkflowSummaryResult DescribeServiceWorkflowSummaryResult

// NewDescribeServiceWorkflowSummaryResult instantiates a new DescribeServiceWorkflowSummaryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceWorkflowSummaryResult(environmentId string, serviceId string) *DescribeServiceWorkflowSummaryResult {
	this := DescribeServiceWorkflowSummaryResult{}
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	return &this
}

// NewDescribeServiceWorkflowSummaryResultWithDefaults instantiates a new DescribeServiceWorkflowSummaryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceWorkflowSummaryResultWithDefaults() *DescribeServiceWorkflowSummaryResult {
	this := DescribeServiceWorkflowSummaryResult{}
	return &this
}

// GetActiveWorkflowCount returns the ActiveWorkflowCount field value if set, zero value otherwise.
func (o *DescribeServiceWorkflowSummaryResult) GetActiveWorkflowCount() int64 {
	if o == nil || IsNil(o.ActiveWorkflowCount) {
		var ret int64
		return ret
	}
	return *o.ActiveWorkflowCount
}

// GetActiveWorkflowCountOk returns a tuple with the ActiveWorkflowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowSummaryResult) GetActiveWorkflowCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveWorkflowCount) {
		return nil, false
	}
	return o.ActiveWorkflowCount, true
}

// HasActiveWorkflowCount returns a boolean if a field has been set.
func (o *DescribeServiceWorkflowSummaryResult) HasActiveWorkflowCount() bool {
	if o != nil && !IsNil(o.ActiveWorkflowCount) {
		return true
	}

	return false
}

// SetActiveWorkflowCount gets a reference to the given int64 and assigns it to the ActiveWorkflowCount field.
func (o *DescribeServiceWorkflowSummaryResult) SetActiveWorkflowCount(v int64) {
	o.ActiveWorkflowCount = &v
}

// GetCompletedWorkflowCount returns the CompletedWorkflowCount field value if set, zero value otherwise.
func (o *DescribeServiceWorkflowSummaryResult) GetCompletedWorkflowCount() int64 {
	if o == nil || IsNil(o.CompletedWorkflowCount) {
		var ret int64
		return ret
	}
	return *o.CompletedWorkflowCount
}

// GetCompletedWorkflowCountOk returns a tuple with the CompletedWorkflowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowSummaryResult) GetCompletedWorkflowCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletedWorkflowCount) {
		return nil, false
	}
	return o.CompletedWorkflowCount, true
}

// HasCompletedWorkflowCount returns a boolean if a field has been set.
func (o *DescribeServiceWorkflowSummaryResult) HasCompletedWorkflowCount() bool {
	if o != nil && !IsNil(o.CompletedWorkflowCount) {
		return true
	}

	return false
}

// SetCompletedWorkflowCount gets a reference to the given int64 and assigns it to the CompletedWorkflowCount field.
func (o *DescribeServiceWorkflowSummaryResult) SetCompletedWorkflowCount(v int64) {
	o.CompletedWorkflowCount = &v
}

// GetFailedWorkflowCount returns the FailedWorkflowCount field value if set, zero value otherwise.
func (o *DescribeServiceWorkflowSummaryResult) GetFailedWorkflowCount() int64 {
	if o == nil || IsNil(o.FailedWorkflowCount) {
		var ret int64
		return ret
	}
	return *o.FailedWorkflowCount
}

// GetFailedWorkflowCountOk returns a tuple with the FailedWorkflowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowSummaryResult) GetFailedWorkflowCountOk() (*int64, bool) {
	if o == nil || IsNil(o.FailedWorkflowCount) {
		return nil, false
	}
	return o.FailedWorkflowCount, true
}

// HasFailedWorkflowCount returns a boolean if a field has been set.
func (o *DescribeServiceWorkflowSummaryResult) HasFailedWorkflowCount() bool {
	if o != nil && !IsNil(o.FailedWorkflowCount) {
		return true
	}

	return false
}

// SetFailedWorkflowCount gets a reference to the given int64 and assigns it to the FailedWorkflowCount field.
func (o *DescribeServiceWorkflowSummaryResult) SetFailedWorkflowCount(v int64) {
	o.FailedWorkflowCount = &v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *DescribeServiceWorkflowSummaryResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowSummaryResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *DescribeServiceWorkflowSummaryResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceWorkflowSummaryResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowSummaryResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceWorkflowSummaryResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o DescribeServiceWorkflowSummaryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceWorkflowSummaryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveWorkflowCount) {
		toSerialize["ActiveWorkflowCount"] = o.ActiveWorkflowCount
	}
	if !IsNil(o.CompletedWorkflowCount) {
		toSerialize["CompletedWorkflowCount"] = o.CompletedWorkflowCount
	}
	if !IsNil(o.FailedWorkflowCount) {
		toSerialize["FailedWorkflowCount"] = o.FailedWorkflowCount
	}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["serviceId"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceWorkflowSummaryResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
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

	varDescribeServiceWorkflowSummaryResult := _DescribeServiceWorkflowSummaryResult{}

	err = json.Unmarshal(data, &varDescribeServiceWorkflowSummaryResult)

	if err != nil {
		return err
	}

	*o = DescribeServiceWorkflowSummaryResult(varDescribeServiceWorkflowSummaryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ActiveWorkflowCount")
		delete(additionalProperties, "CompletedWorkflowCount")
		delete(additionalProperties, "FailedWorkflowCount")
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "serviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceWorkflowSummaryResult struct {
	value *DescribeServiceWorkflowSummaryResult
	isSet bool
}

func (v NullableDescribeServiceWorkflowSummaryResult) Get() *DescribeServiceWorkflowSummaryResult {
	return v.value
}

func (v *NullableDescribeServiceWorkflowSummaryResult) Set(val *DescribeServiceWorkflowSummaryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceWorkflowSummaryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceWorkflowSummaryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceWorkflowSummaryResult(val *DescribeServiceWorkflowSummaryResult) *NullableDescribeServiceWorkflowSummaryResult {
	return &NullableDescribeServiceWorkflowSummaryResult{value: val, isSet: true}
}

func (v NullableDescribeServiceWorkflowSummaryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceWorkflowSummaryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


