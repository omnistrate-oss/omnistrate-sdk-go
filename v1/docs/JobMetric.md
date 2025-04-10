# JobMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]interface{}** | Additional metadata about the job | [optional] 
**EndTime** | Pointer to **string** | When the job completed | [optional] 
**MetricType** | **string** | Type of job metric being tracked | 
**StartTime** | **string** | When the job started running | 
**Value** | **float64** | Value of the metric | 

## Methods

### NewJobMetric

`func NewJobMetric(metricType string, startTime string, value float64, ) *JobMetric`

NewJobMetric instantiates a new JobMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobMetricWithDefaults

`func NewJobMetricWithDefaults() *JobMetric`

NewJobMetricWithDefaults instantiates a new JobMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *JobMetric) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *JobMetric) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *JobMetric) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *JobMetric) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetEndTime

`func (o *JobMetric) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobMetric) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobMetric) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobMetric) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMetricType

`func (o *JobMetric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *JobMetric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *JobMetric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetStartTime

`func (o *JobMetric) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobMetric) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobMetric) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetValue

`func (o *JobMetric) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JobMetric) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JobMetric) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


