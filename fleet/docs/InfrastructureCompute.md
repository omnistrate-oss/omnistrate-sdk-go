# InfrastructureCompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedObjects** | Pointer to [**[]InfrastructureAppliedObject**](InfrastructureAppliedObject.md) | Operator: objects the workflow steps applied that are not drawn among owners, pods, services or claims | [optional] 
**AppliedObjectsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Nodes** | Pointer to [**[]InfrastructureNode**](InfrastructureNode.md) |  | [optional] 
**NodesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Owners** | Pointer to [**[]InfrastructureOwner**](InfrastructureOwner.md) |  | [optional] 
**OwnersTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**PdbsAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**PdbsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**PodDisruptionBudgets** | Pointer to [**[]InfrastructurePDB**](InfrastructurePDB.md) |  | [optional] 
**Pods** | Pointer to [**[]InfrastructurePod**](InfrastructurePod.md) |  | [optional] 
**PodsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**SortOrder** | Pointer to **string** |  | [optional] 
**StringsTruncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInfrastructureCompute

`func NewInfrastructureCompute() *InfrastructureCompute`

NewInfrastructureCompute instantiates a new InfrastructureCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureComputeWithDefaults

`func NewInfrastructureComputeWithDefaults() *InfrastructureCompute`

NewInfrastructureComputeWithDefaults instantiates a new InfrastructureCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedObjects

`func (o *InfrastructureCompute) GetAppliedObjects() []InfrastructureAppliedObject`

GetAppliedObjects returns the AppliedObjects field if non-nil, zero value otherwise.

### GetAppliedObjectsOk

`func (o *InfrastructureCompute) GetAppliedObjectsOk() (*[]InfrastructureAppliedObject, bool)`

GetAppliedObjectsOk returns a tuple with the AppliedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedObjects

`func (o *InfrastructureCompute) SetAppliedObjects(v []InfrastructureAppliedObject)`

SetAppliedObjects sets AppliedObjects field to given value.

### HasAppliedObjects

`func (o *InfrastructureCompute) HasAppliedObjects() bool`

HasAppliedObjects returns a boolean if a field has been set.

### GetAppliedObjectsTruncation

`func (o *InfrastructureCompute) GetAppliedObjectsTruncation() CheckpointTruncation`

GetAppliedObjectsTruncation returns the AppliedObjectsTruncation field if non-nil, zero value otherwise.

### GetAppliedObjectsTruncationOk

`func (o *InfrastructureCompute) GetAppliedObjectsTruncationOk() (*CheckpointTruncation, bool)`

GetAppliedObjectsTruncationOk returns a tuple with the AppliedObjectsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedObjectsTruncation

`func (o *InfrastructureCompute) SetAppliedObjectsTruncation(v CheckpointTruncation)`

SetAppliedObjectsTruncation sets AppliedObjectsTruncation field to given value.

### HasAppliedObjectsTruncation

`func (o *InfrastructureCompute) HasAppliedObjectsTruncation() bool`

HasAppliedObjectsTruncation returns a boolean if a field has been set.

### GetAvailability

`func (o *InfrastructureCompute) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureCompute) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureCompute) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureCompute) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetNodes

`func (o *InfrastructureCompute) GetNodes() []InfrastructureNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InfrastructureCompute) GetNodesOk() (*[]InfrastructureNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InfrastructureCompute) SetNodes(v []InfrastructureNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InfrastructureCompute) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNodesTruncation

`func (o *InfrastructureCompute) GetNodesTruncation() CheckpointTruncation`

GetNodesTruncation returns the NodesTruncation field if non-nil, zero value otherwise.

### GetNodesTruncationOk

`func (o *InfrastructureCompute) GetNodesTruncationOk() (*CheckpointTruncation, bool)`

GetNodesTruncationOk returns a tuple with the NodesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesTruncation

`func (o *InfrastructureCompute) SetNodesTruncation(v CheckpointTruncation)`

SetNodesTruncation sets NodesTruncation field to given value.

### HasNodesTruncation

`func (o *InfrastructureCompute) HasNodesTruncation() bool`

HasNodesTruncation returns a boolean if a field has been set.

### GetOwners

