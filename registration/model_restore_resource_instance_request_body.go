/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RestoreResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreResourceInstanceRequestBody{}

// RestoreResourceInstanceRequestBody struct for RestoreResourceInstanceRequestBody
type RestoreResourceInstanceRequestBody struct {
	// The network type
	NetworkType *string `json:"network_type,omitempty"`
	// The target restore time
	TargetRestoreTime string `json:"targetRestoreTime"`
}

type _RestoreResourceInstanceRequestBody RestoreResourceInstanceRequestBody

// NewRestoreResourceInstanceRequestBody instantiates a new RestoreResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreResourceInstanceRequestBody(targetRestoreTime string) *RestoreResourceInstanceRequestBody {
	this := RestoreResourceInstanceRequestBody{}
	this.TargetRestoreTime = targetRestoreTime
	return &this
}

// NewRestoreResourceInstanceRequestBodyWithDefaults instantiates a new RestoreResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreResourceInstanceRequestBodyWithDefaults() *RestoreResourceInstanceRequestBody {
	this := RestoreResourceInstanceRequestBody{}
	return &this
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *RestoreResourceInstanceRequestBody) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceRequestBody) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *RestoreResourceInstanceRequestBody) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetTargetRestoreTime returns the TargetRestoreTime field value
func (o *RestoreResourceInstanceRequestBody) GetTargetRestoreTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRestoreTime
}

// GetTargetRestoreTimeOk returns a tuple with the TargetRestoreTime field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceRequestBody) GetTargetRestoreTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetRestoreTime, true
}

// SetTargetRestoreTime sets field value
func (o *RestoreResourceInstanceRequestBody) SetTargetRestoreTime(v string) {
	o.TargetRestoreTime = v
}

func (o RestoreResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}
	toSerialize["targetRestoreTime"] = o.TargetRestoreTime
	return toSerialize, nil
}

func (o *RestoreResourceInstanceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetRestoreTime",
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

	varRestoreResourceInstanceRequestBody := _RestoreResourceInstanceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestoreResourceInstanceRequestBody)

	if err != nil {
		return err
	}

	*o = RestoreResourceInstanceRequestBody(varRestoreResourceInstanceRequestBody)

	return err
}

type NullableRestoreResourceInstanceRequestBody struct {
	value *RestoreResourceInstanceRequestBody
	isSet bool
}

func (v NullableRestoreResourceInstanceRequestBody) Get() *RestoreResourceInstanceRequestBody {
	return v.value
}

func (v *NullableRestoreResourceInstanceRequestBody) Set(val *RestoreResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreResourceInstanceRequestBody(val *RestoreResourceInstanceRequestBody) *NullableRestoreResourceInstanceRequestBody {
	return &NullableRestoreResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableRestoreResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


