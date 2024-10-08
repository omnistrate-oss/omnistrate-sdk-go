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

// checks if the DescribeStorageConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeStorageConfigResult{}

// DescribeStorageConfigResult struct for DescribeStorageConfigResult
type DescribeStorageConfigResult struct {
	// Description of the storage config
	Description string `json:"description"`
	// The storage config ID
	Id string `json:"id"`
	// The list of infra config IDs associated with the compute config.
	InfraConfigIDs []string `json:"infraConfigIDs,omitempty"`
	// Name of the storage config
	Name string `json:"name"`
	// The service ID
	ServiceId string `json:"serviceId"`
	// The storage volume config IDs and the corresponding mount path
	Volumes map[string][]string `json:"volumes"`
}

type _DescribeStorageConfigResult DescribeStorageConfigResult

// NewDescribeStorageConfigResult instantiates a new DescribeStorageConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeStorageConfigResult(description string, id string, name string, serviceId string, volumes map[string][]string) *DescribeStorageConfigResult {
	this := DescribeStorageConfigResult{}
	this.Description = description
	this.Id = id
	this.Name = name
	this.ServiceId = serviceId
	this.Volumes = volumes
	return &this
}

// NewDescribeStorageConfigResultWithDefaults instantiates a new DescribeStorageConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeStorageConfigResultWithDefaults() *DescribeStorageConfigResult {
	this := DescribeStorageConfigResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *DescribeStorageConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeStorageConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeStorageConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeStorageConfigResult) SetId(v string) {
	o.Id = v
}

// GetInfraConfigIDs returns the InfraConfigIDs field value if set, zero value otherwise.
func (o *DescribeStorageConfigResult) GetInfraConfigIDs() []string {
	if o == nil || IsNil(o.InfraConfigIDs) {
		var ret []string
		return ret
	}
	return o.InfraConfigIDs
}

// GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetInfraConfigIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.InfraConfigIDs) {
		return nil, false
	}
	return o.InfraConfigIDs, true
}

// HasInfraConfigIDs returns a boolean if a field has been set.
func (o *DescribeStorageConfigResult) HasInfraConfigIDs() bool {
	if o != nil && !IsNil(o.InfraConfigIDs) {
		return true
	}

	return false
}

// SetInfraConfigIDs gets a reference to the given []string and assigns it to the InfraConfigIDs field.
func (o *DescribeStorageConfigResult) SetInfraConfigIDs(v []string) {
	o.InfraConfigIDs = v
}

// GetName returns the Name field value
func (o *DescribeStorageConfigResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeStorageConfigResult) SetName(v string) {
	o.Name = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeStorageConfigResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeStorageConfigResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetVolumes returns the Volumes field value
func (o *DescribeStorageConfigResult) GetVolumes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *DescribeStorageConfigResult) GetVolumesOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value
func (o *DescribeStorageConfigResult) SetVolumes(v map[string][]string) {
	o.Volumes = v
}

func (o DescribeStorageConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeStorageConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	if !IsNil(o.InfraConfigIDs) {
		toSerialize["infraConfigIDs"] = o.InfraConfigIDs
	}
	toSerialize["name"] = o.Name
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["volumes"] = o.Volumes
	return toSerialize, nil
}

func (o *DescribeStorageConfigResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"name",
		"serviceId",
		"volumes",
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

	varDescribeStorageConfigResult := _DescribeStorageConfigResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeStorageConfigResult)

	if err != nil {
		return err
	}

	*o = DescribeStorageConfigResult(varDescribeStorageConfigResult)

	return err
}

type NullableDescribeStorageConfigResult struct {
	value *DescribeStorageConfigResult
	isSet bool
}

func (v NullableDescribeStorageConfigResult) Get() *DescribeStorageConfigResult {
	return v.value
}

func (v *NullableDescribeStorageConfigResult) Set(val *DescribeStorageConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeStorageConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeStorageConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeStorageConfigResult(val *DescribeStorageConfigResult) *NullableDescribeStorageConfigResult {
	return &NullableDescribeStorageConfigResult{value: val, isSet: true}
}

func (v NullableDescribeStorageConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeStorageConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


