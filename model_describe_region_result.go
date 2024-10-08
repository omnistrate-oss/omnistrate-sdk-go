/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeRegionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeRegionResult{}

// DescribeRegionResult struct for DescribeRegionResult
type DescribeRegionResult struct {
	CloudProviderId string `json:"cloudProviderId"`
	// The cloud provider for this compute instance type config
	CloudProviderName string `json:"cloudProviderName"`
	// Cloud-provider native region code
	Code string `json:"code"`
	// Description of the Region
	Description string `json:"description"`
	// ID of the Region
	Id string `json:"id"`
}

type _DescribeRegionResult DescribeRegionResult

// NewDescribeRegionResult instantiates a new DescribeRegionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeRegionResult(cloudProviderId string, cloudProviderName string, code string, description string, id string) *DescribeRegionResult {
	this := DescribeRegionResult{}
	this.CloudProviderId = cloudProviderId
	this.CloudProviderName = cloudProviderName
	this.Code = code
	this.Description = description
	this.Id = id
	return &this
}

// NewDescribeRegionResultWithDefaults instantiates a new DescribeRegionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeRegionResultWithDefaults() *DescribeRegionResult {
	this := DescribeRegionResult{}
	return &this
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *DescribeRegionResult) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *DescribeRegionResult) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *DescribeRegionResult) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *DescribeRegionResult) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *DescribeRegionResult) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *DescribeRegionResult) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCode returns the Code field value
func (o *DescribeRegionResult) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DescribeRegionResult) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DescribeRegionResult) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *DescribeRegionResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeRegionResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeRegionResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeRegionResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeRegionResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeRegionResult) SetId(v string) {
	o.Id = v
}

func (o DescribeRegionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeRegionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderId"] = o.CloudProviderId
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["code"] = o.Code
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *DescribeRegionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderId",
		"cloudProviderName",
		"code",
		"description",
		"id",
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

	varDescribeRegionResult := _DescribeRegionResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeRegionResult)

	if err != nil {
		return err
	}

	*o = DescribeRegionResult(varDescribeRegionResult)

	return err
}

type NullableDescribeRegionResult struct {
	value *DescribeRegionResult
	isSet bool
}

func (v NullableDescribeRegionResult) Get() *DescribeRegionResult {
	return v.value
}

func (v *NullableDescribeRegionResult) Set(val *DescribeRegionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeRegionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeRegionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeRegionResult(val *DescribeRegionResult) *NullableDescribeRegionResult {
	return &NullableDescribeRegionResult{value: val, isSet: true}
}

func (v NullableDescribeRegionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeRegionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


