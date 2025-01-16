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

// checks if the DescribeAvailabilityZoneResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeAvailabilityZoneResult{}

// DescribeAvailabilityZoneResult struct for DescribeAvailabilityZoneResult
type DescribeAvailabilityZoneResult struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	// Cloud-provider native availability zone code
	Code string `json:"code"`
	// Description of the AvailabilityZone
	Description string `json:"description"`
	// ID of an AZ
	Id string `json:"id"`
	// Cloud-provider native region code
	RegionCode string `json:"regionCode"`
	AdditionalProperties map[string]interface{}
}

type _DescribeAvailabilityZoneResult DescribeAvailabilityZoneResult

// NewDescribeAvailabilityZoneResult instantiates a new DescribeAvailabilityZoneResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAvailabilityZoneResult(cloudProviderName string, code string, description string, id string, regionCode string) *DescribeAvailabilityZoneResult {
	this := DescribeAvailabilityZoneResult{}
	this.CloudProviderName = cloudProviderName
	this.Code = code
	this.Description = description
	this.Id = id
	this.RegionCode = regionCode
	return &this
}

// NewDescribeAvailabilityZoneResultWithDefaults instantiates a new DescribeAvailabilityZoneResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAvailabilityZoneResultWithDefaults() *DescribeAvailabilityZoneResult {
	this := DescribeAvailabilityZoneResult{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *DescribeAvailabilityZoneResult) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *DescribeAvailabilityZoneResult) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *DescribeAvailabilityZoneResult) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCode returns the Code field value
func (o *DescribeAvailabilityZoneResult) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DescribeAvailabilityZoneResult) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DescribeAvailabilityZoneResult) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *DescribeAvailabilityZoneResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeAvailabilityZoneResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeAvailabilityZoneResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeAvailabilityZoneResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeAvailabilityZoneResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeAvailabilityZoneResult) SetId(v string) {
	o.Id = v
}

// GetRegionCode returns the RegionCode field value
func (o *DescribeAvailabilityZoneResult) GetRegionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value
// and a boolean to check if the value has been set.
func (o *DescribeAvailabilityZoneResult) GetRegionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionCode, true
}

// SetRegionCode sets field value
func (o *DescribeAvailabilityZoneResult) SetRegionCode(v string) {
	o.RegionCode = v
}

func (o DescribeAvailabilityZoneResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAvailabilityZoneResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["code"] = o.Code
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["regionCode"] = o.RegionCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeAvailabilityZoneResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"code",
		"description",
		"id",
		"regionCode",
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

	varDescribeAvailabilityZoneResult := _DescribeAvailabilityZoneResult{}

	err = json.Unmarshal(data, &varDescribeAvailabilityZoneResult)

	if err != nil {
		return err
	}

	*o = DescribeAvailabilityZoneResult(varDescribeAvailabilityZoneResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "code")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "regionCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeAvailabilityZoneResult struct {
	value *DescribeAvailabilityZoneResult
	isSet bool
}

func (v NullableDescribeAvailabilityZoneResult) Get() *DescribeAvailabilityZoneResult {
	return v.value
}

func (v *NullableDescribeAvailabilityZoneResult) Set(val *DescribeAvailabilityZoneResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAvailabilityZoneResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAvailabilityZoneResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAvailabilityZoneResult(val *DescribeAvailabilityZoneResult) *NullableDescribeAvailabilityZoneResult {
	return &NullableDescribeAvailabilityZoneResult{value: val, isSet: true}
}

func (v NullableDescribeAvailabilityZoneResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAvailabilityZoneResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


