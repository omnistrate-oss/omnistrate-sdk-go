/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the GetUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsageRequest{}

// GetUsageRequest struct for GetUsageRequest
type GetUsageRequest struct {
	// End time of the window in ISO 8601 format. If omitted, the filter is open-ended at the start.
	EndDate *time.Time `json:"endDate,omitempty"`
	// Start time of the window in ISO 8601 format. If omitted, the filter is open-ended at the start.
	StartDate *time.Time `json:"startDate,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _GetUsageRequest GetUsageRequest

// NewGetUsageRequest instantiates a new GetUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsageRequest(token string) *GetUsageRequest {
	this := GetUsageRequest{}
	this.Token = token
	return &this
}

// NewGetUsageRequestWithDefaults instantiates a new GetUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsageRequestWithDefaults() *GetUsageRequest {
	this := GetUsageRequest{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetUsageRequest) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetUsageRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GetUsageRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetUsageRequest) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetUsageRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *GetUsageRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetToken returns the Token field value
func (o *GetUsageRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetUsageRequest) SetToken(v string) {
	o.Token = v
}

func (o GetUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUsageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGetUsageRequest := _GetUsageRequest{}

	err = json.Unmarshal(data, &varGetUsageRequest)

	if err != nil {
		return err
	}

	*o = GetUsageRequest(varGetUsageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUsageRequest struct {
	value *GetUsageRequest
	isSet bool
}

func (v NullableGetUsageRequest) Get() *GetUsageRequest {
	return v.value
}

func (v *NullableGetUsageRequest) Set(val *GetUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsageRequest(val *GetUsageRequest) *NullableGetUsageRequest {
	return &NullableGetUsageRequest{value: val, isSet: true}
}

func (v NullableGetUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


