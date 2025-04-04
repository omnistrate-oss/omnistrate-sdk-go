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

// checks if the PerDeploymentCellCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerDeploymentCellCost{}

// PerDeploymentCellCost struct for PerDeploymentCellCost
type PerDeploymentCellCost struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	Cost []CostDataPerDate `json:"cost"`
	InstancesCost []PerInstanceCost `json:"instancesCost,omitempty"`
	// The name of the region
	RegionName string `json:"regionName"`
	// The total cost of the fleet in the deployment cell
	TotalCost float64 `json:"totalCost"`
	AdditionalProperties map[string]interface{}
}

type _PerDeploymentCellCost PerDeploymentCellCost

// NewPerDeploymentCellCost instantiates a new PerDeploymentCellCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerDeploymentCellCost(cloudProviderName string, cost []CostDataPerDate, regionName string, totalCost float64) *PerDeploymentCellCost {
	this := PerDeploymentCellCost{}
	this.CloudProviderName = cloudProviderName
	this.Cost = cost
	this.RegionName = regionName
	this.TotalCost = totalCost
	return &this
}

// NewPerDeploymentCellCostWithDefaults instantiates a new PerDeploymentCellCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerDeploymentCellCostWithDefaults() *PerDeploymentCellCost {
	this := PerDeploymentCellCost{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *PerDeploymentCellCost) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *PerDeploymentCellCost) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *PerDeploymentCellCost) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCost returns the Cost field value
func (o *PerDeploymentCellCost) GetCost() []CostDataPerDate {
	if o == nil {
		var ret []CostDataPerDate
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *PerDeploymentCellCost) GetCostOk() ([]CostDataPerDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost, true
}

// SetCost sets field value
func (o *PerDeploymentCellCost) SetCost(v []CostDataPerDate) {
	o.Cost = v
}

// GetInstancesCost returns the InstancesCost field value if set, zero value otherwise.
func (o *PerDeploymentCellCost) GetInstancesCost() []PerInstanceCost {
	if o == nil || IsNil(o.InstancesCost) {
		var ret []PerInstanceCost
		return ret
	}
	return o.InstancesCost
}

// GetInstancesCostOk returns a tuple with the InstancesCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerDeploymentCellCost) GetInstancesCostOk() ([]PerInstanceCost, bool) {
	if o == nil || IsNil(o.InstancesCost) {
		return nil, false
	}
	return o.InstancesCost, true
}

// HasInstancesCost returns a boolean if a field has been set.
func (o *PerDeploymentCellCost) HasInstancesCost() bool {
	if o != nil && !IsNil(o.InstancesCost) {
		return true
	}

	return false
}

// SetInstancesCost gets a reference to the given []PerInstanceCost and assigns it to the InstancesCost field.
func (o *PerDeploymentCellCost) SetInstancesCost(v []PerInstanceCost) {
	o.InstancesCost = v
}

// GetRegionName returns the RegionName field value
func (o *PerDeploymentCellCost) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *PerDeploymentCellCost) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *PerDeploymentCellCost) SetRegionName(v string) {
	o.RegionName = v
}

// GetTotalCost returns the TotalCost field value
func (o *PerDeploymentCellCost) GetTotalCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value
// and a boolean to check if the value has been set.
func (o *PerDeploymentCellCost) GetTotalCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCost, true
}

// SetTotalCost sets field value
func (o *PerDeploymentCellCost) SetTotalCost(v float64) {
	o.TotalCost = v
}

func (o PerDeploymentCellCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerDeploymentCellCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["cost"] = o.Cost
	if !IsNil(o.InstancesCost) {
		toSerialize["instancesCost"] = o.InstancesCost
	}
	toSerialize["regionName"] = o.RegionName
	toSerialize["totalCost"] = o.TotalCost

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PerDeploymentCellCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"cost",
		"regionName",
		"totalCost",
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

	varPerDeploymentCellCost := _PerDeploymentCellCost{}

	err = json.Unmarshal(data, &varPerDeploymentCellCost)

	if err != nil {
		return err
	}

	*o = PerDeploymentCellCost(varPerDeploymentCellCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "instancesCost")
		delete(additionalProperties, "regionName")
		delete(additionalProperties, "totalCost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePerDeploymentCellCost struct {
	value *PerDeploymentCellCost
	isSet bool
}

func (v NullablePerDeploymentCellCost) Get() *PerDeploymentCellCost {
	return v.value
}

func (v *NullablePerDeploymentCellCost) Set(val *PerDeploymentCellCost) {
	v.value = val
	v.isSet = true
}

func (v NullablePerDeploymentCellCost) IsSet() bool {
	return v.isSet
}

func (v *NullablePerDeploymentCellCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerDeploymentCellCost(val *PerDeploymentCellCost) *NullablePerDeploymentCellCost {
	return &NullablePerDeploymentCellCost{value: val, isSet: true}
}

func (v NullablePerDeploymentCellCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerDeploymentCellCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


