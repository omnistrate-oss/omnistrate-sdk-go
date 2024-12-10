/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the NodeNetworkTopologyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeNetworkTopologyResult{}

// NodeNetworkTopologyResult struct for NodeNetworkTopologyResult
type NodeNetworkTopologyResult struct {
	// The availability zone of the node
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	DetailedHealth *DetailedNodeHealthResult `json:"detailedHealth,omitempty"`
	// The endpoint of the node
	Endpoint *string `json:"endpoint,omitempty"`
	// The heath status of a resource
	HealthStatus *string `json:"healthStatus,omitempty"`
	// The ID of the node
	Id *string `json:"id,omitempty"`
	KubernetesDashboardEndpoint *KubernetesDashboardEndpoint `json:"kubernetesDashboardEndpoint,omitempty"`
	// The ports that this node exposes
	Ports []int64 `json:"ports,omitempty"`
	// The status of an operation
	Status *string `json:"status,omitempty"`
	// The storage size of the node in GiB
	StorageSize *int64 `json:"storageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NodeNetworkTopologyResult NodeNetworkTopologyResult

// NewNodeNetworkTopologyResult instantiates a new NodeNetworkTopologyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeNetworkTopologyResult() *NodeNetworkTopologyResult {
	this := NodeNetworkTopologyResult{}
	return &this
}

// NewNodeNetworkTopologyResultWithDefaults instantiates a new NodeNetworkTopologyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeNetworkTopologyResultWithDefaults() *NodeNetworkTopologyResult {
	this := NodeNetworkTopologyResult{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *NodeNetworkTopologyResult) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetDetailedHealth returns the DetailedHealth field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetDetailedHealth() DetailedNodeHealthResult {
	if o == nil || IsNil(o.DetailedHealth) {
		var ret DetailedNodeHealthResult
		return ret
	}
	return *o.DetailedHealth
}

// GetDetailedHealthOk returns a tuple with the DetailedHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetDetailedHealthOk() (*DetailedNodeHealthResult, bool) {
	if o == nil || IsNil(o.DetailedHealth) {
		return nil, false
	}
	return o.DetailedHealth, true
}

// HasDetailedHealth returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasDetailedHealth() bool {
	if o != nil && !IsNil(o.DetailedHealth) {
		return true
	}

	return false
}

// SetDetailedHealth gets a reference to the given DetailedNodeHealthResult and assigns it to the DetailedHealth field.
func (o *NodeNetworkTopologyResult) SetDetailedHealth(v DetailedNodeHealthResult) {
	o.DetailedHealth = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *NodeNetworkTopologyResult) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHealthStatus returns the HealthStatus field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetHealthStatus() string {
	if o == nil || IsNil(o.HealthStatus) {
		var ret string
		return ret
	}
	return *o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetHealthStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HealthStatus) {
		return nil, false
	}
	return o.HealthStatus, true
}

// HasHealthStatus returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasHealthStatus() bool {
	if o != nil && !IsNil(o.HealthStatus) {
		return true
	}

	return false
}

// SetHealthStatus gets a reference to the given string and assigns it to the HealthStatus field.
func (o *NodeNetworkTopologyResult) SetHealthStatus(v string) {
	o.HealthStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NodeNetworkTopologyResult) SetId(v string) {
	o.Id = &v
}

// GetKubernetesDashboardEndpoint returns the KubernetesDashboardEndpoint field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetKubernetesDashboardEndpoint() KubernetesDashboardEndpoint {
	if o == nil || IsNil(o.KubernetesDashboardEndpoint) {
		var ret KubernetesDashboardEndpoint
		return ret
	}
	return *o.KubernetesDashboardEndpoint
}

// GetKubernetesDashboardEndpointOk returns a tuple with the KubernetesDashboardEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetKubernetesDashboardEndpointOk() (*KubernetesDashboardEndpoint, bool) {
	if o == nil || IsNil(o.KubernetesDashboardEndpoint) {
		return nil, false
	}
	return o.KubernetesDashboardEndpoint, true
}

// HasKubernetesDashboardEndpoint returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasKubernetesDashboardEndpoint() bool {
	if o != nil && !IsNil(o.KubernetesDashboardEndpoint) {
		return true
	}

	return false
}

// SetKubernetesDashboardEndpoint gets a reference to the given KubernetesDashboardEndpoint and assigns it to the KubernetesDashboardEndpoint field.
func (o *NodeNetworkTopologyResult) SetKubernetesDashboardEndpoint(v KubernetesDashboardEndpoint) {
	o.KubernetesDashboardEndpoint = &v
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

// HasPorts returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int64 and assigns it to the Ports field.
func (o *NodeNetworkTopologyResult) SetPorts(v []int64) {
	o.Ports = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NodeNetworkTopologyResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkTopologyResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NodeNetworkTopologyResult) SetStatus(v string) {
	o.Status = &v
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

// HasStorageSize returns a boolean if a field has been set.
func (o *NodeNetworkTopologyResult) HasStorageSize() bool {
	if o != nil && !IsNil(o.StorageSize) {
		return true
	}

	return false
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
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.DetailedHealth) {
		toSerialize["detailedHealth"] = o.DetailedHealth
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.HealthStatus) {
		toSerialize["healthStatus"] = o.HealthStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.KubernetesDashboardEndpoint) {
		toSerialize["kubernetesDashboardEndpoint"] = o.KubernetesDashboardEndpoint
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StorageSize) {
		toSerialize["storageSize"] = o.StorageSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NodeNetworkTopologyResult) UnmarshalJSON(data []byte) (err error) {
	varNodeNetworkTopologyResult := _NodeNetworkTopologyResult{}

	err = json.Unmarshal(data, &varNodeNetworkTopologyResult)

	if err != nil {
		return err
	}

	*o = NodeNetworkTopologyResult(varNodeNetworkTopologyResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "availabilityZone")
		delete(additionalProperties, "detailedHealth")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "healthStatus")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kubernetesDashboardEndpoint")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "status")
		delete(additionalProperties, "storageSize")
		o.AdditionalProperties = additionalProperties
	}

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


