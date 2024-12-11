# NodeHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** | The availability zone of the node | [optional] 
**ConnectivityStatus** | Pointer to **string** | The health status of the network endpoints | [optional] 
**DiskHealth** | Pointer to **string** | The health status of the disk | [optional] 
**Endpoint** | Pointer to **string** | The endpoint of the node | [optional] 
**Events** | Pointer to **[]string** | The list of process events | [optional] 
**IntegrationsHealth** | Pointer to [**IntegrationsHealth**](IntegrationsHealth.md) |  | [optional] 
**LoadHealth** | Pointer to **string** | (deprecated) The load status of the pod | [optional] 
**LoadStatus** | Pointer to **string** | The load status of the pod | [optional] 
**NodeHealth** | Pointer to **string** | The health status of the machine hosting the service | [optional] 
**NodeName** | Pointer to **string** | The name of the node | [optional] 
**Ports** | Pointer to **[]int64** | The ports that this node exposes | [optional] 
**ProcessHealth** | Pointer to **string** | The health status of the process | [optional] 
**ProcessLiveness** | Pointer to **string** | The liveness status of the process | [optional] 
**RecentLogs** | Pointer to **string** | The recent logs of the process | [optional] 
**Status** | Pointer to **string** | The overall status of the node | [optional] 

## Methods

### NewNodeHealthSummary

`func NewNodeHealthSummary() *NodeHealthSummary`

NewNodeHealthSummary instantiates a new NodeHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHealthSummaryWithDefaults

`func NewNodeHealthSummaryWithDefaults() *NodeHealthSummary`

NewNodeHealthSummaryWithDefaults instantiates a new NodeHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *NodeHealthSummary) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *NodeHealthSummary) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *NodeHealthSummary) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *NodeHealthSummary) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetConnectivityStatus

`func (o *NodeHealthSummary) GetConnectivityStatus() string`

GetConnectivityStatus returns the ConnectivityStatus field if non-nil, zero value otherwise.

### GetConnectivityStatusOk

`func (o *NodeHealthSummary) GetConnectivityStatusOk() (*string, bool)`

GetConnectivityStatusOk returns a tuple with the ConnectivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityStatus

`func (o *NodeHealthSummary) SetConnectivityStatus(v string)`

SetConnectivityStatus sets ConnectivityStatus field to given value.

### HasConnectivityStatus

`func (o *NodeHealthSummary) HasConnectivityStatus() bool`

HasConnectivityStatus returns a boolean if a field has been set.

### GetDiskHealth

`func (o *NodeHealthSummary) GetDiskHealth() string`

GetDiskHealth returns the DiskHealth field if non-nil, zero value otherwise.

### GetDiskHealthOk

`func (o *NodeHealthSummary) GetDiskHealthOk() (*string, bool)`

GetDiskHealthOk returns a tuple with the DiskHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskHealth

`func (o *NodeHealthSummary) SetDiskHealth(v string)`

SetDiskHealth sets DiskHealth field to given value.

### HasDiskHealth

`func (o *NodeHealthSummary) HasDiskHealth() bool`

HasDiskHealth returns a boolean if a field has been set.

### GetEndpoint

`func (o *NodeHealthSummary) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NodeHealthSummary) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NodeHealthSummary) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *NodeHealthSummary) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEvents

`func (o *NodeHealthSummary) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NodeHealthSummary) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NodeHealthSummary) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NodeHealthSummary) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIntegrationsHealth

`func (o *NodeHealthSummary) GetIntegrationsHealth() IntegrationsHealth`

GetIntegrationsHealth returns the IntegrationsHealth field if non-nil, zero value otherwise.

### GetIntegrationsHealthOk

`func (o *NodeHealthSummary) GetIntegrationsHealthOk() (*IntegrationsHealth, bool)`

GetIntegrationsHealthOk returns a tuple with the IntegrationsHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsHealth

`func (o *NodeHealthSummary) SetIntegrationsHealth(v IntegrationsHealth)`

SetIntegrationsHealth sets IntegrationsHealth field to given value.

### HasIntegrationsHealth

`func (o *NodeHealthSummary) HasIntegrationsHealth() bool`

HasIntegrationsHealth returns a boolean if a field has been set.

