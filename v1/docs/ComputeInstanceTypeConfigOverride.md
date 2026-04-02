# ComputeInstanceTypeConfigOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceleratorConfiguration** | Pointer to [**AcceleratorConfiguration**](AcceleratorConfiguration.md) |  | [optional] 
**EphemeralStorageLocalSsdConfig** | Pointer to [**EphemeralStorageLocalSsdConfig**](EphemeralStorageLocalSsdConfig.md) |  | [optional] 
**GpuClusterID** | Pointer to **string** | GPU cluster attachment identifier for providers that support explicit GPU cluster placement (for example Nebius) | [optional] 
**InstanceLifeCycleType** | Pointer to **string** | The instance life cycle type for this compute instance type config | [optional] 
**Labels** | Pointer to **map[string]string** | Labels for the compute instance type config | [optional] 
**LocalNvmeSsdBlockConfig** | Pointer to [**LocalNvmeSsdBlockConfig**](LocalNvmeSsdBlockConfig.md) |  | [optional] 
**OsFamily** | Pointer to **string** | Operating system family for AWS instance type override (supported only for AWS) | [optional] 
**Platform** | Pointer to **string** | First-class platform selector for cloud providers that require more than a single instance type string (for example Nebius) | [optional] 
**RootVolumeSizeGi** | Pointer to **int64** | Size of the root volume in Gi | [optional] 
**RootVolumeSizeGiAPIParam** | Pointer to **string** | Size of the root volume in Gi as a string | [optional] 
**Taints** | Pointer to [**[]TaintConfiguration**](TaintConfiguration.md) | Taints for the compute instance type config | [optional] 
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

### GetAcceleratorConfiguration

`func (o *ComputeInstanceTypeConfigOverride) GetAcceleratorConfiguration() AcceleratorConfiguration`

GetAcceleratorConfiguration returns the AcceleratorConfiguration field if non-nil, zero value otherwise.

### GetAcceleratorConfigurationOk

`func (o *ComputeInstanceTypeConfigOverride) GetAcceleratorConfigurationOk() (*AcceleratorConfiguration, bool)`

GetAcceleratorConfigurationOk returns a tuple with the AcceleratorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorConfiguration

`func (o *ComputeInstanceTypeConfigOverride) SetAcceleratorConfiguration(v AcceleratorConfiguration)`

SetAcceleratorConfiguration sets AcceleratorConfiguration field to given value.

### HasAcceleratorConfiguration

`func (o *ComputeInstanceTypeConfigOverride) HasAcceleratorConfiguration() bool`

HasAcceleratorConfiguration returns a boolean if a field has been set.

### GetEphemeralStorageLocalSsdConfig

`func (o *ComputeInstanceTypeConfigOverride) GetEphemeralStorageLocalSsdConfig() EphemeralStorageLocalSsdConfig`

GetEphemeralStorageLocalSsdConfig returns the EphemeralStorageLocalSsdConfig field if non-nil, zero value otherwise.

### GetEphemeralStorageLocalSsdConfigOk

`func (o *ComputeInstanceTypeConfigOverride) GetEphemeralStorageLocalSsdConfigOk() (*EphemeralStorageLocalSsdConfig, bool)`

GetEphemeralStorageLocalSsdConfigOk returns a tuple with the EphemeralStorageLocalSsdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageLocalSsdConfig

`func (o *ComputeInstanceTypeConfigOverride) SetEphemeralStorageLocalSsdConfig(v EphemeralStorageLocalSsdConfig)`

SetEphemeralStorageLocalSsdConfig sets EphemeralStorageLocalSsdConfig field to given value.

### HasEphemeralStorageLocalSsdConfig

`func (o *ComputeInstanceTypeConfigOverride) HasEphemeralStorageLocalSsdConfig() bool`

HasEphemeralStorageLocalSsdConfig returns a boolean if a field has been set.

### GetGpuClusterID

`func (o *ComputeInstanceTypeConfigOverride) GetGpuClusterID() string`

