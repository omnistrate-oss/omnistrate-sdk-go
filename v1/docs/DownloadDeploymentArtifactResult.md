# DownloadDeploymentArtifactResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to **string** | Content-Disposition header for the artifact download. | [optional] 
**ContentType** | Pointer to **string** | Content-Type header for the artifact download. | [optional] [default to "application/gzip"]
**Length** | Pointer to **int64** | Length is the downloaded content length in bytes. | [optional] 

## Methods

### NewDownloadDeploymentArtifactResult

`func NewDownloadDeploymentArtifactResult() *DownloadDeploymentArtifactResult`

NewDownloadDeploymentArtifactResult instantiates a new DownloadDeploymentArtifactResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadDeploymentArtifactResultWithDefaults

`func NewDownloadDeploymentArtifactResultWithDefaults() *DownloadDeploymentArtifactResult`

NewDownloadDeploymentArtifactResultWithDefaults instantiates a new DownloadDeploymentArtifactResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *DownloadDeploymentArtifactResult) GetContentDisposition() string`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *DownloadDeploymentArtifactResult) GetContentDispositionOk() (*string, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *DownloadDeploymentArtifactResult) SetContentDisposition(v string)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *DownloadDeploymentArtifactResult) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetContentType

`func (o *DownloadDeploymentArtifactResult) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DownloadDeploymentArtifactResult) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DownloadDeploymentArtifactResult) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DownloadDeploymentArtifactResult) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetLength

`func (o *DownloadDeploymentArtifactResult) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *DownloadDeploymentArtifactResult) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *DownloadDeploymentArtifactResult) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *DownloadDeploymentArtifactResult) HasLength() bool`

HasLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


