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

// checks if the GetConsumptionUsageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConsumptionUsageResult{}

// GetConsumptionUsageResult struct for GetConsumptionUsageResult
type GetConsumptionUsageResult struct {
	// End timestamp of usage
	EndTime *string `json:"endTime,omitempty"`
	// Start timestamp of usage
	StartTime *string `json:"startTime,omitempty"`
	// Usage for the current plan
	Usage []UsagePerDimension `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetConsumptionUsageResult GetConsumptionUsageResult

// NewGetConsumptionUsageResult instantiates a new GetConsumptionUsageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConsumptionUsageResult() *GetConsumptionUsageResult {
	this := GetConsumptionUsageResult{}
	return &this
}

// NewGetConsumptionUsageResultWithDefaults instantiates a new GetConsumptionUsageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConsumptionUsageResultWithDefaults() *GetConsumptionUsageResult {
	this := GetConsumptionUsageResult{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetConsumptionUsageResult) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsumptionUsageResult) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetConsumptionUsageResult) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *GetConsumptionUsageResult) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetConsumptionUsageResult) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsumptionUsageResult) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetConsumptionUsageResult) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GetConsumptionUsageResult) SetStartTime(v string) {
	o.StartTime = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetConsumptionUsageResult) GetUsage() []UsagePerDimension {
	if o == nil || IsNil(o.Usage) {
		var ret []UsagePerDimension
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsumptionUsageResult) GetUsageOk() ([]UsagePerDimension, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetConsumptionUsageResult) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsagePerDimension and assigns it to the Usage field.
func (o *GetConsumptionUsageResult) SetUsage(v []UsagePerDimension) {
	o.Usage = v
}

func (o GetConsumptionUsageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConsumptionUsageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
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

func (o *GetConsumptionUsageResult) UnmarshalJSON(data []byte) (err error) {
	varGetConsumptionUsageResult := _GetConsumptionUsageResult{}

	err = json.Unmarshal(data, &varGetConsumptionUsageResult)

	if err != nil {
		return err
	}

	*o = GetConsumptionUsageResult(varGetConsumptionUsageResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetConsumptionUsageResult struct {
	value *GetConsumptionUsageResult
	isSet bool
}

func (v NullableGetConsumptionUsageResult) Get() *GetConsumptionUsageResult {
	return v.value
}

func (v *NullableGetConsumptionUsageResult) Set(val *GetConsumptionUsageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConsumptionUsageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConsumptionUsageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConsumptionUsageResult(val *GetConsumptionUsageResult) *NullableGetConsumptionUsageResult {
	return &NullableGetConsumptionUsageResult{value: val, isSet: true}
}

func (v NullableGetConsumptionUsageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConsumptionUsageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


