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

// checks if the ListServiceWorkflowsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceWorkflowsResult{}

// ListServiceWorkflowsResult struct for ListServiceWorkflowsResult
type ListServiceWorkflowsResult struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// List of workflows.
	Workflows []ServiceWorkflow `json:"workflows"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceWorkflowsResult ListServiceWorkflowsResult

// NewListServiceWorkflowsResult instantiates a new ListServiceWorkflowsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceWorkflowsResult(environmentId string, serviceId string, workflows []ServiceWorkflow) *ListServiceWorkflowsResult {
	this := ListServiceWorkflowsResult{}
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	this.Workflows = workflows
	return &this
}

// NewListServiceWorkflowsResultWithDefaults instantiates a new ListServiceWorkflowsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceWorkflowsResultWithDefaults() *ListServiceWorkflowsResult {
	this := ListServiceWorkflowsResult{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *ListServiceWorkflowsResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *ListServiceWorkflowsResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceWorkflowsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListServiceWorkflowsResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceWorkflowsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetServiceId returns the ServiceId field value
func (o *ListServiceWorkflowsResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListServiceWorkflowsResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetWorkflows returns the Workflows field value
func (o *ListServiceWorkflowsResult) GetWorkflows() []ServiceWorkflow {
	if o == nil {
		var ret []ServiceWorkflow
		return ret
	}

	return o.Workflows
}

// GetWorkflowsOk returns a tuple with the Workflows field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsResult) GetWorkflowsOk() ([]ServiceWorkflow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workflows, true
}

// SetWorkflows sets field value
func (o *ListServiceWorkflowsResult) SetWorkflows(v []ServiceWorkflow) {
	o.Workflows = v
}

func (o ListServiceWorkflowsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceWorkflowsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["workflows"] = o.Workflows

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceWorkflowsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"serviceId",
		"workflows",
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

	varListServiceWorkflowsResult := _ListServiceWorkflowsResult{}

	err = json.Unmarshal(data, &varListServiceWorkflowsResult)

	if err != nil {
		return err
	}

	*o = ListServiceWorkflowsResult(varListServiceWorkflowsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "workflows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceWorkflowsResult struct {
	value *ListServiceWorkflowsResult
	isSet bool
}

func (v NullableListServiceWorkflowsResult) Get() *ListServiceWorkflowsResult {
	return v.value
}

func (v *NullableListServiceWorkflowsResult) Set(val *ListServiceWorkflowsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceWorkflowsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceWorkflowsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceWorkflowsResult(val *ListServiceWorkflowsResult) *NullableListServiceWorkflowsResult {
	return &NullableListServiceWorkflowsResult{value: val, isSet: true}
}

func (v NullableListServiceWorkflowsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceWorkflowsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


