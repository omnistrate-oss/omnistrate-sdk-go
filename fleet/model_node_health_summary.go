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

// checks if the NodeHealthSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeHealthSummary{}

// NodeHealthSummary struct for NodeHealthSummary
type NodeHealthSummary struct {
	// The availability zone of the node
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// The heath status of a resource
	ConnectivityStatus *string `json:"connectivityStatus,omitempty"`
	// The heath status of a resource
	DiskHealth *string `json:"diskHealth,omitempty"`
	// The endpoint of the node
	Endpoint *string `json:"endpoint,omitempty"`
	IntegrationsHealth *IntegrationsHealth `json:"integrationsHealth,omitempty"`
	// The load status of a pod
	LoadStatus *string `json:"loadStatus,omitempty"`
	// The heath status of a resource
	NodeHealth *string `json:"nodeHealth,omitempty"`
	// The name of the node
	NodeName *string `json:"nodeName,omitempty"`
	// The ports that this node exposes
	Ports []int64 `json:"ports,omitempty"`
	// The heath status of a resource
	ProcessHealth *string `json:"processHealth,omitempty"`
	// The heath status of a resource
	ProcessLiveness *string `json:"processLiveness,omitempty"`
	// The heath status of a resource
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NodeHealthSummary NodeHealthSummary

// NewNodeHealthSummary instantiates a new NodeHealthSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeHealthSummary() *NodeHealthSummary {
	this := NodeHealthSummary{}
	return &this
}

// NewNodeHealthSummaryWithDefaults instantiates a new NodeHealthSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeHealthSummaryWithDefaults() *NodeHealthSummary {
	this := NodeHealthSummary{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *NodeHealthSummary) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetConnectivityStatus returns the ConnectivityStatus field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetConnectivityStatus() string {
	if o == nil || IsNil(o.ConnectivityStatus) {
		var ret string
		return ret
	}
	return *o.ConnectivityStatus
}

// GetConnectivityStatusOk returns a tuple with the ConnectivityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetConnectivityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectivityStatus) {
		return nil, false
	}
	return o.ConnectivityStatus, true
}

// HasConnectivityStatus returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasConnectivityStatus() bool {
	if o != nil && !IsNil(o.ConnectivityStatus) {
		return true
	}

	return false
}

// SetConnectivityStatus gets a reference to the given string and assigns it to the ConnectivityStatus field.
func (o *NodeHealthSummary) SetConnectivityStatus(v string) {
	o.ConnectivityStatus = &v
}

// GetDiskHealth returns the DiskHealth field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetDiskHealth() string {
	if o == nil || IsNil(o.DiskHealth) {
		var ret string
		return ret
	}
	return *o.DiskHealth
}

// GetDiskHealthOk returns a tuple with the DiskHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetDiskHealthOk() (*string, bool) {
	if o == nil || IsNil(o.DiskHealth) {
		return nil, false
	}
	return o.DiskHealth, true
}

// HasDiskHealth returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasDiskHealth() bool {
	if o != nil && !IsNil(o.DiskHealth) {
		return true
	}

	return false
}

// SetDiskHealth gets a reference to the given string and assigns it to the DiskHealth field.
func (o *NodeHealthSummary) SetDiskHealth(v string) {
	o.DiskHealth = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *NodeHealthSummary) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetIntegrationsHealth returns the IntegrationsHealth field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetIntegrationsHealth() IntegrationsHealth {
	if o == nil || IsNil(o.IntegrationsHealth) {
		var ret IntegrationsHealth
		return ret
	}
	return *o.IntegrationsHealth
}

// GetIntegrationsHealthOk returns a tuple with the IntegrationsHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetIntegrationsHealthOk() (*IntegrationsHealth, bool) {
	if o == nil || IsNil(o.IntegrationsHealth) {
		return nil, false
	}
	return o.IntegrationsHealth, true
}

// HasIntegrationsHealth returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasIntegrationsHealth() bool {
	if o != nil && !IsNil(o.IntegrationsHealth) {
		return true
	}

	return false
}

// SetIntegrationsHealth gets a reference to the given IntegrationsHealth and assigns it to the IntegrationsHealth field.
func (o *NodeHealthSummary) SetIntegrationsHealth(v IntegrationsHealth) {
	o.IntegrationsHealth = &v
}

// GetLoadStatus returns the LoadStatus field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetLoadStatus() string {
	if o == nil || IsNil(o.LoadStatus) {
		var ret string
		return ret
	}
	return *o.LoadStatus
}

// GetLoadStatusOk returns a tuple with the LoadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetLoadStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LoadStatus) {
		return nil, false
	}
	return o.LoadStatus, true
}

// HasLoadStatus returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasLoadStatus() bool {
	if o != nil && !IsNil(o.LoadStatus) {
		return true
	}

	return false
}

// SetLoadStatus gets a reference to the given string and assigns it to the LoadStatus field.
func (o *NodeHealthSummary) SetLoadStatus(v string) {
	o.LoadStatus = &v
}

