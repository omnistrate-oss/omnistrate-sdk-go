# WorkspaceArtifactsDownloadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to **string** | Content-Disposition header for the workspace artifact archive download. | [optional] 
**ContentType** | Pointer to **string** | Content-Type header for the workspace artifact archive download. | [optional] [default to "application/gzip"]
**Length** | Pointer to **int64** | Length is the downloaded content length in bytes. | [optional] 

## Methods

### NewWorkspaceArtifactsDownloadResult

`func NewWorkspaceArtifactsDownloadResult() *WorkspaceArtifactsDownloadResult`

NewWorkspaceArtifactsDownloadResult instantiates a new WorkspaceArtifactsDownloadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceArtifactsDownloadResultWithDefaults

`func NewWorkspaceArtifactsDownloadResultWithDefaults() *WorkspaceArtifactsDownloadResult`

NewWorkspaceArtifactsDownloadResultWithDefaults instantiates a new WorkspaceArtifactsDownloadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *WorkspaceArtifactsDownloadResult) GetContentDisposition() string`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *WorkspaceArtifactsDownloadResult) GetContentDispositionOk() (*string, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *WorkspaceArtifactsDownloadResult) SetContentDisposition(v string)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *WorkspaceArtifactsDownloadResult) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetContentType

`func (o *WorkspaceArtifactsDownloadResult) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkspaceArtifactsDownloadResult) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkspaceArtifactsDownloadResult) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WorkspaceArtifactsDownloadResult) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetLength

`func (o *WorkspaceArtifactsDownloadResult) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WorkspaceArtifactsDownloadResult) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WorkspaceArtifactsDownloadResult) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *WorkspaceArtifactsDownloadResult) HasLength() bool`

HasLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


