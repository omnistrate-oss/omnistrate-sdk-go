# DeploymentCellAmenityArtifactSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmenityName** | **string** | The amenity name | 
**ArtifactKind** | **string** | The artifact kind | 
**ContentEncoding** | Pointer to **string** | The artifact payload encoding | [optional] 
**ContentType** | Pointer to **string** | The artifact content type | [optional] 
**Generation** | **int64** | The amenity generation for this artifact | 
**PayloadBase64** | Pointer to **string** | Base64-encoded uncompressed artifact payload for debug views | [optional] 
**SecretMasked** | Pointer to **bool** | Whether secrets have been masked in this artifact | [optional] 
**Sha256** | Pointer to **string** | The SHA256 hash of the uncompressed artifact payload | [optional] 
**SizeBytes** | Pointer to **int64** | The uncompressed artifact size | [optional] 

## Methods

### NewDeploymentCellAmenityArtifactSummary

`func NewDeploymentCellAmenityArtifactSummary(amenityName string, artifactKind string, generation int64, ) *DeploymentCellAmenityArtifactSummary`

NewDeploymentCellAmenityArtifactSummary instantiates a new DeploymentCellAmenityArtifactSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellAmenityArtifactSummaryWithDefaults

`func NewDeploymentCellAmenityArtifactSummaryWithDefaults() *DeploymentCellAmenityArtifactSummary`

NewDeploymentCellAmenityArtifactSummaryWithDefaults instantiates a new DeploymentCellAmenityArtifactSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenityName

`func (o *DeploymentCellAmenityArtifactSummary) GetAmenityName() string`

GetAmenityName returns the AmenityName field if non-nil, zero value otherwise.

### GetAmenityNameOk

`func (o *DeploymentCellAmenityArtifactSummary) GetAmenityNameOk() (*string, bool)`

GetAmenityNameOk returns a tuple with the AmenityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityName

`func (o *DeploymentCellAmenityArtifactSummary) SetAmenityName(v string)`

SetAmenityName sets AmenityName field to given value.


### GetArtifactKind

`func (o *DeploymentCellAmenityArtifactSummary) GetArtifactKind() string`

GetArtifactKind returns the ArtifactKind field if non-nil, zero value otherwise.

### GetArtifactKindOk

`func (o *DeploymentCellAmenityArtifactSummary) GetArtifactKindOk() (*string, bool)`

GetArtifactKindOk returns a tuple with the ArtifactKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKind

`func (o *DeploymentCellAmenityArtifactSummary) SetArtifactKind(v string)`

SetArtifactKind sets ArtifactKind field to given value.


### GetContentEncoding

`func (o *DeploymentCellAmenityArtifactSummary) GetContentEncoding() string`

GetContentEncoding returns the ContentEncoding field if non-nil, zero value otherwise.

### GetContentEncodingOk

`func (o *DeploymentCellAmenityArtifactSummary) GetContentEncodingOk() (*string, bool)`

GetContentEncodingOk returns a tuple with the ContentEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEncoding

`func (o *DeploymentCellAmenityArtifactSummary) SetContentEncoding(v string)`

SetContentEncoding sets ContentEncoding field to given value.

### HasContentEncoding

`func (o *DeploymentCellAmenityArtifactSummary) HasContentEncoding() bool`

HasContentEncoding returns a boolean if a field has been set.

### GetContentType

`func (o *DeploymentCellAmenityArtifactSummary) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DeploymentCellAmenityArtifactSummary) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DeploymentCellAmenityArtifactSummary) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DeploymentCellAmenityArtifactSummary) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetGeneration

`func (o *DeploymentCellAmenityArtifactSummary) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *DeploymentCellAmenityArtifactSummary) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *DeploymentCellAmenityArtifactSummary) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetPayloadBase64

`func (o *DeploymentCellAmenityArtifactSummary) GetPayloadBase64() string`

GetPayloadBase64 returns the PayloadBase64 field if non-nil, zero value otherwise.

### GetPayloadBase64Ok

`func (o *DeploymentCellAmenityArtifactSummary) GetPayloadBase64Ok() (*string, bool)`

GetPayloadBase64Ok returns a tuple with the PayloadBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadBase64

`func (o *DeploymentCellAmenityArtifactSummary) SetPayloadBase64(v string)`

SetPayloadBase64 sets PayloadBase64 field to given value.

### HasPayloadBase64

`func (o *DeploymentCellAmenityArtifactSummary) HasPayloadBase64() bool`

HasPayloadBase64 returns a boolean if a field has been set.

### GetSecretMasked

`func (o *DeploymentCellAmenityArtifactSummary) GetSecretMasked() bool`

GetSecretMasked returns the SecretMasked field if non-nil, zero value otherwise.

### GetSecretMaskedOk

`func (o *DeploymentCellAmenityArtifactSummary) GetSecretMaskedOk() (*bool, bool)`

GetSecretMaskedOk returns a tuple with the SecretMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretMasked

`func (o *DeploymentCellAmenityArtifactSummary) SetSecretMasked(v bool)`

SetSecretMasked sets SecretMasked field to given value.

### HasSecretMasked

`func (o *DeploymentCellAmenityArtifactSummary) HasSecretMasked() bool`

HasSecretMasked returns a boolean if a field has been set.

### GetSha256

`func (o *DeploymentCellAmenityArtifactSummary) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *DeploymentCellAmenityArtifactSummary) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *DeploymentCellAmenityArtifactSummary) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *DeploymentCellAmenityArtifactSummary) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSizeBytes

`func (o *DeploymentCellAmenityArtifactSummary) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *DeploymentCellAmenityArtifactSummary) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *DeploymentCellAmenityArtifactSummary) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *DeploymentCellAmenityArtifactSummary) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


