# CreateComputeConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalingPolicy** | Pointer to [**AutoscalingPolicy**](AutoscalingPolicy.md) |  | [optional] 
**CpuArchitecture** | Pointer to **string** | Processor architecture | [optional] 
**Description** | **string** | Description of the compute config | 
**Name** | **string** | Name of the compute config | 
**ReplicaCount** | Pointer to **string** | Number of replicas to provision for this logical pool of nodes per instance of the resource | [optional] 
**Resources** | Pointer to [**ResourceSpec**](ResourceSpec.md) |  | [optional] 
**RootVolumeSizeGi** | Pointer to **int64** | Size of the root volume in Gi | [optional] 
**WarmPoolConfiguration** | Pointer to [**WarmPoolConfiguration**](WarmPoolConfiguration.md) |  | [optional] 

## Methods

### NewCreateComputeConfigRequestBody

`func NewCreateComputeConfigRequestBody(description string, name string, ) *CreateComputeConfigRequestBody`

NewCreateComputeConfigRequestBody instantiates a new CreateComputeConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComputeConfigRequestBodyWithDefaults

`func NewCreateComputeConfigRequestBodyWithDefaults() *CreateComputeConfigRequestBody`

NewCreateComputeConfigRequestBodyWithDefaults instantiates a new CreateComputeConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalingPolicy

`func (o *CreateComputeConfigRequestBody) GetAutoscalingPolicy() AutoscalingPolicy`

GetAutoscalingPolicy returns the AutoscalingPolicy field if non-nil, zero value otherwise.

### GetAutoscalingPolicyOk

`func (o *CreateComputeConfigRequestBody) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool)`

GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingPolicy

`func (o *CreateComputeConfigRequestBody) SetAutoscalingPolicy(v AutoscalingPolicy)`

SetAutoscalingPolicy sets AutoscalingPolicy field to given value.

### HasAutoscalingPolicy

`func (o *CreateComputeConfigRequestBody) HasAutoscalingPolicy() bool`

HasAutoscalingPolicy returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *CreateComputeConfigRequestBody) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *CreateComputeConfigRequestBody) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *CreateComputeConfigRequestBody) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *CreateComputeConfigRequestBody) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetDescription

`func (o *CreateComputeConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateComputeConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateComputeConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateComputeConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateComputeConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateComputeConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetReplicaCount

`func (o *CreateComputeConfigRequestBody) GetReplicaCount() string`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *CreateComputeConfigRequestBody) GetReplicaCountOk() (*string, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *CreateComputeConfigRequestBody) SetReplicaCount(v string)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *CreateComputeConfigRequestBody) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetResources

`func (o *CreateComputeConfigRequestBody) GetResources() ResourceSpec`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateComputeConfigRequestBody) GetResourcesOk() (*ResourceSpec, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateComputeConfigRequestBody) SetResources(v ResourceSpec)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CreateComputeConfigRequestBody) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRootVolumeSizeGi

`func (o *CreateComputeConfigRequestBody) GetRootVolumeSizeGi() int64`

GetRootVolumeSizeGi returns the RootVolumeSizeGi field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiOk

`func (o *CreateComputeConfigRequestBody) GetRootVolumeSizeGiOk() (*int64, bool)`

GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGi

`func (o *CreateComputeConfigRequestBody) SetRootVolumeSizeGi(v int64)`

SetRootVolumeSizeGi sets RootVolumeSizeGi field to given value.

### HasRootVolumeSizeGi

`func (o *CreateComputeConfigRequestBody) HasRootVolumeSizeGi() bool`

HasRootVolumeSizeGi returns a boolean if a field has been set.

### GetWarmPoolConfiguration

`func (o *CreateComputeConfigRequestBody) GetWarmPoolConfiguration() WarmPoolConfiguration`

GetWarmPoolConfiguration returns the WarmPoolConfiguration field if non-nil, zero value otherwise.

### GetWarmPoolConfigurationOk

`func (o *CreateComputeConfigRequestBody) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool)`

GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolConfiguration

`func (o *CreateComputeConfigRequestBody) SetWarmPoolConfiguration(v WarmPoolConfiguration)`

SetWarmPoolConfiguration sets WarmPoolConfiguration field to given value.

### HasWarmPoolConfiguration

`func (o *CreateComputeConfigRequestBody) HasWarmPoolConfiguration() bool`

HasWarmPoolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


