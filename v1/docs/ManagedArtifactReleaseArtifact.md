# ManagedArtifactReleaseArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmenityName** | **string** | Canonical Base Amenity name associated with this artifact. Clients must use this field instead of parsing artifactKey or name. | 
**ArtifactKey** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**RelativePath** | **string** |  | 
**SourceChecksum** | Pointer to **string** |  | [optional] 
**SourceDigest** | Pointer to **string** |  | [optional] 
**SourceRef** | **string** | Canonical source reference for the artifact. | 
**Type** | **string** |  | 
**Version** | **string** | Artifact version or image tag derived by the backend from the canonical source reference. | 

## Methods

### NewManagedArtifactReleaseArtifact

`func NewManagedArtifactReleaseArtifact(amenityName string, artifactKey string, relativePath string, sourceRef string, type_ string, version string, ) *ManagedArtifactReleaseArtifact`

NewManagedArtifactReleaseArtifact instantiates a new ManagedArtifactReleaseArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactReleaseArtifactWithDefaults

`func NewManagedArtifactReleaseArtifactWithDefaults() *ManagedArtifactReleaseArtifact`

NewManagedArtifactReleaseArtifactWithDefaults instantiates a new ManagedArtifactReleaseArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenityName

`func (o *ManagedArtifactReleaseArtifact) GetAmenityName() string`

GetAmenityName returns the AmenityName field if non-nil, zero value otherwise.

### GetAmenityNameOk

`func (o *ManagedArtifactReleaseArtifact) GetAmenityNameOk() (*string, bool)`

GetAmenityNameOk returns a tuple with the AmenityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityName

`func (o *ManagedArtifactReleaseArtifact) SetAmenityName(v string)`

SetAmenityName sets AmenityName field to given value.


### GetArtifactKey

`func (o *ManagedArtifactReleaseArtifact) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *ManagedArtifactReleaseArtifact) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *ManagedArtifactReleaseArtifact) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.


### GetName

`func (o *ManagedArtifactReleaseArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedArtifactReleaseArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedArtifactReleaseArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedArtifactReleaseArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatforms

`func (o *ManagedArtifactReleaseArtifact) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ManagedArtifactReleaseArtifact) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ManagedArtifactReleaseArtifact) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ManagedArtifactReleaseArtifact) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRelativePath

`func (o *ManagedArtifactReleaseArtifact) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ManagedArtifactReleaseArtifact) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ManagedArtifactReleaseArtifact) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetSourceChecksum

`func (o *ManagedArtifactReleaseArtifact) GetSourceChecksum() string`

GetSourceChecksum returns the SourceChecksum field if non-nil, zero value otherwise.

### GetSourceChecksumOk

`func (o *ManagedArtifactReleaseArtifact) GetSourceChecksumOk() (*string, bool)`

GetSourceChecksumOk returns a tuple with the SourceChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChecksum

`func (o *ManagedArtifactReleaseArtifact) SetSourceChecksum(v string)`

SetSourceChecksum sets SourceChecksum field to given value.

### HasSourceChecksum

`func (o *ManagedArtifactReleaseArtifact) HasSourceChecksum() bool`

HasSourceChecksum returns a boolean if a field has been set.

### GetSourceDigest

`func (o *ManagedArtifactReleaseArtifact) GetSourceDigest() string`

GetSourceDigest returns the SourceDigest field if non-nil, zero value otherwise.

### GetSourceDigestOk

`func (o *ManagedArtifactReleaseArtifact) GetSourceDigestOk() (*string, bool)`

GetSourceDigestOk returns a tuple with the SourceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDigest

`func (o *ManagedArtifactReleaseArtifact) SetSourceDigest(v string)`

SetSourceDigest sets SourceDigest field to given value.

### HasSourceDigest

`func (o *ManagedArtifactReleaseArtifact) HasSourceDigest() bool`

HasSourceDigest returns a boolean if a field has been set.

### GetSourceRef

`func (o *ManagedArtifactReleaseArtifact) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ManagedArtifactReleaseArtifact) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ManagedArtifactReleaseArtifact) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetType

`func (o *ManagedArtifactReleaseArtifact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedArtifactReleaseArtifact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedArtifactReleaseArtifact) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *ManagedArtifactReleaseArtifact) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedArtifactReleaseArtifact) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedArtifactReleaseArtifact) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


