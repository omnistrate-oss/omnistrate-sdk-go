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

// checks if the OnboardingStage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingStage{}

// OnboardingStage struct for OnboardingStage
type OnboardingStage struct {
	// The name of the stage.
	Name string `json:"name"`
	// The status of the stage.
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnboardingStage OnboardingStage

// NewOnboardingStage instantiates a new OnboardingStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingStage(name string) *OnboardingStage {
	this := OnboardingStage{}
	this.Name = name
	return &this
}

// NewOnboardingStageWithDefaults instantiates a new OnboardingStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingStageWithDefaults() *OnboardingStage {
	this := OnboardingStage{}
	return &this
}

// GetName returns the Name field value
func (o *OnboardingStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OnboardingStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OnboardingStage) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OnboardingStage) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingStage) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OnboardingStage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OnboardingStage) SetStatus(v string) {
	o.Status = &v
}

func (o OnboardingStage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnboardingStage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varOnboardingStage := _OnboardingStage{}

	err = json.Unmarshal(data, &varOnboardingStage)

	if err != nil {
		return err
	}

	*o = OnboardingStage(varOnboardingStage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnboardingStage struct {
	value *OnboardingStage
	isSet bool
}

func (v NullableOnboardingStage) Get() *OnboardingStage {
	return v.value
}

func (v *NullableOnboardingStage) Set(val *OnboardingStage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingStage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingStage(val *OnboardingStage) *NullableOnboardingStage {
	return &NullableOnboardingStage{value: val, isSet: true}
}

func (v NullableOnboardingStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


