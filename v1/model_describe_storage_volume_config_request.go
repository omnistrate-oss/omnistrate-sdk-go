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

// checks if the DescribeStorageVolumeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeStorageVolumeConfigRequest{}

// DescribeStorageVolumeConfigRequest struct for DescribeStorageVolumeConfigRequest
type DescribeStorageVolumeConfigRequest struct {
	// ID of a Product Tier
	ProductTierId *string `json:"ProductTierId,omitempty"`
	// The product tier version of the infra config to describe. If not specified, the latest version is described.
	ProductTierVersion *string `json:"ProductTierVersion,omitempty"`
	// ID of a Storage Volume Config
	Id string `json:"id"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeStorageVolumeConfigRequest DescribeStorageVolumeConfigRequest

// NewDescribeStorageVolumeConfigRequest instantiates a new DescribeStorageVolumeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeStorageVolumeConfigRequest(id string, serviceId string, token string) *DescribeStorageVolumeConfigRequest {
	this := DescribeStorageVolumeConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewDescribeStorageVolumeConfigRequestWithDefaults instantiates a new DescribeStorageVolumeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeStorageVolumeConfigRequestWithDefaults() *DescribeStorageVolumeConfigRequest {
	this := DescribeStorageVolumeConfigRequest{}
	return &this
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *DescribeStorageVolumeConfigRequest) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeStorageVolumeConfigRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *DescribeStorageVolumeConfigRequest) SetProductTierId(v string) {
	o.ProductTierId = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *DescribeStorageVolumeConfigRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeStorageVolumeConfigRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *DescribeStorageVolumeConfigRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetId returns the Id field value
func (o *DescribeStorageVolumeConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageVolumeConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeStorageVolumeConfigRequest) SetId(v string) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeStorageVolumeConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageVolumeConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeStorageVolumeConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *DescribeStorageVolumeConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageVolumeConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeStorageVolumeConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeStorageVolumeConfigRequest) ToMap() (map[string]interface{}, error) {
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

func (o *DescribeStorageVolumeConfigRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeStorageVolumeConfigRequest := _DescribeStorageVolumeConfigRequest{}

	err = json.Unmarshal(data, &varDescribeStorageVolumeConfigRequest)

	if err != nil {
		return err
	}

	*o = DescribeStorageVolumeConfigRequest(varDescribeStorageVolumeConfigRequest)

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

type NullableDescribeStorageVolumeConfigRequest struct {
	value *DescribeStorageVolumeConfigRequest
	isSet bool
}

func (v NullableDescribeStorageVolumeConfigRequest) Get() *DescribeStorageVolumeConfigRequest {
	return v.value
}

func (v *NullableDescribeStorageVolumeConfigRequest) Set(val *DescribeStorageVolumeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeStorageVolumeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeStorageVolumeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeStorageVolumeConfigRequest(val *DescribeStorageVolumeConfigRequest) *NullableDescribeStorageVolumeConfigRequest {
	return &NullableDescribeStorageVolumeConfigRequest{value: val, isSet: true}
}

func (v NullableDescribeStorageVolumeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeStorageVolumeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


