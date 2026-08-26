# ManagedReleaseArtifact3

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

### NewManagedReleaseArtifact3

`func NewManagedReleaseArtifact3(artifactKey string, consumers []string, ownerComponent string, relativePath string, sourceRef string, type_ string, ) *ManagedReleaseArtifact3`

NewManagedReleaseArtifact3 instantiates a new ManagedReleaseArtifact3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseArtifact3WithDefaults

`func NewManagedReleaseArtifact3WithDefaults() *ManagedReleaseArtifact3`

NewManagedReleaseArtifact3WithDefaults instantiates a new ManagedReleaseArtifact3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactKey

`func (o *ManagedReleaseArtifact3) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *ManagedReleaseArtifact3) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *ManagedReleaseArtifact3) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.


### GetConsumers

`func (o *ManagedReleaseArtifact3) GetConsumers() []string`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ManagedReleaseArtifact3) GetConsumersOk() (*[]string, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ManagedReleaseArtifact3) SetConsumers(v []string)`

SetConsumers sets Consumers field to given value.


### GetName

`func (o *ManagedReleaseArtifact3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedReleaseArtifact3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedReleaseArtifact3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedReleaseArtifact3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerComponent

`func (o *ManagedReleaseArtifact3) GetOwnerComponent() string`

GetOwnerComponent returns the OwnerComponent field if non-nil, zero value otherwise.

### GetOwnerComponentOk

`func (o *ManagedReleaseArtifact3) GetOwnerComponentOk() (*string, bool)`

GetOwnerComponentOk returns a tuple with the OwnerComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerComponent

`func (o *ManagedReleaseArtifact3) SetOwnerComponent(v string)`

SetOwnerComponent sets OwnerComponent field to given value.


### GetPlatforms

`func (o *ManagedReleaseArtifact3) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ManagedReleaseArtifact3) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ManagedReleaseArtifact3) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ManagedReleaseArtifact3) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRelativePath

`func (o *ManagedReleaseArtifact3) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ManagedReleaseArtifact3) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ManagedReleaseArtifact3) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetSourceChecksum

`func (o *ManagedReleaseArtifact3) GetSourceChecksum() string`

GetSourceChecksum returns the SourceChecksum field if non-nil, zero value otherwise.

### GetSourceChecksumOk

`func (o *ManagedReleaseArtifact3) GetSourceChecksumOk() (*string, bool)`

GetSourceChecksumOk returns a tuple with the SourceChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChecksum

`func (o *ManagedReleaseArtifact3) SetSourceChecksum(v string)`

SetSourceChecksum sets SourceChecksum field to given value.

### HasSourceChecksum

`func (o *ManagedReleaseArtifact3) HasSourceChecksum() bool`

HasSourceChecksum returns a boolean if a field has been set.

### GetSourceDigest

`func (o *ManagedReleaseArtifact3) GetSourceDigest() string`

GetSourceDigest returns the SourceDigest field if non-nil, zero value otherwise.

### GetSourceDigestOk

`func (o *ManagedReleaseArtifact3) GetSourceDigestOk() (*string, bool)`

GetSourceDigestOk returns a tuple with the SourceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDigest

`func (o *ManagedReleaseArtifact3) SetSourceDigest(v string)`

SetSourceDigest sets SourceDigest field to given value.

### HasSourceDigest

`func (o *ManagedReleaseArtifact3) HasSourceDigest() bool`

HasSourceDigest returns a boolean if a field has been set.

### GetSourceRef

`func (o *ManagedReleaseArtifact3) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *ManagedReleaseArtifact3) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *ManagedReleaseArtifact3) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetType

`func (o *ManagedReleaseArtifact3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedReleaseArtifact3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedReleaseArtifact3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


