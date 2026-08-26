# ManagedReleaseArtifact2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactKey** | **string** |  | 
**Consumers** | **[]string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**OwnerComponent** | **string** |  | 
**Platforms** | Pointer to **[]string** |  | [optional] 
**RelativePath** | **string** |  | 
**SourceChecksum** | Pointer to **string** |  | [optional] 
**SourceDigest** | Pointer to **string** |  | [optional] 
**SourceRef** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewManagedReleaseArtifact2

`func NewManagedReleaseArtifact2(artifactKey string, consumers []string, ownerComponent string, relativePath string, sourceRef string, type_ string, ) *ManagedReleaseArtifact2`

NewManagedReleaseArtifact2 instantiates a new ManagedReleaseArtifact2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseArtifact2WithDefaults

`func NewManagedReleaseArtifact2WithDefaults() *ManagedReleaseArtifact2`

NewManagedReleaseArtifact2WithDefaults instantiates a new ManagedReleaseArtifact2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactKey

`func (o *ManagedReleaseArtifact2) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *ManagedReleaseArtifact2) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *ManagedReleaseArtifact2) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.


### GetConsumers

`func (o *ManagedReleaseArtifact2) GetConsumers() []string`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ManagedReleaseArtifact2) GetConsumersOk() (*[]string, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ManagedReleaseArtifact2) SetConsumers(v []string)`

SetConsumers sets Consumers field to given value.


### GetName

`func (o *ManagedReleaseArtifact2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedReleaseArtifact2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedReleaseArtifact2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedReleaseArtifact2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerComponent

`func (o *ManagedReleaseArtifact2) GetOwnerComponent() string`

GetOwnerComponent returns the OwnerComponent field if non-nil, zero value otherwise.

### GetOwnerComponentOk

`func (o *ManagedReleaseArtifact2) GetOwnerComponentOk() (*string, bool)`

GetOwnerComponentOk returns a tuple with the OwnerComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerComponent

`func (o *ManagedReleaseArtifact2) SetOwnerComponent(v string)`

SetOwnerComponent sets OwnerComponent field to given value.


### GetPlatforms

`func (o *ManagedReleaseArtifact2) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ManagedReleaseArtifact2) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ManagedReleaseArtifact2) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ManagedReleaseArtifact2) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRelativePath

`func (o *ManagedReleaseArtifact2) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ManagedReleaseArtifact2) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ManagedReleaseArtifact2) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetSourceChecksum

`func (o *ManagedReleaseArtifact2) GetSourceChecksum() string`

GetSourceChecksum returns the SourceChecksum field if non-nil, zero value otherwise.

### GetSourceChecksumOk

`func (o *ManagedReleaseArtifact2) GetSourceChecksumOk() (*string, bool)`

GetSourceChecksumOk returns a tuple with the SourceChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChecksum

`func (o *ManagedReleaseArtifact2) SetSourceChecksum(v string)`

SetSourceChecksum sets SourceChecksum field to given value.

### HasSourceChecksum

`func (o *ManagedReleaseArtifact2) HasSourceChecksum() bool`

HasSourceChecksum returns a boolean if a field has been set.

### GetSourceDigest

`func (o *ManagedReleaseArtifact2) GetSourceDigest() string`

GetSourceDigest returns the SourceDigest field if non-nil, zero value otherwise.

### GetSourceDigestOk

`func (o *ManagedReleaseArtifact2) GetSourceDigestOk() (*string, bool)`

GetSourceDigestOk returns a tuple with the SourceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDigest

`func (o *ManagedReleaseArtifact2) SetSourceDigest(v string)`

SetSourceDigest sets SourceDigest field to given value.

### HasSourceDigest

`func (o *ManagedReleaseArtifact2) HasSourceDigest() bool`

HasSourceDigest returns a boolean if a field has been set.

### GetSourceRef

`func (o *ManagedReleaseArtifact2) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ManagedReleaseArtifact2) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ManagedReleaseArtifact2) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetType

`func (o *ManagedReleaseArtifact2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedReleaseArtifact2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedReleaseArtifact2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


