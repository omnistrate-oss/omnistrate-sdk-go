/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
)

// checks if the UpdateServiceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceRequestBody{}

// UpdateServiceRequestBody struct for UpdateServiceRequestBody
type UpdateServiceRequestBody struct {
	// A brief description of the service
	Description *string `json:"description,omitempty"`
	// Name of the Service
	Name *string `json:"name,omitempty"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
}

// NewUpdateServiceRequestBody instantiates a new UpdateServiceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceRequestBody() *UpdateServiceRequestBody {
	this := UpdateServiceRequestBody{}
	return &this
}

// NewUpdateServiceRequestBodyWithDefaults instantiates a new UpdateServiceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceRequestBodyWithDefaults() *UpdateServiceRequestBody {
	this := UpdateServiceRequestBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateServiceRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateServiceRequestBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateServiceRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServiceRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServiceRequestBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServiceRequestBody) SetName(v string) {
	o.Name = &v
}

// GetServiceLogoURL returns the ServiceLogoURL field value if set, zero value otherwise.
func (o *UpdateServiceRequestBody) GetServiceLogoURL() string {
	if o == nil || IsNil(o.ServiceLogoURL) {
		var ret string
		return ret
	}
	return *o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequestBody) GetServiceLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogoURL) {
		return nil, false
	}
	return o.ServiceLogoURL, true
}

// HasServiceLogoURL returns a boolean if a field has been set.
func (o *UpdateServiceRequestBody) HasServiceLogoURL() bool {
	if o != nil && !IsNil(o.ServiceLogoURL) {
		return true
	}

	return false
}

// SetServiceLogoURL gets a reference to the given string and assigns it to the ServiceLogoURL field.
func (o *UpdateServiceRequestBody) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = &v
}

func (o UpdateServiceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	return toSerialize, nil
}

type NullableUpdateServiceRequestBody struct {
	value *UpdateServiceRequestBody
	isSet bool
}

func (v NullableUpdateServiceRequestBody) Get() *UpdateServiceRequestBody {
	return v.value
}

func (v *NullableUpdateServiceRequestBody) Set(val *UpdateServiceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceRequestBody(val *UpdateServiceRequestBody) *NullableUpdateServiceRequestBody {
	return &NullableUpdateServiceRequestBody{value: val, isSet: true}
}

func (v NullableUpdateServiceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


