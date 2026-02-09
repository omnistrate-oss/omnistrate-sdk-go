# DescribeNetworkConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the network config | 
**EndpointPerReplica** | **bool** | Generates a DNS endpoint per-replica for this network config | 
**Id** | **string** | ID of a Network Config | 
**InfraConfigIDs** | Pointer to **[]string** | The list of infra config IDs associated with the compute config. | [optional] 
**Internal** | **bool** | Restrict access to this network config to the internal network | 
**Name** | **string** | Name of the network config | 
**NamedOpenPorts** | Pointer to [**map[string]NamedPortSpec**](NamedPortSpec.md) | Map of port names to port specifications | [optional] 
**OpenPorts** | **[]int64** | Ports to map to the generated DNS endpoint | 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StableEgressIP** | **bool** | Create an external node load balancer per node rather than expose the node ip directly | 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**ZoneConfiguration** | **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | 

## Methods

### NewDescribeNetworkConfigResult

`func NewDescribeNetworkConfigResult(description string, endpointPerReplica bool, id string, internal bool, name string, openPorts []int64, serviceId string, stableEgressIP bool, zoneConfiguration string, ) *DescribeNetworkConfigResult`

NewDescribeNetworkConfigResult instantiates a new DescribeNetworkConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNetworkConfigResultWithDefaults

`func NewDescribeNetworkConfigResultWithDefaults() *DescribeNetworkConfigResult`

NewDescribeNetworkConfigResultWithDefaults instantiates a new DescribeNetworkConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeNetworkConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeNetworkConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeNetworkConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndpointPerReplica

`func (o *DescribeNetworkConfigResult) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *DescribeNetworkConfigResult) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *DescribeNetworkConfigResult) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.


### GetId

`func (o *DescribeNetworkConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeNetworkConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeNetworkConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetInfraConfigIDs

`func (o *DescribeNetworkConfigResult) GetInfraConfigIDs() []string`

GetInfraConfigIDs returns the InfraConfigIDs field if non-nil, zero value otherwise.

### GetInfraConfigIDsOk

`func (o *DescribeNetworkConfigResult) GetInfraConfigIDsOk() (*[]string, bool)`

GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigIDs

`func (o *DescribeNetworkConfigResult) SetInfraConfigIDs(v []string)`

SetInfraConfigIDs sets InfraConfigIDs field to given value.

### HasInfraConfigIDs

`func (o *DescribeNetworkConfigResult) HasInfraConfigIDs() bool`

HasInfraConfigIDs returns a boolean if a field has been set.

### GetInternal

`func (o *DescribeNetworkConfigResult) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *DescribeNetworkConfigResult) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *DescribeNetworkConfigResult) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetName

`func (o *DescribeNetworkConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeNetworkConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeNetworkConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetNamedOpenPorts

`func (o *DescribeNetworkConfigResult) GetNamedOpenPorts() map[string]NamedPortSpec`

GetNamedOpenPorts returns the NamedOpenPorts field if non-nil, zero value otherwise.

### GetNamedOpenPortsOk

`func (o *DescribeNetworkConfigResult) GetNamedOpenPortsOk() (*map[string]NamedPortSpec, bool)`

GetNamedOpenPortsOk returns a tuple with the NamedOpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedOpenPorts

`func (o *DescribeNetworkConfigResult) SetNamedOpenPorts(v map[string]NamedPortSpec)`

SetNamedOpenPorts sets NamedOpenPorts field to given value.

### HasNamedOpenPorts

`func (o *DescribeNetworkConfigResult) HasNamedOpenPorts() bool`

HasNamedOpenPorts returns a boolean if a field has been set.

### GetOpenPorts

`func (o *DescribeNetworkConfigResult) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *DescribeNetworkConfigResult) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *DescribeNetworkConfigResult) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.


### GetPrivateNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *DescribeNetworkConfigResult) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *DescribeNetworkConfigResult) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *DescribeNetworkConfigResult) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeNetworkConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeNetworkConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeNetworkConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStableEgressIP

`func (o *DescribeNetworkConfigResult) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *DescribeNetworkConfigResult) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *DescribeNetworkConfigResult) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.


### GetTlsTerminationPort

`func (o *DescribeNetworkConfigResult) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *DescribeNetworkConfigResult) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *DescribeNetworkConfigResult) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *DescribeNetworkConfigResult) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetZoneConfiguration

`func (o *DescribeNetworkConfigResult) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *DescribeNetworkConfigResult) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *DescribeNetworkConfigResult) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


