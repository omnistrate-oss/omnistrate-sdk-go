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

// checks if the ResourceEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceEntity{}

// ResourceEntity struct for ResourceEntity
type ResourceEntity struct {
	// Whether the resource has autoscaling enabled
	IsAutoscalingEnabled *bool `json:"isAutoscalingEnabled,omitempty"`
	// Whether the resource has backup enabled
	IsBackupEnabled bool `json:"isBackupEnabled"`
	// Whether the service offering is deprecated
	IsDeprecated bool `json:"isDeprecated"`
	// The resource name
	Name string `json:"name"`
	// ID of a resource
	ResourceId string `json:"resourceId"`
	// The resource type
	ResourceType *string `json:"resourceType,omitempty"`
	// The resource URL key
	UrlKey string `json:"urlKey"`
	AdditionalProperties map[string]interface{}
}

type _ResourceEntity ResourceEntity

// NewResourceEntity instantiates a new ResourceEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceEntity(isBackupEnabled bool, isDeprecated bool, name string, resourceId string, urlKey string) *ResourceEntity {
	this := ResourceEntity{}
	this.IsBackupEnabled = isBackupEnabled
	this.IsDeprecated = isDeprecated
	this.Name = name
	this.ResourceId = resourceId
	this.UrlKey = urlKey
	return &this
}

// NewResourceEntityWithDefaults instantiates a new ResourceEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceEntityWithDefaults() *ResourceEntity {
	this := ResourceEntity{}
	var isDeprecated bool = false
	this.IsDeprecated = isDeprecated
	return &this
}

// GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field value if set, zero value otherwise.
func (o *ResourceEntity) GetIsAutoscalingEnabled() bool {
	if o == nil || IsNil(o.IsAutoscalingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAutoscalingEnabled
}

// GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetIsAutoscalingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoscalingEnabled) {
		return nil, false
	}
	return o.IsAutoscalingEnabled, true
}

// HasIsAutoscalingEnabled returns a boolean if a field has been set.
func (o *ResourceEntity) HasIsAutoscalingEnabled() bool {
	if o != nil && !IsNil(o.IsAutoscalingEnabled) {
		return true
	}

	return false
}

// SetIsAutoscalingEnabled gets a reference to the given bool and assigns it to the IsAutoscalingEnabled field.
func (o *ResourceEntity) SetIsAutoscalingEnabled(v bool) {
	o.IsAutoscalingEnabled = &v
}

// GetIsBackupEnabled returns the IsBackupEnabled field value
func (o *ResourceEntity) GetIsBackupEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBackupEnabled
}

// GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field value
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetIsBackupEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBackupEnabled, true
}

// SetIsBackupEnabled sets field value
func (o *ResourceEntity) SetIsBackupEnabled(v bool) {
	o.IsBackupEnabled = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *ResourceEntity) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *ResourceEntity) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetName returns the Name field value
func (o *ResourceEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceEntity) SetName(v string) {
	o.Name = v
}

// GetResourceId returns the ResourceId field value
func (o *ResourceEntity) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ResourceEntity) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceEntity) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceEntity) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourceEntity) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetUrlKey returns the UrlKey field value
func (o *ResourceEntity) GetUrlKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value
// and a boolean to check if the value has been set.
func (o *ResourceEntity) GetUrlKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlKey, true
}

// SetUrlKey sets field value
func (o *ResourceEntity) SetUrlKey(v string) {
	o.UrlKey = v
}

func (o ResourceEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAutoscalingEnabled) {
		toSerialize["isAutoscalingEnabled"] = o.IsAutoscalingEnabled
	}
	toSerialize["isBackupEnabled"] = o.IsBackupEnabled
	toSerialize["isDeprecated"] = o.IsDeprecated
	toSerialize["name"] = o.Name
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	toSerialize["urlKey"] = o.UrlKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isBackupEnabled",
		"isDeprecated",
		"name",
		"resourceId",
		"urlKey",
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

	varResourceEntity := _ResourceEntity{}

	err = json.Unmarshal(data, &varResourceEntity)

	if err != nil {
		return err
	}

	*o = ResourceEntity(varResourceEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isAutoscalingEnabled")
		delete(additionalProperties, "isBackupEnabled")
		delete(additionalProperties, "isDeprecated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "urlKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceEntity struct {
	value *ResourceEntity
	isSet bool
}

func (v NullableResourceEntity) Get() *ResourceEntity {
	return v.value
}

func (v *NullableResourceEntity) Set(val *ResourceEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceEntity(val *ResourceEntity) *NullableResourceEntity {
	return &NullableResourceEntity{value: val, isSet: true}
}

func (v NullableResourceEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


