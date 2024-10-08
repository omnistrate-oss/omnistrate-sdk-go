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

// checks if the CreateStorageVolumeConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStorageVolumeConfigRequestBody{}

// CreateStorageVolumeConfigRequestBody struct for CreateStorageVolumeConfigRequestBody
type CreateStorageVolumeConfigRequestBody struct {
	// The type of the fixed storage for the cluster
	ClusterStorageType *string `json:"clusterStorageType,omitempty"`
	// A brief description of the context for the storage volume pool
	Description string `json:"description"`
	// Disable backup for the storage volume
	DisableBackup *bool `json:"disableBackup,omitempty"`
	// The IOPS provisioned for the configured instance storage type
	InstanceStorageIops *string `json:"instanceStorageIops,omitempty"`
	// The storage size (in Gi) provisioned for the configured instance storage type
	InstanceStorageSizeGi *string `json:"instanceStorageSizeGi,omitempty"`
	// The throughput (in MiBps) provisioned for the configured instance storage type
	InstanceStorageThroughputMiBps *string `json:"instanceStorageThroughputMiBps,omitempty"`
	// The type of the storage for a compute instance
	InstanceStorageType *string `json:"instanceStorageType,omitempty"`
	// Name of the storage volume pool
	Name string `json:"name"`
	// The storage resource ID
	StorageResourceID *string `json:"storageResourceID,omitempty"`
}

type _CreateStorageVolumeConfigRequestBody CreateStorageVolumeConfigRequestBody

// NewCreateStorageVolumeConfigRequestBody instantiates a new CreateStorageVolumeConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStorageVolumeConfigRequestBody(description string, name string) *CreateStorageVolumeConfigRequestBody {
	this := CreateStorageVolumeConfigRequestBody{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateStorageVolumeConfigRequestBodyWithDefaults instantiates a new CreateStorageVolumeConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStorageVolumeConfigRequestBodyWithDefaults() *CreateStorageVolumeConfigRequestBody {
	this := CreateStorageVolumeConfigRequestBody{}
	return &this
}

// GetClusterStorageType returns the ClusterStorageType field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetClusterStorageType() string {
	if o == nil || IsNil(o.ClusterStorageType) {
		var ret string
		return ret
	}
	return *o.ClusterStorageType
}

// GetClusterStorageTypeOk returns a tuple with the ClusterStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetClusterStorageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterStorageType) {
		return nil, false
	}
	return o.ClusterStorageType, true
}

// HasClusterStorageType returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasClusterStorageType() bool {
	if o != nil && !IsNil(o.ClusterStorageType) {
		return true
	}

	return false
}

// SetClusterStorageType gets a reference to the given string and assigns it to the ClusterStorageType field.
func (o *CreateStorageVolumeConfigRequestBody) SetClusterStorageType(v string) {
	o.ClusterStorageType = &v
}

// GetDescription returns the Description field value
func (o *CreateStorageVolumeConfigRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateStorageVolumeConfigRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetDisableBackup returns the DisableBackup field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetDisableBackup() bool {
	if o == nil || IsNil(o.DisableBackup) {
		var ret bool
		return ret
	}
	return *o.DisableBackup
}

// GetDisableBackupOk returns a tuple with the DisableBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetDisableBackupOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableBackup) {
		return nil, false
	}
	return o.DisableBackup, true
}

// HasDisableBackup returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasDisableBackup() bool {
	if o != nil && !IsNil(o.DisableBackup) {
		return true
	}

	return false
}

// SetDisableBackup gets a reference to the given bool and assigns it to the DisableBackup field.
func (o *CreateStorageVolumeConfigRequestBody) SetDisableBackup(v bool) {
	o.DisableBackup = &v
}

// GetInstanceStorageIops returns the InstanceStorageIops field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageIops() string {
	if o == nil || IsNil(o.InstanceStorageIops) {
		var ret string
		return ret
	}
	return *o.InstanceStorageIops
}

// GetInstanceStorageIopsOk returns a tuple with the InstanceStorageIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageIopsOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageIops) {
		return nil, false
	}
	return o.InstanceStorageIops, true
}

// HasInstanceStorageIops returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageIops() bool {
	if o != nil && !IsNil(o.InstanceStorageIops) {
		return true
	}

	return false
}

// SetInstanceStorageIops gets a reference to the given string and assigns it to the InstanceStorageIops field.
func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageIops(v string) {
	o.InstanceStorageIops = &v
}

// GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageSizeGi() string {
	if o == nil || IsNil(o.InstanceStorageSizeGi) {
		var ret string
		return ret
	}
	return *o.InstanceStorageSizeGi
}

// GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageSizeGiOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageSizeGi) {
		return nil, false
	}
	return o.InstanceStorageSizeGi, true
}

// HasInstanceStorageSizeGi returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageSizeGi() bool {
	if o != nil && !IsNil(o.InstanceStorageSizeGi) {
		return true
	}

	return false
}

// SetInstanceStorageSizeGi gets a reference to the given string and assigns it to the InstanceStorageSizeGi field.
func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageSizeGi(v string) {
	o.InstanceStorageSizeGi = &v
}

