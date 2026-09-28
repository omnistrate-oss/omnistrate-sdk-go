# DeploymentProgressDiagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The explicit action the operation is performing; see DeploymentCheckpoint.action. | [optional] 
**Blocker** | Pointer to [**CheckpointBlocker**](CheckpointBlocker.md) |  | [optional] 
**CheckpointTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Checkpoints** | Pointer to [**[]DebugDeploymentCheckpoint**](DebugDeploymentCheckpoint.md) | The observed checkpoints for the operation, ordered by observed entry time. Only checkpoints with observed evidence appear; declared but unobserved phases are omitted. checkpointTruncation discloses whether more exist. | [optional] 
**Counts** | Pointer to [**CheckpointCounts**](CheckpointCounts.md) |  | [optional] 
**ObservedAt** | Pointer to **string** | The time this projection was assembled, in RFC3339 format. This is a server-side observation time and must never be used as a checkpoint transition time. | [optional] 
**OperationCorrelationId** | Pointer to **string** | The retry-stable operation correlator, opaque to consumers. | [optional] 
**OperationId** | Pointer to **string** | Attempt-scoped operation evidence, opaque to consumers. | [optional] 
**Provenance** | Pointer to **string** | How a projected checkpoint was obtained. &#39;reported&#39; means a producer reported it for this operation, &#39;persisted&#39; means it was read back from a durable control-plane record, and &#39;derived&#39; means the control plane inferred it from other observations. Live-derived and sample content is never labelled &#39;persisted&#39;. This is an open string so that new sources can be added without breaking existing consumers; an unrecognized value must be treated as &#39;derived&#39;, and a checkpoint with no provenance at all must be discarded rather than rendered. | [optional] 
**ResourceType** | Pointer to **string** | The resource type the checkpoints belong to: operatorCRD|genericCRD|helm|terraform|workload|cloudInfra|job|infraStack. | [optional] 
**SchemaVersion** | Pointer to **int64** | The schema version of this diagnosis block. | [optional] 
**SourceAvailable** | Pointer to **bool** | False when a progress source was consulted and explicitly could not answer, so consumers must render an unknown state rather than zero progress. Absent means availability was not evaluated. | [optional] 
**SourceUnavailableCode** | Pointer to **string** | Bounded code explaining why no progress could be reported: ProgressUnavailable|ProgressExpired|ExecutionNotObserved|UnsupportedExecutionMode|FeatureDisabled. Present only when sourceAvailable is false. | [optional] 

## Methods

### NewDeploymentProgressDiagnosis

`func NewDeploymentProgressDiagnosis() *DeploymentProgressDiagnosis`

NewDeploymentProgressDiagnosis instantiates a new DeploymentProgressDiagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProgressDiagnosisWithDefaults

`func NewDeploymentProgressDiagnosisWithDefaults() *DeploymentProgressDiagnosis`

NewDeploymentProgressDiagnosisWithDefaults instantiates a new DeploymentProgressDiagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentProgressDiagnosis) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentProgressDiagnosis) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentProgressDiagnosis) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentProgressDiagnosis) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlocker

`func (o *DeploymentProgressDiagnosis) GetBlocker() CheckpointBlocker`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DeploymentProgressDiagnosis) GetBlockerOk() (*CheckpointBlocker, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DeploymentProgressDiagnosis) SetBlocker(v CheckpointBlocker)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DeploymentProgressDiagnosis) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetCheckpointTruncation

`func (o *DeploymentProgressDiagnosis) GetCheckpointTruncation() CheckpointTruncation`

GetCheckpointTruncation returns the CheckpointTruncation field if non-nil, zero value otherwise.

### GetCheckpointTruncationOk

`func (o *DeploymentProgressDiagnosis) GetCheckpointTruncationOk() (*CheckpointTruncation, bool)`

GetCheckpointTruncationOk returns a tuple with the CheckpointTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointTruncation

`func (o *DeploymentProgressDiagnosis) SetCheckpointTruncation(v CheckpointTruncation)`

SetCheckpointTruncation sets CheckpointTruncation field to given value.

### HasCheckpointTruncation

`func (o *DeploymentProgressDiagnosis) HasCheckpointTruncation() bool`

