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

// checks if the DescribeServiceAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceAPIRequest{}

// DescribeServiceAPIRequest struct for DescribeServiceAPIRequest
type DescribeServiceAPIRequest struct {
	// ID of a Service API
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceAPIRequest DescribeServiceAPIRequest

// NewDescribeServiceAPIRequest instantiates a new DescribeServiceAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceAPIRequest(id string, serviceId string, token string) *DescribeServiceAPIRequest {
	this := DescribeServiceAPIRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeServiceAPIRequestWithDefaults instantiates a new DescribeServiceAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceAPIRequestWithDefaults() *DescribeServiceAPIRequest {
	this := DescribeServiceAPIRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DescribeServiceAPIRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceAPIRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeServiceAPIRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceAPIRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceAPIRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceAPIRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeServiceAPIRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceAPIRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeServiceAPIRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeServiceAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceAPIRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeServiceAPIRequest := _DescribeServiceAPIRequest{}

	err = json.Unmarshal(data, &varDescribeServiceAPIRequest)

	if err != nil {
		return err
	}

	*o = DescribeServiceAPIRequest(varDescribeServiceAPIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceAPIRequest struct {
	value *DescribeServiceAPIRequest
	isSet bool
}

func (v NullableDescribeServiceAPIRequest) Get() *DescribeServiceAPIRequest {
	return v.value
}

func (v *NullableDescribeServiceAPIRequest) Set(val *DescribeServiceAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceAPIRequest(val *DescribeServiceAPIRequest) *NullableDescribeServiceAPIRequest {
	return &NullableDescribeServiceAPIRequest{value: val, isSet: true}
}

func (v NullableDescribeServiceAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


