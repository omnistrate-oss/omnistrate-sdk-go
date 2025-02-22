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

// checks if the EFSFileSystemConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EFSFileSystemConfiguration{}

// EFSFileSystemConfiguration struct for EFSFileSystemConfiguration
type EFSFileSystemConfiguration struct {
	// The performance mode of the EFS file system
	PerformanceMode string `json:"PerformanceMode"`
	// The throughput, measured in MiBps, that you want to provision for the EFS file system, if throughput mode if provisioned
	ProvisionedThroughputInMibps *float64 `json:"ProvisionedThroughputInMibps,omitempty"`
	// The throughput mode of the EFS file system
	ThroughputMode string `json:"ThroughputMode"`
	AdditionalProperties map[string]interface{}
}

type _EFSFileSystemConfiguration EFSFileSystemConfiguration

// NewEFSFileSystemConfiguration instantiates a new EFSFileSystemConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEFSFileSystemConfiguration(performanceMode string, throughputMode string) *EFSFileSystemConfiguration {
	this := EFSFileSystemConfiguration{}
	this.PerformanceMode = performanceMode
	this.ThroughputMode = throughputMode
	return &this
}

// NewEFSFileSystemConfigurationWithDefaults instantiates a new EFSFileSystemConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEFSFileSystemConfigurationWithDefaults() *EFSFileSystemConfiguration {
	this := EFSFileSystemConfiguration{}
	return &this
}

// GetPerformanceMode returns the PerformanceMode field value
func (o *EFSFileSystemConfiguration) GetPerformanceMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PerformanceMode
}

// GetPerformanceModeOk returns a tuple with the PerformanceMode field value
// and a boolean to check if the value has been set.
func (o *EFSFileSystemConfiguration) GetPerformanceModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerformanceMode, true
}

// SetPerformanceMode sets field value
func (o *EFSFileSystemConfiguration) SetPerformanceMode(v string) {
	o.PerformanceMode = v
}

// GetProvisionedThroughputInMibps returns the ProvisionedThroughputInMibps field value if set, zero value otherwise.
func (o *EFSFileSystemConfiguration) GetProvisionedThroughputInMibps() float64 {
	if o == nil || IsNil(o.ProvisionedThroughputInMibps) {
		var ret float64
		return ret
	}
	return *o.ProvisionedThroughputInMibps
}

// GetProvisionedThroughputInMibpsOk returns a tuple with the ProvisionedThroughputInMibps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EFSFileSystemConfiguration) GetProvisionedThroughputInMibpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionedThroughputInMibps) {
		return nil, false
	}
	return o.ProvisionedThroughputInMibps, true
}

// SetProvisionedThroughputInMibps gets a reference to the given float64 and assigns it to the ProvisionedThroughputInMibps field.
func (o *EFSFileSystemConfiguration) SetProvisionedThroughputInMibps(v float64) {
	o.ProvisionedThroughputInMibps = &v
}

// GetThroughputMode returns the ThroughputMode field value
func (o *EFSFileSystemConfiguration) GetThroughputMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThroughputMode
}

// GetThroughputModeOk returns a tuple with the ThroughputMode field value
// and a boolean to check if the value has been set.
func (o *EFSFileSystemConfiguration) GetThroughputModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThroughputMode, true
}

// SetThroughputMode sets field value
func (o *EFSFileSystemConfiguration) SetThroughputMode(v string) {
	o.ThroughputMode = v
}

func (o EFSFileSystemConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EFSFileSystemConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PerformanceMode"] = o.PerformanceMode
	if !IsNil(o.ProvisionedThroughputInMibps) {
		toSerialize["ProvisionedThroughputInMibps"] = o.ProvisionedThroughputInMibps
	}
	toSerialize["ThroughputMode"] = o.ThroughputMode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EFSFileSystemConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PerformanceMode",
		"ThroughputMode",
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

	varEFSFileSystemConfiguration := _EFSFileSystemConfiguration{}

	err = json.Unmarshal(data, &varEFSFileSystemConfiguration)

	if err != nil {
		return err
	}

	*o = EFSFileSystemConfiguration(varEFSFileSystemConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "PerformanceMode")
		delete(additionalProperties, "ProvisionedThroughputInMibps")
		delete(additionalProperties, "ThroughputMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEFSFileSystemConfiguration struct {
	value *EFSFileSystemConfiguration
	isSet bool
}

func (v NullableEFSFileSystemConfiguration) Get() *EFSFileSystemConfiguration {
	return v.value
}

func (v *NullableEFSFileSystemConfiguration) Set(val *EFSFileSystemConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableEFSFileSystemConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableEFSFileSystemConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEFSFileSystemConfiguration(val *EFSFileSystemConfiguration) *NullableEFSFileSystemConfiguration {
	return &NullableEFSFileSystemConfiguration{value: val, isSet: true}
}

func (v NullableEFSFileSystemConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEFSFileSystemConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


