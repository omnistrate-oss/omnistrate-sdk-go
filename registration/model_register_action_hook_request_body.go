/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegisterActionHookRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterActionHookRequestBody{}

// RegisterActionHookRequestBody struct for RegisterActionHookRequestBody
type RegisterActionHookRequestBody struct {
	// The Base64 encoded command template to execute
	CommandTemplate string `json:"commandTemplate"`
	// The scope of the hook
	Scope string `json:"scope"`
	// The type of hook to execute
	Type string `json:"type"`
}

type _RegisterActionHookRequestBody RegisterActionHookRequestBody

// NewRegisterActionHookRequestBody instantiates a new RegisterActionHookRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterActionHookRequestBody(commandTemplate string, scope string, type_ string) *RegisterActionHookRequestBody {
	this := RegisterActionHookRequestBody{}
	this.CommandTemplate = commandTemplate
	this.Scope = scope
	this.Type = type_
	return &this
}

// NewRegisterActionHookRequestBodyWithDefaults instantiates a new RegisterActionHookRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterActionHookRequestBodyWithDefaults() *RegisterActionHookRequestBody {
	this := RegisterActionHookRequestBody{}
	return &this
}

// GetCommandTemplate returns the CommandTemplate field value
func (o *RegisterActionHookRequestBody) GetCommandTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandTemplate
}

// GetCommandTemplateOk returns a tuple with the CommandTemplate field value
// and a boolean to check if the value has been set.
func (o *RegisterActionHookRequestBody) GetCommandTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandTemplate, true
}

// SetCommandTemplate sets field value
func (o *RegisterActionHookRequestBody) SetCommandTemplate(v string) {
	o.CommandTemplate = v
}

// GetScope returns the Scope field value
func (o *RegisterActionHookRequestBody) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *RegisterActionHookRequestBody) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *RegisterActionHookRequestBody) SetScope(v string) {
	o.Scope = v
}

// GetType returns the Type field value
func (o *RegisterActionHookRequestBody) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RegisterActionHookRequestBody) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RegisterActionHookRequestBody) SetType(v string) {
	o.Type = v
}

func (o RegisterActionHookRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterActionHookRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandTemplate"] = o.CommandTemplate
	toSerialize["scope"] = o.Scope
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *RegisterActionHookRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commandTemplate",
		"scope",
		"type",
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

	varRegisterActionHookRequestBody := _RegisterActionHookRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterActionHookRequestBody)

	if err != nil {
		return err
	}

	*o = RegisterActionHookRequestBody(varRegisterActionHookRequestBody)

	return err
}

type NullableRegisterActionHookRequestBody struct {
	value *RegisterActionHookRequestBody
	isSet bool
}

func (v NullableRegisterActionHookRequestBody) Get() *RegisterActionHookRequestBody {
	return v.value
}

func (v *NullableRegisterActionHookRequestBody) Set(val *RegisterActionHookRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterActionHookRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterActionHookRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterActionHookRequestBody(val *RegisterActionHookRequestBody) *NullableRegisterActionHookRequestBody {
	return &NullableRegisterActionHookRequestBody{value: val, isSet: true}
}

func (v NullableRegisterActionHookRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterActionHookRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


