# InfraRollConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfraRolloutStrategy** | Pointer to **string** | Rollout strategy to use for infra config updates | [optional] 
**MaximumNumberOfResourceInstancesPerBatch** | Pointer to **int64** | Maximum number of resource instances to update in a single batch | [optional] [default to 1]
**MaximumUnavailableReplicas** | Pointer to **int64** | Maximum number of replicas of this resource that can be unavailable before halting the rollout | [optional] [default to 1]

## Methods

### NewInfraRollConfiguration

`func NewInfraRollConfiguration() *InfraRollConfiguration`

NewInfraRollConfiguration instantiates a new InfraRollConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraRollConfigurationWithDefaults

`func NewInfraRollConfigurationWithDefaults() *InfraRollConfiguration`

NewInfraRollConfigurationWithDefaults instantiates a new InfraRollConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfraRolloutStrategy

`func (o *InfraRollConfiguration) GetInfraRolloutStrategy() string`

GetInfraRolloutStrategy returns the InfraRolloutStrategy field if non-nil, zero value otherwise.

### GetInfraRolloutStrategyOk

`func (o *InfraRollConfiguration) GetInfraRolloutStrategyOk() (*string, bool)`

GetInfraRolloutStrategyOk returns a tuple with the InfraRolloutStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRolloutStrategy

`func (o *InfraRollConfiguration) SetInfraRolloutStrategy(v string)`

SetInfraRolloutStrategy sets InfraRolloutStrategy field to given value.

### HasInfraRolloutStrategy

`func (o *InfraRollConfiguration) HasInfraRolloutStrategy() bool`

HasInfraRolloutStrategy returns a boolean if a field has been set.

### GetMaximumNumberOfResourceInstancesPerBatch

`func (o *InfraRollConfiguration) GetMaximumNumberOfResourceInstancesPerBatch() int64`

GetMaximumNumberOfResourceInstancesPerBatch returns the MaximumNumberOfResourceInstancesPerBatch field if non-nil, zero value otherwise.

### GetMaximumNumberOfResourceInstancesPerBatchOk

`func (o *InfraRollConfiguration) GetMaximumNumberOfResourceInstancesPerBatchOk() (*int64, bool)`

GetMaximumNumberOfResourceInstancesPerBatchOk returns a tuple with the MaximumNumberOfResourceInstancesPerBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumberOfResourceInstancesPerBatch

`func (o *InfraRollConfiguration) SetMaximumNumberOfResourceInstancesPerBatch(v int64)`

SetMaximumNumberOfResourceInstancesPerBatch sets MaximumNumberOfResourceInstancesPerBatch field to given value.

### HasMaximumNumberOfResourceInstancesPerBatch

`func (o *InfraRollConfiguration) HasMaximumNumberOfResourceInstancesPerBatch() bool`

HasMaximumNumberOfResourceInstancesPerBatch returns a boolean if a field has been set.

### GetMaximumUnavailableReplicas

`func (o *InfraRollConfiguration) GetMaximumUnavailableReplicas() int64`

GetMaximumUnavailableReplicas returns the MaximumUnavailableReplicas field if non-nil, zero value otherwise.

### GetMaximumUnavailableReplicasOk

`func (o *InfraRollConfiguration) GetMaximumUnavailableReplicasOk() (*int64, bool)`

GetMaximumUnavailableReplicasOk returns a tuple with the MaximumUnavailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumUnavailableReplicas

`func (o *InfraRollConfiguration) SetMaximumUnavailableReplicas(v int64)`

SetMaximumUnavailableReplicas sets MaximumUnavailableReplicas field to given value.

### HasMaximumUnavailableReplicas

`func (o *InfraRollConfiguration) HasMaximumUnavailableReplicas() bool`

HasMaximumUnavailableReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


