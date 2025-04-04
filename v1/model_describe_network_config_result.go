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

// checks if the DescribeNetworkConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeNetworkConfigResult{}

// DescribeNetworkConfigResult struct for DescribeNetworkConfigResult
type DescribeNetworkConfigResult struct {
	// A brief description of the network config
	Description string `json:"description"`
	// Generates a DNS endpoint per-replica for this network config
	EndpointPerReplica bool `json:"endpointPerReplica"`
	// ID of a Network Config
	Id string `json:"id"`
	// The list of infra config IDs associated with the compute config.
	InfraConfigIDs []string `json:"infraConfigIDs,omitempty"`
	// Restrict access to this network config to the internal network
	Internal bool `json:"internal"`
	// Name of the network config
	Name string `json:"name"`
	// Ports to map to the generated DNS endpoint
	OpenPorts []int64 `json:"openPorts"`
	PrivateNetworkingConfiguration *PrivateNetworkingConfiguration `json:"privateNetworkingConfiguration,omitempty"`
	PublicNetworkingConfiguration *PublicNetworkingConfiguration `json:"publicNetworkingConfiguration,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// Create an external node load balancer per node rather than expose the node ip directly
	StableEgressIP bool `json:"stableEgressIP"`
	// The port that hosts the reverse proxy for TLS termination
	TlsTerminationPort *int64 `json:"tlsTerminationPort,omitempty"`
	// The preferred type of zonal availability for this resource and the specific zone(s) to deploy in
	ZoneConfiguration string `json:"zoneConfiguration"`
	AdditionalProperties map[string]interface{}
}

type _DescribeNetworkConfigResult DescribeNetworkConfigResult

// NewDescribeNetworkConfigResult instantiates a new DescribeNetworkConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeNetworkConfigResult(description string, endpointPerReplica bool, id string, internal bool, name string, openPorts []int64, serviceId string, stableEgressIP bool, zoneConfiguration string) *DescribeNetworkConfigResult {
	this := DescribeNetworkConfigResult{}
	this.Description = description
	this.EndpointPerReplica = endpointPerReplica
	this.Id = id
	this.Internal = internal
	this.Name = name
	this.OpenPorts = openPorts
	this.ServiceId = serviceId
	this.StableEgressIP = stableEgressIP
	this.ZoneConfiguration = zoneConfiguration
	return &this
}

// NewDescribeNetworkConfigResultWithDefaults instantiates a new DescribeNetworkConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeNetworkConfigResultWithDefaults() *DescribeNetworkConfigResult {
	this := DescribeNetworkConfigResult{}
	return &this
}

// GetDescription returns the Description field value
func (o *DescribeNetworkConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeNetworkConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetEndpointPerReplica returns the EndpointPerReplica field value
func (o *DescribeNetworkConfigResult) GetEndpointPerReplica() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EndpointPerReplica
}

// GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetEndpointPerReplicaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointPerReplica, true
}

// SetEndpointPerReplica sets field value
func (o *DescribeNetworkConfigResult) SetEndpointPerReplica(v bool) {
	o.EndpointPerReplica = v
}

// GetId returns the Id field value
func (o *DescribeNetworkConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeNetworkConfigResult) SetId(v string) {
	o.Id = v
}

// GetInfraConfigIDs returns the InfraConfigIDs field value if set, zero value otherwise.
func (o *DescribeNetworkConfigResult) GetInfraConfigIDs() []string {
	if o == nil || IsNil(o.InfraConfigIDs) {
		var ret []string
		return ret
	}
	return o.InfraConfigIDs
}

// GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetInfraConfigIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.InfraConfigIDs) {
		return nil, false
	}
	return o.InfraConfigIDs, true
}

// SetInfraConfigIDs gets a reference to the given []string and assigns it to the InfraConfigIDs field.
func (o *DescribeNetworkConfigResult) SetInfraConfigIDs(v []string) {
	o.InfraConfigIDs = v
}

// GetInternal returns the Internal field value
func (o *DescribeNetworkConfigResult) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *DescribeNetworkConfigResult) SetInternal(v bool) {
	o.Internal = v
}

// GetName returns the Name field value
func (o *DescribeNetworkConfigResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeNetworkConfigResult) SetName(v string) {
	o.Name = v
}

// GetOpenPorts returns the OpenPorts field value
func (o *DescribeNetworkConfigResult) GetOpenPorts() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OpenPorts
}

// GetOpenPortsOk returns a tuple with the OpenPorts field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetOpenPortsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenPorts, true
}

// SetOpenPorts sets field value
func (o *DescribeNetworkConfigResult) SetOpenPorts(v []int64) {
	o.OpenPorts = v
}

// GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field value if set, zero value otherwise.
func (o *DescribeNetworkConfigResult) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		var ret PrivateNetworkingConfiguration
		return ret
	}
	return *o.PrivateNetworkingConfiguration
}

// GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		return nil, false
	}
	return o.PrivateNetworkingConfiguration, true
}

// SetPrivateNetworkingConfiguration gets a reference to the given PrivateNetworkingConfiguration and assigns it to the PrivateNetworkingConfiguration field.
func (o *DescribeNetworkConfigResult) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration) {
	o.PrivateNetworkingConfiguration = &v
}

// GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field value if set, zero value otherwise.
func (o *DescribeNetworkConfigResult) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		var ret PublicNetworkingConfiguration
		return ret
	}
	return *o.PublicNetworkingConfiguration
}

// GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		return nil, false
	}
	return o.PublicNetworkingConfiguration, true
}

// SetPublicNetworkingConfiguration gets a reference to the given PublicNetworkingConfiguration and assigns it to the PublicNetworkingConfiguration field.
func (o *DescribeNetworkConfigResult) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration) {
	o.PublicNetworkingConfiguration = &v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeNetworkConfigResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeNetworkConfigResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetStableEgressIP returns the StableEgressIP field value
func (o *DescribeNetworkConfigResult) GetStableEgressIP() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StableEgressIP
}

// GetStableEgressIPOk returns a tuple with the StableEgressIP field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetStableEgressIPOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StableEgressIP, true
}

// SetStableEgressIP sets field value
func (o *DescribeNetworkConfigResult) SetStableEgressIP(v bool) {
	o.StableEgressIP = v
}

// GetTlsTerminationPort returns the TlsTerminationPort field value if set, zero value otherwise.
func (o *DescribeNetworkConfigResult) GetTlsTerminationPort() int64 {
	if o == nil || IsNil(o.TlsTerminationPort) {
		var ret int64
		return ret
	}
	return *o.TlsTerminationPort
}

// GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetTlsTerminationPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TlsTerminationPort) {
		return nil, false
	}
	return o.TlsTerminationPort, true
}

// SetTlsTerminationPort gets a reference to the given int64 and assigns it to the TlsTerminationPort field.
func (o *DescribeNetworkConfigResult) SetTlsTerminationPort(v int64) {
	o.TlsTerminationPort = &v
}

// GetZoneConfiguration returns the ZoneConfiguration field value
func (o *DescribeNetworkConfigResult) GetZoneConfiguration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneConfiguration
}

// GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field value
// and a boolean to check if the value has been set.
func (o *DescribeNetworkConfigResult) GetZoneConfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneConfiguration, true
}

// SetZoneConfiguration sets field value
func (o *DescribeNetworkConfigResult) SetZoneConfiguration(v string) {
	o.ZoneConfiguration = v
}

func (o DescribeNetworkConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeNetworkConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["endpointPerReplica"] = o.EndpointPerReplica
	toSerialize["id"] = o.Id
	if !IsNil(o.InfraConfigIDs) {
		toSerialize["infraConfigIDs"] = o.InfraConfigIDs
	}
	toSerialize["internal"] = o.Internal
	toSerialize["name"] = o.Name
	toSerialize["openPorts"] = o.OpenPorts
	if !IsNil(o.PrivateNetworkingConfiguration) {
		toSerialize["privateNetworkingConfiguration"] = o.PrivateNetworkingConfiguration
	}
	if !IsNil(o.PublicNetworkingConfiguration) {
		toSerialize["publicNetworkingConfiguration"] = o.PublicNetworkingConfiguration
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["stableEgressIP"] = o.StableEgressIP
	if !IsNil(o.TlsTerminationPort) {
		toSerialize["tlsTerminationPort"] = o.TlsTerminationPort
	}
	toSerialize["zoneConfiguration"] = o.ZoneConfiguration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeNetworkConfigResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"endpointPerReplica",
		"id",
		"internal",
		"name",
		"openPorts",
		"serviceId",
		"stableEgressIP",
		"zoneConfiguration",
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

	varDescribeNetworkConfigResult := _DescribeNetworkConfigResult{}

	err = json.Unmarshal(data, &varDescribeNetworkConfigResult)

	if err != nil {
		return err
	}

	*o = DescribeNetworkConfigResult(varDescribeNetworkConfigResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "endpointPerReplica")
		delete(additionalProperties, "id")
		delete(additionalProperties, "infraConfigIDs")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "name")
		delete(additionalProperties, "openPorts")
		delete(additionalProperties, "privateNetworkingConfiguration")
		delete(additionalProperties, "publicNetworkingConfiguration")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "stableEgressIP")
		delete(additionalProperties, "tlsTerminationPort")
		delete(additionalProperties, "zoneConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeNetworkConfigResult struct {
	value *DescribeNetworkConfigResult
	isSet bool
}

func (v NullableDescribeNetworkConfigResult) Get() *DescribeNetworkConfigResult {
	return v.value
}

func (v *NullableDescribeNetworkConfigResult) Set(val *DescribeNetworkConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeNetworkConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeNetworkConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeNetworkConfigResult(val *DescribeNetworkConfigResult) *NullableDescribeNetworkConfigResult {
	return &NullableDescribeNetworkConfigResult{value: val, isSet: true}
}

func (v NullableDescribeNetworkConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeNetworkConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


