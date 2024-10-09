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

// checks if the AddCapacityToResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCapacityToResourceInstanceRequestBody{}

// AddCapacityToResourceInstanceRequestBody struct for AddCapacityToResourceInstanceRequestBody
type AddCapacityToResourceInstanceRequestBody struct {
	// Number of replicas to be added
	CapacityToBeAdded int64 `json:"capacityToBeAdded"`
	// The resource ID.
	ResourceId string `json:"resourceId"`
}

type _AddCapacityToResourceInstanceRequestBody AddCapacityToResourceInstanceRequestBody

// NewAddCapacityToResourceInstanceRequestBody instantiates a new AddCapacityToResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCapacityToResourceInstanceRequestBody(capacityToBeAdded int64, resourceId string) *AddCapacityToResourceInstanceRequestBody {
	this := AddCapacityToResourceInstanceRequestBody{}
	this.CapacityToBeAdded = capacityToBeAdded
	this.ResourceId = resourceId
	return &this
}

// NewAddCapacityToResourceInstanceRequestBodyWithDefaults instantiates a new AddCapacityToResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCapacityToResourceInstanceRequestBodyWithDefaults() *AddCapacityToResourceInstanceRequestBody {
	this := AddCapacityToResourceInstanceRequestBody{}
	return &this
}

// GetCapacityToBeAdded returns the CapacityToBeAdded field value
func (o *AddCapacityToResourceInstanceRequestBody) GetCapacityToBeAdded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityToBeAdded
}

// GetCapacityToBeAddedOk returns a tuple with the CapacityToBeAdded field value
// and a boolean to check if the value has been set.
func (o *AddCapacityToResourceInstanceRequestBody) GetCapacityToBeAddedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityToBeAdded, true
}

// SetCapacityToBeAdded sets field value
func (o *AddCapacityToResourceInstanceRequestBody) SetCapacityToBeAdded(v int64) {
	o.CapacityToBeAdded = v
}

// GetResourceId returns the ResourceId field value
func (o *AddCapacityToResourceInstanceRequestBody) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AddCapacityToResourceInstanceRequestBody) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AddCapacityToResourceInstanceRequestBody) SetResourceId(v string) {
	o.ResourceId = v
}

func (o AddCapacityToResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCapacityToResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capacityToBeAdded"] = o.CapacityToBeAdded
	toSerialize["resourceId"] = o.ResourceId
	return toSerialize, nil
}

func (o *AddCapacityToResourceInstanceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capacityToBeAdded",
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

	varAddCapacityToResourceInstanceRequestBody := _AddCapacityToResourceInstanceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddCapacityToResourceInstanceRequestBody)

	if err != nil {
		return err
	}

	*o = AddCapacityToResourceInstanceRequestBody(varAddCapacityToResourceInstanceRequestBody)

	return err
}

type NullableAddCapacityToResourceInstanceRequestBody struct {
	value *AddCapacityToResourceInstanceRequestBody
	isSet bool
}

func (v NullableAddCapacityToResourceInstanceRequestBody) Get() *AddCapacityToResourceInstanceRequestBody {
	return v.value
}

func (v *NullableAddCapacityToResourceInstanceRequestBody) Set(val *AddCapacityToResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCapacityToResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCapacityToResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCapacityToResourceInstanceRequestBody(val *AddCapacityToResourceInstanceRequestBody) *NullableAddCapacityToResourceInstanceRequestBody {
	return &NullableAddCapacityToResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableAddCapacityToResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCapacityToResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


