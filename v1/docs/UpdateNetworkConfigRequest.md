# UpdateNetworkConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the network config | [optional] 
**EndpointPerReplica** | Pointer to **bool** | Generates a DNS endpoint per-replica for this network config | [optional] 
**Id** | **string** | ID of a Network Config | 
**Internal** | Pointer to **bool** | Restrict access to this network config to the internal network | [optional] 
**Name** | Pointer to **string** | Name of the network config | [optional] 
**NamedOpenPorts** | Pointer to [**map[string]NamedPortSpec**](NamedPortSpec.md) | Map of port names to port specifications | [optional] 
**OpenPorts** | Pointer to **[]int64** | Ports to map to the generated DNS endpoint | [optional] 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StableEgressIP** | Pointer to **bool** | Create an external node load balancer per node rather than expose the node ip directly | [optional] 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**ZoneConfiguration** | Pointer to **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | [optional] 

## Methods

### NewUpdateNetworkConfigRequest

`func NewUpdateNetworkConfigRequest(id string, serviceId string, token string, ) *UpdateNetworkConfigRequest`

NewUpdateNetworkConfigRequest instantiates a new UpdateNetworkConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkConfigRequestWithDefaults

`func NewUpdateNetworkConfigRequestWithDefaults() *UpdateNetworkConfigRequest`

NewUpdateNetworkConfigRequestWithDefaults instantiates a new UpdateNetworkConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateNetworkConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointPerReplica

`func (o *UpdateNetworkConfigRequest) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *UpdateNetworkConfigRequest) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *UpdateNetworkConfigRequest) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.

### HasEndpointPerReplica

`func (o *UpdateNetworkConfigRequest) HasEndpointPerReplica() bool`

HasEndpointPerReplica returns a boolean if a field has been set.

### GetId

`func (o *UpdateNetworkConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInternal

`func (o *UpdateNetworkConfigRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *UpdateNetworkConfigRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *UpdateNetworkConfigRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *UpdateNetworkConfigRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamedOpenPorts

`func (o *UpdateNetworkConfigRequest) GetNamedOpenPorts() map[string]NamedPortSpec`

GetNamedOpenPorts returns the NamedOpenPorts field if non-nil, zero value otherwise.

### GetNamedOpenPortsOk

`func (o *UpdateNetworkConfigRequest) GetNamedOpenPortsOk() (*map[string]NamedPortSpec, bool)`

GetNamedOpenPortsOk returns a tuple with the NamedOpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedOpenPorts

`func (o *UpdateNetworkConfigRequest) SetNamedOpenPorts(v map[string]NamedPortSpec)`

SetNamedOpenPorts sets NamedOpenPorts field to given value.

### HasNamedOpenPorts

`func (o *UpdateNetworkConfigRequest) HasNamedOpenPorts() bool`

HasNamedOpenPorts returns a boolean if a field has been set.

### GetOpenPorts

`func (o *UpdateNetworkConfigRequest) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *UpdateNetworkConfigRequest) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *UpdateNetworkConfigRequest) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *UpdateNetworkConfigRequest) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequest) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequest) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequest) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateNetworkConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateNetworkConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateNetworkConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStableEgressIP

`func (o *UpdateNetworkConfigRequest) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *UpdateNetworkConfigRequest) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *UpdateNetworkConfigRequest) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.

### HasStableEgressIP

`func (o *UpdateNetworkConfigRequest) HasStableEgressIP() bool`

HasStableEgressIP returns a boolean if a field has been set.

### GetTlsTerminationPort

`func (o *UpdateNetworkConfigRequest) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *UpdateNetworkConfigRequest) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *UpdateNetworkConfigRequest) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *UpdateNetworkConfigRequest) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetToken

`func (o *UpdateNetworkConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateNetworkConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateNetworkConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetZoneConfiguration

`func (o *UpdateNetworkConfigRequest) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *UpdateNetworkConfigRequest) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *UpdateNetworkConfigRequest) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.

### HasZoneConfiguration

`func (o *UpdateNetworkConfigRequest) HasZoneConfiguration() bool`

HasZoneConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


