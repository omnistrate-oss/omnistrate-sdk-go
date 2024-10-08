/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateImageRegistryRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageRegistryRequestBody{}

// CreateImageRegistryRequestBody struct for CreateImageRegistryRequestBody
type CreateImageRegistryRequestBody struct {
	// A brief description of the Image Registry
	Description string `json:"description"`
	// The Image Registry host
	Host string `json:"host"`
	// Name of the Image Registry
	Name string `json:"name"`
	// The password to use when authenticating to the Image Registry
	Password *string `json:"password,omitempty"`
	// The username to use when authenticating to the Image Registry
	Username *string `json:"username,omitempty"`
}

type _CreateImageRegistryRequestBody CreateImageRegistryRequestBody

// NewCreateImageRegistryRequestBody instantiates a new CreateImageRegistryRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageRegistryRequestBody(description string, host string, name string) *CreateImageRegistryRequestBody {
	this := CreateImageRegistryRequestBody{}
	this.Description = description
	this.Host = host
	this.Name = name
	return &this
}

// NewCreateImageRegistryRequestBodyWithDefaults instantiates a new CreateImageRegistryRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageRegistryRequestBodyWithDefaults() *CreateImageRegistryRequestBody {
	this := CreateImageRegistryRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateImageRegistryRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateImageRegistryRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetHost returns the Host field value
func (o *CreateImageRegistryRequestBody) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequestBody) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateImageRegistryRequestBody) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *CreateImageRegistryRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateImageRegistryRequestBody) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateImageRegistryRequestBody) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequestBody) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateImageRegistryRequestBody) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateImageRegistryRequestBody) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateImageRegistryRequestBody) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRegistryRequestBody) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateImageRegistryRequestBody) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateImageRegistryRequestBody) SetUsername(v string) {
	o.Username = &v
}

func (o CreateImageRegistryRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageRegistryRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *CreateImageRegistryRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"host",
		"name",
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

	varCreateImageRegistryRequestBody := _CreateImageRegistryRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateImageRegistryRequestBody)

	if err != nil {
		return err
	}

	*o = CreateImageRegistryRequestBody(varCreateImageRegistryRequestBody)

	return err
}

type NullableCreateImageRegistryRequestBody struct {
	value *CreateImageRegistryRequestBody
	isSet bool
}

func (v NullableCreateImageRegistryRequestBody) Get() *CreateImageRegistryRequestBody {
	return v.value
}

func (v *NullableCreateImageRegistryRequestBody) Set(val *CreateImageRegistryRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageRegistryRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageRegistryRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageRegistryRequestBody(val *CreateImageRegistryRequestBody) *NullableCreateImageRegistryRequestBody {
	return &NullableCreateImageRegistryRequestBody{value: val, isSet: true}
}

func (v NullableCreateImageRegistryRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageRegistryRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


