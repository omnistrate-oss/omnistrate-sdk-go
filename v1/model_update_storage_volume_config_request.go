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

// checks if the UpdateStorageVolumeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageVolumeConfigRequest{}

// UpdateStorageVolumeConfigRequest struct for UpdateStorageVolumeConfigRequest
type UpdateStorageVolumeConfigRequest struct {
	// A brief description of the context for the storage volume pool
	Description *string `json:"description,omitempty"`
	// Disable backup for the storage volume
	DisableBackup *bool `json:"disableBackup,omitempty"`
	// ID of a Storage Volume Config
	Id string `json:"id"`
	// Name of the storage volume pool
	Name *string `json:"name,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorageVolumeConfigRequest UpdateStorageVolumeConfigRequest

// NewUpdateStorageVolumeConfigRequest instantiates a new UpdateStorageVolumeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageVolumeConfigRequest(id string, serviceId string, token string) *UpdateStorageVolumeConfigRequest {
	this := UpdateStorageVolumeConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateStorageVolumeConfigRequestWithDefaults instantiates a new UpdateStorageVolumeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageVolumeConfigRequestWithDefaults() *UpdateStorageVolumeConfigRequest {
	this := UpdateStorageVolumeConfigRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateStorageVolumeConfigRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateStorageVolumeConfigRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDisableBackup returns the DisableBackup field value if set, zero value otherwise.
func (o *UpdateStorageVolumeConfigRequest) GetDisableBackup() bool {
	if o == nil || IsNil(o.DisableBackup) {
		var ret bool
		return ret
	}
	return *o.DisableBackup
}

// GetDisableBackupOk returns a tuple with the DisableBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetDisableBackupOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableBackup) {
		return nil, false
	}
	return o.DisableBackup, true
}

// SetDisableBackup gets a reference to the given bool and assigns it to the DisableBackup field.
func (o *UpdateStorageVolumeConfigRequest) SetDisableBackup(v bool) {
	o.DisableBackup = &v
}

// GetId returns the Id field value
func (o *UpdateStorageVolumeConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateStorageVolumeConfigRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateStorageVolumeConfigRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateStorageVolumeConfigRequest) SetName(v string) {
	o.Name = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateStorageVolumeConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateStorageVolumeConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *UpdateStorageVolumeConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumeConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateStorageVolumeConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageVolumeConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisableBackup) {
		toSerialize["disableBackup"] = o.DisableBackup
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorageVolumeConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateStorageVolumeConfigRequest := _UpdateStorageVolumeConfigRequest{}

	err = json.Unmarshal(data, &varUpdateStorageVolumeConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdateStorageVolumeConfigRequest(varUpdateStorageVolumeConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "disableBackup")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorageVolumeConfigRequest struct {
	value *UpdateStorageVolumeConfigRequest
	isSet bool
}

func (v NullableUpdateStorageVolumeConfigRequest) Get() *UpdateStorageVolumeConfigRequest {
	return v.value
}

func (v *NullableUpdateStorageVolumeConfigRequest) Set(val *UpdateStorageVolumeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorageVolumeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorageVolumeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorageVolumeConfigRequest(val *UpdateStorageVolumeConfigRequest) *NullableUpdateStorageVolumeConfigRequest {
	return &NullableUpdateStorageVolumeConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorageVolumeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


