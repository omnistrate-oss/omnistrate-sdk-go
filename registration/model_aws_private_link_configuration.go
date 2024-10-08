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

// checks if the AWSPrivateLinkConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSPrivateLinkConfiguration{}

// AWSPrivateLinkConfiguration struct for AWSPrivateLinkConfiguration
type AWSPrivateLinkConfiguration struct {
	// The port of the target group
	Port int64 `json:"port"`
	// The target group name
	TargetGroupName string `json:"targetGroupName"`
}

type _AWSPrivateLinkConfiguration AWSPrivateLinkConfiguration

// NewAWSPrivateLinkConfiguration instantiates a new AWSPrivateLinkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSPrivateLinkConfiguration(port int64, targetGroupName string) *AWSPrivateLinkConfiguration {
	this := AWSPrivateLinkConfiguration{}
	this.Port = port
	this.TargetGroupName = targetGroupName
	return &this
}

// NewAWSPrivateLinkConfigurationWithDefaults instantiates a new AWSPrivateLinkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSPrivateLinkConfigurationWithDefaults() *AWSPrivateLinkConfiguration {
	this := AWSPrivateLinkConfiguration{}
	return &this
}

// GetPort returns the Port field value
func (o *AWSPrivateLinkConfiguration) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConfiguration) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *AWSPrivateLinkConfiguration) SetPort(v int64) {
	o.Port = v
}

// GetTargetGroupName returns the TargetGroupName field value
func (o *AWSPrivateLinkConfiguration) GetTargetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroupName
}

// GetTargetGroupNameOk returns a tuple with the TargetGroupName field value
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConfiguration) GetTargetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetGroupName, true
}

// SetTargetGroupName sets field value
func (o *AWSPrivateLinkConfiguration) SetTargetGroupName(v string) {
	o.TargetGroupName = v
}

func (o AWSPrivateLinkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSPrivateLinkConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["targetGroupName"] = o.TargetGroupName
	return toSerialize, nil
}

func (o *AWSPrivateLinkConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"targetGroupName",
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

	varAWSPrivateLinkConfiguration := _AWSPrivateLinkConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAWSPrivateLinkConfiguration)

	if err != nil {
		return err
	}

	*o = AWSPrivateLinkConfiguration(varAWSPrivateLinkConfiguration)

	return err
}

type NullableAWSPrivateLinkConfiguration struct {
	value *AWSPrivateLinkConfiguration
	isSet bool
}

func (v NullableAWSPrivateLinkConfiguration) Get() *AWSPrivateLinkConfiguration {
	return v.value
}

func (v *NullableAWSPrivateLinkConfiguration) Set(val *AWSPrivateLinkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSPrivateLinkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSPrivateLinkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSPrivateLinkConfiguration(val *AWSPrivateLinkConfiguration) *NullableAWSPrivateLinkConfiguration {
	return &NullableAWSPrivateLinkConfiguration{value: val, isSet: true}
}

func (v NullableAWSPrivateLinkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSPrivateLinkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


