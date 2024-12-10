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

// checks if the CreateIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIdentityProviderRequest{}

// CreateIdentityProviderRequest struct for CreateIdentityProviderRequest
type CreateIdentityProviderRequest struct {
	// The Client ID of the Identity Provider
	ClientId string `json:"clientId"`
	// The Client Secret of the Identity Provider
	ClientSecret string `json:"clientSecret"`
	// The name of the identity provider
	IdentityProviderName string `json:"identityProviderName"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _CreateIdentityProviderRequest CreateIdentityProviderRequest

// NewCreateIdentityProviderRequest instantiates a new CreateIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIdentityProviderRequest(clientId string, clientSecret string, identityProviderName string, token string) *CreateIdentityProviderRequest {
	this := CreateIdentityProviderRequest{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.IdentityProviderName = identityProviderName
	this.Token = token
	return &this
}

// NewCreateIdentityProviderRequestWithDefaults instantiates a new CreateIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIdentityProviderRequestWithDefaults() *CreateIdentityProviderRequest {
	this := CreateIdentityProviderRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *CreateIdentityProviderRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CreateIdentityProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *CreateIdentityProviderRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *CreateIdentityProviderRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *CreateIdentityProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *CreateIdentityProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetIdentityProviderName returns the IdentityProviderName field value
func (o *CreateIdentityProviderRequest) GetIdentityProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value
// and a boolean to check if the value has been set.
func (o *CreateIdentityProviderRequest) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderName, true
}

// SetIdentityProviderName sets field value
func (o *CreateIdentityProviderRequest) SetIdentityProviderName(v string) {
	o.IdentityProviderName = v
}

// GetToken returns the Token field value
func (o *CreateIdentityProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateIdentityProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateIdentityProviderRequest) SetToken(v string) {
	o.Token = v
}

func (o CreateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["identityProviderName"] = o.IdentityProviderName
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"clientSecret",
		"identityProviderName",
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

	varCreateIdentityProviderRequest := _CreateIdentityProviderRequest{}

	err = json.Unmarshal(data, &varCreateIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = CreateIdentityProviderRequest(varCreateIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "identityProviderName")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateIdentityProviderRequest struct {
	value *CreateIdentityProviderRequest
	isSet bool
}

func (v NullableCreateIdentityProviderRequest) Get() *CreateIdentityProviderRequest {
	return v.value
}

func (v *NullableCreateIdentityProviderRequest) Set(val *CreateIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIdentityProviderRequest(val *CreateIdentityProviderRequest) *NullableCreateIdentityProviderRequest {
	return &NullableCreateIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableCreateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


