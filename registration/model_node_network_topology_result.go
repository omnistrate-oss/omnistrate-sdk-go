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

// checks if the NodeNetworkTopologyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeNetworkTopologyResult{}

// NodeNetworkTopologyResult struct for NodeNetworkTopologyResult
type NodeNetworkTopologyResult struct {
	// The availability zone of the node
	AvailabilityZone string `json:"availabilityZone"`
	DetailedHealth DetailedNodeHealthResult `json:"detailedHealth"`
	// The endpoint of the node
	Endpoint string `json:"endpoint"`
	// The heath status of the node
	HealthStatus string `json:"healthStatus"`
	// The ID of the node
	Id string `json:"id"`
	// The ports that this node exposes
	Ports []int64 `json:"ports,omitempty"`
	// The status of the node
	Status string `json:"status"`
	// The storage size of the node in GiB
	StorageSize *int64 `json:"storageSize,omitempty"`
}

type _NodeNetworkTopologyResult NodeNetworkTopologyResult

// NewNodeNetworkTopologyResult instantiates a new NodeNetworkTopologyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeNetworkTopologyResult(availabilityZone string, detailedHealth DetailedNodeHealthResult, endpoint string, healthStatus string, id string, status string) *NodeNetworkTopologyResult {
	this := NodeNetworkTopologyResult{}
	this.AvailabilityZone = availabilityZone
	this.DetailedHealth = detailedHealth
	this.Endpoint = endpoint
	this.HealthStatus = healthStatus
	this.Id = id
	this.Status = status
	return &this
}

// NewNodeNetworkTopologyResultWithDefaults instantiates a new NodeNetworkTopologyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeNetworkTopologyResultWithDefaults() *NodeNetworkTopologyResult {
	this := NodeNetworkTopologyResult{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *NodeNetworkTopologyResult) GetAvailabilityZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *NodeNetworkTopologyResult) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetDetailedHealth returns the DetailedHealth field value
func (o *NodeNetworkTopologyResult) GetDetailedHealth() DetailedNodeHealthResult {
	if o == nil {
		var ret DetailedNodeHealthResult
		return ret
	}

	return o.DetailedHealth
}

// GetDetailedHealthOk returns a tuple with the DetailedHealth field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetDetailedHealthOk() (*DetailedNodeHealthResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailedHealth, true
}

// SetDetailedHealth sets field value
func (o *NodeNetworkTopologyResult) SetDetailedHealth(v DetailedNodeHealthResult) {
	o.DetailedHealth = v
}

// GetEndpoint returns the Endpoint field value
func (o *NodeNetworkTopologyResult) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *NodeNetworkTopologyResult) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetHealthStatus returns the HealthStatus field value
func (o *NodeNetworkTopologyResult) GetHealthStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetHealthStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthStatus, true
}

// SetHealthStatus sets field value
func (o *NodeNetworkTopologyResult) SetHealthStatus(v string) {
	o.HealthStatus = v
}

// GetId returns the Id field value
func (o *NodeNetworkTopologyResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NodeNetworkTopologyResult) SetId(v string) {
	o.Id = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetPorts() []int64 {
	if o == nil || IsNil(o.Ports) {
		var ret []int64
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts gets a reference to the given []int64 and assigns it to the Ports field.
func (o *NodeNetworkTopologyResult) SetPorts(v []int64) {
	o.Ports = v
}

// GetStatus returns the Status field value
func (o *NodeNetworkTopologyResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NodeNetworkTopologyResult) SetStatus(v string) {
	o.Status = v
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetStorageSize() int64 {
	if o == nil || IsNil(o.StorageSize) {
		var ret int64
		return ret
	}
	return *o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetStorageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageSize) {
		return nil, false
	}
	return o.StorageSize, true
}

// SetStorageSize gets a reference to the given int64 and assigns it to the StorageSize field.
func (o *NodeNetworkTopologyResult) SetStorageSize(v int64) {
	o.StorageSize = &v
}

func (o NodeNetworkTopologyResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeNetworkTopologyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["availabilityZone"] = o.AvailabilityZone
	toSerialize["detailedHealth"] = o.DetailedHealth
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["healthStatus"] = o.HealthStatus
	toSerialize["id"] = o.Id
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.StorageSize) {
		toSerialize["storageSize"] = o.StorageSize
	}
	return toSerialize, nil
}

func (o *NodeNetworkTopologyResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"availabilityZone",
		"detailedHealth",
		"endpoint",
		"healthStatus",
		"id",
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

	varNodeNetworkTopologyResult := _NodeNetworkTopologyResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeNetworkTopologyResult)

	if err != nil {
		return err
	}

	*o = NodeNetworkTopologyResult(varNodeNetworkTopologyResult)

	return err
}

type NullableNodeNetworkTopologyResult struct {
	value *NodeNetworkTopologyResult
	isSet bool
}

func (v NullableNodeNetworkTopologyResult) Get() *NodeNetworkTopologyResult {
	return v.value
}

func (v *NullableNodeNetworkTopologyResult) Set(val *NodeNetworkTopologyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeNetworkTopologyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeNetworkTopologyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeNetworkTopologyResult(val *NodeNetworkTopologyResult) *NullableNodeNetworkTopologyResult {
	return &NullableNodeNetworkTopologyResult{value: val, isSet: true}
}

func (v NullableNodeNetworkTopologyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeNetworkTopologyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


