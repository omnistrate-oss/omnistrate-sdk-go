# NebiusFileSystemConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemType** | **string** | The Nebius filesystem type | 
**SizeGibibytes** | **int64** | The size of the Nebius filesystem in GiB | 

## Methods

### NewNebiusFileSystemConfiguration

`func NewNebiusFileSystemConfiguration(filesystemType string, sizeGibibytes int64, ) *NebiusFileSystemConfiguration`

NewNebiusFileSystemConfiguration instantiates a new NebiusFileSystemConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNebiusFileSystemConfigurationWithDefaults

`func NewNebiusFileSystemConfigurationWithDefaults() *NebiusFileSystemConfiguration`

NewNebiusFileSystemConfigurationWithDefaults instantiates a new NebiusFileSystemConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemType

`func (o *NebiusFileSystemConfiguration) GetFilesystemType() string`

GetFilesystemType returns the FilesystemType field if non-nil, zero value otherwise.

### GetFilesystemTypeOk

`func (o *NebiusFileSystemConfiguration) GetFilesystemTypeOk() (*string, bool)`

GetFilesystemTypeOk returns a tuple with the FilesystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemType

`func (o *NebiusFileSystemConfiguration) SetFilesystemType(v string)`

SetFilesystemType sets FilesystemType field to given value.


### GetSizeGibibytes

`func (o *NebiusFileSystemConfiguration) GetSizeGibibytes() int64`

GetSizeGibibytes returns the SizeGibibytes field if non-nil, zero value otherwise.

### GetSizeGibibytesOk

`func (o *NebiusFileSystemConfiguration) GetSizeGibibytesOk() (*int64, bool)`

GetSizeGibibytesOk returns a tuple with the SizeGibibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGibibytes

`func (o *NebiusFileSystemConfiguration) SetSizeGibibytes(v int64)`

SetSizeGibibytes sets SizeGibibytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


