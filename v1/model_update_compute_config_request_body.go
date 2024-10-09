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

// checks if the UpdateComputeConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputeConfigRequestBody{}

// UpdateComputeConfigRequestBody struct for UpdateComputeConfigRequestBody
type UpdateComputeConfigRequestBody struct {
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
}

// NewUpdateComputeConfigRequestBody instantiates a new UpdateComputeConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputeConfigRequestBody() *UpdateComputeConfigRequestBody {
	this := UpdateComputeConfigRequestBody{}
	return &this
}

// NewUpdateComputeConfigRequestBodyWithDefaults instantiates a new UpdateComputeConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputeConfigRequestBodyWithDefaults() *UpdateComputeConfigRequestBody {
	this := UpdateComputeConfigRequestBody{}
	return &this
}

// GetAutoscalingPolicy returns the AutoscalingPolicy field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetAutoscalingPolicy() AutoscalingPolicy {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		var ret AutoscalingPolicy
		return ret
	}
	return *o.AutoscalingPolicy
}

// GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool) {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		return nil, false
	}
	return o.AutoscalingPolicy, true
}

// SetAutoscalingPolicy gets a reference to the given AutoscalingPolicy and assigns it to the AutoscalingPolicy field.
func (o *UpdateComputeConfigRequestBody) SetAutoscalingPolicy(v AutoscalingPolicy) {
	o.AutoscalingPolicy = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetCpuArchitecture() string {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret string
		return ret
	}
	return *o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetCpuArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// SetCpuArchitecture gets a reference to the given string and assigns it to the CpuArchitecture field.
func (o *UpdateComputeConfigRequestBody) SetCpuArchitecture(v string) {
	o.CpuArchitecture = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateComputeConfigRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateComputeConfigRequestBody) SetName(v string) {
	o.Name = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetReplicaCount() string {
	if o == nil || IsNil(o.ReplicaCount) {
		var ret string
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetReplicaCountOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaCount) {
		return nil, false
	}
	return o.ReplicaCount, true
}

// SetReplicaCount gets a reference to the given string and assigns it to the ReplicaCount field.
func (o *UpdateComputeConfigRequestBody) SetReplicaCount(v string) {
	o.ReplicaCount = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetResources() ResourceSpec {
	if o == nil || IsNil(o.Resources) {
		var ret ResourceSpec
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetResourcesOk() (*ResourceSpec, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// SetResources gets a reference to the given ResourceSpec and assigns it to the Resources field.
func (o *UpdateComputeConfigRequestBody) SetResources(v ResourceSpec) {
	o.Resources = &v
}

// GetRootVolumeSizeGi returns the RootVolumeSizeGi field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetRootVolumeSizeGi() int64 {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		var ret int64
		return ret
	}
	return *o.RootVolumeSizeGi
}

// GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetRootVolumeSizeGiOk() (*int64, bool) {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		return nil, false
	}
	return o.RootVolumeSizeGi, true
}

// SetRootVolumeSizeGi gets a reference to the given int64 and assigns it to the RootVolumeSizeGi field.
func (o *UpdateComputeConfigRequestBody) SetRootVolumeSizeGi(v int64) {
	o.RootVolumeSizeGi = &v
}

// GetWarmPoolConfiguration returns the WarmPoolConfiguration field value if set, zero value otherwise.
func (o *UpdateComputeConfigRequestBody) GetWarmPoolConfiguration() WarmPoolConfiguration {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		var ret WarmPoolConfiguration
		return ret
	}
	return *o.WarmPoolConfiguration
}

// GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeConfigRequestBody) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool) {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		return nil, false
	}
	return o.WarmPoolConfiguration, true
}

// SetWarmPoolConfiguration gets a reference to the given WarmPoolConfiguration and assigns it to the WarmPoolConfiguration field.
func (o *UpdateComputeConfigRequestBody) SetWarmPoolConfiguration(v WarmPoolConfiguration) {
	o.WarmPoolConfiguration = &v
}

func (o UpdateComputeConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputeConfigRequestBody) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableUpdateComputeConfigRequestBody struct {
	value *UpdateComputeConfigRequestBody
	isSet bool
}

func (v NullableUpdateComputeConfigRequestBody) Get() *UpdateComputeConfigRequestBody {
	return v.value
}

func (v *NullableUpdateComputeConfigRequestBody) Set(val *UpdateComputeConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputeConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputeConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputeConfigRequestBody(val *UpdateComputeConfigRequestBody) *NullableUpdateComputeConfigRequestBody {
	return &NullableUpdateComputeConfigRequestBody{value: val, isSet: true}
}

func (v NullableUpdateComputeConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputeConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

