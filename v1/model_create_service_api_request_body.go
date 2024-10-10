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

// checks if the CreateServiceAPIRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceAPIRequestBody{}

// CreateServiceAPIRequestBody struct for CreateServiceAPIRequestBody
type CreateServiceAPIRequestBody struct {
	// A brief description of the service API bundle
	Description string `json:"description"`
	// The service environment ID
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	AdditionalProperties map[string]interface{}
}

type _CreateServiceAPIRequestBody CreateServiceAPIRequestBody

// NewCreateServiceAPIRequestBody instantiates a new CreateServiceAPIRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceAPIRequestBody(description string, serviceEnvironmentId string) *CreateServiceAPIRequestBody {
	this := CreateServiceAPIRequestBody{}
	this.Description = description
	this.ServiceEnvironmentId = serviceEnvironmentId
	return &this
}

// NewCreateServiceAPIRequestBodyWithDefaults instantiates a new CreateServiceAPIRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceAPIRequestBodyWithDefaults() *CreateServiceAPIRequestBody {
	this := CreateServiceAPIRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateServiceAPIRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAPIRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateServiceAPIRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *CreateServiceAPIRequestBody) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAPIRequestBody) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *CreateServiceAPIRequestBody) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

func (o CreateServiceAPIRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceAPIRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateServiceAPIRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"serviceEnvironmentId",
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

	varCreateServiceAPIRequestBody := _CreateServiceAPIRequestBody{}

	err = json.Unmarshal(data, &varCreateServiceAPIRequestBody)

	if err != nil {
		return err
	}

	*o = CreateServiceAPIRequestBody(varCreateServiceAPIRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "serviceEnvironmentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateServiceAPIRequestBody struct {
	value *CreateServiceAPIRequestBody
	isSet bool
}

func (v NullableCreateServiceAPIRequestBody) Get() *CreateServiceAPIRequestBody {
	return v.value
}

func (v *NullableCreateServiceAPIRequestBody) Set(val *CreateServiceAPIRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceAPIRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceAPIRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceAPIRequestBody(val *CreateServiceAPIRequestBody) *NullableCreateServiceAPIRequestBody {
	return &NullableCreateServiceAPIRequestBody{value: val, isSet: true}
}

func (v NullableCreateServiceAPIRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceAPIRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


