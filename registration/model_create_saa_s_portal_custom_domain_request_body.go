/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateSaaSPortalCustomDomainRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSaaSPortalCustomDomainRequestBody{}

// CreateSaaSPortalCustomDomainRequestBody struct for CreateSaaSPortalCustomDomainRequestBody
type CreateSaaSPortalCustomDomainRequestBody struct {
	// The custom domain
	CustomDomain string `json:"customDomain"`
	// The custom domain description
	Description string `json:"description"`
	// The environment type for the custom domain
	EnvironmentType string `json:"environmentType"`
	// The custom domain name
	Name string `json:"name"`
}

type _CreateSaaSPortalCustomDomainRequestBody CreateSaaSPortalCustomDomainRequestBody

// NewCreateSaaSPortalCustomDomainRequestBody instantiates a new CreateSaaSPortalCustomDomainRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSaaSPortalCustomDomainRequestBody(customDomain string, description string, environmentType string, name string) *CreateSaaSPortalCustomDomainRequestBody {
	this := CreateSaaSPortalCustomDomainRequestBody{}
	this.CustomDomain = customDomain
	this.Description = description
	this.EnvironmentType = environmentType
	this.Name = name
	return &this
}

// NewCreateSaaSPortalCustomDomainRequestBodyWithDefaults instantiates a new CreateSaaSPortalCustomDomainRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSaaSPortalCustomDomainRequestBodyWithDefaults() *CreateSaaSPortalCustomDomainRequestBody {
	this := CreateSaaSPortalCustomDomainRequestBody{}
	return &this
}

// GetCustomDomain returns the CustomDomain field value
func (o *CreateSaaSPortalCustomDomainRequestBody) GetCustomDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDomain
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value
// and a boolean to check if the value has been set.
func (o *CreateSaaSPortalCustomDomainRequestBody) GetCustomDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDomain, true
}

// SetCustomDomain sets field value
func (o *CreateSaaSPortalCustomDomainRequestBody) SetCustomDomain(v string) {
	o.CustomDomain = v
}

// GetDescription returns the Description field value
func (o *CreateSaaSPortalCustomDomainRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateSaaSPortalCustomDomainRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateSaaSPortalCustomDomainRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetEnvironmentType returns the EnvironmentType field value
func (o *CreateSaaSPortalCustomDomainRequestBody) GetEnvironmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value
// and a boolean to check if the value has been set.
func (o *CreateSaaSPortalCustomDomainRequestBody) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentType, true
}

// SetEnvironmentType sets field value
func (o *CreateSaaSPortalCustomDomainRequestBody) SetEnvironmentType(v string) {
	o.EnvironmentType = v
}

// GetName returns the Name field value
func (o *CreateSaaSPortalCustomDomainRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSaaSPortalCustomDomainRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSaaSPortalCustomDomainRequestBody) SetName(v string) {
	o.Name = v
}

func (o CreateSaaSPortalCustomDomainRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSaaSPortalCustomDomainRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customDomain"] = o.CustomDomain
	toSerialize["description"] = o.Description
	toSerialize["environmentType"] = o.EnvironmentType
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CreateSaaSPortalCustomDomainRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customDomain",
		"description",
		"environmentType",
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

	varCreateSaaSPortalCustomDomainRequestBody := _CreateSaaSPortalCustomDomainRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSaaSPortalCustomDomainRequestBody)

	if err != nil {
		return err
	}

	*o = CreateSaaSPortalCustomDomainRequestBody(varCreateSaaSPortalCustomDomainRequestBody)

	return err
}

type NullableCreateSaaSPortalCustomDomainRequestBody struct {
	value *CreateSaaSPortalCustomDomainRequestBody
	isSet bool
}

func (v NullableCreateSaaSPortalCustomDomainRequestBody) Get() *CreateSaaSPortalCustomDomainRequestBody {
	return v.value
}

func (v *NullableCreateSaaSPortalCustomDomainRequestBody) Set(val *CreateSaaSPortalCustomDomainRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSaaSPortalCustomDomainRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSaaSPortalCustomDomainRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSaaSPortalCustomDomainRequestBody(val *CreateSaaSPortalCustomDomainRequestBody) *NullableCreateSaaSPortalCustomDomainRequestBody {
	return &NullableCreateSaaSPortalCustomDomainRequestBody{value: val, isSet: true}
}

func (v NullableCreateSaaSPortalCustomDomainRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSaaSPortalCustomDomainRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


