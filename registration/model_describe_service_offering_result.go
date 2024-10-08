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

// checks if the DescribeServiceOfferingResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceOfferingResult{}

// DescribeServiceOfferingResult struct for DescribeServiceOfferingResult
type DescribeServiceOfferingResult struct {
	// The time the service was created
	CreatedAt string `json:"createdAt"`
	// Whether the service offering is deprecated
	IsDeprecated bool `json:"isDeprecated"`
	// The service offerings
	Offerings []ServiceOffering `json:"offerings"`
	// The service description
	ServiceDescription string `json:"serviceDescription"`
	// The service id
	ServiceId string `json:"serviceId"`
	// The service name
	ServiceName string `json:"serviceName"`
	// The org id of the service
	ServiceOrgId string `json:"serviceOrgId"`
	// The id of the service provider
	ServiceProviderId string `json:"serviceProviderId"`
	// The name of the service provider
	ServiceProviderName string `json:"serviceProviderName"`
	// The service URL key
	ServiceURLKey string `json:"serviceURLKey"`
}

type _DescribeServiceOfferingResult DescribeServiceOfferingResult

// NewDescribeServiceOfferingResult instantiates a new DescribeServiceOfferingResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceOfferingResult(createdAt string, isDeprecated bool, offerings []ServiceOffering, serviceDescription string, serviceId string, serviceName string, serviceOrgId string, serviceProviderId string, serviceProviderName string, serviceURLKey string) *DescribeServiceOfferingResult {
	this := DescribeServiceOfferingResult{}
	this.CreatedAt = createdAt
	this.IsDeprecated = isDeprecated
	this.Offerings = offerings
	this.ServiceDescription = serviceDescription
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.ServiceOrgId = serviceOrgId
	this.ServiceProviderId = serviceProviderId
	this.ServiceProviderName = serviceProviderName
	this.ServiceURLKey = serviceURLKey
	return &this
}

// NewDescribeServiceOfferingResultWithDefaults instantiates a new DescribeServiceOfferingResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceOfferingResultWithDefaults() *DescribeServiceOfferingResult {
	this := DescribeServiceOfferingResult{}
	var isDeprecated bool = false
	this.IsDeprecated = isDeprecated
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *DescribeServiceOfferingResult) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DescribeServiceOfferingResult) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *DescribeServiceOfferingResult) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *DescribeServiceOfferingResult) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetOfferings returns the Offerings field value
func (o *DescribeServiceOfferingResult) GetOfferings() []ServiceOffering {
	if o == nil {
		var ret []ServiceOffering
		return ret
	}

	return o.Offerings
}

// GetOfferingsOk returns a tuple with the Offerings field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetOfferingsOk() ([]ServiceOffering, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offerings, true
}

// SetOfferings sets field value
func (o *DescribeServiceOfferingResult) SetOfferings(v []ServiceOffering) {
	o.Offerings = v
}

// GetServiceDescription returns the ServiceDescription field value
func (o *DescribeServiceOfferingResult) GetServiceDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceDescription
}

// GetServiceDescriptionOk returns a tuple with the ServiceDescription field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDescription, true
}

// SetServiceDescription sets field value
func (o *DescribeServiceOfferingResult) SetServiceDescription(v string) {
	o.ServiceDescription = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceOfferingResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceOfferingResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value
func (o *DescribeServiceOfferingResult) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *DescribeServiceOfferingResult) SetServiceName(v string) {
	o.ServiceName = v
}

// GetServiceOrgId returns the ServiceOrgId field value
func (o *DescribeServiceOfferingResult) GetServiceOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceOrgId
}

// GetServiceOrgIdOk returns a tuple with the ServiceOrgId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceOrgId, true
}

// SetServiceOrgId sets field value
func (o *DescribeServiceOfferingResult) SetServiceOrgId(v string) {
	o.ServiceOrgId = v
}

// GetServiceProviderId returns the ServiceProviderId field value
func (o *DescribeServiceOfferingResult) GetServiceProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProviderId
}

// GetServiceProviderIdOk returns a tuple with the ServiceProviderId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProviderId, true
}

// SetServiceProviderId sets field value
func (o *DescribeServiceOfferingResult) SetServiceProviderId(v string) {
	o.ServiceProviderId = v
}

// GetServiceProviderName returns the ServiceProviderName field value
func (o *DescribeServiceOfferingResult) GetServiceProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProviderName
}

// GetServiceProviderNameOk returns a tuple with the ServiceProviderName field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProviderName, true
}

// SetServiceProviderName sets field value
func (o *DescribeServiceOfferingResult) SetServiceProviderName(v string) {
	o.ServiceProviderName = v
}

// GetServiceURLKey returns the ServiceURLKey field value
func (o *DescribeServiceOfferingResult) GetServiceURLKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceURLKey
}

// GetServiceURLKeyOk returns a tuple with the ServiceURLKey field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceOfferingResult) GetServiceURLKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceURLKey, true
}

// SetServiceURLKey sets field value
func (o *DescribeServiceOfferingResult) SetServiceURLKey(v string) {
	o.ServiceURLKey = v
}

func (o DescribeServiceOfferingResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceOfferingResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["isDeprecated"] = o.IsDeprecated
	toSerialize["offerings"] = o.Offerings
	toSerialize["serviceDescription"] = o.ServiceDescription
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["serviceOrgId"] = o.ServiceOrgId
	toSerialize["serviceProviderId"] = o.ServiceProviderId
	toSerialize["serviceProviderName"] = o.ServiceProviderName
	toSerialize["serviceURLKey"] = o.ServiceURLKey
	return toSerialize, nil
}

func (o *DescribeServiceOfferingResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"isDeprecated",
		"offerings",
		"serviceDescription",
		"serviceId",
		"serviceName",
		"serviceOrgId",
		"serviceProviderId",
		"serviceProviderName",
		"serviceURLKey",
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

	varDescribeServiceOfferingResult := _DescribeServiceOfferingResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeServiceOfferingResult)

	if err != nil {
		return err
	}

	*o = DescribeServiceOfferingResult(varDescribeServiceOfferingResult)

	return err
}

type NullableDescribeServiceOfferingResult struct {
	value *DescribeServiceOfferingResult
	isSet bool
}

func (v NullableDescribeServiceOfferingResult) Get() *DescribeServiceOfferingResult {
	return v.value
}

func (v *NullableDescribeServiceOfferingResult) Set(val *DescribeServiceOfferingResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceOfferingResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceOfferingResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceOfferingResult(val *DescribeServiceOfferingResult) *NullableDescribeServiceOfferingResult {
	return &NullableDescribeServiceOfferingResult{value: val, isSet: true}
}

func (v NullableDescribeServiceOfferingResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceOfferingResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


