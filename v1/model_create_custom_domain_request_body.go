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

// checks if the CreateCustomDomainRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomDomainRequestBody{}

// CreateCustomDomainRequestBody struct for CreateCustomDomainRequestBody
type CreateCustomDomainRequestBody struct {
	// The root domain of the domain to use as suffix
	CustomDomain string `json:"customDomain"`
	// The description for the domain
	Description string `json:"description"`
	// The name of the custom domain
	Name string `json:"name"`
	Route53Configuration Route53Configuration `json:"route53Configuration"`
}

type _CreateCustomDomainRequestBody CreateCustomDomainRequestBody

// NewCreateCustomDomainRequestBody instantiates a new CreateCustomDomainRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomDomainRequestBody(customDomain string, description string, name string, route53Configuration Route53Configuration) *CreateCustomDomainRequestBody {
	this := CreateCustomDomainRequestBody{}
	this.CustomDomain = customDomain
	this.Description = description
	this.Name = name
	this.Route53Configuration = route53Configuration
	return &this
}

// NewCreateCustomDomainRequestBodyWithDefaults instantiates a new CreateCustomDomainRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomDomainRequestBodyWithDefaults() *CreateCustomDomainRequestBody {
	this := CreateCustomDomainRequestBody{}
	return &this
}

// GetCustomDomain returns the CustomDomain field value
func (o *CreateCustomDomainRequestBody) GetCustomDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDomain
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainRequestBody) GetCustomDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDomain, true
}

// SetCustomDomain sets field value
func (o *CreateCustomDomainRequestBody) SetCustomDomain(v string) {
	o.CustomDomain = v
}

// GetDescription returns the Description field value
func (o *CreateCustomDomainRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateCustomDomainRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateCustomDomainRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomDomainRequestBody) SetName(v string) {
	o.Name = v
}

// GetRoute53Configuration returns the Route53Configuration field value
func (o *CreateCustomDomainRequestBody) GetRoute53Configuration() Route53Configuration {
	if o == nil {
		var ret Route53Configuration
		return ret
	}

	return o.Route53Configuration
}

// GetRoute53ConfigurationOk returns a tuple with the Route53Configuration field value
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainRequestBody) GetRoute53ConfigurationOk() (*Route53Configuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Route53Configuration, true
}

// SetRoute53Configuration sets field value
func (o *CreateCustomDomainRequestBody) SetRoute53Configuration(v Route53Configuration) {
	o.Route53Configuration = v
}

func (o CreateCustomDomainRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomDomainRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customDomain"] = o.CustomDomain
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["route53Configuration"] = o.Route53Configuration
	return toSerialize, nil
}

func (o *CreateCustomDomainRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customDomain",
		"description",
		"name",
		"route53Configuration",
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

	varCreateCustomDomainRequestBody := _CreateCustomDomainRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCustomDomainRequestBody)

	if err != nil {
		return err
	}

	*o = CreateCustomDomainRequestBody(varCreateCustomDomainRequestBody)

	return err
}

type NullableCreateCustomDomainRequestBody struct {
	value *CreateCustomDomainRequestBody
	isSet bool
}

func (v NullableCreateCustomDomainRequestBody) Get() *CreateCustomDomainRequestBody {
	return v.value
}

func (v *NullableCreateCustomDomainRequestBody) Set(val *CreateCustomDomainRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomDomainRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomDomainRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomDomainRequestBody(val *CreateCustomDomainRequestBody) *NullableCreateCustomDomainRequestBody {
	return &NullableCreateCustomDomainRequestBody{value: val, isSet: true}
}

func (v NullableCreateCustomDomainRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomDomainRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

