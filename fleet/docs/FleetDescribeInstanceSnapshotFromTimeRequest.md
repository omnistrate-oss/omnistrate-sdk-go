# FleetDescribeInstanceSnapshotFromTimeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**TargetRestoreTime** | **string** | The target restore time | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetDescribeInstanceSnapshotFromTimeRequest

`func NewFleetDescribeInstanceSnapshotFromTimeRequest(environmentId string, instanceId string, serviceId string, targetRestoreTime string, token string, ) *FleetDescribeInstanceSnapshotFromTimeRequest`

NewFleetDescribeInstanceSnapshotFromTimeRequest instantiates a new FleetDescribeInstanceSnapshotFromTimeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeInstanceSnapshotFromTimeRequestWithDefaults

`func NewFleetDescribeInstanceSnapshotFromTimeRequestWithDefaults() *FleetDescribeInstanceSnapshotFromTimeRequest`

NewFleetDescribeInstanceSnapshotFromTimeRequestWithDefaults instantiates a new FleetDescribeInstanceSnapshotFromTimeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetRestoreTime

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetTargetRestoreTime() string`

GetTargetRestoreTime returns the TargetRestoreTime field if non-nil, zero value otherwise.

### GetTargetRestoreTimeOk

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetTargetRestoreTimeOk() (*string, bool)`

GetTargetRestoreTimeOk returns a tuple with the TargetRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRestoreTime

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) SetTargetRestoreTime(v string)`

SetTargetRestoreTime sets TargetRestoreTime field to given value.


### GetToken

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetDescribeInstanceSnapshotFromTimeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


