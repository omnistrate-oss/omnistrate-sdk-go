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

// checks if the CreateOutputParameterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOutputParameterRequest{}

// CreateOutputParameterRequest struct for CreateOutputParameterRequest
type CreateOutputParameterRequest struct {
	// Description of the output variable being exported
	Description string `json:"description"`
	// Key of the output variable being exported
	Key string `json:"key" validate:"regexp=^[a-zA-Z][a-zA-Z0-9_]*$"`
	// External name of the output variable being exported
	Name string `json:"name"`
	// ID of a resource
	ResourceId string `json:"resourceId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// Value of the output variable being exported
	Value *string `json:"value,omitempty"`
	// Reference to another variable that will be used to set the value of the output variable being exported
	ValueRef *string `json:"valueRef,omitempty"`
	// Type of the variable encoding the value
	ValueType *string `json:"valueType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOutputParameterRequest CreateOutputParameterRequest

// NewCreateOutputParameterRequest instantiates a new CreateOutputParameterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutputParameterRequest(description string, key string, name string, resourceId string, serviceId string, token string) *CreateOutputParameterRequest {
	this := CreateOutputParameterRequest{}
	this.Description = description
	this.Key = key
	this.Name = name
	this.ResourceId = resourceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewCreateOutputParameterRequestWithDefaults instantiates a new CreateOutputParameterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutputParameterRequestWithDefaults() *CreateOutputParameterRequest {
	this := CreateOutputParameterRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateOutputParameterRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateOutputParameterRequest) SetDescription(v string) {
	o.Description = v
}

// GetKey returns the Key field value
func (o *CreateOutputParameterRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateOutputParameterRequest) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *CreateOutputParameterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOutputParameterRequest) SetName(v string) {
	o.Name = v
}

// GetResourceId returns the ResourceId field value
func (o *CreateOutputParameterRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateOutputParameterRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetServiceId returns the ServiceId field value
func (o *CreateOutputParameterRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CreateOutputParameterRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *CreateOutputParameterRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateOutputParameterRequest) SetToken(v string) {
	o.Token = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateOutputParameterRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateOutputParameterRequest) SetValue(v string) {
	o.Value = &v
}

// GetValueRef returns the ValueRef field value if set, zero value otherwise.
func (o *CreateOutputParameterRequest) GetValueRef() string {
	if o == nil || IsNil(o.ValueRef) {
		var ret string
		return ret
	}
	return *o.ValueRef
}

// GetValueRefOk returns a tuple with the ValueRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetValueRefOk() (*string, bool) {
	if o == nil || IsNil(o.ValueRef) {
		return nil, false
	}
	return o.ValueRef, true
}

// SetValueRef gets a reference to the given string and assigns it to the ValueRef field.
func (o *CreateOutputParameterRequest) SetValueRef(v string) {
	o.ValueRef = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *CreateOutputParameterRequest) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputParameterRequest) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *CreateOutputParameterRequest) SetValueType(v string) {
	o.ValueType = &v
}

func (o CreateOutputParameterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutputParameterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueRef) {
		toSerialize["valueRef"] = o.ValueRef
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOutputParameterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"key",
		"name",
		"resourceId",
		"serviceId",
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

	varCreateOutputParameterRequest := _CreateOutputParameterRequest{}

	err = json.Unmarshal(data, &varCreateOutputParameterRequest)

	if err != nil {
		return err
	}

	*o = CreateOutputParameterRequest(varCreateOutputParameterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "value")
		delete(additionalProperties, "valueRef")
		delete(additionalProperties, "valueType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOutputParameterRequest struct {
	value *CreateOutputParameterRequest
	isSet bool
}

func (v NullableCreateOutputParameterRequest) Get() *CreateOutputParameterRequest {
	return v.value
}

func (v *NullableCreateOutputParameterRequest) Set(val *CreateOutputParameterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutputParameterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutputParameterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutputParameterRequest(val *CreateOutputParameterRequest) *NullableCreateOutputParameterRequest {
	return &NullableCreateOutputParameterRequest{value: val, isSet: true}
}

func (v NullableCreateOutputParameterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutputParameterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


