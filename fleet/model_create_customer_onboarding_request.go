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

// checks if the CreateCustomerOnboardingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerOnboardingRequest{}

// CreateCustomerOnboardingRequest struct for CreateCustomerOnboardingRequest
type CreateCustomerOnboardingRequest struct {
	// DEPRECATED: Name will be generated automatically.
	Name *string `json:"name,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _CreateCustomerOnboardingRequest CreateCustomerOnboardingRequest

// NewCreateCustomerOnboardingRequest instantiates a new CreateCustomerOnboardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerOnboardingRequest(token string) *CreateCustomerOnboardingRequest {
	this := CreateCustomerOnboardingRequest{}
	this.Token = token
	return &this
}

// NewCreateCustomerOnboardingRequestWithDefaults instantiates a new CreateCustomerOnboardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerOnboardingRequestWithDefaults() *CreateCustomerOnboardingRequest {
	this := CreateCustomerOnboardingRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateCustomerOnboardingRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerOnboardingRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateCustomerOnboardingRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCustomerOnboardingRequest) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value
func (o *CreateCustomerOnboardingRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerOnboardingRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateCustomerOnboardingRequest) SetToken(v string) {
	o.Token = v
}

func (o CreateCustomerOnboardingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerOnboardingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCustomerOnboardingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateCustomerOnboardingRequest := _CreateCustomerOnboardingRequest{}

	err = json.Unmarshal(data, &varCreateCustomerOnboardingRequest)

	if err != nil {
		return err
	}

	*o = CreateCustomerOnboardingRequest(varCreateCustomerOnboardingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCustomerOnboardingRequest struct {
	value *CreateCustomerOnboardingRequest
	isSet bool
}

func (v NullableCreateCustomerOnboardingRequest) Get() *CreateCustomerOnboardingRequest {
	return v.value
}

func (v *NullableCreateCustomerOnboardingRequest) Set(val *CreateCustomerOnboardingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerOnboardingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerOnboardingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerOnboardingRequest(val *CreateCustomerOnboardingRequest) *NullableCreateCustomerOnboardingRequest {
	return &NullableCreateCustomerOnboardingRequest{value: val, isSet: true}
}

func (v NullableCreateCustomerOnboardingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerOnboardingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


