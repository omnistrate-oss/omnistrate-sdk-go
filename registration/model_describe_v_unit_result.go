/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeVUnitResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeVUnitResult{}

// DescribeVUnitResult struct for DescribeVUnitResult
type DescribeVUnitResult struct {
	// The infra / cloud provider name
	CloudProvider string `json:"cloudProvider"`
	// VUnit to operate on
	Id *string `json:"id,omitempty"`
	// List of network IDs in the given context
	NetworkIds []string `json:"networkIds"`
	// Region code specific to the cloud-provider
	Region string `json:"region"`
	// Service ID for the VUnit
	ServiceId *string `json:"serviceId,omitempty"`
	// Service Model ID for the VUnit
	ServiceModelId string `json:"serviceModelId"`
}

type _DescribeVUnitResult DescribeVUnitResult

// NewDescribeVUnitResult instantiates a new DescribeVUnitResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeVUnitResult(cloudProvider string, networkIds []string, region string, serviceModelId string) *DescribeVUnitResult {
	this := DescribeVUnitResult{}
	this.CloudProvider = cloudProvider
	this.NetworkIds = networkIds
	this.Region = region
	this.ServiceModelId = serviceModelId
	return &this
}

// NewDescribeVUnitResultWithDefaults instantiates a new DescribeVUnitResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeVUnitResultWithDefaults() *DescribeVUnitResult {
	this := DescribeVUnitResult{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *DescribeVUnitResult) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DescribeVUnitResult) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DescribeVUnitResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DescribeVUnitResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DescribeVUnitResult) SetId(v string) {
	o.Id = &v
}

// GetNetworkIds returns the NetworkIds field value
func (o *DescribeVUnitResult) GetNetworkIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetNetworkIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkIds, true
}

// SetNetworkIds sets field value
func (o *DescribeVUnitResult) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

// GetRegion returns the Region field value
func (o *DescribeVUnitResult) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *DescribeVUnitResult) SetRegion(v string) {
	o.Region = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DescribeVUnitResult) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DescribeVUnitResult) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DescribeVUnitResult) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *DescribeVUnitResult) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *DescribeVUnitResult) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *DescribeVUnitResult) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

func (o DescribeVUnitResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeVUnitResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["networkIds"] = o.NetworkIds
	toSerialize["region"] = o.Region
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	toSerialize["serviceModelId"] = o.ServiceModelId
	return toSerialize, nil
}

func (o *DescribeVUnitResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProvider",
		"networkIds",
		"region",
		"serviceModelId",
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

	varDescribeVUnitResult := _DescribeVUnitResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeVUnitResult)

	if err != nil {
		return err
	}

	*o = DescribeVUnitResult(varDescribeVUnitResult)

	return err
}

type NullableDescribeVUnitResult struct {
	value *DescribeVUnitResult
	isSet bool
}

func (v NullableDescribeVUnitResult) Get() *DescribeVUnitResult {
	return v.value
}

func (v *NullableDescribeVUnitResult) Set(val *DescribeVUnitResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeVUnitResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeVUnitResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeVUnitResult(val *DescribeVUnitResult) *NullableDescribeVUnitResult {
	return &NullableDescribeVUnitResult{value: val, isSet: true}
}

func (v NullableDescribeVUnitResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeVUnitResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


