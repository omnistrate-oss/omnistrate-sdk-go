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

// checks if the OmnistrateServiceVersionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmnistrateServiceVersionResult{}

// OmnistrateServiceVersionResult struct for OmnistrateServiceVersionResult
type OmnistrateServiceVersionResult struct {
	// Version of the Entity to operate on
	ApiVersion string `json:"apiVersion"`
	// The commit hash of the build
	BuildCommitSHA string `json:"buildCommitSHA"`
	// The timestamp of the build
	BuildTimestamp string `json:"buildTimestamp"`
	AdditionalProperties map[string]interface{}
}

type _OmnistrateServiceVersionResult OmnistrateServiceVersionResult

// NewOmnistrateServiceVersionResult instantiates a new OmnistrateServiceVersionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmnistrateServiceVersionResult(apiVersion string, buildCommitSHA string, buildTimestamp string) *OmnistrateServiceVersionResult {
	this := OmnistrateServiceVersionResult{}
	this.ApiVersion = apiVersion
	this.BuildCommitSHA = buildCommitSHA
	this.BuildTimestamp = buildTimestamp
	return &this
}

// NewOmnistrateServiceVersionResultWithDefaults instantiates a new OmnistrateServiceVersionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmnistrateServiceVersionResultWithDefaults() *OmnistrateServiceVersionResult {
	this := OmnistrateServiceVersionResult{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *OmnistrateServiceVersionResult) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *OmnistrateServiceVersionResult) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *OmnistrateServiceVersionResult) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetBuildCommitSHA returns the BuildCommitSHA field value
func (o *OmnistrateServiceVersionResult) GetBuildCommitSHA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildCommitSHA
}

// GetBuildCommitSHAOk returns a tuple with the BuildCommitSHA field value
// and a boolean to check if the value has been set.
func (o *OmnistrateServiceVersionResult) GetBuildCommitSHAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildCommitSHA, true
}

// SetBuildCommitSHA sets field value
func (o *OmnistrateServiceVersionResult) SetBuildCommitSHA(v string) {
	o.BuildCommitSHA = v
}

// GetBuildTimestamp returns the BuildTimestamp field value
func (o *OmnistrateServiceVersionResult) GetBuildTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildTimestamp
}

// GetBuildTimestampOk returns a tuple with the BuildTimestamp field value
// and a boolean to check if the value has been set.
func (o *OmnistrateServiceVersionResult) GetBuildTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildTimestamp, true
}

// SetBuildTimestamp sets field value
func (o *OmnistrateServiceVersionResult) SetBuildTimestamp(v string) {
	o.BuildTimestamp = v
}

func (o OmnistrateServiceVersionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmnistrateServiceVersionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["buildCommitSHA"] = o.BuildCommitSHA
	toSerialize["buildTimestamp"] = o.BuildTimestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OmnistrateServiceVersionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"buildCommitSHA",
		"buildTimestamp",
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

	varOmnistrateServiceVersionResult := _OmnistrateServiceVersionResult{}

	err = json.Unmarshal(data, &varOmnistrateServiceVersionResult)

	if err != nil {
		return err
	}

	*o = OmnistrateServiceVersionResult(varOmnistrateServiceVersionResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiVersion")
		delete(additionalProperties, "buildCommitSHA")
		delete(additionalProperties, "buildTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOmnistrateServiceVersionResult struct {
	value *OmnistrateServiceVersionResult
	isSet bool
}

func (v NullableOmnistrateServiceVersionResult) Get() *OmnistrateServiceVersionResult {
	return v.value
}

func (v *NullableOmnistrateServiceVersionResult) Set(val *OmnistrateServiceVersionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOmnistrateServiceVersionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOmnistrateServiceVersionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmnistrateServiceVersionResult(val *OmnistrateServiceVersionResult) *NullableOmnistrateServiceVersionResult {
	return &NullableOmnistrateServiceVersionResult{value: val, isSet: true}
}

func (v NullableOmnistrateServiceVersionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmnistrateServiceVersionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


