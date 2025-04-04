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

// checks if the RemoveResourceDependencyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveResourceDependencyRequest{}

// RemoveResourceDependencyRequest struct for RemoveResourceDependencyRequest
type RemoveResourceDependencyRequest struct {
	// ID of a resource
	Id string `json:"id"`
	// ID of a resource
	ResourceDependencyId string `json:"resourceDependencyId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RemoveResourceDependencyRequest RemoveResourceDependencyRequest

// NewRemoveResourceDependencyRequest instantiates a new RemoveResourceDependencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveResourceDependencyRequest(id string, resourceDependencyId string, serviceId string, token string) *RemoveResourceDependencyRequest {
	this := RemoveResourceDependencyRequest{}
	this.Id = id
	this.ResourceDependencyId = resourceDependencyId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewRemoveResourceDependencyRequestWithDefaults instantiates a new RemoveResourceDependencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveResourceDependencyRequestWithDefaults() *RemoveResourceDependencyRequest {
	this := RemoveResourceDependencyRequest{}
	return &this
}

// GetId returns the Id field value
func (o *RemoveResourceDependencyRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RemoveResourceDependencyRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RemoveResourceDependencyRequest) SetId(v string) {
	o.Id = v
}

// GetResourceDependencyId returns the ResourceDependencyId field value
func (o *RemoveResourceDependencyRequest) GetResourceDependencyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceDependencyId
}

// GetResourceDependencyIdOk returns a tuple with the ResourceDependencyId field value
// and a boolean to check if the value has been set.
func (o *RemoveResourceDependencyRequest) GetResourceDependencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceDependencyId, true
}

// SetResourceDependencyId sets field value
func (o *RemoveResourceDependencyRequest) SetResourceDependencyId(v string) {
	o.ResourceDependencyId = v
}

// GetServiceId returns the ServiceId field value
func (o *RemoveResourceDependencyRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *RemoveResourceDependencyRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *RemoveResourceDependencyRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *RemoveResourceDependencyRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RemoveResourceDependencyRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RemoveResourceDependencyRequest) SetToken(v string) {
	o.Token = v
}

func (o RemoveResourceDependencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveResourceDependencyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resourceDependencyId"] = o.ResourceDependencyId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveResourceDependencyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"resourceDependencyId",
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

	varRemoveResourceDependencyRequest := _RemoveResourceDependencyRequest{}

	err = json.Unmarshal(data, &varRemoveResourceDependencyRequest)

	if err != nil {
		return err
	}

	*o = RemoveResourceDependencyRequest(varRemoveResourceDependencyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "resourceDependencyId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveResourceDependencyRequest struct {
	value *RemoveResourceDependencyRequest
	isSet bool
}

func (v NullableRemoveResourceDependencyRequest) Get() *RemoveResourceDependencyRequest {
	return v.value
}

func (v *NullableRemoveResourceDependencyRequest) Set(val *RemoveResourceDependencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveResourceDependencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveResourceDependencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveResourceDependencyRequest(val *RemoveResourceDependencyRequest) *NullableRemoveResourceDependencyRequest {
	return &NullableRemoveResourceDependencyRequest{value: val, isSet: true}
}

func (v NullableRemoveResourceDependencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveResourceDependencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


