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

// checks if the DeleteImageConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteImageConfigRequest{}

// DeleteImageConfigRequest Delete an image configuration
type DeleteImageConfigRequest struct {
	// ID of an Image Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeleteImageConfigRequest DeleteImageConfigRequest

// NewDeleteImageConfigRequest instantiates a new DeleteImageConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteImageConfigRequest(id string, serviceId string, token string) *DeleteImageConfigRequest {
	this := DeleteImageConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDeleteImageConfigRequestWithDefaults instantiates a new DeleteImageConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteImageConfigRequestWithDefaults() *DeleteImageConfigRequest {
	this := DeleteImageConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteImageConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteImageConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteImageConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DeleteImageConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DeleteImageConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DeleteImageConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DeleteImageConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeleteImageConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeleteImageConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DeleteImageConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteImageConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteImageConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteImageConfigRequest := _DeleteImageConfigRequest{}

	err = json.Unmarshal(data, &varDeleteImageConfigRequest)

	if err != nil {
		return err
	}

	*o = DeleteImageConfigRequest(varDeleteImageConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteImageConfigRequest struct {
	value *DeleteImageConfigRequest
	isSet bool
}

func (v NullableDeleteImageConfigRequest) Get() *DeleteImageConfigRequest {
	return v.value
}

func (v *NullableDeleteImageConfigRequest) Set(val *DeleteImageConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteImageConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteImageConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteImageConfigRequest(val *DeleteImageConfigRequest) *NullableDeleteImageConfigRequest {
	return &NullableDeleteImageConfigRequest{value: val, isSet: true}
}

func (v NullableDeleteImageConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteImageConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


