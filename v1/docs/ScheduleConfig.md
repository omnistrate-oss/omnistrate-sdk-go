# ScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron expression for advanced scheduling (e.g., &#39;0 2 * * *&#39; for daily at 2 AM). Takes precedence over simpleInterval if both are provided. | [optional] 
**SimpleInterval** | Pointer to **string** | Simple interval scheduling using human-readable format (e.g., &#39;5m&#39;, &#39;1h&#39;, &#39;30m&#39;, &#39;2h&#39;). Supports: s(seconds), m(minutes), h(hours), d(days) | [optional] 
**Timezone** | Pointer to **string** | Timezone for cron scheduling (IANA Time Zone format, e.g., &#39;UTC&#39;, &#39;America/New_York&#39;). Only applies to cronExpression. | [optional] [default to "UTC"]

## Methods

### NewScheduleConfig

`func NewScheduleConfig() *ScheduleConfig`

NewScheduleConfig instantiates a new ScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleConfigWithDefaults

`func NewScheduleConfigWithDefaults() *ScheduleConfig`

NewScheduleConfigWithDefaults instantiates a new ScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *ScheduleConfig) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleConfig) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleConfig) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ScheduleConfig) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetSimpleInterval

`func (o *ScheduleConfig) GetSimpleInterval() string`

GetSimpleInterval returns the SimpleInterval field if non-nil, zero value otherwise.

### GetSimpleIntervalOk

`func (o *ScheduleConfig) GetSimpleIntervalOk() (*string, bool)`

GetSimpleIntervalOk returns a tuple with the SimpleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleInterval

`func (o *ScheduleConfig) SetSimpleInterval(v string)`

SetSimpleInterval sets SimpleInterval field to given value.

### HasSimpleInterval

`func (o *ScheduleConfig) HasSimpleInterval() bool`

HasSimpleInterval returns a boolean if a field has been set.

### GetTimezone

`func (o *ScheduleConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ScheduleConfig) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


