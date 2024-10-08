# ListFilesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]FileMetadata**](FileMetadata.md) | List of files | 

## Methods

### NewListFilesResult

`func NewListFilesResult(files []FileMetadata, ) *ListFilesResult`

NewListFilesResult instantiates a new ListFilesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFilesResultWithDefaults

`func NewListFilesResultWithDefaults() *ListFilesResult`

NewListFilesResultWithDefaults instantiates a new ListFilesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ListFilesResult) GetFiles() []FileMetadata`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ListFilesResult) GetFilesOk() (*[]FileMetadata, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ListFilesResult) SetFiles(v []FileMetadata)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


