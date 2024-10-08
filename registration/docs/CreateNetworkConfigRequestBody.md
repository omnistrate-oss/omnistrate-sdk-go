# CreateNetworkConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the network config | 
**EndpointPerReplica** | **bool** | Generates a DNS endpoint per-replica for this network config | 
**Internal** | Pointer to **bool** | Restrict access to this network config to the internal network | [optional] [default to false]
**Name** | **string** | Name of the network config | 
**OpenPorts** | Pointer to **[]int64** | Ports to map to the generated DNS endpoint | [optional] 
**PrivateNetworkingConfiguration** | Pointer to [**PrivateNetworkingConfiguration**](PrivateNetworkingConfiguration.md) |  | [optional] 
**PublicNetworkingConfiguration** | Pointer to [**PublicNetworkingConfiguration**](PublicNetworkingConfiguration.md) |  | [optional] 
**StableEgressIP** | Pointer to **bool** | Enable stable egress IP | [optional] 
**TlsTerminationPort** | Pointer to **int64** | The port that hosts the reverse proxy for TLS termination | [optional] 
**ZoneConfiguration** | Pointer to **string** | The preferred type of zonal availability for this resource and the specific zone(s) to deploy in | [optional] 

## Methods

### NewCreateNetworkConfigRequestBody

`func NewCreateNetworkConfigRequestBody(description string, endpointPerReplica bool, name string, ) *CreateNetworkConfigRequestBody`

NewCreateNetworkConfigRequestBody instantiates a new CreateNetworkConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkConfigRequestBodyWithDefaults

`func NewCreateNetworkConfigRequestBodyWithDefaults() *CreateNetworkConfigRequestBody`

NewCreateNetworkConfigRequestBodyWithDefaults instantiates a new CreateNetworkConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateNetworkConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndpointPerReplica

`func (o *CreateNetworkConfigRequestBody) GetEndpointPerReplica() bool`

GetEndpointPerReplica returns the EndpointPerReplica field if non-nil, zero value otherwise.

### GetEndpointPerReplicaOk

`func (o *CreateNetworkConfigRequestBody) GetEndpointPerReplicaOk() (*bool, bool)`

GetEndpointPerReplicaOk returns a tuple with the EndpointPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPerReplica

`func (o *CreateNetworkConfigRequestBody) SetEndpointPerReplica(v bool)`

SetEndpointPerReplica sets EndpointPerReplica field to given value.


### GetInternal

`func (o *CreateNetworkConfigRequestBody) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateNetworkConfigRequestBody) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateNetworkConfigRequestBody) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateNetworkConfigRequestBody) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetOpenPorts

`func (o *CreateNetworkConfigRequestBody) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *CreateNetworkConfigRequestBody) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *CreateNetworkConfigRequestBody) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *CreateNetworkConfigRequestBody) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) GetPrivateNetworkingConfiguration() PrivateNetworkingConfiguration`

GetPrivateNetworkingConfiguration returns the PrivateNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkingConfigurationOk

`func (o *CreateNetworkConfigRequestBody) GetPrivateNetworkingConfigurationOk() (*PrivateNetworkingConfiguration, bool)`

GetPrivateNetworkingConfigurationOk returns a tuple with the PrivateNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) SetPrivateNetworkingConfiguration(v PrivateNetworkingConfiguration)`

SetPrivateNetworkingConfiguration sets PrivateNetworkingConfiguration field to given value.

### HasPrivateNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) HasPrivateNetworkingConfiguration() bool`

HasPrivateNetworkingConfiguration returns a boolean if a field has been set.

### GetPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) GetPublicNetworkingConfiguration() PublicNetworkingConfiguration`

GetPublicNetworkingConfiguration returns the PublicNetworkingConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkingConfigurationOk

`func (o *CreateNetworkConfigRequestBody) GetPublicNetworkingConfigurationOk() (*PublicNetworkingConfiguration, bool)`

GetPublicNetworkingConfigurationOk returns a tuple with the PublicNetworkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) SetPublicNetworkingConfiguration(v PublicNetworkingConfiguration)`

SetPublicNetworkingConfiguration sets PublicNetworkingConfiguration field to given value.

### HasPublicNetworkingConfiguration

`func (o *CreateNetworkConfigRequestBody) HasPublicNetworkingConfiguration() bool`

HasPublicNetworkingConfiguration returns a boolean if a field has been set.

### GetStableEgressIP

`func (o *CreateNetworkConfigRequestBody) GetStableEgressIP() bool`

GetStableEgressIP returns the StableEgressIP field if non-nil, zero value otherwise.

### GetStableEgressIPOk

`func (o *CreateNetworkConfigRequestBody) GetStableEgressIPOk() (*bool, bool)`

GetStableEgressIPOk returns a tuple with the StableEgressIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableEgressIP

`func (o *CreateNetworkConfigRequestBody) SetStableEgressIP(v bool)`

SetStableEgressIP sets StableEgressIP field to given value.

### HasStableEgressIP

`func (o *CreateNetworkConfigRequestBody) HasStableEgressIP() bool`

HasStableEgressIP returns a boolean if a field has been set.

### GetTlsTerminationPort

`func (o *CreateNetworkConfigRequestBody) GetTlsTerminationPort() int64`

GetTlsTerminationPort returns the TlsTerminationPort field if non-nil, zero value otherwise.

### GetTlsTerminationPortOk

`func (o *CreateNetworkConfigRequestBody) GetTlsTerminationPortOk() (*int64, bool)`

GetTlsTerminationPortOk returns a tuple with the TlsTerminationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationPort

`func (o *CreateNetworkConfigRequestBody) SetTlsTerminationPort(v int64)`

SetTlsTerminationPort sets TlsTerminationPort field to given value.

### HasTlsTerminationPort

`func (o *CreateNetworkConfigRequestBody) HasTlsTerminationPort() bool`

HasTlsTerminationPort returns a boolean if a field has been set.

### GetZoneConfiguration

`func (o *CreateNetworkConfigRequestBody) GetZoneConfiguration() string`

GetZoneConfiguration returns the ZoneConfiguration field if non-nil, zero value otherwise.

### GetZoneConfigurationOk

`func (o *CreateNetworkConfigRequestBody) GetZoneConfigurationOk() (*string, bool)`

GetZoneConfigurationOk returns a tuple with the ZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneConfiguration

`func (o *CreateNetworkConfigRequestBody) SetZoneConfiguration(v string)`

SetZoneConfiguration sets ZoneConfiguration field to given value.

### HasZoneConfiguration

`func (o *CreateNetworkConfigRequestBody) HasZoneConfiguration() bool`

HasZoneConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


