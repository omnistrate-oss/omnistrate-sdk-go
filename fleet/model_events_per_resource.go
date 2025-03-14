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

// checks if the EventsPerResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsPerResource{}

// EventsPerResource struct for EventsPerResource
type EventsPerResource struct {
	// ID of a resource
	ResourceId string `json:"resourceId"`
	// Resource Key
	ResourceKey string `json:"resourceKey"`
	// Resource Name
	ResourceName string `json:"resourceName"`
	// Per step workflow events for the resource
	WorkflowSteps []EventsPerWorkflowStep `json:"workflowSteps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventsPerResource EventsPerResource

// NewEventsPerResource instantiates a new EventsPerResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsPerResource(resourceId string, resourceKey string, resourceName string) *EventsPerResource {
	this := EventsPerResource{}
	this.ResourceId = resourceId
	this.ResourceKey = resourceKey
	this.ResourceName = resourceName
	return &this
}

// NewEventsPerResourceWithDefaults instantiates a new EventsPerResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsPerResourceWithDefaults() *EventsPerResource {
	this := EventsPerResource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *EventsPerResource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *EventsPerResource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *EventsPerResource) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceKey returns the ResourceKey field value
func (o *EventsPerResource) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *EventsPerResource) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *EventsPerResource) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetResourceName returns the ResourceName field value
func (o *EventsPerResource) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *EventsPerResource) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *EventsPerResource) SetResourceName(v string) {
	o.ResourceName = v
}

// GetWorkflowSteps returns the WorkflowSteps field value if set, zero value otherwise.
func (o *EventsPerResource) GetWorkflowSteps() []EventsPerWorkflowStep {
	if o == nil || IsNil(o.WorkflowSteps) {
		var ret []EventsPerWorkflowStep
		return ret
	}
	return o.WorkflowSteps
}

// GetWorkflowStepsOk returns a tuple with the WorkflowSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPerResource) GetWorkflowStepsOk() ([]EventsPerWorkflowStep, bool) {
	if o == nil || IsNil(o.WorkflowSteps) {
		return nil, false
	}
	return o.WorkflowSteps, true
}

// HasWorkflowSteps returns a boolean if a field has been set.
func (o *EventsPerResource) HasWorkflowSteps() bool {
	if o != nil && !IsNil(o.WorkflowSteps) {
		return true
	}

	return false
}

// SetWorkflowSteps gets a reference to the given []EventsPerWorkflowStep and assigns it to the WorkflowSteps field.
func (o *EventsPerResource) SetWorkflowSteps(v []EventsPerWorkflowStep) {
	o.WorkflowSteps = v
}

func (o EventsPerResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsPerResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceKey"] = o.ResourceKey
	toSerialize["resourceName"] = o.ResourceName
	if !IsNil(o.WorkflowSteps) {
		toSerialize["workflowSteps"] = o.WorkflowSteps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventsPerResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
		"resourceKey",
		"resourceName",
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

	varEventsPerResource := _EventsPerResource{}

	err = json.Unmarshal(data, &varEventsPerResource)

	if err != nil {
		return err
	}

	*o = EventsPerResource(varEventsPerResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceKey")
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "workflowSteps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventsPerResource struct {
	value *EventsPerResource
	isSet bool
}

func (v NullableEventsPerResource) Get() *EventsPerResource {
	return v.value
}

func (v *NullableEventsPerResource) Set(val *EventsPerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsPerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsPerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsPerResource(val *EventsPerResource) *NullableEventsPerResource {
	return &NullableEventsPerResource{value: val, isSet: true}
}

func (v NullableEventsPerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsPerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


