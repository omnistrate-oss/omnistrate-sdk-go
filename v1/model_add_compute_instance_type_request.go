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

// checks if the AddComputeInstanceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddComputeInstanceTypeRequest{}

// AddComputeInstanceTypeRequest struct for AddComputeInstanceTypeRequest
type AddComputeInstanceTypeRequest struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	ConfigOverride *ComputeInstanceTypeConfigOverride `json:"configOverride,omitempty"`
	// ID of a Compute Config
	Id string `json:"id"`
	// The instance type for this compute instance type config
	InstanceType string `json:"instanceType"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _AddComputeInstanceTypeRequest AddComputeInstanceTypeRequest

// NewAddComputeInstanceTypeRequest instantiates a new AddComputeInstanceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddComputeInstanceTypeRequest(cloudProviderName string, id string, instanceType string, serviceId string, token string) *AddComputeInstanceTypeRequest {
	this := AddComputeInstanceTypeRequest{}
	this.CloudProviderName = cloudProviderName
	this.Id = id
	this.InstanceType = instanceType
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewAddComputeInstanceTypeRequestWithDefaults instantiates a new AddComputeInstanceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddComputeInstanceTypeRequestWithDefaults() *AddComputeInstanceTypeRequest {
	this := AddComputeInstanceTypeRequest{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *AddComputeInstanceTypeRequest) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *AddComputeInstanceTypeRequest) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetConfigOverride returns the ConfigOverride field value if set, zero value otherwise.
func (o *AddComputeInstanceTypeRequest) GetConfigOverride() ComputeInstanceTypeConfigOverride {
	if o == nil || IsNil(o.ConfigOverride) {
		var ret ComputeInstanceTypeConfigOverride
		return ret
	}
	return *o.ConfigOverride
}

// GetConfigOverrideOk returns a tuple with the ConfigOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetConfigOverrideOk() (*ComputeInstanceTypeConfigOverride, bool) {
	if o == nil || IsNil(o.ConfigOverride) {
		return nil, false
	}
	return o.ConfigOverride, true
}

// SetConfigOverride gets a reference to the given ComputeInstanceTypeConfigOverride and assigns it to the ConfigOverride field.
func (o *AddComputeInstanceTypeRequest) SetConfigOverride(v ComputeInstanceTypeConfigOverride) {
	o.ConfigOverride = &v
}

// GetId returns the Id field value
func (o *AddComputeInstanceTypeRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddComputeInstanceTypeRequest) SetId(v string) {
	o.Id = v
}

// GetInstanceType returns the InstanceType field value
func (o *AddComputeInstanceTypeRequest) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *AddComputeInstanceTypeRequest) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetServiceId returns the ServiceId field value
func (o *AddComputeInstanceTypeRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *AddComputeInstanceTypeRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *AddComputeInstanceTypeRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AddComputeInstanceTypeRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AddComputeInstanceTypeRequest) SetToken(v string) {
	o.Token = v
}

func (o AddComputeInstanceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddComputeInstanceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	if !IsNil(o.ConfigOverride) {
		toSerialize["configOverride"] = o.ConfigOverride
	}
	toSerialize["id"] = o.Id
	toSerialize["instanceType"] = o.InstanceType
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddComputeInstanceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"id",
		"instanceType",
		"serviceId",
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

	varAddComputeInstanceTypeRequest := _AddComputeInstanceTypeRequest{}

	err = json.Unmarshal(data, &varAddComputeInstanceTypeRequest)

	if err != nil {
		return err
	}

	*o = AddComputeInstanceTypeRequest(varAddComputeInstanceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "configOverride")
		delete(additionalProperties, "id")
		delete(additionalProperties, "instanceType")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddComputeInstanceTypeRequest struct {
	value *AddComputeInstanceTypeRequest
	isSet bool
}

func (v NullableAddComputeInstanceTypeRequest) Get() *AddComputeInstanceTypeRequest {
	return v.value
}

func (v *NullableAddComputeInstanceTypeRequest) Set(val *AddComputeInstanceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddComputeInstanceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddComputeInstanceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddComputeInstanceTypeRequest(val *AddComputeInstanceTypeRequest) *NullableAddComputeInstanceTypeRequest {
	return &NullableAddComputeInstanceTypeRequest{value: val, isSet: true}
}

func (v NullableAddComputeInstanceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddComputeInstanceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


