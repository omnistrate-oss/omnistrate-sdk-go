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

// checks if the SignupRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignupRequestBody{}

// SignupRequestBody struct for SignupRequestBody
type SignupRequestBody struct {
	CompanyDescription *string `json:"companyDescription,omitempty"`
	CompanyUrl *string `json:"companyUrl,omitempty"`
	Email string `json:"email"`
	LegalCompanyName *string `json:"legalCompanyName,omitempty"`
	Name string `json:"name"`
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _SignupRequestBody SignupRequestBody

// NewSignupRequestBody instantiates a new SignupRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignupRequestBody(email string, name string, password string) *SignupRequestBody {
	this := SignupRequestBody{}
	this.Email = email
	var legalCompanyName string = ""
	this.LegalCompanyName = &legalCompanyName
	this.Name = name
	this.Password = password
	return &this
}

// NewSignupRequestBodyWithDefaults instantiates a new SignupRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignupRequestBodyWithDefaults() *SignupRequestBody {
	this := SignupRequestBody{}
	var legalCompanyName string = ""
	this.LegalCompanyName = &legalCompanyName
	return &this
}

// GetCompanyDescription returns the CompanyDescription field value if set, zero value otherwise.
func (o *SignupRequestBody) GetCompanyDescription() string {
	if o == nil || IsNil(o.CompanyDescription) {
		var ret string
		return ret
	}
	return *o.CompanyDescription
}

// GetCompanyDescriptionOk returns a tuple with the CompanyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetCompanyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyDescription) {
		return nil, false
	}
	return o.CompanyDescription, true
}

// SetCompanyDescription gets a reference to the given string and assigns it to the CompanyDescription field.
func (o *SignupRequestBody) SetCompanyDescription(v string) {
	o.CompanyDescription = &v
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise.
func (o *SignupRequestBody) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl) {
		var ret string
		return ret
	}
	return *o.CompanyUrl
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetCompanyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyUrl) {
		return nil, false
	}
	return o.CompanyUrl, true
}

// SetCompanyUrl gets a reference to the given string and assigns it to the CompanyUrl field.
func (o *SignupRequestBody) SetCompanyUrl(v string) {
	o.CompanyUrl = &v
}

// GetEmail returns the Email field value
func (o *SignupRequestBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SignupRequestBody) SetEmail(v string) {
	o.Email = v
}

// GetLegalCompanyName returns the LegalCompanyName field value if set, zero value otherwise.
func (o *SignupRequestBody) GetLegalCompanyName() string {
	if o == nil || IsNil(o.LegalCompanyName) {
		var ret string
		return ret
	}
	return *o.LegalCompanyName
}

// GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetLegalCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalCompanyName) {
		return nil, false
	}
	return o.LegalCompanyName, true
}

// SetLegalCompanyName gets a reference to the given string and assigns it to the LegalCompanyName field.
func (o *SignupRequestBody) SetLegalCompanyName(v string) {
	o.LegalCompanyName = &v
}

// GetName returns the Name field value
func (o *SignupRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignupRequestBody) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value
func (o *SignupRequestBody) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SignupRequestBody) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SignupRequestBody) SetPassword(v string) {
	o.Password = v
}

func (o SignupRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignupRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyDescription) {
		toSerialize["companyDescription"] = o.CompanyDescription
	}
	if !IsNil(o.CompanyUrl) {
		toSerialize["companyUrl"] = o.CompanyUrl
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.LegalCompanyName) {
		toSerialize["legalCompanyName"] = o.LegalCompanyName
	}
	toSerialize["name"] = o.Name
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignupRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varSignupRequestBody := _SignupRequestBody{}

	err = json.Unmarshal(data, &varSignupRequestBody)

	if err != nil {
		return err
	}

	*o = SignupRequestBody(varSignupRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "companyDescription")
		delete(additionalProperties, "companyUrl")
		delete(additionalProperties, "email")
		delete(additionalProperties, "legalCompanyName")
		delete(additionalProperties, "name")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignupRequestBody struct {
	value *SignupRequestBody
	isSet bool
}

func (v NullableSignupRequestBody) Get() *SignupRequestBody {
	return v.value
}

func (v *NullableSignupRequestBody) Set(val *SignupRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupRequestBody(val *SignupRequestBody) *NullableSignupRequestBody {
	return &NullableSignupRequestBody{value: val, isSet: true}
}

func (v NullableSignupRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


