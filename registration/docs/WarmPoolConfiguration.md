# WarmPoolConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumNodesInPool** | Pointer to **int64** | Minimum number of compute nodes in pool | [optional] [default to 1]
**MinimumPercentageNodesInPool** | Pointer to **int64** | Minimum percentage of compute nodes in pool | [optional] [default to 0]

## Methods

### NewWarmPoolConfiguration

`func NewWarmPoolConfiguration() *WarmPoolConfiguration`

NewWarmPoolConfiguration instantiates a new WarmPoolConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarmPoolConfigurationWithDefaults

`func NewWarmPoolConfigurationWithDefaults() *WarmPoolConfiguration`

NewWarmPoolConfigurationWithDefaults instantiates a new WarmPoolConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumNodesInPool

`func (o *WarmPoolConfiguration) GetMinimumNodesInPool() int64`

GetMinimumNodesInPool returns the MinimumNodesInPool field if non-nil, zero value otherwise.

### GetMinimumNodesInPoolOk

`func (o *WarmPoolConfiguration) GetMinimumNodesInPoolOk() (*int64, bool)`

GetMinimumNodesInPoolOk returns a tuple with the MinimumNodesInPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodesInPool

`func (o *WarmPoolConfiguration) SetMinimumNodesInPool(v int64)`

SetMinimumNodesInPool sets MinimumNodesInPool field to given value.

### HasMinimumNodesInPool

`func (o *WarmPoolConfiguration) HasMinimumNodesInPool() bool`

HasMinimumNodesInPool returns a boolean if a field has been set.

### GetMinimumPercentageNodesInPool

`func (o *WarmPoolConfiguration) GetMinimumPercentageNodesInPool() int64`

GetMinimumPercentageNodesInPool returns the MinimumPercentageNodesInPool field if non-nil, zero value otherwise.

### GetMinimumPercentageNodesInPoolOk

`func (o *WarmPoolConfiguration) GetMinimumPercentageNodesInPoolOk() (*int64, bool)`

GetMinimumPercentageNodesInPoolOk returns a tuple with the MinimumPercentageNodesInPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPercentageNodesInPool

`func (o *WarmPoolConfiguration) SetMinimumPercentageNodesInPool(v int64)`

SetMinimumPercentageNodesInPool sets MinimumPercentageNodesInPool field to given value.

### HasMinimumPercentageNodesInPool

`func (o *WarmPoolConfiguration) HasMinimumPercentageNodesInPool() bool`

HasMinimumPercentageNodesInPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