// GetNodeHealth returns the NodeHealth field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetNodeHealth() string {
	if o == nil || IsNil(o.NodeHealth) {
		var ret string
		return ret
	}
	return *o.NodeHealth
}

// GetNodeHealthOk returns a tuple with the NodeHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetNodeHealthOk() (*string, bool) {
	if o == nil || IsNil(o.NodeHealth) {
		return nil, false
	}
	return o.NodeHealth, true
}

// HasNodeHealth returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasNodeHealth() bool {
	if o != nil && !IsNil(o.NodeHealth) {
		return true
	}

	return false
}

// SetNodeHealth gets a reference to the given string and assigns it to the NodeHealth field.
func (o *NodeHealthSummary) SetNodeHealth(v string) {
	o.NodeHealth = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *NodeHealthSummary) SetNodeName(v string) {
	o.NodeName = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetPorts() []int64 {
	if o == nil || IsNil(o.Ports) {
		var ret []int64
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int64 and assigns it to the Ports field.
func (o *NodeHealthSummary) SetPorts(v []int64) {
	o.Ports = v
}

// GetProcessHealth returns the ProcessHealth field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetProcessHealth() string {
	if o == nil || IsNil(o.ProcessHealth) {
		var ret string
		return ret
	}
	return *o.ProcessHealth
}

// GetProcessHealthOk returns a tuple with the ProcessHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetProcessHealthOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessHealth) {
		return nil, false
	}
	return o.ProcessHealth, true
}

// HasProcessHealth returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasProcessHealth() bool {
	if o != nil && !IsNil(o.ProcessHealth) {
		return true
	}

	return false
}

// SetProcessHealth gets a reference to the given string and assigns it to the ProcessHealth field.
func (o *NodeHealthSummary) SetProcessHealth(v string) {
	o.ProcessHealth = &v
}

// GetProcessLiveness returns the ProcessLiveness field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetProcessLiveness() string {
	if o == nil || IsNil(o.ProcessLiveness) {
		var ret string
		return ret
	}
	return *o.ProcessLiveness
}

// GetProcessLivenessOk returns a tuple with the ProcessLiveness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetProcessLivenessOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessLiveness) {
		return nil, false
	}
	return o.ProcessLiveness, true
}

// HasProcessLiveness returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasProcessLiveness() bool {
	if o != nil && !IsNil(o.ProcessLiveness) {
		return true
	}

	return false
}

// SetProcessLiveness gets a reference to the given string and assigns it to the ProcessLiveness field.
func (o *NodeHealthSummary) SetProcessLiveness(v string) {
	o.ProcessLiveness = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NodeHealthSummary) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthSummary) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NodeHealthSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NodeHealthSummary) SetStatus(v string) {
	o.Status = &v
}

func (o NodeHealthSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeHealthSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.ConnectivityStatus) {
		toSerialize["connectivityStatus"] = o.ConnectivityStatus
	}
	if !IsNil(o.DiskHealth) {
		toSerialize["diskHealth"] = o.DiskHealth
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.IntegrationsHealth) {
		toSerialize["integrationsHealth"] = o.IntegrationsHealth
	}
	if !IsNil(o.LoadStatus) {
		toSerialize["loadStatus"] = o.LoadStatus
	}
	if !IsNil(o.NodeHealth) {
		toSerialize["nodeHealth"] = o.NodeHealth
	}
	if !IsNil(o.NodeName) {
		toSerialize["nodeName"] = o.NodeName
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.ProcessHealth) {
		toSerialize["processHealth"] = o.ProcessHealth
	}
	if !IsNil(o.ProcessLiveness) {
		toSerialize["processLiveness"] = o.ProcessLiveness
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NodeHealthSummary) UnmarshalJSON(data []byte) (err error) {
	varNodeHealthSummary := _NodeHealthSummary{}

	err = json.Unmarshal(data, &varNodeHealthSummary)

	if err != nil {
		return err
	}

	*o = NodeHealthSummary(varNodeHealthSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "availabilityZone")
		delete(additionalProperties, "connectivityStatus")
		delete(additionalProperties, "diskHealth")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "integrationsHealth")
		delete(additionalProperties, "loadStatus")
		delete(additionalProperties, "nodeHealth")
		delete(additionalProperties, "nodeName")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "processHealth")
		delete(additionalProperties, "processLiveness")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNodeHealthSummary struct {
	value *NodeHealthSummary
	isSet bool
}

func (v NullableNodeHealthSummary) Get() *NodeHealthSummary {
	return v.value
}

func (v *NullableNodeHealthSummary) Set(val *NodeHealthSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeHealthSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeHealthSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeHealthSummary(val *NodeHealthSummary) *NullableNodeHealthSummary {
	return &NullableNodeHealthSummary{value: val, isSet: true}
}

func (v NullableNodeHealthSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeHealthSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


