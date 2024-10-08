# FileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the file | 
**FileId** | **string** | The ID of the file | 
**MountPath** | **string** | The mount path of the file | 
**Name** | **string** | The name of the file | 
**Size** | **int64** | The size of the file | 
**Type** | **string** | The type of the file | 
**UploadTime** | **string** | The time the file was uploaded | 
**UploadedBy** | **string** | The user who uploaded the file | 

## Methods

### NewFileMetadata

`func NewFileMetadata(description string, fileId string, mountPath string, name string, size int64, type_ string, uploadTime string, uploadedBy string, ) *FileMetadata`

NewFileMetadata instantiates a new FileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataWithDefaults

`func NewFileMetadataWithDefaults() *FileMetadata`

NewFileMetadataWithDefaults instantiates a new FileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FileMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFileId

`func (o *FileMetadata) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileMetadata) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileMetadata) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetMountPath

`func (o *FileMetadata) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *FileMetadata) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *FileMetadata) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetName

`func (o *FileMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *FileMetadata) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileMetadata) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileMetadata) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *FileMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetUploadTime

`func (o *FileMetadata) GetUploadTime() string`

GetUploadTime returns the UploadTime field if non-nil, zero value otherwise.

### GetUploadTimeOk

`func (o *FileMetadata) GetUploadTimeOk() (*string, bool)`

GetUploadTimeOk returns a tuple with the UploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTime

`func (o *FileMetadata) SetUploadTime(v string)`

SetUploadTime sets UploadTime field to given value.


### GetUploadedBy

`func (o *FileMetadata) GetUploadedBy() string`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *FileMetadata) GetUploadedByOk() (*string, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *FileMetadata) SetUploadedBy(v string)`

SetUploadedBy sets UploadedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


