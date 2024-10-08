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

// checks if the DescribeInventorySummaryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeInventorySummaryResult{}

// DescribeInventorySummaryResult struct for DescribeInventorySummaryResult
type DescribeInventorySummaryResult struct {
	// The number of environments in the service.
	EnvironmentCount int64 `json:"EnvironmentCount"`
	// The number of images in the service.
	ImagesCount int64 `json:"ImagesCount"`
	// The number of infrastructure configurations in the service.
	InfraConfigCount int64 `json:"InfraConfigCount"`
	// The number of instances in the service.
	InstancesCount int64 `json:"InstancesCount"`
	// The number of active organizations using the service.
	OrganizationCount int64 `json:"OrganizationCount"`
	// The number of resources in the service.
	ResourceCount int64 `json:"ResourceCount"`
	// The number of active users using the service.
	UserCount int64 `json:"UserCount"`
	// The service environment ID this workflow belongs to.
	EnvironmentId string `json:"environmentId"`
	// The service ID this workflow belongs to.
	ServiceId string `json:"serviceId"`
}

type _DescribeInventorySummaryResult DescribeInventorySummaryResult

// NewDescribeInventorySummaryResult instantiates a new DescribeInventorySummaryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeInventorySummaryResult(environmentCount int64, imagesCount int64, infraConfigCount int64, instancesCount int64, organizationCount int64, resourceCount int64, userCount int64, environmentId string, serviceId string) *DescribeInventorySummaryResult {
	this := DescribeInventorySummaryResult{}
	this.EnvironmentCount = environmentCount
	this.ImagesCount = imagesCount
	this.InfraConfigCount = infraConfigCount
	this.InstancesCount = instancesCount
	this.OrganizationCount = organizationCount
	this.ResourceCount = resourceCount
	this.UserCount = userCount
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	return &this
}

// NewDescribeInventorySummaryResultWithDefaults instantiates a new DescribeInventorySummaryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeInventorySummaryResultWithDefaults() *DescribeInventorySummaryResult {
	this := DescribeInventorySummaryResult{}
	return &this
}

// GetEnvironmentCount returns the EnvironmentCount field value
func (o *DescribeInventorySummaryResult) GetEnvironmentCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EnvironmentCount
}

// GetEnvironmentCountOk returns a tuple with the EnvironmentCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetEnvironmentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentCount, true
}

// SetEnvironmentCount sets field value
func (o *DescribeInventorySummaryResult) SetEnvironmentCount(v int64) {
	o.EnvironmentCount = v
}

// GetImagesCount returns the ImagesCount field value
func (o *DescribeInventorySummaryResult) GetImagesCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ImagesCount
}

// GetImagesCountOk returns a tuple with the ImagesCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetImagesCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImagesCount, true
}

// SetImagesCount sets field value
func (o *DescribeInventorySummaryResult) SetImagesCount(v int64) {
	o.ImagesCount = v
}

// GetInfraConfigCount returns the InfraConfigCount field value
func (o *DescribeInventorySummaryResult) GetInfraConfigCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InfraConfigCount
}

// GetInfraConfigCountOk returns a tuple with the InfraConfigCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetInfraConfigCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfraConfigCount, true
}

// SetInfraConfigCount sets field value
func (o *DescribeInventorySummaryResult) SetInfraConfigCount(v int64) {
	o.InfraConfigCount = v
}

// GetInstancesCount returns the InstancesCount field value
func (o *DescribeInventorySummaryResult) GetInstancesCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstancesCount
}

// GetInstancesCountOk returns a tuple with the InstancesCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetInstancesCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstancesCount, true
}

// SetInstancesCount sets field value
func (o *DescribeInventorySummaryResult) SetInstancesCount(v int64) {
	o.InstancesCount = v
}

// GetOrganizationCount returns the OrganizationCount field value
func (o *DescribeInventorySummaryResult) GetOrganizationCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationCount
}

// GetOrganizationCountOk returns a tuple with the OrganizationCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetOrganizationCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationCount, true
}

// SetOrganizationCount sets field value
func (o *DescribeInventorySummaryResult) SetOrganizationCount(v int64) {
	o.OrganizationCount = v
}

// GetResourceCount returns the ResourceCount field value
func (o *DescribeInventorySummaryResult) GetResourceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetResourceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCount, true
}

// SetResourceCount sets field value
func (o *DescribeInventorySummaryResult) SetResourceCount(v int64) {
	o.ResourceCount = v
}

// GetUserCount returns the UserCount field value
func (o *DescribeInventorySummaryResult) GetUserCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetUserCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCount, true
}

// SetUserCount sets field value
func (o *DescribeInventorySummaryResult) SetUserCount(v int64) {
	o.UserCount = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *DescribeInventorySummaryResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *DescribeInventorySummaryResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeInventorySummaryResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeInventorySummaryResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeInventorySummaryResult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o DescribeInventorySummaryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeInventorySummaryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["EnvironmentCount"] = o.EnvironmentCount
	toSerialize["ImagesCount"] = o.ImagesCount
	toSerialize["InfraConfigCount"] = o.InfraConfigCount
	toSerialize["InstancesCount"] = o.InstancesCount
	toSerialize["OrganizationCount"] = o.OrganizationCount
	toSerialize["ResourceCount"] = o.ResourceCount
	toSerialize["UserCount"] = o.UserCount
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *DescribeInventorySummaryResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"EnvironmentCount",
		"ImagesCount",
		"InfraConfigCount",
		"InstancesCount",
		"OrganizationCount",
		"ResourceCount",
		"UserCount",
		"environmentId",
		"serviceId",
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

	varDescribeInventorySummaryResult := _DescribeInventorySummaryResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeInventorySummaryResult)

	if err != nil {
		return err
	}

	*o = DescribeInventorySummaryResult(varDescribeInventorySummaryResult)

	return err
}

type NullableDescribeInventorySummaryResult struct {
	value *DescribeInventorySummaryResult
	isSet bool
}

func (v NullableDescribeInventorySummaryResult) Get() *DescribeInventorySummaryResult {
	return v.value
}

func (v *NullableDescribeInventorySummaryResult) Set(val *DescribeInventorySummaryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeInventorySummaryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeInventorySummaryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeInventorySummaryResult(val *DescribeInventorySummaryResult) *NullableDescribeInventorySummaryResult {
	return &NullableDescribeInventorySummaryResult{value: val, isSet: true}
}

func (v NullableDescribeInventorySummaryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeInventorySummaryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


