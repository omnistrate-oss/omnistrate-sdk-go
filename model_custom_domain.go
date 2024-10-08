/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDomain{}

// CustomDomain struct for CustomDomain
type CustomDomain struct {
	// The cluster endpoint for the saas portal instance
	ClusterEndpoint string `json:"clusterEndpoint"`
	// The custom domain
	CustomDomain string `json:"customDomain"`
	// The custom domain description
	Description string `json:"description"`
	// The environment type for the custom domain
	EnvironmentType string `json:"environmentType"`
	// The custom domain name
	Name string `json:"name"`
	// The custom domain status
	Status string `json:"status"`
}

type _CustomDomain CustomDomain

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain(clusterEndpoint string, customDomain string, description string, environmentType string, name string, status string) *CustomDomain {
	this := CustomDomain{}
	this.ClusterEndpoint = clusterEndpoint
	this.CustomDomain = customDomain
	this.Description = description
	this.EnvironmentType = environmentType
	this.Name = name
	this.Status = status
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetClusterEndpoint returns the ClusterEndpoint field value
func (o *CustomDomain) GetClusterEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterEndpoint
}

// GetClusterEndpointOk returns a tuple with the ClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetClusterEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterEndpoint, true
}

// SetClusterEndpoint sets field value
func (o *CustomDomain) SetClusterEndpoint(v string) {
	o.ClusterEndpoint = v
}

// GetCustomDomain returns the CustomDomain field value
func (o *CustomDomain) GetCustomDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDomain
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCustomDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDomain, true
}

// SetCustomDomain sets field value
func (o *CustomDomain) SetCustomDomain(v string) {
	o.CustomDomain = v
}

// GetDescription returns the Description field value
func (o *CustomDomain) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CustomDomain) SetDescription(v string) {
	o.Description = v
}

// GetEnvironmentType returns the EnvironmentType field value
func (o *CustomDomain) GetEnvironmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentType, true
}

// SetEnvironmentType sets field value
func (o *CustomDomain) SetEnvironmentType(v string) {
	o.EnvironmentType = v
}

// GetName returns the Name field value
func (o *CustomDomain) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomDomain) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CustomDomain) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CustomDomain) SetStatus(v string) {
	o.Status = v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterEndpoint"] = o.ClusterEndpoint
	toSerialize["customDomain"] = o.CustomDomain
	toSerialize["description"] = o.Description
	toSerialize["environmentType"] = o.EnvironmentType
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *CustomDomain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterEndpoint",
		"customDomain",
		"description",
		"environmentType",
		"name",
		"status",
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

	varCustomDomain := _CustomDomain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomDomain)

	if err != nil {
		return err
	}

	*o = CustomDomain(varCustomDomain)

	return err
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


