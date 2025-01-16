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

// checks if the PromoteTierVersionSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromoteTierVersionSetRequest{}

// PromoteTierVersionSetRequest struct for PromoteTierVersionSetRequest
type PromoteTierVersionSetRequest struct {
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// The version number for the specific version set.
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _PromoteTierVersionSetRequest PromoteTierVersionSetRequest

// NewPromoteTierVersionSetRequest instantiates a new PromoteTierVersionSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromoteTierVersionSetRequest(productTierId string, serviceId string, token string, version string) *PromoteTierVersionSetRequest {
	this := PromoteTierVersionSetRequest{}
	this.ProductTierId = productTierId
	this.ServiceId = serviceId
	this.Token = token
	this.Version = version
	return &this
}

// NewPromoteTierVersionSetRequestWithDefaults instantiates a new PromoteTierVersionSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromoteTierVersionSetRequestWithDefaults() *PromoteTierVersionSetRequest {
	this := PromoteTierVersionSetRequest{}
	return &this
}

// GetProductTierId returns the ProductTierId field value
func (o *PromoteTierVersionSetRequest) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *PromoteTierVersionSetRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *PromoteTierVersionSetRequest) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetServiceId returns the ServiceId field value
func (o *PromoteTierVersionSetRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *PromoteTierVersionSetRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *PromoteTierVersionSetRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *PromoteTierVersionSetRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PromoteTierVersionSetRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PromoteTierVersionSetRequest) SetToken(v string) {
	o.Token = v
}

// GetVersion returns the Version field value
func (o *PromoteTierVersionSetRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PromoteTierVersionSetRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PromoteTierVersionSetRequest) SetVersion(v string) {
	o.Version = v
}

func (o PromoteTierVersionSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromoteTierVersionSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PromoteTierVersionSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierId",
		"serviceId",
		"token",
		"version",
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

	varPromoteTierVersionSetRequest := _PromoteTierVersionSetRequest{}

	err = json.Unmarshal(data, &varPromoteTierVersionSetRequest)

	if err != nil {
		return err
	}

	*o = PromoteTierVersionSetRequest(varPromoteTierVersionSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePromoteTierVersionSetRequest struct {
	value *PromoteTierVersionSetRequest
	isSet bool
}

func (v NullablePromoteTierVersionSetRequest) Get() *PromoteTierVersionSetRequest {
	return v.value
}

func (v *NullablePromoteTierVersionSetRequest) Set(val *PromoteTierVersionSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromoteTierVersionSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromoteTierVersionSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromoteTierVersionSetRequest(val *PromoteTierVersionSetRequest) *NullablePromoteTierVersionSetRequest {
	return &NullablePromoteTierVersionSetRequest{value: val, isSet: true}
}

func (v NullablePromoteTierVersionSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromoteTierVersionSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


