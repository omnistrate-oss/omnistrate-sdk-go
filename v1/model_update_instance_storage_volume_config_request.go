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

// checks if the UpdateInstanceStorageVolumeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceStorageVolumeConfigRequest{}

// UpdateInstanceStorageVolumeConfigRequest struct for UpdateInstanceStorageVolumeConfigRequest
type UpdateInstanceStorageVolumeConfigRequest struct {
	// ID of a Storage Volume Config
	Id string `json:"id"`
	// The IOPS provisioned for the configured instance storage type
	InstanceStorageIops *string `json:"instanceStorageIops,omitempty"`
	// The throughput (in MiBps) provisioned for the configured instance storage type
	InstanceStorageThroughputMiBps *string `json:"instanceStorageThroughputMiBps,omitempty"`
	// The type of the storage for a compute instance
	InstanceStorageType *string `json:"instanceStorageType,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateInstanceStorageVolumeConfigRequest UpdateInstanceStorageVolumeConfigRequest

// NewUpdateInstanceStorageVolumeConfigRequest instantiates a new UpdateInstanceStorageVolumeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceStorageVolumeConfigRequest(id string, serviceId string, token string) *UpdateInstanceStorageVolumeConfigRequest {
	this := UpdateInstanceStorageVolumeConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateInstanceStorageVolumeConfigRequestWithDefaults instantiates a new UpdateInstanceStorageVolumeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceStorageVolumeConfigRequestWithDefaults() *UpdateInstanceStorageVolumeConfigRequest {
	this := UpdateInstanceStorageVolumeConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateInstanceStorageVolumeConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateInstanceStorageVolumeConfigRequest) SetId(v string) {
	o.Id = v
}

// GetInstanceStorageIops returns the InstanceStorageIops field value if set, zero value otherwise.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageIops() string {
	if o == nil || IsNil(o.InstanceStorageIops) {
		var ret string
		return ret
	}
	return *o.InstanceStorageIops
}

// GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageIopsOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageIops) {
		return nil, false
	}
	return o.InstanceStorageIops, true
}

// SetInstanceStorageIops gets a reference to the given string and assigns it to the InstanceStorageIops field.
func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageIops(v string) {
	o.InstanceStorageIops = &v
}

// GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field value if set, zero value otherwise.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBps() string {
	if o == nil || IsNil(o.InstanceStorageThroughputMiBps) {
		var ret string
		return ret
	}
	return *o.InstanceStorageThroughputMiBps
}

// GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageThroughputMiBpsOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageThroughputMiBps) {
		return nil, false
	}
	return o.InstanceStorageThroughputMiBps, true
}

// SetInstanceStorageThroughputMiBps gets a reference to the given string and assigns it to the InstanceStorageThroughputMiBps field.
func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageThroughputMiBps(v string) {
	o.InstanceStorageThroughputMiBps = &v
}

// GetInstanceStorageType returns the InstanceStorageType field value if set, zero value otherwise.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageType() string {
	if o == nil || IsNil(o.InstanceStorageType) {
		var ret string
		return ret
	}
	return *o.InstanceStorageType
}

// GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetInstanceStorageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageType) {
		return nil, false
	}
	return o.InstanceStorageType, true
}

// SetInstanceStorageType gets a reference to the given string and assigns it to the InstanceStorageType field.
func (o *UpdateInstanceStorageVolumeConfigRequest) SetInstanceStorageType(v string) {
	o.InstanceStorageType = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateInstanceStorageVolumeConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateInstanceStorageVolumeConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *UpdateInstanceStorageVolumeConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateInstanceStorageVolumeConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateInstanceStorageVolumeConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateInstanceStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstanceStorageVolumeConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.InstanceStorageIops) {
		toSerialize["instanceStorageIops"] = o.InstanceStorageIops
	}
	if !IsNil(o.InstanceStorageThroughputMiBps) {
		toSerialize["instanceStorageThroughputMiBps"] = o.InstanceStorageThroughputMiBps
	}
	if !IsNil(o.InstanceStorageType) {
		toSerialize["instanceStorageType"] = o.InstanceStorageType
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateInstanceStorageVolumeConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateInstanceStorageVolumeConfigRequest := _UpdateInstanceStorageVolumeConfigRequest{}

	err = json.Unmarshal(data, &varUpdateInstanceStorageVolumeConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdateInstanceStorageVolumeConfigRequest(varUpdateInstanceStorageVolumeConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "instanceStorageIops")
		delete(additionalProperties, "instanceStorageThroughputMiBps")
		delete(additionalProperties, "instanceStorageType")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateInstanceStorageVolumeConfigRequest struct {
	value *UpdateInstanceStorageVolumeConfigRequest
	isSet bool
}

func (v NullableUpdateInstanceStorageVolumeConfigRequest) Get() *UpdateInstanceStorageVolumeConfigRequest {
	return v.value
}

func (v *NullableUpdateInstanceStorageVolumeConfigRequest) Set(val *UpdateInstanceStorageVolumeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInstanceStorageVolumeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInstanceStorageVolumeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInstanceStorageVolumeConfigRequest(val *UpdateInstanceStorageVolumeConfigRequest) *NullableUpdateInstanceStorageVolumeConfigRequest {
	return &NullableUpdateInstanceStorageVolumeConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateInstanceStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInstanceStorageVolumeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


