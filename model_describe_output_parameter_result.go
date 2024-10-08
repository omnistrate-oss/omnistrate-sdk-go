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

// checks if the DescribeOutputParameterResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeOutputParameterResult{}

// DescribeOutputParameterResult struct for DescribeOutputParameterResult
type DescribeOutputParameterResult struct {
	// Description of the output variable being exported
	Description string `json:"description"`
	// The ID of the output parameter
	Id string `json:"id"`
	// Key of the output variable being exported
	Key string `json:"key"`
	// External name of the output variable being exported
	Name string `json:"name"`
	// The ID of the resource that this input parameter belongs to
	ResourceId string `json:"resourceId"`
	// The ID of the service that this output parameter belongs to
	ServiceId string `json:"serviceId"`
	// Value of the output variable being exported
	Value *string `json:"value,omitempty"`
	// Reference to an input variable that will be used to set the value of the output variable being exported
	ValueRef *string `json:"valueRef,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
}

type _DescribeOutputParameterResult DescribeOutputParameterResult

// NewDescribeOutputParameterResult instantiates a new DescribeOutputParameterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeOutputParameterResult(description string, id string, key string, name string, resourceId string, serviceId string) *DescribeOutputParameterResult {
	this := DescribeOutputParameterResult{}
	this.Description = description
	this.Id = id
	this.Key = key
	this.Name = name
	this.ResourceId = resourceId
	this.ServiceId = serviceId
	return &this
}

// NewDescribeOutputParameterResultWithDefaults instantiates a new DescribeOutputParameterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeOutputParameterResultWithDefaults() *DescribeOutputParameterResult {
	this := DescribeOutputParameterResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *DescribeOutputParameterResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeOutputParameterResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeOutputParameterResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeOutputParameterResult) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *DescribeOutputParameterResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeOutputParameterResult) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *DescribeOutputParameterResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeOutputParameterResult) SetName(v string) {
	o.Name = v
}

// GetResourceId returns the ResourceId field value
func (o *DescribeOutputParameterResult) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *DescribeOutputParameterResult) SetResourceId(v string) {
	o.ResourceId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeOutputParameterResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeOutputParameterResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DescribeOutputParameterResult) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DescribeOutputParameterResult) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DescribeOutputParameterResult) SetValue(v string) {
	o.Value = &v
}

// GetValueRef returns the ValueRef field value if set, zero value otherwise.
func (o *DescribeOutputParameterResult) GetValueRef() string {
	if o == nil || IsNil(o.ValueRef) {
		var ret string
		return ret
	}
	return *o.ValueRef
}

// GetValueRefOk returns a tuple with the ValueRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetValueRefOk() (*string, bool) {
	if o == nil || IsNil(o.ValueRef) {
		return nil, false
	}
	return o.ValueRef, true
}

// HasValueRef returns a boolean if a field has been set.
func (o *DescribeOutputParameterResult) HasValueRef() bool {
	if o != nil && !IsNil(o.ValueRef) {
		return true
	}

	return false
}

// SetValueRef gets a reference to the given string and assigns it to the ValueRef field.
func (o *DescribeOutputParameterResult) SetValueRef(v string) {
	o.ValueRef = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *DescribeOutputParameterResult) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeOutputParameterResult) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *DescribeOutputParameterResult) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *DescribeOutputParameterResult) SetValueType(v string) {
	o.ValueType = &v
}

func (o DescribeOutputParameterResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeOutputParameterResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["serviceId"] = o.ServiceId
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

func (o *DescribeOutputParameterResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"key",
		"name",
		"resourceId",
		"serviceId",
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

	varDescribeOutputParameterResult := _DescribeOutputParameterResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeOutputParameterResult)

	if err != nil {
		return err
	}

	*o = DescribeOutputParameterResult(varDescribeOutputParameterResult)

	return err
}

type NullableDescribeOutputParameterResult struct {
	value *DescribeOutputParameterResult
	isSet bool
}

func (v NullableDescribeOutputParameterResult) Get() *DescribeOutputParameterResult {
	return v.value
}

func (v *NullableDescribeOutputParameterResult) Set(val *DescribeOutputParameterResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeOutputParameterResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeOutputParameterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeOutputParameterResult(val *DescribeOutputParameterResult) *NullableDescribeOutputParameterResult {
	return &NullableDescribeOutputParameterResult{value: val, isSet: true}
}

func (v NullableDescribeOutputParameterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeOutputParameterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


