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

// checks if the DescribeCustomerOnboardingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeCustomerOnboardingRequest{}

// DescribeCustomerOnboardingRequest struct for DescribeCustomerOnboardingRequest
type DescribeCustomerOnboardingRequest struct {
	// ID of a Customer Onboarding
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeCustomerOnboardingRequest DescribeCustomerOnboardingRequest

// NewDescribeCustomerOnboardingRequest instantiates a new DescribeCustomerOnboardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCustomerOnboardingRequest(id string, token string) *DescribeCustomerOnboardingRequest {
	this := DescribeCustomerOnboardingRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewDescribeCustomerOnboardingRequestWithDefaults instantiates a new DescribeCustomerOnboardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCustomerOnboardingRequestWithDefaults() *DescribeCustomerOnboardingRequest {
	this := DescribeCustomerOnboardingRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DescribeCustomerOnboardingRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeCustomerOnboardingRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeCustomerOnboardingRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *DescribeCustomerOnboardingRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeCustomerOnboardingRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeCustomerOnboardingRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeCustomerOnboardingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCustomerOnboardingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeCustomerOnboardingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varDescribeCustomerOnboardingRequest := _DescribeCustomerOnboardingRequest{}

	err = json.Unmarshal(data, &varDescribeCustomerOnboardingRequest)

	if err != nil {
		return err
	}

	*o = DescribeCustomerOnboardingRequest(varDescribeCustomerOnboardingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeCustomerOnboardingRequest struct {
	value *DescribeCustomerOnboardingRequest
	isSet bool
}

func (v NullableDescribeCustomerOnboardingRequest) Get() *DescribeCustomerOnboardingRequest {
	return v.value
}

func (v *NullableDescribeCustomerOnboardingRequest) Set(val *DescribeCustomerOnboardingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCustomerOnboardingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCustomerOnboardingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCustomerOnboardingRequest(val *DescribeCustomerOnboardingRequest) *NullableDescribeCustomerOnboardingRequest {
	return &NullableDescribeCustomerOnboardingRequest{value: val, isSet: true}
}

func (v NullableDescribeCustomerOnboardingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCustomerOnboardingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


