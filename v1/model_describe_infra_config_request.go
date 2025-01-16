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

// checks if the DescribeInfraConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeInfraConfigRequest{}

// DescribeInfraConfigRequest struct for DescribeInfraConfigRequest
type DescribeInfraConfigRequest struct {
	// ID of a Product Tier
	ProductTierId *string `json:"ProductTierId,omitempty"`
	// The product tier version of the infra config to describe. If not specified, the latest version is described.
	ProductTierVersion *string `json:"ProductTierVersion,omitempty"`
	// ID of an Infra Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeInfraConfigRequest DescribeInfraConfigRequest

// NewDescribeInfraConfigRequest instantiates a new DescribeInfraConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeInfraConfigRequest(id string, serviceId string, token string) *DescribeInfraConfigRequest {
	this := DescribeInfraConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeInfraConfigRequestWithDefaults instantiates a new DescribeInfraConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeInfraConfigRequestWithDefaults() *DescribeInfraConfigRequest {
	this := DescribeInfraConfigRequest{}
	return &this
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *DescribeInfraConfigRequest) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInfraConfigRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *DescribeInfraConfigRequest) SetProductTierId(v string) {
	o.ProductTierId = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *DescribeInfraConfigRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInfraConfigRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *DescribeInfraConfigRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetId returns the Id field value
func (o *DescribeInfraConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeInfraConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeInfraConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeInfraConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeInfraConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeInfraConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeInfraConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeInfraConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeInfraConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeInfraConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeInfraConfigRequest) ToMap() (map[string]interface{}, error) {
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

func (o *DescribeInfraConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeInfraConfigRequest := _DescribeInfraConfigRequest{}

	err = json.Unmarshal(data, &varDescribeInfraConfigRequest)

	if err != nil {
		return err
	}

	*o = DescribeInfraConfigRequest(varDescribeInfraConfigRequest)

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

type NullableDescribeInfraConfigRequest struct {
	value *DescribeInfraConfigRequest
	isSet bool
}

func (v NullableDescribeInfraConfigRequest) Get() *DescribeInfraConfigRequest {
	return v.value
}

func (v *NullableDescribeInfraConfigRequest) Set(val *DescribeInfraConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeInfraConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeInfraConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeInfraConfigRequest(val *DescribeInfraConfigRequest) *NullableDescribeInfraConfigRequest {
	return &NullableDescribeInfraConfigRequest{value: val, isSet: true}
}

func (v NullableDescribeInfraConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeInfraConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


