# DeploymentCheckpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The explicit action the operation is performing: install|upgrade|rollback|uninstall for Helm, init|plan|apply|output|destroy for Terraform. Consumers must not infer the action from state, and must suppress rather than render progress for an operation whose action is absent when the distinction matters, because a destroy is otherwise indistinguishable from an apply. | [optional] 
**Blocker** | Pointer to [**CheckpointBlocker**](CheckpointBlocker.md) |  | [optional] 
**ChildTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Children** | Pointer to [**[]DeploymentCheckpointChild**](DeploymentCheckpointChild.md) | The child units observed inside this checkpoint - per-resource rows for infrastructure and Terraform, hook Jobs and workloads for Helm. Only children actually returned are listed, and childTruncation discloses whether more exist; a child with no observed evidence is omitted rather than reported as pending. The list is bounded by the producer and not by this schema: the list-shaped caps and the truncation disclosure that pairs with them live on the producer side, as every other cap in this contract does, because a required maximum here would make a generated client reject an entire fleet API response over a telemetry field. | [optional] 
**CompletedAt** | Pointer to **string** | The observed time this checkpoint completed, in RFC3339 format. Absent while the checkpoint is in flight or when no completion evidence exists. | [optional] 
**Counts** | Pointer to [**CheckpointCounts**](CheckpointCounts.md) |  | [optional] 
**EnteredAt** | Pointer to **string** | The observed time this checkpoint was entered, in RFC3339 format. Absent when no entry evidence exists; never synthesized from the current poll. | [optional] 
**Id** | Pointer to **string** | The deterministic transition identity, composed of resource type, stable resource identifier, operation identity, and checkpoint occurrence. It is stable across retries and never incorporates a poll time, an attempt counter, or a process-local value. Opaque to consumers: use it only for equality, deduplication, and merge. A checkpoint without an id must be discarded rather than merged. | [optional] 
**Name** | Pointer to **string** | The checkpoint identifier. Open string. For Terraform the vocabulary is the observed CLI action sequence - init|plan|apply|output|destroy - and a declared but unobserved action is omitted rather than reported as pending. For Helm the vocabulary covers preparation|install|upgrade|hooks|readiness. | [optional] 
**OperationCorrelationId** | Pointer to **string** | The retry-stable correlator for the operation, opaque to consumers. For Terraform this is the generation-hash prefix; for Helm it is derived from the release revision and status. Consumers must discard a checkpoint whose correlator does not match the operation being viewed rather than ranking it, so that a newer revision cannot contaminate an older workflow. | [optional] 
**OperationId** | Pointer to **string** | Attempt-scoped operation evidence, opaque to consumers. For Terraform this is the full generation-hash-and-nonce operation identifier, which is re-minted on every retry of a monolithic apply and is therefore not a correlator. | [optional] 
**Provenance** | Pointer to **string** | How a projected checkpoint was obtained. &#39;reported&#39; means a producer reported it for this operation, &#39;persisted&#39; means it was read back from a durable control-plane record, and &#39;derived&#39; means the control plane inferred it from other observations. Live-derived and sample content is never labelled &#39;persisted&#39;. This is an open string so that new sources can be added without breaking existing consumers; an unrecognized value must be treated as &#39;derived&#39;, and a checkpoint with no provenance at all must be discarded rather than rendered. | [optional] 
**SchemaVersion** | Pointer to **int64** | The schema version of this checkpoint block. Consumers must accept an unknown version by ignoring fields they do not recognize, never by discarding the response. | [optional] 
**State** | Pointer to **string** | The observed checkpoint state, using the existing task lifecycle vocabulary rather than a second state model: Pending|Applying|AwaitingCondition|DriftMismatch|Failed|Succeeded. A checkpoint with no observed entry evidence is omitted entirely. | [optional] 

## Methods

### NewDeploymentCheckpoint

`func NewDeploymentCheckpoint() *DeploymentCheckpoint`

NewDeploymentCheckpoint instantiates a new DeploymentCheckpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCheckpointWithDefaults

`func NewDeploymentCheckpointWithDefaults() *DeploymentCheckpoint`

