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

// checks if the DeleteIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteIdentityProviderRequest{}

// DeleteIdentityProviderRequest struct for DeleteIdentityProviderRequest
type DeleteIdentityProviderRequest struct {
	// ID of an Identity Provider
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeleteIdentityProviderRequest DeleteIdentityProviderRequest

// NewDeleteIdentityProviderRequest instantiates a new DeleteIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteIdentityProviderRequest(id string, token string) *DeleteIdentityProviderRequest {
	this := DeleteIdentityProviderRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewDeleteIdentityProviderRequestWithDefaults instantiates a new DeleteIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteIdentityProviderRequestWithDefaults() *DeleteIdentityProviderRequest {
	this := DeleteIdentityProviderRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteIdentityProviderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteIdentityProviderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteIdentityProviderRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *DeleteIdentityProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeleteIdentityProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeleteIdentityProviderRequest) SetToken(v string) {
	o.Token = v
}

func (o DeleteIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteIdentityProviderRequest := _DeleteIdentityProviderRequest{}

	err = json.Unmarshal(data, &varDeleteIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = DeleteIdentityProviderRequest(varDeleteIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteIdentityProviderRequest struct {
	value *DeleteIdentityProviderRequest
	isSet bool
}

func (v NullableDeleteIdentityProviderRequest) Get() *DeleteIdentityProviderRequest {
	return v.value
}

func (v *NullableDeleteIdentityProviderRequest) Set(val *DeleteIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteIdentityProviderRequest(val *DeleteIdentityProviderRequest) *NullableDeleteIdentityProviderRequest {
	return &NullableDeleteIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableDeleteIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


