/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
)

// checks if the InfraConfigChangeSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraConfigChangeSummary{}

// InfraConfigChangeSummary struct for InfraConfigChangeSummary
type InfraConfigChangeSummary struct {
	// The pending change state for the infra configuration
	ChangeState *string `json:"changeState,omitempty"`
	// The ID of the infrastructure configuration that this resource refers to
	InfraConfigId *string `json:"infraConfigId,omitempty"`
	// The name of the infra config
	Name *string `json:"name,omitempty"`
}

// NewInfraConfigChangeSummary instantiates a new InfraConfigChangeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraConfigChangeSummary() *InfraConfigChangeSummary {
	this := InfraConfigChangeSummary{}
	return &this
}

// NewInfraConfigChangeSummaryWithDefaults instantiates a new InfraConfigChangeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraConfigChangeSummaryWithDefaults() *InfraConfigChangeSummary {
	this := InfraConfigChangeSummary{}
	return &this
}

// GetChangeState returns the ChangeState field value if set, zero value otherwise.
func (o *InfraConfigChangeSummary) GetChangeState() string {
	if o == nil || IsNil(o.ChangeState) {
		var ret string
		return ret
	}
	return *o.ChangeState
}

// GetChangeStateOk returns a tuple with the ChangeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraConfigChangeSummary) GetChangeStateOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeState) {
		return nil, false
	}
	return o.ChangeState, true
}

// SetChangeState gets a reference to the given string and assigns it to the ChangeState field.
func (o *InfraConfigChangeSummary) SetChangeState(v string) {
	o.ChangeState = &v
}

// GetInfraConfigId returns the InfraConfigId field value if set, zero value otherwise.
func (o *InfraConfigChangeSummary) GetInfraConfigId() string {
	if o == nil || IsNil(o.InfraConfigId) {
		var ret string
		return ret
	}
	return *o.InfraConfigId
}

// GetInfraConfigIdOk returns a tuple with the InfraConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraConfigChangeSummary) GetInfraConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfraConfigId) {
		return nil, false
	}
	return o.InfraConfigId, true
}

// SetInfraConfigId gets a reference to the given string and assigns it to the InfraConfigId field.
func (o *InfraConfigChangeSummary) SetInfraConfigId(v string) {
	o.InfraConfigId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfraConfigChangeSummary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraConfigChangeSummary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfraConfigChangeSummary) SetName(v string) {
	o.Name = &v
}

func (o InfraConfigChangeSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraConfigChangeSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangeState) {
		toSerialize["changeState"] = o.ChangeState
	}
	if !IsNil(o.InfraConfigId) {
		toSerialize["infraConfigId"] = o.InfraConfigId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableInfraConfigChangeSummary struct {
	value *InfraConfigChangeSummary
	isSet bool
}

func (v NullableInfraConfigChangeSummary) Get() *InfraConfigChangeSummary {
	return v.value
}

func (v *NullableInfraConfigChangeSummary) Set(val *InfraConfigChangeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraConfigChangeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraConfigChangeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraConfigChangeSummary(val *InfraConfigChangeSummary) *NullableInfraConfigChangeSummary {
	return &NullableInfraConfigChangeSummary{value: val, isSet: true}
}

func (v NullableInfraConfigChangeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraConfigChangeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


