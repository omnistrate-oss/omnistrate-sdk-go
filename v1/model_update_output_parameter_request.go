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

// checks if the UpdateOutputParameterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOutputParameterRequest{}

// UpdateOutputParameterRequest struct for UpdateOutputParameterRequest
type UpdateOutputParameterRequest struct {
	// Description of the output variable being exported
	Description *string `json:"description,omitempty"`
	// ID of an Output Parameter
	Id string `json:"id"`
	// External name of the output variable being exported
	Name *string `json:"name,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// Value of the output variable being exported
	Value *string `json:"value,omitempty"`
	// Reference to an input variable that will be used to set the value of the output variable being exported
	ValueRef *string `json:"valueRef,omitempty"`
	// Type of the variable encoding the value
	ValueType *string `json:"valueType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateOutputParameterRequest UpdateOutputParameterRequest

// NewUpdateOutputParameterRequest instantiates a new UpdateOutputParameterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOutputParameterRequest(id string, serviceId string, token string) *UpdateOutputParameterRequest {
	this := UpdateOutputParameterRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateOutputParameterRequestWithDefaults instantiates a new UpdateOutputParameterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOutputParameterRequestWithDefaults() *UpdateOutputParameterRequest {
	this := UpdateOutputParameterRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOutputParameterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOutputParameterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *UpdateOutputParameterRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateOutputParameterRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOutputParameterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOutputParameterRequest) SetName(v string) {
	o.Name = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateOutputParameterRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateOutputParameterRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *UpdateOutputParameterRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateOutputParameterRequest) SetToken(v string) {
	o.Token = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UpdateOutputParameterRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UpdateOutputParameterRequest) SetValue(v string) {
	o.Value = &v
}

// GetValueRef returns the ValueRef field value if set, zero value otherwise.
func (o *UpdateOutputParameterRequest) GetValueRef() string {
	if o == nil || IsNil(o.ValueRef) {
		var ret string
		return ret
	}
	return *o.ValueRef
}

// GetValueRefOk returns a tuple with the ValueRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetValueRefOk() (*string, bool) {
	if o == nil || IsNil(o.ValueRef) {
		return nil, false
	}
	return o.ValueRef, true
}

// SetValueRef gets a reference to the given string and assigns it to the ValueRef field.
func (o *UpdateOutputParameterRequest) SetValueRef(v string) {
	o.ValueRef = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *UpdateOutputParameterRequest) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutputParameterRequest) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *UpdateOutputParameterRequest) SetValueType(v string) {
	o.ValueType = &v
}

func (o UpdateOutputParameterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOutputParameterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

func (o *UpdateOutputParameterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUpdateOutputParameterRequest := _UpdateOutputParameterRequest{}

	err = json.Unmarshal(data, &varUpdateOutputParameterRequest)

	if err != nil {
		return err
	}

	*o = UpdateOutputParameterRequest(varUpdateOutputParameterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "value")
		delete(additionalProperties, "valueRef")
		delete(additionalProperties, "valueType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateOutputParameterRequest struct {
	value *UpdateOutputParameterRequest
	isSet bool
}

func (v NullableUpdateOutputParameterRequest) Get() *UpdateOutputParameterRequest {
	return v.value
}

func (v *NullableUpdateOutputParameterRequest) Set(val *UpdateOutputParameterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOutputParameterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOutputParameterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOutputParameterRequest(val *UpdateOutputParameterRequest) *NullableUpdateOutputParameterRequest {
	return &NullableUpdateOutputParameterRequest{value: val, isSet: true}
}

func (v NullableUpdateOutputParameterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOutputParameterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


