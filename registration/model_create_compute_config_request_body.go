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

// checks if the CreateComputeConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateComputeConfigRequestBody{}

// CreateComputeConfigRequestBody struct for CreateComputeConfigRequestBody
type CreateComputeConfigRequestBody struct {
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy,omitempty"`
	// Processor architecture
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`
	// Description of the compute config
	Description string `json:"description"`
	// Name of the compute config
	Name string `json:"name"`
	// Number of replicas to provision for this logical pool of nodes per instance of the resource
	ReplicaCount *string `json:"replicaCount,omitempty"`
	Resources *ResourceSpec `json:"resources,omitempty"`
	// Size of the root volume in Gi
	RootVolumeSizeGi *int64 `json:"rootVolumeSizeGi,omitempty"`
	WarmPoolConfiguration *WarmPoolConfiguration `json:"warmPoolConfiguration,omitempty"`
}

type _CreateComputeConfigRequestBody CreateComputeConfigRequestBody

// NewCreateComputeConfigRequestBody instantiates a new CreateComputeConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateComputeConfigRequestBody(description string, name string) *CreateComputeConfigRequestBody {
	this := CreateComputeConfigRequestBody{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateComputeConfigRequestBodyWithDefaults instantiates a new CreateComputeConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateComputeConfigRequestBodyWithDefaults() *CreateComputeConfigRequestBody {
	this := CreateComputeConfigRequestBody{}
	return &this
}

// GetAutoscalingPolicy returns the AutoscalingPolicy field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetAutoscalingPolicy() AutoscalingPolicy {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		var ret AutoscalingPolicy
		return ret
	}
	return *o.AutoscalingPolicy
}

// GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool) {
	if o == nil || IsNil(o.AutoscalingPolicy) {
		return nil, false
	}
	return o.AutoscalingPolicy, true
}

// SetAutoscalingPolicy gets a reference to the given AutoscalingPolicy and assigns it to the AutoscalingPolicy field.
func (o *CreateComputeConfigRequestBody) SetAutoscalingPolicy(v AutoscalingPolicy) {
	o.AutoscalingPolicy = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetCpuArchitecture() string {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret string
		return ret
	}
	return *o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetCpuArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// SetCpuArchitecture gets a reference to the given string and assigns it to the CpuArchitecture field.
func (o *CreateComputeConfigRequestBody) SetCpuArchitecture(v string) {
	o.CpuArchitecture = &v
}

// GetDescription returns the Description field value
func (o *CreateComputeConfigRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateComputeConfigRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateComputeConfigRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateComputeConfigRequestBody) SetName(v string) {
	o.Name = v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetReplicaCount() string {
	if o == nil || IsNil(o.ReplicaCount) {
		var ret string
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetReplicaCountOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaCount) {
		return nil, false
	}
	return o.ReplicaCount, true
}

// SetReplicaCount gets a reference to the given string and assigns it to the ReplicaCount field.
func (o *CreateComputeConfigRequestBody) SetReplicaCount(v string) {
	o.ReplicaCount = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetResources() ResourceSpec {
	if o == nil || IsNil(o.Resources) {
		var ret ResourceSpec
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetResourcesOk() (*ResourceSpec, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// SetResources gets a reference to the given ResourceSpec and assigns it to the Resources field.
func (o *CreateComputeConfigRequestBody) SetResources(v ResourceSpec) {
	o.Resources = &v
}

// GetRootVolumeSizeGi returns the RootVolumeSizeGi field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetRootVolumeSizeGi() int64 {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		var ret int64
		return ret
	}
	return *o.RootVolumeSizeGi
}

// GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetRootVolumeSizeGiOk() (*int64, bool) {
	if o == nil || IsNil(o.RootVolumeSizeGi) {
		return nil, false
	}
	return o.RootVolumeSizeGi, true
}

// SetRootVolumeSizeGi gets a reference to the given int64 and assigns it to the RootVolumeSizeGi field.
func (o *CreateComputeConfigRequestBody) SetRootVolumeSizeGi(v int64) {
	o.RootVolumeSizeGi = &v
}

// GetWarmPoolConfiguration returns the WarmPoolConfiguration field value if set, zero value otherwise.
func (o *CreateComputeConfigRequestBody) GetWarmPoolConfiguration() WarmPoolConfiguration {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		var ret WarmPoolConfiguration
		return ret
	}
	return *o.WarmPoolConfiguration
}

// GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputeConfigRequestBody) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool) {
	if o == nil || IsNil(o.WarmPoolConfiguration) {
		return nil, false
	}
	return o.WarmPoolConfiguration, true
}

// SetWarmPoolConfiguration gets a reference to the given WarmPoolConfiguration and assigns it to the WarmPoolConfiguration field.
func (o *CreateComputeConfigRequestBody) SetWarmPoolConfiguration(v WarmPoolConfiguration) {
	o.WarmPoolConfiguration = &v
}

func (o CreateComputeConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateComputeConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoscalingPolicy) {
		toSerialize["autoscalingPolicy"] = o.AutoscalingPolicy
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
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

func (o *CreateComputeConfigRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateComputeConfigRequestBody := _CreateComputeConfigRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateComputeConfigRequestBody)

	if err != nil {
		return err
	}

	*o = CreateComputeConfigRequestBody(varCreateComputeConfigRequestBody)

	return err
}

type NullableCreateComputeConfigRequestBody struct {
	value *CreateComputeConfigRequestBody
	isSet bool
}

func (v NullableCreateComputeConfigRequestBody) Get() *CreateComputeConfigRequestBody {
	return v.value
}

func (v *NullableCreateComputeConfigRequestBody) Set(val *CreateComputeConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateComputeConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateComputeConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateComputeConfigRequestBody(val *CreateComputeConfigRequestBody) *NullableCreateComputeConfigRequestBody {
	return &NullableCreateComputeConfigRequestBody{value: val, isSet: true}
}

func (v NullableCreateComputeConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateComputeConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


