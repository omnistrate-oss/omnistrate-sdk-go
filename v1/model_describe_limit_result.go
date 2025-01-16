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

// checks if the DescribeLimitResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeLimitResult{}

// DescribeLimitResult struct for DescribeLimitResult
type DescribeLimitResult struct {
	// A brief description of the limit
	Description string `json:"description"`
	// The limit family
	Family string `json:"family"`
	// Unique key to identify the limit
	Key string `json:"key"`
	// Whether the limit can be modified
	Modifiable bool `json:"modifiable"`
	// Name of the limit
	Name string `json:"name"`
	// Value of the limit being enforced
	Value int64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _DescribeLimitResult DescribeLimitResult

// NewDescribeLimitResult instantiates a new DescribeLimitResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeLimitResult(description string, family string, key string, modifiable bool, name string, value int64) *DescribeLimitResult {
	this := DescribeLimitResult{}
	this.Description = description
	this.Family = family
	this.Key = key
	this.Modifiable = modifiable
	this.Name = name
	this.Value = value
	return &this
}

// NewDescribeLimitResultWithDefaults instantiates a new DescribeLimitResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeLimitResultWithDefaults() *DescribeLimitResult {
	this := DescribeLimitResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *DescribeLimitResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeLimitResult) SetDescription(v string) {
	o.Description = v
}

// GetFamily returns the Family field value
func (o *DescribeLimitResult) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *DescribeLimitResult) SetFamily(v string) {
	o.Family = v
}

// GetKey returns the Key field value
func (o *DescribeLimitResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeLimitResult) SetKey(v string) {
	o.Key = v
}

// GetModifiable returns the Modifiable field value
func (o *DescribeLimitResult) GetModifiable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Modifiable
}

// GetModifiableOk returns a tuple with the Modifiable field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetModifiableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modifiable, true
}

// SetModifiable sets field value
func (o *DescribeLimitResult) SetModifiable(v bool) {
	o.Modifiable = v
}

// GetName returns the Name field value
func (o *DescribeLimitResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeLimitResult) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *DescribeLimitResult) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DescribeLimitResult) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DescribeLimitResult) SetValue(v int64) {
	o.Value = v
}

func (o DescribeLimitResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeLimitResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["family"] = o.Family
	toSerialize["key"] = o.Key
	toSerialize["modifiable"] = o.Modifiable
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeLimitResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"family",
		"key",
		"modifiable",
		"name",
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

	varDescribeLimitResult := _DescribeLimitResult{}

	err = json.Unmarshal(data, &varDescribeLimitResult)

	if err != nil {
		return err
	}

	*o = DescribeLimitResult(varDescribeLimitResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "family")
		delete(additionalProperties, "key")
		delete(additionalProperties, "modifiable")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeLimitResult struct {
	value *DescribeLimitResult
	isSet bool
}

func (v NullableDescribeLimitResult) Get() *DescribeLimitResult {
	return v.value
}

func (v *NullableDescribeLimitResult) Set(val *DescribeLimitResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeLimitResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeLimitResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeLimitResult(val *DescribeLimitResult) *NullableDescribeLimitResult {
	return &NullableDescribeLimitResult{value: val, isSet: true}
}

func (v NullableDescribeLimitResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeLimitResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


