/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListFleetResourceInstancesResultInternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFleetResourceInstancesResultInternal{}

// ListFleetResourceInstancesResultInternal struct for ListFleetResourceInstancesResultInternal
type ListFleetResourceInstancesResultInternal struct {
	// The service environment ID this workflow belongs to.
	EnvironmentId *string `json:"environmentId,omitempty"`
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The list of resource instances.
	ResourceInstances []ResourceInstance `json:"resourceInstances"`
	// The service ID this workflow belongs to.
	ServiceId *string `json:"serviceId,omitempty"`
}

type _ListFleetResourceInstancesResultInternal ListFleetResourceInstancesResultInternal

// NewListFleetResourceInstancesResultInternal instantiates a new ListFleetResourceInstancesResultInternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFleetResourceInstancesResultInternal(resourceInstances []ResourceInstance) *ListFleetResourceInstancesResultInternal {
	this := ListFleetResourceInstancesResultInternal{}
	this.ResourceInstances = resourceInstances
	return &this
}

// NewListFleetResourceInstancesResultInternalWithDefaults instantiates a new ListFleetResourceInstancesResultInternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFleetResourceInstancesResultInternalWithDefaults() *ListFleetResourceInstancesResultInternal {
	this := ListFleetResourceInstancesResultInternal{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ListFleetResourceInstancesResultInternal) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFleetResourceInstancesResultInternal) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ListFleetResourceInstancesResultInternal) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ListFleetResourceInstancesResultInternal) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListFleetResourceInstancesResultInternal) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFleetResourceInstancesResultInternal) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListFleetResourceInstancesResultInternal) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListFleetResourceInstancesResultInternal) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetResourceInstances returns the ResourceInstances field value
func (o *ListFleetResourceInstancesResultInternal) GetResourceInstances() []ResourceInstance {
	if o == nil {
		var ret []ResourceInstance
		return ret
	}

	return o.ResourceInstances
}

// GetResourceInstancesOk returns a tuple with the ResourceInstances field value
// and a boolean to check if the value has been set.
func (o *ListFleetResourceInstancesResultInternal) GetResourceInstancesOk() ([]ResourceInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceInstances, true
}

// SetResourceInstances sets field value
func (o *ListFleetResourceInstancesResultInternal) SetResourceInstances(v []ResourceInstance) {
	o.ResourceInstances = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ListFleetResourceInstancesResultInternal) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFleetResourceInstancesResultInternal) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ListFleetResourceInstancesResultInternal) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ListFleetResourceInstancesResultInternal) SetServiceId(v string) {
	o.ServiceId = &v
}

func (o ListFleetResourceInstancesResultInternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFleetResourceInstancesResultInternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["resourceInstances"] = o.ResourceInstances
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	return toSerialize, nil
}

func (o *ListFleetResourceInstancesResultInternal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceInstances",
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

	varListFleetResourceInstancesResultInternal := _ListFleetResourceInstancesResultInternal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListFleetResourceInstancesResultInternal)

	if err != nil {
		return err
	}

	*o = ListFleetResourceInstancesResultInternal(varListFleetResourceInstancesResultInternal)

	return err
}

type NullableListFleetResourceInstancesResultInternal struct {
	value *ListFleetResourceInstancesResultInternal
	isSet bool
}

func (v NullableListFleetResourceInstancesResultInternal) Get() *ListFleetResourceInstancesResultInternal {
	return v.value
}

func (v *NullableListFleetResourceInstancesResultInternal) Set(val *ListFleetResourceInstancesResultInternal) {
	v.value = val
	v.isSet = true
}

func (v NullableListFleetResourceInstancesResultInternal) IsSet() bool {
	return v.isSet
}

func (v *NullableListFleetResourceInstancesResultInternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFleetResourceInstancesResultInternal(val *ListFleetResourceInstancesResultInternal) *NullableListFleetResourceInstancesResultInternal {
	return &NullableListFleetResourceInstancesResultInternal{value: val, isSet: true}
}

func (v NullableListFleetResourceInstancesResultInternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFleetResourceInstancesResultInternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

