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

// checks if the ListResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourcesRequest{}

// ListResourcesRequest struct for ListResourcesRequest
type ListResourcesRequest struct {
	// The product tier version of the infra config to describe. If not specified, the latest version is described.
	ProductTierVersion *string `json:"ProductTierVersion,omitempty"`
	// Is resource managed by omnistrate
	Managed *bool `json:"managed,omitempty"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListResourcesRequest ListResourcesRequest

// NewListResourcesRequest instantiates a new ListResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourcesRequest(productTierId string, serviceId string, token string) *ListResourcesRequest {
	this := ListResourcesRequest{}
	this.ProductTierId = productTierId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewListResourcesRequestWithDefaults instantiates a new ListResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourcesRequestWithDefaults() *ListResourcesRequest {
	this := ListResourcesRequest{}
	return &this
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *ListResourcesRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *ListResourcesRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ListResourcesRequest) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ListResourcesRequest) SetManaged(v bool) {
	o.Managed = &v
}

// GetProductTierId returns the ProductTierId field value
func (o *ListResourcesRequest) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *ListResourcesRequest) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetServiceId returns the ServiceId field value
func (o *ListResourcesRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListResourcesRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *ListResourcesRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListResourcesRequest) SetToken(v string) {
	o.Token = v
}

func (o ListResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductTierVersion) {
		toSerialize["ProductTierVersion"] = o.ProductTierVersion
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListResourcesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierId",
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

	varListResourcesRequest := _ListResourcesRequest{}

	err = json.Unmarshal(data, &varListResourcesRequest)

	if err != nil {
		return err
	}

	*o = ListResourcesRequest(varListResourcesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ProductTierVersion")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListResourcesRequest struct {
	value *ListResourcesRequest
	isSet bool
}

func (v NullableListResourcesRequest) Get() *ListResourcesRequest {
	return v.value
}

func (v *NullableListResourcesRequest) Set(val *ListResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourcesRequest(val *ListResourcesRequest) *NullableListResourcesRequest {
	return &NullableListResourcesRequest{value: val, isSet: true}
}

func (v NullableListResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