NewDeploymentCheckpointWithDefaults instantiates a new DeploymentCheckpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentCheckpoint) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentCheckpoint) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentCheckpoint) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentCheckpoint) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlocker

`func (o *DeploymentCheckpoint) GetBlocker() CheckpointBlocker`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DeploymentCheckpoint) GetBlockerOk() (*CheckpointBlocker, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DeploymentCheckpoint) SetBlocker(v CheckpointBlocker)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DeploymentCheckpoint) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetChildTruncation

`func (o *DeploymentCheckpoint) GetChildTruncation() CheckpointTruncation`

GetChildTruncation returns the ChildTruncation field if non-nil, zero value otherwise.

### GetChildTruncationOk

`func (o *DeploymentCheckpoint) GetChildTruncationOk() (*CheckpointTruncation, bool)`

GetChildTruncationOk returns a tuple with the ChildTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTruncation

`func (o *DeploymentCheckpoint) SetChildTruncation(v CheckpointTruncation)`

SetChildTruncation sets ChildTruncation field to given value.

### HasChildTruncation

`func (o *DeploymentCheckpoint) HasChildTruncation() bool`

HasChildTruncation returns a boolean if a field has been set.

### GetChildren

`func (o *DeploymentCheckpoint) GetChildren() []DeploymentCheckpointChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DeploymentCheckpoint) GetChildrenOk() (*[]DeploymentCheckpointChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DeploymentCheckpoint) SetChildren(v []DeploymentCheckpointChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *DeploymentCheckpoint) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DeploymentCheckpoint) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DeploymentCheckpoint) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DeploymentCheckpoint) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DeploymentCheckpoint) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCounts

`func (o *DeploymentCheckpoint) GetCounts() CheckpointCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DeploymentCheckpoint) GetCountsOk() (*CheckpointCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DeploymentCheckpoint) SetCounts(v CheckpointCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DeploymentCheckpoint) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetEnteredAt

`func (o *DeploymentCheckpoint) GetEnteredAt() string`

GetEnteredAt returns the EnteredAt field if non-nil, zero value otherwise.

### GetEnteredAtOk

`func (o *DeploymentCheckpoint) GetEnteredAtOk() (*string, bool)`

GetEnteredAtOk returns a tuple with the EnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredAt

`func (o *DeploymentCheckpoint) SetEnteredAt(v string)`

SetEnteredAt sets EnteredAt field to given value.

### HasEnteredAt

`func (o *DeploymentCheckpoint) HasEnteredAt() bool`

HasEnteredAt returns a boolean if a field has been set.

### GetId

`func (o *DeploymentCheckpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentCheckpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentCheckpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentCheckpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentCheckpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentCheckpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentCheckpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentCheckpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationCorrelationId

`func (o *DeploymentCheckpoint) GetOperationCorrelationId() string`

GetOperationCorrelationId returns the OperationCorrelationId field if non-nil, zero value otherwise.

### GetOperationCorrelationIdOk

`func (o *DeploymentCheckpoint) GetOperationCorrelationIdOk() (*string, bool)`

GetOperationCorrelationIdOk returns a tuple with the OperationCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCorrelationId

`func (o *DeploymentCheckpoint) SetOperationCorrelationId(v string)`

SetOperationCorrelationId sets OperationCorrelationId field to given value.

### HasOperationCorrelationId

`func (o *DeploymentCheckpoint) HasOperationCorrelationId() bool`

HasOperationCorrelationId returns a boolean if a field has been set.

### GetOperationId

`func (o *DeploymentCheckpoint) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *DeploymentCheckpoint) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *DeploymentCheckpoint) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *DeploymentCheckpoint) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetProvenance

`func (o *DeploymentCheckpoint) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DeploymentCheckpoint) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DeploymentCheckpoint) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DeploymentCheckpoint) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DeploymentCheckpoint) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeploymentCheckpoint) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeploymentCheckpoint) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DeploymentCheckpoint) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetState

`func (o *DeploymentCheckpoint) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentCheckpoint) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentCheckpoint) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentCheckpoint) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


