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

// checks if the ContactusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactusRequest{}

// ContactusRequest struct for ContactusRequest
type ContactusRequest struct {
	Company string `json:"company"`
	Email string `json:"email"`
	Message string `json:"message"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ContactusRequest ContactusRequest

// NewContactusRequest instantiates a new ContactusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactusRequest(company string, email string, message string, name string) *ContactusRequest {
	this := ContactusRequest{}
	this.Company = company
	this.Email = email
	this.Message = message
	this.Name = name
	return &this
}

// NewContactusRequestWithDefaults instantiates a new ContactusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactusRequestWithDefaults() *ContactusRequest {
	this := ContactusRequest{}
	return &this
}

// GetCompany returns the Company field value
func (o *ContactusRequest) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *ContactusRequest) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *ContactusRequest) SetCompany(v string) {
	o.Company = v
}

// GetEmail returns the Email field value
func (o *ContactusRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ContactusRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ContactusRequest) SetEmail(v string) {
	o.Email = v
}

// GetMessage returns the Message field value
func (o *ContactusRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ContactusRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ContactusRequest) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value
func (o *ContactusRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContactusRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContactusRequest) SetName(v string) {
	o.Name = v
}

func (o ContactusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	toSerialize["email"] = o.Email
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContactusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
		"email",
		"message",
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

	varContactusRequest := _ContactusRequest{}

	err = json.Unmarshal(data, &varContactusRequest)

	if err != nil {
		return err
	}

	*o = ContactusRequest(varContactusRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company")
		delete(additionalProperties, "email")
		delete(additionalProperties, "message")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactusRequest struct {
	value *ContactusRequest
	isSet bool
}

func (v NullableContactusRequest) Get() *ContactusRequest {
	return v.value
}

func (v *NullableContactusRequest) Set(val *ContactusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactusRequest(val *ContactusRequest) *NullableContactusRequest {
	return &NullableContactusRequest{value: val, isSet: true}
}

func (v NullableContactusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


