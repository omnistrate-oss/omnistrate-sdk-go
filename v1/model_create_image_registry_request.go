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

// checks if the CreateImageRegistryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageRegistryRequest{}

// CreateImageRegistryRequest Create an HTTP API v2 Docker Image Registry
type CreateImageRegistryRequest struct {
	// A brief description of the Image Registry
	Description string `json:"description"`
	// The Image Registry host
	Host string `json:"host"`
	// Name of the Image Registry
	Name string `json:"name"`
	// The password to use when authenticating to the Image Registry
	Password *string `json:"password,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// The username to use when authenticating to the Image Registry
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateImageRegistryRequest CreateImageRegistryRequest

// NewCreateImageRegistryRequest instantiates a new CreateImageRegistryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageRegistryRequest(description string, host string, name string, token string) *CreateImageRegistryRequest {
	this := CreateImageRegistryRequest{}
	this.Description = description
	this.Host = host
	this.Name = name
	this.Token = token
	return &this
}

// NewCreateImageRegistryRequestWithDefaults instantiates a new CreateImageRegistryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageRegistryRequestWithDefaults() *CreateImageRegistryRequest {
	this := CreateImageRegistryRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateImageRegistryRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateImageRegistryRequest) SetDescription(v string) {
	o.Description = v
}

// GetHost returns the Host field value
func (o *CreateImageRegistryRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateImageRegistryRequest) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *CreateImageRegistryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateImageRegistryRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateImageRegistryRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateImageRegistryRequest) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value
func (o *CreateImageRegistryRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateImageRegistryRequest) SetToken(v string) {
	o.Token = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateImageRegistryRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateImageRegistryRequest) SetUsername(v string) {
	o.Username = &v
}

func (o CreateImageRegistryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageRegistryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateImageRegistryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"host",
		"name",
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

	varCreateImageRegistryRequest := _CreateImageRegistryRequest{}

	err = json.Unmarshal(data, &varCreateImageRegistryRequest)

	if err != nil {
		return err
	}

	*o = CreateImageRegistryRequest(varCreateImageRegistryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "host")
		delete(additionalProperties, "name")
		delete(additionalProperties, "password")
		delete(additionalProperties, "token")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateImageRegistryRequest struct {
	value *CreateImageRegistryRequest
	isSet bool
}

func (v NullableCreateImageRegistryRequest) Get() *CreateImageRegistryRequest {
	return v.value
}

func (v *NullableCreateImageRegistryRequest) Set(val *CreateImageRegistryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageRegistryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageRegistryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageRegistryRequest(val *CreateImageRegistryRequest) *NullableCreateImageRegistryRequest {
	return &NullableCreateImageRegistryRequest{value: val, isSet: true}
}

func (v NullableCreateImageRegistryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageRegistryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