`func (o *InfrastructureCompute) GetOwners() []InfrastructureOwner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *InfrastructureCompute) GetOwnersOk() (*[]InfrastructureOwner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *InfrastructureCompute) SetOwners(v []InfrastructureOwner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *InfrastructureCompute) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetOwnersTruncation

`func (o *InfrastructureCompute) GetOwnersTruncation() CheckpointTruncation`

GetOwnersTruncation returns the OwnersTruncation field if non-nil, zero value otherwise.

### GetOwnersTruncationOk

`func (o *InfrastructureCompute) GetOwnersTruncationOk() (*CheckpointTruncation, bool)`

GetOwnersTruncationOk returns a tuple with the OwnersTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnersTruncation

`func (o *InfrastructureCompute) SetOwnersTruncation(v CheckpointTruncation)`

SetOwnersTruncation sets OwnersTruncation field to given value.

### HasOwnersTruncation

`func (o *InfrastructureCompute) HasOwnersTruncation() bool`

HasOwnersTruncation returns a boolean if a field has been set.

### GetPdbsAvailability

`func (o *InfrastructureCompute) GetPdbsAvailability() InfrastructureAvailability`

GetPdbsAvailability returns the PdbsAvailability field if non-nil, zero value otherwise.

### GetPdbsAvailabilityOk

`func (o *InfrastructureCompute) GetPdbsAvailabilityOk() (*InfrastructureAvailability, bool)`

GetPdbsAvailabilityOk returns a tuple with the PdbsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbsAvailability

`func (o *InfrastructureCompute) SetPdbsAvailability(v InfrastructureAvailability)`

SetPdbsAvailability sets PdbsAvailability field to given value.

### HasPdbsAvailability

`func (o *InfrastructureCompute) HasPdbsAvailability() bool`

HasPdbsAvailability returns a boolean if a field has been set.

### GetPdbsTruncation

`func (o *InfrastructureCompute) GetPdbsTruncation() CheckpointTruncation`

GetPdbsTruncation returns the PdbsTruncation field if non-nil, zero value otherwise.

### GetPdbsTruncationOk

`func (o *InfrastructureCompute) GetPdbsTruncationOk() (*CheckpointTruncation, bool)`

GetPdbsTruncationOk returns a tuple with the PdbsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbsTruncation

`func (o *InfrastructureCompute) SetPdbsTruncation(v CheckpointTruncation)`

SetPdbsTruncation sets PdbsTruncation field to given value.

### HasPdbsTruncation

`func (o *InfrastructureCompute) HasPdbsTruncation() bool`

HasPdbsTruncation returns a boolean if a field has been set.

### GetPodDisruptionBudgets

`func (o *InfrastructureCompute) GetPodDisruptionBudgets() []InfrastructurePDB`

GetPodDisruptionBudgets returns the PodDisruptionBudgets field if non-nil, zero value otherwise.

### GetPodDisruptionBudgetsOk

`func (o *InfrastructureCompute) GetPodDisruptionBudgetsOk() (*[]InfrastructurePDB, bool)`

GetPodDisruptionBudgetsOk returns a tuple with the PodDisruptionBudgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodDisruptionBudgets

`func (o *InfrastructureCompute) SetPodDisruptionBudgets(v []InfrastructurePDB)`

SetPodDisruptionBudgets sets PodDisruptionBudgets field to given value.

### HasPodDisruptionBudgets

`func (o *InfrastructureCompute) HasPodDisruptionBudgets() bool`

HasPodDisruptionBudgets returns a boolean if a field has been set.

### GetPods

`func (o *InfrastructureCompute) GetPods() []InfrastructurePod`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *InfrastructureCompute) GetPodsOk() (*[]InfrastructurePod, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *InfrastructureCompute) SetPods(v []InfrastructurePod)`

SetPods sets Pods field to given value.

### HasPods

`func (o *InfrastructureCompute) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetPodsTruncation

`func (o *InfrastructureCompute) GetPodsTruncation() CheckpointTruncation`

GetPodsTruncation returns the PodsTruncation field if non-nil, zero value otherwise.

### GetPodsTruncationOk

`func (o *InfrastructureCompute) GetPodsTruncationOk() (*CheckpointTruncation, bool)`

GetPodsTruncationOk returns a tuple with the PodsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsTruncation

`func (o *InfrastructureCompute) SetPodsTruncation(v CheckpointTruncation)`

SetPodsTruncation sets PodsTruncation field to given value.

### HasPodsTruncation

`func (o *InfrastructureCompute) HasPodsTruncation() bool`

HasPodsTruncation returns a boolean if a field has been set.

### GetSortOrder

`func (o *InfrastructureCompute) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InfrastructureCompute) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InfrastructureCompute) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InfrastructureCompute) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStringsTruncated

`func (o *InfrastructureCompute) GetStringsTruncated() bool`

GetStringsTruncated returns the StringsTruncated field if non-nil, zero value otherwise.

### GetStringsTruncatedOk

`func (o *InfrastructureCompute) GetStringsTruncatedOk() (*bool, bool)`

GetStringsTruncatedOk returns a tuple with the StringsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringsTruncated

`func (o *InfrastructureCompute) SetStringsTruncated(v bool)`

SetStringsTruncated sets StringsTruncated field to given value.

### HasStringsTruncated

`func (o *InfrastructureCompute) HasStringsTruncated() bool`

HasStringsTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


