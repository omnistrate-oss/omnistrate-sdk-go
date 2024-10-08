/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateConsumptionUserRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConsumptionUserRequestBody{}

// CreateConsumptionUserRequestBody struct for CreateConsumptionUserRequestBody
type CreateConsumptionUserRequestBody struct {
	// Email address of the user
	Email string `json:"email"`
	// Legal company name of the user.
	LegalCompanyName string `json:"legalCompanyName"`
	// Name of the user
	Name string `json:"name"`
	// Password of the user
	Password string `json:"password"`
}

type _CreateConsumptionUserRequestBody CreateConsumptionUserRequestBody

// NewCreateConsumptionUserRequestBody instantiates a new CreateConsumptionUserRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConsumptionUserRequestBody(email string, legalCompanyName string, name string, password string) *CreateConsumptionUserRequestBody {
	this := CreateConsumptionUserRequestBody{}
	this.Email = email
	this.LegalCompanyName = legalCompanyName
	this.Name = name
	this.Password = password
	return &this
}

// NewCreateConsumptionUserRequestBodyWithDefaults instantiates a new CreateConsumptionUserRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConsumptionUserRequestBodyWithDefaults() *CreateConsumptionUserRequestBody {
	this := CreateConsumptionUserRequestBody{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateConsumptionUserRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateConsumptionUserRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateConsumptionUserRequestBody) SetEmail(v string) {
	o.Email = v
}

// GetLegalCompanyName returns the LegalCompanyName field value
func (o *CreateConsumptionUserRequestBody) GetLegalCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalCompanyName
}

// GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field value
// and a boolean to check if the value has been set.
func (o *CreateConsumptionUserRequestBody) GetLegalCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalCompanyName, true
}

// SetLegalCompanyName sets field value
func (o *CreateConsumptionUserRequestBody) SetLegalCompanyName(v string) {
	o.LegalCompanyName = v
}

// GetName returns the Name field value
func (o *CreateConsumptionUserRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateConsumptionUserRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateConsumptionUserRequestBody) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value
func (o *CreateConsumptionUserRequestBody) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateConsumptionUserRequestBody) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateConsumptionUserRequestBody) SetPassword(v string) {
	o.Password = v
}

func (o CreateConsumptionUserRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConsumptionUserRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["legalCompanyName"] = o.LegalCompanyName
	toSerialize["name"] = o.Name
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *CreateConsumptionUserRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"legalCompanyName",
		"name",
		"password",
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

	varCreateConsumptionUserRequestBody := _CreateConsumptionUserRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateConsumptionUserRequestBody)

	if err != nil {
		return err
	}

	*o = CreateConsumptionUserRequestBody(varCreateConsumptionUserRequestBody)

	return err
}

type NullableCreateConsumptionUserRequestBody struct {
	value *CreateConsumptionUserRequestBody
	isSet bool
}

func (v NullableCreateConsumptionUserRequestBody) Get() *CreateConsumptionUserRequestBody {
	return v.value
}

func (v *NullableCreateConsumptionUserRequestBody) Set(val *CreateConsumptionUserRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConsumptionUserRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConsumptionUserRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConsumptionUserRequestBody(val *CreateConsumptionUserRequestBody) *NullableCreateConsumptionUserRequestBody {
	return &NullableCreateConsumptionUserRequestBody{value: val, isSet: true}
}

func (v NullableCreateConsumptionUserRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConsumptionUserRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


