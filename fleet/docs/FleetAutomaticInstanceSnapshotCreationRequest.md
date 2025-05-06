# FleetAutomaticInstanceSnapshotCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetAutomaticInstanceSnapshotCreationRequest

`func NewFleetAutomaticInstanceSnapshotCreationRequest(environmentId string, instanceId string, serviceId string, token string, ) *FleetAutomaticInstanceSnapshotCreationRequest`

NewFleetAutomaticInstanceSnapshotCreationRequest instantiates a new FleetAutomaticInstanceSnapshotCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAutomaticInstanceSnapshotCreationRequestWithDefaults

`func NewFleetAutomaticInstanceSnapshotCreationRequestWithDefaults() *FleetAutomaticInstanceSnapshotCreationRequest`

NewFleetAutomaticInstanceSnapshotCreationRequestWithDefaults instantiates a new FleetAutomaticInstanceSnapshotCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetAutomaticInstanceSnapshotCreationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


