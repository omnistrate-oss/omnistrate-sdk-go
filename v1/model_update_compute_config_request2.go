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

// checks if the UpdateComputeConfigRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputeConfigRequest2{}

// UpdateComputeConfigRequest2 struct for UpdateComputeConfigRequest2
type UpdateComputeConfigRequest2 struct {
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy,omitempty"`
	// Processor architecture
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`
	// Description of the compute config
	Description *string `json:"description,omitempty"`
	// Name of the compute config
	Name *string `json:"name,omitempty"`
	// Number of replicas to provision for this logical pool of nodes per instance of the resource
	ReplicaCount *string `json:"replicaCount,omitempty"`
	Resources *ResourceSpec `json:"resources,omitempty"`
	// Size of the root volume in Gi
	RootVolumeSizeGi *int64 `json:"rootVolumeSizeGi,omitempty"`
	WarmPoolConfiguration *WarmPoolConfiguration `json:"warmPoolConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateComputeConfigRequest2 UpdateComputeConfigRequest2

// NewUpdateComputeConfigRequest2 instantiates a new UpdateComputeConfigRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputeConfigRequest2() *UpdateComputeConfigRequest2 {
	this := UpdateComputeConfigRequest2{}
	return &this
}

// NewUpdateComputeConfigRequest2WithDefaults instantiates a new UpdateComputeConfigRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputeConfigRequest2WithDefaults() *UpdateComputeConfigRequest2 {
	this := UpdateComputeConfigRequest2{}
	return &this
}

// GetAutoscalingPolicy returns the AutoscalingPolicy field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetAutoscalingPolicy() AutoscalingPolicy {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		var ret AutoscalingPolicy
		return ret
	}
	return *o.AutoscalingPolicy
}

// GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool) {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		return nil, false
	}
	return o.AutoscalingPolicy, true
}

// SetAutoscalingPolicy gets a reference to the given AutoscalingPolicy and assigns it to the AutoscalingPolicy field.
func (o *UpdateComputeConfigRequest2) SetAutoscalingPolicy(v AutoscalingPolicy) {
	o.AutoscalingPolicy = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetCpuArchitecture() string {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret string
		return ret
	}
	return *o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetCpuArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// SetCpuArchitecture gets a reference to the given string and assigns it to the CpuArchitecture field.
func (o *UpdateComputeConfigRequest2) SetCpuArchitecture(v string) {
	o.CpuArchitecture = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateComputeConfigRequest2) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateComputeConfigRequest2) SetName(v string) {
	o.Name = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetReplicaCount() string {
	if o == nil || IsNil(o.ReplicaCount) {
		var ret string
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetReplicaCountOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaCount) {
		return nil, false
	}
	return o.ReplicaCount, true
}

// SetReplicaCount gets a reference to the given string and assigns it to the ReplicaCount field.
func (o *UpdateComputeConfigRequest2) SetReplicaCount(v string) {
	o.ReplicaCount = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetResources() ResourceSpec {
	if o == nil || IsNil(o.Resources) {
		var ret ResourceSpec
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetResourcesOk() (*ResourceSpec, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// SetResources gets a reference to the given ResourceSpec and assigns it to the Resources field.
func (o *UpdateComputeConfigRequest2) SetResources(v ResourceSpec) {
	o.Resources = &v
}

// GetRootVolumeSizeGi returns the RootVolumeSizeGi field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetRootVolumeSizeGi() int64 {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		var ret int64
		return ret
	}
	return *o.RootVolumeSizeGi
}

// GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetRootVolumeSizeGiOk() (*int64, bool) {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		return nil, false
	}
	return o.RootVolumeSizeGi, true
}

// SetRootVolumeSizeGi gets a reference to the given int64 and assigns it to the RootVolumeSizeGi field.
func (o *UpdateComputeConfigRequest2) SetRootVolumeSizeGi(v int64) {
	o.RootVolumeSizeGi = &v
}

// GetWarmPoolConfiguration returns the WarmPoolConfiguration field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequest2) GetWarmPoolConfiguration() WarmPoolConfiguration {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		var ret WarmPoolConfiguration
		return ret
	}
	return *o.WarmPoolConfiguration
}

// GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequest2) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool) {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		return nil, false
	}
	return o.WarmPoolConfiguration, true
}

// SetWarmPoolConfiguration gets a reference to the given WarmPoolConfiguration and assigns it to the WarmPoolConfiguration field.
func (o *UpdateComputeConfigRequest2) SetWarmPoolConfiguration(v WarmPoolConfiguration) {
	o.WarmPoolConfiguration = &v
}

func (o UpdateComputeConfigRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputeConfigRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoscalingPolicy) {
		toSerialize["autoscalingPolicy"] = o.AutoscalingPolicy
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReplicaCount) {
		toSerialize["replicaCount"] = o.ReplicaCount
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.RootVolumeSizeGi) {
		toSerialize["rootVolumeSizeGi"] = o.RootVolumeSizeGi
	}
	if !IsNil(o.WarmPoolConfiguration) {
		toSerialize["warmPoolConfiguration"] = o.WarmPoolConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateComputeConfigRequest2) UnmarshalJSON(data []byte) (err error) {
	varUpdateComputeConfigRequest2 := _UpdateComputeConfigRequest2{}

	err = json.Unmarshal(data, &varUpdateComputeConfigRequest2)

	if err != nil {
		return err
	}

	*o = UpdateComputeConfigRequest2(varUpdateComputeConfigRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoscalingPolicy")
		delete(additionalProperties, "cpuArchitecture")
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "replicaCount")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "rootVolumeSizeGi")
		delete(additionalProperties, "warmPoolConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateComputeConfigRequest2 struct {
	value *UpdateComputeConfigRequest2
	isSet bool
}

func (v NullableUpdateComputeConfigRequest2) Get() *UpdateComputeConfigRequest2 {
	return v.value
}

func (v *NullableUpdateComputeConfigRequest2) Set(val *UpdateComputeConfigRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputeConfigRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputeConfigRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputeConfigRequest2(val *UpdateComputeConfigRequest2) *NullableUpdateComputeConfigRequest2 {
	return &NullableUpdateComputeConfigRequest2{value: val, isSet: true}
}

func (v NullableUpdateComputeConfigRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputeConfigRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