GetGpuClusterID returns the GpuClusterID field if non-nil, zero value otherwise.

### GetGpuClusterIDOk

`func (o *ComputeInstanceTypeConfigOverride) GetGpuClusterIDOk() (*string, bool)`

GetGpuClusterIDOk returns a tuple with the GpuClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClusterID

`func (o *ComputeInstanceTypeConfigOverride) SetGpuClusterID(v string)`

SetGpuClusterID sets GpuClusterID field to given value.

### HasGpuClusterID

`func (o *ComputeInstanceTypeConfigOverride) HasGpuClusterID() bool`

HasGpuClusterID returns a boolean if a field has been set.

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

### GetLabels

`func (o *ComputeInstanceTypeConfigOverride) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ComputeInstanceTypeConfigOverride) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ComputeInstanceTypeConfigOverride) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ComputeInstanceTypeConfigOverride) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocalNvmeSsdBlockConfig

`func (o *ComputeInstanceTypeConfigOverride) GetLocalNvmeSsdBlockConfig() LocalNvmeSsdBlockConfig`

GetLocalNvmeSsdBlockConfig returns the LocalNvmeSsdBlockConfig field if non-nil, zero value otherwise.

### GetLocalNvmeSsdBlockConfigOk

`func (o *ComputeInstanceTypeConfigOverride) GetLocalNvmeSsdBlockConfigOk() (*LocalNvmeSsdBlockConfig, bool)`

GetLocalNvmeSsdBlockConfigOk returns a tuple with the LocalNvmeSsdBlockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNvmeSsdBlockConfig

`func (o *ComputeInstanceTypeConfigOverride) SetLocalNvmeSsdBlockConfig(v LocalNvmeSsdBlockConfig)`

SetLocalNvmeSsdBlockConfig sets LocalNvmeSsdBlockConfig field to given value.

### HasLocalNvmeSsdBlockConfig

`func (o *ComputeInstanceTypeConfigOverride) HasLocalNvmeSsdBlockConfig() bool`

HasLocalNvmeSsdBlockConfig returns a boolean if a field has been set.

### GetOsFamily

`func (o *ComputeInstanceTypeConfigOverride) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ComputeInstanceTypeConfigOverride) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ComputeInstanceTypeConfigOverride) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *ComputeInstanceTypeConfigOverride) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetPlatform

`func (o *ComputeInstanceTypeConfigOverride) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ComputeInstanceTypeConfigOverride) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ComputeInstanceTypeConfigOverride) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ComputeInstanceTypeConfigOverride) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

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

### GetRootVolumeSizeGiAPIParam

`func (o *ComputeInstanceTypeConfigOverride) GetRootVolumeSizeGiAPIParam() string`

GetRootVolumeSizeGiAPIParam returns the RootVolumeSizeGiAPIParam field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiAPIParamOk

`func (o *ComputeInstanceTypeConfigOverride) GetRootVolumeSizeGiAPIParamOk() (*string, bool)`

GetRootVolumeSizeGiAPIParamOk returns a tuple with the RootVolumeSizeGiAPIParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGiAPIParam

`func (o *ComputeInstanceTypeConfigOverride) SetRootVolumeSizeGiAPIParam(v string)`

SetRootVolumeSizeGiAPIParam sets RootVolumeSizeGiAPIParam field to given value.

### HasRootVolumeSizeGiAPIParam

`func (o *ComputeInstanceTypeConfigOverride) HasRootVolumeSizeGiAPIParam() bool`

HasRootVolumeSizeGiAPIParam returns a boolean if a field has been set.

### GetTaints

`func (o *ComputeInstanceTypeConfigOverride) GetTaints() []TaintConfiguration`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *ComputeInstanceTypeConfigOverride) GetTaintsOk() (*[]TaintConfiguration, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *ComputeInstanceTypeConfigOverride) SetTaints(v []TaintConfiguration)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *ComputeInstanceTypeConfigOverride) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

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