HasCheckpointTruncation returns a boolean if a field has been set.

### GetCheckpoints

`func (o *DeploymentProgressDiagnosis) GetCheckpoints() []DebugDeploymentCheckpoint`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *DeploymentProgressDiagnosis) GetCheckpointsOk() (*[]DebugDeploymentCheckpoint, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *DeploymentProgressDiagnosis) SetCheckpoints(v []DebugDeploymentCheckpoint)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *DeploymentProgressDiagnosis) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.

### GetCounts

`func (o *DeploymentProgressDiagnosis) GetCounts() CheckpointCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *DeploymentProgressDiagnosis) GetCountsOk() (*CheckpointCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *DeploymentProgressDiagnosis) SetCounts(v CheckpointCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *DeploymentProgressDiagnosis) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetObservedAt

`func (o *DeploymentProgressDiagnosis) GetObservedAt() string`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *DeploymentProgressDiagnosis) GetObservedAtOk() (*string, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *DeploymentProgressDiagnosis) SetObservedAt(v string)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *DeploymentProgressDiagnosis) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetOperationCorrelationId

`func (o *DeploymentProgressDiagnosis) GetOperationCorrelationId() string`

GetOperationCorrelationId returns the OperationCorrelationId field if non-nil, zero value otherwise.

### GetOperationCorrelationIdOk

`func (o *DeploymentProgressDiagnosis) GetOperationCorrelationIdOk() (*string, bool)`

GetOperationCorrelationIdOk returns a tuple with the OperationCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCorrelationId

`func (o *DeploymentProgressDiagnosis) SetOperationCorrelationId(v string)`

SetOperationCorrelationId sets OperationCorrelationId field to given value.

### HasOperationCorrelationId

`func (o *DeploymentProgressDiagnosis) HasOperationCorrelationId() bool`

HasOperationCorrelationId returns a boolean if a field has been set.

### GetOperationId

`func (o *DeploymentProgressDiagnosis) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *DeploymentProgressDiagnosis) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *DeploymentProgressDiagnosis) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *DeploymentProgressDiagnosis) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetProvenance

`func (o *DeploymentProgressDiagnosis) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DeploymentProgressDiagnosis) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DeploymentProgressDiagnosis) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DeploymentProgressDiagnosis) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetResourceType

`func (o *DeploymentProgressDiagnosis) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DeploymentProgressDiagnosis) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DeploymentProgressDiagnosis) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DeploymentProgressDiagnosis) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DeploymentProgressDiagnosis) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeploymentProgressDiagnosis) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeploymentProgressDiagnosis) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DeploymentProgressDiagnosis) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSourceAvailable

`func (o *DeploymentProgressDiagnosis) GetSourceAvailable() bool`

GetSourceAvailable returns the SourceAvailable field if non-nil, zero value otherwise.

### GetSourceAvailableOk

`func (o *DeploymentProgressDiagnosis) GetSourceAvailableOk() (*bool, bool)`

GetSourceAvailableOk returns a tuple with the SourceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailable

`func (o *DeploymentProgressDiagnosis) SetSourceAvailable(v bool)`

SetSourceAvailable sets SourceAvailable field to given value.

### HasSourceAvailable

`func (o *DeploymentProgressDiagnosis) HasSourceAvailable() bool`

HasSourceAvailable returns a boolean if a field has been set.

### GetSourceUnavailableCode

`func (o *DeploymentProgressDiagnosis) GetSourceUnavailableCode() string`

GetSourceUnavailableCode returns the SourceUnavailableCode field if non-nil, zero value otherwise.

### GetSourceUnavailableCodeOk

`func (o *DeploymentProgressDiagnosis) GetSourceUnavailableCodeOk() (*string, bool)`

GetSourceUnavailableCodeOk returns a tuple with the SourceUnavailableCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUnavailableCode

`func (o *DeploymentProgressDiagnosis) SetSourceUnavailableCode(v string)`

SetSourceUnavailableCode sets SourceUnavailableCode field to given value.

### HasSourceUnavailableCode

`func (o *DeploymentProgressDiagnosis) HasSourceUnavailableCode() bool`

HasSourceUnavailableCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


