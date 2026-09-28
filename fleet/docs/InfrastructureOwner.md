# InfrastructureOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedBy** | Pointer to [**InfrastructureAppliedBy**](InfrastructureAppliedBy.md) |  | [optional] 
**Controller** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**DesiredReplicas** | Pointer to **int64** |  | [optional] 
**ReadyReplicas** | Pointer to **int64** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 

## Methods

### NewInfrastructureOwner

`func NewInfrastructureOwner() *InfrastructureOwner`

NewInfrastructureOwner instantiates a new InfrastructureOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureOwnerWithDefaults

`func NewInfrastructureOwnerWithDefaults() *InfrastructureOwner`

NewInfrastructureOwnerWithDefaults instantiates a new InfrastructureOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedBy

`func (o *InfrastructureOwner) GetAppliedBy() InfrastructureAppliedBy`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *InfrastructureOwner) GetAppliedByOk() (*InfrastructureAppliedBy, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *InfrastructureOwner) SetAppliedBy(v InfrastructureAppliedBy)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *InfrastructureOwner) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetController

`func (o *InfrastructureOwner) GetController() InfrastructureObjectReference`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InfrastructureOwner) GetControllerOk() (*InfrastructureObjectReference, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InfrastructureOwner) SetController(v InfrastructureObjectReference)`

SetController sets Controller field to given value.

### HasController

`func (o *InfrastructureOwner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *InfrastructureOwner) GetDesiredReplicas() int64`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *InfrastructureOwner) GetDesiredReplicasOk() (*int64, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *InfrastructureOwner) SetDesiredReplicas(v int64)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *InfrastructureOwner) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *InfrastructureOwner) GetReadyReplicas() int64`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *InfrastructureOwner) GetReadyReplicasOk() (*int64, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *InfrastructureOwner) SetReadyReplicas(v int64)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *InfrastructureOwner) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureOwner) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureOwner) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureOwner) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureOwner) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


