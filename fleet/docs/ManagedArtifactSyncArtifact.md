# ManagedArtifactSyncArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmenityName** | **string** | Canonical Base Amenity name associated with this artifact. Clients must use this field instead of parsing artifactKey. | 
**ArtifactKey** | **string** |  | 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**DestinationDigest** | Pointer to **string** |  | [optional] 
**DestinationRef** | Pointer to **string** | Materialized destination artifact reference without credentials. | [optional] 
**FailureCategory** | Pointer to **string** | Sanitized category for a managed artifact synchronization failure. | [optional] 
**FailureMessage** | Pointer to **string** | Sanitized operator-facing failure summary. | [optional] 
**SourceDigest** | Pointer to **string** |  | [optional] 
**SourceRef** | Pointer to **string** | Canonical source artifact reference without credentials or signed query parameters. | [optional] 
**Status** | **string** | Public synchronization status for one artifact in a managed release. | 
**Version** | **string** | Artifact version from the linked release revision, derived by the backend from its canonical source reference. | 

## Methods

### NewManagedArtifactSyncArtifact

`func NewManagedArtifactSyncArtifact(amenityName string, artifactKey string, status string, version string, ) *ManagedArtifactSyncArtifact`

NewManagedArtifactSyncArtifact instantiates a new ManagedArtifactSyncArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactSyncArtifactWithDefaults

`func NewManagedArtifactSyncArtifactWithDefaults() *ManagedArtifactSyncArtifact`

NewManagedArtifactSyncArtifactWithDefaults instantiates a new ManagedArtifactSyncArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenityName

`func (o *ManagedArtifactSyncArtifact) GetAmenityName() string`

GetAmenityName returns the AmenityName field if non-nil, zero value otherwise.

### GetAmenityNameOk

`func (o *ManagedArtifactSyncArtifact) GetAmenityNameOk() (*string, bool)`

GetAmenityNameOk returns a tuple with the AmenityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityName

`func (o *ManagedArtifactSyncArtifact) SetAmenityName(v string)`

SetAmenityName sets AmenityName field to given value.


### GetArtifactKey

`func (o *ManagedArtifactSyncArtifact) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *ManagedArtifactSyncArtifact) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *ManagedArtifactSyncArtifact) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.


### GetCompletedAt

`func (o *ManagedArtifactSyncArtifact) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ManagedArtifactSyncArtifact) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ManagedArtifactSyncArtifact) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ManagedArtifactSyncArtifact) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetDestinationDigest

`func (o *ManagedArtifactSyncArtifact) GetDestinationDigest() string`

GetDestinationDigest returns the DestinationDigest field if non-nil, zero value otherwise.

### GetDestinationDigestOk

`func (o *ManagedArtifactSyncArtifact) GetDestinationDigestOk() (*string, bool)`

GetDestinationDigestOk returns a tuple with the DestinationDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDigest

`func (o *ManagedArtifactSyncArtifact) SetDestinationDigest(v string)`

SetDestinationDigest sets DestinationDigest field to given value.

### HasDestinationDigest

`func (o *ManagedArtifactSyncArtifact) HasDestinationDigest() bool`

HasDestinationDigest returns a boolean if a field has been set.

### GetDestinationRef

`func (o *ManagedArtifactSyncArtifact) GetDestinationRef() string`

GetDestinationRef returns the DestinationRef field if non-nil, zero value otherwise.

### GetDestinationRefOk

`func (o *ManagedArtifactSyncArtifact) GetDestinationRefOk() (*string, bool)`

GetDestinationRefOk returns a tuple with the DestinationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRef

`func (o *ManagedArtifactSyncArtifact) SetDestinationRef(v string)`

SetDestinationRef sets DestinationRef field to given value.

### HasDestinationRef

`func (o *ManagedArtifactSyncArtifact) HasDestinationRef() bool`

HasDestinationRef returns a boolean if a field has been set.

### GetFailureCategory

`func (o *ManagedArtifactSyncArtifact) GetFailureCategory() string`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *ManagedArtifactSyncArtifact) GetFailureCategoryOk() (*string, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *ManagedArtifactSyncArtifact) SetFailureCategory(v string)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *ManagedArtifactSyncArtifact) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ManagedArtifactSyncArtifact) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ManagedArtifactSyncArtifact) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ManagedArtifactSyncArtifact) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ManagedArtifactSyncArtifact) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetSourceDigest

`func (o *ManagedArtifactSyncArtifact) GetSourceDigest() string`

GetSourceDigest returns the SourceDigest field if non-nil, zero value otherwise.

### GetSourceDigestOk

`func (o *ManagedArtifactSyncArtifact) GetSourceDigestOk() (*string, bool)`

GetSourceDigestOk returns a tuple with the SourceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDigest

`func (o *ManagedArtifactSyncArtifact) SetSourceDigest(v string)`

SetSourceDigest sets SourceDigest field to given value.

### HasSourceDigest

`func (o *ManagedArtifactSyncArtifact) HasSourceDigest() bool`

HasSourceDigest returns a boolean if a field has been set.

### GetSourceRef

`func (o *ManagedArtifactSyncArtifact) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ManagedArtifactSyncArtifact) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ManagedArtifactSyncArtifact) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *ManagedArtifactSyncArtifact) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedArtifactSyncArtifact) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedArtifactSyncArtifact) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedArtifactSyncArtifact) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *ManagedArtifactSyncArtifact) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedArtifactSyncArtifact) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedArtifactSyncArtifact) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


