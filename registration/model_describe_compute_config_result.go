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

// checks if the DescribeComputeConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeComputeConfigResult{}

// DescribeComputeConfigResult struct for DescribeComputeConfigResult
type DescribeComputeConfigResult struct {
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy,omitempty"`
	// Processor architecture
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`
	// Description of the compute config
	Description string `json:"description"`
	// ID of the compute config
	Id string `json:"id"`
	// The list of infra config IDs associated with the compute config.
	InfraConfigIDs []string `json:"infraConfigIDs,omitempty"`
	// The instance types for this compute config
	InstanceTypes *map[string][]string `json:"instanceTypes,omitempty"`
	// Name of the compute config
	Name string `json:"name"`
	// Number of replicas to provision for this logical pool of nodes per instance of the resource
	ReplicaCount string `json:"replicaCount"`
	Resources *ResourceSpec `json:"resources,omitempty"`
	// Size of the root volume in Gi
	RootVolumeSizeGi *int64 `json:"rootVolumeSizeGi,omitempty"`
	// The service ID
	ServiceId string `json:"serviceId"`
	WarmPoolConfiguration *WarmPoolConfiguration `json:"warmPoolConfiguration,omitempty"`
}

type _DescribeComputeConfigResult DescribeComputeConfigResult

// NewDescribeComputeConfigResult instantiates a new DescribeComputeConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeComputeConfigResult(description string, id string, name string, replicaCount string, serviceId string) *DescribeComputeConfigResult {
	this := DescribeComputeConfigResult{}
	this.Description = description
	this.Id = id
	this.Name = name
	this.ReplicaCount = replicaCount
	this.ServiceId = serviceId
	return &this
}

// NewDescribeComputeConfigResultWithDefaults instantiates a new DescribeComputeConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeComputeConfigResultWithDefaults() *DescribeComputeConfigResult {
	this := DescribeComputeConfigResult{}
	return &this
}

// GetAutoscalingPolicy returns the AutoscalingPolicy field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetAutoscalingPolicy() AutoscalingPolicy {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		var ret AutoscalingPolicy
		return ret
	}
	return *o.AutoscalingPolicy
}

// GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool) {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		return nil, false
	}
	return o.AutoscalingPolicy, true
}

// HasAutoscalingPolicy returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasAutoscalingPolicy() bool {
	if o != nil && !IsNil(o.AutoscalingPolicy) {
		return true
	}

	return false
}

// SetAutoscalingPolicy gets a reference to the given AutoscalingPolicy and assigns it to the AutoscalingPolicy field.
func (o *DescribeComputeConfigResult) SetAutoscalingPolicy(v AutoscalingPolicy) {
	o.AutoscalingPolicy = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetCpuArchitecture() string {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret string
		return ret
	}
	return *o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetCpuArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// HasCpuArchitecture returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasCpuArchitecture() bool {
	if o != nil && !IsNil(o.CpuArchitecture) {
		return true
	}

	return false
}

// SetCpuArchitecture gets a reference to the given string and assigns it to the CpuArchitecture field.
func (o *DescribeComputeConfigResult) SetCpuArchitecture(v string) {
	o.CpuArchitecture = &v
}

// GetDescription returns the Description field value
func (o *DescribeComputeConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeComputeConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeComputeConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeComputeConfigResult) SetId(v string) {
	o.Id = v
}

// GetInfraConfigIDs returns the InfraConfigIDs field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetInfraConfigIDs() []string {
	if o == nil || IsNil(o.InfraConfigIDs) {
		var ret []string
		return ret
	}
	return o.InfraConfigIDs
}

// GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetInfraConfigIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.InfraConfigIDs) {
		return nil, false
	}
	return o.InfraConfigIDs, true
}

// HasInfraConfigIDs returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasInfraConfigIDs() bool {
	if o != nil && !IsNil(o.InfraConfigIDs) {
		return true
	}

	return false
}

// SetInfraConfigIDs gets a reference to the given []string and assigns it to the InfraConfigIDs field.
func (o *DescribeComputeConfigResult) SetInfraConfigIDs(v []string) {
	o.InfraConfigIDs = v
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetInstanceTypes() map[string][]string {
	if o == nil || IsNil(o.InstanceTypes) {
		var ret map[string][]string
		return ret
	}
	return *o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetInstanceTypesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.InstanceTypes) {
		return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasInstanceTypes() bool {
	if o != nil && !IsNil(o.InstanceTypes) {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given map[string][]string and assigns it to the InstanceTypes field.
func (o *DescribeComputeConfigResult) SetInstanceTypes(v map[string][]string) {
	o.InstanceTypes = &v
}

// GetName returns the Name field value
func (o *DescribeComputeConfigResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeComputeConfigResult) SetName(v string) {
	o.Name = v
}

// GetReplicaCount returns the ReplicaCount field value
func (o *DescribeComputeConfigResult) GetReplicaCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetReplicaCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaCount, true
}

// SetReplicaCount sets field value
func (o *DescribeComputeConfigResult) SetReplicaCount(v string) {
	o.ReplicaCount = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetResources() ResourceSpec {
	if o == nil || IsNil(o.Resources) {
		var ret ResourceSpec
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetResourcesOk() (*ResourceSpec, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given ResourceSpec and assigns it to the Resources field.
func (o *DescribeComputeConfigResult) SetResources(v ResourceSpec) {
	o.Resources = &v
}

// GetRootVolumeSizeGi returns the RootVolumeSizeGi field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetRootVolumeSizeGi() int64 {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		var ret int64
		return ret
	}
	return *o.RootVolumeSizeGi
}

// GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetRootVolumeSizeGiOk() (*int64, bool) {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		return nil, false
	}
	return o.RootVolumeSizeGi, true
}

// HasRootVolumeSizeGi returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasRootVolumeSizeGi() bool {
	if o != nil && !IsNil(o.RootVolumeSizeGi) {
		return true
	}

	return false
}

// SetRootVolumeSizeGi gets a reference to the given int64 and assigns it to the RootVolumeSizeGi field.
func (o *DescribeComputeConfigResult) SetRootVolumeSizeGi(v int64) {
	o.RootVolumeSizeGi = &v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeComputeConfigResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeComputeConfigResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetWarmPoolConfiguration returns the WarmPoolConfiguration field value if set, zero value otherwise.
func (o *DescribeComputeConfigResult) GetWarmPoolConfiguration() WarmPoolConfiguration {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		var ret WarmPoolConfiguration
		return ret
	}
	return *o.WarmPoolConfiguration
}

// GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeComputeConfigResult) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool) {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		return nil, false
	}
	return o.WarmPoolConfiguration, true
}

// HasWarmPoolConfiguration returns a boolean if a field has been set.
func (o *DescribeComputeConfigResult) HasWarmPoolConfiguration() bool {
	if o != nil && !IsNil(o.WarmPoolConfiguration) {
		return true
	}

	return false
}

// SetWarmPoolConfiguration gets a reference to the given WarmPoolConfiguration and assigns it to the WarmPoolConfiguration field.
func (o *DescribeComputeConfigResult) SetWarmPoolConfiguration(v WarmPoolConfiguration) {
	o.WarmPoolConfiguration = &v
}

func (o DescribeComputeConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeComputeConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoscalingPolicy) {
		toSerialize["autoscalingPolicy"] = o.AutoscalingPolicy
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	if !IsNil(o.InfraConfigIDs) {
		toSerialize["infraConfigIDs"] = o.InfraConfigIDs
	}
	if !IsNil(o.InstanceTypes) {
		toSerialize["instanceTypes"] = o.InstanceTypes
	}
	toSerialize["name"] = o.Name
	toSerialize["replicaCount"] = o.ReplicaCount
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.RootVolumeSizeGi) {
		toSerialize["rootVolumeSizeGi"] = o.RootVolumeSizeGi
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.WarmPoolConfiguration) {
		toSerialize["warmPoolConfiguration"] = o.WarmPoolConfiguration
	}
	return toSerialize, nil
}

func (o *DescribeComputeConfigResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"name",
		"replicaCount",
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

	varDescribeComputeConfigResult := _DescribeComputeConfigResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeComputeConfigResult)

	if err != nil {
		return err
	}

	*o = DescribeComputeConfigResult(varDescribeComputeConfigResult)

	return err
}

type NullableDescribeComputeConfigResult struct {
	value *DescribeComputeConfigResult
	isSet bool
}

func (v NullableDescribeComputeConfigResult) Get() *DescribeComputeConfigResult {
	return v.value
}

func (v *NullableDescribeComputeConfigResult) Set(val *DescribeComputeConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeComputeConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeComputeConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeComputeConfigResult(val *DescribeComputeConfigResult) *NullableDescribeComputeConfigResult {
	return &NullableDescribeComputeConfigResult{value: val, isSet: true}
}

func (v NullableDescribeComputeConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeComputeConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


