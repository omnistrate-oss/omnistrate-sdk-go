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

// checks if the DetailedNodeHealthResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedNodeHealthResult{}

// DetailedNodeHealthResult struct for DetailedNodeHealthResult
type DetailedNodeHealthResult struct {
	// The health status of the network endpoints
	ConnectivityStatus string `json:"ConnectivityStatus"`
	// The health status of the disk
	DiskHealth string `json:"DiskHealth"`
	// The load status of the pod
	LoadHealth string `json:"LoadHealth"`
	// The health status of the machine hosting the service
	NodeHealth string `json:"NodeHealth"`
	// The health status of the process
	ProcessHealth string `json:"ProcessHealth"`
	// The liveness status of the process
	ProcessLiveness string `json:"ProcessLiveness"`
	AdditionalProperties map[string]interface{}
}

type _DetailedNodeHealthResult DetailedNodeHealthResult

// NewDetailedNodeHealthResult instantiates a new DetailedNodeHealthResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedNodeHealthResult(connectivityStatus string, diskHealth string, loadHealth string, nodeHealth string, processHealth string, processLiveness string) *DetailedNodeHealthResult {
	this := DetailedNodeHealthResult{}
	this.ConnectivityStatus = connectivityStatus
	this.DiskHealth = diskHealth
	this.LoadHealth = loadHealth
	this.NodeHealth = nodeHealth
	this.ProcessHealth = processHealth
	this.ProcessLiveness = processLiveness
	return &this
}

// NewDetailedNodeHealthResultWithDefaults instantiates a new DetailedNodeHealthResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedNodeHealthResultWithDefaults() *DetailedNodeHealthResult {
	this := DetailedNodeHealthResult{}
	return &this
}

// GetConnectivityStatus returns the ConnectivityStatus field value
func (o *DetailedNodeHealthResult) GetConnectivityStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectivityStatus
}

// GetConnectivityStatusOk returns a tuple with the ConnectivityStatus field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetConnectivityStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectivityStatus, true
}

// SetConnectivityStatus sets field value
func (o *DetailedNodeHealthResult) SetConnectivityStatus(v string) {
	o.ConnectivityStatus = v
}

// GetDiskHealth returns the DiskHealth field value
func (o *DetailedNodeHealthResult) GetDiskHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskHealth
}

// GetDiskHealthOk returns a tuple with the DiskHealth field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetDiskHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskHealth, true
}

// SetDiskHealth sets field value
func (o *DetailedNodeHealthResult) SetDiskHealth(v string) {
	o.DiskHealth = v
}

// GetLoadHealth returns the LoadHealth field value
func (o *DetailedNodeHealthResult) GetLoadHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadHealth
}

// GetLoadHealthOk returns a tuple with the LoadHealth field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetLoadHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadHealth, true
}

// SetLoadHealth sets field value
func (o *DetailedNodeHealthResult) SetLoadHealth(v string) {
	o.LoadHealth = v
}

// GetNodeHealth returns the NodeHealth field value
func (o *DetailedNodeHealthResult) GetNodeHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeHealth
}

// GetNodeHealthOk returns a tuple with the NodeHealth field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetNodeHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeHealth, true
}

// SetNodeHealth sets field value
func (o *DetailedNodeHealthResult) SetNodeHealth(v string) {
	o.NodeHealth = v
}

// GetProcessHealth returns the ProcessHealth field value
func (o *DetailedNodeHealthResult) GetProcessHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessHealth
}

// GetProcessHealthOk returns a tuple with the ProcessHealth field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetProcessHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessHealth, true
}

// SetProcessHealth sets field value
func (o *DetailedNodeHealthResult) SetProcessHealth(v string) {
	o.ProcessHealth = v
}

// GetProcessLiveness returns the ProcessLiveness field value
func (o *DetailedNodeHealthResult) GetProcessLiveness() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessLiveness
}

// GetProcessLivenessOk returns a tuple with the ProcessLiveness field value
// and a boolean to check if the value has been set.
func (o *DetailedNodeHealthResult) GetProcessLivenessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessLiveness, true
}

// SetProcessLiveness sets field value
func (o *DetailedNodeHealthResult) SetProcessLiveness(v string) {
	o.ProcessLiveness = v
}

func (o DetailedNodeHealthResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedNodeHealthResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ConnectivityStatus"] = o.ConnectivityStatus
	toSerialize["DiskHealth"] = o.DiskHealth
	toSerialize["LoadHealth"] = o.LoadHealth
	toSerialize["NodeHealth"] = o.NodeHealth
	toSerialize["ProcessHealth"] = o.ProcessHealth
	toSerialize["ProcessLiveness"] = o.ProcessLiveness

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetailedNodeHealthResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ConnectivityStatus",
		"DiskHealth",
		"LoadHealth",
		"NodeHealth",
		"ProcessHealth",
		"ProcessLiveness",
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

	varDetailedNodeHealthResult := _DetailedNodeHealthResult{}

	err = json.Unmarshal(data, &varDetailedNodeHealthResult)

	if err != nil {
		return err
	}

	*o = DetailedNodeHealthResult(varDetailedNodeHealthResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ConnectivityStatus")
		delete(additionalProperties, "DiskHealth")
		delete(additionalProperties, "LoadHealth")
		delete(additionalProperties, "NodeHealth")
		delete(additionalProperties, "ProcessHealth")
		delete(additionalProperties, "ProcessLiveness")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetailedNodeHealthResult struct {
	value *DetailedNodeHealthResult
	isSet bool
}

func (v NullableDetailedNodeHealthResult) Get() *DetailedNodeHealthResult {
	return v.value
}

func (v *NullableDetailedNodeHealthResult) Set(val *DetailedNodeHealthResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedNodeHealthResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedNodeHealthResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedNodeHealthResult(val *DetailedNodeHealthResult) *NullableDetailedNodeHealthResult {
	return &NullableDetailedNodeHealthResult{value: val, isSet: true}
}

func (v NullableDetailedNodeHealthResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedNodeHealthResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


