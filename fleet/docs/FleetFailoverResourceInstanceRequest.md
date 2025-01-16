# FleetFailoverResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**FailedReplicaAction** | Pointer to **string** | The action to take for the replica that has failed | [optional] 
**FailedReplicaID** | **string** | The failed replica ID | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetFailoverResourceInstanceRequest

`func NewFleetFailoverResourceInstanceRequest(environmentId string, failedReplicaID string, instanceId string, serviceId string, token string, ) *FleetFailoverResourceInstanceRequest`

NewFleetFailoverResourceInstanceRequest instantiates a new FleetFailoverResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetFailoverResourceInstanceRequestWithDefaults

`func NewFleetFailoverResourceInstanceRequestWithDefaults() *FleetFailoverResourceInstanceRequest`

NewFleetFailoverResourceInstanceRequestWithDefaults instantiates a new FleetFailoverResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetFailoverResourceInstanceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetFailoverResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetFailoverResourceInstanceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest) GetFailedReplicaAction() string`

GetFailedReplicaAction returns the FailedReplicaAction field if non-nil, zero value otherwise.

### GetFailedReplicaActionOk

`func (o *FleetFailoverResourceInstanceRequest) GetFailedReplicaActionOk() (*string, bool)`

GetFailedReplicaActionOk returns a tuple with the FailedReplicaAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest) SetFailedReplicaAction(v string)`

SetFailedReplicaAction sets FailedReplicaAction field to given value.

### HasFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest) HasFailedReplicaAction() bool`

HasFailedReplicaAction returns a boolean if a field has been set.

### GetFailedReplicaID

`func (o *FleetFailoverResourceInstanceRequest) GetFailedReplicaID() string`

GetFailedReplicaID returns the FailedReplicaID field if non-nil, zero value otherwise.

### GetFailedReplicaIDOk

`func (o *FleetFailoverResourceInstanceRequest) GetFailedReplicaIDOk() (*string, bool)`

GetFailedReplicaIDOk returns a tuple with the FailedReplicaID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaID

`func (o *FleetFailoverResourceInstanceRequest) SetFailedReplicaID(v string)`

SetFailedReplicaID sets FailedReplicaID field to given value.


### GetInstanceId

`func (o *FleetFailoverResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetFailoverResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetFailoverResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetFailoverResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetFailoverResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetFailoverResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetFailoverResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetFailoverResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetFailoverResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


