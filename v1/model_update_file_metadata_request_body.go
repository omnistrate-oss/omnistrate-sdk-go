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

// checks if the UpdateFileMetadataRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileMetadataRequestBody{}

// UpdateFileMetadataRequestBody struct for UpdateFileMetadataRequestBody
type UpdateFileMetadataRequestBody struct {
	// The description of the file
	Description *string `json:"description,omitempty"`
	// The mount path of the file
	MountPath *string `json:"mountPath,omitempty"`
	// The name of the file
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFileMetadataRequestBody UpdateFileMetadataRequestBody

// NewUpdateFileMetadataRequestBody instantiates a new UpdateFileMetadataRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileMetadataRequestBody() *UpdateFileMetadataRequestBody {
	this := UpdateFileMetadataRequestBody{}
	return &this
}

// NewUpdateFileMetadataRequestBodyWithDefaults instantiates a new UpdateFileMetadataRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileMetadataRequestBodyWithDefaults() *UpdateFileMetadataRequestBody {
	this := UpdateFileMetadataRequestBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFileMetadataRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileMetadataRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFileMetadataRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *UpdateFileMetadataRequestBody) GetMountPath() string {
	if o == nil || IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileMetadataRequestBody) GetMountPathOk() (*string, bool) {
	if o == nil || IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath gets a reference to the given string and assigns it to the MountPath field.
func (o *UpdateFileMetadataRequestBody) SetMountPath(v string) {
	o.MountPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFileMetadataRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileMetadataRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFileMetadataRequestBody) SetName(v string) {
	o.Name = &v
}

func (o UpdateFileMetadataRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileMetadataRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MountPath) {
		toSerialize["mountPath"] = o.MountPath
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFileMetadataRequestBody) UnmarshalJSON(data []byte) (err error) {
	varUpdateFileMetadataRequestBody := _UpdateFileMetadataRequestBody{}

	err = json.Unmarshal(data, &varUpdateFileMetadataRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateFileMetadataRequestBody(varUpdateFileMetadataRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "mountPath")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFileMetadataRequestBody struct {
	value *UpdateFileMetadataRequestBody
	isSet bool
}

func (v NullableUpdateFileMetadataRequestBody) Get() *UpdateFileMetadataRequestBody {
	return v.value
}

func (v *NullableUpdateFileMetadataRequestBody) Set(val *UpdateFileMetadataRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFileMetadataRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFileMetadataRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFileMetadataRequestBody(val *UpdateFileMetadataRequestBody) *NullableUpdateFileMetadataRequestBody {
	return &NullableUpdateFileMetadataRequestBody{value: val, isSet: true}
}

func (v NullableUpdateFileMetadataRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFileMetadataRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


