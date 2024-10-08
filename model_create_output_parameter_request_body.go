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

// checks if the CreateOutputParameterRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOutputParameterRequestBody{}

// CreateOutputParameterRequestBody struct for CreateOutputParameterRequestBody
type CreateOutputParameterRequestBody struct {
	// Description of the output variable being exported
	Description string `json:"description"`
	// Key of the output variable being exported
	Key string `json:"key" validate:"regexp=^[a-zA-Z][a-zA-Z0-9_]*$"`
	// External name of the output variable being exported
	Name string `json:"name"`
	// The ID of the resource that this output parameter belongs to
	ResourceId string `json:"resourceId"`
	// Value of the output variable being exported
	Value *string `json:"value,omitempty"`
	// Reference to another variable that will be used to set the value of the output variable being exported
	ValueRef *string `json:"valueRef,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
}

type _CreateOutputParameterRequestBody CreateOutputParameterRequestBody

// NewCreateOutputParameterRequestBody instantiates a new CreateOutputParameterRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutputParameterRequestBody(description string, key string, name string, resourceId string) *CreateOutputParameterRequestBody {
	this := CreateOutputParameterRequestBody{}
	this.Description = description
	this.Key = key
	this.Name = name
	this.ResourceId = resourceId
	return &this
}

// NewCreateOutputParameterRequestBodyWithDefaults instantiates a new CreateOutputParameterRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutputParameterRequestBodyWithDefaults() *CreateOutputParameterRequestBody {
	this := CreateOutputParameterRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateOutputParameterRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateOutputParameterRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetKey returns the Key field value
func (o *CreateOutputParameterRequestBody) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateOutputParameterRequestBody) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *CreateOutputParameterRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOutputParameterRequestBody) SetName(v string) {
	o.Name = v
}

// GetResourceId returns the ResourceId field value
func (o *CreateOutputParameterRequestBody) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateOutputParameterRequestBody) SetResourceId(v string) {
	o.ResourceId = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateOutputParameterRequestBody) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateOutputParameterRequestBody) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateOutputParameterRequestBody) SetValue(v string) {
	o.Value = &v
}

// GetValueRef returns the ValueRef field value if set, zero value otherwise.
func (o *CreateOutputParameterRequestBody) GetValueRef() string {
	if o == nil || IsNil(o.ValueRef) {
		var ret string
		return ret
	}
	return *o.ValueRef
}

// GetValueRefOk returns a tuple with the ValueRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetValueRefOk() (*string, bool) {
	if o == nil || IsNil(o.ValueRef) {
		return nil, false
	}
	return o.ValueRef, true
}

// HasValueRef returns a boolean if a field has been set.
func (o *CreateOutputParameterRequestBody) HasValueRef() bool {
	if o != nil && !IsNil(o.ValueRef) {
		return true
	}

	return false
}

// SetValueRef gets a reference to the given string and assigns it to the ValueRef field.
func (o *CreateOutputParameterRequestBody) SetValueRef(v string) {
	o.ValueRef = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *CreateOutputParameterRequestBody) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequestBody) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *CreateOutputParameterRequestBody) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *CreateOutputParameterRequestBody) SetValueType(v string) {
	o.ValueType = &v
}

func (o CreateOutputParameterRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutputParameterRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueRef) {
		toSerialize["valueRef"] = o.ValueRef
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	return toSerialize, nil
}

func (o *CreateOutputParameterRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"key",
		"name",
		"resourceId",
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

	varCreateOutputParameterRequestBody := _CreateOutputParameterRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOutputParameterRequestBody)

	if err != nil {
		return err
	}

	*o = CreateOutputParameterRequestBody(varCreateOutputParameterRequestBody)

	return err
}

type NullableCreateOutputParameterRequestBody struct {
	value *CreateOutputParameterRequestBody
	isSet bool
}

func (v NullableCreateOutputParameterRequestBody) Get() *CreateOutputParameterRequestBody {
	return v.value
}

func (v *NullableCreateOutputParameterRequestBody) Set(val *CreateOutputParameterRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutputParameterRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutputParameterRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutputParameterRequestBody(val *CreateOutputParameterRequestBody) *NullableCreateOutputParameterRequestBody {
	return &NullableCreateOutputParameterRequestBody{value: val, isSet: true}
}

func (v NullableCreateOutputParameterRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutputParameterRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


