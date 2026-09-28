# ManagedArtifactSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactCount** | **int64** | Number of managed charts and images included in this public sync view. | 
**Artifacts** | Pointer to [**[]ManagedArtifactSyncArtifact**](ManagedArtifactSyncArtifact.md) |  | [optional] 
**BundleVersion** | **string** |  | 
**CompletedArtifactCount** | **int64** | Number of managed charts and images copied successfully. | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FailedArtifactCount** | **int64** | Number of managed charts and images that failed to copy. | 
**FailureCategory** | Pointer to **string** | Sanitized category for a managed artifact synchronization failure. | [optional] 
**FailureMessage** | Pointer to **string** | Sanitized operator-facing failure summary. | [optional] 
**Id** | **string** |  | 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** | Lifecycle status for one provisioner-target synchronization. | 
**Target** | [**ManagedArtifactTarget**](ManagedArtifactTarget.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewManagedArtifactSync

`func NewManagedArtifactSync(artifactCount int64, bundleVersion string, completedArtifactCount int64, failedArtifactCount int64, id string, status string, target ManagedArtifactTarget, ) *ManagedArtifactSync`

NewManagedArtifactSync instantiates a new ManagedArtifactSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactSyncWithDefaults

`func NewManagedArtifactSyncWithDefaults() *ManagedArtifactSync`

NewManagedArtifactSyncWithDefaults instantiates a new ManagedArtifactSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactCount

`func (o *ManagedArtifactSync) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *ManagedArtifactSync) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *ManagedArtifactSync) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetArtifacts

`func (o *ManagedArtifactSync) GetArtifacts() []ManagedArtifactSyncArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ManagedArtifactSync) GetArtifactsOk() (*[]ManagedArtifactSyncArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ManagedArtifactSync) SetArtifacts(v []ManagedArtifactSyncArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ManagedArtifactSync) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBundleVersion

`func (o *ManagedArtifactSync) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedArtifactSync) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedArtifactSync) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetCompletedArtifactCount

`func (o *ManagedArtifactSync) GetCompletedArtifactCount() int64`

GetCompletedArtifactCount returns the CompletedArtifactCount field if non-nil, zero value otherwise.

### GetCompletedArtifactCountOk

`func (o *ManagedArtifactSync) GetCompletedArtifactCountOk() (*int64, bool)`

GetCompletedArtifactCountOk returns a tuple with the CompletedArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedArtifactCount

`func (o *ManagedArtifactSync) SetCompletedArtifactCount(v int64)`

SetCompletedArtifactCount sets CompletedArtifactCount field to given value.


### GetCreatedAt

`func (o *ManagedArtifactSync) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedArtifactSync) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedArtifactSync) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedArtifactSync) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFailedArtifactCount

`func (o *ManagedArtifactSync) GetFailedArtifactCount() int64`

GetFailedArtifactCount returns the FailedArtifactCount field if non-nil, zero value otherwise.

### GetFailedArtifactCountOk

`func (o *ManagedArtifactSync) GetFailedArtifactCountOk() (*int64, bool)`

GetFailedArtifactCountOk returns a tuple with the FailedArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedArtifactCount

`func (o *ManagedArtifactSync) SetFailedArtifactCount(v int64)`

SetFailedArtifactCount sets FailedArtifactCount field to given value.


### GetFailureCategory

`func (o *ManagedArtifactSync) GetFailureCategory() string`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *ManagedArtifactSync) GetFailureCategoryOk() (*string, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *ManagedArtifactSync) SetFailureCategory(v string)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *ManagedArtifactSync) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ManagedArtifactSync) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ManagedArtifactSync) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ManagedArtifactSync) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ManagedArtifactSync) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *ManagedArtifactSync) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedArtifactSync) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedArtifactSync) SetId(v string)`

SetId sets Id field to given value.


### GetLastTransitionTime

`func (o *ManagedArtifactSync) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedArtifactSync) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedArtifactSync) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedArtifactSync) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedArtifactSync) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedArtifactSync) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedArtifactSync) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTarget

`func (o *ManagedArtifactSync) GetTarget() ManagedArtifactTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ManagedArtifactSync) GetTargetOk() (*ManagedArtifactTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ManagedArtifactSync) SetTarget(v ManagedArtifactTarget)`

SetTarget sets Target field to given value.


### GetUpdatedAt

`func (o *ManagedArtifactSync) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedArtifactSync) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedArtifactSync) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManagedArtifactSync) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


