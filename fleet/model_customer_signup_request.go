/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the CustomerSignupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSignupRequest{}

// CustomerSignupRequest struct for CustomerSignupRequest
type CustomerSignupRequest struct {
	CompanyDescription *string `json:"companyDescription,omitempty"`
	CompanyUrl *string `json:"companyUrl,omitempty"`
	// Email address of the end-user
	Email string `json:"email"`
	LegalCompanyName *string `json:"legalCompanyName,omitempty"`
	// Name of the end-user
	Name string `json:"name"`
	Password string `json:"password"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _CustomerSignupRequest CustomerSignupRequest

// NewCustomerSignupRequest instantiates a new CustomerSignupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSignupRequest(email string, name string, password string, token string) *CustomerSignupRequest {
	this := CustomerSignupRequest{}
	this.Email = email
	var legalCompanyName string = ""
	this.LegalCompanyName = &legalCompanyName
	this.Name = name
	this.Password = password
	this.Token = token
	return &this
}

// NewCustomerSignupRequestWithDefaults instantiates a new CustomerSignupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSignupRequestWithDefaults() *CustomerSignupRequest {
	this := CustomerSignupRequest{}
	var legalCompanyName string = ""
	this.LegalCompanyName = &legalCompanyName
	return &this
}

// GetCompanyDescription returns the CompanyDescription field value if set, zero value otherwise.
func (o *CustomerSignupRequest) GetCompanyDescription() string {
	if o == nil || IsNil(o.CompanyDescription) {
		var ret string
		return ret
	}
	return *o.CompanyDescription
}

// GetCompanyDescriptionOk returns a tuple with the CompanyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetCompanyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyDescription) {
		return nil, false
	}
	return o.CompanyDescription, true
}

// HasCompanyDescription returns a boolean if a field has been set.
func (o *CustomerSignupRequest) HasCompanyDescription() bool {
	if o != nil && !IsNil(o.CompanyDescription) {
		return true
	}

	return false
}

// SetCompanyDescription gets a reference to the given string and assigns it to the CompanyDescription field.
func (o *CustomerSignupRequest) SetCompanyDescription(v string) {
	o.CompanyDescription = &v
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise.
func (o *CustomerSignupRequest) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl) {
		var ret string
		return ret
	}
	return *o.CompanyUrl
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetCompanyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyUrl) {
		return nil, false
	}
	return o.CompanyUrl, true
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *CustomerSignupRequest) HasCompanyUrl() bool {
	if o != nil && !IsNil(o.CompanyUrl) {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given string and assigns it to the CompanyUrl field.
func (o *CustomerSignupRequest) SetCompanyUrl(v string) {
	o.CompanyUrl = &v
}

// GetEmail returns the Email field value
func (o *CustomerSignupRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerSignupRequest) SetEmail(v string) {
	o.Email = v
}

// GetLegalCompanyName returns the LegalCompanyName field value if set, zero value otherwise.
func (o *CustomerSignupRequest) GetLegalCompanyName() string {
	if o == nil || IsNil(o.LegalCompanyName) {
		var ret string
		return ret
	}
	return *o.LegalCompanyName
}

// GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetLegalCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalCompanyName) {
		return nil, false
	}
	return o.LegalCompanyName, true
}

// HasLegalCompanyName returns a boolean if a field has been set.
func (o *CustomerSignupRequest) HasLegalCompanyName() bool {
	if o != nil && !IsNil(o.LegalCompanyName) {
		return true
	}

	return false
}

// SetLegalCompanyName gets a reference to the given string and assigns it to the LegalCompanyName field.
func (o *CustomerSignupRequest) SetLegalCompanyName(v string) {
	o.LegalCompanyName = &v
}

// GetName returns the Name field value
func (o *CustomerSignupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerSignupRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value
func (o *CustomerSignupRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CustomerSignupRequest) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *CustomerSignupRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CustomerSignupRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CustomerSignupRequest) SetToken(v string) {
	o.Token = v
}

func (o CustomerSignupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSignupRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerSignupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"name",
		"password",
		"token",
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

	varCustomerSignupRequest := _CustomerSignupRequest{}

	err = json.Unmarshal(data, &varCustomerSignupRequest)

	if err != nil {
		return err
	}

	*o = CustomerSignupRequest(varCustomerSignupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "companyDescription")
		delete(additionalProperties, "companyUrl")
		delete(additionalProperties, "email")
		delete(additionalProperties, "legalCompanyName")
		delete(additionalProperties, "name")
		delete(additionalProperties, "password")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerSignupRequest struct {
	value *CustomerSignupRequest
	isSet bool
}

func (v NullableCustomerSignupRequest) Get() *CustomerSignupRequest {
	return v.value
}

func (v *NullableCustomerSignupRequest) Set(val *CustomerSignupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSignupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSignupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSignupRequest(val *CustomerSignupRequest) *NullableCustomerSignupRequest {
	return &NullableCustomerSignupRequest{value: val, isSet: true}
}

func (v NullableCustomerSignupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSignupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


