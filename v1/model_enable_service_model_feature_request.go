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

// checks if the EnableServiceModelFeatureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableServiceModelFeatureRequest{}

// EnableServiceModelFeatureRequest struct for EnableServiceModelFeatureRequest
type EnableServiceModelFeatureRequest struct {
	// The configuration parameters of the service model feature
	Configuration map[string]interface{} `json:"configuration"`
	// Name of the service model feature
	Feature string `json:"feature"`
	// ID of a Service Model
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _EnableServiceModelFeatureRequest EnableServiceModelFeatureRequest

// NewEnableServiceModelFeatureRequest instantiates a new EnableServiceModelFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableServiceModelFeatureRequest(configuration map[string]interface{}, feature string, id string, serviceId string, token string) *EnableServiceModelFeatureRequest {
	this := EnableServiceModelFeatureRequest{}
	this.Configuration = configuration
	this.Feature = feature
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewEnableServiceModelFeatureRequestWithDefaults instantiates a new EnableServiceModelFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableServiceModelFeatureRequestWithDefaults() *EnableServiceModelFeatureRequest {
	this := EnableServiceModelFeatureRequest{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *EnableServiceModelFeatureRequest) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *EnableServiceModelFeatureRequest) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value
func (o *EnableServiceModelFeatureRequest) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetFeature returns the Feature field value
func (o *EnableServiceModelFeatureRequest) GetFeature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
func (o *EnableServiceModelFeatureRequest) GetFeatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feature, true
}

// SetFeature sets field value
func (o *EnableServiceModelFeatureRequest) SetFeature(v string) {
	o.Feature = v
}

// GetId returns the Id field value
func (o *EnableServiceModelFeatureRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnableServiceModelFeatureRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnableServiceModelFeatureRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *EnableServiceModelFeatureRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *EnableServiceModelFeatureRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *EnableServiceModelFeatureRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *EnableServiceModelFeatureRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *EnableServiceModelFeatureRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *EnableServiceModelFeatureRequest) SetToken(v string) {
	o.Token = v
}

func (o EnableServiceModelFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableServiceModelFeatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	toSerialize["feature"] = o.Feature
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableServiceModelFeatureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configuration",
		"feature",
		"id",
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

	varEnableServiceModelFeatureRequest := _EnableServiceModelFeatureRequest{}

	err = json.Unmarshal(data, &varEnableServiceModelFeatureRequest)

	if err != nil {
		return err
	}

	*o = EnableServiceModelFeatureRequest(varEnableServiceModelFeatureRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "feature")
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableServiceModelFeatureRequest struct {
	value *EnableServiceModelFeatureRequest
	isSet bool
}

func (v NullableEnableServiceModelFeatureRequest) Get() *EnableServiceModelFeatureRequest {
	return v.value
}

func (v *NullableEnableServiceModelFeatureRequest) Set(val *EnableServiceModelFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableServiceModelFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableServiceModelFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableServiceModelFeatureRequest(val *EnableServiceModelFeatureRequest) *NullableEnableServiceModelFeatureRequest {
	return &NullableEnableServiceModelFeatureRequest{value: val, isSet: true}
}

func (v NullableEnableServiceModelFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableServiceModelFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


