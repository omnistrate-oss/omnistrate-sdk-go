/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SaveHelmPackageRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveHelmPackageRequestBody{}

// SaveHelmPackageRequestBody struct for SaveHelmPackageRequestBody
type SaveHelmPackageRequestBody struct {
	HelmPackage HelmPackage `json:"helmPackage"`
}

type _SaveHelmPackageRequestBody SaveHelmPackageRequestBody

// NewSaveHelmPackageRequestBody instantiates a new SaveHelmPackageRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveHelmPackageRequestBody(helmPackage HelmPackage) *SaveHelmPackageRequestBody {
	this := SaveHelmPackageRequestBody{}
	this.HelmPackage = helmPackage
	return &this
}

// NewSaveHelmPackageRequestBodyWithDefaults instantiates a new SaveHelmPackageRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveHelmPackageRequestBodyWithDefaults() *SaveHelmPackageRequestBody {
	this := SaveHelmPackageRequestBody{}
	return &this
}

// GetHelmPackage returns the HelmPackage field value
func (o *SaveHelmPackageRequestBody) GetHelmPackage() HelmPackage {
	if o == nil {
		var ret HelmPackage
		return ret
	}

	return o.HelmPackage
}

// GetHelmPackageOk returns a tuple with the HelmPackage field value
// and a boolean to check if the value has been set.
func (o *SaveHelmPackageRequestBody) GetHelmPackageOk() (*HelmPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HelmPackage, true
}

// SetHelmPackage sets field value
func (o *SaveHelmPackageRequestBody) SetHelmPackage(v HelmPackage) {
	o.HelmPackage = v
}

func (o SaveHelmPackageRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveHelmPackageRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["helmPackage"] = o.HelmPackage
	return toSerialize, nil
}

func (o *SaveHelmPackageRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"helmPackage",
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

	varSaveHelmPackageRequestBody := _SaveHelmPackageRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSaveHelmPackageRequestBody)

	if err != nil {
		return err
	}

	*o = SaveHelmPackageRequestBody(varSaveHelmPackageRequestBody)

	return err
}

type NullableSaveHelmPackageRequestBody struct {
	value *SaveHelmPackageRequestBody
	isSet bool
}

func (v NullableSaveHelmPackageRequestBody) Get() *SaveHelmPackageRequestBody {
	return v.value
}

func (v *NullableSaveHelmPackageRequestBody) Set(val *SaveHelmPackageRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveHelmPackageRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveHelmPackageRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveHelmPackageRequestBody(val *SaveHelmPackageRequestBody) *NullableSaveHelmPackageRequestBody {
	return &NullableSaveHelmPackageRequestBody{value: val, isSet: true}
}

func (v NullableSaveHelmPackageRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveHelmPackageRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

