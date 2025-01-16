# UpdateComputeConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalingPolicy** | Pointer to [**AutoscalingPolicy**](AutoscalingPolicy.md) |  | [optional] 
**CpuArchitecture** | Pointer to **string** | Processor architecture | [optional] 
**Description** | Pointer to **string** | Description of the compute config | [optional] 
**Name** | Pointer to **string** | Name of the compute config | [optional] 
**ReplicaCount** | Pointer to **string** | Number of replicas to provision for this logical pool of nodes per instance of the resource | [optional] 
**Resources** | Pointer to [**ResourceSpec**](ResourceSpec.md) |  | [optional] 
**RootVolumeSizeGi** | Pointer to **int64** | Size of the root volume in Gi | [optional] 
**WarmPoolConfiguration** | Pointer to [**WarmPoolConfiguration**](WarmPoolConfiguration.md) |  | [optional] 

## Methods

### NewUpdateComputeConfigRequest2

`func NewUpdateComputeConfigRequest2() *UpdateComputeConfigRequest2`

NewUpdateComputeConfigRequest2 instantiates a new UpdateComputeConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateComputeConfigRequest2WithDefaults

`func NewUpdateComputeConfigRequest2WithDefaults() *UpdateComputeConfigRequest2`

NewUpdateComputeConfigRequest2WithDefaults instantiates a new UpdateComputeConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalingPolicy

`func (o *UpdateComputeConfigRequest2) GetAutoscalingPolicy() AutoscalingPolicy`

GetAutoscalingPolicy returns the AutoscalingPolicy field if non-nil, zero value otherwise.

### GetAutoscalingPolicyOk

`func (o *UpdateComputeConfigRequest2) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool)`

GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingPolicy

`func (o *UpdateComputeConfigRequest2) SetAutoscalingPolicy(v AutoscalingPolicy)`

SetAutoscalingPolicy sets AutoscalingPolicy field to given value.

### HasAutoscalingPolicy

`func (o *UpdateComputeConfigRequest2) HasAutoscalingPolicy() bool`

HasAutoscalingPolicy returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *UpdateComputeConfigRequest2) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *UpdateComputeConfigRequest2) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *UpdateComputeConfigRequest2) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *UpdateComputeConfigRequest2) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateComputeConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateComputeConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateComputeConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateComputeConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateComputeConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateComputeConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateComputeConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateComputeConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicaCount

`func (o *UpdateComputeConfigRequest2) GetReplicaCount() string`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *UpdateComputeConfigRequest2) GetReplicaCountOk() (*string, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *UpdateComputeConfigRequest2) SetReplicaCount(v string)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *UpdateComputeConfigRequest2) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetResources

`func (o *UpdateComputeConfigRequest2) GetResources() ResourceSpec`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UpdateComputeConfigRequest2) GetResourcesOk() (*ResourceSpec, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UpdateComputeConfigRequest2) SetResources(v ResourceSpec)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UpdateComputeConfigRequest2) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRootVolumeSizeGi

`func (o *UpdateComputeConfigRequest2) GetRootVolumeSizeGi() int64`

GetRootVolumeSizeGi returns the RootVolumeSizeGi field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiOk

`func (o *UpdateComputeConfigRequest2) GetRootVolumeSizeGiOk() (*int64, bool)`

GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGi

`func (o *UpdateComputeConfigRequest2) SetRootVolumeSizeGi(v int64)`

SetRootVolumeSizeGi sets RootVolumeSizeGi field to given value.

### HasRootVolumeSizeGi

`func (o *UpdateComputeConfigRequest2) HasRootVolumeSizeGi() bool`

HasRootVolumeSizeGi returns a boolean if a field has been set.

### GetWarmPoolConfiguration

`func (o *UpdateComputeConfigRequest2) GetWarmPoolConfiguration() WarmPoolConfiguration`

GetWarmPoolConfiguration returns the WarmPoolConfiguration field if non-nil, zero value otherwise.

### GetWarmPoolConfigurationOk

`func (o *UpdateComputeConfigRequest2) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool)`

GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolConfiguration

`func (o *UpdateComputeConfigRequest2) SetWarmPoolConfiguration(v WarmPoolConfiguration)`

SetWarmPoolConfiguration sets WarmPoolConfiguration field to given value.

### HasWarmPoolConfiguration

`func (o *UpdateComputeConfigRequest2) HasWarmPoolConfiguration() bool`

HasWarmPoolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


