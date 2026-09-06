# ValidationArtifactInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveContent** | **string** | Base64 encoded gzip-compressed tar archive holding the exact local bytes to validate. Never logged, echoed back or persisted. | 
**CompressedSizeBytes** | **int64** | Length in bytes of the decoded compressed archive. Recomputed and verified by the server. | 
**Encoding** | **string** | The encoding of archiveContent. Exactly one encoding is permitted: a gzip-compressed tar archive, base64 encoded. | 
**LogicalPath** | **string** | The canonical local artifact root this archive provides, as reported in requiredArtifacts[].logicalPath. The workspace root is the canonical path &#39;.&#39;. This is an identifier on the wire, never a path the server creates on disk. | 
**Sha256** | **string** | Lowercase hex SHA-256 of the decoded compressed archive bytes, not of the base64 text and not of a mutable source directory. The server recomputes the digest and the size and rejects the request when either disagrees; neither supplied value is trusted. | 

## Methods

### NewValidationArtifactInput

`func NewValidationArtifactInput(archiveContent string, compressedSizeBytes int64, encoding string, logicalPath string, sha256 string, ) *ValidationArtifactInput`

NewValidationArtifactInput instantiates a new ValidationArtifactInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationArtifactInputWithDefaults

`func NewValidationArtifactInputWithDefaults() *ValidationArtifactInput`

NewValidationArtifactInputWithDefaults instantiates a new ValidationArtifactInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveContent

`func (o *ValidationArtifactInput) GetArchiveContent() string`

GetArchiveContent returns the ArchiveContent field if non-nil, zero value otherwise.

### GetArchiveContentOk

`func (o *ValidationArtifactInput) GetArchiveContentOk() (*string, bool)`

GetArchiveContentOk returns a tuple with the ArchiveContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveContent

`func (o *ValidationArtifactInput) SetArchiveContent(v string)`

SetArchiveContent sets ArchiveContent field to given value.


### GetCompressedSizeBytes

`func (o *ValidationArtifactInput) GetCompressedSizeBytes() int64`

GetCompressedSizeBytes returns the CompressedSizeBytes field if non-nil, zero value otherwise.

### GetCompressedSizeBytesOk

`func (o *ValidationArtifactInput) GetCompressedSizeBytesOk() (*int64, bool)`

GetCompressedSizeBytesOk returns a tuple with the CompressedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedSizeBytes

`func (o *ValidationArtifactInput) SetCompressedSizeBytes(v int64)`

SetCompressedSizeBytes sets CompressedSizeBytes field to given value.


### GetEncoding

`func (o *ValidationArtifactInput) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ValidationArtifactInput) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ValidationArtifactInput) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetLogicalPath

`func (o *ValidationArtifactInput) GetLogicalPath() string`

GetLogicalPath returns the LogicalPath field if non-nil, zero value otherwise.

### GetLogicalPathOk

`func (o *ValidationArtifactInput) GetLogicalPathOk() (*string, bool)`

GetLogicalPathOk returns a tuple with the LogicalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalPath

`func (o *ValidationArtifactInput) SetLogicalPath(v string)`

SetLogicalPath sets LogicalPath field to given value.


### GetSha256

`func (o *ValidationArtifactInput) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ValidationArtifactInput) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ValidationArtifactInput) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


