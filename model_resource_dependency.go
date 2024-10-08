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

// checks if the ResourceDependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceDependency{}

// ResourceDependency Resource dependency relation specification
type ResourceDependency struct {
	// A map of the source parameter to the resource dependency parameter
	ParameterMap *map[string]string `json:"parameterMap,omitempty"`
	// The ID of the resource
	ResourceId string `json:"resourceId"`
}

type _ResourceDependency ResourceDependency

// NewResourceDependency instantiates a new ResourceDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDependency(resourceId string) *ResourceDependency {
	this := ResourceDependency{}
	this.ResourceId = resourceId
	return &this
}

// NewResourceDependencyWithDefaults instantiates a new ResourceDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDependencyWithDefaults() *ResourceDependency {
	this := ResourceDependency{}
	return &this
}

// GetParameterMap returns the ParameterMap field value if set, zero value otherwise.
func (o *ResourceDependency) GetParameterMap() map[string]string {
	if o == nil || IsNil(o.ParameterMap) {
		var ret map[string]string
		return ret
	}
	return *o.ParameterMap
}

// GetParameterMapOk returns a tuple with the ParameterMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceDependency) GetParameterMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ParameterMap) {
		return nil, false
	}
	return o.ParameterMap, true
}

// HasParameterMap returns a boolean if a field has been set.
func (o *ResourceDependency) HasParameterMap() bool {
	if o != nil && !IsNil(o.ParameterMap) {
		return true
	}

	return false
}

// SetParameterMap gets a reference to the given map[string]string and assigns it to the ParameterMap field.
func (o *ResourceDependency) SetParameterMap(v map[string]string) {
	o.ParameterMap = &v
}

// GetResourceId returns the ResourceId field value
func (o *ResourceDependency) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ResourceDependency) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ResourceDependency) SetResourceId(v string) {
	o.ResourceId = v
}

func (o ResourceDependency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceDependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterMap) {
		toSerialize["parameterMap"] = o.ParameterMap
	}
	toSerialize["resourceId"] = o.ResourceId
	return toSerialize, nil
}

func (o *ResourceDependency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varResourceDependency := _ResourceDependency{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceDependency)

	if err != nil {
		return err
	}

	*o = ResourceDependency(varResourceDependency)

	return err
}

type NullableResourceDependency struct {
	value *ResourceDependency
	isSet bool
}

func (v NullableResourceDependency) Get() *ResourceDependency {
	return v.value
}

func (v *NullableResourceDependency) Set(val *ResourceDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDependency(val *ResourceDependency) *NullableResourceDependency {
	return &NullableResourceDependency{value: val, isSet: true}
}

func (v NullableResourceDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


