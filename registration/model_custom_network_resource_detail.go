/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
)

// checks if the CustomNetworkResourceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomNetworkResourceDetail{}

// CustomNetworkResourceDetail Details of custom network used for instance
type CustomNetworkResourceDetail struct {
	// CIDR block of the network
	Cidr *string `json:"cidr,omitempty"`
	// ID of a custom network
	Id *string `json:"id,omitempty"`
	// User friendly network name
	Name *string `json:"name,omitempty"`
}

// NewCustomNetworkResourceDetail instantiates a new CustomNetworkResourceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomNetworkResourceDetail() *CustomNetworkResourceDetail {
	this := CustomNetworkResourceDetail{}
	return &this
}

// NewCustomNetworkResourceDetailWithDefaults instantiates a new CustomNetworkResourceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomNetworkResourceDetailWithDefaults() *CustomNetworkResourceDetail {
	this := CustomNetworkResourceDetail{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CustomNetworkResourceDetail) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkResourceDetail) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CustomNetworkResourceDetail) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CustomNetworkResourceDetail) SetCidr(v string) {
	o.Cidr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomNetworkResourceDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkResourceDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomNetworkResourceDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomNetworkResourceDetail) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomNetworkResourceDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkResourceDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomNetworkResourceDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomNetworkResourceDetail) SetName(v string) {
	o.Name = &v
}

func (o CustomNetworkResourceDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomNetworkResourceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCustomNetworkResourceDetail struct {
	value *CustomNetworkResourceDetail
	isSet bool
}

func (v NullableCustomNetworkResourceDetail) Get() *CustomNetworkResourceDetail {
	return v.value
}

func (v *NullableCustomNetworkResourceDetail) Set(val *CustomNetworkResourceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomNetworkResourceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomNetworkResourceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomNetworkResourceDetail(val *CustomNetworkResourceDetail) *NullableCustomNetworkResourceDetail {
	return &NullableCustomNetworkResourceDetail{value: val, isSet: true}
}

func (v NullableCustomNetworkResourceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomNetworkResourceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


