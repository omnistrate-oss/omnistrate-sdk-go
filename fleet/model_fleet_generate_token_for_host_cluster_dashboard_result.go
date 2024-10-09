/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FleetGenerateTokenForHostClusterDashboardResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetGenerateTokenForHostClusterDashboardResult{}

// FleetGenerateTokenForHostClusterDashboardResult struct for FleetGenerateTokenForHostClusterDashboardResult
type FleetGenerateTokenForHostClusterDashboardResult struct {
	// The timestamp when the token expires
	ExpirationTimestamp string `json:"expirationTimestamp"`
	// The token to access the dashboard
	Token string `json:"token"`
}

type _FleetGenerateTokenForHostClusterDashboardResult FleetGenerateTokenForHostClusterDashboardResult

// NewFleetGenerateTokenForHostClusterDashboardResult instantiates a new FleetGenerateTokenForHostClusterDashboardResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetGenerateTokenForHostClusterDashboardResult(expirationTimestamp string, token string) *FleetGenerateTokenForHostClusterDashboardResult {
	this := FleetGenerateTokenForHostClusterDashboardResult{}
	this.ExpirationTimestamp = expirationTimestamp
	this.Token = token
	return &this
}

// NewFleetGenerateTokenForHostClusterDashboardResultWithDefaults instantiates a new FleetGenerateTokenForHostClusterDashboardResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetGenerateTokenForHostClusterDashboardResultWithDefaults() *FleetGenerateTokenForHostClusterDashboardResult {
	this := FleetGenerateTokenForHostClusterDashboardResult{}
	return &this
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value
func (o *FleetGenerateTokenForHostClusterDashboardResult) GetExpirationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value
// and a boolean to check if the value has been set.
func (o *FleetGenerateTokenForHostClusterDashboardResult) GetExpirationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestamp, true
}

// SetExpirationTimestamp sets field value
func (o *FleetGenerateTokenForHostClusterDashboardResult) SetExpirationTimestamp(v string) {
	o.ExpirationTimestamp = v
}

// GetToken returns the Token field value
func (o *FleetGenerateTokenForHostClusterDashboardResult) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FleetGenerateTokenForHostClusterDashboardResult) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FleetGenerateTokenForHostClusterDashboardResult) SetToken(v string) {
	o.Token = v
}

func (o FleetGenerateTokenForHostClusterDashboardResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetGenerateTokenForHostClusterDashboardResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expirationTimestamp"] = o.ExpirationTimestamp
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *FleetGenerateTokenForHostClusterDashboardResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expirationTimestamp",
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

	varFleetGenerateTokenForHostClusterDashboardResult := _FleetGenerateTokenForHostClusterDashboardResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFleetGenerateTokenForHostClusterDashboardResult)

	if err != nil {
		return err
	}

	*o = FleetGenerateTokenForHostClusterDashboardResult(varFleetGenerateTokenForHostClusterDashboardResult)

	return err
}

type NullableFleetGenerateTokenForHostClusterDashboardResult struct {
	value *FleetGenerateTokenForHostClusterDashboardResult
	isSet bool
}

func (v NullableFleetGenerateTokenForHostClusterDashboardResult) Get() *FleetGenerateTokenForHostClusterDashboardResult {
	return v.value
}

func (v *NullableFleetGenerateTokenForHostClusterDashboardResult) Set(val *FleetGenerateTokenForHostClusterDashboardResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetGenerateTokenForHostClusterDashboardResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetGenerateTokenForHostClusterDashboardResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetGenerateTokenForHostClusterDashboardResult(val *FleetGenerateTokenForHostClusterDashboardResult) *NullableFleetGenerateTokenForHostClusterDashboardResult {
	return &NullableFleetGenerateTokenForHostClusterDashboardResult{value: val, isSet: true}
}

func (v NullableFleetGenerateTokenForHostClusterDashboardResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetGenerateTokenForHostClusterDashboardResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


