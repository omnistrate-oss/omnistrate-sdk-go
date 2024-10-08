/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServiceHealthSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceHealthSummary{}

// ServiceHealthSummary struct for ServiceHealthSummary
type ServiceHealthSummary struct {
	// The summary of health by cloud provider
	CloudProviderHealthSummary map[string]CloudProviderHealthSummary `json:"cloudProviderHealthSummary"`
	// The number of instances currently deploying
	DeployingInstances int64 `json:"deployingInstances"`
	// The number of healthy instances in the region
	HealthyInstances int64 `json:"healthyInstances"`
	// The status message
	Message string `json:"message"`
	// The ID of the service environment
	ServiceEnvironmentID string `json:"serviceEnvironmentID"`
	// The ID of the service
	ServiceID string `json:"serviceID"`
	// The status of the service
	Status string `json:"status"`
	// The total number of instances in the region
	TotalInstances int64 `json:"totalInstances"`
	// The number of unhealthy instances in the region
	UnhealthyInstances int64 `json:"unhealthyInstances"`
}

type _ServiceHealthSummary ServiceHealthSummary

// NewServiceHealthSummary instantiates a new ServiceHealthSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHealthSummary(cloudProviderHealthSummary map[string]CloudProviderHealthSummary, deployingInstances int64, healthyInstances int64, message string, serviceEnvironmentID string, serviceID string, status string, totalInstances int64, unhealthyInstances int64) *ServiceHealthSummary {
	this := ServiceHealthSummary{}
	this.CloudProviderHealthSummary = cloudProviderHealthSummary
	this.DeployingInstances = deployingInstances
	this.HealthyInstances = healthyInstances
	this.Message = message
	this.ServiceEnvironmentID = serviceEnvironmentID
	this.ServiceID = serviceID
	this.Status = status
	this.TotalInstances = totalInstances
	this.UnhealthyInstances = unhealthyInstances
	return &this
}

// NewServiceHealthSummaryWithDefaults instantiates a new ServiceHealthSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHealthSummaryWithDefaults() *ServiceHealthSummary {
	this := ServiceHealthSummary{}
	return &this
}

// GetCloudProviderHealthSummary returns the CloudProviderHealthSummary field value
func (o *ServiceHealthSummary) GetCloudProviderHealthSummary() map[string]CloudProviderHealthSummary {
	if o == nil {
		var ret map[string]CloudProviderHealthSummary
		return ret
	}

	return o.CloudProviderHealthSummary
}

// GetCloudProviderHealthSummaryOk returns a tuple with the CloudProviderHealthSummary field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetCloudProviderHealthSummaryOk() (*map[string]CloudProviderHealthSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderHealthSummary, true
}

// SetCloudProviderHealthSummary sets field value
func (o *ServiceHealthSummary) SetCloudProviderHealthSummary(v map[string]CloudProviderHealthSummary) {
	o.CloudProviderHealthSummary = v
}

// GetDeployingInstances returns the DeployingInstances field value
func (o *ServiceHealthSummary) GetDeployingInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeployingInstances
}

// GetDeployingInstancesOk returns a tuple with the DeployingInstances field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetDeployingInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployingInstances, true
}

// SetDeployingInstances sets field value
func (o *ServiceHealthSummary) SetDeployingInstances(v int64) {
	o.DeployingInstances = v
}

// GetHealthyInstances returns the HealthyInstances field value
func (o *ServiceHealthSummary) GetHealthyInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HealthyInstances
}

// GetHealthyInstancesOk returns a tuple with the HealthyInstances field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetHealthyInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthyInstances, true
}

// SetHealthyInstances sets field value
func (o *ServiceHealthSummary) SetHealthyInstances(v int64) {
	o.HealthyInstances = v
}

// GetMessage returns the Message field value
func (o *ServiceHealthSummary) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServiceHealthSummary) SetMessage(v string) {
	o.Message = v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value
func (o *ServiceHealthSummary) GetServiceEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID sets field value
func (o *ServiceHealthSummary) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = v
}

// GetServiceID returns the ServiceID field value
func (o *ServiceHealthSummary) GetServiceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetServiceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceID, true
}

// SetServiceID sets field value
func (o *ServiceHealthSummary) SetServiceID(v string) {
	o.ServiceID = v
}

// GetStatus returns the Status field value
func (o *ServiceHealthSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ServiceHealthSummary) SetStatus(v string) {
	o.Status = v
}

// GetTotalInstances returns the TotalInstances field value
func (o *ServiceHealthSummary) GetTotalInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalInstances
}

// GetTotalInstancesOk returns a tuple with the TotalInstances field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetTotalInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInstances, true
}

// SetTotalInstances sets field value
func (o *ServiceHealthSummary) SetTotalInstances(v int64) {
	o.TotalInstances = v
}

// GetUnhealthyInstances returns the UnhealthyInstances field value
func (o *ServiceHealthSummary) GetUnhealthyInstances() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnhealthyInstances
}

// GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field value
// and a boolean to check if the value has been set.
func (o *ServiceHealthSummary) GetUnhealthyInstancesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnhealthyInstances, true
}

// SetUnhealthyInstances sets field value
func (o *ServiceHealthSummary) SetUnhealthyInstances(v int64) {
	o.UnhealthyInstances = v
}

func (o ServiceHealthSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHealthSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderHealthSummary"] = o.CloudProviderHealthSummary
	toSerialize["deployingInstances"] = o.DeployingInstances
	toSerialize["healthyInstances"] = o.HealthyInstances
	toSerialize["message"] = o.Message
	toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	toSerialize["serviceID"] = o.ServiceID
	toSerialize["status"] = o.Status
	toSerialize["totalInstances"] = o.TotalInstances
	toSerialize["unhealthyInstances"] = o.UnhealthyInstances
	return toSerialize, nil
}

func (o *ServiceHealthSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderHealthSummary",
		"deployingInstances",
		"healthyInstances",
		"message",
		"serviceEnvironmentID",
		"serviceID",
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

	varServiceHealthSummary := _ServiceHealthSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceHealthSummary)

	if err != nil {
		return err
	}

	*o = ServiceHealthSummary(varServiceHealthSummary)

	return err
}

type NullableServiceHealthSummary struct {
	value *ServiceHealthSummary
	isSet bool
}

func (v NullableServiceHealthSummary) Get() *ServiceHealthSummary {
	return v.value
}

func (v *NullableServiceHealthSummary) Set(val *ServiceHealthSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHealthSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHealthSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHealthSummary(val *ServiceHealthSummary) *NullableServiceHealthSummary {
	return &NullableServiceHealthSummary{value: val, isSet: true}
}

func (v NullableServiceHealthSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHealthSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


