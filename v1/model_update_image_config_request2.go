/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the UpdateImageConfigRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateImageConfigRequest2{}

// UpdateImageConfigRequest2 struct for UpdateImageConfigRequest2
type UpdateImageConfigRequest2 struct {
	CustomImageCommandsAndArgs *CustomImageCommandsAndArgs `json:"customImageCommandsAndArgs,omitempty"`
	// A brief description of the image configuration
	Description *string `json:"description,omitempty"`
	// Name of the container image
	ImageName *string `json:"imageName,omitempty"`
	// The image registry ID to use for the infra
	ImageRegistryId *string `json:"imageRegistryId,omitempty"`
	// PEM-encoded Public key part of the key used to sign the container image
	ImageSignaturePublicKeyPEM *string `json:"imageSignaturePublicKeyPEM,omitempty"`
	// Tag representing the software image version that is currently preferred
	ImageTag *string `json:"imageTag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateImageConfigRequest2 UpdateImageConfigRequest2

// NewUpdateImageConfigRequest2 instantiates a new UpdateImageConfigRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageConfigRequest2() *UpdateImageConfigRequest2 {
	this := UpdateImageConfigRequest2{}
	return &this
}

// NewUpdateImageConfigRequest2WithDefaults instantiates a new UpdateImageConfigRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageConfigRequest2WithDefaults() *UpdateImageConfigRequest2 {
	this := UpdateImageConfigRequest2{}
	return &this
}

// GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs {
	if o == nil || IsNil(o.CustomImageCommandsAndArgs) {
		var ret CustomImageCommandsAndArgs
		return ret
	}
	return *o.CustomImageCommandsAndArgs
}

// GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool) {
	if o == nil || IsNil(o.CustomImageCommandsAndArgs) {
		return nil, false
	}
	return o.CustomImageCommandsAndArgs, true
}

// SetCustomImageCommandsAndArgs gets a reference to the given CustomImageCommandsAndArgs and assigns it to the CustomImageCommandsAndArgs field.
func (o *UpdateImageConfigRequest2) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs) {
	o.CustomImageCommandsAndArgs = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateImageConfigRequest2) SetDescription(v string) {
	o.Description = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *UpdateImageConfigRequest2) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageRegistryId returns the ImageRegistryId field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetImageRegistryId() string {
	if o == nil || IsNil(o.ImageRegistryId) {
		var ret string
		return ret
	}
	return *o.ImageRegistryId
}

// GetImageRegistryIdOk returns a tuple with the ImageRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetImageRegistryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageRegistryId) {
		return nil, false
	}
	return o.ImageRegistryId, true
}

// SetImageRegistryId gets a reference to the given string and assigns it to the ImageRegistryId field.
func (o *UpdateImageConfigRequest2) SetImageRegistryId(v string) {
	o.ImageRegistryId = &v
}

// GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetImageSignaturePublicKeyPEM() string {
	if o == nil || IsNil(o.ImageSignaturePublicKeyPEM) {
		var ret string
		return ret
	}
	return *o.ImageSignaturePublicKeyPEM
}

// GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetImageSignaturePublicKeyPEMOk() (*string, bool) {
	if o == nil || IsNil(o.ImageSignaturePublicKeyPEM) {
		return nil, false
	}
	return o.ImageSignaturePublicKeyPEM, true
}

// SetImageSignaturePublicKeyPEM gets a reference to the given string and assigns it to the ImageSignaturePublicKeyPEM field.
func (o *UpdateImageConfigRequest2) SetImageSignaturePublicKeyPEM(v string) {
	o.ImageSignaturePublicKeyPEM = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *UpdateImageConfigRequest2) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag) {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageConfigRequest2) GetImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTag) {
		return nil, false
	}
	return o.ImageTag, true
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *UpdateImageConfigRequest2) SetImageTag(v string) {
	o.ImageTag = &v
}

func (o UpdateImageConfigRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateImageConfigRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomImageCommandsAndArgs) {
		toSerialize["customImageCommandsAndArgs"] = o.CustomImageCommandsAndArgs
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageName) {
		toSerialize["imageName"] = o.ImageName
	}
	if !IsNil(o.ImageRegistryId) {
		toSerialize["imageRegistryId"] = o.ImageRegistryId
	}
	if !IsNil(o.ImageSignaturePublicKeyPEM) {
		toSerialize["imageSignaturePublicKeyPEM"] = o.ImageSignaturePublicKeyPEM
	}
	if !IsNil(o.ImageTag) {
		toSerialize["imageTag"] = o.ImageTag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateImageConfigRequest2) UnmarshalJSON(data []byte) (err error) {
	varUpdateImageConfigRequest2 := _UpdateImageConfigRequest2{}

	err = json.Unmarshal(data, &varUpdateImageConfigRequest2)

	if err != nil {
		return err
	}

	*o = UpdateImageConfigRequest2(varUpdateImageConfigRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customImageCommandsAndArgs")
		delete(additionalProperties, "description")
		delete(additionalProperties, "imageName")
		delete(additionalProperties, "imageRegistryId")
		delete(additionalProperties, "imageSignaturePublicKeyPEM")
		delete(additionalProperties, "imageTag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateImageConfigRequest2 struct {
	value *UpdateImageConfigRequest2
	isSet bool
}

func (v NullableUpdateImageConfigRequest2) Get() *UpdateImageConfigRequest2 {
	return v.value
}

func (v *NullableUpdateImageConfigRequest2) Set(val *UpdateImageConfigRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageConfigRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageConfigRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageConfigRequest2(val *UpdateImageConfigRequest2) *NullableUpdateImageConfigRequest2 {
	return &NullableUpdateImageConfigRequest2{value: val, isSet: true}
}

func (v NullableUpdateImageConfigRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageConfigRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


