/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the ApproveSubscriptionRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveSubscriptionRequestRequest{}

// ApproveSubscriptionRequestRequest struct for ApproveSubscriptionRequestRequest
type ApproveSubscriptionRequestRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Subscription Request
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ApproveSubscriptionRequestRequest ApproveSubscriptionRequestRequest

// NewApproveSubscriptionRequestRequest instantiates a new ApproveSubscriptionRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveSubscriptionRequestRequest(environmentId string, id string, serviceId string, token string) *ApproveSubscriptionRequestRequest {
	this := ApproveSubscriptionRequestRequest{}
	this.EnvironmentId = environmentId
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewApproveSubscriptionRequestRequestWithDefaults instantiates a new ApproveSubscriptionRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveSubscriptionRequestRequestWithDefaults() *ApproveSubscriptionRequestRequest {
	this := ApproveSubscriptionRequestRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *ApproveSubscriptionRequestRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionRequestRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *ApproveSubscriptionRequestRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *ApproveSubscriptionRequestRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionRequestRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApproveSubscriptionRequestRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *ApproveSubscriptionRequestRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionRequestRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ApproveSubscriptionRequestRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *ApproveSubscriptionRequestRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionRequestRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ApproveSubscriptionRequestRequest) SetToken(v string) {
	o.Token = v
}

func (o ApproveSubscriptionRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproveSubscriptionRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApproveSubscriptionRequestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
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

	varApproveSubscriptionRequestRequest := _ApproveSubscriptionRequestRequest{}

	err = json.Unmarshal(data, &varApproveSubscriptionRequestRequest)

	if err != nil {
		return err
	}

	*o = ApproveSubscriptionRequestRequest(varApproveSubscriptionRequestRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApproveSubscriptionRequestRequest struct {
	value *ApproveSubscriptionRequestRequest
	isSet bool
}

func (v NullableApproveSubscriptionRequestRequest) Get() *ApproveSubscriptionRequestRequest {
	return v.value
}

func (v *NullableApproveSubscriptionRequestRequest) Set(val *ApproveSubscriptionRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveSubscriptionRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveSubscriptionRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveSubscriptionRequestRequest(val *ApproveSubscriptionRequestRequest) *NullableApproveSubscriptionRequestRequest {
	return &NullableApproveSubscriptionRequestRequest{value: val, isSet: true}
}

func (v NullableApproveSubscriptionRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveSubscriptionRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


