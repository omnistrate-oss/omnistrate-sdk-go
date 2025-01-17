/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the DescribeServiceOfferingResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceOfferingResourceRequest{}

// DescribeServiceOfferingResourceRequest struct for DescribeServiceOfferingResourceRequest
type DescribeServiceOfferingResourceRequest struct {
	// The instance ID
	InstanceId *string `json:"instanceId,omitempty"`
	// ID of a resource
	ResourceId string `json:"resourceId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceOfferingResourceRequest DescribeServiceOfferingResourceRequest

// NewDescribeServiceOfferingResourceRequest instantiates a new DescribeServiceOfferingResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceOfferingResourceRequest(resourceId string, serviceId string, token string) *DescribeServiceOfferingResourceRequest {
	this := DescribeServiceOfferingResourceRequest{}
	var instanceId string = "none"
	this.InstanceId = &instanceId
	this.ResourceId = resourceId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeServiceOfferingResourceRequestWithDefaults instantiates a new DescribeServiceOfferingResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceOfferingResourceRequestWithDefaults() *DescribeServiceOfferingResourceRequest {
	this := DescribeServiceOfferingResourceRequest{}
	var instanceId string = "none"
	this.InstanceId = &instanceId
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *DescribeServiceOfferingResourceRequest) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResourceRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *DescribeServiceOfferingResourceRequest) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *DescribeServiceOfferingResourceRequest) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetResourceId returns the ResourceId field value
func (o *DescribeServiceOfferingResourceRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResourceRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *DescribeServiceOfferingResourceRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceOfferingResourceRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResourceRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceOfferingResourceRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeServiceOfferingResourceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResourceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeServiceOfferingResourceRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeServiceOfferingResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceOfferingResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceOfferingResourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
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

	varDescribeServiceOfferingResourceRequest := _DescribeServiceOfferingResourceRequest{}

	err = json.Unmarshal(data, &varDescribeServiceOfferingResourceRequest)

	if err != nil {
		return err
	}

	*o = DescribeServiceOfferingResourceRequest(varDescribeServiceOfferingResourceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceOfferingResourceRequest struct {
	value *DescribeServiceOfferingResourceRequest
	isSet bool
}

func (v NullableDescribeServiceOfferingResourceRequest) Get() *DescribeServiceOfferingResourceRequest {
	return v.value
}

func (v *NullableDescribeServiceOfferingResourceRequest) Set(val *DescribeServiceOfferingResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceOfferingResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceOfferingResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceOfferingResourceRequest(val *DescribeServiceOfferingResourceRequest) *NullableDescribeServiceOfferingResourceRequest {
	return &NullableDescribeServiceOfferingResourceRequest{value: val, isSet: true}
}

func (v NullableDescribeServiceOfferingResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceOfferingResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


