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

// checks if the DescribeResourceMetricsConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeResourceMetricsConfigRequest{}

// DescribeResourceMetricsConfigRequest struct for DescribeResourceMetricsConfigRequest
type DescribeResourceMetricsConfigRequest struct {
	// ID of a Product Tier
	ProductTierId *string `json:"ProductTierId,omitempty"`
	// The product tier version of the infra config to describe. If not specified, the latest version is described.
	ProductTierVersion *string `json:"ProductTierVersion,omitempty"`
	// ID of a resource
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeResourceMetricsConfigRequest DescribeResourceMetricsConfigRequest

// NewDescribeResourceMetricsConfigRequest instantiates a new DescribeResourceMetricsConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeResourceMetricsConfigRequest(id string, serviceId string, token string) *DescribeResourceMetricsConfigRequest {
	this := DescribeResourceMetricsConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeResourceMetricsConfigRequestWithDefaults instantiates a new DescribeResourceMetricsConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeResourceMetricsConfigRequestWithDefaults() *DescribeResourceMetricsConfigRequest {
	this := DescribeResourceMetricsConfigRequest{}
	return &this
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *DescribeResourceMetricsConfigRequest) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceMetricsConfigRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *DescribeResourceMetricsConfigRequest) SetProductTierId(v string) {
	o.ProductTierId = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *DescribeResourceMetricsConfigRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceMetricsConfigRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *DescribeResourceMetricsConfigRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetId returns the Id field value
func (o *DescribeResourceMetricsConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceMetricsConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeResourceMetricsConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeResourceMetricsConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceMetricsConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeResourceMetricsConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeResourceMetricsConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceMetricsConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeResourceMetricsConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeResourceMetricsConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeResourceMetricsConfigRequest) ToMap() (map[string]interface{}, error) {
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

func (o *DescribeResourceMetricsConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeResourceMetricsConfigRequest := _DescribeResourceMetricsConfigRequest{}

	err = json.Unmarshal(data, &varDescribeResourceMetricsConfigRequest)

	if err != nil {
		return err
	}

	*o = DescribeResourceMetricsConfigRequest(varDescribeResourceMetricsConfigRequest)

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

type NullableDescribeResourceMetricsConfigRequest struct {
	value *DescribeResourceMetricsConfigRequest
	isSet bool
}

func (v NullableDescribeResourceMetricsConfigRequest) Get() *DescribeResourceMetricsConfigRequest {
	return v.value
}

func (v *NullableDescribeResourceMetricsConfigRequest) Set(val *DescribeResourceMetricsConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeResourceMetricsConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeResourceMetricsConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeResourceMetricsConfigRequest(val *DescribeResourceMetricsConfigRequest) *NullableDescribeResourceMetricsConfigRequest {
	return &NullableDescribeResourceMetricsConfigRequest{value: val, isSet: true}
}

func (v NullableDescribeResourceMetricsConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeResourceMetricsConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


