# UpdateNetworkConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the network config | [optional] 
**EndpointPerReplica** | Pointer to **bool** | Generates a DNS endpoint per-replica for this network config | [optional] 
**Internal** | Pointer to **bool** | Restrict access to this network config to the internal network | [optional] 
**Name** | Pointer to **string** | Name of the network config | [optional] 
**NamedOpenPorts** | Pointer to [**map[string]NamedPortSpec**](NamedPortSpec.md) | Named ports to map to the generated DNS endpoint | [optional] 
**OpenPorts** | Pointer to **[]int64** | Ports to map to the generated DNS endpoint | [optional] 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**StableEgressIP** | Pointer to **bool** | Create an external node load balancer per node rather than expose the node ip directly | [optional] 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**ZoneConfiguration** | Pointer to **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | [optional] 

## Methods

### NewUpdateNetworkConfigRequest2

`func NewUpdateNetworkConfigRequest2() *UpdateNetworkConfigRequest2`

NewUpdateNetworkConfigRequest2 instantiates a new UpdateNetworkConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkConfigRequest2WithDefaults

`func NewUpdateNetworkConfigRequest2WithDefaults() *UpdateNetworkConfigRequest2`

NewUpdateNetworkConfigRequest2WithDefaults instantiates a new UpdateNetworkConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateNetworkConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointPerReplica

`func (o *UpdateNetworkConfigRequest2) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *UpdateNetworkConfigRequest2) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *UpdateNetworkConfigRequest2) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.

### HasEndpointPerReplica

`func (o *UpdateNetworkConfigRequest2) HasEndpointPerReplica() bool`

HasEndpointPerReplica returns a boolean if a field has been set.

### GetInternal

`func (o *UpdateNetworkConfigRequest2) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *UpdateNetworkConfigRequest2) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *UpdateNetworkConfigRequest2) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *UpdateNetworkConfigRequest2) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamedOpenPorts

`func (o *UpdateNetworkConfigRequest2) GetNamedOpenPorts() map[string]NamedPortSpec`

GetNamedOpenPorts returns the NamedOpenPorts field if non-nil, zero value otherwise.

### GetNamedOpenPortsOk

`func (o *UpdateNetworkConfigRequest2) GetNamedOpenPortsOk() (*map[string]NamedPortSpec, bool)`

GetNamedOpenPortsOk returns a tuple with the NamedOpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedOpenPorts

`func (o *UpdateNetworkConfigRequest2) SetNamedOpenPorts(v map[string]NamedPortSpec)`

SetNamedOpenPorts sets NamedOpenPorts field to given value.

### HasNamedOpenPorts

`func (o *UpdateNetworkConfigRequest2) HasNamedOpenPorts() bool`

HasNamedOpenPorts returns a boolean if a field has been set.

### GetOpenPorts

`func (o *UpdateNetworkConfigRequest2) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *UpdateNetworkConfigRequest2) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *UpdateNetworkConfigRequest2) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *UpdateNetworkConfigRequest2) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequest2) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequest2) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest2) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetStableEgressIP

`func (o *UpdateNetworkConfigRequest2) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *UpdateNetworkConfigRequest2) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *UpdateNetworkConfigRequest2) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.

### HasStableEgressIP

`func (o *UpdateNetworkConfigRequest2) HasStableEgressIP() bool`

HasStableEgressIP returns a boolean if a field has been set.

### GetTlsTerminationPort

`func (o *UpdateNetworkConfigRequest2) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *UpdateNetworkConfigRequest2) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *UpdateNetworkConfigRequest2) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *UpdateNetworkConfigRequest2) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetZoneConfiguration

`func (o *UpdateNetworkConfigRequest2) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *UpdateNetworkConfigRequest2) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *UpdateNetworkConfigRequest2) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.

### HasZoneConfiguration

`func (o *UpdateNetworkConfigRequest2) HasZoneConfiguration() bool`

HasZoneConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


