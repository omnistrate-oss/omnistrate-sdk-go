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

// checks if the ReportHealthRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportHealthRequest{}

// ReportHealthRequest struct for ReportHealthRequest
type ReportHealthRequest struct {
	// ID of a Service
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ReportHealthRequest ReportHealthRequest

// NewReportHealthRequest instantiates a new ReportHealthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportHealthRequest(id string, token string) *ReportHealthRequest {
	this := ReportHealthRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewReportHealthRequestWithDefaults instantiates a new ReportHealthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportHealthRequestWithDefaults() *ReportHealthRequest {
	this := ReportHealthRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ReportHealthRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportHealthRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReportHealthRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *ReportHealthRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ReportHealthRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ReportHealthRequest) SetToken(v string) {
	o.Token = v
}

func (o ReportHealthRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportHealthRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportHealthRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varReportHealthRequest := _ReportHealthRequest{}

	err = json.Unmarshal(data, &varReportHealthRequest)

	if err != nil {
		return err
	}

	*o = ReportHealthRequest(varReportHealthRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportHealthRequest struct {
	value *ReportHealthRequest
	isSet bool
}

func (v NullableReportHealthRequest) Get() *ReportHealthRequest {
	return v.value
}

func (v *NullableReportHealthRequest) Set(val *ReportHealthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportHealthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportHealthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportHealthRequest(val *ReportHealthRequest) *NullableReportHealthRequest {
	return &NullableReportHealthRequest{value: val, isSet: true}
}

func (v NullableReportHealthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportHealthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


