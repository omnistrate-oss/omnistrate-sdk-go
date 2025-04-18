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

// checks if the ListBYOAConfigRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBYOAConfigRequest2{}

// ListBYOAConfigRequest2 struct for ListBYOAConfigRequest2
type ListBYOAConfigRequest2 struct {
	// Cloud Provider name to filter on
	CloudProviderName string `json:"cloudProviderName"`
	AdditionalProperties map[string]interface{}
}

type _ListBYOAConfigRequest2 ListBYOAConfigRequest2

// NewListBYOAConfigRequest2 instantiates a new ListBYOAConfigRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBYOAConfigRequest2(cloudProviderName string) *ListBYOAConfigRequest2 {
	this := ListBYOAConfigRequest2{}
	this.CloudProviderName = cloudProviderName
	return &this
}

// NewListBYOAConfigRequest2WithDefaults instantiates a new ListBYOAConfigRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBYOAConfigRequest2WithDefaults() *ListBYOAConfigRequest2 {
	this := ListBYOAConfigRequest2{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *ListBYOAConfigRequest2) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *ListBYOAConfigRequest2) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *ListBYOAConfigRequest2) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

func (o ListBYOAConfigRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBYOAConfigRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListBYOAConfigRequest2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
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

	varListBYOAConfigRequest2 := _ListBYOAConfigRequest2{}

	err = json.Unmarshal(data, &varListBYOAConfigRequest2)

	if err != nil {
		return err
	}

	*o = ListBYOAConfigRequest2(varListBYOAConfigRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListBYOAConfigRequest2 struct {
	value *ListBYOAConfigRequest2
	isSet bool
}

func (v NullableListBYOAConfigRequest2) Get() *ListBYOAConfigRequest2 {
	return v.value
}

func (v *NullableListBYOAConfigRequest2) Set(val *ListBYOAConfigRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableListBYOAConfigRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableListBYOAConfigRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBYOAConfigRequest2(val *ListBYOAConfigRequest2) *NullableListBYOAConfigRequest2 {
	return &NullableListBYOAConfigRequest2{value: val, isSet: true}
}

func (v NullableListBYOAConfigRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBYOAConfigRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


