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

// checks if the FleetListLinkedInstancesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetListLinkedInstancesResult{}

// FleetListLinkedInstancesResult struct for FleetListLinkedInstancesResult
type FleetListLinkedInstancesResult struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The list of resource instances.
	ResourceInstances []ResourceInstance `json:"resourceInstances"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	AdditionalProperties map[string]interface{}
}

type _FleetListLinkedInstancesResult FleetListLinkedInstancesResult

// NewFleetListLinkedInstancesResult instantiates a new FleetListLinkedInstancesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetListLinkedInstancesResult(environmentId string, resourceInstances []ResourceInstance, serviceId string) *FleetListLinkedInstancesResult {
	this := FleetListLinkedInstancesResult{}
	this.EnvironmentId = environmentId
	this.ResourceInstances = resourceInstances
	this.ServiceId = serviceId
	return &this
}

// NewFleetListLinkedInstancesResultWithDefaults instantiates a new FleetListLinkedInstancesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetListLinkedInstancesResultWithDefaults() *FleetListLinkedInstancesResult {
	this := FleetListLinkedInstancesResult{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetListLinkedInstancesResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetListLinkedInstancesResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *FleetListLinkedInstancesResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *FleetListLinkedInstancesResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *FleetListLinkedInstancesResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetResourceInstances returns the ResourceInstances field value
func (o *FleetListLinkedInstancesResult) GetResourceInstances() []ResourceInstance {
	if o == nil {
		var ret []ResourceInstance
		return ret
	}

	return o.ResourceInstances
}

// GetResourceInstancesOk returns a tuple with the ResourceInstances field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesResult) GetResourceInstancesOk() ([]ResourceInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceInstances, true
}

// SetResourceInstances sets field value
func (o *FleetListLinkedInstancesResult) SetResourceInstances(v []ResourceInstance) {
	o.ResourceInstances = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetListLinkedInstancesResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetListLinkedInstancesResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetListLinkedInstancesResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o FleetListLinkedInstancesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetListLinkedInstancesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["resourceInstances"] = o.ResourceInstances
	toSerialize["serviceId"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetListLinkedInstancesResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"resourceInstances",
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

	varFleetListLinkedInstancesResult := _FleetListLinkedInstancesResult{}

	err = json.Unmarshal(data, &varFleetListLinkedInstancesResult)

	if err != nil {
		return err
	}

	*o = FleetListLinkedInstancesResult(varFleetListLinkedInstancesResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "resourceInstances")
		delete(additionalProperties, "serviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetListLinkedInstancesResult struct {
	value *FleetListLinkedInstancesResult
	isSet bool
}

func (v NullableFleetListLinkedInstancesResult) Get() *FleetListLinkedInstancesResult {
	return v.value
}

func (v *NullableFleetListLinkedInstancesResult) Set(val *FleetListLinkedInstancesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetListLinkedInstancesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetListLinkedInstancesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetListLinkedInstancesResult(val *FleetListLinkedInstancesResult) *NullableFleetListLinkedInstancesResult {
	return &NullableFleetListLinkedInstancesResult{value: val, isSet: true}
}

func (v NullableFleetListLinkedInstancesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetListLinkedInstancesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


