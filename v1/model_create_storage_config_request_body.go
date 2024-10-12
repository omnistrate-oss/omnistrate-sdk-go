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

// checks if the CreateStorageConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStorageConfigRequestBody{}

// CreateStorageConfigRequestBody struct for CreateStorageConfigRequestBody
type CreateStorageConfigRequestBody struct {
	// Description of the storage config
	Description string `json:"description"`
	// Name of the storage config
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreateStorageConfigRequestBody CreateStorageConfigRequestBody

// NewCreateStorageConfigRequestBody instantiates a new CreateStorageConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStorageConfigRequestBody(description string, name string) *CreateStorageConfigRequestBody {
	this := CreateStorageConfigRequestBody{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateStorageConfigRequestBodyWithDefaults instantiates a new CreateStorageConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStorageConfigRequestBodyWithDefaults() *CreateStorageConfigRequestBody {
	this := CreateStorageConfigRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateStorageConfigRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateStorageConfigRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateStorageConfigRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateStorageConfigRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStorageConfigRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStorageConfigRequestBody) SetName(v string) {
	o.Name = v
}

func (o CreateStorageConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStorageConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateStorageConfigRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varCreateStorageConfigRequestBody := _CreateStorageConfigRequestBody{}

	err = json.Unmarshal(data, &varCreateStorageConfigRequestBody)

	if err != nil {
		return err
	}

	*o = CreateStorageConfigRequestBody(varCreateStorageConfigRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateStorageConfigRequestBody struct {
	value *CreateStorageConfigRequestBody
	isSet bool
}

func (v NullableCreateStorageConfigRequestBody) Get() *CreateStorageConfigRequestBody {
	return v.value
}

func (v *NullableCreateStorageConfigRequestBody) Set(val *CreateStorageConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStorageConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStorageConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStorageConfigRequestBody(val *CreateStorageConfigRequestBody) *NullableCreateStorageConfigRequestBody {
	return &NullableCreateStorageConfigRequestBody{value: val, isSet: true}
}

func (v NullableCreateStorageConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStorageConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


