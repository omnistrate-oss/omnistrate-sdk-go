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

// checks if the UpdateNetworkConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkConfigRequest{}

// UpdateNetworkConfigRequest struct for UpdateNetworkConfigRequest
type UpdateNetworkConfigRequest struct {
	// A brief description of the network config
	Description *string `json:"description,omitempty"`
	// Generates a DNS endpoint per-replica for this network config
	EndpointPerReplica *bool `json:"endpointPerReplica,omitempty"`
	// ID of a Network Config
	Id string `json:"id"`
	// Restrict access to this network config to the internal network
	Internal *bool `json:"internal,omitempty"`
	// Name of the network config
	Name *string `json:"name,omitempty"`
	// Ports to map to the generated DNS endpoint
	OpenPorts []int64 `json:"openPorts,omitempty"`
	PrivateNetworkingConfiguration *PrivateNetworkingConfiguration `json:"privateNetworkingConfiguration,omitempty"`
	PublicNetworkingConfiguration *PublicNetworkingConfiguration `json:"publicNetworkingConfiguration,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// Create an external node load balancer per node rather than expose the node ip directly
	StableEgressIP *bool `json:"stableEgressIP,omitempty"`
	// The port that hosts the reverse proxy for TLS termination
	TlsTerminationPort *int64 `json:"tlsTerminationPort,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// The preferred type of zonal availability for this resource and the specific zone(s) to deploy in
	ZoneConfiguration *string `json:"zoneConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworkConfigRequest UpdateNetworkConfigRequest

// NewUpdateNetworkConfigRequest instantiates a new UpdateNetworkConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkConfigRequest(id string, serviceId string, token string) *UpdateNetworkConfigRequest {
	this := UpdateNetworkConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateNetworkConfigRequestWithDefaults instantiates a new UpdateNetworkConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkConfigRequestWithDefaults() *UpdateNetworkConfigRequest {
	this := UpdateNetworkConfigRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkConfigRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEndpointPerReplica returns the EndpointPerReplica field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetEndpointPerReplica() bool {
	if o == nil || IsNil(o.EndpointPerReplica) {
		var ret bool
		return ret
	}
	return *o.EndpointPerReplica
}

// GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetEndpointPerReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.EndpointPerReplica) {
		return nil, false
	}
	return o.EndpointPerReplica, true
}

// SetEndpointPerReplica gets a reference to the given bool and assigns it to the EndpointPerReplica field.
func (o *UpdateNetworkConfigRequest) SetEndpointPerReplica(v bool) {
	o.EndpointPerReplica = &v
}

// GetId returns the Id field value
func (o *UpdateNetworkConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkConfigRequest) SetId(v string) {
	o.Id = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *UpdateNetworkConfigRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkConfigRequest) SetName(v string) {
	o.Name = &v
}

// GetOpenPorts returns the OpenPorts field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetOpenPorts() []int64 {
	if o == nil || IsNil(o.OpenPorts) {
		var ret []int64
		return ret
	}
	return o.OpenPorts
}

// GetOpenPortsOk returns a tuple with the OpenPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetOpenPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.OpenPorts) {
		return nil, false
	}
	return o.OpenPorts, true
}

// SetOpenPorts gets a reference to the given []int64 and assigns it to the OpenPorts field.
func (o *UpdateNetworkConfigRequest) SetOpenPorts(v []int64) {
	o.OpenPorts = v
}

// GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		var ret PrivateNetworkingConfiguration
		return ret
	}
	return *o.PrivateNetworkingConfiguration
}

// GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PrivateNetworkingConfiguration) {
		return nil, false
	}
	return o.PrivateNetworkingConfiguration, true
}

// SetPrivateNetworkingConfiguration gets a reference to the given PrivateNetworkingConfiguration and assigns it to the PrivateNetworkingConfiguration field.
func (o *UpdateNetworkConfigRequest) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration) {
	o.PrivateNetworkingConfiguration = &v
}

// GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		var ret PublicNetworkingConfiguration
		return ret
	}
	return *o.PublicNetworkingConfiguration
}

// GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool) {
	if o == nil || IsNil(o.PublicNetworkingConfiguration) {
		return nil, false
	}
	return o.PublicNetworkingConfiguration, true
}

// SetPublicNetworkingConfiguration gets a reference to the given PublicNetworkingConfiguration and assigns it to the PublicNetworkingConfiguration field.
func (o *UpdateNetworkConfigRequest) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration) {
	o.PublicNetworkingConfiguration = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateNetworkConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateNetworkConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetStableEgressIP returns the StableEgressIP field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetStableEgressIP() bool {
	if o == nil || IsNil(o.StableEgressIP) {
		var ret bool
		return ret
	}
	return *o.StableEgressIP
}

// GetStableEgressIPOk returns a tuple with the StableEgressIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetStableEgressIPOk() (*bool, bool) {
	if o == nil || IsNil(o.StableEgressIP) {
		return nil, false
	}
	return o.StableEgressIP, true
}

// SetStableEgressIP gets a reference to the given bool and assigns it to the StableEgressIP field.
func (o *UpdateNetworkConfigRequest) SetStableEgressIP(v bool) {
	o.StableEgressIP = &v
}

// GetTlsTerminationPort returns the TlsTerminationPort field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetTlsTerminationPort() int64 {
	if o == nil || IsNil(o.TlsTerminationPort) {
		var ret int64
		return ret
	}
	return *o.TlsTerminationPort
}

// GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetTlsTerminationPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TlsTerminationPort) {
		return nil, false
	}
	return o.TlsTerminationPort, true
}

// SetTlsTerminationPort gets a reference to the given int64 and assigns it to the TlsTerminationPort field.
func (o *UpdateNetworkConfigRequest) SetTlsTerminationPort(v int64) {
	o.TlsTerminationPort = &v
}

// GetToken returns the Token field value
func (o *UpdateNetworkConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateNetworkConfigRequest) SetToken(v string) {
	o.Token = v
}

// GetZoneConfiguration returns the ZoneConfiguration field value if set, zero value otherwise.
func (o *UpdateNetworkConfigRequest) GetZoneConfiguration() string {
	if o == nil || IsNil(o.ZoneConfiguration) {
		var ret string
		return ret
	}
	return *o.ZoneConfiguration
}

// GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkConfigRequest) GetZoneConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneConfiguration) {
		return nil, false
	}
	return o.ZoneConfiguration, true
}

// SetZoneConfiguration gets a reference to the given string and assigns it to the ZoneConfiguration field.
func (o *UpdateNetworkConfigRequest) SetZoneConfiguration(v string) {
	o.ZoneConfiguration = &v
}

func (o UpdateNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EndpointPerReplica) {
		toSerialize["endpointPerReplica"] = o.EndpointPerReplica
	}
	toSerialize["id"] = o.Id
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
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.StableEgressIP) {
		toSerialize["stableEgressIP"] = o.StableEgressIP
	}
	if !IsNil(o.TlsTerminationPort) {
		toSerialize["tlsTerminationPort"] = o.TlsTerminationPort
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.ZoneConfiguration) {
		toSerialize["zoneConfiguration"] = o.ZoneConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworkConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serviceId",
		"token",
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

	varUpdateNetworkConfigRequest := _UpdateNetworkConfigRequest{}

	err = json.Unmarshal(data, &varUpdateNetworkConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdateNetworkConfigRequest(varUpdateNetworkConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "endpointPerReplica")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "name")
		delete(additionalProperties, "openPorts")
		delete(additionalProperties, "privateNetworkingConfiguration")
		delete(additionalProperties, "publicNetworkingConfiguration")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "stableEgressIP")
		delete(additionalProperties, "tlsTerminationPort")
		delete(additionalProperties, "token")
		delete(additionalProperties, "zoneConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworkConfigRequest struct {
	value *UpdateNetworkConfigRequest
	isSet bool
}

func (v NullableUpdateNetworkConfigRequest) Get() *UpdateNetworkConfigRequest {
	return v.value
}

func (v *NullableUpdateNetworkConfigRequest) Set(val *UpdateNetworkConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkConfigRequest(val *UpdateNetworkConfigRequest) *NullableUpdateNetworkConfigRequest {
	return &NullableUpdateNetworkConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


