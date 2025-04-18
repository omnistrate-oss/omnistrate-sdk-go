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

// checks if the VerifyIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyIdentityProviderRequest{}

// VerifyIdentityProviderRequest struct for VerifyIdentityProviderRequest
type VerifyIdentityProviderRequest struct {
	// ID of an Identity Provider
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _VerifyIdentityProviderRequest VerifyIdentityProviderRequest

// NewVerifyIdentityProviderRequest instantiates a new VerifyIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyIdentityProviderRequest(id string, token string) *VerifyIdentityProviderRequest {
	this := VerifyIdentityProviderRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewVerifyIdentityProviderRequestWithDefaults instantiates a new VerifyIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyIdentityProviderRequestWithDefaults() *VerifyIdentityProviderRequest {
	this := VerifyIdentityProviderRequest{}
	return &this
}

// GetId returns the Id field value
func (o *VerifyIdentityProviderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VerifyIdentityProviderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VerifyIdentityProviderRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *VerifyIdentityProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *VerifyIdentityProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *VerifyIdentityProviderRequest) SetToken(v string) {
	o.Token = v
}

func (o VerifyIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifyIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varVerifyIdentityProviderRequest := _VerifyIdentityProviderRequest{}

	err = json.Unmarshal(data, &varVerifyIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = VerifyIdentityProviderRequest(varVerifyIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifyIdentityProviderRequest struct {
	value *VerifyIdentityProviderRequest
	isSet bool
}

func (v NullableVerifyIdentityProviderRequest) Get() *VerifyIdentityProviderRequest {
	return v.value
}

func (v *NullableVerifyIdentityProviderRequest) Set(val *VerifyIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyIdentityProviderRequest(val *VerifyIdentityProviderRequest) *NullableVerifyIdentityProviderRequest {
	return &NullableVerifyIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableVerifyIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


