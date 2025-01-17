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

// checks if the ListAllSubscriptionUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllSubscriptionUsersRequest{}

// ListAllSubscriptionUsersRequest struct for ListAllSubscriptionUsersRequest
type ListAllSubscriptionUsersRequest struct {
	// The type of service environment
	EnvironmentType *string `json:"environmentType,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListAllSubscriptionUsersRequest ListAllSubscriptionUsersRequest

// NewListAllSubscriptionUsersRequest instantiates a new ListAllSubscriptionUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllSubscriptionUsersRequest(token string) *ListAllSubscriptionUsersRequest {
	this := ListAllSubscriptionUsersRequest{}
	this.Token = token
	return &this
}

// NewListAllSubscriptionUsersRequestWithDefaults instantiates a new ListAllSubscriptionUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllSubscriptionUsersRequestWithDefaults() *ListAllSubscriptionUsersRequest {
	this := ListAllSubscriptionUsersRequest{}
	return &this
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *ListAllSubscriptionUsersRequest) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllSubscriptionUsersRequest) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *ListAllSubscriptionUsersRequest) HasEnvironmentType() bool {
	if o != nil && !IsNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *ListAllSubscriptionUsersRequest) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetToken returns the Token field value
func (o *ListAllSubscriptionUsersRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListAllSubscriptionUsersRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListAllSubscriptionUsersRequest) SetToken(v string) {
	o.Token = v
}

func (o ListAllSubscriptionUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllSubscriptionUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAllSubscriptionUsersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varListAllSubscriptionUsersRequest := _ListAllSubscriptionUsersRequest{}

	err = json.Unmarshal(data, &varListAllSubscriptionUsersRequest)

	if err != nil {
		return err
	}

	*o = ListAllSubscriptionUsersRequest(varListAllSubscriptionUsersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAllSubscriptionUsersRequest struct {
	value *ListAllSubscriptionUsersRequest
	isSet bool
}

func (v NullableListAllSubscriptionUsersRequest) Get() *ListAllSubscriptionUsersRequest {
	return v.value
}

func (v *NullableListAllSubscriptionUsersRequest) Set(val *ListAllSubscriptionUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllSubscriptionUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllSubscriptionUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllSubscriptionUsersRequest(val *ListAllSubscriptionUsersRequest) *NullableListAllSubscriptionUsersRequest {
	return &NullableListAllSubscriptionUsersRequest{value: val, isSet: true}
}

func (v NullableListAllSubscriptionUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllSubscriptionUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


