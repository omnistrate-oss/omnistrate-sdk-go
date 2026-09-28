# DeploymentCellManagedArtifactSyncEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactCount** | **int64** |  | 
**ArtifactIndex** | **int64** |  | 
**ArtifactKey** | Pointer to **string** |  | [optional] 
**ArtifactType** | Pointer to **string** |  | [optional] 
**BundleRevisionId** | **string** |  | 
**BundleVersion** | **string** |  | 
**DestinationRef** | Pointer to **string** | Resolved destination reference written by the synchronization workflow. | [optional] 
**OwnerComponent** | Pointer to **string** |  | [optional] 
**RelativePath** | Pointer to **string** | Artifact path relative to the destination registry. | [optional] 
**SourceRef** | Pointer to **string** | Canonical source reference for the artifact. | [optional] 
**Status** | **string** | Per-artifact synchronization state. | 
**SyncId** | **string** |  | 
**Version** | Pointer to **string** | Artifact version or image tag derived from the canonical source reference. | [optional] 

## Methods

### NewDeploymentCellManagedArtifactSyncEventDetail

`func NewDeploymentCellManagedArtifactSyncEventDetail(artifactCount int64, artifactIndex int64, bundleRevisionId string, bundleVersion string, status string, syncId string, ) *DeploymentCellManagedArtifactSyncEventDetail`

NewDeploymentCellManagedArtifactSyncEventDetail instantiates a new DeploymentCellManagedArtifactSyncEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellManagedArtifactSyncEventDetailWithDefaults

`func NewDeploymentCellManagedArtifactSyncEventDetailWithDefaults() *DeploymentCellManagedArtifactSyncEventDetail`

NewDeploymentCellManagedArtifactSyncEventDetailWithDefaults instantiates a new DeploymentCellManagedArtifactSyncEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactCount

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetArtifactIndex

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactIndex() int64`

GetArtifactIndex returns the ArtifactIndex field if non-nil, zero value otherwise.

### GetArtifactIndexOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactIndexOk() (*int64, bool)`

GetArtifactIndexOk returns a tuple with the ArtifactIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactIndex

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetArtifactIndex(v int64)`

SetArtifactIndex sets ArtifactIndex field to given value.


### GetArtifactKey

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.

### HasArtifactKey

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasArtifactKey() bool`

HasArtifactKey returns a boolean if a field has been set.

### GetArtifactType

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetBundleRevisionId

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetBundleVersion

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetDestinationRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetDestinationRef() string`

GetDestinationRef returns the DestinationRef field if non-nil, zero value otherwise.

### GetDestinationRefOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetDestinationRefOk() (*string, bool)`

GetDestinationRefOk returns a tuple with the DestinationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetDestinationRef(v string)`

SetDestinationRef sets DestinationRef field to given value.

### HasDestinationRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasDestinationRef() bool`

HasDestinationRef returns a boolean if a field has been set.

### GetOwnerComponent

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetOwnerComponent() string`

GetOwnerComponent returns the OwnerComponent field if non-nil, zero value otherwise.

### GetOwnerComponentOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetOwnerComponentOk() (*string, bool)`

GetOwnerComponentOk returns a tuple with the OwnerComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerComponent

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetOwnerComponent(v string)`

SetOwnerComponent sets OwnerComponent field to given value.

### HasOwnerComponent

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasOwnerComponent() bool`

HasOwnerComponent returns a boolean if a field has been set.

### GetRelativePath

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetSourceRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSyncId

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.


### GetVersion

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentCellManagedArtifactSyncEventDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentCellManagedArtifactSyncEventDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentCellManagedArtifactSyncEventDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


