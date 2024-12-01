# BlobStorageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GCSConfiguration** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBlobStorageConfiguration

`func NewBlobStorageConfiguration() *BlobStorageConfiguration`

NewBlobStorageConfiguration instantiates a new BlobStorageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobStorageConfigurationWithDefaults

`func NewBlobStorageConfigurationWithDefaults() *BlobStorageConfiguration`

NewBlobStorageConfigurationWithDefaults instantiates a new BlobStorageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGCSConfiguration

`func (o *BlobStorageConfiguration) GetGCSConfiguration() map[string]interface{}`

GetGCSConfiguration returns the GCSConfiguration field if non-nil, zero value otherwise.

### GetGCSConfigurationOk

`func (o *BlobStorageConfiguration) GetGCSConfigurationOk() (*map[string]interface{}, bool)`

GetGCSConfigurationOk returns a tuple with the GCSConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGCSConfiguration

`func (o *BlobStorageConfiguration) SetGCSConfiguration(v map[string]interface{})`

SetGCSConfiguration sets GCSConfiguration field to given value.

### HasGCSConfiguration

`func (o *BlobStorageConfiguration) HasGCSConfiguration() bool`

HasGCSConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


