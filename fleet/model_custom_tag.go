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

// checks if the CustomTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTag{}

// CustomTag Custom tag
type CustomTag struct {
	// Key of the custom tag
	Key string `json:"key"`
	// Value of the custom tag
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CustomTag CustomTag

// NewCustomTag instantiates a new CustomTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTag(key string, value string) *CustomTag {
	this := CustomTag{}
	this.Key = key
	this.Value = value
	return &this
}

// NewCustomTagWithDefaults instantiates a new CustomTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTagWithDefaults() *CustomTag {
	this := CustomTag{}
	return &this
}

// GetKey returns the Key field value
func (o *CustomTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CustomTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CustomTag) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *CustomTag) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CustomTag) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CustomTag) SetValue(v string) {
	o.Value = v
}

func (o CustomTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varCustomTag := _CustomTag{}

	err = json.Unmarshal(data, &varCustomTag)

	if err != nil {
		return err
	}

	*o = CustomTag(varCustomTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTag struct {
	value *CustomTag
	isSet bool
}

func (v NullableCustomTag) Get() *CustomTag {
	return v.value
}

func (v *NullableCustomTag) Set(val *CustomTag) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTag) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTag(val *CustomTag) *NullableCustomTag {
	return &NullableCustomTag{value: val, isSet: true}
}

func (v NullableCustomTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


