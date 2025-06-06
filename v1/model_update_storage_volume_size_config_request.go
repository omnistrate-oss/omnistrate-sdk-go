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

// checks if the UpdateStorageVolumeSizeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageVolumeSizeConfigRequest{}

// UpdateStorageVolumeSizeConfigRequest struct for UpdateStorageVolumeSizeConfigRequest
type UpdateStorageVolumeSizeConfigRequest struct {
	// ID of a Storage Volume Config
	Id string `json:"id"`
	// The storage size (in Gi) provisioned for the configured instance storage type
	InstanceStorageSizeGi string `json:"instanceStorageSizeGi"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorageVolumeSizeConfigRequest UpdateStorageVolumeSizeConfigRequest

// NewUpdateStorageVolumeSizeConfigRequest instantiates a new UpdateStorageVolumeSizeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageVolumeSizeConfigRequest(id string, instanceStorageSizeGi string, serviceId string, token string) *UpdateStorageVolumeSizeConfigRequest {
	this := UpdateStorageVolumeSizeConfigRequest{}
	this.Id = id
	this.InstanceStorageSizeGi = instanceStorageSizeGi
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateStorageVolumeSizeConfigRequestWithDefaults instantiates a new UpdateStorageVolumeSizeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageVolumeSizeConfigRequestWithDefaults() *UpdateStorageVolumeSizeConfigRequest {
	this := UpdateStorageVolumeSizeConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateStorageVolumeSizeConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeSizeConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateStorageVolumeSizeConfigRequest) SetId(v string) {
	o.Id = v
}

// GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field value
func (o *UpdateStorageVolumeSizeConfigRequest) GetInstanceStorageSizeGi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceStorageSizeGi
}

// GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeSizeConfigRequest) GetInstanceStorageSizeGiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceStorageSizeGi, true
}

// SetInstanceStorageSizeGi sets field value
func (o *UpdateStorageVolumeSizeConfigRequest) SetInstanceStorageSizeGi(v string) {
	o.InstanceStorageSizeGi = v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateStorageVolumeSizeConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeSizeConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateStorageVolumeSizeConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *UpdateStorageVolumeSizeConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeSizeConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateStorageVolumeSizeConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateStorageVolumeSizeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageVolumeSizeConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["instanceStorageSizeGi"] = o.InstanceStorageSizeGi
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorageVolumeSizeConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"instanceStorageSizeGi",
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

	varUpdateStorageVolumeSizeConfigRequest := _UpdateStorageVolumeSizeConfigRequest{}

	err = json.Unmarshal(data, &varUpdateStorageVolumeSizeConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdateStorageVolumeSizeConfigRequest(varUpdateStorageVolumeSizeConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "instanceStorageSizeGi")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorageVolumeSizeConfigRequest struct {
	value *UpdateStorageVolumeSizeConfigRequest
	isSet bool
}

func (v NullableUpdateStorageVolumeSizeConfigRequest) Get() *UpdateStorageVolumeSizeConfigRequest {
	return v.value
}

func (v *NullableUpdateStorageVolumeSizeConfigRequest) Set(val *UpdateStorageVolumeSizeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorageVolumeSizeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorageVolumeSizeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorageVolumeSizeConfigRequest(val *UpdateStorageVolumeSizeConfigRequest) *NullableUpdateStorageVolumeSizeConfigRequest {
	return &NullableUpdateStorageVolumeSizeConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateStorageVolumeSizeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorageVolumeSizeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


