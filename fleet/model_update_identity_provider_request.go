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

// checks if the UpdateIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIdentityProviderRequest{}

// UpdateIdentityProviderRequest struct for UpdateIdentityProviderRequest
type UpdateIdentityProviderRequest struct {
	// The Client ID of the Identity Provider
	ClientId *string `json:"clientId,omitempty"`
	// The Client Secret of the Identity Provider
	ClientSecret *string `json:"clientSecret,omitempty"`
	// ID of an Identity Provider
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIdentityProviderRequest UpdateIdentityProviderRequest

// NewUpdateIdentityProviderRequest instantiates a new UpdateIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIdentityProviderRequest(id string, token string) *UpdateIdentityProviderRequest {
	this := UpdateIdentityProviderRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewUpdateIdentityProviderRequestWithDefaults instantiates a new UpdateIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIdentityProviderRequestWithDefaults() *UpdateIdentityProviderRequest {
	this := UpdateIdentityProviderRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UpdateIdentityProviderRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UpdateIdentityProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *UpdateIdentityProviderRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *UpdateIdentityProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetId returns the Id field value
func (o *UpdateIdentityProviderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateIdentityProviderRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *UpdateIdentityProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateIdentityProviderRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateIdentityProviderRequest := _UpdateIdentityProviderRequest{}

	err = json.Unmarshal(data, &varUpdateIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = UpdateIdentityProviderRequest(varUpdateIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIdentityProviderRequest struct {
	value *UpdateIdentityProviderRequest
	isSet bool
}

func (v NullableUpdateIdentityProviderRequest) Get() *UpdateIdentityProviderRequest {
	return v.value
}

func (v *NullableUpdateIdentityProviderRequest) Set(val *UpdateIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIdentityProviderRequest(val *UpdateIdentityProviderRequest) *NullableUpdateIdentityProviderRequest {
	return &NullableUpdateIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableUpdateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


