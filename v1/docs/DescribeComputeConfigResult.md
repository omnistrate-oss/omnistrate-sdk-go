# DescribeComputeConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalingPolicy** | Pointer to [**AutoscalingPolicy**](AutoscalingPolicy.md) |  | [optional] 
**ComputeInstanceTypeConfigOverrides** | Pointer to [**map[string]map[string]ComputeInstanceTypeConfigOverride**](map.md) | The compute instance type config overrides for this compute config | [optional] 
**CpuArchitecture** | Pointer to **string** | Processor architecture | [optional] 
**Description** | **string** | Description of the compute config | 
**Id** | **string** | ID of a Compute Config | 
**InfraConfigIDs** | Pointer to **[]string** | The list of infra config IDs associated with the compute config. | [optional] 
**InstanceTypes** | Pointer to **map[string][]string** | The instance types for this compute config | [optional] 
**Name** | **string** | Name of the compute config | 
**ReplicaCount** | **string** | Number of replicas to provision for this logical pool of nodes per instance of the resource | 
**Resources** | Pointer to [**ResourceSpec**](ResourceSpec.md) |  | [optional] 
**RootVolumeSizeGi** | Pointer to **int64** | Size of the root volume in Gi | [optional] 
**ServiceId** | **string** | ID of a Service | 
**WarmPoolConfiguration** | Pointer to [**WarmPoolConfiguration**](WarmPoolConfiguration.md) |  | [optional] 

## Methods

### NewDescribeComputeConfigResult

`func NewDescribeComputeConfigResult(description string, id string, name string, replicaCount string, serviceId string, ) *DescribeComputeConfigResult`

NewDescribeComputeConfigResult instantiates a new DescribeComputeConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeComputeConfigResultWithDefaults

`func NewDescribeComputeConfigResultWithDefaults() *DescribeComputeConfigResult`

NewDescribeComputeConfigResultWithDefaults instantiates a new DescribeComputeConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalingPolicy

`func (o *DescribeComputeConfigResult) GetAutoscalingPolicy() AutoscalingPolicy`

GetAutoscalingPolicy returns the AutoscalingPolicy field if non-nil, zero value otherwise.

### GetAutoscalingPolicyOk

`func (o *DescribeComputeConfigResult) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool)`

GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingPolicy

`func (o *DescribeComputeConfigResult) SetAutoscalingPolicy(v AutoscalingPolicy)`

SetAutoscalingPolicy sets AutoscalingPolicy field to given value.

### HasAutoscalingPolicy

`func (o *DescribeComputeConfigResult) HasAutoscalingPolicy() bool`

HasAutoscalingPolicy returns a boolean if a field has been set.

### GetComputeInstanceTypeConfigOverrides

`func (o *DescribeComputeConfigResult) GetComputeInstanceTypeConfigOverrides() map[string]map[string]ComputeInstanceTypeConfigOverride`

GetComputeInstanceTypeConfigOverrides returns the ComputeInstanceTypeConfigOverrides field if non-nil, zero value otherwise.

### GetComputeInstanceTypeConfigOverridesOk

`func (o *DescribeComputeConfigResult) GetComputeInstanceTypeConfigOverridesOk() (*map[string]map[string]ComputeInstanceTypeConfigOverride, bool)`

GetComputeInstanceTypeConfigOverridesOk returns a tuple with the ComputeInstanceTypeConfigOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInstanceTypeConfigOverrides

`func (o *DescribeComputeConfigResult) SetComputeInstanceTypeConfigOverrides(v map[string]map[string]ComputeInstanceTypeConfigOverride)`

SetComputeInstanceTypeConfigOverrides sets ComputeInstanceTypeConfigOverrides field to given value.

### HasComputeInstanceTypeConfigOverrides

`func (o *DescribeComputeConfigResult) HasComputeInstanceTypeConfigOverrides() bool`

HasComputeInstanceTypeConfigOverrides returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DescribeComputeConfigResult) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DescribeComputeConfigResult) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DescribeComputeConfigResult) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DescribeComputeConfigResult) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeComputeConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeComputeConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeComputeConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeComputeConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeComputeConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeComputeConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetInfraConfigIDs

`func (o *DescribeComputeConfigResult) GetInfraConfigIDs() []string`

GetInfraConfigIDs returns the InfraConfigIDs field if non-nil, zero value otherwise.

### GetInfraConfigIDsOk

`func (o *DescribeComputeConfigResult) GetInfraConfigIDsOk() (*[]string, bool)`

GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigIDs

`func (o *DescribeComputeConfigResult) SetInfraConfigIDs(v []string)`

SetInfraConfigIDs sets InfraConfigIDs field to given value.

### HasInfraConfigIDs

`func (o *DescribeComputeConfigResult) HasInfraConfigIDs() bool`

HasInfraConfigIDs returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *DescribeComputeConfigResult) GetInstanceTypes() map[string][]string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *DescribeComputeConfigResult) GetInstanceTypesOk() (*map[string][]string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *DescribeComputeConfigResult) SetInstanceTypes(v map[string][]string)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *DescribeComputeConfigResult) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetName

`func (o *DescribeComputeConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeComputeConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeComputeConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetReplicaCount

`func (o *DescribeComputeConfigResult) GetReplicaCount() string`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *DescribeComputeConfigResult) GetReplicaCountOk() (*string, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *DescribeComputeConfigResult) SetReplicaCount(v string)`

SetReplicaCount sets ReplicaCount field to given value.


### GetResources

`func (o *DescribeComputeConfigResult) GetResources() ResourceSpec`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DescribeComputeConfigResult) GetResourcesOk() (*ResourceSpec, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DescribeComputeConfigResult) SetResources(v ResourceSpec)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DescribeComputeConfigResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRootVolumeSizeGi

`func (o *DescribeComputeConfigResult) GetRootVolumeSizeGi() int64`

GetRootVolumeSizeGi returns the RootVolumeSizeGi field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiOk

`func (o *DescribeComputeConfigResult) GetRootVolumeSizeGiOk() (*int64, bool)`

GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGi

`func (o *DescribeComputeConfigResult) SetRootVolumeSizeGi(v int64)`

SetRootVolumeSizeGi sets RootVolumeSizeGi field to given value.

### HasRootVolumeSizeGi

`func (o *DescribeComputeConfigResult) HasRootVolumeSizeGi() bool`

HasRootVolumeSizeGi returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeComputeConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeComputeConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeComputeConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetWarmPoolConfiguration

`func (o *DescribeComputeConfigResult) GetWarmPoolConfiguration() WarmPoolConfiguration`

GetWarmPoolConfiguration returns the WarmPoolConfiguration field if non-nil, zero value otherwise.

### GetWarmPoolConfigurationOk

`func (o *DescribeComputeConfigResult) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool)`

GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolConfiguration

`func (o *DescribeComputeConfigResult) SetWarmPoolConfiguration(v WarmPoolConfiguration)`

SetWarmPoolConfiguration sets WarmPoolConfiguration field to given value.

### HasWarmPoolConfiguration

`func (o *DescribeComputeConfigResult) HasWarmPoolConfiguration() bool`

HasWarmPoolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


