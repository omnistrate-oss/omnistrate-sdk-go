/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the ResourceSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSummary{}

// ResourceSummary struct for ResourceSummary
type ResourceSummary struct {
	// A brief description of the resource
	Description string `json:"description"`
	// ID of a resource
	Id string `json:"id"`
	// ID of an Image Config
	ImageConfigId *string `json:"imageConfigId,omitempty"`
	// ID of an Infra Config
	InfraConfigId *string `json:"infraConfigId,omitempty"`
	// Whether the resource is external
	IsExternal bool `json:"isExternal"`
	// The managed resource type of instance
	ManagedResourceType *string `json:"managedResourceType,omitempty"`
	// The name of the resource
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSummary ResourceSummary

// NewResourceSummary instantiates a new ResourceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSummary(description string, id string, isExternal bool, name string) *ResourceSummary {
	this := ResourceSummary{}
	this.Description = description
	this.Id = id
	this.IsExternal = isExternal
	this.Name = name
	return &this
}

// NewResourceSummaryWithDefaults instantiates a new ResourceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSummaryWithDefaults() *ResourceSummary {
	this := ResourceSummary{}
	return &this
}

// GetDescription returns the Description field value
func (o *ResourceSummary) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ResourceSummary) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *ResourceSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceSummary) SetId(v string) {
	o.Id = v
}

// GetImageConfigId returns the ImageConfigId field value if set, zero value otherwise.
func (o *ResourceSummary) GetImageConfigId() string {
	if o == nil || IsNil(o.ImageConfigId) {
		var ret string
		return ret
	}
	return *o.ImageConfigId
}

// GetImageConfigIdOk returns a tuple with the ImageConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetImageConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageConfigId) {
		return nil, false
	}
	return o.ImageConfigId, true
}

// SetImageConfigId gets a reference to the given string and assigns it to the ImageConfigId field.
func (o *ResourceSummary) SetImageConfigId(v string) {
	o.ImageConfigId = &v
}

// GetInfraConfigId returns the InfraConfigId field value if set, zero value otherwise.
func (o *ResourceSummary) GetInfraConfigId() string {
	if o == nil || IsNil(o.InfraConfigId) {
		var ret string
		return ret
	}
	return *o.InfraConfigId
}

// GetInfraConfigIdOk returns a tuple with the InfraConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetInfraConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfraConfigId) {
		return nil, false
	}
	return o.InfraConfigId, true
}

// SetInfraConfigId gets a reference to the given string and assigns it to the InfraConfigId field.
func (o *ResourceSummary) SetInfraConfigId(v string) {
	o.InfraConfigId = &v
}

// GetIsExternal returns the IsExternal field value
func (o *ResourceSummary) GetIsExternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetIsExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExternal, true
}

// SetIsExternal sets field value
func (o *ResourceSummary) SetIsExternal(v bool) {
	o.IsExternal = v
}

// GetManagedResourceType returns the ManagedResourceType field value if set, zero value otherwise.
func (o *ResourceSummary) GetManagedResourceType() string {
	if o == nil || IsNil(o.ManagedResourceType) {
		var ret string
		return ret
	}
	return *o.ManagedResourceType
}

// GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetManagedResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedResourceType) {
		return nil, false
	}
	return o.ManagedResourceType, true
}

// SetManagedResourceType gets a reference to the given string and assigns it to the ManagedResourceType field.
func (o *ResourceSummary) SetManagedResourceType(v string) {
	o.ManagedResourceType = &v
}

// GetName returns the Name field value
func (o *ResourceSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceSummary) SetName(v string) {
	o.Name = v
}

func (o ResourceSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	if !IsNil(o.ImageConfigId) {
		toSerialize["imageConfigId"] = o.ImageConfigId
	}
	if !IsNil(o.InfraConfigId) {
		toSerialize["infraConfigId"] = o.InfraConfigId
	}
	toSerialize["isExternal"] = o.IsExternal
	if !IsNil(o.ManagedResourceType) {
		toSerialize["managedResourceType"] = o.ManagedResourceType
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"isExternal",
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

	varResourceSummary := _ResourceSummary{}

	err = json.Unmarshal(data, &varResourceSummary)

	if err != nil {
		return err
	}

	*o = ResourceSummary(varResourceSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "imageConfigId")
		delete(additionalProperties, "infraConfigId")
		delete(additionalProperties, "isExternal")
		delete(additionalProperties, "managedResourceType")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSummary struct {
	value *ResourceSummary
	isSet bool
}

func (v NullableResourceSummary) Get() *ResourceSummary {
	return v.value
}

func (v *NullableResourceSummary) Set(val *ResourceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSummary(val *ResourceSummary) *NullableResourceSummary {
	return &NullableResourceSummary{value: val, isSet: true}
}

func (v NullableResourceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


