# FleetListLinkedInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExcludeHAStatus** | Pointer to **bool** | Whether to exclude high availability and autoscaling status from the response | [optional] 
**ExcludeIntegrations** | Pointer to **bool** | Whether to exclude integration statuses (e.g. OpenTelemetry) from the response | [optional] 
**ExcludeNetworkTopology** | Pointer to **bool** | Whether to exclude network topology details from the response | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListLinkedInstancesRequest

`func NewFleetListLinkedInstancesRequest(environmentId string, instanceId string, serviceId string, token string, ) *FleetListLinkedInstancesRequest`

NewFleetListLinkedInstancesRequest instantiates a new FleetListLinkedInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListLinkedInstancesRequestWithDefaults

`func NewFleetListLinkedInstancesRequestWithDefaults() *FleetListLinkedInstancesRequest`

NewFleetListLinkedInstancesRequestWithDefaults instantiates a new FleetListLinkedInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListLinkedInstancesRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListLinkedInstancesRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListLinkedInstancesRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExcludeHAStatus

`func (o *FleetListLinkedInstancesRequest) GetExcludeHAStatus() bool`

GetExcludeHAStatus returns the ExcludeHAStatus field if non-nil, zero value otherwise.

### GetExcludeHAStatusOk

`func (o *FleetListLinkedInstancesRequest) GetExcludeHAStatusOk() (*bool, bool)`

GetExcludeHAStatusOk returns a tuple with the ExcludeHAStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHAStatus

`func (o *FleetListLinkedInstancesRequest) SetExcludeHAStatus(v bool)`

SetExcludeHAStatus sets ExcludeHAStatus field to given value.

### HasExcludeHAStatus

`func (o *FleetListLinkedInstancesRequest) HasExcludeHAStatus() bool`

HasExcludeHAStatus returns a boolean if a field has been set.

### GetExcludeIntegrations

`func (o *FleetListLinkedInstancesRequest) GetExcludeIntegrations() bool`

GetExcludeIntegrations returns the ExcludeIntegrations field if non-nil, zero value otherwise.

### GetExcludeIntegrationsOk

`func (o *FleetListLinkedInstancesRequest) GetExcludeIntegrationsOk() (*bool, bool)`

GetExcludeIntegrationsOk returns a tuple with the ExcludeIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIntegrations

`func (o *FleetListLinkedInstancesRequest) SetExcludeIntegrations(v bool)`

SetExcludeIntegrations sets ExcludeIntegrations field to given value.

### HasExcludeIntegrations

`func (o *FleetListLinkedInstancesRequest) HasExcludeIntegrations() bool`

HasExcludeIntegrations returns a boolean if a field has been set.

### GetExcludeNetworkTopology

`func (o *FleetListLinkedInstancesRequest) GetExcludeNetworkTopology() bool`

GetExcludeNetworkTopology returns the ExcludeNetworkTopology field if non-nil, zero value otherwise.

### GetExcludeNetworkTopologyOk

`func (o *FleetListLinkedInstancesRequest) GetExcludeNetworkTopologyOk() (*bool, bool)`

GetExcludeNetworkTopologyOk returns a tuple with the ExcludeNetworkTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeNetworkTopology

`func (o *FleetListLinkedInstancesRequest) SetExcludeNetworkTopology(v bool)`

SetExcludeNetworkTopology sets ExcludeNetworkTopology field to given value.

### HasExcludeNetworkTopology

`func (o *FleetListLinkedInstancesRequest) HasExcludeNetworkTopology() bool`

HasExcludeNetworkTopology returns a boolean if a field has been set.

### GetInstanceId

`func (o *FleetListLinkedInstancesRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetListLinkedInstancesRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetListLinkedInstancesRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetListLinkedInstancesRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListLinkedInstancesRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListLinkedInstancesRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetListLinkedInstancesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListLinkedInstancesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListLinkedInstancesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


