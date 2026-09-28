# DeploymentCheckpointChild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocker** | Pointer to [**CheckpointBlocker**](CheckpointBlocker.md) |  | [optional] 
**CompletedAt** | Pointer to **string** | The observed time the child completed, in RFC3339 format. Absent while the child is still in flight or when no completion evidence exists. | [optional] 
**Counts** | Pointer to [**CheckpointCounts**](CheckpointCounts.md) |  | [optional] 
**EnteredAt** | Pointer to **string** | The observed time the child entered its current state, in RFC3339 format. Absent when no entry evidence exists. | [optional] 
**Id** | Pointer to **string** | Stable identity of the child within its checkpoint, used only for equality and deduplication. Opaque to consumers. | [optional] 
**Kind** | Pointer to **string** | The kind of child unit: resource|hook|workload|module. Open string. | [optional] 
**Name** | Pointer to **string** | The operator-facing name of the child unit. | [optional] 
**State** | Pointer to **string** | The observed child state, using the existing task lifecycle vocabulary: Pending|Applying|AwaitingCondition|DriftMismatch|Failed|Succeeded. A child with no observed evidence is omitted rather than reported as Pending. | [optional] 
**Weight** | Pointer to **int64** | The child&#39;s execution-ordering weight, where the engine has one. Helm runs hook Jobs sequentially in ascending weight order, so a pending hook&#39;s weight is the explanation of why it has not started: it is waiting behind a lower weight rather than failing. Absent when the producer&#39;s engine has no ordering weight - Terraform resources never carry one - and a reported zero is a real weight, the default a chart gets when it declares none, so consumers must not read an absent weight as zero. | [optional] 

## Methods

### NewDeploymentCheckpointChild

`func NewDeploymentCheckpointChild() *DeploymentCheckpointChild`

NewDeploymentCheckpointChild instantiates a new DeploymentCheckpointChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCheckpointChildWithDefaults

`func NewDeploymentCheckpointChildWithDefaults() *DeploymentCheckpointChild`

NewDeploymentCheckpointChildWithDefaults instantiates a new DeploymentCheckpointChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocker

`func (o *DeploymentCheckpointChild) GetBlocker() CheckpointBlocker`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DeploymentCheckpointChild) GetBlockerOk() (*CheckpointBlocker, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DeploymentCheckpointChild) SetBlocker(v CheckpointBlocker)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DeploymentCheckpointChild) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DeploymentCheckpointChild) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DeploymentCheckpointChild) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DeploymentCheckpointChild) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DeploymentCheckpointChild) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCounts

`func (o *DeploymentCheckpointChild) GetCounts() CheckpointCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DeploymentCheckpointChild) GetCountsOk() (*CheckpointCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DeploymentCheckpointChild) SetCounts(v CheckpointCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DeploymentCheckpointChild) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetEnteredAt

`func (o *DeploymentCheckpointChild) GetEnteredAt() string`

GetEnteredAt returns the EnteredAt field if non-nil, zero value otherwise.

### GetEnteredAtOk

`func (o *DeploymentCheckpointChild) GetEnteredAtOk() (*string, bool)`

GetEnteredAtOk returns a tuple with the EnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredAt

`func (o *DeploymentCheckpointChild) SetEnteredAt(v string)`

SetEnteredAt sets EnteredAt field to given value.

### HasEnteredAt

`func (o *DeploymentCheckpointChild) HasEnteredAt() bool`

HasEnteredAt returns a boolean if a field has been set.

### GetId

`func (o *DeploymentCheckpointChild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentCheckpointChild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentCheckpointChild) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentCheckpointChild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DeploymentCheckpointChild) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeploymentCheckpointChild) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeploymentCheckpointChild) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeploymentCheckpointChild) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DeploymentCheckpointChild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentCheckpointChild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentCheckpointChild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentCheckpointChild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *DeploymentCheckpointChild) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentCheckpointChild) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentCheckpointChild) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentCheckpointChild) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWeight

`func (o *DeploymentCheckpointChild) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DeploymentCheckpointChild) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DeploymentCheckpointChild) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DeploymentCheckpointChild) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


