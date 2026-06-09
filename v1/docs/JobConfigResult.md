# JobConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | **int64** | The duration in seconds relative to the startTime that the job may be active before the system tries to terminate it | 
**BackoffLimit** | **int32** | The number of retries before marking the job as failed | 
**ScheduleConfig** | Pointer to [**ScheduleConfig**](ScheduleConfig.md) |  | [optional] 

## Methods

### NewJobConfigResult

`func NewJobConfigResult(activeDeadlineSeconds int64, backoffLimit int32, ) *JobConfigResult`

NewJobConfigResult instantiates a new JobConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigResultWithDefaults

`func NewJobConfigResultWithDefaults() *JobConfigResult`

NewJobConfigResultWithDefaults instantiates a new JobConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *JobConfigResult) GetActiveDeadlineSeconds() int64`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *JobConfigResult) GetActiveDeadlineSecondsOk() (*int64, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *JobConfigResult) SetActiveDeadlineSeconds(v int64)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.


### GetBackoffLimit

`func (o *JobConfigResult) GetBackoffLimit() int32`

GetBackoffLimit returns the BackoffLimit field if non-nil, zero value otherwise.

### GetBackoffLimitOk

`func (o *JobConfigResult) GetBackoffLimitOk() (*int32, bool)`

GetBackoffLimitOk returns a tuple with the BackoffLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffLimit

`func (o *JobConfigResult) SetBackoffLimit(v int32)`

SetBackoffLimit sets BackoffLimit field to given value.


### GetScheduleConfig

`func (o *JobConfigResult) GetScheduleConfig() ScheduleConfig`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *JobConfigResult) GetScheduleConfigOk() (*ScheduleConfig, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *JobConfigResult) SetScheduleConfig(v ScheduleConfig)`

SetScheduleConfig sets ScheduleConfig field to given value.

### HasScheduleConfig

`func (o *JobConfigResult) HasScheduleConfig() bool`

HasScheduleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


