# DeploymentCheckpointSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The explicit action the operation is performing; see DeploymentCheckpoint.action. | [optional] 
**ActiveCheckpoint** | Pointer to **string** | The identifier of the checkpoint currently active, when one is. | [optional] 
**ActiveCheckpointId** | Pointer to **string** | The transition identity of the checkpoint currently active, when one is. | [optional] 
**ActiveSince** | Pointer to **string** | The observed time the active checkpoint was entered, in RFC3339 format. Absent when no entry evidence exists; never synthesized from the current poll. | [optional] 
**ActiveState** | Pointer to **string** | The state of the currently active checkpoint, using the existing task lifecycle vocabulary. | [optional] 
**Blocker** | Pointer to [**CheckpointBlocker**](CheckpointBlocker.md) |  | [optional] 
**CheckpointTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Counts** | Pointer to [**CheckpointCounts**](CheckpointCounts.md) |  | [optional] 
**OperationCorrelationId** | Pointer to **string** | The retry-stable operation correlator, opaque to consumers. Consumers must discard a summary whose correlator does not match the operation being viewed. | [optional] 
**OperationId** | Pointer to **string** | Attempt-scoped operation evidence, opaque to consumers. | [optional] 
**Provenance** | Pointer to **string** | How a projected checkpoint was obtained. &#39;reported&#39; means a producer reported it for this operation, &#39;persisted&#39; means it was read back from a durable control-plane record, and &#39;derived&#39; means the control plane inferred it from other observations. Live-derived and sample content is never labelled &#39;persisted&#39;. This is an open string so that new sources can be added without breaking existing consumers; an unrecognized value must be treated as &#39;derived&#39;, and a checkpoint with no provenance at all must be discarded rather than rendered. | [optional] 
**SchemaVersion** | Pointer to **int64** | The schema version of this summary block. | [optional] 
**SourceAvailable** | Pointer to **bool** | False when a progress source was consulted and explicitly could not answer, so consumers must render an unknown state rather than zero progress. Absent means availability was not evaluated, which is distinct from an explicit unavailable answer. | [optional] 
**SourceUnavailableCode** | Pointer to **string** | Bounded code explaining why no progress could be reported: ProgressUnavailable|ProgressExpired|ExecutionNotObserved|UnsupportedExecutionMode|FeatureDisabled. Present only when sourceAvailable is false. | [optional] 

## Methods

### NewDeploymentCheckpointSummary

`func NewDeploymentCheckpointSummary() *DeploymentCheckpointSummary`

NewDeploymentCheckpointSummary instantiates a new DeploymentCheckpointSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCheckpointSummaryWithDefaults

`func NewDeploymentCheckpointSummaryWithDefaults() *DeploymentCheckpointSummary`

NewDeploymentCheckpointSummaryWithDefaults instantiates a new DeploymentCheckpointSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentCheckpointSummary) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentCheckpointSummary) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentCheckpointSummary) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentCheckpointSummary) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActiveCheckpoint

`func (o *DeploymentCheckpointSummary) GetActiveCheckpoint() string`

GetActiveCheckpoint returns the ActiveCheckpoint field if non-nil, zero value otherwise.

### GetActiveCheckpointOk

`func (o *DeploymentCheckpointSummary) GetActiveCheckpointOk() (*string, bool)`

GetActiveCheckpointOk returns a tuple with the ActiveCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCheckpoint

`func (o *DeploymentCheckpointSummary) SetActiveCheckpoint(v string)`

SetActiveCheckpoint sets ActiveCheckpoint field to given value.

### HasActiveCheckpoint

`func (o *DeploymentCheckpointSummary) HasActiveCheckpoint() bool`

HasActiveCheckpoint returns a boolean if a field has been set.

### GetActiveCheckpointId

`func (o *DeploymentCheckpointSummary) GetActiveCheckpointId() string`

GetActiveCheckpointId returns the ActiveCheckpointId field if non-nil, zero value otherwise.

### GetActiveCheckpointIdOk

`func (o *DeploymentCheckpointSummary) GetActiveCheckpointIdOk() (*string, bool)`

GetActiveCheckpointIdOk returns a tuple with the ActiveCheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCheckpointId

`func (o *DeploymentCheckpointSummary) SetActiveCheckpointId(v string)`

SetActiveCheckpointId sets ActiveCheckpointId field to given value.

### HasActiveCheckpointId

`func (o *DeploymentCheckpointSummary) HasActiveCheckpointId() bool`

HasActiveCheckpointId returns a boolean if a field has been set.

### GetActiveSince

`func (o *DeploymentCheckpointSummary) GetActiveSince() string`

GetActiveSince returns the ActiveSince field if non-nil, zero value otherwise.

### GetActiveSinceOk

`func (o *DeploymentCheckpointSummary) GetActiveSinceOk() (*string, bool)`

GetActiveSinceOk returns a tuple with the ActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSince

`func (o *DeploymentCheckpointSummary) SetActiveSince(v string)`

SetActiveSince sets ActiveSince field to given value.

### HasActiveSince

