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

// checks if the RemoveStorageVolumeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveStorageVolumeConfigRequest{}

// RemoveStorageVolumeConfigRequest struct for RemoveStorageVolumeConfigRequest
type RemoveStorageVolumeConfigRequest struct {
	// ID of a Storage Config
	Id string `json:"id"`
	// The specific mount path to remove. If not specified, all mount paths for the storage volume config will be removed
	MountPath *string `json:"mountPath,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Storage Volume Config
	StorageVolumeConfigId string `json:"storageVolumeConfigId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RemoveStorageVolumeConfigRequest RemoveStorageVolumeConfigRequest

// NewRemoveStorageVolumeConfigRequest instantiates a new RemoveStorageVolumeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveStorageVolumeConfigRequest(id string, serviceId string, storageVolumeConfigId string, token string) *RemoveStorageVolumeConfigRequest {
	this := RemoveStorageVolumeConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.StorageVolumeConfigId = storageVolumeConfigId
	this.Token = token
	return &this
}

// NewRemoveStorageVolumeConfigRequestWithDefaults instantiates a new RemoveStorageVolumeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveStorageVolumeConfigRequestWithDefaults() *RemoveStorageVolumeConfigRequest {
	this := RemoveStorageVolumeConfigRequest{}
	return &this
}

// GetId returns the Id field value
func (o *RemoveStorageVolumeConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RemoveStorageVolumeConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RemoveStorageVolumeConfigRequest) SetId(v string) {
	o.Id = v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *RemoveStorageVolumeConfigRequest) GetMountPath() string {
	if o == nil || IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveStorageVolumeConfigRequest) GetMountPathOk() (*string, bool) {
	if o == nil || IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath gets a reference to the given string and assigns it to the MountPath field.
func (o *RemoveStorageVolumeConfigRequest) SetMountPath(v string) {
	o.MountPath = &v
}

// GetServiceId returns the ServiceId field value
func (o *RemoveStorageVolumeConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *RemoveStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *RemoveStorageVolumeConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetStorageVolumeConfigId returns the StorageVolumeConfigId field value
func (o *RemoveStorageVolumeConfigRequest) GetStorageVolumeConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageVolumeConfigId
}

// GetStorageVolumeConfigIdOk returns a tuple with the StorageVolumeConfigId field value
// and a boolean to check if the value has been set.
func (o *RemoveStorageVolumeConfigRequest) GetStorageVolumeConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageVolumeConfigId, true
}

// SetStorageVolumeConfigId sets field value
func (o *RemoveStorageVolumeConfigRequest) SetStorageVolumeConfigId(v string) {
	o.StorageVolumeConfigId = v
}

// GetToken returns the Token field value
func (o *RemoveStorageVolumeConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RemoveStorageVolumeConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RemoveStorageVolumeConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o RemoveStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveStorageVolumeConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["storageVolumeConfigId"] = o.StorageVolumeConfigId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveStorageVolumeConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serviceId",
		"storageVolumeConfigId",
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

	varRemoveStorageVolumeConfigRequest := _RemoveStorageVolumeConfigRequest{}

	err = json.Unmarshal(data, &varRemoveStorageVolumeConfigRequest)

	if err != nil {
		return err
	}

	*o = RemoveStorageVolumeConfigRequest(varRemoveStorageVolumeConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mountPath")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "storageVolumeConfigId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveStorageVolumeConfigRequest struct {
	value *RemoveStorageVolumeConfigRequest
	isSet bool
}

func (v NullableRemoveStorageVolumeConfigRequest) Get() *RemoveStorageVolumeConfigRequest {
	return v.value
}

func (v *NullableRemoveStorageVolumeConfigRequest) Set(val *RemoveStorageVolumeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveStorageVolumeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveStorageVolumeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveStorageVolumeConfigRequest(val *RemoveStorageVolumeConfigRequest) *NullableRemoveStorageVolumeConfigRequest {
	return &NullableRemoveStorageVolumeConfigRequest{value: val, isSet: true}
}

func (v NullableRemoveStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveStorageVolumeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


