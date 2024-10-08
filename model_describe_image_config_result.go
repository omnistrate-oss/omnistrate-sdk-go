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

// checks if the DescribeImageConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeImageConfigResult{}

// DescribeImageConfigResult struct for DescribeImageConfigResult
type DescribeImageConfigResult struct {
	CustomImageCommandsAndArgs *CustomImageCommandsAndArgs `json:"customImageCommandsAndArgs,omitempty"`
	// A brief description of the image configuration
	Description string `json:"description"`
	// The image configuration ID
	Id string `json:"id"`
	// Name of the container image
	ImageName string `json:"imageName"`
	// The image registry ID to use for the infra
	ImageRegistryId string `json:"imageRegistryId"`
	// PEM-encoded Public key part of the key used to sign the container image
	ImageSignaturePublicKeyPEM *string `json:"imageSignaturePublicKeyPEM,omitempty"`
	// Tag representing the software image version that is currently preferred
	ImageTag string `json:"imageTag"`
	// The service environment ID
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// The service ID
	ServiceId string `json:"serviceId"`
}

type _DescribeImageConfigResult DescribeImageConfigResult

// NewDescribeImageConfigResult instantiates a new DescribeImageConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeImageConfigResult(description string, id string, imageName string, imageRegistryId string, imageTag string, serviceEnvironmentId string, serviceId string) *DescribeImageConfigResult {
	this := DescribeImageConfigResult{}
	this.Description = description
	this.Id = id
	this.ImageName = imageName
	this.ImageRegistryId = imageRegistryId
	this.ImageTag = imageTag
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceId = serviceId
	return &this
}

// NewDescribeImageConfigResultWithDefaults instantiates a new DescribeImageConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeImageConfigResultWithDefaults() *DescribeImageConfigResult {
	this := DescribeImageConfigResult{}
	var imageTag string = "latest"
	this.ImageTag = imageTag
	return &this
}

// GetCustomImageCommandsAndArgs returns the CustomImageCommandsAndArgs field value if set, zero value otherwise.
func (o *DescribeImageConfigResult) GetCustomImageCommandsAndArgs() CustomImageCommandsAndArgs {
	if o == nil || IsNil(o.CustomImageCommandsAndArgs) {
		var ret CustomImageCommandsAndArgs
		return ret
	}
	return *o.CustomImageCommandsAndArgs
}

// GetCustomImageCommandsAndArgsOk returns a tuple with the CustomImageCommandsAndArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetCustomImageCommandsAndArgsOk() (*CustomImageCommandsAndArgs, bool) {
	if o == nil || IsNil(o.CustomImageCommandsAndArgs) {
		return nil, false
	}
	return o.CustomImageCommandsAndArgs, true
}

// HasCustomImageCommandsAndArgs returns a boolean if a field has been set.
func (o *DescribeImageConfigResult) HasCustomImageCommandsAndArgs() bool {
	if o != nil && !IsNil(o.CustomImageCommandsAndArgs) {
		return true
	}

	return false
}

// SetCustomImageCommandsAndArgs gets a reference to the given CustomImageCommandsAndArgs and assigns it to the CustomImageCommandsAndArgs field.
func (o *DescribeImageConfigResult) SetCustomImageCommandsAndArgs(v CustomImageCommandsAndArgs) {
	o.CustomImageCommandsAndArgs = &v
}

// GetDescription returns the Description field value
func (o *DescribeImageConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeImageConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeImageConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeImageConfigResult) SetId(v string) {
	o.Id = v
}

// GetImageName returns the ImageName field value
func (o *DescribeImageConfigResult) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *DescribeImageConfigResult) SetImageName(v string) {
	o.ImageName = v
}

// GetImageRegistryId returns the ImageRegistryId field value
func (o *DescribeImageConfigResult) GetImageRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageRegistryId
}

// GetImageRegistryIdOk returns a tuple with the ImageRegistryId field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetImageRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageRegistryId, true
}

// SetImageRegistryId sets field value
func (o *DescribeImageConfigResult) SetImageRegistryId(v string) {
	o.ImageRegistryId = v
}

// GetImageSignaturePublicKeyPEM returns the ImageSignaturePublicKeyPEM field value if set, zero value otherwise.
func (o *DescribeImageConfigResult) GetImageSignaturePublicKeyPEM() string {
	if o == nil || IsNil(o.ImageSignaturePublicKeyPEM) {
		var ret string
		return ret
	}
	return *o.ImageSignaturePublicKeyPEM
}

// GetImageSignaturePublicKeyPEMOk returns a tuple with the ImageSignaturePublicKeyPEM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetImageSignaturePublicKeyPEMOk() (*string, bool) {
	if o == nil || IsNil(o.ImageSignaturePublicKeyPEM) {
		return nil, false
	}
	return o.ImageSignaturePublicKeyPEM, true
}

// HasImageSignaturePublicKeyPEM returns a boolean if a field has been set.
func (o *DescribeImageConfigResult) HasImageSignaturePublicKeyPEM() bool {
	if o != nil && !IsNil(o.ImageSignaturePublicKeyPEM) {
		return true
	}

	return false
}

// SetImageSignaturePublicKeyPEM gets a reference to the given string and assigns it to the ImageSignaturePublicKeyPEM field.
func (o *DescribeImageConfigResult) SetImageSignaturePublicKeyPEM(v string) {
	o.ImageSignaturePublicKeyPEM = &v
}

// GetImageTag returns the ImageTag field value
func (o *DescribeImageConfigResult) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *DescribeImageConfigResult) SetImageTag(v string) {
	o.ImageTag = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *DescribeImageConfigResult) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *DescribeImageConfigResult) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeImageConfigResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeImageConfigResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeImageConfigResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o DescribeImageConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeImageConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomImageCommandsAndArgs) {
		toSerialize["customImageCommandsAndArgs"] = o.CustomImageCommandsAndArgs
	}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["imageName"] = o.ImageName
	toSerialize["imageRegistryId"] = o.ImageRegistryId
	if !IsNil(o.ImageSignaturePublicKeyPEM) {
		toSerialize["imageSignaturePublicKeyPEM"] = o.ImageSignaturePublicKeyPEM
	}
	toSerialize["imageTag"] = o.ImageTag
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *DescribeImageConfigResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"imageName",
		"imageRegistryId",
		"imageTag",
		"serviceEnvironmentId",
		"serviceId",
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

	varDescribeImageConfigResult := _DescribeImageConfigResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeImageConfigResult)

	if err != nil {
		return err
	}

	*o = DescribeImageConfigResult(varDescribeImageConfigResult)

	return err
}

type NullableDescribeImageConfigResult struct {
	value *DescribeImageConfigResult
	isSet bool
}

func (v NullableDescribeImageConfigResult) Get() *DescribeImageConfigResult {
	return v.value
}

func (v *NullableDescribeImageConfigResult) Set(val *DescribeImageConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeImageConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeImageConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeImageConfigResult(val *DescribeImageConfigResult) *NullableDescribeImageConfigResult {
	return &NullableDescribeImageConfigResult{value: val, isSet: true}
}

func (v NullableDescribeImageConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeImageConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


