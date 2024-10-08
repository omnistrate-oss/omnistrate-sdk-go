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

// checks if the AddCustomDNSToResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCustomDNSToResourceInstanceRequestBody{}

// AddCustomDNSToResourceInstanceRequestBody struct for AddCustomDNSToResourceInstanceRequestBody
type AddCustomDNSToResourceInstanceRequestBody struct {
	// The custom DNS to add
	CustomDNS string `json:"customDNS"`
	// The target port
	TargetPort *int64 `json:"targetPort,omitempty"`
}

type _AddCustomDNSToResourceInstanceRequestBody AddCustomDNSToResourceInstanceRequestBody

// NewAddCustomDNSToResourceInstanceRequestBody instantiates a new AddCustomDNSToResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCustomDNSToResourceInstanceRequestBody(customDNS string) *AddCustomDNSToResourceInstanceRequestBody {
	this := AddCustomDNSToResourceInstanceRequestBody{}
	this.CustomDNS = customDNS
	return &this
}

// NewAddCustomDNSToResourceInstanceRequestBodyWithDefaults instantiates a new AddCustomDNSToResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCustomDNSToResourceInstanceRequestBodyWithDefaults() *AddCustomDNSToResourceInstanceRequestBody {
	this := AddCustomDNSToResourceInstanceRequestBody{}
	return &this
}

// GetCustomDNS returns the CustomDNS field value
func (o *AddCustomDNSToResourceInstanceRequestBody) GetCustomDNS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDNS
}

// GetCustomDNSOk returns a tuple with the CustomDNS field value
// and a boolean to check if the value has been set.
func (o *AddCustomDNSToResourceInstanceRequestBody) GetCustomDNSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDNS, true
}

// SetCustomDNS sets field value
func (o *AddCustomDNSToResourceInstanceRequestBody) SetCustomDNS(v string) {
	o.CustomDNS = v
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *AddCustomDNSToResourceInstanceRequestBody) GetTargetPort() int64 {
	if o == nil || IsNil(o.TargetPort) {
		var ret int64
		return ret
	}
	return *o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomDNSToResourceInstanceRequestBody) GetTargetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetPort) {
		return nil, false
	}
	return o.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *AddCustomDNSToResourceInstanceRequestBody) HasTargetPort() bool {
	if o != nil && !IsNil(o.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given int64 and assigns it to the TargetPort field.
func (o *AddCustomDNSToResourceInstanceRequestBody) SetTargetPort(v int64) {
	o.TargetPort = &v
}

func (o AddCustomDNSToResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCustomDNSToResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customDNS"] = o.CustomDNS
	if !IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}
	return toSerialize, nil
}

func (o *AddCustomDNSToResourceInstanceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customDNS",
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

	varAddCustomDNSToResourceInstanceRequestBody := _AddCustomDNSToResourceInstanceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddCustomDNSToResourceInstanceRequestBody)

	if err != nil {
		return err
	}

	*o = AddCustomDNSToResourceInstanceRequestBody(varAddCustomDNSToResourceInstanceRequestBody)

	return err
}

type NullableAddCustomDNSToResourceInstanceRequestBody struct {
	value *AddCustomDNSToResourceInstanceRequestBody
	isSet bool
}

func (v NullableAddCustomDNSToResourceInstanceRequestBody) Get() *AddCustomDNSToResourceInstanceRequestBody {
	return v.value
}

func (v *NullableAddCustomDNSToResourceInstanceRequestBody) Set(val *AddCustomDNSToResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCustomDNSToResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCustomDNSToResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCustomDNSToResourceInstanceRequestBody(val *AddCustomDNSToResourceInstanceRequestBody) *NullableAddCustomDNSToResourceInstanceRequestBody {
	return &NullableAddCustomDNSToResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableAddCustomDNSToResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCustomDNSToResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


