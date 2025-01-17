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

// checks if the DeleteServicesOrchestrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteServicesOrchestrationRequest{}

// DeleteServicesOrchestrationRequest struct for DeleteServicesOrchestrationRequest
type DeleteServicesOrchestrationRequest struct {
	// ID of a Services Orchestration
	Id string `json:"id"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeleteServicesOrchestrationRequest DeleteServicesOrchestrationRequest

// NewDeleteServicesOrchestrationRequest instantiates a new DeleteServicesOrchestrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteServicesOrchestrationRequest(id string, token string) *DeleteServicesOrchestrationRequest {
	this := DeleteServicesOrchestrationRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewDeleteServicesOrchestrationRequestWithDefaults instantiates a new DeleteServicesOrchestrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteServicesOrchestrationRequestWithDefaults() *DeleteServicesOrchestrationRequest {
	this := DeleteServicesOrchestrationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteServicesOrchestrationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteServicesOrchestrationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteServicesOrchestrationRequest) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *DeleteServicesOrchestrationRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeleteServicesOrchestrationRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeleteServicesOrchestrationRequest) SetToken(v string) {
	o.Token = v
}

func (o DeleteServicesOrchestrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteServicesOrchestrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteServicesOrchestrationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteServicesOrchestrationRequest := _DeleteServicesOrchestrationRequest{}

	err = json.Unmarshal(data, &varDeleteServicesOrchestrationRequest)

	if err != nil {
		return err
	}

	*o = DeleteServicesOrchestrationRequest(varDeleteServicesOrchestrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteServicesOrchestrationRequest struct {
	value *DeleteServicesOrchestrationRequest
	isSet bool
}

func (v NullableDeleteServicesOrchestrationRequest) Get() *DeleteServicesOrchestrationRequest {
	return v.value
}

func (v *NullableDeleteServicesOrchestrationRequest) Set(val *DeleteServicesOrchestrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteServicesOrchestrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteServicesOrchestrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteServicesOrchestrationRequest(val *DeleteServicesOrchestrationRequest) *NullableDeleteServicesOrchestrationRequest {
	return &NullableDeleteServicesOrchestrationRequest{value: val, isSet: true}
}

func (v NullableDeleteServicesOrchestrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteServicesOrchestrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


