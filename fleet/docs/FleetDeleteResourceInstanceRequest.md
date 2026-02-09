# FleetDeleteResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**SkipFinalSnapshot** | Pointer to **bool** | If true, skip the automatic final snapshot before deletion even if snapshotBeforeDeletion is enabled for the resource. | [optional] [default to false]
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetDeleteResourceInstanceRequest

`func NewFleetDeleteResourceInstanceRequest(environmentId string, instanceId string, resourceId string, serviceId string, token string, ) *FleetDeleteResourceInstanceRequest`

NewFleetDeleteResourceInstanceRequest instantiates a new FleetDeleteResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDeleteResourceInstanceRequestWithDefaults

`func NewFleetDeleteResourceInstanceRequestWithDefaults() *FleetDeleteResourceInstanceRequest`

NewFleetDeleteResourceInstanceRequestWithDefaults instantiates a new FleetDeleteResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetDeleteResourceInstanceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDeleteResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDeleteResourceInstanceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetDeleteResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetDeleteResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetDeleteResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetResourceId

`func (o *FleetDeleteResourceInstanceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetDeleteResourceInstanceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetDeleteResourceInstanceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *FleetDeleteResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDeleteResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDeleteResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest) GetSkipFinalSnapshot() bool`

GetSkipFinalSnapshot returns the SkipFinalSnapshot field if non-nil, zero value otherwise.

### GetSkipFinalSnapshotOk

`func (o *FleetDeleteResourceInstanceRequest) GetSkipFinalSnapshotOk() (*bool, bool)`

GetSkipFinalSnapshotOk returns a tuple with the SkipFinalSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest) SetSkipFinalSnapshot(v bool)`

SetSkipFinalSnapshot sets SkipFinalSnapshot field to given value.

### HasSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest) HasSkipFinalSnapshot() bool`

HasSkipFinalSnapshot returns a boolean if a field has been set.

### GetToken

`func (o *FleetDeleteResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetDeleteResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetDeleteResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


