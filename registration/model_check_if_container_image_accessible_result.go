/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CheckIfContainerImageAccessibleResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckIfContainerImageAccessibleResult{}

// CheckIfContainerImageAccessibleResult struct for CheckIfContainerImageAccessibleResult
type CheckIfContainerImageAccessibleResult struct {
	// Error message if the image is not accessible
	ErrorMsg *string `json:"errorMsg,omitempty"`
	// Indicates if the image is accessible
	ImageAccessible bool `json:"imageAccessible"`
}

type _CheckIfContainerImageAccessibleResult CheckIfContainerImageAccessibleResult

// NewCheckIfContainerImageAccessibleResult instantiates a new CheckIfContainerImageAccessibleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckIfContainerImageAccessibleResult(imageAccessible bool) *CheckIfContainerImageAccessibleResult {
	this := CheckIfContainerImageAccessibleResult{}
	this.ImageAccessible = imageAccessible
	return &this
}

// NewCheckIfContainerImageAccessibleResultWithDefaults instantiates a new CheckIfContainerImageAccessibleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckIfContainerImageAccessibleResultWithDefaults() *CheckIfContainerImageAccessibleResult {
	this := CheckIfContainerImageAccessibleResult{}
	return &this
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *CheckIfContainerImageAccessibleResult) GetErrorMsg() string {
	if o == nil || IsNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckIfContainerImageAccessibleResult) GetErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *CheckIfContainerImageAccessibleResult) HasErrorMsg() bool {
	if o != nil && !IsNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *CheckIfContainerImageAccessibleResult) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetImageAccessible returns the ImageAccessible field value
func (o *CheckIfContainerImageAccessibleResult) GetImageAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageAccessible
}

// GetImageAccessibleOk returns a tuple with the ImageAccessible field value
// and a boolean to check if the value has been set.
func (o *CheckIfContainerImageAccessibleResult) GetImageAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageAccessible, true
}

// SetImageAccessible sets field value
func (o *CheckIfContainerImageAccessibleResult) SetImageAccessible(v bool) {
	o.ImageAccessible = v
}

func (o CheckIfContainerImageAccessibleResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckIfContainerImageAccessibleResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMsg) {
		toSerialize["errorMsg"] = o.ErrorMsg
	}
	toSerialize["imageAccessible"] = o.ImageAccessible
	return toSerialize, nil
}

func (o *CheckIfContainerImageAccessibleResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageAccessible",
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

	varCheckIfContainerImageAccessibleResult := _CheckIfContainerImageAccessibleResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCheckIfContainerImageAccessibleResult)

	if err != nil {
		return err
	}

	*o = CheckIfContainerImageAccessibleResult(varCheckIfContainerImageAccessibleResult)

	return err
}

type NullableCheckIfContainerImageAccessibleResult struct {
	value *CheckIfContainerImageAccessibleResult
	isSet bool
}

func (v NullableCheckIfContainerImageAccessibleResult) Get() *CheckIfContainerImageAccessibleResult {
	return v.value
}

func (v *NullableCheckIfContainerImageAccessibleResult) Set(val *CheckIfContainerImageAccessibleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckIfContainerImageAccessibleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckIfContainerImageAccessibleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckIfContainerImageAccessibleResult(val *CheckIfContainerImageAccessibleResult) *NullableCheckIfContainerImageAccessibleResult {
	return &NullableCheckIfContainerImageAccessibleResult{value: val, isSet: true}
}

func (v NullableCheckIfContainerImageAccessibleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckIfContainerImageAccessibleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


