# FleetCopyResourceInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**TargetRegion** | **string** | The target region to copy the snapshot to | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCopyResourceInstanceSnapshotRequest

`func NewFleetCopyResourceInstanceSnapshotRequest(environmentId string, instanceId string, serviceId string, targetRegion string, token string, ) *FleetCopyResourceInstanceSnapshotRequest`

NewFleetCopyResourceInstanceSnapshotRequest instantiates a new FleetCopyResourceInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCopyResourceInstanceSnapshotRequestWithDefaults

`func NewFleetCopyResourceInstanceSnapshotRequestWithDefaults() *FleetCopyResourceInstanceSnapshotRequest`

NewFleetCopyResourceInstanceSnapshotRequestWithDefaults instantiates a new FleetCopyResourceInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetCopyResourceInstanceSnapshotRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetCopyResourceInstanceSnapshotRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetCopyResourceInstanceSnapshotRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetRegion

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *FleetCopyResourceInstanceSnapshotRequest) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


### GetToken

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCopyResourceInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCopyResourceInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


