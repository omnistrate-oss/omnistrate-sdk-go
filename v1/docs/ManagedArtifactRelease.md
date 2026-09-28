# ManagedArtifactRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmenityCount** | **int64** | Number of managed Base Amenities represented by this release. | 
**ArtifactCount** | **int64** | Number of managed charts and images in this public release view. | 
**Artifacts** | Pointer to [**[]ManagedArtifactReleaseArtifact**](ManagedArtifactReleaseArtifact.md) | Managed Base Amenity charts and images only. Internal control-plane artifacts are excluded. | [optional] 
**BundleVersion** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**ReleaseSequence** | **int64** |  | 
**ReleasedAt** | **time.Time** |  | 
**TargetStatus** | [**ManagedArtifactTargetStatusSummary**](ManagedArtifactTargetStatusSummary.md) |  | 

## Methods

### NewManagedArtifactRelease

`func NewManagedArtifactRelease(amenityCount int64, artifactCount int64, bundleVersion string, releaseSequence int64, releasedAt time.Time, targetStatus ManagedArtifactTargetStatusSummary, ) *ManagedArtifactRelease`

NewManagedArtifactRelease instantiates a new ManagedArtifactRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactReleaseWithDefaults

`func NewManagedArtifactReleaseWithDefaults() *ManagedArtifactRelease`

NewManagedArtifactReleaseWithDefaults instantiates a new ManagedArtifactRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenityCount

`func (o *ManagedArtifactRelease) GetAmenityCount() int64`

GetAmenityCount returns the AmenityCount field if non-nil, zero value otherwise.

### GetAmenityCountOk

`func (o *ManagedArtifactRelease) GetAmenityCountOk() (*int64, bool)`

GetAmenityCountOk returns a tuple with the AmenityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityCount

`func (o *ManagedArtifactRelease) SetAmenityCount(v int64)`

SetAmenityCount sets AmenityCount field to given value.


### GetArtifactCount

`func (o *ManagedArtifactRelease) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *ManagedArtifactRelease) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *ManagedArtifactRelease) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetArtifacts

`func (o *ManagedArtifactRelease) GetArtifacts() []ManagedArtifactReleaseArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ManagedArtifactRelease) GetArtifactsOk() (*[]ManagedArtifactReleaseArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ManagedArtifactRelease) SetArtifacts(v []ManagedArtifactReleaseArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ManagedArtifactRelease) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBundleVersion

`func (o *ManagedArtifactRelease) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedArtifactRelease) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedArtifactRelease) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetCreatedAt

`func (o *ManagedArtifactRelease) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedArtifactRelease) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedArtifactRelease) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedArtifactRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *ManagedArtifactRelease) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedArtifactRelease) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedArtifactRelease) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedArtifactRelease) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReleaseSequence

`func (o *ManagedArtifactRelease) GetReleaseSequence() int64`

GetReleaseSequence returns the ReleaseSequence field if non-nil, zero value otherwise.

### GetReleaseSequenceOk

`func (o *ManagedArtifactRelease) GetReleaseSequenceOk() (*int64, bool)`

GetReleaseSequenceOk returns a tuple with the ReleaseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSequence

`func (o *ManagedArtifactRelease) SetReleaseSequence(v int64)`

SetReleaseSequence sets ReleaseSequence field to given value.


### GetReleasedAt

`func (o *ManagedArtifactRelease) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *ManagedArtifactRelease) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *ManagedArtifactRelease) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.


### GetTargetStatus

`func (o *ManagedArtifactRelease) GetTargetStatus() ManagedArtifactTargetStatusSummary`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *ManagedArtifactRelease) GetTargetStatusOk() (*ManagedArtifactTargetStatusSummary, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *ManagedArtifactRelease) SetTargetStatus(v ManagedArtifactTargetStatusSummary)`

SetTargetStatus sets TargetStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


