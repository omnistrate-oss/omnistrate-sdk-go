/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetWorkflowEventsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWorkflowEventsResult{}

// GetWorkflowEventsResult struct for GetWorkflowEventsResult
type GetWorkflowEventsResult struct {
	// The service environment ID this workflow belongs to.
	EnvironmentId string `json:"environmentId"`
	// ID of the ServiceWorkflow
	Id string `json:"id"`
	// List of resources with workflow events.
	Resources []EventsPerResource `json:"resources"`
	// The service ID this workflow belongs to.
	ServiceId string `json:"serviceId"`
}

type _GetWorkflowEventsResult GetWorkflowEventsResult

// NewGetWorkflowEventsResult instantiates a new GetWorkflowEventsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWorkflowEventsResult(environmentId string, id string, resources []EventsPerResource, serviceId string) *GetWorkflowEventsResult {
	this := GetWorkflowEventsResult{}
	this.EnvironmentId = environmentId
	this.Id = id
	this.Resources = resources
	this.ServiceId = serviceId
	return &this
}

// NewGetWorkflowEventsResultWithDefaults instantiates a new GetWorkflowEventsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWorkflowEventsResultWithDefaults() *GetWorkflowEventsResult {
	this := GetWorkflowEventsResult{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *GetWorkflowEventsResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowEventsResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *GetWorkflowEventsResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *GetWorkflowEventsResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowEventsResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetWorkflowEventsResult) SetId(v string) {
	o.Id = v
}

// GetResources returns the Resources field value
func (o *GetWorkflowEventsResult) GetResources() []EventsPerResource {
	if o == nil {
		var ret []EventsPerResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowEventsResult) GetResourcesOk() ([]EventsPerResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *GetWorkflowEventsResult) SetResources(v []EventsPerResource) {
	o.Resources = v
}

// GetServiceId returns the ServiceId field value
func (o *GetWorkflowEventsResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowEventsResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *GetWorkflowEventsResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o GetWorkflowEventsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWorkflowEventsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["id"] = o.Id
	toSerialize["resources"] = o.Resources
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *GetWorkflowEventsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"id",
		"resources",
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

	varGetWorkflowEventsResult := _GetWorkflowEventsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetWorkflowEventsResult)

	if err != nil {
		return err
	}

	*o = GetWorkflowEventsResult(varGetWorkflowEventsResult)

	return err
}

type NullableGetWorkflowEventsResult struct {
	value *GetWorkflowEventsResult
	isSet bool
}

func (v NullableGetWorkflowEventsResult) Get() *GetWorkflowEventsResult {
	return v.value
}

func (v *NullableGetWorkflowEventsResult) Set(val *GetWorkflowEventsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWorkflowEventsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWorkflowEventsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWorkflowEventsResult(val *GetWorkflowEventsResult) *NullableGetWorkflowEventsResult {
	return &NullableGetWorkflowEventsResult{value: val, isSet: true}
}

func (v NullableGetWorkflowEventsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWorkflowEventsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


