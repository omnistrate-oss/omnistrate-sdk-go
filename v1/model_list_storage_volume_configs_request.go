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

// checks if the ListStorageVolumeConfigsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStorageVolumeConfigsRequest{}

// ListStorageVolumeConfigsRequest struct for ListStorageVolumeConfigsRequest
type ListStorageVolumeConfigsRequest struct {
	// Is storage volume config managed by omnistrate
	Managed *bool `json:"managed,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListStorageVolumeConfigsRequest ListStorageVolumeConfigsRequest

// NewListStorageVolumeConfigsRequest instantiates a new ListStorageVolumeConfigsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStorageVolumeConfigsRequest(serviceId string, token string) *ListStorageVolumeConfigsRequest {
	this := ListStorageVolumeConfigsRequest{}
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewListStorageVolumeConfigsRequestWithDefaults instantiates a new ListStorageVolumeConfigsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStorageVolumeConfigsRequestWithDefaults() *ListStorageVolumeConfigsRequest {
	this := ListStorageVolumeConfigsRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ListStorageVolumeConfigsRequest) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStorageVolumeConfigsRequest) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ListStorageVolumeConfigsRequest) SetManaged(v bool) {
	o.Managed = &v
}

// GetServiceId returns the ServiceId field value
func (o *ListStorageVolumeConfigsRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListStorageVolumeConfigsRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListStorageVolumeConfigsRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *ListStorageVolumeConfigsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListStorageVolumeConfigsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListStorageVolumeConfigsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListStorageVolumeConfigsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStorageVolumeConfigsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListStorageVolumeConfigsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varListStorageVolumeConfigsRequest := _ListStorageVolumeConfigsRequest{}

	err = json.Unmarshal(data, &varListStorageVolumeConfigsRequest)

	if err != nil {
		return err
	}

	*o = ListStorageVolumeConfigsRequest(varListStorageVolumeConfigsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managed")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListStorageVolumeConfigsRequest struct {
	value *ListStorageVolumeConfigsRequest
	isSet bool
}

func (v NullableListStorageVolumeConfigsRequest) Get() *ListStorageVolumeConfigsRequest {
	return v.value
}

func (v *NullableListStorageVolumeConfigsRequest) Set(val *ListStorageVolumeConfigsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListStorageVolumeConfigsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListStorageVolumeConfigsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStorageVolumeConfigsRequest(val *ListStorageVolumeConfigsRequest) *NullableListStorageVolumeConfigsRequest {
	return &NullableListStorageVolumeConfigsRequest{value: val, isSet: true}
}

func (v NullableListStorageVolumeConfigsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStorageVolumeConfigsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


