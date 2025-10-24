# GCPFilestoreConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **string** | The capacity of the GCP filestore in GiB | 
**MaxIopsPerTb** | Pointer to **int64** | The maximum IOPS per TB for the performance tier | [optional] 
**Tier** | **string** | The tier of the GCP filestore | 

## Methods

### NewGCPFilestoreConfiguration

`func NewGCPFilestoreConfiguration(capacity string, tier string, ) *GCPFilestoreConfiguration`

NewGCPFilestoreConfiguration instantiates a new GCPFilestoreConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPFilestoreConfigurationWithDefaults

`func NewGCPFilestoreConfigurationWithDefaults() *GCPFilestoreConfiguration`

NewGCPFilestoreConfigurationWithDefaults instantiates a new GCPFilestoreConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *GCPFilestoreConfiguration) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *GCPFilestoreConfiguration) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *GCPFilestoreConfiguration) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.


### GetMaxIopsPerTb

`func (o *GCPFilestoreConfiguration) GetMaxIopsPerTb() int64`

GetMaxIopsPerTb returns the MaxIopsPerTb field if non-nil, zero value otherwise.

### GetMaxIopsPerTbOk

`func (o *GCPFilestoreConfiguration) GetMaxIopsPerTbOk() (*int64, bool)`

GetMaxIopsPerTbOk returns a tuple with the MaxIopsPerTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIopsPerTb

`func (o *GCPFilestoreConfiguration) SetMaxIopsPerTb(v int64)`

SetMaxIopsPerTb sets MaxIopsPerTb field to given value.

### HasMaxIopsPerTb

`func (o *GCPFilestoreConfiguration) HasMaxIopsPerTb() bool`

HasMaxIopsPerTb returns a boolean if a field has been set.

### GetTier

`func (o *GCPFilestoreConfiguration) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *GCPFilestoreConfiguration) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *GCPFilestoreConfiguration) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


