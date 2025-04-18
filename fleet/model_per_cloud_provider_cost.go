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

// checks if the PerCloudProviderCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerCloudProviderCost{}

// PerCloudProviderCost struct for PerCloudProviderCost
type PerCloudProviderCost struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	Cost []CostDataPerDate `json:"cost"`
	// The total cost of the fleet on the cloud provider
	TotalCost float64 `json:"totalCost"`
	AdditionalProperties map[string]interface{}
}

type _PerCloudProviderCost PerCloudProviderCost

// NewPerCloudProviderCost instantiates a new PerCloudProviderCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerCloudProviderCost(cloudProviderName string, cost []CostDataPerDate, totalCost float64) *PerCloudProviderCost {
	this := PerCloudProviderCost{}
	this.CloudProviderName = cloudProviderName
	this.Cost = cost
	this.TotalCost = totalCost
	return &this
}

// NewPerCloudProviderCostWithDefaults instantiates a new PerCloudProviderCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerCloudProviderCostWithDefaults() *PerCloudProviderCost {
	this := PerCloudProviderCost{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *PerCloudProviderCost) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *PerCloudProviderCost) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *PerCloudProviderCost) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCost returns the Cost field value
func (o *PerCloudProviderCost) GetCost() []CostDataPerDate {
	if o == nil {
		var ret []CostDataPerDate
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *PerCloudProviderCost) GetCostOk() ([]CostDataPerDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost, true
}

// SetCost sets field value
func (o *PerCloudProviderCost) SetCost(v []CostDataPerDate) {
	o.Cost = v
}

// GetTotalCost returns the TotalCost field value
func (o *PerCloudProviderCost) GetTotalCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value
// and a boolean to check if the value has been set.
func (o *PerCloudProviderCost) GetTotalCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCost, true
}

// SetTotalCost sets field value
func (o *PerCloudProviderCost) SetTotalCost(v float64) {
	o.TotalCost = v
}

func (o PerCloudProviderCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerCloudProviderCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["cost"] = o.Cost
	toSerialize["totalCost"] = o.TotalCost

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PerCloudProviderCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"cost",
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

	varPerCloudProviderCost := _PerCloudProviderCost{}

	err = json.Unmarshal(data, &varPerCloudProviderCost)

	if err != nil {
		return err
	}

	*o = PerCloudProviderCost(varPerCloudProviderCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "totalCost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePerCloudProviderCost struct {
	value *PerCloudProviderCost
	isSet bool
}

func (v NullablePerCloudProviderCost) Get() *PerCloudProviderCost {
	return v.value
}

func (v *NullablePerCloudProviderCost) Set(val *PerCloudProviderCost) {
	v.value = val
	v.isSet = true
}

func (v NullablePerCloudProviderCost) IsSet() bool {
	return v.isSet
}

func (v *NullablePerCloudProviderCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerCloudProviderCost(val *PerCloudProviderCost) *NullablePerCloudProviderCost {
	return &NullablePerCloudProviderCost{value: val, isSet: true}
}

func (v NullablePerCloudProviderCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerCloudProviderCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


