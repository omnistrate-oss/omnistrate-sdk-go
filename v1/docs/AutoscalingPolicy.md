# AutoscalingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleMinutesBeforeScalingDown** | Pointer to **int64** | Minutes before scaling down the compute nodes in idle state | [optional] [default to 5]
**IdleThreshold** | Pointer to **int64** | Metric value below threshold will be considered to be idle | [optional] [default to 20]
**ManagementPolicy** | Pointer to **string** | The autoscaling policy type | [optional] [default to "managed"]
**MaxReplicas** | **string** | Maximum number of compute nodes to provision | 
**MinReplicas** | **string** | Minimum number of compute nodes to provision | 
**OverUtilizedMinutesBeforeScalingUp** | Pointer to **int64** | Minutes before scaling up the compute nodes in overUtilized state | [optional] [default to 5]
**OverUtilizedThreshold** | Pointer to **int64** | Metric value beyond threshold will be considered to be overUtilized | [optional] [default to 80]
**ScalingMetric** | Pointer to [**AutoScalingMetricSpec**](AutoScalingMetricSpec.md) |  | [optional] 

## Methods

### NewAutoscalingPolicy

`func NewAutoscalingPolicy(maxReplicas string, minReplicas string, ) *AutoscalingPolicy`

NewAutoscalingPolicy instantiates a new AutoscalingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingPolicyWithDefaults

`func NewAutoscalingPolicyWithDefaults() *AutoscalingPolicy`

NewAutoscalingPolicyWithDefaults instantiates a new AutoscalingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleMinutesBeforeScalingDown

`func (o *AutoscalingPolicy) GetIdleMinutesBeforeScalingDown() int64`

GetIdleMinutesBeforeScalingDown returns the IdleMinutesBeforeScalingDown field if non-nil, zero value otherwise.

### GetIdleMinutesBeforeScalingDownOk

`func (o *AutoscalingPolicy) GetIdleMinutesBeforeScalingDownOk() (*int64, bool)`

GetIdleMinutesBeforeScalingDownOk returns a tuple with the IdleMinutesBeforeScalingDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleMinutesBeforeScalingDown

`func (o *AutoscalingPolicy) SetIdleMinutesBeforeScalingDown(v int64)`

SetIdleMinutesBeforeScalingDown sets IdleMinutesBeforeScalingDown field to given value.

### HasIdleMinutesBeforeScalingDown

`func (o *AutoscalingPolicy) HasIdleMinutesBeforeScalingDown() bool`

HasIdleMinutesBeforeScalingDown returns a boolean if a field has been set.

### GetIdleThreshold

`func (o *AutoscalingPolicy) GetIdleThreshold() int64`

GetIdleThreshold returns the IdleThreshold field if non-nil, zero value otherwise.

### GetIdleThresholdOk

`func (o *AutoscalingPolicy) GetIdleThresholdOk() (*int64, bool)`

GetIdleThresholdOk returns a tuple with the IdleThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleThreshold

`func (o *AutoscalingPolicy) SetIdleThreshold(v int64)`

SetIdleThreshold sets IdleThreshold field to given value.

### HasIdleThreshold

`func (o *AutoscalingPolicy) HasIdleThreshold() bool`

HasIdleThreshold returns a boolean if a field has been set.

### GetManagementPolicy

`func (o *AutoscalingPolicy) GetManagementPolicy() string`

GetManagementPolicy returns the ManagementPolicy field if non-nil, zero value otherwise.

### GetManagementPolicyOk

`func (o *AutoscalingPolicy) GetManagementPolicyOk() (*string, bool)`

GetManagementPolicyOk returns a tuple with the ManagementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPolicy

`func (o *AutoscalingPolicy) SetManagementPolicy(v string)`

SetManagementPolicy sets ManagementPolicy field to given value.

### HasManagementPolicy

`func (o *AutoscalingPolicy) HasManagementPolicy() bool`

HasManagementPolicy returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *AutoscalingPolicy) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *AutoscalingPolicy) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *AutoscalingPolicy) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMinReplicas

`func (o *AutoscalingPolicy) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *AutoscalingPolicy) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *AutoscalingPolicy) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.


### GetOverUtilizedMinutesBeforeScalingUp

`func (o *AutoscalingPolicy) GetOverUtilizedMinutesBeforeScalingUp() int64`

GetOverUtilizedMinutesBeforeScalingUp returns the OverUtilizedMinutesBeforeScalingUp field if non-nil, zero value otherwise.

### GetOverUtilizedMinutesBeforeScalingUpOk

`func (o *AutoscalingPolicy) GetOverUtilizedMinutesBeforeScalingUpOk() (*int64, bool)`

GetOverUtilizedMinutesBeforeScalingUpOk returns a tuple with the OverUtilizedMinutesBeforeScalingUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverUtilizedMinutesBeforeScalingUp

`func (o *AutoscalingPolicy) SetOverUtilizedMinutesBeforeScalingUp(v int64)`

SetOverUtilizedMinutesBeforeScalingUp sets OverUtilizedMinutesBeforeScalingUp field to given value.

### HasOverUtilizedMinutesBeforeScalingUp

`func (o *AutoscalingPolicy) HasOverUtilizedMinutesBeforeScalingUp() bool`

HasOverUtilizedMinutesBeforeScalingUp returns a boolean if a field has been set.

### GetOverUtilizedThreshold

`func (o *AutoscalingPolicy) GetOverUtilizedThreshold() int64`

GetOverUtilizedThreshold returns the OverUtilizedThreshold field if non-nil, zero value otherwise.

### GetOverUtilizedThresholdOk

`func (o *AutoscalingPolicy) GetOverUtilizedThresholdOk() (*int64, bool)`

GetOverUtilizedThresholdOk returns a tuple with the OverUtilizedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverUtilizedThreshold

`func (o *AutoscalingPolicy) SetOverUtilizedThreshold(v int64)`

SetOverUtilizedThreshold sets OverUtilizedThreshold field to given value.

### HasOverUtilizedThreshold

`func (o *AutoscalingPolicy) HasOverUtilizedThreshold() bool`

HasOverUtilizedThreshold returns a boolean if a field has been set.

### GetScalingMetric

`func (o *AutoscalingPolicy) GetScalingMetric() AutoScalingMetricSpec`

GetScalingMetric returns the ScalingMetric field if non-nil, zero value otherwise.

### GetScalingMetricOk

`func (o *AutoscalingPolicy) GetScalingMetricOk() (*AutoScalingMetricSpec, bool)`

GetScalingMetricOk returns a tuple with the ScalingMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingMetric

`func (o *AutoscalingPolicy) SetScalingMetric(v AutoScalingMetricSpec)`

SetScalingMetric sets ScalingMetric field to given value.

### HasScalingMetric

`func (o *AutoscalingPolicy) HasScalingMetric() bool`

HasScalingMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


