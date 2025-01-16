# CreateComputeConfigRequest

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
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**WarmPoolConfiguration** | Pointer to [**WarmPoolConfiguration**](WarmPoolConfiguration.md) |  | [optional] 

## Methods

### NewCreateComputeConfigRequest

`func NewCreateComputeConfigRequest(description string, name string, serviceId string, token string, ) *CreateComputeConfigRequest`

NewCreateComputeConfigRequest instantiates a new CreateComputeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComputeConfigRequestWithDefaults

`func NewCreateComputeConfigRequestWithDefaults() *CreateComputeConfigRequest`

NewCreateComputeConfigRequestWithDefaults instantiates a new CreateComputeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalingPolicy

`func (o *CreateComputeConfigRequest) GetAutoscalingPolicy() AutoscalingPolicy`

GetAutoscalingPolicy returns the AutoscalingPolicy field if non-nil, zero value otherwise.

### GetAutoscalingPolicyOk

`func (o *CreateComputeConfigRequest) GetAutoscalingPolicyOk() (*AutoscalingPolicy, bool)`

GetAutoscalingPolicyOk returns a tuple with the AutoscalingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingPolicy

`func (o *CreateComputeConfigRequest) SetAutoscalingPolicy(v AutoscalingPolicy)`

SetAutoscalingPolicy sets AutoscalingPolicy field to given value.

### HasAutoscalingPolicy

`func (o *CreateComputeConfigRequest) HasAutoscalingPolicy() bool`

HasAutoscalingPolicy returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *CreateComputeConfigRequest) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *CreateComputeConfigRequest) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *CreateComputeConfigRequest) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *CreateComputeConfigRequest) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetDescription

`func (o *CreateComputeConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateComputeConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateComputeConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateComputeConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateComputeConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateComputeConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetReplicaCount

`func (o *CreateComputeConfigRequest) GetReplicaCount() string`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *CreateComputeConfigRequest) GetReplicaCountOk() (*string, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *CreateComputeConfigRequest) SetReplicaCount(v string)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *CreateComputeConfigRequest) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetResources

`func (o *CreateComputeConfigRequest) GetResources() ResourceSpec`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateComputeConfigRequest) GetResourcesOk() (*ResourceSpec, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateComputeConfigRequest) SetResources(v ResourceSpec)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CreateComputeConfigRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRootVolumeSizeGi

`func (o *CreateComputeConfigRequest) GetRootVolumeSizeGi() int64`

GetRootVolumeSizeGi returns the RootVolumeSizeGi field if non-nil, zero value otherwise.

### GetRootVolumeSizeGiOk

`func (o *CreateComputeConfigRequest) GetRootVolumeSizeGiOk() (*int64, bool)`

GetRootVolumeSizeGiOk returns a tuple with the RootVolumeSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSizeGi

`func (o *CreateComputeConfigRequest) SetRootVolumeSizeGi(v int64)`

SetRootVolumeSizeGi sets RootVolumeSizeGi field to given value.

### HasRootVolumeSizeGi

`func (o *CreateComputeConfigRequest) HasRootVolumeSizeGi() bool`

HasRootVolumeSizeGi returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateComputeConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateComputeConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateComputeConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreateComputeConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateComputeConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateComputeConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWarmPoolConfiguration

`func (o *CreateComputeConfigRequest) GetWarmPoolConfiguration() WarmPoolConfiguration`

GetWarmPoolConfiguration returns the WarmPoolConfiguration field if non-nil, zero value otherwise.

### GetWarmPoolConfigurationOk

`func (o *CreateComputeConfigRequest) GetWarmPoolConfigurationOk() (*WarmPoolConfiguration, bool)`

GetWarmPoolConfigurationOk returns a tuple with the WarmPoolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolConfiguration

`func (o *CreateComputeConfigRequest) SetWarmPoolConfiguration(v WarmPoolConfiguration)`

SetWarmPoolConfiguration sets WarmPoolConfiguration field to given value.

### HasWarmPoolConfiguration

`func (o *CreateComputeConfigRequest) HasWarmPoolConfiguration() bool`

HasWarmPoolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


