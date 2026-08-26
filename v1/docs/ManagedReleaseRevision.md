# ManagedReleaseRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactCount** | **int64** |  | 
**Artifacts** | Pointer to [**[]ManagedReleaseArtifact**](ManagedReleaseArtifact.md) |  | [optional] 
**BundleRevisionId** | **string** |  | 
**BundleVersion** | **string** |  | 
**ComponentRevisions** | [**[]ManagedReleaseComponentRevision**](ManagedReleaseComponentRevision.md) |  | 
**ContentHash** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**ReleaseComponent** | **string** | Component pipeline that owns a managed artifact subset. | 
**ReleaseSequence** | **int64** |  | 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** | Lifecycle status for a managed release revision or provisioner-target sync. | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedReleaseRevision

`func NewManagedReleaseRevision(artifactCount int64, bundleRevisionId string, bundleVersion string, componentRevisions []ManagedReleaseComponentRevision, releaseComponent string, releaseSequence int64, status string, ) *ManagedReleaseRevision`

NewManagedReleaseRevision instantiates a new ManagedReleaseRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseRevisionWithDefaults

`func NewManagedReleaseRevisionWithDefaults() *ManagedReleaseRevision`

NewManagedReleaseRevisionWithDefaults instantiates a new ManagedReleaseRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactCount

`func (o *ManagedReleaseRevision) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *ManagedReleaseRevision) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *ManagedReleaseRevision) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetArtifacts

`func (o *ManagedReleaseRevision) GetArtifacts() []ManagedReleaseArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ManagedReleaseRevision) GetArtifactsOk() (*[]ManagedReleaseArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ManagedReleaseRevision) SetArtifacts(v []ManagedReleaseArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ManagedReleaseRevision) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBundleRevisionId

`func (o *ManagedReleaseRevision) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *ManagedReleaseRevision) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *ManagedReleaseRevision) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetBundleVersion

`func (o *ManagedReleaseRevision) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedReleaseRevision) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedReleaseRevision) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetComponentRevisions

`func (o *ManagedReleaseRevision) GetComponentRevisions() []ManagedReleaseComponentRevision`

GetComponentRevisions returns the ComponentRevisions field if non-nil, zero value otherwise.

### GetComponentRevisionsOk

`func (o *ManagedReleaseRevision) GetComponentRevisionsOk() (*[]ManagedReleaseComponentRevision, bool)`

GetComponentRevisionsOk returns a tuple with the ComponentRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentRevisions

`func (o *ManagedReleaseRevision) SetComponentRevisions(v []ManagedReleaseComponentRevision)`

SetComponentRevisions sets ComponentRevisions field to given value.


### GetContentHash

`func (o *ManagedReleaseRevision) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *ManagedReleaseRevision) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *ManagedReleaseRevision) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *ManagedReleaseRevision) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManagedReleaseRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedReleaseRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedReleaseRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedReleaseRevision) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastError

`func (o *ManagedReleaseRevision) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ManagedReleaseRevision) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ManagedReleaseRevision) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ManagedReleaseRevision) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *ManagedReleaseRevision) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedReleaseRevision) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedReleaseRevision) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedReleaseRevision) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReleaseComponent

`func (o *ManagedReleaseRevision) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *ManagedReleaseRevision) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *ManagedReleaseRevision) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.


### GetReleaseSequence

`func (o *ManagedReleaseRevision) GetReleaseSequence() int64`

GetReleaseSequence returns the ReleaseSequence field if non-nil, zero value otherwise.

### GetReleaseSequenceOk

`func (o *ManagedReleaseRevision) GetReleaseSequenceOk() (*int64, bool)`

GetReleaseSequenceOk returns a tuple with the ReleaseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSequence

`func (o *ManagedReleaseRevision) SetReleaseSequence(v int64)`

SetReleaseSequence sets ReleaseSequence field to given value.


### GetReleasedAt

`func (o *ManagedReleaseRevision) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *ManagedReleaseRevision) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *ManagedReleaseRevision) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *ManagedReleaseRevision) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedReleaseRevision) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedReleaseRevision) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedReleaseRevision) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ManagedReleaseRevision) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedReleaseRevision) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedReleaseRevision) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManagedReleaseRevision) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ManagedReleaseRevision) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ManagedReleaseRevision) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ManagedReleaseRevision) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ManagedReleaseRevision) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowRunId

`func (o *ManagedReleaseRevision) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *ManagedReleaseRevision) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *ManagedReleaseRevision) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *ManagedReleaseRevision) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


