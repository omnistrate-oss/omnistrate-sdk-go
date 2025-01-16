/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the DescribeServiceWorkflowResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceWorkflowResult{}

// DescribeServiceWorkflowResult struct for DescribeServiceWorkflowResult
type DescribeServiceWorkflowResult struct {
	// List of resources with deployment status.
	Resources []ResourceDeploymentStatus `json:"Resources,omitempty"`
	Workflow ServiceWorkflow `json:"Workflow"`
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceWorkflowResult DescribeServiceWorkflowResult

// NewDescribeServiceWorkflowResult instantiates a new DescribeServiceWorkflowResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceWorkflowResult(workflow ServiceWorkflow, environmentId string, serviceId string) *DescribeServiceWorkflowResult {
	this := DescribeServiceWorkflowResult{}
	this.Workflow = workflow
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	return &this
}

// NewDescribeServiceWorkflowResultWithDefaults instantiates a new DescribeServiceWorkflowResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceWorkflowResultWithDefaults() *DescribeServiceWorkflowResult {
	this := DescribeServiceWorkflowResult{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DescribeServiceWorkflowResult) GetResources() []ResourceDeploymentStatus {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourceDeploymentStatus
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowResult) GetResourcesOk() ([]ResourceDeploymentStatus, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DescribeServiceWorkflowResult) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceDeploymentStatus and assigns it to the Resources field.
func (o *DescribeServiceWorkflowResult) SetResources(v []ResourceDeploymentStatus) {
	o.Resources = v
}

// GetWorkflow returns the Workflow field value
func (o *DescribeServiceWorkflowResult) GetWorkflow() ServiceWorkflow {
	if o == nil {
		var ret ServiceWorkflow
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowResult) GetWorkflowOk() (*ServiceWorkflow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *DescribeServiceWorkflowResult) SetWorkflow(v ServiceWorkflow) {
	o.Workflow = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *DescribeServiceWorkflowResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *DescribeServiceWorkflowResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceWorkflowResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceWorkflowResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceWorkflowResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o DescribeServiceWorkflowResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceWorkflowResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["Resources"] = o.Resources
	}
	toSerialize["Workflow"] = o.Workflow
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["serviceId"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceWorkflowResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Workflow",
		"environmentId",
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

	varDescribeServiceWorkflowResult := _DescribeServiceWorkflowResult{}

	err = json.Unmarshal(data, &varDescribeServiceWorkflowResult)

	if err != nil {
		return err
	}

	*o = DescribeServiceWorkflowResult(varDescribeServiceWorkflowResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Resources")
		delete(additionalProperties, "Workflow")
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "serviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceWorkflowResult struct {
	value *DescribeServiceWorkflowResult
	isSet bool
}

func (v NullableDescribeServiceWorkflowResult) Get() *DescribeServiceWorkflowResult {
	return v.value
}

func (v *NullableDescribeServiceWorkflowResult) Set(val *DescribeServiceWorkflowResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceWorkflowResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceWorkflowResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceWorkflowResult(val *DescribeServiceWorkflowResult) *NullableDescribeServiceWorkflowResult {
	return &NullableDescribeServiceWorkflowResult{value: val, isSet: true}
}

func (v NullableDescribeServiceWorkflowResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceWorkflowResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


