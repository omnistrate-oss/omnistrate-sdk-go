# ManagedReleaseRevision3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactCount** | **int64** |  | 
**Artifacts** | Pointer to [**[]ManagedReleaseArtifact4**](ManagedReleaseArtifact4.md) |  | [optional] 
**BundleRevisionId** | **string** |  | 
**BundleVersion** | **string** |  | 
**ComponentRevisions** | [**[]ManagedReleaseComponentRevision**](ManagedReleaseComponentRevision.md) |  | 
**ContentHash** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**ReleaseComponent** | **string** |  | 
**ReleaseSequence** | **int64** |  | 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedReleaseRevision3

`func NewManagedReleaseRevision3(artifactCount int64, bundleRevisionId string, bundleVersion string, componentRevisions []ManagedReleaseComponentRevision, releaseComponent string, releaseSequence int64, status string, ) *ManagedReleaseRevision3`

NewManagedReleaseRevision3 instantiates a new ManagedReleaseRevision3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseRevision3WithDefaults

`func NewManagedReleaseRevision3WithDefaults() *ManagedReleaseRevision3`

NewManagedReleaseRevision3WithDefaults instantiates a new ManagedReleaseRevision3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactCount

`func (o *ManagedReleaseRevision3) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *ManagedReleaseRevision3) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *ManagedReleaseRevision3) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetArtifacts

`func (o *ManagedReleaseRevision3) GetArtifacts() []ManagedReleaseArtifact4`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ManagedReleaseRevision3) GetArtifactsOk() (*[]ManagedReleaseArtifact4, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ManagedReleaseRevision3) SetArtifacts(v []ManagedReleaseArtifact4)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ManagedReleaseRevision3) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBundleRevisionId

`func (o *ManagedReleaseRevision3) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *ManagedReleaseRevision3) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *ManagedReleaseRevision3) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetBundleVersion

`func (o *ManagedReleaseRevision3) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedReleaseRevision3) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedReleaseRevision3) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetComponentRevisions

`func (o *ManagedReleaseRevision3) GetComponentRevisions() []ManagedReleaseComponentRevision`

GetComponentRevisions returns the ComponentRevisions field if non-nil, zero value otherwise.

### GetComponentRevisionsOk

`func (o *ManagedReleaseRevision3) GetComponentRevisionsOk() (*[]ManagedReleaseComponentRevision, bool)`

GetComponentRevisionsOk returns a tuple with the ComponentRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentRevisions

`func (o *ManagedReleaseRevision3) SetComponentRevisions(v []ManagedReleaseComponentRevision)`

SetComponentRevisions sets ComponentRevisions field to given value.


### GetContentHash

`func (o *ManagedReleaseRevision3) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *ManagedReleaseRevision3) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *ManagedReleaseRevision3) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *ManagedReleaseRevision3) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManagedReleaseRevision3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedReleaseRevision3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedReleaseRevision3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedReleaseRevision3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastError

`func (o *ManagedReleaseRevision3) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ManagedReleaseRevision3) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ManagedReleaseRevision3) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ManagedReleaseRevision3) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *ManagedReleaseRevision3) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedReleaseRevision3) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedReleaseRevision3) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedReleaseRevision3) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReleaseComponent

`func (o *ManagedReleaseRevision3) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *ManagedReleaseRevision3) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *ManagedReleaseRevision3) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.


### GetReleaseSequence

`func (o *ManagedReleaseRevision3) GetReleaseSequence() int64`

GetReleaseSequence returns the ReleaseSequence field if non-nil, zero value otherwise.

### GetReleaseSequenceOk

`func (o *ManagedReleaseRevision3) GetReleaseSequenceOk() (*int64, bool)`

GetReleaseSequenceOk returns a tuple with the ReleaseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSequence

`func (o *ManagedReleaseRevision3) SetReleaseSequence(v int64)`

SetReleaseSequence sets ReleaseSequence field to given value.


### GetReleasedAt

`func (o *ManagedReleaseRevision3) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *ManagedReleaseRevision3) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *ManagedReleaseRevision3) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *ManagedReleaseRevision3) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedReleaseRevision3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedReleaseRevision3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedReleaseRevision3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ManagedReleaseRevision3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedReleaseRevision3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedReleaseRevision3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManagedReleaseRevision3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ManagedReleaseRevision3) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ManagedReleaseRevision3) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ManagedReleaseRevision3) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ManagedReleaseRevision3) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowRunId

`func (o *ManagedReleaseRevision3) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *ManagedReleaseRevision3) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *ManagedReleaseRevision3) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *ManagedReleaseRevision3) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


