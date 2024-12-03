/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the RestoreResourceInstanceFromSnapshotRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreResourceInstanceFromSnapshotRequestBody{}

// RestoreResourceInstanceFromSnapshotRequestBody struct for RestoreResourceInstanceFromSnapshotRequestBody
type RestoreResourceInstanceFromSnapshotRequestBody struct {
	// The network type
	NetworkType *string `json:"network_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RestoreResourceInstanceFromSnapshotRequestBody RestoreResourceInstanceFromSnapshotRequestBody

// NewRestoreResourceInstanceFromSnapshotRequestBody instantiates a new RestoreResourceInstanceFromSnapshotRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreResourceInstanceFromSnapshotRequestBody() *RestoreResourceInstanceFromSnapshotRequestBody {
	this := RestoreResourceInstanceFromSnapshotRequestBody{}
	return &this
}

// NewRestoreResourceInstanceFromSnapshotRequestBodyWithDefaults instantiates a new RestoreResourceInstanceFromSnapshotRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreResourceInstanceFromSnapshotRequestBodyWithDefaults() *RestoreResourceInstanceFromSnapshotRequestBody {
	this := RestoreResourceInstanceFromSnapshotRequestBody{}
	return &this
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *RestoreResourceInstanceFromSnapshotRequestBody) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequestBody) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *RestoreResourceInstanceFromSnapshotRequestBody) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *RestoreResourceInstanceFromSnapshotRequestBody) SetNetworkType(v string) {
	o.NetworkType = &v
}

func (o RestoreResourceInstanceFromSnapshotRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreResourceInstanceFromSnapshotRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RestoreResourceInstanceFromSnapshotRequestBody) UnmarshalJSON(data []byte) (err error) {
	varRestoreResourceInstanceFromSnapshotRequestBody := _RestoreResourceInstanceFromSnapshotRequestBody{}

	err = json.Unmarshal(data, &varRestoreResourceInstanceFromSnapshotRequestBody)

	if err != nil {
		return err
	}

	*o = RestoreResourceInstanceFromSnapshotRequestBody(varRestoreResourceInstanceFromSnapshotRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRestoreResourceInstanceFromSnapshotRequestBody struct {
	value *RestoreResourceInstanceFromSnapshotRequestBody
	isSet bool
}

func (v NullableRestoreResourceInstanceFromSnapshotRequestBody) Get() *RestoreResourceInstanceFromSnapshotRequestBody {
	return v.value
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequestBody) Set(val *RestoreResourceInstanceFromSnapshotRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreResourceInstanceFromSnapshotRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreResourceInstanceFromSnapshotRequestBody(val *RestoreResourceInstanceFromSnapshotRequestBody) *NullableRestoreResourceInstanceFromSnapshotRequestBody {
	return &NullableRestoreResourceInstanceFromSnapshotRequestBody{value: val, isSet: true}
}

func (v NullableRestoreResourceInstanceFromSnapshotRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


