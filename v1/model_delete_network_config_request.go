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

// checks if the DeleteNetworkConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteNetworkConfigRequest{}

// DeleteNetworkConfigRequest struct for DeleteNetworkConfigRequest
type DeleteNetworkConfigRequest struct {
	// ID of a Network Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeleteNetworkConfigRequest DeleteNetworkConfigRequest

// NewDeleteNetworkConfigRequest instantiates a new DeleteNetworkConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNetworkConfigRequest(id string, serviceId string, token string) *DeleteNetworkConfigRequest {
	this := DeleteNetworkConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDeleteNetworkConfigRequestWithDefaults instantiates a new DeleteNetworkConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNetworkConfigRequestWithDefaults() *DeleteNetworkConfigRequest {
	this := DeleteNetworkConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteNetworkConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteNetworkConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteNetworkConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DeleteNetworkConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DeleteNetworkConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DeleteNetworkConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DeleteNetworkConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeleteNetworkConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeleteNetworkConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DeleteNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteNetworkConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteNetworkConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteNetworkConfigRequest := _DeleteNetworkConfigRequest{}

	err = json.Unmarshal(data, &varDeleteNetworkConfigRequest)

	if err != nil {
		return err
	}

	*o = DeleteNetworkConfigRequest(varDeleteNetworkConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteNetworkConfigRequest struct {
	value *DeleteNetworkConfigRequest
	isSet bool
}

func (v NullableDeleteNetworkConfigRequest) Get() *DeleteNetworkConfigRequest {
	return v.value
}

func (v *NullableDeleteNetworkConfigRequest) Set(val *DeleteNetworkConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNetworkConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNetworkConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNetworkConfigRequest(val *DeleteNetworkConfigRequest) *NullableDeleteNetworkConfigRequest {
	return &NullableDeleteNetworkConfigRequest{value: val, isSet: true}
}

func (v NullableDeleteNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNetworkConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


