# ComputeInstanceTypeConfigOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceLifeCycleType** | Pointer to **string** | The instance life cycle type for this compute instance type config | [optional] 
**RootVolumeSizeGi** | Pointer to **int64** | Size of the root volume in Gi | [optional] 
**WarmPoolConfiguration** | Pointer to [**WarmPoolConfiguration**](WarmPoolConfiguration.md) |  | [optional] 

## Methods

### NewComputeInstanceTypeConfigOverride

`func NewComputeInstanceTypeConfigOverride() *ComputeInstanceTypeConfigOverride`

NewComputeInstanceTypeConfigOverride instantiates a new ComputeInstanceTypeConfigOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeInstanceTypeConfigOverrideWithDefaults

`func NewComputeInstanceTypeConfigOverrideWithDefaults() *ComputeInstanceTypeConfigOverride`

NewComputeInstanceTypeConfigOverrideWithDefaults instantiates a new ComputeInstanceTypeConfigOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceLifeCycleType

`func (o *ComputeInstanceTypeConfigOverride) GetInstanceLifeCycleType() string`

GetInstanceLifeCycleType returns the InstanceLifeCycleType field if non-nil, zero value otherwise.

### GetInstanceLifeCycleTypeOk

`func (o *ComputeInstanceTypeConfigOverride) GetInstanceLifeCycleTypeOk() (*string, bool)`

GetInstanceLifeCycleTypeOk returns a tuple with the InstanceLifeCycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLifeCycleType

`func (o *ComputeInstanceTypeConfigOverride) SetInstanceLifeCycleType(v string)`

SetInstanceLifeCycleType sets InstanceLifeCycleType field to given value.

### HasInstanceLifeCycleType

`func (o *ComputeInstanceTypeConfigOverride) HasInstanceLifeCycleType() bool`

HasInstanceLifeCycleType returns a boolean if a field has been set.

### GetRootVolumeSizeGi

`func (o *ComputeInstanceTypeConfigOverride) GetRootVolumeSizeGi() int64`

GetRootVolumeSizeGi returns the RootVolumeSizeGi field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiOk

`func (o *ComputeInstanceTypeConfigOverride) GetRootVolumeSizeGiOk() (*int64, bool)`

GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGi

`func (o *ComputeInstanceTypeConfigOverride) SetRootVolumeSizeGi(v int64)`

SetRootVolumeSizeGi sets RootVolumeSizeGi field to given value.

### HasRootVolumeSizeGi

`func (o *ComputeInstanceTypeConfigOverride) HasRootVolumeSizeGi() bool`

HasRootVolumeSizeGi returns a boolean if a field has been set.

### GetWarmPoolConfiguration

`func (o *ComputeInstanceTypeConfigOverride) GetWarmPoolConfiguration() WarmPoolConfiguration`

GetWarmPoolConfiguration returns the WarmPoolConfiguration field if non-nil, zero value otherwise.

### GetWarmPoolConfigurationOk

`func (o *ComputeInstanceTypeConfigOverride) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool)`

GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolConfiguration

`func (o *ComputeInstanceTypeConfigOverride) SetWarmPoolConfiguration(v WarmPoolConfiguration)`

SetWarmPoolConfiguration sets WarmPoolConfiguration field to given value.

### HasWarmPoolConfiguration

`func (o *ComputeInstanceTypeConfigOverride) HasWarmPoolConfiguration() bool`

HasWarmPoolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


