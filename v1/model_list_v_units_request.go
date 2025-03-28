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

// checks if the ListVUnitsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVUnitsRequest{}

// ListVUnitsRequest struct for ListVUnitsRequest
type ListVUnitsRequest struct {
	// Name of the Infra Provider
	CloudProvider string `json:"cloudProvider"`
	// Region code specific to the cloud-provider
	Region string `json:"region"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Service Model
	ServiceModelId string `json:"serviceModelId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListVUnitsRequest ListVUnitsRequest

// NewListVUnitsRequest instantiates a new ListVUnitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVUnitsRequest(cloudProvider string, region string, serviceId string, serviceModelId string, token string) *ListVUnitsRequest {
	this := ListVUnitsRequest{}
	this.CloudProvider = cloudProvider
	this.Region = region
	this.ServiceId = serviceId
	this.ServiceModelId = serviceModelId
	this.Token = token
	return &this
}

// NewListVUnitsRequestWithDefaults instantiates a new ListVUnitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVUnitsRequestWithDefaults() *ListVUnitsRequest {
	this := ListVUnitsRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *ListVUnitsRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ListVUnitsRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ListVUnitsRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *ListVUnitsRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ListVUnitsRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ListVUnitsRequest) SetRegion(v string) {
	o.Region = v
}

// GetServiceId returns the ServiceId field value
func (o *ListVUnitsRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListVUnitsRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListVUnitsRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *ListVUnitsRequest) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *ListVUnitsRequest) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *ListVUnitsRequest) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

// GetToken returns the Token field value
func (o *ListVUnitsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListVUnitsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListVUnitsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListVUnitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVUnitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["region"] = o.Region
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceModelId"] = o.ServiceModelId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListVUnitsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProvider",
		"region",
		"serviceId",
		"serviceModelId",
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

	varListVUnitsRequest := _ListVUnitsRequest{}

	err = json.Unmarshal(data, &varListVUnitsRequest)

	if err != nil {
		return err
	}

	*o = ListVUnitsRequest(varListVUnitsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProvider")
		delete(additionalProperties, "region")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "serviceModelId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListVUnitsRequest struct {
	value *ListVUnitsRequest
	isSet bool
}

func (v NullableListVUnitsRequest) Get() *ListVUnitsRequest {
	return v.value
}

func (v *NullableListVUnitsRequest) Set(val *ListVUnitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListVUnitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListVUnitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVUnitsRequest(val *ListVUnitsRequest) *NullableListVUnitsRequest {
	return &NullableListVUnitsRequest{value: val, isSet: true}
}

func (v NullableListVUnitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVUnitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