// GetInstanceStorageThroughputMiBps returns the InstanceStorageThroughputMiBps field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageThroughputMiBps() string {
	if o == nil || IsNil(o.InstanceStorageThroughputMiBps) {
		var ret string
		return ret
	}
	return *o.InstanceStorageThroughputMiBps
}

// GetInstanceStorageThroughputMiBpsOk returns a tuple with the InstanceStorageThroughputMiBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageThroughputMiBpsOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageThroughputMiBps) {
		return nil, false
	}
	return o.InstanceStorageThroughputMiBps, true
}

// HasInstanceStorageThroughputMiBps returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageThroughputMiBps() bool {
	if o != nil && !IsNil(o.InstanceStorageThroughputMiBps) {
		return true
	}

	return false
}

// SetInstanceStorageThroughputMiBps gets a reference to the given string and assigns it to the InstanceStorageThroughputMiBps field.
func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageThroughputMiBps(v string) {
	o.InstanceStorageThroughputMiBps = &v
}

// GetInstanceStorageType returns the InstanceStorageType field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageType() string {
	if o == nil || IsNil(o.InstanceStorageType) {
		var ret string
		return ret
	}
	return *o.InstanceStorageType
}

// GetInstanceStorageTypeOk returns a tuple with the InstanceStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetInstanceStorageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceStorageType) {
		return nil, false
	}
	return o.InstanceStorageType, true
}

// HasInstanceStorageType returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasInstanceStorageType() bool {
	if o != nil && !IsNil(o.InstanceStorageType) {
		return true
	}

	return false
}

// SetInstanceStorageType gets a reference to the given string and assigns it to the InstanceStorageType field.
func (o *CreateStorageVolumeConfigRequestBody) SetInstanceStorageType(v string) {
	o.InstanceStorageType = &v
}

// GetName returns the Name field value
func (o *CreateStorageVolumeConfigRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStorageVolumeConfigRequestBody) SetName(v string) {
	o.Name = v
}

// GetStorageResourceID returns the StorageResourceID field value if set, zero value otherwise.
func (o *CreateStorageVolumeConfigRequestBody) GetStorageResourceID() string {
	if o == nil || IsNil(o.StorageResourceID) {
		var ret string
		return ret
	}
	return *o.StorageResourceID
}

// GetStorageResourceIDOk returns a tuple with the StorageResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorageVolumeConfigRequestBody) GetStorageResourceIDOk() (*string, bool) {
	if o == nil || IsNil(o.StorageResourceID) {
		return nil, false
	}
	return o.StorageResourceID, true
}

// HasStorageResourceID returns a boolean if a field has been set.
func (o *CreateStorageVolumeConfigRequestBody) HasStorageResourceID() bool {
	if o != nil && !IsNil(o.StorageResourceID) {
		return true
	}

	return false
}

// SetStorageResourceID gets a reference to the given string and assigns it to the StorageResourceID field.
func (o *CreateStorageVolumeConfigRequestBody) SetStorageResourceID(v string) {
	o.StorageResourceID = &v
}

func (o CreateStorageVolumeConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStorageVolumeConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterStorageType) {
		toSerialize["clusterStorageType"] = o.ClusterStorageType
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DisableBackup) {
		toSerialize["disableBackup"] = o.DisableBackup
	}
	if !IsNil(o.InstanceStorageIops) {
		toSerialize["instanceStorageIops"] = o.InstanceStorageIops
	}
	if !IsNil(o.InstanceStorageSizeGi) {
		toSerialize["instanceStorageSizeGi"] = o.InstanceStorageSizeGi
	}
	if !IsNil(o.InstanceStorageThroughputMiBps) {
		toSerialize["instanceStorageThroughputMiBps"] = o.InstanceStorageThroughputMiBps
	}
	if !IsNil(o.InstanceStorageType) {
		toSerialize["instanceStorageType"] = o.InstanceStorageType
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.StorageResourceID) {
		toSerialize["storageResourceID"] = o.StorageResourceID
	}
	return toSerialize, nil
}

func (o *CreateStorageVolumeConfigRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
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

	varCreateStorageVolumeConfigRequestBody := _CreateStorageVolumeConfigRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStorageVolumeConfigRequestBody)

	if err != nil {
		return err
	}

	*o = CreateStorageVolumeConfigRequestBody(varCreateStorageVolumeConfigRequestBody)

	return err
}

type NullableCreateStorageVolumeConfigRequestBody struct {
	value *CreateStorageVolumeConfigRequestBody
	isSet bool
}

func (v NullableCreateStorageVolumeConfigRequestBody) Get() *CreateStorageVolumeConfigRequestBody {
	return v.value
}

func (v *NullableCreateStorageVolumeConfigRequestBody) Set(val *CreateStorageVolumeConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStorageVolumeConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStorageVolumeConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStorageVolumeConfigRequestBody(val *CreateStorageVolumeConfigRequestBody) *NullableCreateStorageVolumeConfigRequestBody {
	return &NullableCreateStorageVolumeConfigRequestBody{value: val, isSet: true}
}

func (v NullableCreateStorageVolumeConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStorageVolumeConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


