# FleetFailoverResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedReplicaAction** | Pointer to **string** | The failed replica action | [optional] 
**FailedReplicaID** | **string** | The failed replica ID | 

## Methods

### NewFleetFailoverResourceInstanceRequest2

`func NewFleetFailoverResourceInstanceRequest2(failedReplicaID string, ) *FleetFailoverResourceInstanceRequest2`

NewFleetFailoverResourceInstanceRequest2 instantiates a new FleetFailoverResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetFailoverResourceInstanceRequest2WithDefaults

`func NewFleetFailoverResourceInstanceRequest2WithDefaults() *FleetFailoverResourceInstanceRequest2`

NewFleetFailoverResourceInstanceRequest2WithDefaults instantiates a new FleetFailoverResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest2) GetFailedReplicaAction() string`

GetFailedReplicaAction returns the FailedReplicaAction field if non-nil, zero value otherwise.

### GetFailedReplicaActionOk

`func (o *FleetFailoverResourceInstanceRequest2) GetFailedReplicaActionOk() (*string, bool)`

GetFailedReplicaActionOk returns a tuple with the FailedReplicaAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest2) SetFailedReplicaAction(v string)`

SetFailedReplicaAction sets FailedReplicaAction field to given value.

### HasFailedReplicaAction

`func (o *FleetFailoverResourceInstanceRequest2) HasFailedReplicaAction() bool`

HasFailedReplicaAction returns a boolean if a field has been set.

### GetFailedReplicaID

`func (o *FleetFailoverResourceInstanceRequest2) GetFailedReplicaID() string`

GetFailedReplicaID returns the FailedReplicaID field if non-nil, zero value otherwise.

### GetFailedReplicaIDOk

`func (o *FleetFailoverResourceInstanceRequest2) GetFailedReplicaIDOk() (*string, bool)`

GetFailedReplicaIDOk returns a tuple with the FailedReplicaID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaID

`func (o *FleetFailoverResourceInstanceRequest2) SetFailedReplicaID(v string)`

SetFailedReplicaID sets FailedReplicaID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


