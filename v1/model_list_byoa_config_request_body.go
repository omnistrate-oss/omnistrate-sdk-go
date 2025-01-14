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

// checks if the ListBYOAConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBYOAConfigRequestBody{}

// ListBYOAConfigRequestBody struct for ListBYOAConfigRequestBody
type ListBYOAConfigRequestBody struct {
	// Cloud Provider name to filter on
	CloudProviderName string `json:"cloudProviderName"`
	AdditionalProperties map[string]interface{}
}

type _ListBYOAConfigRequestBody ListBYOAConfigRequestBody

// NewListBYOAConfigRequestBody instantiates a new ListBYOAConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBYOAConfigRequestBody(cloudProviderName string) *ListBYOAConfigRequestBody {
	this := ListBYOAConfigRequestBody{}
	this.CloudProviderName = cloudProviderName
	return &this
}

// NewListBYOAConfigRequestBodyWithDefaults instantiates a new ListBYOAConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBYOAConfigRequestBodyWithDefaults() *ListBYOAConfigRequestBody {
	this := ListBYOAConfigRequestBody{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *ListBYOAConfigRequestBody) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *ListBYOAConfigRequestBody) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *ListBYOAConfigRequestBody) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

func (o ListBYOAConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBYOAConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListBYOAConfigRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varListBYOAConfigRequestBody := _ListBYOAConfigRequestBody{}

	err = json.Unmarshal(data, &varListBYOAConfigRequestBody)

	if err != nil {
		return err
	}

	*o = ListBYOAConfigRequestBody(varListBYOAConfigRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListBYOAConfigRequestBody struct {
	value *ListBYOAConfigRequestBody
	isSet bool
}

func (v NullableListBYOAConfigRequestBody) Get() *ListBYOAConfigRequestBody {
	return v.value
}

func (v *NullableListBYOAConfigRequestBody) Set(val *ListBYOAConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListBYOAConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListBYOAConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBYOAConfigRequestBody(val *ListBYOAConfigRequestBody) *NullableListBYOAConfigRequestBody {
	return &NullableListBYOAConfigRequestBody{value: val, isSet: true}
}

func (v NullableListBYOAConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBYOAConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