### GetLoadHealth

`func (o *NodeHealthSummary) GetLoadHealth() string`

GetLoadHealth returns the LoadHealth field if non-nil, zero value otherwise.

### GetLoadHealthOk

`func (o *NodeHealthSummary) GetLoadHealthOk() (*string, bool)`

GetLoadHealthOk returns a tuple with the LoadHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadHealth

`func (o *NodeHealthSummary) SetLoadHealth(v string)`

SetLoadHealth sets LoadHealth field to given value.

### HasLoadHealth

`func (o *NodeHealthSummary) HasLoadHealth() bool`

HasLoadHealth returns a boolean if a field has been set.

### GetLoadStatus

`func (o *NodeHealthSummary) GetLoadStatus() string`

GetLoadStatus returns the LoadStatus field if non-nil, zero value otherwise.

### GetLoadStatusOk

`func (o *NodeHealthSummary) GetLoadStatusOk() (*string, bool)`

GetLoadStatusOk returns a tuple with the LoadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadStatus

`func (o *NodeHealthSummary) SetLoadStatus(v string)`

SetLoadStatus sets LoadStatus field to given value.

### HasLoadStatus

`func (o *NodeHealthSummary) HasLoadStatus() bool`

HasLoadStatus returns a boolean if a field has been set.

### GetNodeHealth

`func (o *NodeHealthSummary) GetNodeHealth() string`

GetNodeHealth returns the NodeHealth field if non-nil, zero value otherwise.

### GetNodeHealthOk

`func (o *NodeHealthSummary) GetNodeHealthOk() (*string, bool)`

GetNodeHealthOk returns a tuple with the NodeHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHealth

`func (o *NodeHealthSummary) SetNodeHealth(v string)`

SetNodeHealth sets NodeHealth field to given value.

### HasNodeHealth

`func (o *NodeHealthSummary) HasNodeHealth() bool`

HasNodeHealth returns a boolean if a field has been set.

### GetNodeName

`func (o *NodeHealthSummary) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeHealthSummary) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeHealthSummary) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *NodeHealthSummary) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPorts

`func (o *NodeHealthSummary) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NodeHealthSummary) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NodeHealthSummary) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NodeHealthSummary) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcessHealth

`func (o *NodeHealthSummary) GetProcessHealth() string`

GetProcessHealth returns the ProcessHealth field if non-nil, zero value otherwise.

### GetProcessHealthOk

`func (o *NodeHealthSummary) GetProcessHealthOk() (*string, bool)`

GetProcessHealthOk returns a tuple with the ProcessHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessHealth

`func (o *NodeHealthSummary) SetProcessHealth(v string)`

SetProcessHealth sets ProcessHealth field to given value.

### HasProcessHealth

`func (o *NodeHealthSummary) HasProcessHealth() bool`

HasProcessHealth returns a boolean if a field has been set.

### GetProcessLiveness

`func (o *NodeHealthSummary) GetProcessLiveness() string`

GetProcessLiveness returns the ProcessLiveness field if non-nil, zero value otherwise.

### GetProcessLivenessOk

`func (o *NodeHealthSummary) GetProcessLivenessOk() (*string, bool)`

GetProcessLivenessOk returns a tuple with the ProcessLiveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessLiveness

`func (o *NodeHealthSummary) SetProcessLiveness(v string)`

SetProcessLiveness sets ProcessLiveness field to given value.

### HasProcessLiveness

`func (o *NodeHealthSummary) HasProcessLiveness() bool`

HasProcessLiveness returns a boolean if a field has been set.

### GetRecentLogs

`func (o *NodeHealthSummary) GetRecentLogs() string`

GetRecentLogs returns the RecentLogs field if non-nil, zero value otherwise.

### GetRecentLogsOk

`func (o *NodeHealthSummary) GetRecentLogsOk() (*string, bool)`

GetRecentLogsOk returns a tuple with the RecentLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentLogs

`func (o *NodeHealthSummary) SetRecentLogs(v string)`

SetRecentLogs sets RecentLogs field to given value.

### HasRecentLogs

`func (o *NodeHealthSummary) HasRecentLogs() bool`

HasRecentLogs returns a boolean if a field has been set.

### GetStatus

`func (o *NodeHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeHealthSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


