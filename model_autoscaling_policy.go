/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AutoscalingPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoscalingPolicy{}

// AutoscalingPolicy Autoscaling policy for compute nodes
type AutoscalingPolicy struct {
	// Minutes before scaling down the compute nodes in idle state
	IdleMinutesBeforeScalingDown *int64 `json:"idleMinutesBeforeScalingDown,omitempty"`
	// Metric value below threshold will be considered to be idle
	IdleThreshold *int64 `json:"idleThreshold,omitempty"`
	// Maximum number of compute nodes to provision
	MaxReplicas string `json:"maxReplicas"`
	// Minimum number of compute nodes to provision
	MinReplicas string `json:"minReplicas"`
	// Minutes before scaling up the compute nodes in overUtilized state
	OverUtilizedMinutesBeforeScalingUp *int64 `json:"overUtilizedMinutesBeforeScalingUp,omitempty"`
	// Metric value beyond threshold will be considered to be overUtilized
	OverUtilizedThreshold *int64 `json:"overUtilizedThreshold,omitempty"`
	ScalingMetric *AutoScalingMetricSpec `json:"scalingMetric,omitempty"`
}

type _AutoscalingPolicy AutoscalingPolicy

// NewAutoscalingPolicy instantiates a new AutoscalingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoscalingPolicy(maxReplicas string, minReplicas string) *AutoscalingPolicy {
	this := AutoscalingPolicy{}
	var idleMinutesBeforeScalingDown int64 = 5
	this.IdleMinutesBeforeScalingDown = &idleMinutesBeforeScalingDown
	var idleThreshold int64 = 20
	this.IdleThreshold = &idleThreshold
	this.MaxReplicas = maxReplicas
	this.MinReplicas = minReplicas
	var overUtilizedMinutesBeforeScalingUp int64 = 5
	this.OverUtilizedMinutesBeforeScalingUp = &overUtilizedMinutesBeforeScalingUp
	var overUtilizedThreshold int64 = 80
	this.OverUtilizedThreshold = &overUtilizedThreshold
	return &this
}

// NewAutoscalingPolicyWithDefaults instantiates a new AutoscalingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoscalingPolicyWithDefaults() *AutoscalingPolicy {
	this := AutoscalingPolicy{}
	var idleMinutesBeforeScalingDown int64 = 5
	this.IdleMinutesBeforeScalingDown = &idleMinutesBeforeScalingDown
	var idleThreshold int64 = 20
	this.IdleThreshold = &idleThreshold
	var overUtilizedMinutesBeforeScalingUp int64 = 5
	this.OverUtilizedMinutesBeforeScalingUp = &overUtilizedMinutesBeforeScalingUp
	var overUtilizedThreshold int64 = 80
	this.OverUtilizedThreshold = &overUtilizedThreshold
	return &this
}

// GetIdleMinutesBeforeScalingDown returns the IdleMinutesBeforeScalingDown field value if set, zero value otherwise.
func (o *AutoscalingPolicy) GetIdleMinutesBeforeScalingDown() int64 {
	if o == nil || IsNil(o.IdleMinutesBeforeScalingDown) {
		var ret int64
		return ret
	}
	return *o.IdleMinutesBeforeScalingDown
}

// GetIdleMinutesBeforeScalingDownOk returns a tuple with the IdleMinutesBeforeScalingDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetIdleMinutesBeforeScalingDownOk() (*int64, bool) {
	if o == nil || IsNil(o.IdleMinutesBeforeScalingDown) {
		return nil, false
	}
	return o.IdleMinutesBeforeScalingDown, true
}

// HasIdleMinutesBeforeScalingDown returns a boolean if a field has been set.
func (o *AutoscalingPolicy) HasIdleMinutesBeforeScalingDown() bool {
	if o != nil && !IsNil(o.IdleMinutesBeforeScalingDown) {
		return true
	}

	return false
}

// SetIdleMinutesBeforeScalingDown gets a reference to the given int64 and assigns it to the IdleMinutesBeforeScalingDown field.
func (o *AutoscalingPolicy) SetIdleMinutesBeforeScalingDown(v int64) {
	o.IdleMinutesBeforeScalingDown = &v
}

// GetIdleThreshold returns the IdleThreshold field value if set, zero value otherwise.
func (o *AutoscalingPolicy) GetIdleThreshold() int64 {
	if o == nil || IsNil(o.IdleThreshold) {
		var ret int64
		return ret
	}
	return *o.IdleThreshold
}

// GetIdleThresholdOk returns a tuple with the IdleThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetIdleThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.IdleThreshold) {
		return nil, false
	}
	return o.IdleThreshold, true
}

// HasIdleThreshold returns a boolean if a field has been set.
func (o *AutoscalingPolicy) HasIdleThreshold() bool {
	if o != nil && !IsNil(o.IdleThreshold) {
		return true
	}

	return false
}

// SetIdleThreshold gets a reference to the given int64 and assigns it to the IdleThreshold field.
func (o *AutoscalingPolicy) SetIdleThreshold(v int64) {
	o.IdleThreshold = &v
}

// GetMaxReplicas returns the MaxReplicas field value
func (o *AutoscalingPolicy) GetMaxReplicas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxReplicas
}

// GetMaxReplicasOk returns a tuple with the MaxReplicas field value
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetMaxReplicasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxReplicas, true
}

// SetMaxReplicas sets field value
func (o *AutoscalingPolicy) SetMaxReplicas(v string) {
	o.MaxReplicas = v
}

// GetMinReplicas returns the MinReplicas field value
func (o *AutoscalingPolicy) GetMinReplicas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinReplicas
}

// GetMinReplicasOk returns a tuple with the MinReplicas field value
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetMinReplicasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinReplicas, true
}

