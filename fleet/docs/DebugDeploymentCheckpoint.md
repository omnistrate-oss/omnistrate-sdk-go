# DebugDeploymentCheckpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The explicit action the operation is performing; see DeploymentCheckpoint.action. | [optional] 
**Blocker** | Pointer to [**CheckpointBlocker**](CheckpointBlocker.md) |  | [optional] 
**ChildTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Children** | Pointer to [**[]DeploymentCheckpointChild**](DeploymentCheckpointChild.md) | The child units observed inside this checkpoint. Only children actually returned are listed; childTruncation discloses whether more exist. | [optional] 
**CompletedAt** | Pointer to **string** | The observed completion time, in RFC3339 format. Absent when in flight. | [optional] 
**Counts** | Pointer to [**CheckpointCounts**](CheckpointCounts.md) |  | [optional] 
**EnteredAt** | Pointer to **string** | The observed entry time, in RFC3339 format. Absent when no evidence exists. | [optional] 
**Id** | Pointer to **string** | The deterministic transition identity; see DeploymentCheckpoint.id. Opaque. | [optional] 
**Name** | Pointer to **string** | The checkpoint identifier; see DeploymentCheckpoint.name. | [optional] 
**OperationCorrelationId** | Pointer to **string** | The retry-stable operation correlator, opaque to consumers. | [optional] 
**OperationId** | Pointer to **string** | Attempt-scoped operation evidence, opaque to consumers. | [optional] 
**Provenance** | Pointer to **string** | How a projected checkpoint was obtained. &#39;reported&#39; means a producer reported it for this operation, &#39;persisted&#39; means it was read back from a durable control-plane record, and &#39;derived&#39; means the control plane inferred it from other observations. Live-derived and sample content is never labelled &#39;persisted&#39;. This is an open string so that new sources can be added without breaking existing consumers; an unrecognized value must be treated as &#39;derived&#39;, and a checkpoint with no provenance at all must be discarded rather than rendered. | [optional] 
**SchemaVersion** | Pointer to **int64** | The schema version of this checkpoint block. | [optional] 
**State** | Pointer to **string** | The observed checkpoint state; see DeploymentCheckpoint.state. | [optional] 

## Methods

### NewDebugDeploymentCheckpoint

`func NewDebugDeploymentCheckpoint() *DebugDeploymentCheckpoint`

NewDebugDeploymentCheckpoint instantiates a new DebugDeploymentCheckpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugDeploymentCheckpointWithDefaults

`func NewDebugDeploymentCheckpointWithDefaults() *DebugDeploymentCheckpoint`

NewDebugDeploymentCheckpointWithDefaults instantiates a new DebugDeploymentCheckpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DebugDeploymentCheckpoint) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DebugDeploymentCheckpoint) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DebugDeploymentCheckpoint) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DebugDeploymentCheckpoint) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlocker

`func (o *DebugDeploymentCheckpoint) GetBlocker() CheckpointBlocker`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DebugDeploymentCheckpoint) GetBlockerOk() (*CheckpointBlocker, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DebugDeploymentCheckpoint) SetBlocker(v CheckpointBlocker)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DebugDeploymentCheckpoint) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetChildTruncation

`func (o *DebugDeploymentCheckpoint) GetChildTruncation() CheckpointTruncation`

GetChildTruncation returns the ChildTruncation field if non-nil, zero value otherwise.

### GetChildTruncationOk

`func (o *DebugDeploymentCheckpoint) GetChildTruncationOk() (*CheckpointTruncation, bool)`

GetChildTruncationOk returns a tuple with the ChildTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTruncation

`func (o *DebugDeploymentCheckpoint) SetChildTruncation(v CheckpointTruncation)`

SetChildTruncation sets ChildTruncation field to given value.

### HasChildTruncation

`func (o *DebugDeploymentCheckpoint) HasChildTruncation() bool`

HasChildTruncation returns a boolean if a field has been set.

### GetChildren

`func (o *DebugDeploymentCheckpoint) GetChildren() []DeploymentCheckpointChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DebugDeploymentCheckpoint) GetChildrenOk() (*[]DeploymentCheckpointChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DebugDeploymentCheckpoint) SetChildren(v []DeploymentCheckpointChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *DebugDeploymentCheckpoint) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DebugDeploymentCheckpoint) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DebugDeploymentCheckpoint) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DebugDeploymentCheckpoint) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DebugDeploymentCheckpoint) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCounts

`func (o *DebugDeploymentCheckpoint) GetCounts() CheckpointCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DebugDeploymentCheckpoint) GetCountsOk() (*CheckpointCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DebugDeploymentCheckpoint) SetCounts(v CheckpointCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DebugDeploymentCheckpoint) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetEnteredAt

`func (o *DebugDeploymentCheckpoint) GetEnteredAt() string`

GetEnteredAt returns the EnteredAt field if non-nil, zero value otherwise.

### GetEnteredAtOk

`func (o *DebugDeploymentCheckpoint) GetEnteredAtOk() (*string, bool)`

GetEnteredAtOk returns a tuple with the EnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredAt

`func (o *DebugDeploymentCheckpoint) SetEnteredAt(v string)`

SetEnteredAt sets EnteredAt field to given value.

### HasEnteredAt

`func (o *DebugDeploymentCheckpoint) HasEnteredAt() bool`

HasEnteredAt returns a boolean if a field has been set.

### GetId

`func (o *DebugDeploymentCheckpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebugDeploymentCheckpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebugDeploymentCheckpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DebugDeploymentCheckpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DebugDeploymentCheckpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebugDeploymentCheckpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebugDeploymentCheckpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DebugDeploymentCheckpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationCorrelationId

`func (o *DebugDeploymentCheckpoint) GetOperationCorrelationId() string`

GetOperationCorrelationId returns the OperationCorrelationId field if non-nil, zero value otherwise.

### GetOperationCorrelationIdOk

`func (o *DebugDeploymentCheckpoint) GetOperationCorrelationIdOk() (*string, bool)`

GetOperationCorrelationIdOk returns a tuple with the OperationCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCorrelationId

`func (o *DebugDeploymentCheckpoint) SetOperationCorrelationId(v string)`

SetOperationCorrelationId sets OperationCorrelationId field to given value.

### HasOperationCorrelationId

`func (o *DebugDeploymentCheckpoint) HasOperationCorrelationId() bool`

HasOperationCorrelationId returns a boolean if a field has been set.

### GetOperationId

`func (o *DebugDeploymentCheckpoint) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *DebugDeploymentCheckpoint) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *DebugDeploymentCheckpoint) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *DebugDeploymentCheckpoint) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetProvenance

`func (o *DebugDeploymentCheckpoint) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DebugDeploymentCheckpoint) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DebugDeploymentCheckpoint) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DebugDeploymentCheckpoint) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DebugDeploymentCheckpoint) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DebugDeploymentCheckpoint) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DebugDeploymentCheckpoint) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DebugDeploymentCheckpoint) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetState

`func (o *DebugDeploymentCheckpoint) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DebugDeploymentCheckpoint) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DebugDeploymentCheckpoint) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DebugDeploymentCheckpoint) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


