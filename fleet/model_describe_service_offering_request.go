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

// checks if the DescribeServiceOfferingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceOfferingRequest{}

// DescribeServiceOfferingRequest struct for DescribeServiceOfferingRequest
type DescribeServiceOfferingRequest struct {
	// The type of service environment
	EnvironmentType *string `json:"environmentType,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceOfferingRequest DescribeServiceOfferingRequest

// NewDescribeServiceOfferingRequest instantiates a new DescribeServiceOfferingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceOfferingRequest(serviceId string, token string) *DescribeServiceOfferingRequest {
	this := DescribeServiceOfferingRequest{}
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeServiceOfferingRequestWithDefaults instantiates a new DescribeServiceOfferingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceOfferingRequestWithDefaults() *DescribeServiceOfferingRequest {
	this := DescribeServiceOfferingRequest{}
	return &this
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *DescribeServiceOfferingRequest) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingRequest) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *DescribeServiceOfferingRequest) HasEnvironmentType() bool {
	if o != nil && !IsNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *DescribeServiceOfferingRequest) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceOfferingRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceOfferingRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeServiceOfferingRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeServiceOfferingRequest) SetToken(v string) {
	o.Token = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *DescribeServiceOfferingRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *DescribeServiceOfferingRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *DescribeServiceOfferingRequest) SetVisibility(v string) {
	o.Visibility = &v
}

func (o DescribeServiceOfferingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceOfferingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceOfferingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeServiceOfferingRequest := _DescribeServiceOfferingRequest{}

	err = json.Unmarshal(data, &varDescribeServiceOfferingRequest)

	if err != nil {
		return err
	}

	*o = DescribeServiceOfferingRequest(varDescribeServiceOfferingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceOfferingRequest struct {
	value *DescribeServiceOfferingRequest
	isSet bool
}

func (v NullableDescribeServiceOfferingRequest) Get() *DescribeServiceOfferingRequest {
	return v.value
}

func (v *NullableDescribeServiceOfferingRequest) Set(val *DescribeServiceOfferingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceOfferingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceOfferingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceOfferingRequest(val *DescribeServiceOfferingRequest) *NullableDescribeServiceOfferingRequest {
	return &NullableDescribeServiceOfferingRequest{value: val, isSet: true}
}

func (v NullableDescribeServiceOfferingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceOfferingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


