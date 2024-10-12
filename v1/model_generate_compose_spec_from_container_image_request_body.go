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

// checks if the GenerateComposeSpecFromContainerImageRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateComposeSpecFromContainerImageRequestBody{}

// GenerateComposeSpecFromContainerImageRequestBody struct for GenerateComposeSpecFromContainerImageRequestBody
type GenerateComposeSpecFromContainerImageRequestBody struct {
	// Runtime environment variables needed to run the image
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables,omitempty"`
	// Name of the image along with the tag. Include the repository name if the image is not from the official repository
	Image string `json:"image"`
	// Registry where the image is stored
	ImageRegistry string `json:"imageRegistry"`
	// Password to access the image registry
	Password *string `json:"password,omitempty"`
	// Username to access the image registry
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenerateComposeSpecFromContainerImageRequestBody GenerateComposeSpecFromContainerImageRequestBody

// NewGenerateComposeSpecFromContainerImageRequestBody instantiates a new GenerateComposeSpecFromContainerImageRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateComposeSpecFromContainerImageRequestBody(image string, imageRegistry string) *GenerateComposeSpecFromContainerImageRequestBody {
	this := GenerateComposeSpecFromContainerImageRequestBody{}
	this.Image = image
	this.ImageRegistry = imageRegistry
	return &this
}

// NewGenerateComposeSpecFromContainerImageRequestBodyWithDefaults instantiates a new GenerateComposeSpecFromContainerImageRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateComposeSpecFromContainerImageRequestBodyWithDefaults() *GenerateComposeSpecFromContainerImageRequestBody {
	this := GenerateComposeSpecFromContainerImageRequestBody{}
	return &this
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetEnvironmentVariables() []EnvironmentVariable {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []EnvironmentVariable
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetEnvironmentVariablesOk() ([]EnvironmentVariable, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// SetEnvironmentVariables gets a reference to the given []EnvironmentVariable and assigns it to the EnvironmentVariables field.
func (o *GenerateComposeSpecFromContainerImageRequestBody) SetEnvironmentVariables(v []EnvironmentVariable) {
	o.EnvironmentVariables = v
}

// GetImage returns the Image field value
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *GenerateComposeSpecFromContainerImageRequestBody) SetImage(v string) {
	o.Image = v
}

// GetImageRegistry returns the ImageRegistry field value
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageRegistry
}

// GetImageRegistryOk returns a tuple with the ImageRegistry field value
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetImageRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageRegistry, true
}

// SetImageRegistry sets field value
func (o *GenerateComposeSpecFromContainerImageRequestBody) SetImageRegistry(v string) {
	o.ImageRegistry = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GenerateComposeSpecFromContainerImageRequestBody) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateComposeSpecFromContainerImageRequestBody) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GenerateComposeSpecFromContainerImageRequestBody) SetUsername(v string) {
	o.Username = &v
}

func (o GenerateComposeSpecFromContainerImageRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateComposeSpecFromContainerImageRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	toSerialize["image"] = o.Image
	toSerialize["imageRegistry"] = o.ImageRegistry
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenerateComposeSpecFromContainerImageRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
		"imageRegistry",
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

	varGenerateComposeSpecFromContainerImageRequestBody := _GenerateComposeSpecFromContainerImageRequestBody{}

	err = json.Unmarshal(data, &varGenerateComposeSpecFromContainerImageRequestBody)

	if err != nil {
		return err
	}

	*o = GenerateComposeSpecFromContainerImageRequestBody(varGenerateComposeSpecFromContainerImageRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentVariables")
		delete(additionalProperties, "image")
		delete(additionalProperties, "imageRegistry")
		delete(additionalProperties, "password")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenerateComposeSpecFromContainerImageRequestBody struct {
	value *GenerateComposeSpecFromContainerImageRequestBody
	isSet bool
}

func (v NullableGenerateComposeSpecFromContainerImageRequestBody) Get() *GenerateComposeSpecFromContainerImageRequestBody {
	return v.value
}

func (v *NullableGenerateComposeSpecFromContainerImageRequestBody) Set(val *GenerateComposeSpecFromContainerImageRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateComposeSpecFromContainerImageRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateComposeSpecFromContainerImageRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateComposeSpecFromContainerImageRequestBody(val *GenerateComposeSpecFromContainerImageRequestBody) *NullableGenerateComposeSpecFromContainerImageRequestBody {
	return &NullableGenerateComposeSpecFromContainerImageRequestBody{value: val, isSet: true}
}

func (v NullableGenerateComposeSpecFromContainerImageRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateComposeSpecFromContainerImageRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


