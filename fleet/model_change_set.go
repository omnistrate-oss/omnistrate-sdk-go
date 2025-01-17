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

// checks if the ChangeSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeSet{}

// ChangeSet struct for ChangeSet
type ChangeSet struct {
	// Summary of all changes within a resource for each category (ie. image, infra, product tier feature, etc.)
	CategorizedResourceChanges map[string]interface{} `json:"categorizedResourceChanges"`
	ImageConfigChanges *ImageConfigChangeSummary `json:"imageConfigChanges,omitempty"`
	InfraConfigChanges *InfraConfigChangeSummary `json:"infraConfigChanges,omitempty"`
	// Summary status of the changes
	OverallResourceStatus string `json:"overallResourceStatus"`
	// State of the configuration change
	ProductTierFeatureChanges *string `json:"productTierFeatureChanges,omitempty"`
	// State of the configuration change
	ResourceChanges *string `json:"resourceChanges,omitempty"`
	// The name of the resource
	ResourceName *string `json:"resourceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeSet ChangeSet

// NewChangeSet instantiates a new ChangeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeSet(categorizedResourceChanges map[string]interface{}, overallResourceStatus string) *ChangeSet {
	this := ChangeSet{}
	this.CategorizedResourceChanges = categorizedResourceChanges
	this.OverallResourceStatus = overallResourceStatus
	return &this
}

// NewChangeSetWithDefaults instantiates a new ChangeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeSetWithDefaults() *ChangeSet {
	this := ChangeSet{}
	return &this
}

// GetCategorizedResourceChanges returns the CategorizedResourceChanges field value
func (o *ChangeSet) GetCategorizedResourceChanges() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CategorizedResourceChanges
}

// GetCategorizedResourceChangesOk returns a tuple with the CategorizedResourceChanges field value
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetCategorizedResourceChangesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CategorizedResourceChanges, true
}

// SetCategorizedResourceChanges sets field value
func (o *ChangeSet) SetCategorizedResourceChanges(v map[string]interface{}) {
	o.CategorizedResourceChanges = v
}

// GetImageConfigChanges returns the ImageConfigChanges field value if set, zero value otherwise.
func (o *ChangeSet) GetImageConfigChanges() ImageConfigChangeSummary {
	if o == nil || IsNil(o.ImageConfigChanges) {
		var ret ImageConfigChangeSummary
		return ret
	}
	return *o.ImageConfigChanges
}

// GetImageConfigChangesOk returns a tuple with the ImageConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetImageConfigChangesOk() (*ImageConfigChangeSummary, bool) {
	if o == nil || IsNil(o.ImageConfigChanges) {
		return nil, false
	}
	return o.ImageConfigChanges, true
}

// HasImageConfigChanges returns a boolean if a field has been set.
func (o *ChangeSet) HasImageConfigChanges() bool {
	if o != nil && !IsNil(o.ImageConfigChanges) {
		return true
	}

	return false
}

// SetImageConfigChanges gets a reference to the given ImageConfigChangeSummary and assigns it to the ImageConfigChanges field.
func (o *ChangeSet) SetImageConfigChanges(v ImageConfigChangeSummary) {
	o.ImageConfigChanges = &v
}

// GetInfraConfigChanges returns the InfraConfigChanges field value if set, zero value otherwise.
func (o *ChangeSet) GetInfraConfigChanges() InfraConfigChangeSummary {
	if o == nil || IsNil(o.InfraConfigChanges) {
		var ret InfraConfigChangeSummary
		return ret
	}
	return *o.InfraConfigChanges
}

// GetInfraConfigChangesOk returns a tuple with the InfraConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetInfraConfigChangesOk() (*InfraConfigChangeSummary, bool) {
	if o == nil || IsNil(o.InfraConfigChanges) {
		return nil, false
	}
	return o.InfraConfigChanges, true
}

// HasInfraConfigChanges returns a boolean if a field has been set.
func (o *ChangeSet) HasInfraConfigChanges() bool {
	if o != nil && !IsNil(o.InfraConfigChanges) {
		return true
	}

	return false
}

// SetInfraConfigChanges gets a reference to the given InfraConfigChangeSummary and assigns it to the InfraConfigChanges field.
func (o *ChangeSet) SetInfraConfigChanges(v InfraConfigChangeSummary) {
	o.InfraConfigChanges = &v
}

// GetOverallResourceStatus returns the OverallResourceStatus field value
func (o *ChangeSet) GetOverallResourceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OverallResourceStatus
}

// GetOverallResourceStatusOk returns a tuple with the OverallResourceStatus field value
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetOverallResourceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallResourceStatus, true
}

// SetOverallResourceStatus sets field value
func (o *ChangeSet) SetOverallResourceStatus(v string) {
	o.OverallResourceStatus = v
}

// GetProductTierFeatureChanges returns the ProductTierFeatureChanges field value if set, zero value otherwise.
func (o *ChangeSet) GetProductTierFeatureChanges() string {
	if o == nil || IsNil(o.ProductTierFeatureChanges) {
		var ret string
		return ret
	}
	return *o.ProductTierFeatureChanges
}

// GetProductTierFeatureChangesOk returns a tuple with the ProductTierFeatureChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetProductTierFeatureChangesOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierFeatureChanges) {
		return nil, false
	}
	return o.ProductTierFeatureChanges, true
}

// HasProductTierFeatureChanges returns a boolean if a field has been set.
func (o *ChangeSet) HasProductTierFeatureChanges() bool {
	if o != nil && !IsNil(o.ProductTierFeatureChanges) {
		return true
	}

	return false
}

// SetProductTierFeatureChanges gets a reference to the given string and assigns it to the ProductTierFeatureChanges field.
func (o *ChangeSet) SetProductTierFeatureChanges(v string) {
	o.ProductTierFeatureChanges = &v
}

// GetResourceChanges returns the ResourceChanges field value if set, zero value otherwise.
func (o *ChangeSet) GetResourceChanges() string {
	if o == nil || IsNil(o.ResourceChanges) {
		var ret string
		return ret
	}
	return *o.ResourceChanges
}

// GetResourceChangesOk returns a tuple with the ResourceChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetResourceChangesOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceChanges) {
		return nil, false
	}
	return o.ResourceChanges, true
}

// HasResourceChanges returns a boolean if a field has been set.
func (o *ChangeSet) HasResourceChanges() bool {
	if o != nil && !IsNil(o.ResourceChanges) {
		return true
	}

	return false
}

// SetResourceChanges gets a reference to the given string and assigns it to the ResourceChanges field.
func (o *ChangeSet) SetResourceChanges(v string) {
	o.ResourceChanges = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ChangeSet) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ChangeSet) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ChangeSet) SetResourceName(v string) {
	o.ResourceName = &v
}

func (o ChangeSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categorizedResourceChanges"] = o.CategorizedResourceChanges
	if !IsNil(o.ImageConfigChanges) {
		toSerialize["imageConfigChanges"] = o.ImageConfigChanges
	}
	if !IsNil(o.InfraConfigChanges) {
		toSerialize["infraConfigChanges"] = o.InfraConfigChanges
	}
	toSerialize["overallResourceStatus"] = o.OverallResourceStatus
	if !IsNil(o.ProductTierFeatureChanges) {
		toSerialize["productTierFeatureChanges"] = o.ProductTierFeatureChanges
	}
	if !IsNil(o.ResourceChanges) {
		toSerialize["resourceChanges"] = o.ResourceChanges
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categorizedResourceChanges",
		"overallResourceStatus",
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

	varChangeSet := _ChangeSet{}

	err = json.Unmarshal(data, &varChangeSet)

	if err != nil {
		return err
	}

	*o = ChangeSet(varChangeSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "categorizedResourceChanges")
		delete(additionalProperties, "imageConfigChanges")
		delete(additionalProperties, "infraConfigChanges")
		delete(additionalProperties, "overallResourceStatus")
		delete(additionalProperties, "productTierFeatureChanges")
		delete(additionalProperties, "resourceChanges")
		delete(additionalProperties, "resourceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeSet struct {
	value *ChangeSet
	isSet bool
}

func (v NullableChangeSet) Get() *ChangeSet {
	return v.value
}

func (v *NullableChangeSet) Set(val *ChangeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeSet(val *ChangeSet) *NullableChangeSet {
	return &NullableChangeSet{value: val, isSet: true}
}

func (v NullableChangeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


