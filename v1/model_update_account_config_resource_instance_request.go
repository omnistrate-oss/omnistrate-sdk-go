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

// checks if the UpdateAccountConfigResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccountConfigResourceInstanceRequest{}

// UpdateAccountConfigResourceInstanceRequest struct for UpdateAccountConfigResourceInstanceRequest
type UpdateAccountConfigResourceInstanceRequest struct {
	// The instance ID
	Id string `json:"id"`
	// The service ID
	ServiceId string `json:"serviceId"`
	// set account config instance connection
	SetConnection *bool `json:"setConnection,omitempty"`
	// The subscription ID
	SubscriptionId string `json:"subscriptionId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAccountConfigResourceInstanceRequest UpdateAccountConfigResourceInstanceRequest

// NewUpdateAccountConfigResourceInstanceRequest instantiates a new UpdateAccountConfigResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountConfigResourceInstanceRequest(id string, serviceId string, subscriptionId string, token string) *UpdateAccountConfigResourceInstanceRequest {
	this := UpdateAccountConfigResourceInstanceRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.SubscriptionId = subscriptionId
	this.Token = token
	return &this
}

// NewUpdateAccountConfigResourceInstanceRequestWithDefaults instantiates a new UpdateAccountConfigResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountConfigResourceInstanceRequestWithDefaults() *UpdateAccountConfigResourceInstanceRequest {
	this := UpdateAccountConfigResourceInstanceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateAccountConfigResourceInstanceRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountConfigResourceInstanceRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAccountConfigResourceInstanceRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateAccountConfigResourceInstanceRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountConfigResourceInstanceRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateAccountConfigResourceInstanceRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSetConnection returns the SetConnection field value if set, zero value otherwise.
func (o *UpdateAccountConfigResourceInstanceRequest) GetSetConnection() bool {
	if o == nil || IsNil(o.SetConnection) {
		var ret bool
		return ret
	}
	return *o.SetConnection
}

// GetSetConnectionOk returns a tuple with the SetConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountConfigResourceInstanceRequest) GetSetConnectionOk() (*bool, bool) {
	if o == nil || IsNil(o.SetConnection) {
		return nil, false
	}
	return o.SetConnection, true
}

// SetSetConnection gets a reference to the given bool and assigns it to the SetConnection field.
func (o *UpdateAccountConfigResourceInstanceRequest) SetSetConnection(v bool) {
	o.SetConnection = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UpdateAccountConfigResourceInstanceRequest) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountConfigResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UpdateAccountConfigResourceInstanceRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetToken returns the Token field value
func (o *UpdateAccountConfigResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountConfigResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateAccountConfigResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateAccountConfigResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccountConfigResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SetConnection) {
		toSerialize["setConnection"] = o.SetConnection
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAccountConfigResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serviceId",
		"subscriptionId",
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

	varUpdateAccountConfigResourceInstanceRequest := _UpdateAccountConfigResourceInstanceRequest{}

	err = json.Unmarshal(data, &varUpdateAccountConfigResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = UpdateAccountConfigResourceInstanceRequest(varUpdateAccountConfigResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "setConnection")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAccountConfigResourceInstanceRequest struct {
	value *UpdateAccountConfigResourceInstanceRequest
	isSet bool
}

func (v NullableUpdateAccountConfigResourceInstanceRequest) Get() *UpdateAccountConfigResourceInstanceRequest {
	return v.value
}

func (v *NullableUpdateAccountConfigResourceInstanceRequest) Set(val *UpdateAccountConfigResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountConfigResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountConfigResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountConfigResourceInstanceRequest(val *UpdateAccountConfigResourceInstanceRequest) *NullableUpdateAccountConfigResourceInstanceRequest {
	return &NullableUpdateAccountConfigResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableUpdateAccountConfigResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountConfigResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


