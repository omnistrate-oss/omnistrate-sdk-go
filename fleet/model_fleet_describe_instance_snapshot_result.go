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

// checks if the FleetDescribeInstanceSnapshotResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetDescribeInstanceSnapshotResult{}

// FleetDescribeInstanceSnapshotResult struct for FleetDescribeInstanceSnapshotResult
type FleetDescribeInstanceSnapshotResult struct {
	// The snapshot time
	CompleteTime string `json:"completeTime"`
	// The snapshot creation time
	CreatedTime string `json:"createdTime"`
	// Whether the snapshot is encrypted
	Encrypted bool `json:"encrypted"`
	// The service environment ID this workflow belongs to.
	EnvironmentId string `json:"environmentId"`
	// The product tier ID
	ProductTierId string `json:"productTierId"`
	// The product tier version
	ProductTierVersion string `json:"productTierVersion"`
	// The backup progress. 0-100
	Progress int64 `json:"progress"`
	// The service ID this workflow belongs to.
	ServiceId string `json:"serviceId"`
	// The instance snapshot ID
	SnapshotId string `json:"snapshotId"`
	// The source instance ID
	SourceInstanceId string `json:"sourceInstanceId"`
	// The snapshot status
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _FleetDescribeInstanceSnapshotResult FleetDescribeInstanceSnapshotResult

// NewFleetDescribeInstanceSnapshotResult instantiates a new FleetDescribeInstanceSnapshotResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetDescribeInstanceSnapshotResult(completeTime string, createdTime string, encrypted bool, environmentId string, productTierId string, productTierVersion string, progress int64, serviceId string, snapshotId string, sourceInstanceId string, status string) *FleetDescribeInstanceSnapshotResult {
	this := FleetDescribeInstanceSnapshotResult{}
	this.CompleteTime = completeTime
	this.CreatedTime = createdTime
	this.Encrypted = encrypted
	this.EnvironmentId = environmentId
	this.ProductTierId = productTierId
	this.ProductTierVersion = productTierVersion
	this.Progress = progress
	this.ServiceId = serviceId
	this.SnapshotId = snapshotId
	this.SourceInstanceId = sourceInstanceId
	this.Status = status
	return &this
}

// NewFleetDescribeInstanceSnapshotResultWithDefaults instantiates a new FleetDescribeInstanceSnapshotResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetDescribeInstanceSnapshotResultWithDefaults() *FleetDescribeInstanceSnapshotResult {
	this := FleetDescribeInstanceSnapshotResult{}
	return &this
}

// GetCompleteTime returns the CompleteTime field value
func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompleteTime
}

// GetCompleteTimeOk returns a tuple with the CompleteTime field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompleteTime, true
}

// SetCompleteTime sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetCompleteTime(v string) {
	o.CompleteTime = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetEncrypted returns the Encrypted field value
func (o *FleetDescribeInstanceSnapshotResult) GetEncrypted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encrypted, true
}

// SetEncrypted sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetEncrypted(v bool) {
	o.Encrypted = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetProductTierId returns the ProductTierId field value
func (o *FleetDescribeInstanceSnapshotResult) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierVersion returns the ProductTierVersion field value
func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierVersion, true
}

// SetProductTierVersion sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetProductTierVersion(v string) {
	o.ProductTierVersion = v
}

// GetProgress returns the Progress field value
func (o *FleetDescribeInstanceSnapshotResult) GetProgress() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetProgressOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetProgress(v int64) {
	o.Progress = v
}

// GetServiceId returns the ServiceId field value
func (o *FleetDescribeInstanceSnapshotResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetSnapshotId(v string) {
	o.SnapshotId = v
}

// GetSourceInstanceId returns the SourceInstanceId field value
func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceInstanceId
}

// GetSourceInstanceIdOk returns a tuple with the SourceInstanceId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceInstanceId, true
}

// SetSourceInstanceId sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetSourceInstanceId(v string) {
	o.SourceInstanceId = v
}

// GetStatus returns the Status field value
func (o *FleetDescribeInstanceSnapshotResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeInstanceSnapshotResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FleetDescribeInstanceSnapshotResult) SetStatus(v string) {
	o.Status = v
}

func (o FleetDescribeInstanceSnapshotResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetDescribeInstanceSnapshotResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completeTime"] = o.CompleteTime
	toSerialize["createdTime"] = o.CreatedTime
	toSerialize["encrypted"] = o.Encrypted
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["productTierVersion"] = o.ProductTierVersion
	toSerialize["progress"] = o.Progress
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["snapshotId"] = o.SnapshotId
	toSerialize["sourceInstanceId"] = o.SourceInstanceId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetDescribeInstanceSnapshotResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completeTime",
		"createdTime",
		"encrypted",
		"environmentId",
		"productTierId",
		"productTierVersion",
		"progress",
		"serviceId",
		"snapshotId",
		"sourceInstanceId",
		"status",
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

	varFleetDescribeInstanceSnapshotResult := _FleetDescribeInstanceSnapshotResult{}

	err = json.Unmarshal(data, &varFleetDescribeInstanceSnapshotResult)

	if err != nil {
		return err
	}

	*o = FleetDescribeInstanceSnapshotResult(varFleetDescribeInstanceSnapshotResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completeTime")
		delete(additionalProperties, "createdTime")
		delete(additionalProperties, "encrypted")
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "productTierVersion")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "snapshotId")
		delete(additionalProperties, "sourceInstanceId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetDescribeInstanceSnapshotResult struct {
	value *FleetDescribeInstanceSnapshotResult
	isSet bool
}

func (v NullableFleetDescribeInstanceSnapshotResult) Get() *FleetDescribeInstanceSnapshotResult {
	return v.value
}

func (v *NullableFleetDescribeInstanceSnapshotResult) Set(val *FleetDescribeInstanceSnapshotResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetDescribeInstanceSnapshotResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetDescribeInstanceSnapshotResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetDescribeInstanceSnapshotResult(val *FleetDescribeInstanceSnapshotResult) *NullableFleetDescribeInstanceSnapshotResult {
	return &NullableFleetDescribeInstanceSnapshotResult{value: val, isSet: true}
}

func (v NullableFleetDescribeInstanceSnapshotResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetDescribeInstanceSnapshotResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


