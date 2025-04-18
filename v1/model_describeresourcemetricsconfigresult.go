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

// checks if the Describeresourcemetricsconfigresult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Describeresourcemetricsconfigresult{}

// Describeresourcemetricsconfigresult struct for Describeresourcemetricsconfigresult
type Describeresourcemetricsconfigresult struct {
	// ID of a resource
	Id string `json:"id"`
	// The local host endpoint to supply prometheus metric
	MetricEndpoint string `json:"metricEndpoint"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	AdditionalProperties map[string]interface{}
}

type _Describeresourcemetricsconfigresult Describeresourcemetricsconfigresult

// NewDescriberesourcemetricsconfigresult instantiates a new Describeresourcemetricsconfigresult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescriberesourcemetricsconfigresult(id string, metricEndpoint string, serviceId string) *Describeresourcemetricsconfigresult {
	this := Describeresourcemetricsconfigresult{}
	this.Id = id
	this.MetricEndpoint = metricEndpoint
	this.ServiceId = serviceId
	return &this
}

// NewDescriberesourcemetricsconfigresultWithDefaults instantiates a new Describeresourcemetricsconfigresult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescriberesourcemetricsconfigresultWithDefaults() *Describeresourcemetricsconfigresult {
	this := Describeresourcemetricsconfigresult{}
	return &this
}

// GetId returns the Id field value
func (o *Describeresourcemetricsconfigresult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Describeresourcemetricsconfigresult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Describeresourcemetricsconfigresult) SetId(v string) {
	o.Id = v
}

// GetMetricEndpoint returns the MetricEndpoint field value
func (o *Describeresourcemetricsconfigresult) GetMetricEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricEndpoint
}

// GetMetricEndpointOk returns a tuple with the MetricEndpoint field value
// and a boolean to check if the value has been set.
func (o *Describeresourcemetricsconfigresult) GetMetricEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricEndpoint, true
}

// SetMetricEndpoint sets field value
func (o *Describeresourcemetricsconfigresult) SetMetricEndpoint(v string) {
	o.MetricEndpoint = v
}

// GetServiceId returns the ServiceId field value
func (o *Describeresourcemetricsconfigresult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *Describeresourcemetricsconfigresult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *Describeresourcemetricsconfigresult) SetServiceId(v string) {
	o.ServiceId = v
}

func (o Describeresourcemetricsconfigresult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Describeresourcemetricsconfigresult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["metricEndpoint"] = o.MetricEndpoint
	toSerialize["serviceId"] = o.ServiceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Describeresourcemetricsconfigresult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"metricEndpoint",
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

	varDescriberesourcemetricsconfigresult := _Describeresourcemetricsconfigresult{}

	err = json.Unmarshal(data, &varDescriberesourcemetricsconfigresult)

	if err != nil {
		return err
	}

	*o = Describeresourcemetricsconfigresult(varDescriberesourcemetricsconfigresult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "metricEndpoint")
		delete(additionalProperties, "serviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescriberesourcemetricsconfigresult struct {
	value *Describeresourcemetricsconfigresult
	isSet bool
}

func (v NullableDescriberesourcemetricsconfigresult) Get() *Describeresourcemetricsconfigresult {
	return v.value
}

func (v *NullableDescriberesourcemetricsconfigresult) Set(val *Describeresourcemetricsconfigresult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescriberesourcemetricsconfigresult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescriberesourcemetricsconfigresult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescriberesourcemetricsconfigresult(val *Describeresourcemetricsconfigresult) *NullableDescriberesourcemetricsconfigresult {
	return &NullableDescriberesourcemetricsconfigresult{value: val, isSet: true}
}

func (v NullableDescriberesourcemetricsconfigresult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescriberesourcemetricsconfigresult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


