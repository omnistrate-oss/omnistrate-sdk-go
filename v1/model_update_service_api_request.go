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

// checks if the UpdateServiceAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceAPIRequest{}

// UpdateServiceAPIRequest struct for UpdateServiceAPIRequest
type UpdateServiceAPIRequest struct {
	// A brief description of the service API bundle
	Description *string `json:"description,omitempty"`
	// ID of a Service API
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateServiceAPIRequest UpdateServiceAPIRequest

// NewUpdateServiceAPIRequest instantiates a new UpdateServiceAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceAPIRequest(id string, serviceId string, token string) *UpdateServiceAPIRequest {
	this := UpdateServiceAPIRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateServiceAPIRequestWithDefaults instantiates a new UpdateServiceAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceAPIRequestWithDefaults() *UpdateServiceAPIRequest {
	this := UpdateServiceAPIRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateServiceAPIRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceAPIRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateServiceAPIRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *UpdateServiceAPIRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceAPIRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateServiceAPIRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateServiceAPIRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceAPIRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateServiceAPIRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *UpdateServiceAPIRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceAPIRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateServiceAPIRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateServiceAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateServiceAPIRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateServiceAPIRequest := _UpdateServiceAPIRequest{}

	err = json.Unmarshal(data, &varUpdateServiceAPIRequest)

	if err != nil {
		return err
	}

	*o = UpdateServiceAPIRequest(varUpdateServiceAPIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateServiceAPIRequest struct {
	value *UpdateServiceAPIRequest
	isSet bool
}

func (v NullableUpdateServiceAPIRequest) Get() *UpdateServiceAPIRequest {
	return v.value
}

func (v *NullableUpdateServiceAPIRequest) Set(val *UpdateServiceAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceAPIRequest(val *UpdateServiceAPIRequest) *NullableUpdateServiceAPIRequest {
	return &NullableUpdateServiceAPIRequest{value: val, isSet: true}
}

func (v NullableUpdateServiceAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


