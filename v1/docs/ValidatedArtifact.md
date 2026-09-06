# ValidatedArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressedSizeBytes** | **int64** | Length in bytes of the decoded compressed archive, as recomputed by the server | 
**LogicalPath** | **string** | The canonical local artifact root that was validated | 
**Sha256** | **string** | Lowercase hex SHA-256 of the decoded compressed archive bytes, as recomputed by the server | 
**Uses** | [**[]ValidationArtifactUse**](ValidationArtifactUse.md) | The uses whose validators consumed this content | 

## Methods

### NewValidatedArtifact

`func NewValidatedArtifact(compressedSizeBytes int64, logicalPath string, sha256 string, uses []ValidationArtifactUse, ) *ValidatedArtifact`

NewValidatedArtifact instantiates a new ValidatedArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedArtifactWithDefaults

`func NewValidatedArtifactWithDefaults() *ValidatedArtifact`

NewValidatedArtifactWithDefaults instantiates a new ValidatedArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressedSizeBytes

`func (o *ValidatedArtifact) GetCompressedSizeBytes() int64`

GetCompressedSizeBytes returns the CompressedSizeBytes field if non-nil, zero value otherwise.

### GetCompressedSizeBytesOk

`func (o *ValidatedArtifact) GetCompressedSizeBytesOk() (*int64, bool)`

GetCompressedSizeBytesOk returns a tuple with the CompressedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedSizeBytes

`func (o *ValidatedArtifact) SetCompressedSizeBytes(v int64)`

SetCompressedSizeBytes sets CompressedSizeBytes field to given value.


### GetLogicalPath

`func (o *ValidatedArtifact) GetLogicalPath() string`

GetLogicalPath returns the LogicalPath field if non-nil, zero value otherwise.

### GetLogicalPathOk

`func (o *ValidatedArtifact) GetLogicalPathOk() (*string, bool)`

GetLogicalPathOk returns a tuple with the LogicalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalPath

`func (o *ValidatedArtifact) SetLogicalPath(v string)`

SetLogicalPath sets LogicalPath field to given value.


### GetSha256

`func (o *ValidatedArtifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ValidatedArtifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ValidatedArtifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetUses

`func (o *ValidatedArtifact) GetUses() []ValidationArtifactUse`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ValidatedArtifact) GetUsesOk() (*[]ValidationArtifactUse, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ValidatedArtifact) SetUses(v []ValidationArtifactUse)`

SetUses sets Uses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


