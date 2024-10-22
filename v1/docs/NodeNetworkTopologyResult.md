# NodeNetworkTopologyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | **string** | The availability zone of the node | 
**DetailedHealth** | [**DetailedNodeHealthResult**](DetailedNodeHealthResult.md) |  | 
**Endpoint** | **string** | The endpoint of the node | 
**HealthStatus** | **string** | The heath status of the node | 
**Id** | **string** | The ID of the node | 
**Ports** | Pointer to **[]int64** | The ports that this node exposes | [optional] 
**Status** | **string** | The status of the node | 
**StorageSize** | Pointer to **int64** | The storage size of the node in GiB | [optional] 

## Methods

### NewNodeNetworkTopologyResult

`func NewNodeNetworkTopologyResult(availabilityZone string, detailedHealth DetailedNodeHealthResult, endpoint string, healthStatus string, id string, status string, ) *NodeNetworkTopologyResult`

NewNodeNetworkTopologyResult instantiates a new NodeNetworkTopologyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeNetworkTopologyResultWithDefaults

`func NewNodeNetworkTopologyResultWithDefaults() *NodeNetworkTopologyResult`

NewNodeNetworkTopologyResultWithDefaults instantiates a new NodeNetworkTopologyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *NodeNetworkTopologyResult) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *NodeNetworkTopologyResult) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *NodeNetworkTopologyResult) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetDetailedHealth

`func (o *NodeNetworkTopologyResult) GetDetailedHealth() DetailedNodeHealthResult`

GetDetailedHealth returns the DetailedHealth field if non-nil, zero value otherwise.

### GetDetailedHealthOk

`func (o *NodeNetworkTopologyResult) GetDetailedHealthOk() (*DetailedNodeHealthResult, bool)`

GetDetailedHealthOk returns a tuple with the DetailedHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedHealth

`func (o *NodeNetworkTopologyResult) SetDetailedHealth(v DetailedNodeHealthResult)`

SetDetailedHealth sets DetailedHealth field to given value.


### GetEndpoint

`func (o *NodeNetworkTopologyResult) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NodeNetworkTopologyResult) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NodeNetworkTopologyResult) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetHealthStatus

`func (o *NodeNetworkTopologyResult) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *NodeNetworkTopologyResult) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *NodeNetworkTopologyResult) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.


### GetId

`func (o *NodeNetworkTopologyResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeNetworkTopologyResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeNetworkTopologyResult) SetId(v string)`

SetId sets Id field to given value.


### GetPorts

`func (o *NodeNetworkTopologyResult) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NodeNetworkTopologyResult) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NodeNetworkTopologyResult) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NodeNetworkTopologyResult) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatus

`func (o *NodeNetworkTopologyResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeNetworkTopologyResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeNetworkTopologyResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageSize

`func (o *NodeNetworkTopologyResult) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *NodeNetworkTopologyResult) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *NodeNetworkTopologyResult) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *NodeNetworkTopologyResult) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


