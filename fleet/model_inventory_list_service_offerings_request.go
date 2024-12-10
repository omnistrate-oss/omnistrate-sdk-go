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

// checks if the InventoryListServiceOfferingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryListServiceOfferingsRequest{}

// InventoryListServiceOfferingsRequest struct for InventoryListServiceOfferingsRequest
type InventoryListServiceOfferingsRequest struct {
	// ID of an Org
	OrgId *string `json:"orgId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryListServiceOfferingsRequest InventoryListServiceOfferingsRequest

// NewInventoryListServiceOfferingsRequest instantiates a new InventoryListServiceOfferingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListServiceOfferingsRequest(token string) *InventoryListServiceOfferingsRequest {
	this := InventoryListServiceOfferingsRequest{}
	this.Token = token
	return &this
}

// NewInventoryListServiceOfferingsRequestWithDefaults instantiates a new InventoryListServiceOfferingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListServiceOfferingsRequestWithDefaults() *InventoryListServiceOfferingsRequest {
	this := InventoryListServiceOfferingsRequest{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InventoryListServiceOfferingsRequest) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListServiceOfferingsRequest) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InventoryListServiceOfferingsRequest) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *InventoryListServiceOfferingsRequest) SetOrgId(v string) {
	o.OrgId = &v
}

// GetToken returns the Token field value
func (o *InventoryListServiceOfferingsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *InventoryListServiceOfferingsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *InventoryListServiceOfferingsRequest) SetToken(v string) {
	o.Token = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *InventoryListServiceOfferingsRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListServiceOfferingsRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *InventoryListServiceOfferingsRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *InventoryListServiceOfferingsRequest) SetVisibility(v string) {
	o.Visibility = &v
}

func (o InventoryListServiceOfferingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryListServiceOfferingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryListServiceOfferingsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varInventoryListServiceOfferingsRequest := _InventoryListServiceOfferingsRequest{}

	err = json.Unmarshal(data, &varInventoryListServiceOfferingsRequest)

	if err != nil {
		return err
	}

	*o = InventoryListServiceOfferingsRequest(varInventoryListServiceOfferingsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryListServiceOfferingsRequest struct {
	value *InventoryListServiceOfferingsRequest
	isSet bool
}

func (v NullableInventoryListServiceOfferingsRequest) Get() *InventoryListServiceOfferingsRequest {
	return v.value
}

func (v *NullableInventoryListServiceOfferingsRequest) Set(val *InventoryListServiceOfferingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListServiceOfferingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListServiceOfferingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListServiceOfferingsRequest(val *InventoryListServiceOfferingsRequest) *NullableInventoryListServiceOfferingsRequest {
	return &NullableInventoryListServiceOfferingsRequest{value: val, isSet: true}
}

func (v NullableInventoryListServiceOfferingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListServiceOfferingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


