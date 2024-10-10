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

// checks if the CustomULimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomULimits{}

// CustomULimits struct for CustomULimits
type CustomULimits struct {
	// The name of the ulimit
	ULimitsName string `json:"ULimitsName"`
	// The type of the ulimit
	ULimitsType string `json:"ULimitsType"`
	// The value of the ulimit
	ULimitsValue int64 `json:"ULimitsValue"`
	AdditionalProperties map[string]interface{}
}

type _CustomULimits CustomULimits

// NewCustomULimits instantiates a new CustomULimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomULimits(uLimitsName string, uLimitsType string, uLimitsValue int64) *CustomULimits {
	this := CustomULimits{}
	this.ULimitsName = uLimitsName
	this.ULimitsType = uLimitsType
	this.ULimitsValue = uLimitsValue
	return &this
}

// NewCustomULimitsWithDefaults instantiates a new CustomULimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomULimitsWithDefaults() *CustomULimits {
	this := CustomULimits{}
	return &this
}

// GetULimitsName returns the ULimitsName field value
func (o *CustomULimits) GetULimitsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ULimitsName
}

// GetULimitsNameOk returns a tuple with the ULimitsName field value
// and a boolean to check if the value has been set.
func (o *CustomULimits) GetULimitsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ULimitsName, true
}

// SetULimitsName sets field value
func (o *CustomULimits) SetULimitsName(v string) {
	o.ULimitsName = v
}

// GetULimitsType returns the ULimitsType field value
func (o *CustomULimits) GetULimitsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ULimitsType
}

// GetULimitsTypeOk returns a tuple with the ULimitsType field value
// and a boolean to check if the value has been set.
func (o *CustomULimits) GetULimitsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ULimitsType, true
}

// SetULimitsType sets field value
func (o *CustomULimits) SetULimitsType(v string) {
	o.ULimitsType = v
}

// GetULimitsValue returns the ULimitsValue field value
func (o *CustomULimits) GetULimitsValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ULimitsValue
}

// GetULimitsValueOk returns a tuple with the ULimitsValue field value
// and a boolean to check if the value has been set.
func (o *CustomULimits) GetULimitsValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ULimitsValue, true
}

// SetULimitsValue sets field value
func (o *CustomULimits) SetULimitsValue(v int64) {
	o.ULimitsValue = v
}

func (o CustomULimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomULimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ULimitsName"] = o.ULimitsName
	toSerialize["ULimitsType"] = o.ULimitsType
	toSerialize["ULimitsValue"] = o.ULimitsValue

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomULimits) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ULimitsName",
		"ULimitsType",
		"ULimitsValue",
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

	varCustomULimits := _CustomULimits{}

	err = json.Unmarshal(data, &varCustomULimits)

	if err != nil {
		return err
	}

	*o = CustomULimits(varCustomULimits)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ULimitsName")
		delete(additionalProperties, "ULimitsType")
		delete(additionalProperties, "ULimitsValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomULimits struct {
	value *CustomULimits
	isSet bool
}

func (v NullableCustomULimits) Get() *CustomULimits {
	return v.value
}

func (v *NullableCustomULimits) Set(val *CustomULimits) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomULimits) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomULimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomULimits(val *CustomULimits) *NullableCustomULimits {
	return &NullableCustomULimits{value: val, isSet: true}
}

func (v NullableCustomULimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomULimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


