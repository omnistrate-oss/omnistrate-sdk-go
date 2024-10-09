/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateServiceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceRequestBody{}

// CreateServiceRequestBody struct for CreateServiceRequestBody
type CreateServiceRequestBody struct {
	// A brief description of the service
	Description string `json:"description"`
	// Name of the Service
	Name string `json:"name"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
}

type _CreateServiceRequestBody CreateServiceRequestBody

// NewCreateServiceRequestBody instantiates a new CreateServiceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceRequestBody(description string, name string) *CreateServiceRequestBody {
	this := CreateServiceRequestBody{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateServiceRequestBodyWithDefaults instantiates a new CreateServiceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceRequestBodyWithDefaults() *CreateServiceRequestBody {
	this := CreateServiceRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateServiceRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateServiceRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateServiceRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateServiceRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServiceRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServiceRequestBody) SetName(v string) {
	o.Name = v
}

// GetServiceLogoURL returns the ServiceLogoURL field value if set, zero value otherwise.
func (o *CreateServiceRequestBody) GetServiceLogoURL() string {
	if o == nil || IsNil(o.ServiceLogoURL) {
		var ret string
		return ret
	}
	return *o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceRequestBody) GetServiceLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogoURL) {
		return nil, false
	}
	return o.ServiceLogoURL, true
}

// SetServiceLogoURL gets a reference to the given string and assigns it to the ServiceLogoURL field.
func (o *CreateServiceRequestBody) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = &v
}

func (o CreateServiceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	return toSerialize, nil
}

func (o *CreateServiceRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateServiceRequestBody := _CreateServiceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateServiceRequestBody)

	if err != nil {
		return err
	}

	*o = CreateServiceRequestBody(varCreateServiceRequestBody)

	return err
}

type NullableCreateServiceRequestBody struct {
	value *CreateServiceRequestBody
	isSet bool
}

func (v NullableCreateServiceRequestBody) Get() *CreateServiceRequestBody {
	return v.value
}

func (v *NullableCreateServiceRequestBody) Set(val *CreateServiceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceRequestBody(val *CreateServiceRequestBody) *NullableCreateServiceRequestBody {
	return &NullableCreateServiceRequestBody{value: val, isSet: true}
}

func (v NullableCreateServiceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

