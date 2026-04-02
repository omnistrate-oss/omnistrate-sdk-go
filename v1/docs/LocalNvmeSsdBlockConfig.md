# LocalNvmeSsdBlockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSsdCount** | Pointer to **int64** | Number of raw-block Local NVMe SSDs | [optional] 

## Methods

### NewLocalNvmeSsdBlockConfig

`func NewLocalNvmeSsdBlockConfig() *LocalNvmeSsdBlockConfig`

NewLocalNvmeSsdBlockConfig instantiates a new LocalNvmeSsdBlockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalNvmeSsdBlockConfigWithDefaults

`func NewLocalNvmeSsdBlockConfigWithDefaults() *LocalNvmeSsdBlockConfig`

NewLocalNvmeSsdBlockConfigWithDefaults instantiates a new LocalNvmeSsdBlockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSsdCount

`func (o *LocalNvmeSsdBlockConfig) GetLocalSsdCount() int64`

GetLocalSsdCount returns the LocalSsdCount field if non-nil, zero value otherwise.

### GetLocalSsdCountOk

`func (o *LocalNvmeSsdBlockConfig) GetLocalSsdCountOk() (*int64, bool)`

GetLocalSsdCountOk returns a tuple with the LocalSsdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSsdCount

`func (o *LocalNvmeSsdBlockConfig) SetLocalSsdCount(v int64)`

SetLocalSsdCount sets LocalSsdCount field to given value.

### HasLocalSsdCount

`func (o *LocalNvmeSsdBlockConfig) HasLocalSsdCount() bool`

HasLocalSsdCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


