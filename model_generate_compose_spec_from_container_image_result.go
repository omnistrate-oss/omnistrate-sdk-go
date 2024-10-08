/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GenerateComposeSpecFromContainerImageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateComposeSpecFromContainerImageResult{}

// GenerateComposeSpecFromContainerImageResult struct for GenerateComposeSpecFromContainerImageResult
type GenerateComposeSpecFromContainerImageResult struct {
	// Base64 encoded Compose Spec YAML in docker compose format
	FileContent string `json:"fileContent"`
}

type _GenerateComposeSpecFromContainerImageResult GenerateComposeSpecFromContainerImageResult

// NewGenerateComposeSpecFromContainerImageResult instantiates a new GenerateComposeSpecFromContainerImageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateComposeSpecFromContainerImageResult(fileContent string) *GenerateComposeSpecFromContainerImageResult {
	this := GenerateComposeSpecFromContainerImageResult{}
	this.FileContent = fileContent
	return &this
}

// NewGenerateComposeSpecFromContainerImageResultWithDefaults instantiates a new GenerateComposeSpecFromContainerImageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateComposeSpecFromContainerImageResultWithDefaults() *GenerateComposeSpecFromContainerImageResult {
	this := GenerateComposeSpecFromContainerImageResult{}
	return &this
}

// GetFileContent returns the FileContent field value
func (o *GenerateComposeSpecFromContainerImageResult) GetFileContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageResult) GetFileContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileContent, true
}

// SetFileContent sets field value
func (o *GenerateComposeSpecFromContainerImageResult) SetFileContent(v string) {
	o.FileContent = v
}

func (o GenerateComposeSpecFromContainerImageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateComposeSpecFromContainerImageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileContent"] = o.FileContent
	return toSerialize, nil
}

func (o *GenerateComposeSpecFromContainerImageResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileContent",
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

	varGenerateComposeSpecFromContainerImageResult := _GenerateComposeSpecFromContainerImageResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateComposeSpecFromContainerImageResult)

	if err != nil {
		return err
	}

	*o = GenerateComposeSpecFromContainerImageResult(varGenerateComposeSpecFromContainerImageResult)

	return err
}

type NullableGenerateComposeSpecFromContainerImageResult struct {
	value *GenerateComposeSpecFromContainerImageResult
	isSet bool
}

func (v NullableGenerateComposeSpecFromContainerImageResult) Get() *GenerateComposeSpecFromContainerImageResult {
	return v.value
}

func (v *NullableGenerateComposeSpecFromContainerImageResult) Set(val *GenerateComposeSpecFromContainerImageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateComposeSpecFromContainerImageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateComposeSpecFromContainerImageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateComposeSpecFromContainerImageResult(val *GenerateComposeSpecFromContainerImageResult) *NullableGenerateComposeSpecFromContainerImageResult {
	return &NullableGenerateComposeSpecFromContainerImageResult{value: val, isSet: true}
}

func (v NullableGenerateComposeSpecFromContainerImageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateComposeSpecFromContainerImageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


