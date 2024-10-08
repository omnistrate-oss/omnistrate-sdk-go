/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomDNSEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDNSEndpoint{}

// CustomDNSEndpoint struct for CustomDNSEndpoint
type CustomDNSEndpoint struct {
	ARecordTarget *string `json:"aRecordTarget,omitempty"`
	CnameTarget *string `json:"cnameTarget,omitempty"`
	DnsName *string `json:"dnsName,omitempty"`
	Enabled bool `json:"enabled"`
	Status *string `json:"status,omitempty"`
}

type _CustomDNSEndpoint CustomDNSEndpoint

// NewCustomDNSEndpoint instantiates a new CustomDNSEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDNSEndpoint(enabled bool) *CustomDNSEndpoint {
	this := CustomDNSEndpoint{}
	this.Enabled = enabled
	return &this
}

// NewCustomDNSEndpointWithDefaults instantiates a new CustomDNSEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDNSEndpointWithDefaults() *CustomDNSEndpoint {
	this := CustomDNSEndpoint{}
	return &this
}

// GetARecordTarget returns the ARecordTarget field value if set, zero value otherwise.
func (o *CustomDNSEndpoint) GetARecordTarget() string {
	if o == nil || IsNil(o.ARecordTarget) {
		var ret string
		return ret
	}
	return *o.ARecordTarget
}

// GetARecordTargetOk returns a tuple with the ARecordTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDNSEndpoint) GetARecordTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ARecordTarget) {
		return nil, false
	}
	return o.ARecordTarget, true
}

// HasARecordTarget returns a boolean if a field has been set.
func (o *CustomDNSEndpoint) HasARecordTarget() bool {
	if o != nil && !IsNil(o.ARecordTarget) {
		return true
	}

	return false
}

// SetARecordTarget gets a reference to the given string and assigns it to the ARecordTarget field.
func (o *CustomDNSEndpoint) SetARecordTarget(v string) {
	o.ARecordTarget = &v
}

// GetCnameTarget returns the CnameTarget field value if set, zero value otherwise.
func (o *CustomDNSEndpoint) GetCnameTarget() string {
	if o == nil || IsNil(o.CnameTarget) {
		var ret string
		return ret
	}
	return *o.CnameTarget
}

// GetCnameTargetOk returns a tuple with the CnameTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDNSEndpoint) GetCnameTargetOk() (*string, bool) {
	if o == nil || IsNil(o.CnameTarget) {
		return nil, false
	}
	return o.CnameTarget, true
}

// HasCnameTarget returns a boolean if a field has been set.
func (o *CustomDNSEndpoint) HasCnameTarget() bool {
	if o != nil && !IsNil(o.CnameTarget) {
		return true
	}

	return false
}

// SetCnameTarget gets a reference to the given string and assigns it to the CnameTarget field.
func (o *CustomDNSEndpoint) SetCnameTarget(v string) {
	o.CnameTarget = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *CustomDNSEndpoint) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDNSEndpoint) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *CustomDNSEndpoint) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *CustomDNSEndpoint) SetDnsName(v string) {
	o.DnsName = &v
}

// GetEnabled returns the Enabled field value
func (o *CustomDNSEndpoint) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CustomDNSEndpoint) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CustomDNSEndpoint) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomDNSEndpoint) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDNSEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomDNSEndpoint) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomDNSEndpoint) SetStatus(v string) {
	o.Status = &v
}

func (o CustomDNSEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDNSEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ARecordTarget) {
		toSerialize["aRecordTarget"] = o.ARecordTarget
	}
	if !IsNil(o.CnameTarget) {
		toSerialize["cnameTarget"] = o.CnameTarget
	}
	if !IsNil(o.DnsName) {
		toSerialize["dnsName"] = o.DnsName
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *CustomDNSEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varCustomDNSEndpoint := _CustomDNSEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomDNSEndpoint)

	if err != nil {
		return err
	}

	*o = CustomDNSEndpoint(varCustomDNSEndpoint)

	return err
}

type NullableCustomDNSEndpoint struct {
	value *CustomDNSEndpoint
	isSet bool
}

func (v NullableCustomDNSEndpoint) Get() *CustomDNSEndpoint {
	return v.value
}

func (v *NullableCustomDNSEndpoint) Set(val *CustomDNSEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDNSEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDNSEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDNSEndpoint(val *CustomDNSEndpoint) *NullableCustomDNSEndpoint {
	return &NullableCustomDNSEndpoint{value: val, isSet: true}
}

func (v NullableCustomDNSEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDNSEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


