# UpdateNetworkConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the network config | [optional] 
**EndpointPerReplica** | Pointer to **bool** | Generates a DNS endpoint per-replica for this network config | [optional] 
**Internal** | Pointer to **bool** | Restrict access to this network config to the internal network | [optional] 
**Name** | Pointer to **string** | Name of the network config | [optional] 
**OpenPorts** | Pointer to **[]int64** | Ports to map to the generated DNS endpoint | [optional] 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**StableEgressIP** | Pointer to **bool** | Create an external node load balancer per node rather than expose the node ip directly | [optional] 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**ZoneConfiguration** | Pointer to **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | [optional] 

## Methods

### NewUpdateNetworkConfigRequestBody

`func NewUpdateNetworkConfigRequestBody() *UpdateNetworkConfigRequestBody`

NewUpdateNetworkConfigRequestBody instantiates a new UpdateNetworkConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkConfigRequestBodyWithDefaults

`func NewUpdateNetworkConfigRequestBodyWithDefaults() *UpdateNetworkConfigRequestBody`

NewUpdateNetworkConfigRequestBodyWithDefaults instantiates a new UpdateNetworkConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateNetworkConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkConfigRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointPerReplica

`func (o *UpdateNetworkConfigRequestBody) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *UpdateNetworkConfigRequestBody) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *UpdateNetworkConfigRequestBody) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.

### HasEndpointPerReplica

`func (o *UpdateNetworkConfigRequestBody) HasEndpointPerReplica() bool`

HasEndpointPerReplica returns a boolean if a field has been set.

### GetInternal

`func (o *UpdateNetworkConfigRequestBody) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *UpdateNetworkConfigRequestBody) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *UpdateNetworkConfigRequestBody) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *UpdateNetworkConfigRequestBody) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkConfigRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenPorts

`func (o *UpdateNetworkConfigRequestBody) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *UpdateNetworkConfigRequestBody) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *UpdateNetworkConfigRequestBody) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *UpdateNetworkConfigRequestBody) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequestBody) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *UpdateNetworkConfigRequestBody) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *UpdateNetworkConfigRequestBody) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetStableEgressIP

`func (o *UpdateNetworkConfigRequestBody) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *UpdateNetworkConfigRequestBody) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *UpdateNetworkConfigRequestBody) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.

### HasStableEgressIP

`func (o *UpdateNetworkConfigRequestBody) HasStableEgressIP() bool`

HasStableEgressIP returns a boolean if a field has been set.

### GetTlsTerminationPort

`func (o *UpdateNetworkConfigRequestBody) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *UpdateNetworkConfigRequestBody) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *UpdateNetworkConfigRequestBody) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *UpdateNetworkConfigRequestBody) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetZoneConfiguration

`func (o *UpdateNetworkConfigRequestBody) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *UpdateNetworkConfigRequestBody) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *UpdateNetworkConfigRequestBody) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.

### HasZoneConfiguration

`func (o *UpdateNetworkConfigRequestBody) HasZoneConfiguration() bool`

HasZoneConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