// SetMinReplicas sets field value
func (o *AutoscalingPolicy) SetMinReplicas(v string) {
	o.MinReplicas = v
}

// GetOverUtilizedMinutesBeforeScalingUp returns the OverUtilizedMinutesBeforeScalingUp field value if set, zero value otherwise.
func (o *AutoscalingPolicy) GetOverUtilizedMinutesBeforeScalingUp() int64 {
	if o == nil || IsNil(o.OverUtilizedMinutesBeforeScalingUp) {
		var ret int64
		return ret
	}
	return *o.OverUtilizedMinutesBeforeScalingUp
}

// GetOverUtilizedMinutesBeforeScalingUpOk returns a tuple with the OverUtilizedMinutesBeforeScalingUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetOverUtilizedMinutesBeforeScalingUpOk() (*int64, bool) {
	if o == nil || IsNil(o.OverUtilizedMinutesBeforeScalingUp) {
		return nil, false
	}
	return o.OverUtilizedMinutesBeforeScalingUp, true
}

// HasOverUtilizedMinutesBeforeScalingUp returns a boolean if a field has been set.
func (o *AutoscalingPolicy) HasOverUtilizedMinutesBeforeScalingUp() bool {
	if o != nil && !IsNil(o.OverUtilizedMinutesBeforeScalingUp) {
		return true
	}

	return false
}

// SetOverUtilizedMinutesBeforeScalingUp gets a reference to the given int64 and assigns it to the OverUtilizedMinutesBeforeScalingUp field.
func (o *AutoscalingPolicy) SetOverUtilizedMinutesBeforeScalingUp(v int64) {
	o.OverUtilizedMinutesBeforeScalingUp = &v
}

// GetOverUtilizedThreshold returns the OverUtilizedThreshold field value if set, zero value otherwise.
func (o *AutoscalingPolicy) GetOverUtilizedThreshold() int64 {
	if o == nil || IsNil(o.OverUtilizedThreshold) {
		var ret int64
		return ret
	}
	return *o.OverUtilizedThreshold
}

// GetOverUtilizedThresholdOk returns a tuple with the OverUtilizedThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetOverUtilizedThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.OverUtilizedThreshold) {
		return nil, false
	}
	return o.OverUtilizedThreshold, true
}

// HasOverUtilizedThreshold returns a boolean if a field has been set.
func (o *AutoscalingPolicy) HasOverUtilizedThreshold() bool {
	if o != nil && !IsNil(o.OverUtilizedThreshold) {
		return true
	}

	return false
}

// SetOverUtilizedThreshold gets a reference to the given int64 and assigns it to the OverUtilizedThreshold field.
func (o *AutoscalingPolicy) SetOverUtilizedThreshold(v int64) {
	o.OverUtilizedThreshold = &v
}

// GetScalingMetric returns the ScalingMetric field value if set, zero value otherwise.
func (o *AutoscalingPolicy) GetScalingMetric() AutoScalingMetricSpec {
	if o == nil || IsNil(o.ScalingMetric) {
		var ret AutoScalingMetricSpec
		return ret
	}
	return *o.ScalingMetric
}

// GetScalingMetricOk returns a tuple with the ScalingMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoscalingPolicy) GetScalingMetricOk() (*AutoScalingMetricSpec, bool) {
	if o == nil || IsNil(o.ScalingMetric) {
		return nil, false
	}
	return o.ScalingMetric, true
}

// HasScalingMetric returns a boolean if a field has been set.
func (o *AutoscalingPolicy) HasScalingMetric() bool {
	if o != nil && !IsNil(o.ScalingMetric) {
		return true
	}

	return false
}

// SetScalingMetric gets a reference to the given AutoScalingMetricSpec and assigns it to the ScalingMetric field.
func (o *AutoscalingPolicy) SetScalingMetric(v AutoScalingMetricSpec) {
	o.ScalingMetric = &v
}

func (o AutoscalingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoscalingPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdleMinutesBeforeScalingDown) {
		toSerialize["idleMinutesBeforeScalingDown"] = o.IdleMinutesBeforeScalingDown
	}
	if !IsNil(o.IdleThreshold) {
		toSerialize["idleThreshold"] = o.IdleThreshold
	}
	toSerialize["maxReplicas"] = o.MaxReplicas
	toSerialize["minReplicas"] = o.MinReplicas
	if !IsNil(o.OverUtilizedMinutesBeforeScalingUp) {
		toSerialize["overUtilizedMinutesBeforeScalingUp"] = o.OverUtilizedMinutesBeforeScalingUp
	}
	if !IsNil(o.OverUtilizedThreshold) {
		toSerialize["overUtilizedThreshold"] = o.OverUtilizedThreshold
	}
	if !IsNil(o.ScalingMetric) {
		toSerialize["scalingMetric"] = o.ScalingMetric
	}
	return toSerialize, nil
}

func (o *AutoscalingPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maxReplicas",
		"minReplicas",
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

	varAutoscalingPolicy := _AutoscalingPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoscalingPolicy)

	if err != nil {
		return err
	}

	*o = AutoscalingPolicy(varAutoscalingPolicy)

	return err
}

type NullableAutoscalingPolicy struct {
	value *AutoscalingPolicy
	isSet bool
}

func (v NullableAutoscalingPolicy) Get() *AutoscalingPolicy {
	return v.value
}

func (v *NullableAutoscalingPolicy) Set(val *AutoscalingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoscalingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoscalingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoscalingPolicy(val *AutoscalingPolicy) *NullableAutoscalingPolicy {
	return &NullableAutoscalingPolicy{value: val, isSet: true}
}

func (v NullableAutoscalingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoscalingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


