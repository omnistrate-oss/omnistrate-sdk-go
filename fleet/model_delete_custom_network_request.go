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

// checks if the DeleteCustomNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCustomNetworkRequest{}

// DeleteCustomNetworkRequest struct for DeleteCustomNetworkRequest
type DeleteCustomNetworkRequest struct {
	// ID of a custom network
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeleteCustomNetworkRequest DeleteCustomNetworkRequest

// NewDeleteCustomNetworkRequest instantiates a new DeleteCustomNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCustomNetworkRequest(id string, token string) *DeleteCustomNetworkRequest {
	this := DeleteCustomNetworkRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewDeleteCustomNetworkRequestWithDefaults instantiates a new DeleteCustomNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCustomNetworkRequestWithDefaults() *DeleteCustomNetworkRequest {
	this := DeleteCustomNetworkRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteCustomNetworkRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomNetworkRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteCustomNetworkRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *DeleteCustomNetworkRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomNetworkRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeleteCustomNetworkRequest) SetToken(v string) {
	o.Token = v
}

func (o DeleteCustomNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCustomNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteCustomNetworkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteCustomNetworkRequest := _DeleteCustomNetworkRequest{}

	err = json.Unmarshal(data, &varDeleteCustomNetworkRequest)

	if err != nil {
		return err
	}

	*o = DeleteCustomNetworkRequest(varDeleteCustomNetworkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteCustomNetworkRequest struct {
	value *DeleteCustomNetworkRequest
	isSet bool
}

func (v NullableDeleteCustomNetworkRequest) Get() *DeleteCustomNetworkRequest {
	return v.value
}

func (v *NullableDeleteCustomNetworkRequest) Set(val *DeleteCustomNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCustomNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCustomNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCustomNetworkRequest(val *DeleteCustomNetworkRequest) *NullableDeleteCustomNetworkRequest {
	return &NullableDeleteCustomNetworkRequest{value: val, isSet: true}
}

func (v NullableDeleteCustomNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCustomNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


