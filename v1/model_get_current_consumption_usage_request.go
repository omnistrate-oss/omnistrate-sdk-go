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

// checks if the GetCurrentConsumptionUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrentConsumptionUsageRequest{}

// GetCurrentConsumptionUsageRequest struct for GetCurrentConsumptionUsageRequest
type GetCurrentConsumptionUsageRequest struct {
	// ID of a Subscription
	SubscriptionID *string `json:"subscriptionID,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _GetCurrentConsumptionUsageRequest GetCurrentConsumptionUsageRequest

// NewGetCurrentConsumptionUsageRequest instantiates a new GetCurrentConsumptionUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentConsumptionUsageRequest(token string) *GetCurrentConsumptionUsageRequest {
	this := GetCurrentConsumptionUsageRequest{}
	this.Token = token
	return &this
}

// NewGetCurrentConsumptionUsageRequestWithDefaults instantiates a new GetCurrentConsumptionUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentConsumptionUsageRequestWithDefaults() *GetCurrentConsumptionUsageRequest {
	this := GetCurrentConsumptionUsageRequest{}
	return &this
}

// GetSubscriptionID returns the SubscriptionID field value if set, zero value otherwise.
func (o *GetCurrentConsumptionUsageRequest) GetSubscriptionID() string {
	if o == nil || IsNil(o.SubscriptionID) {
		var ret string
		return ret
	}
	return *o.SubscriptionID
}

// GetSubscriptionIDOk returns a tuple with the SubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentConsumptionUsageRequest) GetSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionID) {
		return nil, false
	}
	return o.SubscriptionID, true
}

// SetSubscriptionID gets a reference to the given string and assigns it to the SubscriptionID field.
func (o *GetCurrentConsumptionUsageRequest) SetSubscriptionID(v string) {
	o.SubscriptionID = &v
}

// GetToken returns the Token field value
func (o *GetCurrentConsumptionUsageRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetCurrentConsumptionUsageRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetCurrentConsumptionUsageRequest) SetToken(v string) {
	o.Token = v
}

func (o GetCurrentConsumptionUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentConsumptionUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionID) {
		toSerialize["subscriptionID"] = o.SubscriptionID
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCurrentConsumptionUsageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGetCurrentConsumptionUsageRequest := _GetCurrentConsumptionUsageRequest{}

	err = json.Unmarshal(data, &varGetCurrentConsumptionUsageRequest)

	if err != nil {
		return err
	}

	*o = GetCurrentConsumptionUsageRequest(varGetCurrentConsumptionUsageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscriptionID")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCurrentConsumptionUsageRequest struct {
	value *GetCurrentConsumptionUsageRequest
	isSet bool
}

func (v NullableGetCurrentConsumptionUsageRequest) Get() *GetCurrentConsumptionUsageRequest {
	return v.value
}

func (v *NullableGetCurrentConsumptionUsageRequest) Set(val *GetCurrentConsumptionUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentConsumptionUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentConsumptionUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentConsumptionUsageRequest(val *GetCurrentConsumptionUsageRequest) *NullableGetCurrentConsumptionUsageRequest {
	return &NullableGetCurrentConsumptionUsageRequest{value: val, isSet: true}
}

func (v NullableGetCurrentConsumptionUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentConsumptionUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