`func (o *DeploymentCheckpointSummary) HasActiveSince() bool`

HasActiveSince returns a boolean if a field has been set.

### GetActiveState

`func (o *DeploymentCheckpointSummary) GetActiveState() string`

GetActiveState returns the ActiveState field if non-nil, zero value otherwise.

### GetActiveStateOk

`func (o *DeploymentCheckpointSummary) GetActiveStateOk() (*string, bool)`

GetActiveStateOk returns a tuple with the ActiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveState

`func (o *DeploymentCheckpointSummary) SetActiveState(v string)`

SetActiveState sets ActiveState field to given value.

### HasActiveState

`func (o *DeploymentCheckpointSummary) HasActiveState() bool`

HasActiveState returns a boolean if a field has been set.

### GetBlocker

`func (o *DeploymentCheckpointSummary) GetBlocker() CheckpointBlocker`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DeploymentCheckpointSummary) GetBlockerOk() (*CheckpointBlocker, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DeploymentCheckpointSummary) SetBlocker(v CheckpointBlocker)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DeploymentCheckpointSummary) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetCheckpointTruncation

`func (o *DeploymentCheckpointSummary) GetCheckpointTruncation() CheckpointTruncation`

GetCheckpointTruncation returns the CheckpointTruncation field if non-nil, zero value otherwise.

### GetCheckpointTruncationOk

`func (o *DeploymentCheckpointSummary) GetCheckpointTruncationOk() (*CheckpointTruncation, bool)`

GetCheckpointTruncationOk returns a tuple with the CheckpointTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointTruncation

`func (o *DeploymentCheckpointSummary) SetCheckpointTruncation(v CheckpointTruncation)`

SetCheckpointTruncation sets CheckpointTruncation field to given value.

### HasCheckpointTruncation

`func (o *DeploymentCheckpointSummary) HasCheckpointTruncation() bool`

HasCheckpointTruncation returns a boolean if a field has been set.

### GetCounts

`func (o *DeploymentCheckpointSummary) GetCounts() CheckpointCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DeploymentCheckpointSummary) GetCountsOk() (*CheckpointCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DeploymentCheckpointSummary) SetCounts(v CheckpointCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DeploymentCheckpointSummary) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetOperationCorrelationId

`func (o *DeploymentCheckpointSummary) GetOperationCorrelationId() string`

GetOperationCorrelationId returns the OperationCorrelationId field if non-nil, zero value otherwise.

### GetOperationCorrelationIdOk

`func (o *DeploymentCheckpointSummary) GetOperationCorrelationIdOk() (*string, bool)`

GetOperationCorrelationIdOk returns a tuple with the OperationCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCorrelationId

`func (o *DeploymentCheckpointSummary) SetOperationCorrelationId(v string)`

SetOperationCorrelationId sets OperationCorrelationId field to given value.

### HasOperationCorrelationId

`func (o *DeploymentCheckpointSummary) HasOperationCorrelationId() bool`

HasOperationCorrelationId returns a boolean if a field has been set.

### GetOperationId

`func (o *DeploymentCheckpointSummary) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *DeploymentCheckpointSummary) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *DeploymentCheckpointSummary) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *DeploymentCheckpointSummary) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetProvenance

`func (o *DeploymentCheckpointSummary) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DeploymentCheckpointSummary) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DeploymentCheckpointSummary) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DeploymentCheckpointSummary) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DeploymentCheckpointSummary) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeploymentCheckpointSummary) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeploymentCheckpointSummary) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DeploymentCheckpointSummary) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSourceAvailable

`func (o *DeploymentCheckpointSummary) GetSourceAvailable() bool`

GetSourceAvailable returns the SourceAvailable field if non-nil, zero value otherwise.

### GetSourceAvailableOk

`func (o *DeploymentCheckpointSummary) GetSourceAvailableOk() (*bool, bool)`

GetSourceAvailableOk returns a tuple with the SourceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailable

`func (o *DeploymentCheckpointSummary) SetSourceAvailable(v bool)`

SetSourceAvailable sets SourceAvailable field to given value.

### HasSourceAvailable

`func (o *DeploymentCheckpointSummary) HasSourceAvailable() bool`

HasSourceAvailable returns a boolean if a field has been set.

### GetSourceUnavailableCode

`func (o *DeploymentCheckpointSummary) GetSourceUnavailableCode() string`

GetSourceUnavailableCode returns the SourceUnavailableCode field if non-nil, zero value otherwise.

### GetSourceUnavailableCodeOk

`func (o *DeploymentCheckpointSummary) GetSourceUnavailableCodeOk() (*string, bool)`

GetSourceUnavailableCodeOk returns a tuple with the SourceUnavailableCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUnavailableCode

`func (o *DeploymentCheckpointSummary) SetSourceUnavailableCode(v string)`

SetSourceUnavailableCode sets SourceUnavailableCode field to given value.

### HasSourceUnavailableCode

`func (o *DeploymentCheckpointSummary) HasSourceUnavailableCode() bool`

HasSourceUnavailableCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


