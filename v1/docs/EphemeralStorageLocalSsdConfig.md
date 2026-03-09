# EphemeralStorageLocalSsdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCacheCount** | Pointer to **int64** | Number of local SSDs to use for data cache | [optional] 
**LocalSsdCount** | Pointer to **int64** | Number of local SSDs to use for ephemeral storage | [optional] 

## Methods

### NewEphemeralStorageLocalSsdConfig

`func NewEphemeralStorageLocalSsdConfig() *EphemeralStorageLocalSsdConfig`

NewEphemeralStorageLocalSsdConfig instantiates a new EphemeralStorageLocalSsdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEphemeralStorageLocalSsdConfigWithDefaults

`func NewEphemeralStorageLocalSsdConfigWithDefaults() *EphemeralStorageLocalSsdConfig`

NewEphemeralStorageLocalSsdConfigWithDefaults instantiates a new EphemeralStorageLocalSsdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCacheCount

`func (o *EphemeralStorageLocalSsdConfig) GetDataCacheCount() int64`

GetDataCacheCount returns the DataCacheCount field if non-nil, zero value otherwise.

### GetDataCacheCountOk

`func (o *EphemeralStorageLocalSsdConfig) GetDataCacheCountOk() (*int64, bool)`

GetDataCacheCountOk returns a tuple with the DataCacheCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheCount

`func (o *EphemeralStorageLocalSsdConfig) SetDataCacheCount(v int64)`

SetDataCacheCount sets DataCacheCount field to given value.

### HasDataCacheCount

`func (o *EphemeralStorageLocalSsdConfig) HasDataCacheCount() bool`

HasDataCacheCount returns a boolean if a field has been set.

### GetLocalSsdCount

`func (o *EphemeralStorageLocalSsdConfig) GetLocalSsdCount() int64`

GetLocalSsdCount returns the LocalSsdCount field if non-nil, zero value otherwise.

### GetLocalSsdCountOk

`func (o *EphemeralStorageLocalSsdConfig) GetLocalSsdCountOk() (*int64, bool)`

GetLocalSsdCountOk returns a tuple with the LocalSsdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSsdCount

`func (o *EphemeralStorageLocalSsdConfig) SetLocalSsdCount(v int64)`

SetLocalSsdCount sets LocalSsdCount field to given value.

### HasLocalSsdCount

`func (o *EphemeralStorageLocalSsdConfig) HasLocalSsdCount() bool`

HasLocalSsdCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


