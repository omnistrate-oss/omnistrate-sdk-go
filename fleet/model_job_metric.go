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

// checks if the JobMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobMetric{}

// JobMetric struct for JobMetric
type JobMetric struct {
	// Additional metadata about the job
	AdditionalData map[string]interface{} `json:"additionalData,omitempty"`
	// When the job completed
	EndTime *string `json:"endTime,omitempty"`
	// Type of job metric being tracked
	MetricType string `json:"metricType"`
	// When the job started running
	StartTime string `json:"startTime"`
	// Value of the metric
	Value float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _JobMetric JobMetric

// NewJobMetric instantiates a new JobMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobMetric(metricType string, startTime string, value float64) *JobMetric {
	this := JobMetric{}
	this.MetricType = metricType
	this.StartTime = startTime
	this.Value = value
	return &this
}

// NewJobMetricWithDefaults instantiates a new JobMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobMetricWithDefaults() *JobMetric {
	this := JobMetric{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *JobMetric) GetAdditionalData() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalData) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobMetric) GetAdditionalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *JobMetric) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *JobMetric) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *JobMetric) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobMetric) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *JobMetric) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *JobMetric) SetEndTime(v string) {
	o.EndTime = &v
}

// GetMetricType returns the MetricType field value
func (o *JobMetric) GetMetricType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *JobMetric) GetMetricTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value
func (o *JobMetric) SetMetricType(v string) {
	o.MetricType = v
}

// GetStartTime returns the StartTime field value
func (o *JobMetric) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *JobMetric) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *JobMetric) SetStartTime(v string) {
	o.StartTime = v
}

// GetValue returns the Value field value
func (o *JobMetric) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JobMetric) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *JobMetric) SetValue(v float64) {
	o.Value = v
}

func (o JobMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	toSerialize["metricType"] = o.MetricType
	toSerialize["startTime"] = o.StartTime
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobMetric) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricType",
		"startTime",
		"value",
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

	varJobMetric := _JobMetric{}

	err = json.Unmarshal(data, &varJobMetric)

	if err != nil {
		return err
	}

	*o = JobMetric(varJobMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalData")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "metricType")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobMetric struct {
	value *JobMetric
	isSet bool
}

func (v NullableJobMetric) Get() *JobMetric {
	return v.value
}

func (v *NullableJobMetric) Set(val *JobMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableJobMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableJobMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobMetric(val *JobMetric) *NullableJobMetric {
	return &NullableJobMetric{value: val, isSet: true}
}

func (v NullableJobMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


