# FleetCreateInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**TargetRegion** | Pointer to **string** | The target region to create the snapshot in. If not specified, use the same region as the instance | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateInstanceSnapshotRequest

`func NewFleetCreateInstanceSnapshotRequest(environmentId string, instanceId string, serviceId string, token string, ) *FleetCreateInstanceSnapshotRequest`

NewFleetCreateInstanceSnapshotRequest instantiates a new FleetCreateInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateInstanceSnapshotRequestWithDefaults

`func NewFleetCreateInstanceSnapshotRequestWithDefaults() *FleetCreateInstanceSnapshotRequest`

NewFleetCreateInstanceSnapshotRequestWithDefaults instantiates a new FleetCreateInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetCreateInstanceSnapshotRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetCreateInstanceSnapshotRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetCreateInstanceSnapshotRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetCreateInstanceSnapshotRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetCreateInstanceSnapshotRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetCreateInstanceSnapshotRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetCreateInstanceSnapshotRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetCreateInstanceSnapshotRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetCreateInstanceSnapshotRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetRegion

`func (o *FleetCreateInstanceSnapshotRequest) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *FleetCreateInstanceSnapshotRequest) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *FleetCreateInstanceSnapshotRequest) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.

### HasTargetRegion

`func (o *FleetCreateInstanceSnapshotRequest) HasTargetRegion() bool`

HasTargetRegion returns a boolean if a field has been set.

### GetToken

`func (o *FleetCreateInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


