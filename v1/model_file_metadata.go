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

// checks if the FileMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileMetadata{}

// FileMetadata struct for FileMetadata
type FileMetadata struct {
	// The description of the file
	Description string `json:"description"`
	// The ID of the file
	FileId string `json:"fileId"`
	// The mount path of the file
	MountPath string `json:"mountPath"`
	// The name of the file
	Name string `json:"name"`
	// The size of the file
	Size int64 `json:"size"`
	// The type of the file
	Type string `json:"type"`
	// The time the file was uploaded
	UploadTime string `json:"uploadTime"`
	// The user who uploaded the file
	UploadedBy string `json:"uploadedBy"`
}

type _FileMetadata FileMetadata

// NewFileMetadata instantiates a new FileMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileMetadata(description string, fileId string, mountPath string, name string, size int64, type_ string, uploadTime string, uploadedBy string) *FileMetadata {
	this := FileMetadata{}
	this.Description = description
	this.FileId = fileId
	this.MountPath = mountPath
	this.Name = name
	this.Size = size
	this.Type = type_
	this.UploadTime = uploadTime
	this.UploadedBy = uploadedBy
	return &this
}

// NewFileMetadataWithDefaults instantiates a new FileMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileMetadataWithDefaults() *FileMetadata {
	this := FileMetadata{}
	return &this
}

// GetDescription returns the Description field value
func (o *FileMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FileMetadata) SetDescription(v string) {
	o.Description = v
}

// GetFileId returns the FileId field value
func (o *FileMetadata) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *FileMetadata) SetFileId(v string) {
	o.FileId = v
}

// GetMountPath returns the MountPath field value
func (o *FileMetadata) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *FileMetadata) SetMountPath(v string) {
	o.MountPath = v
}

// GetName returns the Name field value
func (o *FileMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileMetadata) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value
func (o *FileMetadata) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FileMetadata) SetSize(v int64) {
	o.Size = v
}

// GetType returns the Type field value
func (o *FileMetadata) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileMetadata) SetType(v string) {
	o.Type = v
}

// GetUploadTime returns the UploadTime field value
func (o *FileMetadata) GetUploadTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadTime
}

// GetUploadTimeOk returns a tuple with the UploadTime field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetUploadTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadTime, true
}

// SetUploadTime sets field value
func (o *FileMetadata) SetUploadTime(v string) {
	o.UploadTime = v
}

// GetUploadedBy returns the UploadedBy field value
func (o *FileMetadata) GetUploadedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadedBy
}

// GetUploadedByOk returns a tuple with the UploadedBy field value
// and a boolean to check if the value has been set.
func (o *FileMetadata) GetUploadedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadedBy, true
}

// SetUploadedBy sets field value
func (o *FileMetadata) SetUploadedBy(v string) {
	o.UploadedBy = v
}

func (o FileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["fileId"] = o.FileId
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	toSerialize["uploadTime"] = o.UploadTime
	toSerialize["uploadedBy"] = o.UploadedBy
	return toSerialize, nil
}

func (o *FileMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"fileId",
		"mountPath",
		"name",
		"size",
		"type",
		"uploadTime",
		"uploadedBy",
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

	varFileMetadata := _FileMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileMetadata)

	if err != nil {
		return err
	}

	*o = FileMetadata(varFileMetadata)

	return err
}

type NullableFileMetadata struct {
	value *FileMetadata
	isSet bool
}

func (v NullableFileMetadata) Get() *FileMetadata {
	return v.value
}

func (v *NullableFileMetadata) Set(val *FileMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFileMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFileMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileMetadata(val *FileMetadata) *NullableFileMetadata {
	return &NullableFileMetadata{value: val, isSet: true}
}

func (v NullableFileMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

