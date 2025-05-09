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

// checks if the DetachNetworkConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetachNetworkConfigRequest{}

// DetachNetworkConfigRequest struct for DetachNetworkConfigRequest
type DetachNetworkConfigRequest struct {
	// ID of an Infra Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DetachNetworkConfigRequest DetachNetworkConfigRequest

// NewDetachNetworkConfigRequest instantiates a new DetachNetworkConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetachNetworkConfigRequest(id string, serviceId string, token string) *DetachNetworkConfigRequest {
	this := DetachNetworkConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDetachNetworkConfigRequestWithDefaults instantiates a new DetachNetworkConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetachNetworkConfigRequestWithDefaults() *DetachNetworkConfigRequest {
	this := DetachNetworkConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DetachNetworkConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DetachNetworkConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DetachNetworkConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DetachNetworkConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DetachNetworkConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DetachNetworkConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DetachNetworkConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DetachNetworkConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DetachNetworkConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DetachNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetachNetworkConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetachNetworkConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDetachNetworkConfigRequest := _DetachNetworkConfigRequest{}

	err = json.Unmarshal(data, &varDetachNetworkConfigRequest)

	if err != nil {
		return err
	}

	*o = DetachNetworkConfigRequest(varDetachNetworkConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetachNetworkConfigRequest struct {
	value *DetachNetworkConfigRequest
	isSet bool
}

func (v NullableDetachNetworkConfigRequest) Get() *DetachNetworkConfigRequest {
	return v.value
}

func (v *NullableDetachNetworkConfigRequest) Set(val *DetachNetworkConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetachNetworkConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetachNetworkConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetachNetworkConfigRequest(val *DetachNetworkConfigRequest) *NullableDetachNetworkConfigRequest {
	return &NullableDetachNetworkConfigRequest{value: val, isSet: true}
}

func (v NullableDetachNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetachNetworkConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


