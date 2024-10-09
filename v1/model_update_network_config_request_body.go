/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the UpdateNetworkConfigRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkConfigRequestBody{}

// UpdateNetworkConfigRequestBody struct for UpdateNetworkConfigRequestBody
type UpdateNetworkConfigRequestBody struct {
	// A brief description of the network config
	Description *string `json:"description,omitempty"`
	// Generates a DNS endpoint per-replica for this network config
	EndpointPerReplica *bool `json:"endpointPerReplica,omitempty"`
	// Restrict access to this network config to the internal network
	Internal *bool `json:"internal,omitempty"`
	// Name of the network config
	Name *string `json:"name,omitempty"`
	// Ports to map to the generated DNS endpoint
	OpenPorts []int64 `json:"openPorts,omitempty"`
	PrivateNetworkingConfiguration *PrivateNetworkingConfiguration `json:"privateNetworkingConfiguration,omitempty"`
	PublicNetworkingConfiguration *PublicNetworkingConfiguration `json:"publicNetworkingConfiguration,omitempty"`
	// Create an external node load balancer per node rather than expose the node ip directly
	StableEgressIP *bool `json:"stableEgressIP,omitempty"`
	// The port that hosts the reverse proxy for TLS termination
	TlsTerminationPort *int64 `json:"tlsTerminationPort,omitempty"`
	// The preferred type of zonal availability for this resource and the specific zone(s) to deploy in
	ZoneConfiguration *string `json:"zoneConfiguration,omitempty"`
}

// NewUpdateNetworkConfigRequestBody instantiates a new UpdateNetworkConfigRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkConfigRequestBody() *UpdateNetworkConfigRequestBody {
	this := UpdateNetworkConfigRequestBody{}
	return &this
}

// NewUpdateNetworkConfigRequestBodyWithDefaults instantiates a new UpdateNetworkConfigRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkConfigRequestBodyWithDefaults() *UpdateNetworkConfigRequestBody {
	this := UpdateNetworkConfigRequestBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkConfigRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetEndpointPerReplica returns the EndpointPerReplica field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetEndpointPerReplica() bool {
	if o == nil || IsNil(o.EndpointPerReplica) {
		var ret bool
		return ret
	}
	return *o.EndpointPerReplica
}

// GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetEndpointPerReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.EndpointPerReplica) {
		return nil, false
	}
	return o.EndpointPerReplica, true
}

// SetEndpointPerReplica gets a reference to the given bool and assigns it to the EndpointPerReplica field.
func (o *UpdateNetworkConfigRequestBody) SetEndpointPerReplica(v bool) {
	o.EndpointPerReplica = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *UpdateNetworkConfigRequestBody) SetInternal(v bool) {
	o.Internal = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkConfigRequestBody) SetName(v string) {
	o.Name = &v
}

// GetOpenPorts returns the OpenPorts field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetOpenPorts() []int64 {
	if o == nil || IsNil(o.OpenPorts) {
		var ret []int64
		return ret
	}
	return o.OpenPorts
}

// GetOpenPortsOk returns a tuple with the OpenPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetOpenPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.OpenPorts) {
		return nil, false
	}
	return o.OpenPorts, true
}

// SetOpenPorts gets a reference to the given []int64 and assigns it to the OpenPorts field.
func (o *UpdateNetworkConfigRequestBody) SetOpenPorts(v []int64) {
	o.OpenPorts = v
}

// GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		var ret PrivateNetworkingConfiguration
		return ret
	}
	return *o.PrivateNetworkingConfiguration
}

// GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		return nil, false
	}
	return o.PrivateNetworkingConfiguration, true
}

// SetPrivateNetworkingConfiguration gets a reference to the given PrivateNetworkingConfiguration and assigns it to the PrivateNetworkingConfiguration field.
func (o *UpdateNetworkConfigRequestBody) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration) {
	o.PrivateNetworkingConfiguration = &v
}

// GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		var ret PublicNetworkingConfiguration
		return ret
	}
	return *o.PublicNetworkingConfiguration
}

// GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		return nil, false
	}
	return o.PublicNetworkingConfiguration, true
}

// SetPublicNetworkingConfiguration gets a reference to the given PublicNetworkingConfiguration and assigns it to the PublicNetworkingConfiguration field.
func (o *UpdateNetworkConfigRequestBody) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration) {
	o.PublicNetworkingConfiguration = &v
}

// GetStableEgressIP returns the StableEgressIP field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetStableEgressIP() bool {
	if o == nil || IsNil(o.StableEgressIP) {
		var ret bool
		return ret
	}
	return *o.StableEgressIP
}

// GetStableEgressIPOk returns a tuple with the StableEgressIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetStableEgressIPOk() (*bool, bool) {
	if o == nil || IsNil(o.StableEgressIP) {
		return nil, false
	}
	return o.StableEgressIP, true
}

// SetStableEgressIP gets a reference to the given bool and assigns it to the StableEgressIP field.
func (o *UpdateNetworkConfigRequestBody) SetStableEgressIP(v bool) {
	o.StableEgressIP = &v
}

// GetTlsTerminationPort returns the TlsTerminationPort field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetTlsTerminationPort() int64 {
	if o == nil || IsNil(o.TlsTerminationPort) {
		var ret int64
		return ret
	}
	return *o.TlsTerminationPort
}

// GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetTlsTerminationPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TlsTerminationPort) {
		return nil, false
	}
	return o.TlsTerminationPort, true
}

// SetTlsTerminationPort gets a reference to the given int64 and assigns it to the TlsTerminationPort field.
func (o *UpdateNetworkConfigRequestBody) SetTlsTerminationPort(v int64) {
	o.TlsTerminationPort = &v
}

// GetZoneConfiguration returns the ZoneConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequestBody) GetZoneConfiguration() string {
	if o == nil || IsNil(o.ZoneConfiguration) {
		var ret string
		return ret
	}
	return *o.ZoneConfiguration
}

// GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequestBody) GetZoneConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneConfiguration) {
		return nil, false
	}
	return o.ZoneConfiguration, true
}

// SetZoneConfiguration gets a reference to the given string and assigns it to the ZoneConfiguration field.
func (o *UpdateNetworkConfigRequestBody) SetZoneConfiguration(v string) {
	o.ZoneConfiguration = &v
}

func (o UpdateNetworkConfigRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkConfigRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EndpointPerReplica) {
		toSerialize["endpointPerReplica"] = o.EndpointPerReplica
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OpenPorts) {
		toSerialize["openPorts"] = o.OpenPorts
	}
	if !IsNil(o.PrivateNetworkingConfiguration) {
		toSerialize["privateNetworkingConfiguration"] = o.PrivateNetworkingConfiguration
	}
	if !IsNil(o.PublicNetworkingConfiguration) {
		toSerialize["publicNetworkingConfiguration"] = o.PublicNetworkingConfiguration
	}
	if !IsNil(o.StableEgressIP) {
		toSerialize["stableEgressIP"] = o.StableEgressIP
	}
	if !IsNil(o.TlsTerminationPort) {
		toSerialize["tlsTerminationPort"] = o.TlsTerminationPort
	}
	if !IsNil(o.ZoneConfiguration) {
		toSerialize["zoneConfiguration"] = o.ZoneConfiguration
	}
	return toSerialize, nil
}

type NullableUpdateNetworkConfigRequestBody struct {
	value *UpdateNetworkConfigRequestBody
	isSet bool
}

func (v NullableUpdateNetworkConfigRequestBody) Get() *UpdateNetworkConfigRequestBody {
	return v.value
}

func (v *NullableUpdateNetworkConfigRequestBody) Set(val *UpdateNetworkConfigRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkConfigRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkConfigRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkConfigRequestBody(val *UpdateNetworkConfigRequestBody) *NullableUpdateNetworkConfigRequestBody {
	return &NullableUpdateNetworkConfigRequestBody{value: val, isSet: true}
}

func (v NullableUpdateNetworkConfigRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkConfigRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

