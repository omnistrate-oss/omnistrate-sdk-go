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

// checks if the DescribeNetworkConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeNetworkConfigRequest{}

// DescribeNetworkConfigRequest struct for DescribeNetworkConfigRequest
type DescribeNetworkConfigRequest struct {
	// ID of a Product Tier
	ProductTierId *string `json:"ProductTierId,omitempty"`
	// The product tier version of the infra config to describe. If not specified, the latest version is described.
	ProductTierVersion *string `json:"ProductTierVersion,omitempty"`
	// ID of a Network Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeNetworkConfigRequest DescribeNetworkConfigRequest

// NewDescribeNetworkConfigRequest instantiates a new DescribeNetworkConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeNetworkConfigRequest(id string, serviceId string, token string) *DescribeNetworkConfigRequest {
	this := DescribeNetworkConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeNetworkConfigRequestWithDefaults instantiates a new DescribeNetworkConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeNetworkConfigRequestWithDefaults() *DescribeNetworkConfigRequest {
	this := DescribeNetworkConfigRequest{}
	return &this
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *DescribeNetworkConfigRequest) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *DescribeNetworkConfigRequest) SetProductTierId(v string) {
	o.ProductTierId = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *DescribeNetworkConfigRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *DescribeNetworkConfigRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetId returns the Id field value
func (o *DescribeNetworkConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeNetworkConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeNetworkConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeNetworkConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeNetworkConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeNetworkConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeNetworkConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductTierId) {
		toSerialize["ProductTierId"] = o.ProductTierId
	}
	if !IsNil(o.ProductTierVersion) {
		toSerialize["ProductTierVersion"] = o.ProductTierVersion
	}
	toSerialize["id"] = o.Id
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeNetworkConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeNetworkConfigRequest := _DescribeNetworkConfigRequest{}

	err = json.Unmarshal(data, &varDescribeNetworkConfigRequest)

	if err != nil {
		return err
	}

	*o = DescribeNetworkConfigRequest(varDescribeNetworkConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ProductTierId")
		delete(additionalProperties, "ProductTierVersion")
		delete(additionalProperties, "id")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeNetworkConfigRequest struct {
	value *DescribeNetworkConfigRequest
	isSet bool
}

func (v NullableDescribeNetworkConfigRequest) Get() *DescribeNetworkConfigRequest {
	return v.value
}

func (v *NullableDescribeNetworkConfigRequest) Set(val *DescribeNetworkConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeNetworkConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeNetworkConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeNetworkConfigRequest(val *DescribeNetworkConfigRequest) *NullableDescribeNetworkConfigRequest {
	return &NullableDescribeNetworkConfigRequest{value: val, isSet: true}
}

func (v NullableDescribeNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeNetworkConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


