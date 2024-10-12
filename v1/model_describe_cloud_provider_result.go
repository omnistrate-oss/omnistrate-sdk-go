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

// checks if the DescribeCloudProviderResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeCloudProviderResult{}

// DescribeCloudProviderResult struct for DescribeCloudProviderResult
type DescribeCloudProviderResult struct {
	// Description of the CloudProvider
	Description string `json:"description"`
	// ID of the CloudProvider
	Id string `json:"id"`
	// Name of the CloudProvider
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _DescribeCloudProviderResult DescribeCloudProviderResult

// NewDescribeCloudProviderResult instantiates a new DescribeCloudProviderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCloudProviderResult(description string, id string, name string) *DescribeCloudProviderResult {
	this := DescribeCloudProviderResult{}
	this.Description = description
	this.Id = id
	this.Name = name
	return &this
}

// NewDescribeCloudProviderResultWithDefaults instantiates a new DescribeCloudProviderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCloudProviderResultWithDefaults() *DescribeCloudProviderResult {
	this := DescribeCloudProviderResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *DescribeCloudProviderResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeCloudProviderResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeCloudProviderResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeCloudProviderResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeCloudProviderResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeCloudProviderResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DescribeCloudProviderResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeCloudProviderResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeCloudProviderResult) SetName(v string) {
	o.Name = v
}

func (o DescribeCloudProviderResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCloudProviderResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeCloudProviderResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
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

	varDescribeCloudProviderResult := _DescribeCloudProviderResult{}

	err = json.Unmarshal(data, &varDescribeCloudProviderResult)

	if err != nil {
		return err
	}

	*o = DescribeCloudProviderResult(varDescribeCloudProviderResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeCloudProviderResult struct {
	value *DescribeCloudProviderResult
	isSet bool
}

func (v NullableDescribeCloudProviderResult) Get() *DescribeCloudProviderResult {
	return v.value
}

func (v *NullableDescribeCloudProviderResult) Set(val *DescribeCloudProviderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCloudProviderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCloudProviderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCloudProviderResult(val *DescribeCloudProviderResult) *NullableDescribeCloudProviderResult {
	return &NullableDescribeCloudProviderResult{value: val, isSet: true}
}

func (v NullableDescribeCloudProviderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCloudProviderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


