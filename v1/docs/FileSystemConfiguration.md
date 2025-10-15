# FileSystemConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EFSFileSystemConfiguration** | Pointer to [**EFSFileSystemConfiguration**](EFSFileSystemConfiguration.md) |  | [optional] 
**GCPFilestoreConfiguration** | Pointer to [**GCPFilestoreConfiguration**](GCPFilestoreConfiguration.md) |  | [optional] 

## Methods

### NewFileSystemConfiguration

`func NewFileSystemConfiguration() *FileSystemConfiguration`

NewFileSystemConfiguration instantiates a new FileSystemConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemConfigurationWithDefaults

`func NewFileSystemConfigurationWithDefaults() *FileSystemConfiguration`

NewFileSystemConfigurationWithDefaults instantiates a new FileSystemConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEFSFileSystemConfiguration

`func (o *FileSystemConfiguration) GetEFSFileSystemConfiguration() EFSFileSystemConfiguration`

GetEFSFileSystemConfiguration returns the EFSFileSystemConfiguration field if non-nil, zero value otherwise.

### GetEFSFileSystemConfigurationOk

`func (o *FileSystemConfiguration) GetEFSFileSystemConfigurationOk() (*EFSFileSystemConfiguration, bool)`

GetEFSFileSystemConfigurationOk returns a tuple with the EFSFileSystemConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEFSFileSystemConfiguration

`func (o *FileSystemConfiguration) SetEFSFileSystemConfiguration(v EFSFileSystemConfiguration)`

SetEFSFileSystemConfiguration sets EFSFileSystemConfiguration field to given value.

### HasEFSFileSystemConfiguration

`func (o *FileSystemConfiguration) HasEFSFileSystemConfiguration() bool`

HasEFSFileSystemConfiguration returns a boolean if a field has been set.

### GetGCPFilestoreConfiguration

`func (o *FileSystemConfiguration) GetGCPFilestoreConfiguration() GCPFilestoreConfiguration`

GetGCPFilestoreConfiguration returns the GCPFilestoreConfiguration field if non-nil, zero value otherwise.

### GetGCPFilestoreConfigurationOk

`func (o *FileSystemConfiguration) GetGCPFilestoreConfigurationOk() (*GCPFilestoreConfiguration, bool)`

GetGCPFilestoreConfigurationOk returns a tuple with the GCPFilestoreConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGCPFilestoreConfiguration

`func (o *FileSystemConfiguration) SetGCPFilestoreConfiguration(v GCPFilestoreConfiguration)`

SetGCPFilestoreConfiguration sets GCPFilestoreConfiguration field to given value.

### HasGCPFilestoreConfiguration

`func (o *FileSystemConfiguration) HasGCPFilestoreConfiguration() bool`

HasGCPFilestoreConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


