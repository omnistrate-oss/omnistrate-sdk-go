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

// checks if the ListAvailabilityZonesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAvailabilityZonesRequest{}

// ListAvailabilityZonesRequest struct for ListAvailabilityZonesRequest
type ListAvailabilityZonesRequest struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListAvailabilityZonesRequest ListAvailabilityZonesRequest

// NewListAvailabilityZonesRequest instantiates a new ListAvailabilityZonesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAvailabilityZonesRequest(cloudProviderName string, token string) *ListAvailabilityZonesRequest {
	this := ListAvailabilityZonesRequest{}
	this.CloudProviderName = cloudProviderName
	this.Token = token
	return &this
}

// NewListAvailabilityZonesRequestWithDefaults instantiates a new ListAvailabilityZonesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAvailabilityZonesRequestWithDefaults() *ListAvailabilityZonesRequest {
	this := ListAvailabilityZonesRequest{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *ListAvailabilityZonesRequest) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *ListAvailabilityZonesRequest) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *ListAvailabilityZonesRequest) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetToken returns the Token field value
func (o *ListAvailabilityZonesRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListAvailabilityZonesRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListAvailabilityZonesRequest) SetToken(v string) {
	o.Token = v
}

func (o ListAvailabilityZonesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAvailabilityZonesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAvailabilityZonesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
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

	varListAvailabilityZonesRequest := _ListAvailabilityZonesRequest{}

	err = json.Unmarshal(data, &varListAvailabilityZonesRequest)

	if err != nil {
		return err
	}

	*o = ListAvailabilityZonesRequest(varListAvailabilityZonesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAvailabilityZonesRequest struct {
	value *ListAvailabilityZonesRequest
	isSet bool
}

func (v NullableListAvailabilityZonesRequest) Get() *ListAvailabilityZonesRequest {
	return v.value
}

func (v *NullableListAvailabilityZonesRequest) Set(val *ListAvailabilityZonesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListAvailabilityZonesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListAvailabilityZonesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAvailabilityZonesRequest(val *ListAvailabilityZonesRequest) *NullableListAvailabilityZonesRequest {
	return &NullableListAvailabilityZonesRequest{value: val, isSet: true}
}

func (v NullableListAvailabilityZonesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAvailabilityZonesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


