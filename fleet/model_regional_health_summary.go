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

// checks if the RegionalHealthSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionalHealthSummary{}

// RegionalHealthSummary struct for RegionalHealthSummary
type RegionalHealthSummary struct {
	// The number of instances currently deploying
	DeployingInstances int64 `json:"deployingInstances"`
	// The summary of health by deployment cell
	DeploymentCellHealthSummary map[string]DeploymentCellHealthSummary `json:"deploymentCellHealthSummary"`
	// The number of healthy instances in the region
	HealthyInstances int64 `json:"healthyInstances"`
	// The status message
	Message string `json:"message"`
	// The region
	Region string `json:"region"`
	// The status of the region
	Status string `json:"status"`
	// The total number of instances in the region
	TotalInstances int64 `json:"totalInstances"`
	// The number of unhealthy instances in the region
	UnhealthyInstances int64 `json:"unhealthyInstances"`
}

type _RegionalHealthSummary RegionalHealthSummary

// NewRegionalHealthSummary instantiates a new RegionalHealthSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalHealthSummary(deployingInstances int64, deploymentCellHealthSummary map[string]DeploymentCellHealthSummary, healthyInstances int64, message string, region string, status string, totalInstances int64, unhealthyInstances int64) *RegionalHealthSummary {
	this := RegionalHealthSummary{}
	this.DeployingInstances = deployingInstances
	this.DeploymentCellHealthSummary = deploymentCellHealthSummary
	this.HealthyInstances = healthyInstances
	this.Message = message
	this.Region = region
	this.Status = status
	this.TotalInstances = totalInstances
	this.UnhealthyInstances = unhealthyInstances
	return &this
}

// NewRegionalHealthSummaryWithDefaults instantiates a new RegionalHealthSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalHealthSummaryWithDefaults() *RegionalHealthSummary {
	this := RegionalHealthSummary{}
	return &this
}

// GetDeployingInstances returns the DeployingInstances field value
func (o *RegionalHealthSummary) GetDeployingInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeployingInstances
}

// GetDeployingInstancesOk returns a tuple with the DeployingInstances field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetDeployingInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployingInstances, true
}

// SetDeployingInstances sets field value
func (o *RegionalHealthSummary) SetDeployingInstances(v int64) {
	o.DeployingInstances = v
}

// GetDeploymentCellHealthSummary returns the DeploymentCellHealthSummary field value
func (o *RegionalHealthSummary) GetDeploymentCellHealthSummary() map[string]DeploymentCellHealthSummary {
	if o == nil {
		var ret map[string]DeploymentCellHealthSummary
		return ret
	}

	return o.DeploymentCellHealthSummary
}

// GetDeploymentCellHealthSummaryOk returns a tuple with the DeploymentCellHealthSummary field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetDeploymentCellHealthSummaryOk() (*map[string]DeploymentCellHealthSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentCellHealthSummary, true
}

// SetDeploymentCellHealthSummary sets field value
func (o *RegionalHealthSummary) SetDeploymentCellHealthSummary(v map[string]DeploymentCellHealthSummary) {
	o.DeploymentCellHealthSummary = v
}

// GetHealthyInstances returns the HealthyInstances field value
func (o *RegionalHealthSummary) GetHealthyInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HealthyInstances
}

// GetHealthyInstancesOk returns a tuple with the HealthyInstances field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetHealthyInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthyInstances, true
}

// SetHealthyInstances sets field value
func (o *RegionalHealthSummary) SetHealthyInstances(v int64) {
	o.HealthyInstances = v
}

// GetMessage returns the Message field value
func (o *RegionalHealthSummary) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RegionalHealthSummary) SetMessage(v string) {
	o.Message = v
}

// GetRegion returns the Region field value
func (o *RegionalHealthSummary) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *RegionalHealthSummary) SetRegion(v string) {
	o.Region = v
}

// GetStatus returns the Status field value
func (o *RegionalHealthSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RegionalHealthSummary) SetStatus(v string) {
	o.Status = v
}

// GetTotalInstances returns the TotalInstances field value
func (o *RegionalHealthSummary) GetTotalInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalInstances
}

// GetTotalInstancesOk returns a tuple with the TotalInstances field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetTotalInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInstances, true
}

// SetTotalInstances sets field value
func (o *RegionalHealthSummary) SetTotalInstances(v int64) {
	o.TotalInstances = v
}

// GetUnhealthyInstances returns the UnhealthyInstances field value
func (o *RegionalHealthSummary) GetUnhealthyInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnhealthyInstances
}

// GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field value
// and a boolean to check if the value has been set.
func (o *RegionalHealthSummary) GetUnhealthyInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnhealthyInstances, true
}

// SetUnhealthyInstances sets field value
func (o *RegionalHealthSummary) SetUnhealthyInstances(v int64) {
	o.UnhealthyInstances = v
}

func (o RegionalHealthSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionalHealthSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deployingInstances"] = o.DeployingInstances
	toSerialize["deploymentCellHealthSummary"] = o.DeploymentCellHealthSummary
	toSerialize["healthyInstances"] = o.HealthyInstances
	toSerialize["message"] = o.Message
	toSerialize["region"] = o.Region
	toSerialize["status"] = o.Status
	toSerialize["totalInstances"] = o.TotalInstances
	toSerialize["unhealthyInstances"] = o.UnhealthyInstances
	return toSerialize, nil
}

func (o *RegionalHealthSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deployingInstances",
		"deploymentCellHealthSummary",
		"healthyInstances",
		"message",
		"region",
		"status",
		"totalInstances",
		"unhealthyInstances",
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

	varRegionalHealthSummary := _RegionalHealthSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegionalHealthSummary)

	if err != nil {
		return err
	}

	*o = RegionalHealthSummary(varRegionalHealthSummary)

	return err
}

type NullableRegionalHealthSummary struct {
	value *RegionalHealthSummary
	isSet bool
}

func (v NullableRegionalHealthSummary) Get() *RegionalHealthSummary {
	return v.value
}

func (v *NullableRegionalHealthSummary) Set(val *RegionalHealthSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalHealthSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalHealthSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalHealthSummary(val *RegionalHealthSummary) *NullableRegionalHealthSummary {
	return &NullableRegionalHealthSummary{value: val, isSet: true}
}

func (v NullableRegionalHealthSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalHealthSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


