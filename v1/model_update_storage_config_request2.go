/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the UpdateStorageConfigRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageConfigRequest2{}

// UpdateStorageConfigRequest2 struct for UpdateStorageConfigRequest2
type UpdateStorageConfigRequest2 struct {
	// Description of the storage config
	Description *string `json:"description,omitempty"`
	// Name of the storage config
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorageConfigRequest2 UpdateStorageConfigRequest2

// NewUpdateStorageConfigRequest2 instantiates a new UpdateStorageConfigRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageConfigRequest2() *UpdateStorageConfigRequest2 {
	this := UpdateStorageConfigRequest2{}
	return &this
}

// NewUpdateStorageConfigRequest2WithDefaults instantiates a new UpdateStorageConfigRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageConfigRequest2WithDefaults() *UpdateStorageConfigRequest2 {
	this := UpdateStorageConfigRequest2{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateStorageConfigRequest2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageConfigRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateStorageConfigRequest2) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateStorageConfigRequest2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageConfigRequest2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateStorageConfigRequest2) SetName(v string) {
	o.Name = &v
}

func (o UpdateStorageConfigRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageConfigRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorageConfigRequest2) UnmarshalJSON(data []byte) (err error) {
	varUpdateStorageConfigRequest2 := _UpdateStorageConfigRequest2{}

	err = json.Unmarshal(data, &varUpdateStorageConfigRequest2)

	if err != nil {
		return err
	}

	*o = UpdateStorageConfigRequest2(varUpdateStorageConfigRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorageConfigRequest2 struct {
	value *UpdateStorageConfigRequest2
	isSet bool
}

func (v NullableUpdateStorageConfigRequest2) Get() *UpdateStorageConfigRequest2 {
	return v.value
}

func (v *NullableUpdateStorageConfigRequest2) Set(val *UpdateStorageConfigRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorageConfigRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorageConfigRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorageConfigRequest2(val *UpdateStorageConfigRequest2) *NullableUpdateStorageConfigRequest2 {
	return &NullableUpdateStorageConfigRequest2{value: val, isSet: true}
}

func (v NullableUpdateStorageConfigRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorageConfigRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


