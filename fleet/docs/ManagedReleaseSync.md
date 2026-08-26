# ManagedReleaseSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactResultCount** | **int64** |  | 
**ArtifactResults** | Pointer to [**[]ManagedReleaseSyncArtifactResult**](ManagedReleaseSyncArtifactResult.md) |  | [optional] 
**BundleRevisionId** | **string** |  | 
**BundleVersion** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DestinationRegistry** | Pointer to **string** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | **string** |  | 
**ProvisionerTargetId** | **string** |  | 
**Status** | **string** | Lifecycle status for a managed release revision or provisioner-target sync. | 
**SyncId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedReleaseSync

`func NewManagedReleaseSync(artifactResultCount int64, bundleRevisionId string, bundleVersion string, organizationId string, provisionerTargetId string, status string, syncId string, ) *ManagedReleaseSync`

NewManagedReleaseSync instantiates a new ManagedReleaseSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseSyncWithDefaults

`func NewManagedReleaseSyncWithDefaults() *ManagedReleaseSync`

NewManagedReleaseSyncWithDefaults instantiates a new ManagedReleaseSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactResultCount

`func (o *ManagedReleaseSync) GetArtifactResultCount() int64`

GetArtifactResultCount returns the ArtifactResultCount field if non-nil, zero value otherwise.

### GetArtifactResultCountOk

`func (o *ManagedReleaseSync) GetArtifactResultCountOk() (*int64, bool)`

GetArtifactResultCountOk returns a tuple with the ArtifactResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactResultCount

`func (o *ManagedReleaseSync) SetArtifactResultCount(v int64)`

SetArtifactResultCount sets ArtifactResultCount field to given value.


### GetArtifactResults

`func (o *ManagedReleaseSync) GetArtifactResults() []ManagedReleaseSyncArtifactResult`

GetArtifactResults returns the ArtifactResults field if non-nil, zero value otherwise.

### GetArtifactResultsOk

`func (o *ManagedReleaseSync) GetArtifactResultsOk() (*[]ManagedReleaseSyncArtifactResult, bool)`

GetArtifactResultsOk returns a tuple with the ArtifactResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactResults

`func (o *ManagedReleaseSync) SetArtifactResults(v []ManagedReleaseSyncArtifactResult)`

SetArtifactResults sets ArtifactResults field to given value.

### HasArtifactResults

`func (o *ManagedReleaseSync) HasArtifactResults() bool`

HasArtifactResults returns a boolean if a field has been set.

### GetBundleRevisionId

`func (o *ManagedReleaseSync) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *ManagedReleaseSync) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *ManagedReleaseSync) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetBundleVersion

`func (o *ManagedReleaseSync) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedReleaseSync) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedReleaseSync) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetCreatedAt

`func (o *ManagedReleaseSync) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedReleaseSync) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedReleaseSync) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedReleaseSync) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestinationRegistry

`func (o *ManagedReleaseSync) GetDestinationRegistry() string`

GetDestinationRegistry returns the DestinationRegistry field if non-nil, zero value otherwise.

### GetDestinationRegistryOk

`func (o *ManagedReleaseSync) GetDestinationRegistryOk() (*string, bool)`

GetDestinationRegistryOk returns a tuple with the DestinationRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegistry

`func (o *ManagedReleaseSync) SetDestinationRegistry(v string)`

SetDestinationRegistry sets DestinationRegistry field to given value.

### HasDestinationRegistry

`func (o *ManagedReleaseSync) HasDestinationRegistry() bool`

HasDestinationRegistry returns a boolean if a field has been set.

### GetLastError

`func (o *ManagedReleaseSync) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ManagedReleaseSync) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ManagedReleaseSync) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ManagedReleaseSync) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *ManagedReleaseSync) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedReleaseSync) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedReleaseSync) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedReleaseSync) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ManagedReleaseSync) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ManagedReleaseSync) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ManagedReleaseSync) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProvisionerTargetId

`func (o *ManagedReleaseSync) GetProvisionerTargetId() string`

GetProvisionerTargetId returns the ProvisionerTargetId field if non-nil, zero value otherwise.

### GetProvisionerTargetIdOk

`func (o *ManagedReleaseSync) GetProvisionerTargetIdOk() (*string, bool)`

GetProvisionerTargetIdOk returns a tuple with the ProvisionerTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerTargetId

`func (o *ManagedReleaseSync) SetProvisionerTargetId(v string)`

SetProvisionerTargetId sets ProvisionerTargetId field to given value.


### GetStatus

`func (o *ManagedReleaseSync) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedReleaseSync) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedReleaseSync) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSyncId

`func (o *ManagedReleaseSync) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *ManagedReleaseSync) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *ManagedReleaseSync) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.


### GetUpdatedAt

`func (o *ManagedReleaseSync) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedReleaseSync) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedReleaseSync) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManagedReleaseSync) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ManagedReleaseSync) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ManagedReleaseSync) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ManagedReleaseSync) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ManagedReleaseSync) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowRunId

`func (o *ManagedReleaseSync) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *ManagedReleaseSync) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *ManagedReleaseSync) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *ManagedReleaseSync) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


