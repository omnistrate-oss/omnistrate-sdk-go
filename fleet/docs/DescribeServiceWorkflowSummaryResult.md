# DescribeServiceWorkflowSummaryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWorkflowCount** | Pointer to **int64** | Number of active workflows for the given service in the past 24 hours. | [optional] 
**CompletedWorkflowCount** | Pointer to **int64** | Number of completed workflows for the given service in the past 24 hours. | [optional] 
**FailedWorkflowCount** | Pointer to **int64** | Number of failed workflows for the given service in the past 24 hours. | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewDescribeServiceWorkflowSummaryResult

`func NewDescribeServiceWorkflowSummaryResult(environmentId string, serviceId string, ) *DescribeServiceWorkflowSummaryResult`

NewDescribeServiceWorkflowSummaryResult instantiates a new DescribeServiceWorkflowSummaryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceWorkflowSummaryResultWithDefaults

`func NewDescribeServiceWorkflowSummaryResultWithDefaults() *DescribeServiceWorkflowSummaryResult`

NewDescribeServiceWorkflowSummaryResultWithDefaults instantiates a new DescribeServiceWorkflowSummaryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) GetActiveWorkflowCount() int64`

GetActiveWorkflowCount returns the ActiveWorkflowCount field if non-nil, zero value otherwise.

### GetActiveWorkflowCountOk

`func (o *DescribeServiceWorkflowSummaryResult) GetActiveWorkflowCountOk() (*int64, bool)`

GetActiveWorkflowCountOk returns a tuple with the ActiveWorkflowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) SetActiveWorkflowCount(v int64)`

SetActiveWorkflowCount sets ActiveWorkflowCount field to given value.

### HasActiveWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) HasActiveWorkflowCount() bool`

HasActiveWorkflowCount returns a boolean if a field has been set.

### GetCompletedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) GetCompletedWorkflowCount() int64`

GetCompletedWorkflowCount returns the CompletedWorkflowCount field if non-nil, zero value otherwise.

### GetCompletedWorkflowCountOk

`func (o *DescribeServiceWorkflowSummaryResult) GetCompletedWorkflowCountOk() (*int64, bool)`

GetCompletedWorkflowCountOk returns a tuple with the CompletedWorkflowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) SetCompletedWorkflowCount(v int64)`

SetCompletedWorkflowCount sets CompletedWorkflowCount field to given value.

### HasCompletedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) HasCompletedWorkflowCount() bool`

HasCompletedWorkflowCount returns a boolean if a field has been set.

### GetFailedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) GetFailedWorkflowCount() int64`

GetFailedWorkflowCount returns the FailedWorkflowCount field if non-nil, zero value otherwise.

### GetFailedWorkflowCountOk

`func (o *DescribeServiceWorkflowSummaryResult) GetFailedWorkflowCountOk() (*int64, bool)`

GetFailedWorkflowCountOk returns a tuple with the FailedWorkflowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) SetFailedWorkflowCount(v int64)`

SetFailedWorkflowCount sets FailedWorkflowCount field to given value.

### HasFailedWorkflowCount

`func (o *DescribeServiceWorkflowSummaryResult) HasFailedWorkflowCount() bool`

HasFailedWorkflowCount returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DescribeServiceWorkflowSummaryResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeServiceWorkflowSummaryResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeServiceWorkflowSummaryResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServiceId

`func (o *DescribeServiceWorkflowSummaryResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeServiceWorkflowSummaryResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeServiceWorkflowSummaryResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


