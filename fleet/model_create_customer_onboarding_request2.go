/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the CreateCustomerOnboardingRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerOnboardingRequest2{}

// CreateCustomerOnboardingRequest2 struct for CreateCustomerOnboardingRequest2
type CreateCustomerOnboardingRequest2 struct {
	// DEPRECATED: Name will be generated automatically.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateCustomerOnboardingRequest2 CreateCustomerOnboardingRequest2

// NewCreateCustomerOnboardingRequest2 instantiates a new CreateCustomerOnboardingRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerOnboardingRequest2() *CreateCustomerOnboardingRequest2 {
	this := CreateCustomerOnboardingRequest2{}
	return &this
}

// NewCreateCustomerOnboardingRequest2WithDefaults instantiates a new CreateCustomerOnboardingRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerOnboardingRequest2WithDefaults() *CreateCustomerOnboardingRequest2 {
	this := CreateCustomerOnboardingRequest2{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateCustomerOnboardingRequest2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerOnboardingRequest2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateCustomerOnboardingRequest2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCustomerOnboardingRequest2) SetName(v string) {
	o.Name = &v
}

func (o CreateCustomerOnboardingRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerOnboardingRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCustomerOnboardingRequest2) UnmarshalJSON(data []byte) (err error) {
	varCreateCustomerOnboardingRequest2 := _CreateCustomerOnboardingRequest2{}

	err = json.Unmarshal(data, &varCreateCustomerOnboardingRequest2)

	if err != nil {
		return err
	}

	*o = CreateCustomerOnboardingRequest2(varCreateCustomerOnboardingRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCustomerOnboardingRequest2 struct {
	value *CreateCustomerOnboardingRequest2
	isSet bool
}

func (v NullableCreateCustomerOnboardingRequest2) Get() *CreateCustomerOnboardingRequest2 {
	return v.value
}

func (v *NullableCreateCustomerOnboardingRequest2) Set(val *CreateCustomerOnboardingRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerOnboardingRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerOnboardingRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerOnboardingRequest2(val *CreateCustomerOnboardingRequest2) *NullableCreateCustomerOnboardingRequest2 {
	return &NullableCreateCustomerOnboardingRequest2{value: val, isSet: true}
}

func (v NullableCreateCustomerOnboardingRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerOnboardingRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


