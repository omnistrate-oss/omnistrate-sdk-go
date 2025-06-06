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

// checks if the RecentDeploymentFailureStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentDeploymentFailureStatus{}

// RecentDeploymentFailureStatus The status of the recent deployment failure
type RecentDeploymentFailureStatus struct {
	// The time at which the deployment failed
	FailedAt *string `json:"failedAt,omitempty"`
	// The reason for the deployment failure
	FailureReason *string `json:"failureReason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentDeploymentFailureStatus RecentDeploymentFailureStatus

// NewRecentDeploymentFailureStatus instantiates a new RecentDeploymentFailureStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentDeploymentFailureStatus() *RecentDeploymentFailureStatus {
	this := RecentDeploymentFailureStatus{}
	return &this
}

// NewRecentDeploymentFailureStatusWithDefaults instantiates a new RecentDeploymentFailureStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentDeploymentFailureStatusWithDefaults() *RecentDeploymentFailureStatus {
	this := RecentDeploymentFailureStatus{}
	return &this
}

// GetFailedAt returns the FailedAt field value if set, zero value otherwise.
func (o *RecentDeploymentFailureStatus) GetFailedAt() string {
	if o == nil || IsNil(o.FailedAt) {
		var ret string
		return ret
	}
	return *o.FailedAt
}

// GetFailedAtOk returns a tuple with the FailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentDeploymentFailureStatus) GetFailedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FailedAt) {
		return nil, false
	}
	return o.FailedAt, true
}

// HasFailedAt returns a boolean if a field has been set.
func (o *RecentDeploymentFailureStatus) HasFailedAt() bool {
	if o != nil && !IsNil(o.FailedAt) {
		return true
	}

	return false
}

// SetFailedAt gets a reference to the given string and assigns it to the FailedAt field.
func (o *RecentDeploymentFailureStatus) SetFailedAt(v string) {
	o.FailedAt = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *RecentDeploymentFailureStatus) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentDeploymentFailureStatus) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *RecentDeploymentFailureStatus) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *RecentDeploymentFailureStatus) SetFailureReason(v string) {
	o.FailureReason = &v
}

func (o RecentDeploymentFailureStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentDeploymentFailureStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailedAt) {
		toSerialize["failedAt"] = o.FailedAt
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentDeploymentFailureStatus) UnmarshalJSON(data []byte) (err error) {
	varRecentDeploymentFailureStatus := _RecentDeploymentFailureStatus{}

	err = json.Unmarshal(data, &varRecentDeploymentFailureStatus)

	if err != nil {
		return err
	}

	*o = RecentDeploymentFailureStatus(varRecentDeploymentFailureStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "failedAt")
		delete(additionalProperties, "failureReason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentDeploymentFailureStatus struct {
	value *RecentDeploymentFailureStatus
	isSet bool
}

func (v NullableRecentDeploymentFailureStatus) Get() *RecentDeploymentFailureStatus {
	return v.value
}

func (v *NullableRecentDeploymentFailureStatus) Set(val *RecentDeploymentFailureStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentDeploymentFailureStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentDeploymentFailureStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentDeploymentFailureStatus(val *RecentDeploymentFailureStatus) *NullableRecentDeploymentFailureStatus {
	return &NullableRecentDeploymentFailureStatus{value: val, isSet: true}
}

func (v NullableRecentDeploymentFailureStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentDeploymentFailureStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


