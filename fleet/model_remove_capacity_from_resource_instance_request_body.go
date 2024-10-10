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

// checks if the RemoveCapacityFromResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveCapacityFromResourceInstanceRequestBody{}

// RemoveCapacityFromResourceInstanceRequestBody struct for RemoveCapacityFromResourceInstanceRequestBody
type RemoveCapacityFromResourceInstanceRequestBody struct {
	// Number of replicas to be removed
	CapacityToBeRemoved int64 `json:"capacityToBeRemoved"`
	// The resource ID.
	ResourceId string `json:"resourceId"`
	AdditionalProperties map[string]interface{}
}

type _RemoveCapacityFromResourceInstanceRequestBody RemoveCapacityFromResourceInstanceRequestBody

// NewRemoveCapacityFromResourceInstanceRequestBody instantiates a new RemoveCapacityFromResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveCapacityFromResourceInstanceRequestBody(capacityToBeRemoved int64, resourceId string) *RemoveCapacityFromResourceInstanceRequestBody {
	this := RemoveCapacityFromResourceInstanceRequestBody{}
	this.CapacityToBeRemoved = capacityToBeRemoved
	this.ResourceId = resourceId
	return &this
}

// NewRemoveCapacityFromResourceInstanceRequestBodyWithDefaults instantiates a new RemoveCapacityFromResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveCapacityFromResourceInstanceRequestBodyWithDefaults() *RemoveCapacityFromResourceInstanceRequestBody {
	this := RemoveCapacityFromResourceInstanceRequestBody{}
	return &this
}

// GetCapacityToBeRemoved returns the CapacityToBeRemoved field value
func (o *RemoveCapacityFromResourceInstanceRequestBody) GetCapacityToBeRemoved() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityToBeRemoved
}

// GetCapacityToBeRemovedOk returns a tuple with the CapacityToBeRemoved field value
// and a boolean to check if the value has been set.
func (o *RemoveCapacityFromResourceInstanceRequestBody) GetCapacityToBeRemovedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityToBeRemoved, true
}

// SetCapacityToBeRemoved sets field value
func (o *RemoveCapacityFromResourceInstanceRequestBody) SetCapacityToBeRemoved(v int64) {
	o.CapacityToBeRemoved = v
}

// GetResourceId returns the ResourceId field value
func (o *RemoveCapacityFromResourceInstanceRequestBody) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *RemoveCapacityFromResourceInstanceRequestBody) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *RemoveCapacityFromResourceInstanceRequestBody) SetResourceId(v string) {
	o.ResourceId = v
}

func (o RemoveCapacityFromResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveCapacityFromResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capacityToBeRemoved"] = o.CapacityToBeRemoved
	toSerialize["resourceId"] = o.ResourceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveCapacityFromResourceInstanceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capacityToBeRemoved",
		"resourceId",
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

	varRemoveCapacityFromResourceInstanceRequestBody := _RemoveCapacityFromResourceInstanceRequestBody{}

	err = json.Unmarshal(data, &varRemoveCapacityFromResourceInstanceRequestBody)

	if err != nil {
		return err
	}

	*o = RemoveCapacityFromResourceInstanceRequestBody(varRemoveCapacityFromResourceInstanceRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capacityToBeRemoved")
		delete(additionalProperties, "resourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveCapacityFromResourceInstanceRequestBody struct {
	value *RemoveCapacityFromResourceInstanceRequestBody
	isSet bool
}

func (v NullableRemoveCapacityFromResourceInstanceRequestBody) Get() *RemoveCapacityFromResourceInstanceRequestBody {
	return v.value
}

func (v *NullableRemoveCapacityFromResourceInstanceRequestBody) Set(val *RemoveCapacityFromResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveCapacityFromResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveCapacityFromResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveCapacityFromResourceInstanceRequestBody(val *RemoveCapacityFromResourceInstanceRequestBody) *NullableRemoveCapacityFromResourceInstanceRequestBody {
	return &NullableRemoveCapacityFromResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableRemoveCapacityFromResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveCapacityFromResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


