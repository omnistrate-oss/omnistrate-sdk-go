# JobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **int64** | The duration in seconds relative to the startTime that the job may be active before the system tries to terminate it | [optional] [default to 3600]
**BackoffLimit** | Pointer to **int32** | The number of retries before marking the job as failed | [optional] [default to 5]

## Methods

### NewJobConfig

`func NewJobConfig() *JobConfig`

NewJobConfig instantiates a new JobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigWithDefaults

`func NewJobConfigWithDefaults() *JobConfig`

NewJobConfigWithDefaults instantiates a new JobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *JobConfig) GetActiveDeadlineSeconds() int64`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *JobConfig) GetActiveDeadlineSecondsOk() (*int64, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *JobConfig) SetActiveDeadlineSeconds(v int64)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *JobConfig) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetBackoffLimit

`func (o *JobConfig) GetBackoffLimit() int32`

GetBackoffLimit returns the BackoffLimit field if non-nil, zero value otherwise.

### GetBackoffLimitOk

`func (o *JobConfig) GetBackoffLimitOk() (*int32, bool)`

GetBackoffLimitOk returns a tuple with the BackoffLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffLimit

`func (o *JobConfig) SetBackoffLimit(v int32)`

SetBackoffLimit sets BackoffLimit field to given value.

### HasBackoffLimit

`func (o *JobConfig) HasBackoffLimit() bool`

HasBackoffLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


