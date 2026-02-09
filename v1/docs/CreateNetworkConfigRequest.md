# CreateNetworkConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the network config | 
**EndpointPerReplica** | **bool** | Generates a DNS endpoint per-replica for this network config | 
**Internal** | Pointer to **bool** | Restrict access to this network config to the internal network | [optional] [default to false]
**Name** | **string** | Name of the network config | 
**NamedOpenPorts** | Pointer to [**map[string]NamedPortSpec**](NamedPortSpec.md) | Map of port names to port specifications | [optional] 
**OpenPorts** | Pointer to **[]int64** | Ports to map to the generated DNS endpoint | [optional] 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StableEgressIP** | Pointer to **bool** | Enable stable egress IP | [optional] 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**ZoneConfiguration** | Pointer to **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | [optional] 

## Methods

### NewCreateNetworkConfigRequest

`func NewCreateNetworkConfigRequest(description string, endpointPerReplica bool, name string, serviceId string, token string, ) *CreateNetworkConfigRequest`

NewCreateNetworkConfigRequest instantiates a new CreateNetworkConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkConfigRequestWithDefaults

`func NewCreateNetworkConfigRequestWithDefaults() *CreateNetworkConfigRequest`

NewCreateNetworkConfigRequestWithDefaults instantiates a new CreateNetworkConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateNetworkConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndpointPerReplica

`func (o *CreateNetworkConfigRequest) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *CreateNetworkConfigRequest) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *CreateNetworkConfigRequest) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.


### GetInternal

`func (o *CreateNetworkConfigRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateNetworkConfigRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateNetworkConfigRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateNetworkConfigRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNamedOpenPorts

`func (o *CreateNetworkConfigRequest) GetNamedOpenPorts() map[string]NamedPortSpec`

GetNamedOpenPorts returns the NamedOpenPorts field if non-nil, zero value otherwise.

### GetNamedOpenPortsOk

`func (o *CreateNetworkConfigRequest) GetNamedOpenPortsOk() (*map[string]NamedPortSpec, bool)`

GetNamedOpenPortsOk returns a tuple with the NamedOpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedOpenPorts

`func (o *CreateNetworkConfigRequest) SetNamedOpenPorts(v map[string]NamedPortSpec)`

SetNamedOpenPorts sets NamedOpenPorts field to given value.

### HasNamedOpenPorts

`func (o *CreateNetworkConfigRequest) HasNamedOpenPorts() bool`

HasNamedOpenPorts returns a boolean if a field has been set.

### GetOpenPorts

`func (o *CreateNetworkConfigRequest) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *CreateNetworkConfigRequest) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *CreateNetworkConfigRequest) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *CreateNetworkConfigRequest) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *CreateNetworkConfigRequest) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *CreateNetworkConfigRequest) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequest) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateNetworkConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateNetworkConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateNetworkConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStableEgressIP

`func (o *CreateNetworkConfigRequest) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *CreateNetworkConfigRequest) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *CreateNetworkConfigRequest) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.

### HasStableEgressIP

`func (o *CreateNetworkConfigRequest) HasStableEgressIP() bool`

HasStableEgressIP returns a boolean if a field has been set.

### GetTlsTerminationPort

`func (o *CreateNetworkConfigRequest) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *CreateNetworkConfigRequest) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *CreateNetworkConfigRequest) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *CreateNetworkConfigRequest) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetToken

`func (o *CreateNetworkConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateNetworkConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateNetworkConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetZoneConfiguration

`func (o *CreateNetworkConfigRequest) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *CreateNetworkConfigRequest) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *CreateNetworkConfigRequest) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.

### HasZoneConfiguration

`func (o *CreateNetworkConfigRequest) HasZoneConfiguration() bool`

HasZoneConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


