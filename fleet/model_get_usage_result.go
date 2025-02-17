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

// checks if the GetUsageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsageResult{}

// GetUsageResult struct for GetUsageResult
type GetUsageResult struct {
	// End timestamp of usage
	EndTime *string `json:"endTime,omitempty"`
	// This parameter is used to select the appropriate Product Plan
	PlanName *string `json:"planName,omitempty"`
	// Start timestamp of usage
	StartTime *string `json:"startTime,omitempty"`
	// Usage for the current plan
	Usage []UsagePerDimension `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUsageResult GetUsageResult

// NewGetUsageResult instantiates a new GetUsageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsageResult() *GetUsageResult {
	this := GetUsageResult{}
	return &this
}

// NewGetUsageResultWithDefaults instantiates a new GetUsageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsageResultWithDefaults() *GetUsageResult {
	this := GetUsageResult{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetUsageResult) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageResult) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetUsageResult) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *GetUsageResult) SetEndTime(v string) {
	o.EndTime = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *GetUsageResult) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageResult) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *GetUsageResult) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *GetUsageResult) SetPlanName(v string) {
	o.PlanName = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetUsageResult) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageResult) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetUsageResult) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GetUsageResult) SetStartTime(v string) {
	o.StartTime = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetUsageResult) GetUsage() []UsagePerDimension {
	if o == nil || IsNil(o.Usage) {
		var ret []UsagePerDimension
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageResult) GetUsageOk() ([]UsagePerDimension, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetUsageResult) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsagePerDimension and assigns it to the Usage field.
func (o *GetUsageResult) SetUsage(v []UsagePerDimension) {
	o.Usage = v
}

func (o GetUsageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUsageResult) UnmarshalJSON(data []byte) (err error) {
	varGetUsageResult := _GetUsageResult{}

	err = json.Unmarshal(data, &varGetUsageResult)

	if err != nil {
		return err
	}

	*o = GetUsageResult(varGetUsageResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "planName")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUsageResult struct {
	value *GetUsageResult
	isSet bool
}

func (v NullableGetUsageResult) Get() *GetUsageResult {
	return v.value
}

func (v *NullableGetUsageResult) Set(val *GetUsageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsageResult(val *GetUsageResult) *NullableGetUsageResult {
	return &NullableGetUsageResult{value: val, isSet: true}
}

func (v NullableGetUsageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


